/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-satp-hermes

import (
	"encoding/json"
)

// checks if the RecoverV1Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverV1Message{}

// RecoverV1Message struct for RecoverV1Message
type RecoverV1Message struct {
	SessionID string `json:"sessionID"`
	OdapPhase string `json:"odapPhase"`
	SequenceNumber float32 `json:"sequenceNumber"`
	LastLogEntryTimestamp string `json:"lastLogEntryTimestamp"`
	IsBackup bool `json:"isBackup"`
	NewBasePath string `json:"newBasePath"`
	NewGatewayPubKey *string `json:"newGatewayPubKey,omitempty"`
	Signature string `json:"signature"`
}

// NewRecoverV1Message instantiates a new RecoverV1Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverV1Message(sessionID string, odapPhase string, sequenceNumber float32, lastLogEntryTimestamp string, isBackup bool, newBasePath string, signature string) *RecoverV1Message {
	this := RecoverV1Message{}
	this.SessionID = sessionID
	this.OdapPhase = odapPhase
	this.SequenceNumber = sequenceNumber
	this.LastLogEntryTimestamp = lastLogEntryTimestamp
	this.IsBackup = isBackup
	this.NewBasePath = newBasePath
	this.Signature = signature
	return &this
}

// NewRecoverV1MessageWithDefaults instantiates a new RecoverV1Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverV1MessageWithDefaults() *RecoverV1Message {
	this := RecoverV1Message{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *RecoverV1Message) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *RecoverV1Message) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *RecoverV1Message) SetSessionID(v string) {
	o.SessionID = v
}

// GetOdapPhase returns the OdapPhase field value
func (o *RecoverV1Message) GetOdapPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdapPhase
}

// GetOdapPhaseOk returns a tuple with the OdapPhase field value
// and a boolean to check if the value has been set.
func (o *RecoverV1Message) GetOdapPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdapPhase, true
}

// SetOdapPhase sets field value
func (o *RecoverV1Message) SetOdapPhase(v string) {
	o.OdapPhase = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *RecoverV1Message) GetSequenceNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *RecoverV1Message) GetSequenceNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *RecoverV1Message) SetSequenceNumber(v float32) {
	o.SequenceNumber = v
}

// GetLastLogEntryTimestamp returns the LastLogEntryTimestamp field value
func (o *RecoverV1Message) GetLastLogEntryTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastLogEntryTimestamp
}

// GetLastLogEntryTimestampOk returns a tuple with the LastLogEntryTimestamp field value
// and a boolean to check if the value has been set.
func (o *RecoverV1Message) GetLastLogEntryTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastLogEntryTimestamp, true
}

// SetLastLogEntryTimestamp sets field value
func (o *RecoverV1Message) SetLastLogEntryTimestamp(v string) {
	o.LastLogEntryTimestamp = v
}

// GetIsBackup returns the IsBackup field value
func (o *RecoverV1Message) GetIsBackup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBackup
}

// GetIsBackupOk returns a tuple with the IsBackup field value
// and a boolean to check if the value has been set.
func (o *RecoverV1Message) GetIsBackupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBackup, true
}

// SetIsBackup sets field value
func (o *RecoverV1Message) SetIsBackup(v bool) {
	o.IsBackup = v
}

// GetNewBasePath returns the NewBasePath field value
func (o *RecoverV1Message) GetNewBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewBasePath
}

// GetNewBasePathOk returns a tuple with the NewBasePath field value
// and a boolean to check if the value has been set.
func (o *RecoverV1Message) GetNewBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewBasePath, true
}

// SetNewBasePath sets field value
func (o *RecoverV1Message) SetNewBasePath(v string) {
	o.NewBasePath = v
}

// GetNewGatewayPubKey returns the NewGatewayPubKey field value if set, zero value otherwise.
func (o *RecoverV1Message) GetNewGatewayPubKey() string {
	if o == nil || IsNil(o.NewGatewayPubKey) {
		var ret string
		return ret
	}
	return *o.NewGatewayPubKey
}

// GetNewGatewayPubKeyOk returns a tuple with the NewGatewayPubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverV1Message) GetNewGatewayPubKeyOk() (*string, bool) {
	if o == nil || IsNil(o.NewGatewayPubKey) {
		return nil, false
	}
	return o.NewGatewayPubKey, true
}

// HasNewGatewayPubKey returns a boolean if a field has been set.
func (o *RecoverV1Message) HasNewGatewayPubKey() bool {
	if o != nil && !IsNil(o.NewGatewayPubKey) {
		return true
	}

	return false
}

// SetNewGatewayPubKey gets a reference to the given string and assigns it to the NewGatewayPubKey field.
func (o *RecoverV1Message) SetNewGatewayPubKey(v string) {
	o.NewGatewayPubKey = &v
}

// GetSignature returns the Signature field value
func (o *RecoverV1Message) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *RecoverV1Message) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *RecoverV1Message) SetSignature(v string) {
	o.Signature = v
}

func (o RecoverV1Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverV1Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["odapPhase"] = o.OdapPhase
	toSerialize["sequenceNumber"] = o.SequenceNumber
	toSerialize["lastLogEntryTimestamp"] = o.LastLogEntryTimestamp
	toSerialize["isBackup"] = o.IsBackup
	toSerialize["newBasePath"] = o.NewBasePath
	if !IsNil(o.NewGatewayPubKey) {
		toSerialize["newGatewayPubKey"] = o.NewGatewayPubKey
	}
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

type NullableRecoverV1Message struct {
	value *RecoverV1Message
	isSet bool
}

func (v NullableRecoverV1Message) Get() *RecoverV1Message {
	return v.value
}

func (v *NullableRecoverV1Message) Set(val *RecoverV1Message) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverV1Message) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverV1Message) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverV1Message(val *RecoverV1Message) *NullableRecoverV1Message {
	return &NullableRecoverV1Message{value: val, isSet: true}
}

func (v NullableRecoverV1Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverV1Message) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


