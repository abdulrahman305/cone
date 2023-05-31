/*
ConductorOne API

The ConductorOne API is a HTTP API for managing ConductorOne resources.

API version: 0.1.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package c1api

import (
	"encoding/json"
)

// checks if the C1ApiRequestcatalogV2SearchEntitlementsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiRequestcatalogV2SearchEntitlementsRequest{}

// C1ApiRequestcatalogV2SearchEntitlementsRequest The SearchEntitlementsRequest message.
type C1ApiRequestcatalogV2SearchEntitlementsRequest struct {
	// The entitlementAlias field.
	EntitlementAlias interface{} `json:"entitlementAlias,omitempty"`
	// The pageSize field.
	PageSize interface{} `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken interface{} `json:"pageToken,omitempty"`
	// The query field.
	Query interface{} `json:"query,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiRequestcatalogV2SearchEntitlementsRequest C1ApiRequestcatalogV2SearchEntitlementsRequest

// NewC1ApiRequestcatalogV2SearchEntitlementsRequest instantiates a new C1ApiRequestcatalogV2SearchEntitlementsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiRequestcatalogV2SearchEntitlementsRequest() *C1ApiRequestcatalogV2SearchEntitlementsRequest {
	this := C1ApiRequestcatalogV2SearchEntitlementsRequest{}
	return &this
}

// NewC1ApiRequestcatalogV2SearchEntitlementsRequestWithDefaults instantiates a new C1ApiRequestcatalogV2SearchEntitlementsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiRequestcatalogV2SearchEntitlementsRequestWithDefaults() *C1ApiRequestcatalogV2SearchEntitlementsRequest {
	this := C1ApiRequestcatalogV2SearchEntitlementsRequest{}
	return &this
}

// GetEntitlementAlias returns the EntitlementAlias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) GetEntitlementAlias() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EntitlementAlias
}

// GetEntitlementAliasOk returns a tuple with the EntitlementAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) GetEntitlementAliasOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EntitlementAlias) {
		return nil, false
	}
	return &o.EntitlementAlias, true
}

// HasEntitlementAlias returns a boolean if a field has been set.
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) HasEntitlementAlias() bool {
	if o != nil && IsNil(o.EntitlementAlias) {
		return true
	}

	return false
}

// SetEntitlementAlias gets a reference to the given interface{} and assigns it to the EntitlementAlias field.
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) SetEntitlementAlias(v interface{}) {
	o.EntitlementAlias = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) GetPageSize() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) GetPageSizeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return &o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) HasPageSize() bool {
	if o != nil && IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given interface{} and assigns it to the PageSize field.
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) SetPageSize(v interface{}) {
	o.PageSize = v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) GetPageToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) GetPageTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return &o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) HasPageToken() bool {
	if o != nil && IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given interface{} and assigns it to the PageToken field.
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) SetPageToken(v interface{}) {
	o.PageToken = v
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) GetQuery() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) GetQueryOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return &o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) HasQuery() bool {
	if o != nil && IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given interface{} and assigns it to the Query field.
func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) SetQuery(v interface{}) {
	o.Query = v
}

func (o C1ApiRequestcatalogV2SearchEntitlementsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiRequestcatalogV2SearchEntitlementsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EntitlementAlias != nil {
		toSerialize["entitlementAlias"] = o.EntitlementAlias
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["pageToken"] = o.PageToken
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiRequestcatalogV2SearchEntitlementsRequest) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiRequestcatalogV2SearchEntitlementsRequest := _C1ApiRequestcatalogV2SearchEntitlementsRequest{}

	if err = json.Unmarshal(bytes, &varC1ApiRequestcatalogV2SearchEntitlementsRequest); err == nil {
		*o = C1ApiRequestcatalogV2SearchEntitlementsRequest(varC1ApiRequestcatalogV2SearchEntitlementsRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementAlias")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageToken")
		delete(additionalProperties, "query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiRequestcatalogV2SearchEntitlementsRequest struct {
	value *C1ApiRequestcatalogV2SearchEntitlementsRequest
	isSet bool
}

func (v NullableC1ApiRequestcatalogV2SearchEntitlementsRequest) Get() *C1ApiRequestcatalogV2SearchEntitlementsRequest {
	return v.value
}

func (v *NullableC1ApiRequestcatalogV2SearchEntitlementsRequest) Set(val *C1ApiRequestcatalogV2SearchEntitlementsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiRequestcatalogV2SearchEntitlementsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiRequestcatalogV2SearchEntitlementsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiRequestcatalogV2SearchEntitlementsRequest(val *C1ApiRequestcatalogV2SearchEntitlementsRequest) *NullableC1ApiRequestcatalogV2SearchEntitlementsRequest {
	return &NullableC1ApiRequestcatalogV2SearchEntitlementsRequest{value: val, isSet: true}
}

func (v NullableC1ApiRequestcatalogV2SearchEntitlementsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiRequestcatalogV2SearchEntitlementsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


