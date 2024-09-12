/*
Celenium API

Celenium API is a powerful tool to access all blockchain data that is processed and indexed by our proprietary indexer. With Celenium API you can retrieve all historical data, off-chain data, blobs and statistics through our REST API. Celenium API indexer are open source, which allows you to not depend on third-party services. You can clone, build and run them independently, giving you full control over all components. If you have any questions or feature requests, please feel free to contact us. We appreciate your feedback!

API version: 1.0
Contact: celenium@pklabs.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"encoding/json"
	"time"
)

// checks if the ResponsesVestingPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesVestingPeriod{}

// ResponsesVestingPeriod struct for ResponsesVestingPeriod
type ResponsesVestingPeriod struct {
	Amount *string `json:"amount,omitempty"`
	Time *time.Time `json:"time,omitempty"`
}

// NewResponsesVestingPeriod instantiates a new ResponsesVestingPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesVestingPeriod() *ResponsesVestingPeriod {
	this := ResponsesVestingPeriod{}
	return &this
}

// NewResponsesVestingPeriodWithDefaults instantiates a new ResponsesVestingPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesVestingPeriodWithDefaults() *ResponsesVestingPeriod {
	this := ResponsesVestingPeriod{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ResponsesVestingPeriod) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesVestingPeriod) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ResponsesVestingPeriod) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *ResponsesVestingPeriod) SetAmount(v string) {
	o.Amount = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesVestingPeriod) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesVestingPeriod) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesVestingPeriod) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ResponsesVestingPeriod) SetTime(v time.Time) {
	o.Time = &v
}

func (o ResponsesVestingPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesVestingPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableResponsesVestingPeriod struct {
	value *ResponsesVestingPeriod
	isSet bool
}

func (v NullableResponsesVestingPeriod) Get() *ResponsesVestingPeriod {
	return v.value
}

func (v *NullableResponsesVestingPeriod) Set(val *ResponsesVestingPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesVestingPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesVestingPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesVestingPeriod(val *ResponsesVestingPeriod) *NullableResponsesVestingPeriod {
	return &NullableResponsesVestingPeriod{value: val, isSet: true}
}

func (v NullableResponsesVestingPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesVestingPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


