/*
Hyperledger Cactus - Keychain API

Contains/describes the Keychain API types/paths for Hyperledger Cactus.

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-keychain-google-sm

import (
	"encoding/json"
)

// checks if the HasKeychainEntryRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasKeychainEntryRequestV1{}

// HasKeychainEntryRequestV1 struct for HasKeychainEntryRequestV1
type HasKeychainEntryRequestV1 struct {
	// The key to check for presence in the keychain.
	Key string `json:"key"`
}

// NewHasKeychainEntryRequestV1 instantiates a new HasKeychainEntryRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasKeychainEntryRequestV1(key string) *HasKeychainEntryRequestV1 {
	this := HasKeychainEntryRequestV1{}
	this.Key = key
	return &this
}

// NewHasKeychainEntryRequestV1WithDefaults instantiates a new HasKeychainEntryRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasKeychainEntryRequestV1WithDefaults() *HasKeychainEntryRequestV1 {
	this := HasKeychainEntryRequestV1{}
	return &this
}

// GetKey returns the Key field value
func (o *HasKeychainEntryRequestV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *HasKeychainEntryRequestV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *HasKeychainEntryRequestV1) SetKey(v string) {
	o.Key = v
}

func (o HasKeychainEntryRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasKeychainEntryRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableHasKeychainEntryRequestV1 struct {
	value *HasKeychainEntryRequestV1
	isSet bool
}

func (v NullableHasKeychainEntryRequestV1) Get() *HasKeychainEntryRequestV1 {
	return v.value
}

func (v *NullableHasKeychainEntryRequestV1) Set(val *HasKeychainEntryRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableHasKeychainEntryRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableHasKeychainEntryRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasKeychainEntryRequestV1(val *HasKeychainEntryRequestV1) *NullableHasKeychainEntryRequestV1 {
	return &NullableHasKeychainEntryRequestV1{value: val, isSet: true}
}

func (v NullableHasKeychainEntryRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasKeychainEntryRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


