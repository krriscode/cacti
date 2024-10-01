/*
Hyperledger Cactus Plugin - HTLC ETH BESU ERC20

Allows Cactus nodes to interact with HTLC contracts with ERC-20 Tokens

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-htlc-eth-besu-erc20

import (
	"encoding/json"
	"fmt"
)

// NewContractRequestGas - struct for NewContractRequestGas
type NewContractRequestGas struct {
	Float32 *float32
	String *string
}

// float32AsNewContractRequestGas is a convenience function that returns float32 wrapped in NewContractRequestGas
func Float32AsNewContractRequestGas(v *float32) NewContractRequestGas {
	return NewContractRequestGas{
		Float32: v,
	}
}

// stringAsNewContractRequestGas is a convenience function that returns string wrapped in NewContractRequestGas
func StringAsNewContractRequestGas(v *string) NewContractRequestGas {
	return NewContractRequestGas{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NewContractRequestGas) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NewContractRequestGas)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NewContractRequestGas)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NewContractRequestGas) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NewContractRequestGas) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableNewContractRequestGas struct {
	value *NewContractRequestGas
	isSet bool
}

func (v NullableNewContractRequestGas) Get() *NewContractRequestGas {
	return v.value
}

func (v *NullableNewContractRequestGas) Set(val *NewContractRequestGas) {
	v.value = val
	v.isSet = true
}

func (v NullableNewContractRequestGas) IsSet() bool {
	return v.isSet
}

func (v *NullableNewContractRequestGas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewContractRequestGas(val *NewContractRequestGas) *NullableNewContractRequestGas {
	return &NullableNewContractRequestGas{value: val, isSet: true}
}

func (v NullableNewContractRequestGas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewContractRequestGas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


