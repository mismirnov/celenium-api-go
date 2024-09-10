/*
Celenium API

This is docs of Celenium API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"encoding/json"
	"os"
)

// checks if the ResponsesNamespaceUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesNamespaceUsage{}

// ResponsesNamespaceUsage struct for ResponsesNamespaceUsage
type ResponsesNamespaceUsage struct {
	Name *string `json:"name,omitempty"`
	NamespaceId **os.File `json:"namespace_id,omitempty"`
	Size *float32 `json:"size,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewResponsesNamespaceUsage instantiates a new ResponsesNamespaceUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesNamespaceUsage() *ResponsesNamespaceUsage {
	this := ResponsesNamespaceUsage{}
	return &this
}

// NewResponsesNamespaceUsageWithDefaults instantiates a new ResponsesNamespaceUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesNamespaceUsageWithDefaults() *ResponsesNamespaceUsage {
	this := ResponsesNamespaceUsage{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponsesNamespaceUsage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesNamespaceUsage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponsesNamespaceUsage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResponsesNamespaceUsage) SetName(v string) {
	o.Name = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *ResponsesNamespaceUsage) GetNamespaceId() *os.File {
	if o == nil || IsNil(o.NamespaceId) {
		var ret *os.File
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesNamespaceUsage) GetNamespaceIdOk() (**os.File, bool) {
	if o == nil || IsNil(o.NamespaceId) {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *ResponsesNamespaceUsage) HasNamespaceId() bool {
	if o != nil && !IsNil(o.NamespaceId) {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given *os.File and assigns it to the NamespaceId field.
func (o *ResponsesNamespaceUsage) SetNamespaceId(v *os.File) {
	o.NamespaceId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ResponsesNamespaceUsage) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesNamespaceUsage) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ResponsesNamespaceUsage) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ResponsesNamespaceUsage) SetSize(v float32) {
	o.Size = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ResponsesNamespaceUsage) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesNamespaceUsage) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ResponsesNamespaceUsage) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ResponsesNamespaceUsage) SetVersion(v int32) {
	o.Version = &v
}

func (o ResponsesNamespaceUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesNamespaceUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NamespaceId) {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableResponsesNamespaceUsage struct {
	value *ResponsesNamespaceUsage
	isSet bool
}

func (v NullableResponsesNamespaceUsage) Get() *ResponsesNamespaceUsage {
	return v.value
}

func (v *NullableResponsesNamespaceUsage) Set(val *ResponsesNamespaceUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesNamespaceUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesNamespaceUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesNamespaceUsage(val *ResponsesNamespaceUsage) *NullableResponsesNamespaceUsage {
	return &NullableResponsesNamespaceUsage{value: val, isSet: true}
}

func (v NullableResponsesNamespaceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesNamespaceUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


