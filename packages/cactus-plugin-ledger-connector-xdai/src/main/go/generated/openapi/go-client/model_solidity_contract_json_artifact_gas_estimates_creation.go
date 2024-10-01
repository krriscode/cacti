/*
Hyperledger Cactus Plugin - Connector Xdai

Can perform basic tasks on a Xdai ledger

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-xdai

import (
	"encoding/json"
)

// checks if the SolidityContractJsonArtifactGasEstimatesCreation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SolidityContractJsonArtifactGasEstimatesCreation{}

// SolidityContractJsonArtifactGasEstimatesCreation struct for SolidityContractJsonArtifactGasEstimatesCreation
type SolidityContractJsonArtifactGasEstimatesCreation struct {
	CodeDepositCost *string `json:"codeDepositCost,omitempty"`
	ExecutionCost *string `json:"executionCost,omitempty"`
	TotalCost *string `json:"totalCost,omitempty"`
}

// NewSolidityContractJsonArtifactGasEstimatesCreation instantiates a new SolidityContractJsonArtifactGasEstimatesCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolidityContractJsonArtifactGasEstimatesCreation() *SolidityContractJsonArtifactGasEstimatesCreation {
	this := SolidityContractJsonArtifactGasEstimatesCreation{}
	return &this
}

// NewSolidityContractJsonArtifactGasEstimatesCreationWithDefaults instantiates a new SolidityContractJsonArtifactGasEstimatesCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolidityContractJsonArtifactGasEstimatesCreationWithDefaults() *SolidityContractJsonArtifactGasEstimatesCreation {
	this := SolidityContractJsonArtifactGasEstimatesCreation{}
	return &this
}

// GetCodeDepositCost returns the CodeDepositCost field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) GetCodeDepositCost() string {
	if o == nil || IsNil(o.CodeDepositCost) {
		var ret string
		return ret
	}
	return *o.CodeDepositCost
}

// GetCodeDepositCostOk returns a tuple with the CodeDepositCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) GetCodeDepositCostOk() (*string, bool) {
	if o == nil || IsNil(o.CodeDepositCost) {
		return nil, false
	}
	return o.CodeDepositCost, true
}

// HasCodeDepositCost returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) HasCodeDepositCost() bool {
	if o != nil && !IsNil(o.CodeDepositCost) {
		return true
	}

	return false
}

// SetCodeDepositCost gets a reference to the given string and assigns it to the CodeDepositCost field.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) SetCodeDepositCost(v string) {
	o.CodeDepositCost = &v
}

// GetExecutionCost returns the ExecutionCost field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) GetExecutionCost() string {
	if o == nil || IsNil(o.ExecutionCost) {
		var ret string
		return ret
	}
	return *o.ExecutionCost
}

// GetExecutionCostOk returns a tuple with the ExecutionCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) GetExecutionCostOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionCost) {
		return nil, false
	}
	return o.ExecutionCost, true
}

// HasExecutionCost returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) HasExecutionCost() bool {
	if o != nil && !IsNil(o.ExecutionCost) {
		return true
	}

	return false
}

// SetExecutionCost gets a reference to the given string and assigns it to the ExecutionCost field.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) SetExecutionCost(v string) {
	o.ExecutionCost = &v
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) GetTotalCost() string {
	if o == nil || IsNil(o.TotalCost) {
		var ret string
		return ret
	}
	return *o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) GetTotalCostOk() (*string, bool) {
	if o == nil || IsNil(o.TotalCost) {
		return nil, false
	}
	return o.TotalCost, true
}

// HasTotalCost returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) HasTotalCost() bool {
	if o != nil && !IsNil(o.TotalCost) {
		return true
	}

	return false
}

// SetTotalCost gets a reference to the given string and assigns it to the TotalCost field.
func (o *SolidityContractJsonArtifactGasEstimatesCreation) SetTotalCost(v string) {
	o.TotalCost = &v
}

func (o SolidityContractJsonArtifactGasEstimatesCreation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SolidityContractJsonArtifactGasEstimatesCreation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodeDepositCost) {
		toSerialize["codeDepositCost"] = o.CodeDepositCost
	}
	if !IsNil(o.ExecutionCost) {
		toSerialize["executionCost"] = o.ExecutionCost
	}
	if !IsNil(o.TotalCost) {
		toSerialize["totalCost"] = o.TotalCost
	}
	return toSerialize, nil
}

type NullableSolidityContractJsonArtifactGasEstimatesCreation struct {
	value *SolidityContractJsonArtifactGasEstimatesCreation
	isSet bool
}

func (v NullableSolidityContractJsonArtifactGasEstimatesCreation) Get() *SolidityContractJsonArtifactGasEstimatesCreation {
	return v.value
}

func (v *NullableSolidityContractJsonArtifactGasEstimatesCreation) Set(val *SolidityContractJsonArtifactGasEstimatesCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableSolidityContractJsonArtifactGasEstimatesCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableSolidityContractJsonArtifactGasEstimatesCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolidityContractJsonArtifactGasEstimatesCreation(val *SolidityContractJsonArtifactGasEstimatesCreation) *NullableSolidityContractJsonArtifactGasEstimatesCreation {
	return &NullableSolidityContractJsonArtifactGasEstimatesCreation{value: val, isSet: true}
}

func (v NullableSolidityContractJsonArtifactGasEstimatesCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolidityContractJsonArtifactGasEstimatesCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


