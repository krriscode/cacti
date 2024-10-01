/*
Hyperledger Cactus - Keychain API

Contains/describes the Keychain API types/paths for Hyperledger Cactus.

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-keychain-vault

import (
	"encoding/json"
)

// checks if the GetKeychainEntryResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetKeychainEntryResponseV1{}

// GetKeychainEntryResponseV1 struct for GetKeychainEntryResponseV1
type GetKeychainEntryResponseV1 struct {
	// The key that was used to retrieve the value from the keychain.
	Key string `json:"key"`
	// The value associated with the requested key on the keychain.
	Value string `json:"value"`
}

// NewGetKeychainEntryResponseV1 instantiates a new GetKeychainEntryResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKeychainEntryResponseV1(key string, value string) *GetKeychainEntryResponseV1 {
	this := GetKeychainEntryResponseV1{}
	this.Key = key
	this.Value = value
	return &this
}

// NewGetKeychainEntryResponseV1WithDefaults instantiates a new GetKeychainEntryResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKeychainEntryResponseV1WithDefaults() *GetKeychainEntryResponseV1 {
	this := GetKeychainEntryResponseV1{}
	return &this
}

// GetKey returns the Key field value
func (o *GetKeychainEntryResponseV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GetKeychainEntryResponseV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GetKeychainEntryResponseV1) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *GetKeychainEntryResponseV1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetKeychainEntryResponseV1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetKeychainEntryResponseV1) SetValue(v string) {
	o.Value = v
}

func (o GetKeychainEntryResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetKeychainEntryResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableGetKeychainEntryResponseV1 struct {
	value *GetKeychainEntryResponseV1
	isSet bool
}

func (v NullableGetKeychainEntryResponseV1) Get() *GetKeychainEntryResponseV1 {
	return v.value
}

func (v *NullableGetKeychainEntryResponseV1) Set(val *GetKeychainEntryResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetKeychainEntryResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetKeychainEntryResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetKeychainEntryResponseV1(val *GetKeychainEntryResponseV1) *NullableGetKeychainEntryResponseV1 {
	return &NullableGetKeychainEntryResponseV1{value: val, isSet: true}
}

func (v NullableGetKeychainEntryResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetKeychainEntryResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


