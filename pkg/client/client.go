package client

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/viper"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/uhttp"
)

type client struct {
	httpClient *http.Client
	clientName string
	tokenHost  string
	apiClient  *c1api.APIClient
	baseURL    *url.URL
}

func StringFromPtr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func float32Ptr(i int) *float32 {
	f := float32(i)
	return &f
}

type C1Client interface {
	BaseURL() string

	AuthIntrospect(ctx context.Context) (*c1api.C1ApiAuthV1IntrospectResponse, error)
	GetUser(ctx context.Context, userID string) (*c1api.C1ApiUserV1User, error)
	GetEntitlement(ctx context.Context, appID string, entitlementID string) (*c1api.C1ApiAppV1AppEntitlement, error)
	SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) ([]*EntitlementWithBindings, error)
	ExpandEntitlements(ctx context.Context, in []*EntitlementWithBindings) (*Expander, error)
	GetResource(ctx context.Context, appID string, resourceID string, resourceTypeID string) (*c1api.C1ApiAppV1AppResource, error)
	GetResourceType(ctx context.Context, appID string, resourceTypeID string) (*c1api.C1ApiAppV1AppResourceType, error)
	GetApp(ctx context.Context, appID string) (*c1api.C1ApiAppV1App, error)
	GetTask(ctx context.Context, taskId string) (*c1api.C1ApiTaskV1TaskServiceGetResponse, error)
	CreateGrantTask(
		ctx context.Context,
		appId string,
		appEntitlementId string,
		identityUserId string,
		justification string,
		duration string,
	) (*c1api.C1ApiTaskV1TaskServiceCreateGrantResponse, error)
	CreateRevokeTask(
		ctx context.Context,
		appId string,
		appEntitlementId string,
		identityUserId string,
		justification string,
	) (*c1api.C1ApiTaskV1TaskServiceCreateRevokeResponse, error)
	GetGrantsForIdentity(ctx context.Context, appID string, appEntitlementID string, appUserID string) ([]c1api.C1ApiAppV1AppEntitlementUserBinding, error)

	SearchTasks(ctx context.Context, taskFilter c1api.C1ApiTaskV1TaskSearchRequest) (*c1api.C1ApiTaskV1TaskSearchResponse, error)
	CommentOnTask(ctx context.Context, taskID string, comment string) (*c1api.C1ApiTaskV1TaskActionsServiceCommentResponse, error)
}

func (c *client) BaseURL() string {
	return c.baseURL.String()
}

func New(
	ctx context.Context,
	clientId string,
	clientSecret string,
	v *viper.Viper,
) (C1Client, error) {
	tokenSrc, clientName, tokenHost, err := NewC1TokenSource(ctx, clientId, clientSecret)
	if err != nil {
		return nil, err
	}

	uclient, err := uhttp.NewClient(
		ctx,
		uhttp.WithTokenSource(tokenSrc),
		uhttp.WithDebug(v.GetBool("debug")),
	)
	if err != nil {
		return nil, err
	}
	c := &client{
		tokenHost:  tokenHost,
		clientName: clientName,
		httpClient: uclient,
	}

	apiCfg := c1api.NewConfiguration()
	apiCfg.HTTPClient = uclient

	var apiHostname string
	// If the API host is set in the environment, use that instead of the default
	// HACK(jirwin): Instead of using the generated client's server address, use the hostname from the token.
	if apiHost, ok := os.LookupEnv("CONE_API_ENDPOINT"); ok {
		apiHostname = apiHost
	} else {
		apiHostname = c.tokenHost
	}
	apiURL := url.URL{
		Scheme: "https",
		Host:   apiHostname,
	}
	apiCfg.Servers[0].URL = apiURL.String()
	c.apiClient = c1api.NewAPIClient(apiCfg)
	c.baseURL = &apiURL

	return c, nil
}

// The c1api client uses the context to set various configuration options. Do that here.
func (c *client) GetContext(ctx context.Context) context.Context {
	// If the API host is set in the environment, we don't need to populate any server variables
	if _, ok := os.LookupEnv("CONE_API_ENDPOINT"); !ok {
		return ctx
	}

	// TODO(jirwin): if we choose to use this, we will need to parse the tenant name out of the token, and set it as `tenantHost` here.
	// serverVars := map[string]string{
	// 	"tenantHost": c.tokenHost,
	// }
	// return context.WithValue(ctx, c1api.ContextServerVariables, c.clientName)
	return ctx
}
