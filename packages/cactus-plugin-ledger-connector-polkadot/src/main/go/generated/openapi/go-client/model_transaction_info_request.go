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

// checks if the TransactionInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInfoRequest{}

// TransactionInfoRequest struct for TransactionInfoRequest
type TransactionInfoRequest struct {
	AccountAddress string `json:"accountAddress"`
	TransactionExpiration NullableFloat32 `json:"transactionExpiration"`
}

// NewTransactionInfoRequest instantiates a new TransactionInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInfoRequest(accountAddress string, transactionExpiration NullableFloat32) *TransactionInfoRequest {
	this := TransactionInfoRequest{}
	this.AccountAddress = accountAddress
	this.TransactionExpiration = transactionExpiration
	return &this
}

// NewTransactionInfoRequestWithDefaults instantiates a new TransactionInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInfoRequestWithDefaults() *TransactionInfoRequest {
	this := TransactionInfoRequest{}
	return &this
}

// GetAccountAddress returns the AccountAddress field value
func (o *TransactionInfoRequest) GetAccountAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountAddress
}

// GetAccountAddressOk returns a tuple with the AccountAddress field value
// and a boolean to check if the value has been set.
func (o *TransactionInfoRequest) GetAccountAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountAddress, true
}

// SetAccountAddress sets field value
func (o *TransactionInfoRequest) SetAccountAddress(v string) {
	o.AccountAddress = v
}

// GetTransactionExpiration returns the TransactionExpiration field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *TransactionInfoRequest) GetTransactionExpiration() float32 {
	if o == nil || o.TransactionExpiration.Get() == nil {
		var ret float32
		return ret
	}

	return *o.TransactionExpiration.Get()
}

// GetTransactionExpirationOk returns a tuple with the TransactionExpiration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionInfoRequest) GetTransactionExpirationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionExpiration.Get(), o.TransactionExpiration.IsSet()
}

// SetTransactionExpiration sets field value
func (o *TransactionInfoRequest) SetTransactionExpiration(v float32) {
	o.TransactionExpiration.Set(&v)
}

func (o TransactionInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountAddress"] = o.AccountAddress
	toSerialize["transactionExpiration"] = o.TransactionExpiration.Get()
	return toSerialize, nil
}

type NullableTransactionInfoRequest struct {
	value *TransactionInfoRequest
	isSet bool
}

func (v NullableTransactionInfoRequest) Get() *TransactionInfoRequest {
	return v.value
}

func (v *NullableTransactionInfoRequest) Set(val *TransactionInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInfoRequest(val *TransactionInfoRequest) *NullableTransactionInfoRequest {
	return &NullableTransactionInfoRequest{value: val, isSet: true}
}

func (v NullableTransactionInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


