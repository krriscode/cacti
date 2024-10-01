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

// checks if the SignRawTransactionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignRawTransactionRequest{}

// SignRawTransactionRequest struct for SignRawTransactionRequest
type SignRawTransactionRequest struct {
	RawTransaction string `json:"rawTransaction"`
	Mnemonic string `json:"mnemonic"`
	SigningOptions map[string]interface{} `json:"signingOptions,omitempty"`
}

// NewSignRawTransactionRequest instantiates a new SignRawTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignRawTransactionRequest(rawTransaction string, mnemonic string) *SignRawTransactionRequest {
	this := SignRawTransactionRequest{}
	this.RawTransaction = rawTransaction
	this.Mnemonic = mnemonic
	return &this
}

// NewSignRawTransactionRequestWithDefaults instantiates a new SignRawTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignRawTransactionRequestWithDefaults() *SignRawTransactionRequest {
	this := SignRawTransactionRequest{}
	return &this
}

// GetRawTransaction returns the RawTransaction field value
func (o *SignRawTransactionRequest) GetRawTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawTransaction
}

// GetRawTransactionOk returns a tuple with the RawTransaction field value
// and a boolean to check if the value has been set.
func (o *SignRawTransactionRequest) GetRawTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawTransaction, true
}

// SetRawTransaction sets field value
func (o *SignRawTransactionRequest) SetRawTransaction(v string) {
	o.RawTransaction = v
}

// GetMnemonic returns the Mnemonic field value
func (o *SignRawTransactionRequest) GetMnemonic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnemonic
}

// GetMnemonicOk returns a tuple with the Mnemonic field value
// and a boolean to check if the value has been set.
func (o *SignRawTransactionRequest) GetMnemonicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnemonic, true
}

// SetMnemonic sets field value
func (o *SignRawTransactionRequest) SetMnemonic(v string) {
	o.Mnemonic = v
}

// GetSigningOptions returns the SigningOptions field value if set, zero value otherwise.
func (o *SignRawTransactionRequest) GetSigningOptions() map[string]interface{} {
	if o == nil || IsNil(o.SigningOptions) {
		var ret map[string]interface{}
		return ret
	}
	return o.SigningOptions
}

// GetSigningOptionsOk returns a tuple with the SigningOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignRawTransactionRequest) GetSigningOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SigningOptions) {
		return map[string]interface{}{}, false
	}
	return o.SigningOptions, true
}

// HasSigningOptions returns a boolean if a field has been set.
func (o *SignRawTransactionRequest) HasSigningOptions() bool {
	if o != nil && !IsNil(o.SigningOptions) {
		return true
	}

	return false
}

// SetSigningOptions gets a reference to the given map[string]interface{} and assigns it to the SigningOptions field.
func (o *SignRawTransactionRequest) SetSigningOptions(v map[string]interface{}) {
	o.SigningOptions = v
}

func (o SignRawTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignRawTransactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rawTransaction"] = o.RawTransaction
	toSerialize["mnemonic"] = o.Mnemonic
	if !IsNil(o.SigningOptions) {
		toSerialize["signingOptions"] = o.SigningOptions
	}
	return toSerialize, nil
}

type NullableSignRawTransactionRequest struct {
	value *SignRawTransactionRequest
	isSet bool
}

func (v NullableSignRawTransactionRequest) Get() *SignRawTransactionRequest {
	return v.value
}

func (v *NullableSignRawTransactionRequest) Set(val *SignRawTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignRawTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignRawTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignRawTransactionRequest(val *SignRawTransactionRequest) *NullableSignRawTransactionRequest {
	return &NullableSignRawTransactionRequest{value: val, isSet: true}
}

func (v NullableSignRawTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignRawTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


