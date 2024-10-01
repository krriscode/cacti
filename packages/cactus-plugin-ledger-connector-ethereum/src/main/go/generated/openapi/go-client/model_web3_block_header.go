/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
)

// checks if the Web3BlockHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3BlockHeader{}

// Web3BlockHeader struct for Web3BlockHeader
type Web3BlockHeader struct {
	Number *string `json:"number,omitempty"`
	ParentHash *string `json:"parentHash,omitempty"`
	Sha3Uncles string `json:"sha3Uncles"`
	TransactionsRoot *string `json:"transactionsRoot,omitempty"`
	ReceiptsRoot *string `json:"receiptsRoot,omitempty"`
	Difficulty *string `json:"difficulty,omitempty"`
	MixHash *string `json:"mixHash,omitempty"`
	Miner *string `json:"miner,omitempty"`
	GasLimit string `json:"gasLimit"`
	GasUsed string `json:"gasUsed"`
	StateRoot *string `json:"stateRoot,omitempty"`
	LogsBloom *string `json:"logsBloom,omitempty"`
	ExtraData *string `json:"extraData,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewWeb3BlockHeader instantiates a new Web3BlockHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3BlockHeader(sha3Uncles string, gasLimit string, gasUsed string) *Web3BlockHeader {
	this := Web3BlockHeader{}
	this.Sha3Uncles = sha3Uncles
	this.GasLimit = gasLimit
	this.GasUsed = gasUsed
	return &this
}

// NewWeb3BlockHeaderWithDefaults instantiates a new Web3BlockHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3BlockHeaderWithDefaults() *Web3BlockHeader {
	this := Web3BlockHeader{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Web3BlockHeader) SetNumber(v string) {
	o.Number = &v
}

// GetParentHash returns the ParentHash field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetParentHash() string {
	if o == nil || IsNil(o.ParentHash) {
		var ret string
		return ret
	}
	return *o.ParentHash
}

// GetParentHashOk returns a tuple with the ParentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetParentHashOk() (*string, bool) {
	if o == nil || IsNil(o.ParentHash) {
		return nil, false
	}
	return o.ParentHash, true
}

// HasParentHash returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasParentHash() bool {
	if o != nil && !IsNil(o.ParentHash) {
		return true
	}

	return false
}

// SetParentHash gets a reference to the given string and assigns it to the ParentHash field.
func (o *Web3BlockHeader) SetParentHash(v string) {
	o.ParentHash = &v
}

// GetSha3Uncles returns the Sha3Uncles field value
func (o *Web3BlockHeader) GetSha3Uncles() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha3Uncles
}

// GetSha3UnclesOk returns a tuple with the Sha3Uncles field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetSha3UnclesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha3Uncles, true
}

// SetSha3Uncles sets field value
func (o *Web3BlockHeader) SetSha3Uncles(v string) {
	o.Sha3Uncles = v
}

// GetTransactionsRoot returns the TransactionsRoot field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetTransactionsRoot() string {
	if o == nil || IsNil(o.TransactionsRoot) {
		var ret string
		return ret
	}
	return *o.TransactionsRoot
}

// GetTransactionsRootOk returns a tuple with the TransactionsRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetTransactionsRootOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionsRoot) {
		return nil, false
	}
	return o.TransactionsRoot, true
}

// HasTransactionsRoot returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasTransactionsRoot() bool {
	if o != nil && !IsNil(o.TransactionsRoot) {
		return true
	}

	return false
}

// SetTransactionsRoot gets a reference to the given string and assigns it to the TransactionsRoot field.
func (o *Web3BlockHeader) SetTransactionsRoot(v string) {
	o.TransactionsRoot = &v
}

// GetReceiptsRoot returns the ReceiptsRoot field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetReceiptsRoot() string {
	if o == nil || IsNil(o.ReceiptsRoot) {
		var ret string
		return ret
	}
	return *o.ReceiptsRoot
}

// GetReceiptsRootOk returns a tuple with the ReceiptsRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetReceiptsRootOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptsRoot) {
		return nil, false
	}
	return o.ReceiptsRoot, true
}

// HasReceiptsRoot returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasReceiptsRoot() bool {
	if o != nil && !IsNil(o.ReceiptsRoot) {
		return true
	}

	return false
}

// SetReceiptsRoot gets a reference to the given string and assigns it to the ReceiptsRoot field.
func (o *Web3BlockHeader) SetReceiptsRoot(v string) {
	o.ReceiptsRoot = &v
}

// GetDifficulty returns the Difficulty field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetDifficulty() string {
	if o == nil || IsNil(o.Difficulty) {
		var ret string
		return ret
	}
	return *o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetDifficultyOk() (*string, bool) {
	if o == nil || IsNil(o.Difficulty) {
		return nil, false
	}
	return o.Difficulty, true
}

// HasDifficulty returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasDifficulty() bool {
	if o != nil && !IsNil(o.Difficulty) {
		return true
	}

	return false
}

// SetDifficulty gets a reference to the given string and assigns it to the Difficulty field.
func (o *Web3BlockHeader) SetDifficulty(v string) {
	o.Difficulty = &v
}

// GetMixHash returns the MixHash field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetMixHash() string {
	if o == nil || IsNil(o.MixHash) {
		var ret string
		return ret
	}
	return *o.MixHash
}

// GetMixHashOk returns a tuple with the MixHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetMixHashOk() (*string, bool) {
	if o == nil || IsNil(o.MixHash) {
		return nil, false
	}
	return o.MixHash, true
}

// HasMixHash returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasMixHash() bool {
	if o != nil && !IsNil(o.MixHash) {
		return true
	}

	return false
}

// SetMixHash gets a reference to the given string and assigns it to the MixHash field.
func (o *Web3BlockHeader) SetMixHash(v string) {
	o.MixHash = &v
}

// GetMiner returns the Miner field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetMiner() string {
	if o == nil || IsNil(o.Miner) {
		var ret string
		return ret
	}
	return *o.Miner
}

// GetMinerOk returns a tuple with the Miner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetMinerOk() (*string, bool) {
	if o == nil || IsNil(o.Miner) {
		return nil, false
	}
	return o.Miner, true
}

// HasMiner returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasMiner() bool {
	if o != nil && !IsNil(o.Miner) {
		return true
	}

	return false
}

// SetMiner gets a reference to the given string and assigns it to the Miner field.
func (o *Web3BlockHeader) SetMiner(v string) {
	o.Miner = &v
}

// GetGasLimit returns the GasLimit field value
func (o *Web3BlockHeader) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetGasLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *Web3BlockHeader) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasUsed returns the GasUsed field value
func (o *Web3BlockHeader) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *Web3BlockHeader) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetStateRoot returns the StateRoot field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetStateRoot() string {
	if o == nil || IsNil(o.StateRoot) {
		var ret string
		return ret
	}
	return *o.StateRoot
}

// GetStateRootOk returns a tuple with the StateRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetStateRootOk() (*string, bool) {
	if o == nil || IsNil(o.StateRoot) {
		return nil, false
	}
	return o.StateRoot, true
}

// HasStateRoot returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasStateRoot() bool {
	if o != nil && !IsNil(o.StateRoot) {
		return true
	}

	return false
}

// SetStateRoot gets a reference to the given string and assigns it to the StateRoot field.
func (o *Web3BlockHeader) SetStateRoot(v string) {
	o.StateRoot = &v
}

// GetLogsBloom returns the LogsBloom field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetLogsBloom() string {
	if o == nil || IsNil(o.LogsBloom) {
		var ret string
		return ret
	}
	return *o.LogsBloom
}

// GetLogsBloomOk returns a tuple with the LogsBloom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetLogsBloomOk() (*string, bool) {
	if o == nil || IsNil(o.LogsBloom) {
		return nil, false
	}
	return o.LogsBloom, true
}

// HasLogsBloom returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasLogsBloom() bool {
	if o != nil && !IsNil(o.LogsBloom) {
		return true
	}

	return false
}

// SetLogsBloom gets a reference to the given string and assigns it to the LogsBloom field.
func (o *Web3BlockHeader) SetLogsBloom(v string) {
	o.LogsBloom = &v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetExtraData() string {
	if o == nil || IsNil(o.ExtraData) {
		var ret string
		return ret
	}
	return *o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetExtraDataOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraData) {
		return nil, false
	}
	return o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasExtraData() bool {
	if o != nil && !IsNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given string and assigns it to the ExtraData field.
func (o *Web3BlockHeader) SetExtraData(v string) {
	o.ExtraData = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *Web3BlockHeader) SetNonce(v string) {
	o.Nonce = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Web3BlockHeader) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Web3BlockHeader) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *Web3BlockHeader) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o Web3BlockHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3BlockHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.ParentHash) {
		toSerialize["parentHash"] = o.ParentHash
	}
	toSerialize["sha3Uncles"] = o.Sha3Uncles
	if !IsNil(o.TransactionsRoot) {
		toSerialize["transactionsRoot"] = o.TransactionsRoot
	}
	if !IsNil(o.ReceiptsRoot) {
		toSerialize["receiptsRoot"] = o.ReceiptsRoot
	}
	if !IsNil(o.Difficulty) {
		toSerialize["difficulty"] = o.Difficulty
	}
	if !IsNil(o.MixHash) {
		toSerialize["mixHash"] = o.MixHash
	}
	if !IsNil(o.Miner) {
		toSerialize["miner"] = o.Miner
	}
	toSerialize["gasLimit"] = o.GasLimit
	toSerialize["gasUsed"] = o.GasUsed
	if !IsNil(o.StateRoot) {
		toSerialize["stateRoot"] = o.StateRoot
	}
	if !IsNil(o.LogsBloom) {
		toSerialize["logsBloom"] = o.LogsBloom
	}
	if !IsNil(o.ExtraData) {
		toSerialize["extraData"] = o.ExtraData
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableWeb3BlockHeader struct {
	value *Web3BlockHeader
	isSet bool
}

func (v NullableWeb3BlockHeader) Get() *Web3BlockHeader {
	return v.value
}

func (v *NullableWeb3BlockHeader) Set(val *Web3BlockHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3BlockHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3BlockHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3BlockHeader(val *Web3BlockHeader) *NullableWeb3BlockHeader {
	return &NullableWeb3BlockHeader{value: val, isSet: true}
}

func (v NullableWeb3BlockHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3BlockHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


