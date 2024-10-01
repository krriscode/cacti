/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the WatchBlocksV1Progress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksV1Progress{}

// WatchBlocksV1Progress struct for WatchBlocksV1Progress
type WatchBlocksV1Progress struct {
	BlockHeader Web3BlockHeader `json:"blockHeader"`
}

// NewWatchBlocksV1Progress instantiates a new WatchBlocksV1Progress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksV1Progress(blockHeader Web3BlockHeader) *WatchBlocksV1Progress {
	this := WatchBlocksV1Progress{}
	this.BlockHeader = blockHeader
	return &this
}

// NewWatchBlocksV1ProgressWithDefaults instantiates a new WatchBlocksV1Progress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksV1ProgressWithDefaults() *WatchBlocksV1Progress {
	this := WatchBlocksV1Progress{}
	return &this
}

// GetBlockHeader returns the BlockHeader field value
func (o *WatchBlocksV1Progress) GetBlockHeader() Web3BlockHeader {
	if o == nil {
		var ret Web3BlockHeader
		return ret
	}

	return o.BlockHeader
}

// GetBlockHeaderOk returns a tuple with the BlockHeader field value
// and a boolean to check if the value has been set.
func (o *WatchBlocksV1Progress) GetBlockHeaderOk() (*Web3BlockHeader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeader, true
}

// SetBlockHeader sets field value
func (o *WatchBlocksV1Progress) SetBlockHeader(v Web3BlockHeader) {
	o.BlockHeader = v
}

func (o WatchBlocksV1Progress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksV1Progress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blockHeader"] = o.BlockHeader
	return toSerialize, nil
}

type NullableWatchBlocksV1Progress struct {
	value *WatchBlocksV1Progress
	isSet bool
}

func (v NullableWatchBlocksV1Progress) Get() *WatchBlocksV1Progress {
	return v.value
}

func (v *NullableWatchBlocksV1Progress) Set(val *WatchBlocksV1Progress) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksV1Progress) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksV1Progress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksV1Progress(val *WatchBlocksV1Progress) *NullableWatchBlocksV1Progress {
	return &NullableWatchBlocksV1Progress{value: val, isSet: true}
}

func (v NullableWatchBlocksV1Progress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksV1Progress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


