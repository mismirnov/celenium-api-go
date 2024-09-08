/*
Celenium API

This is docs of Celenium API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"encoding/json"
	"time"
)

// checks if the ResponsesHistogramItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesHistogramItem{}

// ResponsesHistogramItem struct for ResponsesHistogramItem
type ResponsesHistogramItem struct {
	Time *time.Time `json:"time,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewResponsesHistogramItem instantiates a new ResponsesHistogramItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesHistogramItem() *ResponsesHistogramItem {
	this := ResponsesHistogramItem{}
	return &this
}

// NewResponsesHistogramItemWithDefaults instantiates a new ResponsesHistogramItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesHistogramItemWithDefaults() *ResponsesHistogramItem {
	this := ResponsesHistogramItem{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesHistogramItem) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHistogramItem) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesHistogramItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ResponsesHistogramItem) SetTime(v time.Time) {
	o.Time = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ResponsesHistogramItem) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHistogramItem) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResponsesHistogramItem) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResponsesHistogramItem) SetValue(v string) {
	o.Value = &v
}

func (o ResponsesHistogramItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesHistogramItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableResponsesHistogramItem struct {
	value *ResponsesHistogramItem
	isSet bool
}

func (v NullableResponsesHistogramItem) Get() *ResponsesHistogramItem {
	return v.value
}

func (v *NullableResponsesHistogramItem) Set(val *ResponsesHistogramItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesHistogramItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesHistogramItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesHistogramItem(val *ResponsesHistogramItem) *NullableResponsesHistogramItem {
	return &NullableResponsesHistogramItem{value: val, isSet: true}
}

func (v NullableResponsesHistogramItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesHistogramItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


