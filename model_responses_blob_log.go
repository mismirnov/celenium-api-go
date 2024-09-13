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
)

// checks if the ResponsesBlobLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesBlobLog{}

// ResponsesBlobLog struct for ResponsesBlobLog
type ResponsesBlobLog struct {
	Commitment *string `json:"commitment,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Height *int32 `json:"height,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Namespace *ResponsesNamespace `json:"namespace,omitempty"`
	Rollup *ResponsesShortRollup `json:"rollup,omitempty"`
	Signer *string `json:"signer,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	Tx *ResponsesTx `json:"tx,omitempty"`
}

// NewResponsesBlobLog instantiates a new ResponsesBlobLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesBlobLog() *ResponsesBlobLog {
	this := ResponsesBlobLog{}
	return &this
}

// NewResponsesBlobLogWithDefaults instantiates a new ResponsesBlobLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesBlobLogWithDefaults() *ResponsesBlobLog {
	this := ResponsesBlobLog{}
	return &this
}

// GetCommitment returns the Commitment field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetCommitment() string {
	if o == nil || IsNil(o.Commitment) {
		var ret string
		return ret
	}
	return *o.Commitment
}

// GetCommitmentOk returns a tuple with the Commitment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetCommitmentOk() (*string, bool) {
	if o == nil || IsNil(o.Commitment) {
		return nil, false
	}
	return o.Commitment, true
}

// HasCommitment returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasCommitment() bool {
	if o != nil && !IsNil(o.Commitment) {
		return true
	}

	return false
}

// SetCommitment gets a reference to the given string and assigns it to the Commitment field.
func (o *ResponsesBlobLog) SetCommitment(v string) {
	o.Commitment = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *ResponsesBlobLog) SetContentType(v string) {
	o.ContentType = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ResponsesBlobLog) SetHeight(v int32) {
	o.Height = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ResponsesBlobLog) SetId(v int32) {
	o.Id = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetNamespace() ResponsesNamespace {
	if o == nil || IsNil(o.Namespace) {
		var ret ResponsesNamespace
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetNamespaceOk() (*ResponsesNamespace, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given ResponsesNamespace and assigns it to the Namespace field.
func (o *ResponsesBlobLog) SetNamespace(v ResponsesNamespace) {
	o.Namespace = &v
}

// GetRollup returns the Rollup field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetRollup() ResponsesShortRollup {
	if o == nil || IsNil(o.Rollup) {
		var ret ResponsesShortRollup
		return ret
	}
	return *o.Rollup
}

// GetRollupOk returns a tuple with the Rollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetRollupOk() (*ResponsesShortRollup, bool) {
	if o == nil || IsNil(o.Rollup) {
		return nil, false
	}
	return o.Rollup, true
}

// HasRollup returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasRollup() bool {
	if o != nil && !IsNil(o.Rollup) {
		return true
	}

	return false
}

// SetRollup gets a reference to the given ResponsesShortRollup and assigns it to the Rollup field.
func (o *ResponsesBlobLog) SetRollup(v ResponsesShortRollup) {
	o.Rollup = &v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetSigner() string {
	if o == nil || IsNil(o.Signer) {
		var ret string
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetSignerOk() (*string, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given string and assigns it to the Signer field.
func (o *ResponsesBlobLog) SetSigner(v string) {
	o.Signer = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ResponsesBlobLog) SetSize(v int32) {
	o.Size = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ResponsesBlobLog) SetTime(v time.Time) {
	o.Time = &v
}

// GetTx returns the Tx field value if set, zero value otherwise.
func (o *ResponsesBlobLog) GetTx() ResponsesTx {
	if o == nil || IsNil(o.Tx) {
		var ret ResponsesTx
		return ret
	}
	return *o.Tx
}

// GetTxOk returns a tuple with the Tx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlobLog) GetTxOk() (*ResponsesTx, bool) {
	if o == nil || IsNil(o.Tx) {
		return nil, false
	}
	return o.Tx, true
}

// HasTx returns a boolean if a field has been set.
func (o *ResponsesBlobLog) HasTx() bool {
	if o != nil && !IsNil(o.Tx) {
		return true
	}

	return false
}

// SetTx gets a reference to the given ResponsesTx and assigns it to the Tx field.
func (o *ResponsesBlobLog) SetTx(v ResponsesTx) {
	o.Tx = &v
}

func (o ResponsesBlobLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesBlobLog) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Rollup) {
		toSerialize["rollup"] = o.Rollup
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
	if !IsNil(o.Tx) {
		toSerialize["tx"] = o.Tx
	}
	return toSerialize, nil
}

type NullableResponsesBlobLog struct {
	value *ResponsesBlobLog
	isSet bool
}

func (v NullableResponsesBlobLog) Get() *ResponsesBlobLog {
	return v.value
}

func (v *NullableResponsesBlobLog) Set(val *ResponsesBlobLog) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesBlobLog) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesBlobLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesBlobLog(val *ResponsesBlobLog) *NullableResponsesBlobLog {
	return &NullableResponsesBlobLog{value: val, isSet: true}
}

func (v NullableResponsesBlobLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesBlobLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


