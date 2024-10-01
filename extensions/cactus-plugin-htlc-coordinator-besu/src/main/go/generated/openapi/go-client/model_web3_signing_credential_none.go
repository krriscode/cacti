/*
Hyperledger Cactus Plugin - HTLC Coordinator

Can exchange assets between networks

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-htlc-coordinator-besu

import (
	"encoding/json"
)

// checks if the Web3SigningCredentialNone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3SigningCredentialNone{}

// Web3SigningCredentialNone Using this denotes that there is no signing required because the transaction is pre-signed.
type Web3SigningCredentialNone struct {
	Type Web3SigningCredentialType `json:"type"`
}

// NewWeb3SigningCredentialNone instantiates a new Web3SigningCredentialNone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3SigningCredentialNone(type_ Web3SigningCredentialType) *Web3SigningCredentialNone {
	this := Web3SigningCredentialNone{}
	this.Type = type_
	return &this
}

// NewWeb3SigningCredentialNoneWithDefaults instantiates a new Web3SigningCredentialNone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3SigningCredentialNoneWithDefaults() *Web3SigningCredentialNone {
	this := Web3SigningCredentialNone{}
	return &this
}

// GetType returns the Type field value
func (o *Web3SigningCredentialNone) GetType() Web3SigningCredentialType {
	if o == nil {
		var ret Web3SigningCredentialType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Web3SigningCredentialNone) GetTypeOk() (*Web3SigningCredentialType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Web3SigningCredentialNone) SetType(v Web3SigningCredentialType) {
	o.Type = v
}

func (o Web3SigningCredentialNone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3SigningCredentialNone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableWeb3SigningCredentialNone struct {
	value *Web3SigningCredentialNone
	isSet bool
}

func (v NullableWeb3SigningCredentialNone) Get() *Web3SigningCredentialNone {
	return v.value
}

func (v *NullableWeb3SigningCredentialNone) Set(val *Web3SigningCredentialNone) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3SigningCredentialNone) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3SigningCredentialNone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3SigningCredentialNone(val *Web3SigningCredentialNone) *NullableWeb3SigningCredentialNone {
	return &NullableWeb3SigningCredentialNone{value: val, isSet: true}
}

func (v NullableWeb3SigningCredentialNone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3SigningCredentialNone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


