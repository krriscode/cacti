/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
	"fmt"
)

// TransactionStatusV1 Status of Iroha V2 transaction.
type TransactionStatusV1 string

// List of TransactionStatusV1
const (
	Submitted TransactionStatusV1 = "submitted"
	Committed TransactionStatusV1 = "committed"
	Rejected TransactionStatusV1 = "rejected"
)

// All allowed values of TransactionStatusV1 enum
var AllowedTransactionStatusV1EnumValues = []TransactionStatusV1{
	"submitted",
	"committed",
	"rejected",
}

func (v *TransactionStatusV1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionStatusV1(value)
	for _, existing := range AllowedTransactionStatusV1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionStatusV1", value)
}

// NewTransactionStatusV1FromValue returns a pointer to a valid TransactionStatusV1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionStatusV1FromValue(v string) (*TransactionStatusV1, error) {
	ev := TransactionStatusV1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionStatusV1: valid values are %v", v, AllowedTransactionStatusV1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionStatusV1) IsValid() bool {
	for _, existing := range AllowedTransactionStatusV1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionStatusV1 value
func (v TransactionStatusV1) Ptr() *TransactionStatusV1 {
	return &v
}

type NullableTransactionStatusV1 struct {
	value *TransactionStatusV1
	isSet bool
}

func (v NullableTransactionStatusV1) Get() *TransactionStatusV1 {
	return v.value
}

func (v *NullableTransactionStatusV1) Set(val *TransactionStatusV1) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStatusV1) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStatusV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStatusV1(val *TransactionStatusV1) *NullableTransactionStatusV1 {
	return &NullableTransactionStatusV1{value: val, isSet: true}
}

func (v NullableTransactionStatusV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStatusV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

