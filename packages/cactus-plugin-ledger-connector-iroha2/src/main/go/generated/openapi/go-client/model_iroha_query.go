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

// IrohaQuery Command names that correspond to Iroha queries (https://hyperledger.github.io/iroha-2-docs/guide/advanced/queries.html)
type IrohaQuery string

// List of IrohaQuery
const (
	FindAllDomains IrohaQuery = "findAllDomains"
	FindDomainById IrohaQuery = "findDomainById"
	FindAssetDefinitionById IrohaQuery = "findAssetDefinitionById"
	FindAllAssetsDefinitions IrohaQuery = "findAllAssetsDefinitions"
	FindAssetById IrohaQuery = "findAssetById"
	FindAllAssets IrohaQuery = "findAllAssets"
	FindAllPeers IrohaQuery = "findAllPeers"
	FindAllBlocks IrohaQuery = "findAllBlocks"
	FindAccountById IrohaQuery = "findAccountById"
	FindAllAccounts IrohaQuery = "findAllAccounts"
	FindAllTransactions IrohaQuery = "findAllTransactions"
	FindTransactionByHash IrohaQuery = "findTransactionByHash"
)

// All allowed values of IrohaQuery enum
var AllowedIrohaQueryEnumValues = []IrohaQuery{
	"findAllDomains",
	"findDomainById",
	"findAssetDefinitionById",
	"findAllAssetsDefinitions",
	"findAssetById",
	"findAllAssets",
	"findAllPeers",
	"findAllBlocks",
	"findAccountById",
	"findAllAccounts",
	"findAllTransactions",
	"findTransactionByHash",
}

func (v *IrohaQuery) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IrohaQuery(value)
	for _, existing := range AllowedIrohaQueryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IrohaQuery", value)
}

// NewIrohaQueryFromValue returns a pointer to a valid IrohaQuery
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIrohaQueryFromValue(v string) (*IrohaQuery, error) {
	ev := IrohaQuery(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IrohaQuery: valid values are %v", v, AllowedIrohaQueryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IrohaQuery) IsValid() bool {
	for _, existing := range AllowedIrohaQueryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IrohaQuery value
func (v IrohaQuery) Ptr() *IrohaQuery {
	return &v
}

type NullableIrohaQuery struct {
	value *IrohaQuery
	isSet bool
}

func (v NullableIrohaQuery) Get() *IrohaQuery {
	return v.value
}

func (v *NullableIrohaQuery) Set(val *IrohaQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableIrohaQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableIrohaQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIrohaQuery(val *IrohaQuery) *NullableIrohaQuery {
	return &NullableIrohaQuery{value: val, isSet: true}
}

func (v NullableIrohaQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIrohaQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

