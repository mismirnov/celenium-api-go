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

// checks if the ResponsesValidator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesValidator{}

// ResponsesValidator struct for ResponsesValidator
type ResponsesValidator struct {
	Address *string `json:"address,omitempty"`
	Commissions *string `json:"commissions,omitempty"`
	ConsAddress *string `json:"cons_address,omitempty"`
	Contacts *string `json:"contacts,omitempty"`
	Delegator *string `json:"delegator,omitempty"`
	Details *string `json:"details,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Identity *string `json:"identity,omitempty"`
	Jailed *bool `json:"jailed,omitempty"`
	MaxChangeRate *string `json:"max_change_rate,omitempty"`
	MaxRate *string `json:"max_rate,omitempty"`
	MinSelfDelegation *string `json:"min_self_delegation,omitempty"`
	Moniker *string `json:"moniker,omitempty"`
	Rate *string `json:"rate,omitempty"`
	Rewards *string `json:"rewards,omitempty"`
	Stake *string `json:"stake,omitempty"`
	VotingPower *string `json:"voting_power,omitempty"`
	Website *string `json:"website,omitempty"`
}

// NewResponsesValidator instantiates a new ResponsesValidator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesValidator() *ResponsesValidator {
	this := ResponsesValidator{}
	return &this
}

// NewResponsesValidatorWithDefaults instantiates a new ResponsesValidator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesValidatorWithDefaults() *ResponsesValidator {
	this := ResponsesValidator{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ResponsesValidator) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ResponsesValidator) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ResponsesValidator) SetAddress(v string) {
	o.Address = &v
}

// GetCommissions returns the Commissions field value if set, zero value otherwise.
func (o *ResponsesValidator) GetCommissions() string {
	if o == nil || IsNil(o.Commissions) {
		var ret string
		return ret
	}
	return *o.Commissions
}

// GetCommissionsOk returns a tuple with the Commissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetCommissionsOk() (*string, bool) {
	if o == nil || IsNil(o.Commissions) {
		return nil, false
	}
	return o.Commissions, true
}

// HasCommissions returns a boolean if a field has been set.
func (o *ResponsesValidator) HasCommissions() bool {
	if o != nil && !IsNil(o.Commissions) {
		return true
	}

	return false
}

// SetCommissions gets a reference to the given string and assigns it to the Commissions field.
func (o *ResponsesValidator) SetCommissions(v string) {
	o.Commissions = &v
}

// GetConsAddress returns the ConsAddress field value if set, zero value otherwise.
func (o *ResponsesValidator) GetConsAddress() string {
	if o == nil || IsNil(o.ConsAddress) {
		var ret string
		return ret
	}
	return *o.ConsAddress
}

// GetConsAddressOk returns a tuple with the ConsAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetConsAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ConsAddress) {
		return nil, false
	}
	return o.ConsAddress, true
}

// HasConsAddress returns a boolean if a field has been set.
func (o *ResponsesValidator) HasConsAddress() bool {
	if o != nil && !IsNil(o.ConsAddress) {
		return true
	}

	return false
}

// SetConsAddress gets a reference to the given string and assigns it to the ConsAddress field.
func (o *ResponsesValidator) SetConsAddress(v string) {
	o.ConsAddress = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *ResponsesValidator) GetContacts() string {
	if o == nil || IsNil(o.Contacts) {
		var ret string
		return ret
	}
	return *o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetContactsOk() (*string, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *ResponsesValidator) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given string and assigns it to the Contacts field.
func (o *ResponsesValidator) SetContacts(v string) {
	o.Contacts = &v
}

// GetDelegator returns the Delegator field value if set, zero value otherwise.
func (o *ResponsesValidator) GetDelegator() string {
	if o == nil || IsNil(o.Delegator) {
		var ret string
		return ret
	}
	return *o.Delegator
}

// GetDelegatorOk returns a tuple with the Delegator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetDelegatorOk() (*string, bool) {
	if o == nil || IsNil(o.Delegator) {
		return nil, false
	}
	return o.Delegator, true
}

// HasDelegator returns a boolean if a field has been set.
func (o *ResponsesValidator) HasDelegator() bool {
	if o != nil && !IsNil(o.Delegator) {
		return true
	}

	return false
}

// SetDelegator gets a reference to the given string and assigns it to the Delegator field.
func (o *ResponsesValidator) SetDelegator(v string) {
	o.Delegator = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ResponsesValidator) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ResponsesValidator) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ResponsesValidator) SetDetails(v string) {
	o.Details = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponsesValidator) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponsesValidator) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ResponsesValidator) SetId(v int32) {
	o.Id = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *ResponsesValidator) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *ResponsesValidator) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *ResponsesValidator) SetIdentity(v string) {
	o.Identity = &v
}

// GetJailed returns the Jailed field value if set, zero value otherwise.
func (o *ResponsesValidator) GetJailed() bool {
	if o == nil || IsNil(o.Jailed) {
		var ret bool
		return ret
	}
	return *o.Jailed
}

// GetJailedOk returns a tuple with the Jailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetJailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Jailed) {
		return nil, false
	}
	return o.Jailed, true
}

// HasJailed returns a boolean if a field has been set.
func (o *ResponsesValidator) HasJailed() bool {
	if o != nil && !IsNil(o.Jailed) {
		return true
	}

	return false
}

// SetJailed gets a reference to the given bool and assigns it to the Jailed field.
func (o *ResponsesValidator) SetJailed(v bool) {
	o.Jailed = &v
}

// GetMaxChangeRate returns the MaxChangeRate field value if set, zero value otherwise.
func (o *ResponsesValidator) GetMaxChangeRate() string {
	if o == nil || IsNil(o.MaxChangeRate) {
		var ret string
		return ret
	}
	return *o.MaxChangeRate
}

// GetMaxChangeRateOk returns a tuple with the MaxChangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetMaxChangeRateOk() (*string, bool) {
	if o == nil || IsNil(o.MaxChangeRate) {
		return nil, false
	}
	return o.MaxChangeRate, true
}

// HasMaxChangeRate returns a boolean if a field has been set.
func (o *ResponsesValidator) HasMaxChangeRate() bool {
	if o != nil && !IsNil(o.MaxChangeRate) {
		return true
	}

	return false
}

// SetMaxChangeRate gets a reference to the given string and assigns it to the MaxChangeRate field.
func (o *ResponsesValidator) SetMaxChangeRate(v string) {
	o.MaxChangeRate = &v
}

// GetMaxRate returns the MaxRate field value if set, zero value otherwise.
func (o *ResponsesValidator) GetMaxRate() string {
	if o == nil || IsNil(o.MaxRate) {
		var ret string
		return ret
	}
	return *o.MaxRate
}

// GetMaxRateOk returns a tuple with the MaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetMaxRateOk() (*string, bool) {
	if o == nil || IsNil(o.MaxRate) {
		return nil, false
	}
	return o.MaxRate, true
}

// HasMaxRate returns a boolean if a field has been set.
func (o *ResponsesValidator) HasMaxRate() bool {
	if o != nil && !IsNil(o.MaxRate) {
		return true
	}

	return false
}

// SetMaxRate gets a reference to the given string and assigns it to the MaxRate field.
func (o *ResponsesValidator) SetMaxRate(v string) {
	o.MaxRate = &v
}

// GetMinSelfDelegation returns the MinSelfDelegation field value if set, zero value otherwise.
func (o *ResponsesValidator) GetMinSelfDelegation() string {
	if o == nil || IsNil(o.MinSelfDelegation) {
		var ret string
		return ret
	}
	return *o.MinSelfDelegation
}

// GetMinSelfDelegationOk returns a tuple with the MinSelfDelegation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetMinSelfDelegationOk() (*string, bool) {
	if o == nil || IsNil(o.MinSelfDelegation) {
		return nil, false
	}
	return o.MinSelfDelegation, true
}

// HasMinSelfDelegation returns a boolean if a field has been set.
func (o *ResponsesValidator) HasMinSelfDelegation() bool {
	if o != nil && !IsNil(o.MinSelfDelegation) {
		return true
	}

	return false
}

// SetMinSelfDelegation gets a reference to the given string and assigns it to the MinSelfDelegation field.
func (o *ResponsesValidator) SetMinSelfDelegation(v string) {
	o.MinSelfDelegation = &v
}

// GetMoniker returns the Moniker field value if set, zero value otherwise.
func (o *ResponsesValidator) GetMoniker() string {
	if o == nil || IsNil(o.Moniker) {
		var ret string
		return ret
	}
	return *o.Moniker
}

// GetMonikerOk returns a tuple with the Moniker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetMonikerOk() (*string, bool) {
	if o == nil || IsNil(o.Moniker) {
		return nil, false
	}
	return o.Moniker, true
}

// HasMoniker returns a boolean if a field has been set.
func (o *ResponsesValidator) HasMoniker() bool {
	if o != nil && !IsNil(o.Moniker) {
		return true
	}

	return false
}

// SetMoniker gets a reference to the given string and assigns it to the Moniker field.
func (o *ResponsesValidator) SetMoniker(v string) {
	o.Moniker = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ResponsesValidator) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ResponsesValidator) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *ResponsesValidator) SetRate(v string) {
	o.Rate = &v
}

// GetRewards returns the Rewards field value if set, zero value otherwise.
func (o *ResponsesValidator) GetRewards() string {
	if o == nil || IsNil(o.Rewards) {
		var ret string
		return ret
	}
	return *o.Rewards
}

// GetRewardsOk returns a tuple with the Rewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetRewardsOk() (*string, bool) {
	if o == nil || IsNil(o.Rewards) {
		return nil, false
	}
	return o.Rewards, true
}

// HasRewards returns a boolean if a field has been set.
func (o *ResponsesValidator) HasRewards() bool {
	if o != nil && !IsNil(o.Rewards) {
		return true
	}

	return false
}

// SetRewards gets a reference to the given string and assigns it to the Rewards field.
func (o *ResponsesValidator) SetRewards(v string) {
	o.Rewards = &v
}

// GetStake returns the Stake field value if set, zero value otherwise.
func (o *ResponsesValidator) GetStake() string {
	if o == nil || IsNil(o.Stake) {
		var ret string
		return ret
	}
	return *o.Stake
}

// GetStakeOk returns a tuple with the Stake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetStakeOk() (*string, bool) {
	if o == nil || IsNil(o.Stake) {
		return nil, false
	}
	return o.Stake, true
}

// HasStake returns a boolean if a field has been set.
func (o *ResponsesValidator) HasStake() bool {
	if o != nil && !IsNil(o.Stake) {
		return true
	}

	return false
}

// SetStake gets a reference to the given string and assigns it to the Stake field.
func (o *ResponsesValidator) SetStake(v string) {
	o.Stake = &v
}

// GetVotingPower returns the VotingPower field value if set, zero value otherwise.
func (o *ResponsesValidator) GetVotingPower() string {
	if o == nil || IsNil(o.VotingPower) {
		var ret string
		return ret
	}
	return *o.VotingPower
}

// GetVotingPowerOk returns a tuple with the VotingPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetVotingPowerOk() (*string, bool) {
	if o == nil || IsNil(o.VotingPower) {
		return nil, false
	}
	return o.VotingPower, true
}

// HasVotingPower returns a boolean if a field has been set.
func (o *ResponsesValidator) HasVotingPower() bool {
	if o != nil && !IsNil(o.VotingPower) {
		return true
	}

	return false
}

// SetVotingPower gets a reference to the given string and assigns it to the VotingPower field.
func (o *ResponsesValidator) SetVotingPower(v string) {
	o.VotingPower = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ResponsesValidator) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesValidator) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ResponsesValidator) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *ResponsesValidator) SetWebsite(v string) {
	o.Website = &v
}

func (o ResponsesValidator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesValidator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Commissions) {
		toSerialize["commissions"] = o.Commissions
	}
	if !IsNil(o.ConsAddress) {
		toSerialize["cons_address"] = o.ConsAddress
	}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.Delegator) {
		toSerialize["delegator"] = o.Delegator
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.Jailed) {
		toSerialize["jailed"] = o.Jailed
	}
	if !IsNil(o.MaxChangeRate) {
		toSerialize["max_change_rate"] = o.MaxChangeRate
	}
	if !IsNil(o.MaxRate) {
		toSerialize["max_rate"] = o.MaxRate
	}
	if !IsNil(o.MinSelfDelegation) {
		toSerialize["min_self_delegation"] = o.MinSelfDelegation
	}
	if !IsNil(o.Moniker) {
		toSerialize["moniker"] = o.Moniker
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Rewards) {
		toSerialize["rewards"] = o.Rewards
	}
	if !IsNil(o.Stake) {
		toSerialize["stake"] = o.Stake
	}
	if !IsNil(o.VotingPower) {
		toSerialize["voting_power"] = o.VotingPower
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	return toSerialize, nil
}

type NullableResponsesValidator struct {
	value *ResponsesValidator
	isSet bool
}

func (v NullableResponsesValidator) Get() *ResponsesValidator {
	return v.value
}

func (v *NullableResponsesValidator) Set(val *ResponsesValidator) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesValidator) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesValidator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesValidator(val *ResponsesValidator) *NullableResponsesValidator {
	return &NullableResponsesValidator{value: val, isSet: true}
}

func (v NullableResponsesValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesValidator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


