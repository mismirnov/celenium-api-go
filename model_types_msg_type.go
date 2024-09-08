/*
Celenium API

This is docs of Celenium API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"encoding/json"
	"fmt"
)

// TypesMsgType the model 'TypesMsgType'
type TypesMsgType string

// List of types.MsgType
const (
	MsgUnknown TypesMsgType = "MsgUnknown"
	MsgSetWithdrawAddress TypesMsgType = "MsgSetWithdrawAddress"
	MsgWithdrawDelegatorReward TypesMsgType = "MsgWithdrawDelegatorReward"
	MsgWithdrawValidatorCommission TypesMsgType = "MsgWithdrawValidatorCommission"
	MsgFundCommunityPool TypesMsgType = "MsgFundCommunityPool"
	MsgCreateValidator TypesMsgType = "MsgCreateValidator"
	MsgEditValidator TypesMsgType = "MsgEditValidator"
	MsgDelegate TypesMsgType = "MsgDelegate"
	MsgBeginRedelegate TypesMsgType = "MsgBeginRedelegate"
	MsgUndelegate TypesMsgType = "MsgUndelegate"
	MsgCancelUnbondingDelegation TypesMsgType = "MsgCancelUnbondingDelegation"
	MsgUnjail TypesMsgType = "MsgUnjail"
	MsgSend TypesMsgType = "MsgSend"
	MsgMultiSend TypesMsgType = "MsgMultiSend"
	MsgCreateVestingAccount TypesMsgType = "MsgCreateVestingAccount"
	MsgCreatePermanentLockedAccount TypesMsgType = "MsgCreatePermanentLockedAccount"
	MsgCreatePeriodicVestingAccount TypesMsgType = "MsgCreatePeriodicVestingAccount"
	MsgPayForBlobs TypesMsgType = "MsgPayForBlobs"
	MsgGrant TypesMsgType = "MsgGrant"
	MsgExec TypesMsgType = "MsgExec"
	MsgRevoke TypesMsgType = "MsgRevoke"
	MsgGrantAllowance TypesMsgType = "MsgGrantAllowance"
	MsgRevokeAllowance TypesMsgType = "MsgRevokeAllowance"
	MsgRegisterEVMAddress TypesMsgType = "MsgRegisterEVMAddress"
	MsgSubmitProposal TypesMsgType = "MsgSubmitProposal"
	MsgExecLegacyContent TypesMsgType = "MsgExecLegacyContent"
	MsgVote TypesMsgType = "MsgVote"
	MsgVoteWeighted TypesMsgType = "MsgVoteWeighted"
	MsgDeposit TypesMsgType = "MsgDeposit"
	IBCTransfer TypesMsgType = "IBCTransfer"
	MsgVerifyInvariant TypesMsgType = "MsgVerifyInvariant"
	MsgSubmitEvidence TypesMsgType = "MsgSubmitEvidence"
	MsgSendNFT TypesMsgType = "MsgSendNFT"
	MsgCreateGroup TypesMsgType = "MsgCreateGroup"
	MsgUpdateGroupMembers TypesMsgType = "MsgUpdateGroupMembers"
	MsgUpdateGroupAdmin TypesMsgType = "MsgUpdateGroupAdmin"
	MsgUpdateGroupMetadata TypesMsgType = "MsgUpdateGroupMetadata"
	MsgCreateGroupPolicy TypesMsgType = "MsgCreateGroupPolicy"
	MsgUpdateGroupPolicyAdmin TypesMsgType = "MsgUpdateGroupPolicyAdmin"
	MsgCreateGroupWithPolicy TypesMsgType = "MsgCreateGroupWithPolicy"
	MsgUpdateGroupPolicyDecisionPolicy TypesMsgType = "MsgUpdateGroupPolicyDecisionPolicy"
	MsgUpdateGroupPolicyMetadata TypesMsgType = "MsgUpdateGroupPolicyMetadata"
	MsgSubmitProposalGroup TypesMsgType = "MsgSubmitProposalGroup"
	MsgWithdrawProposal TypesMsgType = "MsgWithdrawProposal"
	MsgVoteGroup TypesMsgType = "MsgVoteGroup"
	MsgExecGroup TypesMsgType = "MsgExecGroup"
	MsgLeaveGroup TypesMsgType = "MsgLeaveGroup"
	MsgSoftwareUpgrade TypesMsgType = "MsgSoftwareUpgrade"
	MsgCancelUpgrade TypesMsgType = "MsgCancelUpgrade"
	MsgRegisterInterchainAccount TypesMsgType = "MsgRegisterInterchainAccount"
	MsgSendTx TypesMsgType = "MsgSendTx"
	MsgRegisterPayee TypesMsgType = "MsgRegisterPayee"
	MsgRegisterCounterpartyPayee TypesMsgType = "MsgRegisterCounterpartyPayee"
	MsgPayPacketFee TypesMsgType = "MsgPayPacketFee"
	MsgPayPacketFeeAsync TypesMsgType = "MsgPayPacketFeeAsync"
	MsgTransfer TypesMsgType = "MsgTransfer"
	MsgCreateClient TypesMsgType = "MsgCreateClient"
	MsgUpdateClient TypesMsgType = "MsgUpdateClient"
	MsgUpgradeClient TypesMsgType = "MsgUpgradeClient"
	MsgSubmitMisbehaviour TypesMsgType = "MsgSubmitMisbehaviour"
	MsgConnectionOpenInit TypesMsgType = "MsgConnectionOpenInit"
	MsgConnectionOpenTry TypesMsgType = "MsgConnectionOpenTry"
	MsgConnectionOpenAck TypesMsgType = "MsgConnectionOpenAck"
	MsgConnectionOpenConfirm TypesMsgType = "MsgConnectionOpenConfirm"
	MsgChannelOpenInit TypesMsgType = "MsgChannelOpenInit"
	MsgChannelOpenTry TypesMsgType = "MsgChannelOpenTry"
	MsgChannelOpenAck TypesMsgType = "MsgChannelOpenAck"
	MsgChannelOpenConfirm TypesMsgType = "MsgChannelOpenConfirm"
	MsgChannelCloseInit TypesMsgType = "MsgChannelCloseInit"
	MsgChannelCloseConfirm TypesMsgType = "MsgChannelCloseConfirm"
	MsgRecvPacket TypesMsgType = "MsgRecvPacket"
	MsgTimeout TypesMsgType = "MsgTimeout"
	MsgTimeoutOnClose TypesMsgType = "MsgTimeoutOnClose"
	MsgAcknowledgement TypesMsgType = "MsgAcknowledgement"
)

// All allowed values of TypesMsgType enum
var AllowedTypesMsgTypeEnumValues = []TypesMsgType{
	"MsgUnknown",
	"MsgSetWithdrawAddress",
	"MsgWithdrawDelegatorReward",
	"MsgWithdrawValidatorCommission",
	"MsgFundCommunityPool",
	"MsgCreateValidator",
	"MsgEditValidator",
	"MsgDelegate",
	"MsgBeginRedelegate",
	"MsgUndelegate",
	"MsgCancelUnbondingDelegation",
	"MsgUnjail",
	"MsgSend",
	"MsgMultiSend",
	"MsgCreateVestingAccount",
	"MsgCreatePermanentLockedAccount",
	"MsgCreatePeriodicVestingAccount",
	"MsgPayForBlobs",
	"MsgGrant",
	"MsgExec",
	"MsgRevoke",
	"MsgGrantAllowance",
	"MsgRevokeAllowance",
	"MsgRegisterEVMAddress",
	"MsgSubmitProposal",
	"MsgExecLegacyContent",
	"MsgVote",
	"MsgVoteWeighted",
	"MsgDeposit",
	"IBCTransfer",
	"MsgVerifyInvariant",
	"MsgSubmitEvidence",
	"MsgSendNFT",
	"MsgCreateGroup",
	"MsgUpdateGroupMembers",
	"MsgUpdateGroupAdmin",
	"MsgUpdateGroupMetadata",
	"MsgCreateGroupPolicy",
	"MsgUpdateGroupPolicyAdmin",
	"MsgCreateGroupWithPolicy",
	"MsgUpdateGroupPolicyDecisionPolicy",
	"MsgUpdateGroupPolicyMetadata",
	"MsgSubmitProposalGroup",
	"MsgWithdrawProposal",
	"MsgVoteGroup",
	"MsgExecGroup",
	"MsgLeaveGroup",
	"MsgSoftwareUpgrade",
	"MsgCancelUpgrade",
	"MsgRegisterInterchainAccount",
	"MsgSendTx",
	"MsgRegisterPayee",
	"MsgRegisterCounterpartyPayee",
	"MsgPayPacketFee",
	"MsgPayPacketFeeAsync",
	"MsgTransfer",
	"MsgCreateClient",
	"MsgUpdateClient",
	"MsgUpgradeClient",
	"MsgSubmitMisbehaviour",
	"MsgConnectionOpenInit",
	"MsgConnectionOpenTry",
	"MsgConnectionOpenAck",
	"MsgConnectionOpenConfirm",
	"MsgChannelOpenInit",
	"MsgChannelOpenTry",
	"MsgChannelOpenAck",
	"MsgChannelOpenConfirm",
	"MsgChannelCloseInit",
	"MsgChannelCloseConfirm",
	"MsgRecvPacket",
	"MsgTimeout",
	"MsgTimeoutOnClose",
	"MsgAcknowledgement",
}

func (v *TypesMsgType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesMsgType(value)
	for _, existing := range AllowedTypesMsgTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesMsgType", value)
}

// NewTypesMsgTypeFromValue returns a pointer to a valid TypesMsgType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesMsgTypeFromValue(v string) (*TypesMsgType, error) {
	ev := TypesMsgType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesMsgType: valid values are %v", v, AllowedTypesMsgTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesMsgType) IsValid() bool {
	for _, existing := range AllowedTypesMsgTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.MsgType value
func (v TypesMsgType) Ptr() *TypesMsgType {
	return &v
}

type NullableTypesMsgType struct {
	value *TypesMsgType
	isSet bool
}

func (v NullableTypesMsgType) Get() *TypesMsgType {
	return v.value
}

func (v *NullableTypesMsgType) Set(val *TypesMsgType) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesMsgType) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesMsgType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesMsgType(val *TypesMsgType) *NullableTypesMsgType {
	return &NullableTypesMsgType{value: val, isSet: true}
}

func (v NullableTypesMsgType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesMsgType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

