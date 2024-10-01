/*
Hyperledger Core API

Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-core-api

import (
	"encoding/json"
)

// checks if the CactusNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CactusNode{}

// CactusNode A Cactus node can be a single server, or a set of servers behind a load balancer acting as one.
type CactusNode struct {
	NodeApiHost string `json:"nodeApiHost"`
	// The PEM encoded public key that was used to generate the JWS included in the response (the jws property)
	PublicKeyPem string `json:"publicKeyPem"`
	Id string `json:"id"`
	ConsortiumId string `json:"consortiumId"`
	MemberId string `json:"memberId"`
	// Stores an array of Ledger entity IDs that are reachable (routable) via this Cactus Node. This information is used by the client side SDK API client to figure out at runtime where to send API requests that are specific to a certain ledger such as requests to execute transactions.
	LedgerIds []string `json:"ledgerIds"`
	PluginInstanceIds []string `json:"pluginInstanceIds"`
}

// NewCactusNode instantiates a new CactusNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCactusNode(nodeApiHost string, publicKeyPem string, id string, consortiumId string, memberId string, ledgerIds []string, pluginInstanceIds []string) *CactusNode {
	this := CactusNode{}
	this.NodeApiHost = nodeApiHost
	this.PublicKeyPem = publicKeyPem
	this.Id = id
	this.ConsortiumId = consortiumId
	this.MemberId = memberId
	this.LedgerIds = ledgerIds
	this.PluginInstanceIds = pluginInstanceIds
	return &this
}

// NewCactusNodeWithDefaults instantiates a new CactusNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCactusNodeWithDefaults() *CactusNode {
	this := CactusNode{}
	return &this
}

// GetNodeApiHost returns the NodeApiHost field value
func (o *CactusNode) GetNodeApiHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeApiHost
}

// GetNodeApiHostOk returns a tuple with the NodeApiHost field value
// and a boolean to check if the value has been set.
func (o *CactusNode) GetNodeApiHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeApiHost, true
}

// SetNodeApiHost sets field value
func (o *CactusNode) SetNodeApiHost(v string) {
	o.NodeApiHost = v
}

// GetPublicKeyPem returns the PublicKeyPem field value
func (o *CactusNode) GetPublicKeyPem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyPem
}

// GetPublicKeyPemOk returns a tuple with the PublicKeyPem field value
// and a boolean to check if the value has been set.
func (o *CactusNode) GetPublicKeyPemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKeyPem, true
}

// SetPublicKeyPem sets field value
func (o *CactusNode) SetPublicKeyPem(v string) {
	o.PublicKeyPem = v
}

// GetId returns the Id field value
func (o *CactusNode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CactusNode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CactusNode) SetId(v string) {
	o.Id = v
}

// GetConsortiumId returns the ConsortiumId field value
func (o *CactusNode) GetConsortiumId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsortiumId
}

// GetConsortiumIdOk returns a tuple with the ConsortiumId field value
// and a boolean to check if the value has been set.
func (o *CactusNode) GetConsortiumIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsortiumId, true
}

// SetConsortiumId sets field value
func (o *CactusNode) SetConsortiumId(v string) {
	o.ConsortiumId = v
}

// GetMemberId returns the MemberId field value
func (o *CactusNode) GetMemberId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value
// and a boolean to check if the value has been set.
func (o *CactusNode) GetMemberIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberId, true
}

// SetMemberId sets field value
func (o *CactusNode) SetMemberId(v string) {
	o.MemberId = v
}

// GetLedgerIds returns the LedgerIds field value
func (o *CactusNode) GetLedgerIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LedgerIds
}

// GetLedgerIdsOk returns a tuple with the LedgerIds field value
// and a boolean to check if the value has been set.
func (o *CactusNode) GetLedgerIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LedgerIds, true
}

// SetLedgerIds sets field value
func (o *CactusNode) SetLedgerIds(v []string) {
	o.LedgerIds = v
}

// GetPluginInstanceIds returns the PluginInstanceIds field value
func (o *CactusNode) GetPluginInstanceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PluginInstanceIds
}

// GetPluginInstanceIdsOk returns a tuple with the PluginInstanceIds field value
// and a boolean to check if the value has been set.
func (o *CactusNode) GetPluginInstanceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginInstanceIds, true
}

// SetPluginInstanceIds sets field value
func (o *CactusNode) SetPluginInstanceIds(v []string) {
	o.PluginInstanceIds = v
}

func (o CactusNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CactusNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodeApiHost"] = o.NodeApiHost
	toSerialize["publicKeyPem"] = o.PublicKeyPem
	toSerialize["id"] = o.Id
	toSerialize["consortiumId"] = o.ConsortiumId
	toSerialize["memberId"] = o.MemberId
	toSerialize["ledgerIds"] = o.LedgerIds
	toSerialize["pluginInstanceIds"] = o.PluginInstanceIds
	return toSerialize, nil
}

type NullableCactusNode struct {
	value *CactusNode
	isSet bool
}

func (v NullableCactusNode) Get() *CactusNode {
	return v.value
}

func (v *NullableCactusNode) Set(val *CactusNode) {
	v.value = val
	v.isSet = true
}

func (v NullableCactusNode) IsSet() bool {
	return v.isSet
}

func (v *NullableCactusNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCactusNode(val *CactusNode) *NullableCactusNode {
	return &NullableCactusNode{value: val, isSet: true}
}

func (v NullableCactusNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCactusNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


