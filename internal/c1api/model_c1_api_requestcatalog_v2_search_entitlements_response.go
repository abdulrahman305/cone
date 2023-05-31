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

// checks if the C1ApiRequestcatalogV2SearchEntitlementsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiRequestcatalogV2SearchEntitlementsResponse{}

// C1ApiRequestcatalogV2SearchEntitlementsResponse The SearchEntitlementsResponse message.
type C1ApiRequestcatalogV2SearchEntitlementsResponse struct {
	// The list field.
	List interface{} `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken interface{} `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiRequestcatalogV2SearchEntitlementsResponse C1ApiRequestcatalogV2SearchEntitlementsResponse

// NewC1ApiRequestcatalogV2SearchEntitlementsResponse instantiates a new C1ApiRequestcatalogV2SearchEntitlementsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiRequestcatalogV2SearchEntitlementsResponse() *C1ApiRequestcatalogV2SearchEntitlementsResponse {
	this := C1ApiRequestcatalogV2SearchEntitlementsResponse{}
	return &this
}

// NewC1ApiRequestcatalogV2SearchEntitlementsResponseWithDefaults instantiates a new C1ApiRequestcatalogV2SearchEntitlementsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiRequestcatalogV2SearchEntitlementsResponseWithDefaults() *C1ApiRequestcatalogV2SearchEntitlementsResponse {
	this := C1ApiRequestcatalogV2SearchEntitlementsResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiRequestcatalogV2SearchEntitlementsResponse) GetList() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiRequestcatalogV2SearchEntitlementsResponse) GetListOk() (*interface{}, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return &o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *C1ApiRequestcatalogV2SearchEntitlementsResponse) HasList() bool {
	if o != nil && IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given interface{} and assigns it to the List field.
func (o *C1ApiRequestcatalogV2SearchEntitlementsResponse) SetList(v interface{}) {
	o.List = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiRequestcatalogV2SearchEntitlementsResponse) GetNextPageToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiRequestcatalogV2SearchEntitlementsResponse) GetNextPageTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return &o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *C1ApiRequestcatalogV2SearchEntitlementsResponse) HasNextPageToken() bool {
	if o != nil && IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given interface{} and assigns it to the NextPageToken field.
func (o *C1ApiRequestcatalogV2SearchEntitlementsResponse) SetNextPageToken(v interface{}) {
	o.NextPageToken = v
}

func (o C1ApiRequestcatalogV2SearchEntitlementsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiRequestcatalogV2SearchEntitlementsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.List != nil {
		toSerialize["list"] = o.List
	}
	if o.NextPageToken != nil {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiRequestcatalogV2SearchEntitlementsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiRequestcatalogV2SearchEntitlementsResponse := _C1ApiRequestcatalogV2SearchEntitlementsResponse{}

	if err = json.Unmarshal(bytes, &varC1ApiRequestcatalogV2SearchEntitlementsResponse); err == nil {
		*o = C1ApiRequestcatalogV2SearchEntitlementsResponse(varC1ApiRequestcatalogV2SearchEntitlementsResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "list")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiRequestcatalogV2SearchEntitlementsResponse struct {
	value *C1ApiRequestcatalogV2SearchEntitlementsResponse
	isSet bool
}

func (v NullableC1ApiRequestcatalogV2SearchEntitlementsResponse) Get() *C1ApiRequestcatalogV2SearchEntitlementsResponse {
	return v.value
}

func (v *NullableC1ApiRequestcatalogV2SearchEntitlementsResponse) Set(val *C1ApiRequestcatalogV2SearchEntitlementsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiRequestcatalogV2SearchEntitlementsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiRequestcatalogV2SearchEntitlementsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiRequestcatalogV2SearchEntitlementsResponse(val *C1ApiRequestcatalogV2SearchEntitlementsResponse) *NullableC1ApiRequestcatalogV2SearchEntitlementsResponse {
	return &NullableC1ApiRequestcatalogV2SearchEntitlementsResponse{value: val, isSet: true}
}

func (v NullableC1ApiRequestcatalogV2SearchEntitlementsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiRequestcatalogV2SearchEntitlementsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


