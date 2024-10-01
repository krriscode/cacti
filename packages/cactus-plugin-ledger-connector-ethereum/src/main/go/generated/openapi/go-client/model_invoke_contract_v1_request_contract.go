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

// InvokeContractV1RequestContract - struct for InvokeContractV1RequestContract
type InvokeContractV1RequestContract struct {
	ContractKeychainDefinition *ContractKeychainDefinition
	DeployedContractJsonDefinition *DeployedContractJsonDefinition
}

// ContractKeychainDefinitionAsInvokeContractV1RequestContract is a convenience function that returns ContractKeychainDefinition wrapped in InvokeContractV1RequestContract
func ContractKeychainDefinitionAsInvokeContractV1RequestContract(v *ContractKeychainDefinition) InvokeContractV1RequestContract {
	return InvokeContractV1RequestContract{
		ContractKeychainDefinition: v,
	}
}

// DeployedContractJsonDefinitionAsInvokeContractV1RequestContract is a convenience function that returns DeployedContractJsonDefinition wrapped in InvokeContractV1RequestContract
func DeployedContractJsonDefinitionAsInvokeContractV1RequestContract(v *DeployedContractJsonDefinition) InvokeContractV1RequestContract {
	return InvokeContractV1RequestContract{
		DeployedContractJsonDefinition: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InvokeContractV1RequestContract) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContractKeychainDefinition
	err = newStrictDecoder(data).Decode(&dst.ContractKeychainDefinition)
	if err == nil {
		jsonContractKeychainDefinition, _ := json.Marshal(dst.ContractKeychainDefinition)
		if string(jsonContractKeychainDefinition) == "{}" { // empty struct
			dst.ContractKeychainDefinition = nil
		} else {
			match++
		}
	} else {
		dst.ContractKeychainDefinition = nil
	}

	// try to unmarshal data into DeployedContractJsonDefinition
	err = newStrictDecoder(data).Decode(&dst.DeployedContractJsonDefinition)
	if err == nil {
		jsonDeployedContractJsonDefinition, _ := json.Marshal(dst.DeployedContractJsonDefinition)
		if string(jsonDeployedContractJsonDefinition) == "{}" { // empty struct
			dst.DeployedContractJsonDefinition = nil
		} else {
			match++
		}
	} else {
		dst.DeployedContractJsonDefinition = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContractKeychainDefinition = nil
		dst.DeployedContractJsonDefinition = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InvokeContractV1RequestContract)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InvokeContractV1RequestContract)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InvokeContractV1RequestContract) MarshalJSON() ([]byte, error) {
	if src.ContractKeychainDefinition != nil {
		return json.Marshal(&src.ContractKeychainDefinition)
	}

	if src.DeployedContractJsonDefinition != nil {
		return json.Marshal(&src.DeployedContractJsonDefinition)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InvokeContractV1RequestContract) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ContractKeychainDefinition != nil {
		return obj.ContractKeychainDefinition
	}

	if obj.DeployedContractJsonDefinition != nil {
		return obj.DeployedContractJsonDefinition
	}

	// all schemas are nil
	return nil
}

type NullableInvokeContractV1RequestContract struct {
	value *InvokeContractV1RequestContract
	isSet bool
}

func (v NullableInvokeContractV1RequestContract) Get() *InvokeContractV1RequestContract {
	return v.value
}

func (v *NullableInvokeContractV1RequestContract) Set(val *InvokeContractV1RequestContract) {
	v.value = val
	v.isSet = true
}

func (v NullableInvokeContractV1RequestContract) IsSet() bool {
	return v.isSet
}

func (v *NullableInvokeContractV1RequestContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvokeContractV1RequestContract(val *InvokeContractV1RequestContract) *NullableInvokeContractV1RequestContract {
	return &NullableInvokeContractV1RequestContract{value: val, isSet: true}
}

func (v NullableInvokeContractV1RequestContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvokeContractV1RequestContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


