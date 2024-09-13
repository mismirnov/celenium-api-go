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
	"os"
	"time"
)

// checks if the ResponsesTx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesTx{}

// ResponsesTx struct for ResponsesTx
type ResponsesTx struct {
	Codespace *string `json:"codespace,omitempty"`
	Error *string `json:"error,omitempty"`
	EventsCount *int64 `json:"events_count,omitempty"`
	Fee *string `json:"fee,omitempty"`
	GasUsed *int64 `json:"gas_used,omitempty"`
	GasWanted *int64 `json:"gas_wanted,omitempty"`
	Hash **os.File `json:"hash,omitempty"`
	Height *int64 `json:"height,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Memo *string `json:"memo,omitempty"`
	MessageTypes []TypesMsgType `json:"message_types,omitempty"`
	Messages []ResponsesMessage `json:"messages,omitempty"`
	MessagesCount *int64 `json:"messages_count,omitempty"`
	Position *int64 `json:"position,omitempty"`
	Signers []string `json:"signers,omitempty"`
	Status *GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus `json:"status,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	TimeoutHeight *int64 `json:"timeout_height,omitempty"`
}

// NewResponsesTx instantiates a new ResponsesTx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesTx() *ResponsesTx {
	this := ResponsesTx{}
	return &this
}

// NewResponsesTxWithDefaults instantiates a new ResponsesTx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesTxWithDefaults() *ResponsesTx {
	this := ResponsesTx{}
	return &this
}

// GetCodespace returns the Codespace field value if set, zero value otherwise.
func (o *ResponsesTx) GetCodespace() string {
	if o == nil || IsNil(o.Codespace) {
		var ret string
		return ret
	}
	return *o.Codespace
}

// GetCodespaceOk returns a tuple with the Codespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetCodespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Codespace) {
		return nil, false
	}
	return o.Codespace, true
}

// HasCodespace returns a boolean if a field has been set.
func (o *ResponsesTx) HasCodespace() bool {
	if o != nil && !IsNil(o.Codespace) {
		return true
	}

	return false
}

// SetCodespace gets a reference to the given string and assigns it to the Codespace field.
func (o *ResponsesTx) SetCodespace(v string) {
	o.Codespace = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ResponsesTx) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ResponsesTx) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ResponsesTx) SetError(v string) {
	o.Error = &v
}

// GetEventsCount returns the EventsCount field value if set, zero value otherwise.
func (o *ResponsesTx) GetEventsCount() int64 {
	if o == nil || IsNil(o.EventsCount) {
		var ret int64
		return ret
	}
	return *o.EventsCount
}

// GetEventsCountOk returns a tuple with the EventsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetEventsCountOk() (*int64, bool) {
	if o == nil || IsNil(o.EventsCount) {
		return nil, false
	}
	return o.EventsCount, true
}

// HasEventsCount returns a boolean if a field has been set.
func (o *ResponsesTx) HasEventsCount() bool {
	if o != nil && !IsNil(o.EventsCount) {
		return true
	}

	return false
}

// SetEventsCount gets a reference to the given int64 and assigns it to the EventsCount field.
func (o *ResponsesTx) SetEventsCount(v int64) {
	o.EventsCount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *ResponsesTx) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *ResponsesTx) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *ResponsesTx) SetFee(v string) {
	o.Fee = &v
}

// GetGasUsed returns the GasUsed field value if set, zero value otherwise.
func (o *ResponsesTx) GetGasUsed() int64 {
	if o == nil || IsNil(o.GasUsed) {
		var ret int64
		return ret
	}
	return *o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetGasUsedOk() (*int64, bool) {
	if o == nil || IsNil(o.GasUsed) {
		return nil, false
	}
	return o.GasUsed, true
}

// HasGasUsed returns a boolean if a field has been set.
func (o *ResponsesTx) HasGasUsed() bool {
	if o != nil && !IsNil(o.GasUsed) {
		return true
	}

	return false
}

// SetGasUsed gets a reference to the given int64 and assigns it to the GasUsed field.
func (o *ResponsesTx) SetGasUsed(v int64) {
	o.GasUsed = &v
}

// GetGasWanted returns the GasWanted field value if set, zero value otherwise.
func (o *ResponsesTx) GetGasWanted() int64 {
	if o == nil || IsNil(o.GasWanted) {
		var ret int64
		return ret
	}
	return *o.GasWanted
}

// GetGasWantedOk returns a tuple with the GasWanted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetGasWantedOk() (*int64, bool) {
	if o == nil || IsNil(o.GasWanted) {
		return nil, false
	}
	return o.GasWanted, true
}

// HasGasWanted returns a boolean if a field has been set.
func (o *ResponsesTx) HasGasWanted() bool {
	if o != nil && !IsNil(o.GasWanted) {
		return true
	}

	return false
}

// SetGasWanted gets a reference to the given int64 and assigns it to the GasWanted field.
func (o *ResponsesTx) SetGasWanted(v int64) {
	o.GasWanted = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ResponsesTx) GetHash() *os.File {
	if o == nil || IsNil(o.Hash) {
		var ret *os.File
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetHashOk() (**os.File, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ResponsesTx) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given *os.File and assigns it to the Hash field.
func (o *ResponsesTx) SetHash(v *os.File) {
	o.Hash = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ResponsesTx) GetHeight() int64 {
	if o == nil || IsNil(o.Height) {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetHeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ResponsesTx) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *ResponsesTx) SetHeight(v int64) {
	o.Height = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponsesTx) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponsesTx) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ResponsesTx) SetId(v int64) {
	o.Id = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *ResponsesTx) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *ResponsesTx) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *ResponsesTx) SetMemo(v string) {
	o.Memo = &v
}

