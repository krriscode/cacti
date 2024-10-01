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

// checks if the WithdrawCounterpartyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WithdrawCounterpartyRequest{}

// WithdrawCounterpartyRequest struct for WithdrawCounterpartyRequest
type WithdrawCounterpartyRequest struct {
	HtlcPackage HtlcPackage `json:"htlcPackage"`
	// connector Instance Id for the connector plugin
	ConnectorInstanceId string `json:"connectorInstanceId"`
	// keychainId for the keychain plugin
	KeychainId string `json:"keychainId"`
	// contractId for the contract
	ContractId *string `json:"contractId,omitempty"`
	Web3SigningCredential Web3SigningCredential `json:"web3SigningCredential"`
	// Id for the HTLC
	HtlcId string `json:"htlcId"`
	// Counterparty HTLC secret
	Secret string `json:"secret"`
	Gas *float32 `json:"gas,omitempty"`
}

// NewWithdrawCounterpartyRequest instantiates a new WithdrawCounterpartyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawCounterpartyRequest(htlcPackage HtlcPackage, connectorInstanceId string, keychainId string, web3SigningCredential Web3SigningCredential, htlcId string, secret string) *WithdrawCounterpartyRequest {
	this := WithdrawCounterpartyRequest{}
	this.HtlcPackage = htlcPackage
	this.ConnectorInstanceId = connectorInstanceId
	this.KeychainId = keychainId
	this.Web3SigningCredential = web3SigningCredential
	this.HtlcId = htlcId
	this.Secret = secret
	return &this
}

// NewWithdrawCounterpartyRequestWithDefaults instantiates a new WithdrawCounterpartyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawCounterpartyRequestWithDefaults() *WithdrawCounterpartyRequest {
	this := WithdrawCounterpartyRequest{}
	return &this
}

// GetHtlcPackage returns the HtlcPackage field value
func (o *WithdrawCounterpartyRequest) GetHtlcPackage() HtlcPackage {
	if o == nil {
		var ret HtlcPackage
		return ret
	}

	return o.HtlcPackage
}

// GetHtlcPackageOk returns a tuple with the HtlcPackage field value
// and a boolean to check if the value has been set.
func (o *WithdrawCounterpartyRequest) GetHtlcPackageOk() (*HtlcPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtlcPackage, true
}

// SetHtlcPackage sets field value
func (o *WithdrawCounterpartyRequest) SetHtlcPackage(v HtlcPackage) {
	o.HtlcPackage = v
}

// GetConnectorInstanceId returns the ConnectorInstanceId field value
func (o *WithdrawCounterpartyRequest) GetConnectorInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorInstanceId
}

// GetConnectorInstanceIdOk returns a tuple with the ConnectorInstanceId field value
// and a boolean to check if the value has been set.
func (o *WithdrawCounterpartyRequest) GetConnectorInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorInstanceId, true
}

// SetConnectorInstanceId sets field value
func (o *WithdrawCounterpartyRequest) SetConnectorInstanceId(v string) {
	o.ConnectorInstanceId = v
}

// GetKeychainId returns the KeychainId field value
func (o *WithdrawCounterpartyRequest) GetKeychainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeychainId
}

// GetKeychainIdOk returns a tuple with the KeychainId field value
// and a boolean to check if the value has been set.
func (o *WithdrawCounterpartyRequest) GetKeychainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeychainId, true
}

// SetKeychainId sets field value
func (o *WithdrawCounterpartyRequest) SetKeychainId(v string) {
	o.KeychainId = v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *WithdrawCounterpartyRequest) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawCounterpartyRequest) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *WithdrawCounterpartyRequest) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *WithdrawCounterpartyRequest) SetContractId(v string) {
	o.ContractId = &v
}

// GetWeb3SigningCredential returns the Web3SigningCredential field value
func (o *WithdrawCounterpartyRequest) GetWeb3SigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.Web3SigningCredential
}

// GetWeb3SigningCredentialOk returns a tuple with the Web3SigningCredential field value
// and a boolean to check if the value has been set.
func (o *WithdrawCounterpartyRequest) GetWeb3SigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Web3SigningCredential, true
}

// SetWeb3SigningCredential sets field value
func (o *WithdrawCounterpartyRequest) SetWeb3SigningCredential(v Web3SigningCredential) {
	o.Web3SigningCredential = v
}

// GetHtlcId returns the HtlcId field value
func (o *WithdrawCounterpartyRequest) GetHtlcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtlcId
}

// GetHtlcIdOk returns a tuple with the HtlcId field value
// and a boolean to check if the value has been set.
func (o *WithdrawCounterpartyRequest) GetHtlcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtlcId, true
}

// SetHtlcId sets field value
func (o *WithdrawCounterpartyRequest) SetHtlcId(v string) {
	o.HtlcId = v
}

// GetSecret returns the Secret field value
func (o *WithdrawCounterpartyRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *WithdrawCounterpartyRequest) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *WithdrawCounterpartyRequest) SetSecret(v string) {
	o.Secret = v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *WithdrawCounterpartyRequest) GetGas() float32 {
	if o == nil || IsNil(o.Gas) {
		var ret float32
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawCounterpartyRequest) GetGasOk() (*float32, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *WithdrawCounterpartyRequest) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given float32 and assigns it to the Gas field.
func (o *WithdrawCounterpartyRequest) SetGas(v float32) {
	o.Gas = &v
}

func (o WithdrawCounterpartyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithdrawCounterpartyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["htlcPackage"] = o.HtlcPackage
	toSerialize["connectorInstanceId"] = o.ConnectorInstanceId
	toSerialize["keychainId"] = o.KeychainId
	if !IsNil(o.ContractId) {
		toSerialize["contractId"] = o.ContractId
	}
	toSerialize["web3SigningCredential"] = o.Web3SigningCredential
	toSerialize["htlcId"] = o.HtlcId
	toSerialize["secret"] = o.Secret
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	return toSerialize, nil
}

type NullableWithdrawCounterpartyRequest struct {
	value *WithdrawCounterpartyRequest
	isSet bool
}

func (v NullableWithdrawCounterpartyRequest) Get() *WithdrawCounterpartyRequest {
	return v.value
}

func (v *NullableWithdrawCounterpartyRequest) Set(val *WithdrawCounterpartyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawCounterpartyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawCounterpartyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawCounterpartyRequest(val *WithdrawCounterpartyRequest) *NullableWithdrawCounterpartyRequest {
	return &NullableWithdrawCounterpartyRequest{value: val, isSet: true}
}

func (v NullableWithdrawCounterpartyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawCounterpartyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


