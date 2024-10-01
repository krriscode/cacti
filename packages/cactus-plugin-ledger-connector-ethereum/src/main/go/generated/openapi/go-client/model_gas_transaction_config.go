/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
	"fmt"
)

// GasTransactionConfig - Transaction gas settings.
type GasTransactionConfig struct {
	GasTransactionConfigEIP1559 *GasTransactionConfigEIP1559
	GasTransactionConfigLegacy *GasTransactionConfigLegacy
}

// GasTransactionConfigEIP1559AsGasTransactionConfig is a convenience function that returns GasTransactionConfigEIP1559 wrapped in GasTransactionConfig
func GasTransactionConfigEIP1559AsGasTransactionConfig(v *GasTransactionConfigEIP1559) GasTransactionConfig {
	return GasTransactionConfig{
		GasTransactionConfigEIP1559: v,
	}
}

// GasTransactionConfigLegacyAsGasTransactionConfig is a convenience function that returns GasTransactionConfigLegacy wrapped in GasTransactionConfig
func GasTransactionConfigLegacyAsGasTransactionConfig(v *GasTransactionConfigLegacy) GasTransactionConfig {
	return GasTransactionConfig{
		GasTransactionConfigLegacy: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GasTransactionConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GasTransactionConfigEIP1559
	err = newStrictDecoder(data).Decode(&dst.GasTransactionConfigEIP1559)
	if err == nil {
		jsonGasTransactionConfigEIP1559, _ := json.Marshal(dst.GasTransactionConfigEIP1559)
		if string(jsonGasTransactionConfigEIP1559) == "{}" { // empty struct
			dst.GasTransactionConfigEIP1559 = nil
		} else {
			match++
		}
	} else {
		dst.GasTransactionConfigEIP1559 = nil
	}

	// try to unmarshal data into GasTransactionConfigLegacy
	err = newStrictDecoder(data).Decode(&dst.GasTransactionConfigLegacy)
	if err == nil {
		jsonGasTransactionConfigLegacy, _ := json.Marshal(dst.GasTransactionConfigLegacy)
		if string(jsonGasTransactionConfigLegacy) == "{}" { // empty struct
			dst.GasTransactionConfigLegacy = nil
		} else {
			match++
		}
	} else {
		dst.GasTransactionConfigLegacy = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GasTransactionConfigEIP1559 = nil
		dst.GasTransactionConfigLegacy = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GasTransactionConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GasTransactionConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GasTransactionConfig) MarshalJSON() ([]byte, error) {
	if src.GasTransactionConfigEIP1559 != nil {
		return json.Marshal(&src.GasTransactionConfigEIP1559)
	}

	if src.GasTransactionConfigLegacy != nil {
		return json.Marshal(&src.GasTransactionConfigLegacy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GasTransactionConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GasTransactionConfigEIP1559 != nil {
		return obj.GasTransactionConfigEIP1559
	}

	if obj.GasTransactionConfigLegacy != nil {
		return obj.GasTransactionConfigLegacy
	}

	// all schemas are nil
	return nil
}

type NullableGasTransactionConfig struct {
	value *GasTransactionConfig
	isSet bool
}

func (v NullableGasTransactionConfig) Get() *GasTransactionConfig {
	return v.value
}

func (v *NullableGasTransactionConfig) Set(val *GasTransactionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGasTransactionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGasTransactionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGasTransactionConfig(val *GasTransactionConfig) *NullableGasTransactionConfig {
	return &NullableGasTransactionConfig{value: val, isSet: true}
}

func (v NullableGasTransactionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGasTransactionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


