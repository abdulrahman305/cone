// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppOwnersAddRequest struct {
	AddAppOwnerRequest *shared.AddAppOwnerRequest `request:"mediaType=application/json"`
	AppID              string                     `pathParam:"style=simple,explode=false,name=app_id"`
	UserID             string                     `pathParam:"style=simple,explode=false,name=user_id"`
}

func (o *C1APIAppV1AppOwnersAddRequest) GetAddAppOwnerRequest() *shared.AddAppOwnerRequest {
	if o == nil {
		return nil
	}
	return o.AddAppOwnerRequest
}

func (o *C1APIAppV1AppOwnersAddRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppOwnersAddRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type C1APIAppV1AppOwnersAddResponse struct {
	// Successful response
	AddAppOwnerResponse *shared.AddAppOwnerResponse
	ContentType         string
	StatusCode          int
	RawResponse         *http.Response
}

func (o *C1APIAppV1AppOwnersAddResponse) GetAddAppOwnerResponse() *shared.AddAppOwnerResponse {
	if o == nil {
		return nil
	}
	return o.AddAppOwnerResponse
}

func (o *C1APIAppV1AppOwnersAddResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppOwnersAddResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppOwnersAddResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
