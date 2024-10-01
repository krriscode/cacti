/*
Hyperledger Cactus Plugin - Object Store - IPFS 

Contains/describes the Hyperledger Cactus Object Store IPFS plugin.

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-object-store-ipfs

import (
	"encoding/json"
)

// checks if the GetObjectResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetObjectResponseV1{}

// GetObjectResponseV1 struct for GetObjectResponseV1
type GetObjectResponseV1 struct {
	// The key that was used to retrieve the value from the object store.
	Key string `json:"key"`
	// The value associated with the requested key in the object store as a string.
	Value string `json:"value"`
}

// NewGetObjectResponseV1 instantiates a new GetObjectResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetObjectResponseV1(key string, value string) *GetObjectResponseV1 {
	this := GetObjectResponseV1{}
	this.Key = key
	this.Value = value
	return &this
}

// NewGetObjectResponseV1WithDefaults instantiates a new GetObjectResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetObjectResponseV1WithDefaults() *GetObjectResponseV1 {
	this := GetObjectResponseV1{}
	return &this
}

// GetKey returns the Key field value
func (o *GetObjectResponseV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GetObjectResponseV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GetObjectResponseV1) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *GetObjectResponseV1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetObjectResponseV1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetObjectResponseV1) SetValue(v string) {
	o.Value = v
}

func (o GetObjectResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetObjectResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableGetObjectResponseV1 struct {
	value *GetObjectResponseV1
	isSet bool
}

func (v NullableGetObjectResponseV1) Get() *GetObjectResponseV1 {
	return v.value
}

func (v *NullableGetObjectResponseV1) Set(val *GetObjectResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetObjectResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetObjectResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetObjectResponseV1(val *GetObjectResponseV1) *NullableGetObjectResponseV1 {
	return &NullableGetObjectResponseV1{value: val, isSet: true}
}

func (v NullableGetObjectResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetObjectResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


