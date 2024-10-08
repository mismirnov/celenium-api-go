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
	"fmt"
)

// TypesMsgAddressType the model 'TypesMsgAddressType'
type TypesMsgAddressType string

// List of types.MsgAddressType
const (
	MsgAddressTypeValidator TypesMsgAddressType = "validator"
	MsgAddressTypeDelegator TypesMsgAddressType = "delegator"
	MsgAddressTypeDepositor TypesMsgAddressType = "depositor"
	MsgAddressTypeValidatorSrc TypesMsgAddressType = "validatorSrc"
	MsgAddressTypeValidatorDst TypesMsgAddressType = "validatorDst"
	MsgAddressTypeFromAddress TypesMsgAddressType = "fromAddress"
	MsgAddressTypeToAddress TypesMsgAddressType = "toAddress"
	MsgAddressTypeInput TypesMsgAddressType = "input"
	MsgAddressTypeOutput TypesMsgAddressType = "output"
	MsgAddressTypeGrantee TypesMsgAddressType = "grantee"
	MsgAddressTypeGranter TypesMsgAddressType = "granter"
	MsgAddressTypeSigner TypesMsgAddressType = "signer"
	MsgAddressTypeWithdraw TypesMsgAddressType = "withdraw"
	MsgAddressTypeVoter TypesMsgAddressType = "voter"
	MsgAddressTypeProposer TypesMsgAddressType = "proposer"
	MsgAddressTypeAuthority TypesMsgAddressType = "authority"
	MsgAddressTypeSender TypesMsgAddressType = "sender"
	MsgAddressTypeReceiver TypesMsgAddressType = "receiver"
	MsgAddressTypeSubmitter TypesMsgAddressType = "submitter"
	MsgAddressTypeAdmin TypesMsgAddressType = "admin"
	MsgAddressTypeNewAdmin TypesMsgAddressType = "newAdmin"
	MsgAddressTypeGroupPolicyAddress TypesMsgAddressType = "groupPolicyAddress"
	MsgAddressTypeExecutor TypesMsgAddressType = "executor"
	MsgAddressTypeGroupMember TypesMsgAddressType = "groupMember"
	MsgAddressTypeOwner TypesMsgAddressType = "owner"
	MsgAddressTypeRelayer TypesMsgAddressType = "relayer"
	MsgAddressTypePayee TypesMsgAddressType = "payee"
)

// All allowed values of TypesMsgAddressType enum
var AllowedTypesMsgAddressTypeEnumValues = []TypesMsgAddressType{
	"validator",
	"delegator",
	"depositor",
	"validatorSrc",
	"validatorDst",
	"fromAddress",
	"toAddress",
	"input",
	"output",
	"grantee",
	"granter",
	"signer",
	"withdraw",
	"voter",
	"proposer",
	"authority",
	"sender",
	"receiver",
	"submitter",
	"admin",
	"newAdmin",
	"groupPolicyAddress",
	"executor",
	"groupMember",
	"owner",
	"relayer",
	"payee",
}

func (v *TypesMsgAddressType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesMsgAddressType(value)
	for _, existing := range AllowedTypesMsgAddressTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesMsgAddressType", value)
}

// NewTypesMsgAddressTypeFromValue returns a pointer to a valid TypesMsgAddressType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesMsgAddressTypeFromValue(v string) (*TypesMsgAddressType, error) {
	ev := TypesMsgAddressType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesMsgAddressType: valid values are %v", v, AllowedTypesMsgAddressTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesMsgAddressType) IsValid() bool {
	for _, existing := range AllowedTypesMsgAddressTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.MsgAddressType value
func (v TypesMsgAddressType) Ptr() *TypesMsgAddressType {
	return &v
}

type NullableTypesMsgAddressType struct {
	value *TypesMsgAddressType
	isSet bool
}

func (v NullableTypesMsgAddressType) Get() *TypesMsgAddressType {
	return v.value
}

func (v *NullableTypesMsgAddressType) Set(val *TypesMsgAddressType) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesMsgAddressType) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesMsgAddressType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesMsgAddressType(val *TypesMsgAddressType) *NullableTypesMsgAddressType {
	return &NullableTypesMsgAddressType{value: val, isSet: true}
}

func (v NullableTypesMsgAddressType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesMsgAddressType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

