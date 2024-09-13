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
	"time"
	"os"
)

// checks if the ResponsesLightBlobLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesLightBlobLog{}

// ResponsesLightBlobLog struct for ResponsesLightBlobLog
type ResponsesLightBlobLog struct {
	Commitment *string `json:"commitment,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Height *int32 `json:"height,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Signer *string `json:"signer,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	TxHash **os.File `json:"tx_hash,omitempty"`
}

// NewResponsesLightBlobLog instantiates a new ResponsesLightBlobLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesLightBlobLog() *ResponsesLightBlobLog {
	this := ResponsesLightBlobLog{}
	return &this
}

// NewResponsesLightBlobLogWithDefaults instantiates a new ResponsesLightBlobLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesLightBlobLogWithDefaults() *ResponsesLightBlobLog {
	this := ResponsesLightBlobLog{}
	return &this
}

// GetCommitment returns the Commitment field value if set, zero value otherwise.
func (o *ResponsesLightBlobLog) GetCommitment() string {
	if o == nil || IsNil(o.Commitment) {
		var ret string
		return ret
	}
	return *o.Commitment
}

// GetCommitmentOk returns a tuple with the Commitment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesLightBlobLog) GetCommitmentOk() (*string, bool) {
	if o == nil || IsNil(o.Commitment) {
		return nil, false
	}
	return o.Commitment, true
}

// HasCommitment returns a boolean if a field has been set.
func (o *ResponsesLightBlobLog) HasCommitment() bool {
	if o != nil && !IsNil(o.Commitment) {
		return true
	}

	return false
}

// SetCommitment gets a reference to the given string and assigns it to the Commitment field.
func (o *ResponsesLightBlobLog) SetCommitment(v string) {
	o.Commitment = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *ResponsesLightBlobLog) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesLightBlobLog) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *ResponsesLightBlobLog) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *ResponsesLightBlobLog) SetContentType(v string) {
	o.ContentType = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ResponsesLightBlobLog) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesLightBlobLog) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ResponsesLightBlobLog) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ResponsesLightBlobLog) SetHeight(v int32) {
	o.Height = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponsesLightBlobLog) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesLightBlobLog) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponsesLightBlobLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ResponsesLightBlobLog) SetId(v int32) {
	o.Id = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ResponsesLightBlobLog) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesLightBlobLog) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ResponsesLightBlobLog) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ResponsesLightBlobLog) SetNamespace(v string) {
	o.Namespace = &v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *ResponsesLightBlobLog) GetSigner() string {
	if o == nil || IsNil(o.Signer) {
		var ret string
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesLightBlobLog) GetSignerOk() (*string, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *ResponsesLightBlobLog) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given string and assigns it to the Signer field.
func (o *ResponsesLightBlobLog) SetSigner(v string) {
	o.Signer = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ResponsesLightBlobLog) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesLightBlobLog) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ResponsesLightBlobLog) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ResponsesLightBlobLog) SetSize(v int32) {
	o.Size = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesLightBlobLog) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesLightBlobLog) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesLightBlobLog) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ResponsesLightBlobLog) SetTime(v time.Time) {
	o.Time = &v
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *ResponsesLightBlobLog) GetTxHash() *os.File {
	if o == nil || IsNil(o.TxHash) {
		var ret *os.File
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesLightBlobLog) GetTxHashOk() (**os.File, bool) {
	if o == nil || IsNil(o.TxHash) {
		return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *ResponsesLightBlobLog) HasTxHash() bool {
	if o != nil && !IsNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given *os.File and assigns it to the TxHash field.
func (o *ResponsesLightBlobLog) SetTxHash(v *os.File) {
	o.TxHash = &v
}

func (o ResponsesLightBlobLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesLightBlobLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commitment) {
		toSerialize["commitment"] = o.Commitment
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TxHash) {
		toSerialize["tx_hash"] = o.TxHash
	}
	return toSerialize, nil
}

type NullableResponsesLightBlobLog struct {
	value *ResponsesLightBlobLog
	isSet bool
}

func (v NullableResponsesLightBlobLog) Get() *ResponsesLightBlobLog {
	return v.value
}

func (v *NullableResponsesLightBlobLog) Set(val *ResponsesLightBlobLog) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesLightBlobLog) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesLightBlobLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesLightBlobLog(val *ResponsesLightBlobLog) *NullableResponsesLightBlobLog {
	return &NullableResponsesLightBlobLog{value: val, isSet: true}
}

func (v NullableResponsesLightBlobLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesLightBlobLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


