/*
Hyperledger Cactus Plugin - HTLC-ETH Besu

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-htlc-eth-besu

import (
	"encoding/json"
)

// checks if the GetSingleStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleStatusRequest{}

// GetSingleStatusRequest Defines the parameters for retrieving the single status of the HTLC swap.
type GetSingleStatusRequest struct {
	Id string `json:"id"`
	Web3SigningCredential Web3SigningCredential `json:"web3SigningCredential"`
	ConnectorId string `json:"connectorId"`
	KeychainId string `json:"keychainId"`
}

// NewGetSingleStatusRequest instantiates a new GetSingleStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleStatusRequest(id string, web3SigningCredential Web3SigningCredential, connectorId string, keychainId string) *GetSingleStatusRequest {
	this := GetSingleStatusRequest{}
	this.Id = id
	this.Web3SigningCredential = web3SigningCredential
	this.ConnectorId = connectorId
	this.KeychainId = keychainId
	return &this
}

// NewGetSingleStatusRequestWithDefaults instantiates a new GetSingleStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleStatusRequestWithDefaults() *GetSingleStatusRequest {
	this := GetSingleStatusRequest{}
	return &this
}

// GetId returns the Id field value
func (o *GetSingleStatusRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetSingleStatusRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetSingleStatusRequest) SetId(v string) {
	o.Id = v
}

// GetWeb3SigningCredential returns the Web3SigningCredential field value
func (o *GetSingleStatusRequest) GetWeb3SigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.Web3SigningCredential
}

// GetWeb3SigningCredentialOk returns a tuple with the Web3SigningCredential field value
// and a boolean to check if the value has been set.
func (o *GetSingleStatusRequest) GetWeb3SigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Web3SigningCredential, true
}

// SetWeb3SigningCredential sets field value
func (o *GetSingleStatusRequest) SetWeb3SigningCredential(v Web3SigningCredential) {
	o.Web3SigningCredential = v
}

// GetConnectorId returns the ConnectorId field value
func (o *GetSingleStatusRequest) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *GetSingleStatusRequest) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *GetSingleStatusRequest) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetKeychainId returns the KeychainId field value
func (o *GetSingleStatusRequest) GetKeychainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeychainId
}

// GetKeychainIdOk returns a tuple with the KeychainId field value
// and a boolean to check if the value has been set.
func (o *GetSingleStatusRequest) GetKeychainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeychainId, true
}

// SetKeychainId sets field value
func (o *GetSingleStatusRequest) SetKeychainId(v string) {
	o.KeychainId = v
}

func (o GetSingleStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["web3SigningCredential"] = o.Web3SigningCredential
	toSerialize["connectorId"] = o.ConnectorId
	toSerialize["keychainId"] = o.KeychainId
	return toSerialize, nil
}

type NullableGetSingleStatusRequest struct {
	value *GetSingleStatusRequest
	isSet bool
}

func (v NullableGetSingleStatusRequest) Get() *GetSingleStatusRequest {
	return v.value
}

func (v *NullableGetSingleStatusRequest) Set(val *GetSingleStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleStatusRequest(val *GetSingleStatusRequest) *NullableGetSingleStatusRequest {
	return &NullableGetSingleStatusRequest{value: val, isSet: true}
}

func (v NullableGetSingleStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


