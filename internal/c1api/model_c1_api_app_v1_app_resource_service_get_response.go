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

// checks if the C1ApiAppV1AppResourceServiceGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiAppV1AppResourceServiceGetResponse{}

// C1ApiAppV1AppResourceServiceGetResponse The AppResourceServiceGetResponse message.
type C1ApiAppV1AppResourceServiceGetResponse struct {
	AppResourceView *C1ApiAppV1AppResourceView `json:"appResourceView,omitempty"`
	// The expanded field.
	Expanded interface{} `json:"expanded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiAppV1AppResourceServiceGetResponse C1ApiAppV1AppResourceServiceGetResponse

// NewC1ApiAppV1AppResourceServiceGetResponse instantiates a new C1ApiAppV1AppResourceServiceGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiAppV1AppResourceServiceGetResponse() *C1ApiAppV1AppResourceServiceGetResponse {
	this := C1ApiAppV1AppResourceServiceGetResponse{}
	return &this
}

// NewC1ApiAppV1AppResourceServiceGetResponseWithDefaults instantiates a new C1ApiAppV1AppResourceServiceGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiAppV1AppResourceServiceGetResponseWithDefaults() *C1ApiAppV1AppResourceServiceGetResponse {
	this := C1ApiAppV1AppResourceServiceGetResponse{}
	return &this
}

// GetAppResourceView returns the AppResourceView field value if set, zero value otherwise.
func (o *C1ApiAppV1AppResourceServiceGetResponse) GetAppResourceView() C1ApiAppV1AppResourceView {
	if o == nil || IsNil(o.AppResourceView) {
		var ret C1ApiAppV1AppResourceView
		return ret
	}
	return *o.AppResourceView
}

// GetAppResourceViewOk returns a tuple with the AppResourceView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppResourceServiceGetResponse) GetAppResourceViewOk() (*C1ApiAppV1AppResourceView, bool) {
	if o == nil || IsNil(o.AppResourceView) {
		return nil, false
	}
	return o.AppResourceView, true
}

// HasAppResourceView returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceServiceGetResponse) HasAppResourceView() bool {
	if o != nil && !IsNil(o.AppResourceView) {
		return true
	}

	return false
}

// SetAppResourceView gets a reference to the given C1ApiAppV1AppResourceView and assigns it to the AppResourceView field.
func (o *C1ApiAppV1AppResourceServiceGetResponse) SetAppResourceView(v C1ApiAppV1AppResourceView) {
	o.AppResourceView = &v
}

// GetExpanded returns the Expanded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiAppV1AppResourceServiceGetResponse) GetExpanded() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Expanded
}

// GetExpandedOk returns a tuple with the Expanded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiAppV1AppResourceServiceGetResponse) GetExpandedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Expanded) {
		return nil, false
	}
	return &o.Expanded, true
}

// HasExpanded returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceServiceGetResponse) HasExpanded() bool {
	if o != nil && IsNil(o.Expanded) {
		return true
	}

	return false
}

// SetExpanded gets a reference to the given interface{} and assigns it to the Expanded field.
func (o *C1ApiAppV1AppResourceServiceGetResponse) SetExpanded(v interface{}) {
	o.Expanded = v
}

func (o C1ApiAppV1AppResourceServiceGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiAppV1AppResourceServiceGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppResourceView) {
		toSerialize["appResourceView"] = o.AppResourceView
	}
	if o.Expanded != nil {
		toSerialize["expanded"] = o.Expanded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiAppV1AppResourceServiceGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiAppV1AppResourceServiceGetResponse := _C1ApiAppV1AppResourceServiceGetResponse{}

	if err = json.Unmarshal(bytes, &varC1ApiAppV1AppResourceServiceGetResponse); err == nil {
		*o = C1ApiAppV1AppResourceServiceGetResponse(varC1ApiAppV1AppResourceServiceGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appResourceView")
		delete(additionalProperties, "expanded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiAppV1AppResourceServiceGetResponse struct {
	value *C1ApiAppV1AppResourceServiceGetResponse
	isSet bool
}

func (v NullableC1ApiAppV1AppResourceServiceGetResponse) Get() *C1ApiAppV1AppResourceServiceGetResponse {
	return v.value
}

func (v *NullableC1ApiAppV1AppResourceServiceGetResponse) Set(val *C1ApiAppV1AppResourceServiceGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiAppV1AppResourceServiceGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiAppV1AppResourceServiceGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiAppV1AppResourceServiceGetResponse(val *C1ApiAppV1AppResourceServiceGetResponse) *NullableC1ApiAppV1AppResourceServiceGetResponse {
	return &NullableC1ApiAppV1AppResourceServiceGetResponse{value: val, isSet: true}
}

func (v NullableC1ApiAppV1AppResourceServiceGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiAppV1AppResourceServiceGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


