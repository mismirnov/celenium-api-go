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
)

// checks if the ResponsesTPS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesTPS{}

// ResponsesTPS struct for ResponsesTPS
type ResponsesTPS struct {
	ChangeLastHourPct *float32 `json:"change_last_hour_pct,omitempty"`
	Current *float32 `json:"current,omitempty"`
	High *float32 `json:"high,omitempty"`
	Low *float32 `json:"low,omitempty"`
}

// NewResponsesTPS instantiates a new ResponsesTPS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesTPS() *ResponsesTPS {
	this := ResponsesTPS{}
	return &this
}

// NewResponsesTPSWithDefaults instantiates a new ResponsesTPS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesTPSWithDefaults() *ResponsesTPS {
	this := ResponsesTPS{}
	return &this
}

// GetChangeLastHourPct returns the ChangeLastHourPct field value if set, zero value otherwise.
func (o *ResponsesTPS) GetChangeLastHourPct() float32 {
	if o == nil || IsNil(o.ChangeLastHourPct) {
		var ret float32
		return ret
	}
	return *o.ChangeLastHourPct
}

// GetChangeLastHourPctOk returns a tuple with the ChangeLastHourPct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTPS) GetChangeLastHourPctOk() (*float32, bool) {
	if o == nil || IsNil(o.ChangeLastHourPct) {
		return nil, false
	}
	return o.ChangeLastHourPct, true
}

// HasChangeLastHourPct returns a boolean if a field has been set.
func (o *ResponsesTPS) HasChangeLastHourPct() bool {
	if o != nil && !IsNil(o.ChangeLastHourPct) {
		return true
	}

	return false
}

// SetChangeLastHourPct gets a reference to the given float32 and assigns it to the ChangeLastHourPct field.
func (o *ResponsesTPS) SetChangeLastHourPct(v float32) {
	o.ChangeLastHourPct = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *ResponsesTPS) GetCurrent() float32 {
	if o == nil || IsNil(o.Current) {
		var ret float32
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTPS) GetCurrentOk() (*float32, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *ResponsesTPS) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given float32 and assigns it to the Current field.
func (o *ResponsesTPS) SetCurrent(v float32) {
	o.Current = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *ResponsesTPS) GetHigh() float32 {
	if o == nil || IsNil(o.High) {
		var ret float32
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTPS) GetHighOk() (*float32, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *ResponsesTPS) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given float32 and assigns it to the High field.
func (o *ResponsesTPS) SetHigh(v float32) {
	o.High = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *ResponsesTPS) GetLow() float32 {
	if o == nil || IsNil(o.Low) {
		var ret float32
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTPS) GetLowOk() (*float32, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *ResponsesTPS) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given float32 and assigns it to the Low field.
func (o *ResponsesTPS) SetLow(v float32) {
	o.Low = &v
}

func (o ResponsesTPS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesTPS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangeLastHourPct) {
		toSerialize["change_last_hour_pct"] = o.ChangeLastHourPct
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	return toSerialize, nil
}

type NullableResponsesTPS struct {
	value *ResponsesTPS
	isSet bool
}

func (v NullableResponsesTPS) Get() *ResponsesTPS {
	return v.value
}

func (v *NullableResponsesTPS) Set(val *ResponsesTPS) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesTPS) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesTPS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesTPS(val *ResponsesTPS) *NullableResponsesTPS {
	return &NullableResponsesTPS{value: val, isSet: true}
}

func (v NullableResponsesTPS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesTPS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


