/*
Hyperledger Cactus Plugin - Connector Polkadot

Can perform basic tasks on a Polkadot parachain

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-polkadot

import (
	"encoding/json"
)

// checks if the RunTransactionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunTransactionRequest{}

// RunTransactionRequest struct for RunTransactionRequest
type RunTransactionRequest struct {
	Web3SigningCredential Web3SigningCredential `json:"web3SigningCredential"`
	TransactionConfig PolkadotTransactionConfig `json:"transactionConfig"`
}

// NewRunTransactionRequest instantiates a new RunTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunTransactionRequest(web3SigningCredential Web3SigningCredential, transactionConfig PolkadotTransactionConfig) *RunTransactionRequest {
	this := RunTransactionRequest{}
	this.Web3SigningCredential = web3SigningCredential
	this.TransactionConfig = transactionConfig
	return &this
}

// NewRunTransactionRequestWithDefaults instantiates a new RunTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunTransactionRequestWithDefaults() *RunTransactionRequest {
	this := RunTransactionRequest{}
	return &this
}

// GetWeb3SigningCredential returns the Web3SigningCredential field value
func (o *RunTransactionRequest) GetWeb3SigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.Web3SigningCredential
}

// GetWeb3SigningCredentialOk returns a tuple with the Web3SigningCredential field value
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetWeb3SigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Web3SigningCredential, true
}

// SetWeb3SigningCredential sets field value
func (o *RunTransactionRequest) SetWeb3SigningCredential(v Web3SigningCredential) {
	o.Web3SigningCredential = v
}

// GetTransactionConfig returns the TransactionConfig field value
func (o *RunTransactionRequest) GetTransactionConfig() PolkadotTransactionConfig {
	if o == nil {
		var ret PolkadotTransactionConfig
		return ret
	}

	return o.TransactionConfig
}

// GetTransactionConfigOk returns a tuple with the TransactionConfig field value
// and a boolean to check if the value has been set.
func (o *RunTransactionRequest) GetTransactionConfigOk() (*PolkadotTransactionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionConfig, true
}

// SetTransactionConfig sets field value
func (o *RunTransactionRequest) SetTransactionConfig(v PolkadotTransactionConfig) {
	o.TransactionConfig = v
}

func (o RunTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunTransactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["web3SigningCredential"] = o.Web3SigningCredential
	toSerialize["transactionConfig"] = o.TransactionConfig
	return toSerialize, nil
}

type NullableRunTransactionRequest struct {
	value *RunTransactionRequest
	isSet bool
}

func (v NullableRunTransactionRequest) Get() *RunTransactionRequest {
	return v.value
}

func (v *NullableRunTransactionRequest) Set(val *RunTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRunTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRunTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunTransactionRequest(val *RunTransactionRequest) *NullableRunTransactionRequest {
	return &NullableRunTransactionRequest{value: val, isSet: true}
}

func (v NullableRunTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


