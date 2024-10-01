/*
Hyperledger Cacti Plugin - Connector CDL

Can perform basic tasks on Fujitsu CDL service.

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-cdl

import (
	"encoding/json"
)

// checks if the AuthInfoAccessTokenV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthInfoAccessTokenV1{}

// AuthInfoAccessTokenV1 struct for AuthInfoAccessTokenV1
type AuthInfoAccessTokenV1 struct {
	AccessToken string `json:"accessToken"`
	TrustAgentId string `json:"trustAgentId"`
}

// NewAuthInfoAccessTokenV1 instantiates a new AuthInfoAccessTokenV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthInfoAccessTokenV1(accessToken string, trustAgentId string) *AuthInfoAccessTokenV1 {
	this := AuthInfoAccessTokenV1{}
	this.AccessToken = accessToken
	this.TrustAgentId = trustAgentId
	return &this
}

// NewAuthInfoAccessTokenV1WithDefaults instantiates a new AuthInfoAccessTokenV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthInfoAccessTokenV1WithDefaults() *AuthInfoAccessTokenV1 {
	this := AuthInfoAccessTokenV1{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *AuthInfoAccessTokenV1) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AuthInfoAccessTokenV1) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AuthInfoAccessTokenV1) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetTrustAgentId returns the TrustAgentId field value
func (o *AuthInfoAccessTokenV1) GetTrustAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustAgentId
}

// GetTrustAgentIdOk returns a tuple with the TrustAgentId field value
// and a boolean to check if the value has been set.
func (o *AuthInfoAccessTokenV1) GetTrustAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustAgentId, true
}

// SetTrustAgentId sets field value
func (o *AuthInfoAccessTokenV1) SetTrustAgentId(v string) {
	o.TrustAgentId = v
}

func (o AuthInfoAccessTokenV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthInfoAccessTokenV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessToken"] = o.AccessToken
	toSerialize["trustAgentId"] = o.TrustAgentId
	return toSerialize, nil
}

type NullableAuthInfoAccessTokenV1 struct {
	value *AuthInfoAccessTokenV1
	isSet bool
}

func (v NullableAuthInfoAccessTokenV1) Get() *AuthInfoAccessTokenV1 {
	return v.value
}

func (v *NullableAuthInfoAccessTokenV1) Set(val *AuthInfoAccessTokenV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthInfoAccessTokenV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthInfoAccessTokenV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthInfoAccessTokenV1(val *AuthInfoAccessTokenV1) *NullableAuthInfoAccessTokenV1 {
	return &NullableAuthInfoAccessTokenV1{value: val, isSet: true}
}

func (v NullableAuthInfoAccessTokenV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthInfoAccessTokenV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


