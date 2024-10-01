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

// checks if the TransferInitializationV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferInitializationV1Response{}

// TransferInitializationV1Response struct for TransferInitializationV1Response
type TransferInitializationV1Response struct {
	MessageType string `json:"messageType"`
	SessionID string `json:"sessionID"`
	SequenceNumber float32 `json:"sequenceNumber"`
	OdapPhase *string `json:"odapPhase,omitempty"`
	InitialRequestMessageHash string `json:"initialRequestMessageHash"`
	Destination *string `json:"destination,omitempty"`
	TimeStamp string `json:"timeStamp"`
	ProcessedTimeStamp string `json:"processedTimeStamp"`
	ServerIdentityPubkey string `json:"serverIdentityPubkey"`
	Signature string `json:"signature"`
	BackupGatewaysAllowed []string `json:"backupGatewaysAllowed"`
}

// NewTransferInitializationV1Response instantiates a new TransferInitializationV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferInitializationV1Response(messageType string, sessionID string, sequenceNumber float32, initialRequestMessageHash string, timeStamp string, processedTimeStamp string, serverIdentityPubkey string, signature string, backupGatewaysAllowed []string) *TransferInitializationV1Response {
	this := TransferInitializationV1Response{}
	this.MessageType = messageType
	this.SessionID = sessionID
	this.SequenceNumber = sequenceNumber
	this.InitialRequestMessageHash = initialRequestMessageHash
	this.TimeStamp = timeStamp
	this.ProcessedTimeStamp = processedTimeStamp
	this.ServerIdentityPubkey = serverIdentityPubkey
	this.Signature = signature
	this.BackupGatewaysAllowed = backupGatewaysAllowed
	return &this
}

// NewTransferInitializationV1ResponseWithDefaults instantiates a new TransferInitializationV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferInitializationV1ResponseWithDefaults() *TransferInitializationV1Response {
	this := TransferInitializationV1Response{}
	return &this
}

// GetMessageType returns the MessageType field value
func (o *TransferInitializationV1Response) GetMessageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *TransferInitializationV1Response) SetMessageType(v string) {
	o.MessageType = v
}

// GetSessionID returns the SessionID field value
func (o *TransferInitializationV1Response) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *TransferInitializationV1Response) SetSessionID(v string) {
	o.SessionID = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *TransferInitializationV1Response) GetSequenceNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetSequenceNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *TransferInitializationV1Response) SetSequenceNumber(v float32) {
	o.SequenceNumber = v
}

// GetOdapPhase returns the OdapPhase field value if set, zero value otherwise.
func (o *TransferInitializationV1Response) GetOdapPhase() string {
	if o == nil || IsNil(o.OdapPhase) {
		var ret string
		return ret
	}
	return *o.OdapPhase
}

// GetOdapPhaseOk returns a tuple with the OdapPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetOdapPhaseOk() (*string, bool) {
	if o == nil || IsNil(o.OdapPhase) {
		return nil, false
	}
	return o.OdapPhase, true
}

// HasOdapPhase returns a boolean if a field has been set.
func (o *TransferInitializationV1Response) HasOdapPhase() bool {
	if o != nil && !IsNil(o.OdapPhase) {
		return true
	}

	return false
}

// SetOdapPhase gets a reference to the given string and assigns it to the OdapPhase field.
func (o *TransferInitializationV1Response) SetOdapPhase(v string) {
	o.OdapPhase = &v
}

// GetInitialRequestMessageHash returns the InitialRequestMessageHash field value
func (o *TransferInitializationV1Response) GetInitialRequestMessageHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialRequestMessageHash
}

// GetInitialRequestMessageHashOk returns a tuple with the InitialRequestMessageHash field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetInitialRequestMessageHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialRequestMessageHash, true
}

// SetInitialRequestMessageHash sets field value
func (o *TransferInitializationV1Response) SetInitialRequestMessageHash(v string) {
	o.InitialRequestMessageHash = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *TransferInitializationV1Response) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *TransferInitializationV1Response) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *TransferInitializationV1Response) SetDestination(v string) {
	o.Destination = &v
}

// GetTimeStamp returns the TimeStamp field value
func (o *TransferInitializationV1Response) GetTimeStamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetTimeStampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *TransferInitializationV1Response) SetTimeStamp(v string) {
	o.TimeStamp = v
}

// GetProcessedTimeStamp returns the ProcessedTimeStamp field value
func (o *TransferInitializationV1Response) GetProcessedTimeStamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessedTimeStamp
}

// GetProcessedTimeStampOk returns a tuple with the ProcessedTimeStamp field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetProcessedTimeStampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessedTimeStamp, true
}

// SetProcessedTimeStamp sets field value
func (o *TransferInitializationV1Response) SetProcessedTimeStamp(v string) {
	o.ProcessedTimeStamp = v
}

// GetServerIdentityPubkey returns the ServerIdentityPubkey field value
func (o *TransferInitializationV1Response) GetServerIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIdentityPubkey
}

// GetServerIdentityPubkeyOk returns a tuple with the ServerIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetServerIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIdentityPubkey, true
}

// SetServerIdentityPubkey sets field value
func (o *TransferInitializationV1Response) SetServerIdentityPubkey(v string) {
	o.ServerIdentityPubkey = v
}

// GetSignature returns the Signature field value
func (o *TransferInitializationV1Response) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *TransferInitializationV1Response) SetSignature(v string) {
	o.Signature = v
}

// GetBackupGatewaysAllowed returns the BackupGatewaysAllowed field value
func (o *TransferInitializationV1Response) GetBackupGatewaysAllowed() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BackupGatewaysAllowed
}

// GetBackupGatewaysAllowedOk returns a tuple with the BackupGatewaysAllowed field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Response) GetBackupGatewaysAllowedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupGatewaysAllowed, true
}

// SetBackupGatewaysAllowed sets field value
func (o *TransferInitializationV1Response) SetBackupGatewaysAllowed(v []string) {
	o.BackupGatewaysAllowed = v
}

func (o TransferInitializationV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferInitializationV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messageType"] = o.MessageType
	toSerialize["sessionID"] = o.SessionID
	toSerialize["sequenceNumber"] = o.SequenceNumber
	if !IsNil(o.OdapPhase) {
		toSerialize["odapPhase"] = o.OdapPhase
	}
	toSerialize["initialRequestMessageHash"] = o.InitialRequestMessageHash
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	toSerialize["timeStamp"] = o.TimeStamp
	toSerialize["processedTimeStamp"] = o.ProcessedTimeStamp
	toSerialize["serverIdentityPubkey"] = o.ServerIdentityPubkey
	toSerialize["signature"] = o.Signature
	toSerialize["backupGatewaysAllowed"] = o.BackupGatewaysAllowed
	return toSerialize, nil
}

type NullableTransferInitializationV1Response struct {
	value *TransferInitializationV1Response
	isSet bool
}

func (v NullableTransferInitializationV1Response) Get() *TransferInitializationV1Response {
	return v.value
}

func (v *NullableTransferInitializationV1Response) Set(val *TransferInitializationV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferInitializationV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferInitializationV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferInitializationV1Response(val *TransferInitializationV1Response) *NullableTransferInitializationV1Response {
	return &NullableTransferInitializationV1Response{value: val, isSet: true}
}

func (v NullableTransferInitializationV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferInitializationV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


