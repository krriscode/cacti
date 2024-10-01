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

// checks if the RawTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RawTransactionResponse{}

// RawTransactionResponse struct for RawTransactionResponse
type RawTransactionResponse struct {
	ResponseContainer RawTransactionResponseResponseContainer `json:"responseContainer"`
}

// NewRawTransactionResponse instantiates a new RawTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawTransactionResponse(responseContainer RawTransactionResponseResponseContainer) *RawTransactionResponse {
	this := RawTransactionResponse{}
	this.ResponseContainer = responseContainer
	return &this
}

// NewRawTransactionResponseWithDefaults instantiates a new RawTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawTransactionResponseWithDefaults() *RawTransactionResponse {
	this := RawTransactionResponse{}
	return &this
}

// GetResponseContainer returns the ResponseContainer field value
func (o *RawTransactionResponse) GetResponseContainer() RawTransactionResponseResponseContainer {
	if o == nil {
		var ret RawTransactionResponseResponseContainer
		return ret
	}

	return o.ResponseContainer
}

// GetResponseContainerOk returns a tuple with the ResponseContainer field value
// and a boolean to check if the value has been set.
func (o *RawTransactionResponse) GetResponseContainerOk() (*RawTransactionResponseResponseContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseContainer, true
}

// SetResponseContainer sets field value
func (o *RawTransactionResponse) SetResponseContainer(v RawTransactionResponseResponseContainer) {
	o.ResponseContainer = v
}

func (o RawTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RawTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["responseContainer"] = o.ResponseContainer
	return toSerialize, nil
}

type NullableRawTransactionResponse struct {
	value *RawTransactionResponse
	isSet bool
}

func (v NullableRawTransactionResponse) Get() *RawTransactionResponse {
	return v.value
}

func (v *NullableRawTransactionResponse) Set(val *RawTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRawTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRawTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawTransactionResponse(val *RawTransactionResponse) *NullableRawTransactionResponse {
	return &NullableRawTransactionResponse{value: val, isSet: true}
}

func (v NullableRawTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


