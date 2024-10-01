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

// checks if the Web3TransactionReceipt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3TransactionReceipt{}

// Web3TransactionReceipt struct for Web3TransactionReceipt
type Web3TransactionReceipt struct {
	Status bool `json:"status"`
	TransactionHash string `json:"transactionHash"`
	TransactionIndex string `json:"transactionIndex"`
	BlockHash string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	GasUsed string `json:"gasUsed"`
	EffectiveGasPrice *string `json:"effectiveGasPrice,omitempty"`
	ContractAddress NullableString `json:"contractAddress,omitempty"`
	From string `json:"from"`
	To string `json:"to"`
	Logs []interface{} `json:"logs,omitempty"`
	LogsBloom *string `json:"logsBloom,omitempty"`
	RevertReason *string `json:"revertReason,omitempty"`
	Output *string `json:"output,omitempty"`
	CommitmentHash *string `json:"commitmentHash,omitempty"`
	CumulativeGasUsed *string `json:"cumulativeGasUsed,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3TransactionReceipt Web3TransactionReceipt

// NewWeb3TransactionReceipt instantiates a new Web3TransactionReceipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3TransactionReceipt(status bool, transactionHash string, transactionIndex string, blockHash string, blockNumber string, gasUsed string, from string, to string) *Web3TransactionReceipt {
	this := Web3TransactionReceipt{}
	this.Status = status
	this.TransactionHash = transactionHash
	this.TransactionIndex = transactionIndex
	this.BlockHash = blockHash
	this.BlockNumber = blockNumber
	this.GasUsed = gasUsed
	this.From = from
	this.To = to
	return &this
}

// NewWeb3TransactionReceiptWithDefaults instantiates a new Web3TransactionReceipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3TransactionReceiptWithDefaults() *Web3TransactionReceipt {
	this := Web3TransactionReceipt{}
	return &this
}

// GetStatus returns the Status field value
func (o *Web3TransactionReceipt) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Web3TransactionReceipt) SetStatus(v bool) {
	o.Status = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *Web3TransactionReceipt) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *Web3TransactionReceipt) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionIndex returns the TransactionIndex field value
func (o *Web3TransactionReceipt) GetTransactionIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetTransactionIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIndex, true
}

// SetTransactionIndex sets field value
func (o *Web3TransactionReceipt) SetTransactionIndex(v string) {
	o.TransactionIndex = v
}

// GetBlockHash returns the BlockHash field value
func (o *Web3TransactionReceipt) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *Web3TransactionReceipt) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockNumber returns the BlockNumber field value
func (o *Web3TransactionReceipt) GetBlockNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetBlockNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockNumber, true
}

// SetBlockNumber sets field value
func (o *Web3TransactionReceipt) SetBlockNumber(v string) {
	o.BlockNumber = v
}

// GetGasUsed returns the GasUsed field value
func (o *Web3TransactionReceipt) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *Web3TransactionReceipt) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetEffectiveGasPrice returns the EffectiveGasPrice field value if set, zero value otherwise.
func (o *Web3TransactionReceipt) GetEffectiveGasPrice() string {
	if o == nil || IsNil(o.EffectiveGasPrice) {
		var ret string
		return ret
	}
	return *o.EffectiveGasPrice
}

// GetEffectiveGasPriceOk returns a tuple with the EffectiveGasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetEffectiveGasPriceOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveGasPrice) {
		return nil, false
	}
	return o.EffectiveGasPrice, true
}

// HasEffectiveGasPrice returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasEffectiveGasPrice() bool {
	if o != nil && !IsNil(o.EffectiveGasPrice) {
		return true
	}

	return false
}

// SetEffectiveGasPrice gets a reference to the given string and assigns it to the EffectiveGasPrice field.
func (o *Web3TransactionReceipt) SetEffectiveGasPrice(v string) {
	o.EffectiveGasPrice = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TransactionReceipt) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TransactionReceipt) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *Web3TransactionReceipt) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *Web3TransactionReceipt) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *Web3TransactionReceipt) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetFrom returns the From field value
func (o *Web3TransactionReceipt) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *Web3TransactionReceipt) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *Web3TransactionReceipt) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *Web3TransactionReceipt) SetTo(v string) {
	o.To = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *Web3TransactionReceipt) GetLogs() []interface{} {
	if o == nil || IsNil(o.Logs) {
		var ret []interface{}
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetLogsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []interface{} and assigns it to the Logs field.
func (o *Web3TransactionReceipt) SetLogs(v []interface{}) {
	o.Logs = v
}

// GetLogsBloom returns the LogsBloom field value if set, zero value otherwise.
func (o *Web3TransactionReceipt) GetLogsBloom() string {
	if o == nil || IsNil(o.LogsBloom) {
		var ret string
		return ret
	}
	return *o.LogsBloom
}

// GetLogsBloomOk returns a tuple with the LogsBloom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetLogsBloomOk() (*string, bool) {
	if o == nil || IsNil(o.LogsBloom) {
		return nil, false
	}
	return o.LogsBloom, true
}

// HasLogsBloom returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasLogsBloom() bool {
	if o != nil && !IsNil(o.LogsBloom) {
		return true
	}

	return false
}

// SetLogsBloom gets a reference to the given string and assigns it to the LogsBloom field.
func (o *Web3TransactionReceipt) SetLogsBloom(v string) {
	o.LogsBloom = &v
}

// GetRevertReason returns the RevertReason field value if set, zero value otherwise.
func (o *Web3TransactionReceipt) GetRevertReason() string {
	if o == nil || IsNil(o.RevertReason) {
		var ret string
		return ret
	}
	return *o.RevertReason
}

// GetRevertReasonOk returns a tuple with the RevertReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetRevertReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RevertReason) {
		return nil, false
	}
	return o.RevertReason, true
}

// HasRevertReason returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasRevertReason() bool {
	if o != nil && !IsNil(o.RevertReason) {
		return true
	}

	return false
}

// SetRevertReason gets a reference to the given string and assigns it to the RevertReason field.
func (o *Web3TransactionReceipt) SetRevertReason(v string) {
	o.RevertReason = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *Web3TransactionReceipt) GetOutput() string {
	if o == nil || IsNil(o.Output) {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetOutputOk() (*string, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *Web3TransactionReceipt) SetOutput(v string) {
	o.Output = &v
}

// GetCommitmentHash returns the CommitmentHash field value if set, zero value otherwise.
func (o *Web3TransactionReceipt) GetCommitmentHash() string {
	if o == nil || IsNil(o.CommitmentHash) {
		var ret string
		return ret
	}
	return *o.CommitmentHash
}

// GetCommitmentHashOk returns a tuple with the CommitmentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetCommitmentHashOk() (*string, bool) {
	if o == nil || IsNil(o.CommitmentHash) {
		return nil, false
	}
	return o.CommitmentHash, true
}

// HasCommitmentHash returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasCommitmentHash() bool {
	if o != nil && !IsNil(o.CommitmentHash) {
		return true
	}

	return false
}

// SetCommitmentHash gets a reference to the given string and assigns it to the CommitmentHash field.
func (o *Web3TransactionReceipt) SetCommitmentHash(v string) {
	o.CommitmentHash = &v
}

// GetCumulativeGasUsed returns the CumulativeGasUsed field value if set, zero value otherwise.
func (o *Web3TransactionReceipt) GetCumulativeGasUsed() string {
	if o == nil || IsNil(o.CumulativeGasUsed) {
		var ret string
		return ret
	}
	return *o.CumulativeGasUsed
}

// GetCumulativeGasUsedOk returns a tuple with the CumulativeGasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetCumulativeGasUsedOk() (*string, bool) {
	if o == nil || IsNil(o.CumulativeGasUsed) {
		return nil, false
	}
	return o.CumulativeGasUsed, true
}

// HasCumulativeGasUsed returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasCumulativeGasUsed() bool {
	if o != nil && !IsNil(o.CumulativeGasUsed) {
		return true
	}

	return false
}

// SetCumulativeGasUsed gets a reference to the given string and assigns it to the CumulativeGasUsed field.
func (o *Web3TransactionReceipt) SetCumulativeGasUsed(v string) {
	o.CumulativeGasUsed = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Web3TransactionReceipt) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3TransactionReceipt) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Web3TransactionReceipt) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Web3TransactionReceipt) SetType(v string) {
	o.Type = &v
}

func (o Web3TransactionReceipt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3TransactionReceipt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["transactionHash"] = o.TransactionHash
	toSerialize["transactionIndex"] = o.TransactionIndex
	toSerialize["blockHash"] = o.BlockHash
	toSerialize["blockNumber"] = o.BlockNumber
	toSerialize["gasUsed"] = o.GasUsed
	if !IsNil(o.EffectiveGasPrice) {
		toSerialize["effectiveGasPrice"] = o.EffectiveGasPrice
	}
	if o.ContractAddress.IsSet() {
		toSerialize["contractAddress"] = o.ContractAddress.Get()
	}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.LogsBloom) {
		toSerialize["logsBloom"] = o.LogsBloom
	}
	if !IsNil(o.RevertReason) {
		toSerialize["revertReason"] = o.RevertReason
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.CommitmentHash) {
		toSerialize["commitmentHash"] = o.CommitmentHash
	}
	if !IsNil(o.CumulativeGasUsed) {
		toSerialize["cumulativeGasUsed"] = o.CumulativeGasUsed
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3TransactionReceipt) UnmarshalJSON(bytes []byte) (err error) {
	varWeb3TransactionReceipt := _Web3TransactionReceipt{}

	if err = json.Unmarshal(bytes, &varWeb3TransactionReceipt); err == nil {
		*o = Web3TransactionReceipt(varWeb3TransactionReceipt)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "transactionHash")
		delete(additionalProperties, "transactionIndex")
		delete(additionalProperties, "blockHash")
		delete(additionalProperties, "blockNumber")
		delete(additionalProperties, "gasUsed")
		delete(additionalProperties, "effectiveGasPrice")
		delete(additionalProperties, "contractAddress")
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "logs")
		delete(additionalProperties, "logsBloom")
		delete(additionalProperties, "revertReason")
		delete(additionalProperties, "output")
		delete(additionalProperties, "commitmentHash")
		delete(additionalProperties, "cumulativeGasUsed")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3TransactionReceipt struct {
	value *Web3TransactionReceipt
	isSet bool
}

func (v NullableWeb3TransactionReceipt) Get() *Web3TransactionReceipt {
	return v.value
}

func (v *NullableWeb3TransactionReceipt) Set(val *Web3TransactionReceipt) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3TransactionReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3TransactionReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3TransactionReceipt(val *Web3TransactionReceipt) *NullableWeb3TransactionReceipt {
	return &NullableWeb3TransactionReceipt{value: val, isSet: true}
}

func (v NullableWeb3TransactionReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3TransactionReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


