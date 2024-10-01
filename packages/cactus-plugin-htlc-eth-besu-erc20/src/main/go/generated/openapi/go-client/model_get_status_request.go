/*
Hyperledger Cactus Plugin - HTLC ETH BESU ERC20

Allows Cactus nodes to interact with HTLC contracts with ERC-20 Tokens

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-htlc-eth-besu-erc20

import (
	"encoding/json"
)

// checks if the GetStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatusRequest{}

// GetStatusRequest Defines the parameters for retrieving the status of the HTLC swap.
type GetStatusRequest struct {
	Ids []string `json:"ids"`
	Web3SigningCredential Web3SigningCredential `json:"web3SigningCredential"`
	ConnectorId string `json:"connectorId"`
	KeychainId string `json:"keychainId"`
}

// NewGetStatusRequest instantiates a new GetStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatusRequest(ids []string, web3SigningCredential Web3SigningCredential, connectorId string, keychainId string) *GetStatusRequest {
	this := GetStatusRequest{}
	this.Ids = ids
	this.Web3SigningCredential = web3SigningCredential
	this.ConnectorId = connectorId
	this.KeychainId = keychainId
	return &this
}

// NewGetStatusRequestWithDefaults instantiates a new GetStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatusRequestWithDefaults() *GetStatusRequest {
	this := GetStatusRequest{}
	return &this
}

// GetIds returns the Ids field value
func (o *GetStatusRequest) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *GetStatusRequest) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *GetStatusRequest) SetIds(v []string) {
	o.Ids = v
}

// GetWeb3SigningCredential returns the Web3SigningCredential field value
func (o *GetStatusRequest) GetWeb3SigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.Web3SigningCredential
}

// GetWeb3SigningCredentialOk returns a tuple with the Web3SigningCredential field value
// and a boolean to check if the value has been set.
func (o *GetStatusRequest) GetWeb3SigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Web3SigningCredential, true
}

// SetWeb3SigningCredential sets field value
func (o *GetStatusRequest) SetWeb3SigningCredential(v Web3SigningCredential) {
	o.Web3SigningCredential = v
}

// GetConnectorId returns the ConnectorId field value
func (o *GetStatusRequest) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *GetStatusRequest) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *GetStatusRequest) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetKeychainId returns the KeychainId field value
func (o *GetStatusRequest) GetKeychainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeychainId
}

// GetKeychainIdOk returns a tuple with the KeychainId field value
// and a boolean to check if the value has been set.
func (o *GetStatusRequest) GetKeychainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeychainId, true
}

// SetKeychainId sets field value
func (o *GetStatusRequest) SetKeychainId(v string) {
	o.KeychainId = v
}

func (o GetStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	toSerialize["web3SigningCredential"] = o.Web3SigningCredential
	toSerialize["connectorId"] = o.ConnectorId
	toSerialize["keychainId"] = o.KeychainId
	return toSerialize, nil
}

type NullableGetStatusRequest struct {
	value *GetStatusRequest
	isSet bool
}

func (v NullableGetStatusRequest) Get() *GetStatusRequest {
	return v.value
}

func (v *NullableGetStatusRequest) Set(val *GetStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatusRequest(val *GetStatusRequest) *NullableGetStatusRequest {
	return &NullableGetStatusRequest{value: val, isSet: true}
}

func (v NullableGetStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


