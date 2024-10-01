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

// checks if the GetObjectRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetObjectRequestV1{}

// GetObjectRequestV1 struct for GetObjectRequestV1
type GetObjectRequestV1 struct {
	// The key for the entry to get from the object store.
	Key string `json:"key"`
}

// NewGetObjectRequestV1 instantiates a new GetObjectRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetObjectRequestV1(key string) *GetObjectRequestV1 {
	this := GetObjectRequestV1{}
	this.Key = key
	return &this
}

// NewGetObjectRequestV1WithDefaults instantiates a new GetObjectRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetObjectRequestV1WithDefaults() *GetObjectRequestV1 {
	this := GetObjectRequestV1{}
	return &this
}

// GetKey returns the Key field value
func (o *GetObjectRequestV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GetObjectRequestV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GetObjectRequestV1) SetKey(v string) {
	o.Key = v
}

func (o GetObjectRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetObjectRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableGetObjectRequestV1 struct {
	value *GetObjectRequestV1
	isSet bool
}

func (v NullableGetObjectRequestV1) Get() *GetObjectRequestV1 {
	return v.value
}

func (v *NullableGetObjectRequestV1) Set(val *GetObjectRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetObjectRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetObjectRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetObjectRequestV1(val *GetObjectRequestV1) *NullableGetObjectRequestV1 {
	return &NullableGetObjectRequestV1{value: val, isSet: true}
}

func (v NullableGetObjectRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetObjectRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


