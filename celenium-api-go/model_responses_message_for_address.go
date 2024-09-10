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

// checks if the ResponsesMessageForAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesMessageForAddress{}

// ResponsesMessageForAddress struct for ResponsesMessageForAddress
type ResponsesMessageForAddress struct {
	Data map[string]map[string]interface{} `json:"data,omitempty"`
	Height *int64 `json:"height,omitempty"`
	Id *int64 `json:"id,omitempty"`
	InvocationType *TypesMsgAddressType `json:"invocation_type,omitempty"`
	Position *int64 `json:"position,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	Tx *ResponsesTxForAddress `json:"tx,omitempty"`
	TxId *int64 `json:"tx_id,omitempty"`
	Type *TypesMsgType `json:"type,omitempty"`
}

// NewResponsesMessageForAddress instantiates a new ResponsesMessageForAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesMessageForAddress() *ResponsesMessageForAddress {
	this := ResponsesMessageForAddress{}
	return &this
}

// NewResponsesMessageForAddressWithDefaults instantiates a new ResponsesMessageForAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesMessageForAddressWithDefaults() *ResponsesMessageForAddress {
	this := ResponsesMessageForAddress{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetData() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]map[string]interface{} and assigns it to the Data field.
func (o *ResponsesMessageForAddress) SetData(v map[string]map[string]interface{}) {
	o.Data = v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetHeight() int64 {
	if o == nil || IsNil(o.Height) {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetHeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *ResponsesMessageForAddress) SetHeight(v int64) {
	o.Height = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ResponsesMessageForAddress) SetId(v int64) {
	o.Id = &v
}

// GetInvocationType returns the InvocationType field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetInvocationType() TypesMsgAddressType {
	if o == nil || IsNil(o.InvocationType) {
		var ret TypesMsgAddressType
		return ret
	}
	return *o.InvocationType
}

// GetInvocationTypeOk returns a tuple with the InvocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetInvocationTypeOk() (*TypesMsgAddressType, bool) {
	if o == nil || IsNil(o.InvocationType) {
		return nil, false
	}
	return o.InvocationType, true
}

// HasInvocationType returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasInvocationType() bool {
	if o != nil && !IsNil(o.InvocationType) {
		return true
	}

	return false
}

// SetInvocationType gets a reference to the given TypesMsgAddressType and assigns it to the InvocationType field.
func (o *ResponsesMessageForAddress) SetInvocationType(v TypesMsgAddressType) {
	o.InvocationType = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetPosition() int64 {
	if o == nil || IsNil(o.Position) {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *ResponsesMessageForAddress) SetPosition(v int64) {
	o.Position = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ResponsesMessageForAddress) SetSize(v int32) {
	o.Size = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ResponsesMessageForAddress) SetTime(v time.Time) {
	o.Time = &v
}

// GetTx returns the Tx field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetTx() ResponsesTxForAddress {
	if o == nil || IsNil(o.Tx) {
		var ret ResponsesTxForAddress
		return ret
	}
	return *o.Tx
}

// GetTxOk returns a tuple with the Tx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetTxOk() (*ResponsesTxForAddress, bool) {
	if o == nil || IsNil(o.Tx) {
		return nil, false
	}
	return o.Tx, true
}

// HasTx returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasTx() bool {
	if o != nil && !IsNil(o.Tx) {
		return true
	}

	return false
}

// SetTx gets a reference to the given ResponsesTxForAddress and assigns it to the Tx field.
func (o *ResponsesMessageForAddress) SetTx(v ResponsesTxForAddress) {
	o.Tx = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetTxId() int64 {
	if o == nil || IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetTxIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasTxId() bool {
	if o != nil && !IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *ResponsesMessageForAddress) SetTxId(v int64) {
	o.TxId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResponsesMessageForAddress) GetType() TypesMsgType {
	if o == nil || IsNil(o.Type) {
		var ret TypesMsgType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesMessageForAddress) GetTypeOk() (*TypesMsgType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResponsesMessageForAddress) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TypesMsgType and assigns it to the Type field.
func (o *ResponsesMessageForAddress) SetType(v TypesMsgType) {
	o.Type = &v
}

func (o ResponsesMessageForAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesMessageForAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InvocationType) {
		toSerialize["invocation_type"] = o.InvocationType
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Tx) {
		toSerialize["tx"] = o.Tx
	}
	if !IsNil(o.TxId) {
		toSerialize["tx_id"] = o.TxId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableResponsesMessageForAddress struct {
	value *ResponsesMessageForAddress
	isSet bool
}

func (v NullableResponsesMessageForAddress) Get() *ResponsesMessageForAddress {
	return v.value
}

func (v *NullableResponsesMessageForAddress) Set(val *ResponsesMessageForAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesMessageForAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesMessageForAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesMessageForAddress(val *ResponsesMessageForAddress) *NullableResponsesMessageForAddress {
	return &NullableResponsesMessageForAddress{value: val, isSet: true}
}

func (v NullableResponsesMessageForAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesMessageForAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


