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

// checks if the C1ApiAppV1AppResourceTypeView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiAppV1AppResourceTypeView{}

// C1ApiAppV1AppResourceTypeView The AppResourceTypeView message.
type C1ApiAppV1AppResourceTypeView struct {
	// The appPath field.
	AppPath interface{} `json:"appPath,omitempty"`
	AppResourceType *C1ApiAppV1AppResourceType `json:"appResourceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiAppV1AppResourceTypeView C1ApiAppV1AppResourceTypeView

// NewC1ApiAppV1AppResourceTypeView instantiates a new C1ApiAppV1AppResourceTypeView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiAppV1AppResourceTypeView() *C1ApiAppV1AppResourceTypeView {
	this := C1ApiAppV1AppResourceTypeView{}
	return &this
}

// NewC1ApiAppV1AppResourceTypeViewWithDefaults instantiates a new C1ApiAppV1AppResourceTypeView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiAppV1AppResourceTypeViewWithDefaults() *C1ApiAppV1AppResourceTypeView {
	this := C1ApiAppV1AppResourceTypeView{}
	return &this
}

// GetAppPath returns the AppPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiAppV1AppResourceTypeView) GetAppPath() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AppPath
}

// GetAppPathOk returns a tuple with the AppPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiAppV1AppResourceTypeView) GetAppPathOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AppPath) {
		return nil, false
	}
	return &o.AppPath, true
}

// HasAppPath returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceTypeView) HasAppPath() bool {
	if o != nil && IsNil(o.AppPath) {
		return true
	}

	return false
}

// SetAppPath gets a reference to the given interface{} and assigns it to the AppPath field.
func (o *C1ApiAppV1AppResourceTypeView) SetAppPath(v interface{}) {
	o.AppPath = v
}

// GetAppResourceType returns the AppResourceType field value if set, zero value otherwise.
func (o *C1ApiAppV1AppResourceTypeView) GetAppResourceType() C1ApiAppV1AppResourceType {
	if o == nil || IsNil(o.AppResourceType) {
		var ret C1ApiAppV1AppResourceType
		return ret
	}
	return *o.AppResourceType
}

// GetAppResourceTypeOk returns a tuple with the AppResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiAppV1AppResourceTypeView) GetAppResourceTypeOk() (*C1ApiAppV1AppResourceType, bool) {
	if o == nil || IsNil(o.AppResourceType) {
		return nil, false
	}
	return o.AppResourceType, true
}

// HasAppResourceType returns a boolean if a field has been set.
func (o *C1ApiAppV1AppResourceTypeView) HasAppResourceType() bool {
	if o != nil && !IsNil(o.AppResourceType) {
		return true
	}

	return false
}

// SetAppResourceType gets a reference to the given C1ApiAppV1AppResourceType and assigns it to the AppResourceType field.
func (o *C1ApiAppV1AppResourceTypeView) SetAppResourceType(v C1ApiAppV1AppResourceType) {
	o.AppResourceType = &v
}

func (o C1ApiAppV1AppResourceTypeView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiAppV1AppResourceTypeView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppPath != nil {
		toSerialize["appPath"] = o.AppPath
	}
	if !IsNil(o.AppResourceType) {
		toSerialize["appResourceType"] = o.AppResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiAppV1AppResourceTypeView) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiAppV1AppResourceTypeView := _C1ApiAppV1AppResourceTypeView{}

	if err = json.Unmarshal(bytes, &varC1ApiAppV1AppResourceTypeView); err == nil {
		*o = C1ApiAppV1AppResourceTypeView(varC1ApiAppV1AppResourceTypeView)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appPath")
		delete(additionalProperties, "appResourceType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiAppV1AppResourceTypeView struct {
	value *C1ApiAppV1AppResourceTypeView
	isSet bool
}

func (v NullableC1ApiAppV1AppResourceTypeView) Get() *C1ApiAppV1AppResourceTypeView {
	return v.value
}

func (v *NullableC1ApiAppV1AppResourceTypeView) Set(val *C1ApiAppV1AppResourceTypeView) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiAppV1AppResourceTypeView) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiAppV1AppResourceTypeView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiAppV1AppResourceTypeView(val *C1ApiAppV1AppResourceTypeView) *NullableC1ApiAppV1AppResourceTypeView {
	return &NullableC1ApiAppV1AppResourceTypeView{value: val, isSet: true}
}

func (v NullableC1ApiAppV1AppResourceTypeView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiAppV1AppResourceTypeView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