// GetMessageTypes returns the MessageTypes field value if set, zero value otherwise.
func (o *ResponsesTx) GetMessageTypes() []TypesMsgType {
	if o == nil || IsNil(o.MessageTypes) {
		var ret []TypesMsgType
		return ret
	}
	return o.MessageTypes
}

// GetMessageTypesOk returns a tuple with the MessageTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetMessageTypesOk() ([]TypesMsgType, bool) {
	if o == nil || IsNil(o.MessageTypes) {
		return nil, false
	}
	return o.MessageTypes, true
}

// HasMessageTypes returns a boolean if a field has been set.
func (o *ResponsesTx) HasMessageTypes() bool {
	if o != nil && !IsNil(o.MessageTypes) {
		return true
	}

	return false
}

// SetMessageTypes gets a reference to the given []TypesMsgType and assigns it to the MessageTypes field.
func (o *ResponsesTx) SetMessageTypes(v []TypesMsgType) {
	o.MessageTypes = v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *ResponsesTx) GetMessages() []ResponsesMessage {
	if o == nil || IsNil(o.Messages) {
		var ret []ResponsesMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetMessagesOk() ([]ResponsesMessage, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ResponsesTx) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ResponsesMessage and assigns it to the Messages field.
func (o *ResponsesTx) SetMessages(v []ResponsesMessage) {
	o.Messages = v
}

// GetMessagesCount returns the MessagesCount field value if set, zero value otherwise.
func (o *ResponsesTx) GetMessagesCount() int64 {
	if o == nil || IsNil(o.MessagesCount) {
		var ret int64
		return ret
	}
	return *o.MessagesCount
}

// GetMessagesCountOk returns a tuple with the MessagesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetMessagesCountOk() (*int64, bool) {
	if o == nil || IsNil(o.MessagesCount) {
		return nil, false
	}
	return o.MessagesCount, true
}

// HasMessagesCount returns a boolean if a field has been set.
func (o *ResponsesTx) HasMessagesCount() bool {
	if o != nil && !IsNil(o.MessagesCount) {
		return true
	}

	return false
}

// SetMessagesCount gets a reference to the given int64 and assigns it to the MessagesCount field.
func (o *ResponsesTx) SetMessagesCount(v int64) {
	o.MessagesCount = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ResponsesTx) GetPosition() int64 {
	if o == nil || IsNil(o.Position) {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ResponsesTx) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *ResponsesTx) SetPosition(v int64) {
	o.Position = &v
}

// GetSigners returns the Signers field value if set, zero value otherwise.
func (o *ResponsesTx) GetSigners() []string {
	if o == nil || IsNil(o.Signers) {
		var ret []string
		return ret
	}
	return o.Signers
}

// GetSignersOk returns a tuple with the Signers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetSignersOk() ([]string, bool) {
	if o == nil || IsNil(o.Signers) {
		return nil, false
	}
	return o.Signers, true
}

// HasSigners returns a boolean if a field has been set.
func (o *ResponsesTx) HasSigners() bool {
	if o != nil && !IsNil(o.Signers) {
		return true
	}

	return false
}

// SetSigners gets a reference to the given []string and assigns it to the Signers field.
func (o *ResponsesTx) SetSigners(v []string) {
	o.Signers = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponsesTx) GetStatus() GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus {
	if o == nil || IsNil(o.Status) {
		var ret GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetStatusOk() (*GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponsesTx) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus and assigns it to the Status field.
func (o *ResponsesTx) SetStatus(v GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) {
	o.Status = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesTx) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesTx) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ResponsesTx) SetTime(v time.Time) {
	o.Time = &v
}

// GetTimeoutHeight returns the TimeoutHeight field value if set, zero value otherwise.
func (o *ResponsesTx) GetTimeoutHeight() int64 {
	if o == nil || IsNil(o.TimeoutHeight) {
		var ret int64
		return ret
	}
	return *o.TimeoutHeight
}

// GetTimeoutHeightOk returns a tuple with the TimeoutHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesTx) GetTimeoutHeightOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeoutHeight) {
		return nil, false
	}
	return o.TimeoutHeight, true
}

// HasTimeoutHeight returns a boolean if a field has been set.
func (o *ResponsesTx) HasTimeoutHeight() bool {
	if o != nil && !IsNil(o.TimeoutHeight) {
		return true
	}

	return false
}

// SetTimeoutHeight gets a reference to the given int64 and assigns it to the TimeoutHeight field.
func (o *ResponsesTx) SetTimeoutHeight(v int64) {
	o.TimeoutHeight = &v
}

func (o ResponsesTx) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesTx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Codespace) {
		toSerialize["codespace"] = o.Codespace
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.EventsCount) {
		toSerialize["events_count"] = o.EventsCount
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.GasUsed) {
		toSerialize["gas_used"] = o.GasUsed
	}
	if !IsNil(o.GasWanted) {
		toSerialize["gas_wanted"] = o.GasWanted
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.MessageTypes) {
		toSerialize["message_types"] = o.MessageTypes
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.MessagesCount) {
		toSerialize["messages_count"] = o.MessagesCount
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Signers) {
		toSerialize["signers"] = o.Signers
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TimeoutHeight) {
		toSerialize["timeout_height"] = o.TimeoutHeight
	}
	return toSerialize, nil
}

type NullableResponsesTx struct {
	value *ResponsesTx
	isSet bool
}

func (v NullableResponsesTx) Get() *ResponsesTx {
	return v.value
}

func (v *NullableResponsesTx) Set(val *ResponsesTx) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesTx) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesTx(val *ResponsesTx) *NullableResponsesTx {
	return &NullableResponsesTx{value: val, isSet: true}
}

func (v NullableResponsesTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


