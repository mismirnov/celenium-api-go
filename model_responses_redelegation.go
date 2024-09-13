/*
Celenium API

Celenium TEST API is a powerful tool to access all blockchain data that is processed and indexed by our proprietary indexer. With Celenium API you can retrieve all historical data, off-chain data, blobs and statistics through our REST API. Celenium API indexer are open source, which allows you to not depend on third-party services. You can clone, build and run them independently, giving you full control over all components. If you have any questions or feature requests, please feel free to contact us. We appreciate your feedback!

API version: 1.0
Contact: celenium@pklabs.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"encoding/json"
)

// checks if the ResponsesRedelegation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesRedelegation{}

// ResponsesRedelegation struct for ResponsesRedelegation
type ResponsesRedelegation struct {
	Amount *string `json:"amount,omitempty"`
	CompletionTime *string `json:"completion_time,omitempty"`
	Delegator *string `json:"delegator,omitempty"`
	Destination *ResponsesShortValidator `json:"destination,omitempty"`
	Height *int32 `json:"height,omitempty"`
	Source *ResponsesShortValidator `json:"source,omitempty"`
	Time *string `json:"time,omitempty"`
}

// NewResponsesRedelegation instantiates a new ResponsesRedelegation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesRedelegation() *ResponsesRedelegation {
	this := ResponsesRedelegation{}
	return &this
}

// NewResponsesRedelegationWithDefaults instantiates a new ResponsesRedelegation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesRedelegationWithDefaults() *ResponsesRedelegation {
	this := ResponsesRedelegation{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ResponsesRedelegation) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRedelegation) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ResponsesRedelegation) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *ResponsesRedelegation) SetAmount(v string) {
	o.Amount = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *ResponsesRedelegation) GetCompletionTime() string {
	if o == nil || IsNil(o.CompletionTime) {
		var ret string
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRedelegation) GetCompletionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CompletionTime) {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *ResponsesRedelegation) HasCompletionTime() bool {
	if o != nil && !IsNil(o.CompletionTime) {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given string and assigns it to the CompletionTime field.
func (o *ResponsesRedelegation) SetCompletionTime(v string) {
	o.CompletionTime = &v
}

// GetDelegator returns the Delegator field value if set, zero value otherwise.
func (o *ResponsesRedelegation) GetDelegator() string {
	if o == nil || IsNil(o.Delegator) {
		var ret string
		return ret
	}
	return *o.Delegator
}

// GetDelegatorOk returns a tuple with the Delegator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRedelegation) GetDelegatorOk() (*string, bool) {
	if o == nil || IsNil(o.Delegator) {
		return nil, false
	}
	return o.Delegator, true
}

// HasDelegator returns a boolean if a field has been set.
func (o *ResponsesRedelegation) HasDelegator() bool {
	if o != nil && !IsNil(o.Delegator) {
		return true
	}

	return false
}

// SetDelegator gets a reference to the given string and assigns it to the Delegator field.
func (o *ResponsesRedelegation) SetDelegator(v string) {
	o.Delegator = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *ResponsesRedelegation) GetDestination() ResponsesShortValidator {
	if o == nil || IsNil(o.Destination) {
		var ret ResponsesShortValidator
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRedelegation) GetDestinationOk() (*ResponsesShortValidator, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *ResponsesRedelegation) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given ResponsesShortValidator and assigns it to the Destination field.
func (o *ResponsesRedelegation) SetDestination(v ResponsesShortValidator) {
	o.Destination = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ResponsesRedelegation) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRedelegation) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ResponsesRedelegation) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ResponsesRedelegation) SetHeight(v int32) {
	o.Height = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ResponsesRedelegation) GetSource() ResponsesShortValidator {
	if o == nil || IsNil(o.Source) {
		var ret ResponsesShortValidator
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRedelegation) GetSourceOk() (*ResponsesShortValidator, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ResponsesRedelegation) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given ResponsesShortValidator and assigns it to the Source field.
func (o *ResponsesRedelegation) SetSource(v ResponsesShortValidator) {
	o.Source = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesRedelegation) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRedelegation) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesRedelegation) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *ResponsesRedelegation) SetTime(v string) {
	o.Time = &v
}

func (o ResponsesRedelegation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesRedelegation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CompletionTime) {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if !IsNil(o.Delegator) {
		toSerialize["delegator"] = o.Delegator
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableResponsesRedelegation struct {
	value *ResponsesRedelegation
	isSet bool
}

func (v NullableResponsesRedelegation) Get() *ResponsesRedelegation {
	return v.value
}

func (v *NullableResponsesRedelegation) Set(val *ResponsesRedelegation) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesRedelegation) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesRedelegation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesRedelegation(val *ResponsesRedelegation) *NullableResponsesRedelegation {
	return &NullableResponsesRedelegation{value: val, isSet: true}
}

func (v NullableResponsesRedelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesRedelegation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


