/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
)

// checks if the ContractJsonDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractJsonDefinition{}

// ContractJsonDefinition struct for ContractJsonDefinition
type ContractJsonDefinition struct {
	ContractJSON ContractJSON `json:"contractJSON"`
}

// NewContractJsonDefinition instantiates a new ContractJsonDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractJsonDefinition(contractJSON ContractJSON) *ContractJsonDefinition {
	this := ContractJsonDefinition{}
	this.ContractJSON = contractJSON
	return &this
}

// NewContractJsonDefinitionWithDefaults instantiates a new ContractJsonDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractJsonDefinitionWithDefaults() *ContractJsonDefinition {
	this := ContractJsonDefinition{}
	return &this
}

// GetContractJSON returns the ContractJSON field value
func (o *ContractJsonDefinition) GetContractJSON() ContractJSON {
	if o == nil {
		var ret ContractJSON
		return ret
	}

	return o.ContractJSON
}

// GetContractJSONOk returns a tuple with the ContractJSON field value
// and a boolean to check if the value has been set.
func (o *ContractJsonDefinition) GetContractJSONOk() (*ContractJSON, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractJSON, true
}

// SetContractJSON sets field value
func (o *ContractJsonDefinition) SetContractJSON(v ContractJSON) {
	o.ContractJSON = v
}

func (o ContractJsonDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractJsonDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contractJSON"] = o.ContractJSON
	return toSerialize, nil
}

type NullableContractJsonDefinition struct {
	value *ContractJsonDefinition
	isSet bool
}

func (v NullableContractJsonDefinition) Get() *ContractJsonDefinition {
	return v.value
}

func (v *NullableContractJsonDefinition) Set(val *ContractJsonDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableContractJsonDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableContractJsonDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractJsonDefinition(val *ContractJsonDefinition) *NullableContractJsonDefinition {
	return &NullableContractJsonDefinition{value: val, isSet: true}
}

func (v NullableContractJsonDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractJsonDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


