/*
Hyperledger Cacti Plugin - Connector CDL

Can perform basic tasks on Fujitsu CDL service.

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-cdl

import (
	"encoding/json"
)

// checks if the TrailEventDetailsV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrailEventDetailsV1{}

// TrailEventDetailsV1 Details of newly created CDL event.
type TrailEventDetailsV1 struct {
	CdlEvent interface{} `json:"cdl:Event,omitempty"`
	CdlLineage EventLineageV1 `json:"cdl:Lineage"`
	CdlTags interface{} `json:"cdl:Tags"`
	CdlVerification interface{} `json:"cdl:Verification"`
}

// NewTrailEventDetailsV1 instantiates a new TrailEventDetailsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrailEventDetailsV1(cdlLineage EventLineageV1, cdlTags interface{}, cdlVerification interface{}) *TrailEventDetailsV1 {
	this := TrailEventDetailsV1{}
	this.CdlLineage = cdlLineage
	this.CdlTags = cdlTags
	this.CdlVerification = cdlVerification
	return &this
}

// NewTrailEventDetailsV1WithDefaults instantiates a new TrailEventDetailsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrailEventDetailsV1WithDefaults() *TrailEventDetailsV1 {
	this := TrailEventDetailsV1{}
	return &this
}

// GetCdlEvent returns the CdlEvent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailEventDetailsV1) GetCdlEvent() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CdlEvent
}

// GetCdlEventOk returns a tuple with the CdlEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailEventDetailsV1) GetCdlEventOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CdlEvent) {
		return nil, false
	}
	return &o.CdlEvent, true
}

// HasCdlEvent returns a boolean if a field has been set.
func (o *TrailEventDetailsV1) HasCdlEvent() bool {
	if o != nil && IsNil(o.CdlEvent) {
		return true
	}

	return false
}

// SetCdlEvent gets a reference to the given interface{} and assigns it to the CdlEvent field.
func (o *TrailEventDetailsV1) SetCdlEvent(v interface{}) {
	o.CdlEvent = v
}

// GetCdlLineage returns the CdlLineage field value
func (o *TrailEventDetailsV1) GetCdlLineage() EventLineageV1 {
	if o == nil {
		var ret EventLineageV1
		return ret
	}

	return o.CdlLineage
}

// GetCdlLineageOk returns a tuple with the CdlLineage field value
// and a boolean to check if the value has been set.
func (o *TrailEventDetailsV1) GetCdlLineageOk() (*EventLineageV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CdlLineage, true
}

// SetCdlLineage sets field value
func (o *TrailEventDetailsV1) SetCdlLineage(v EventLineageV1) {
	o.CdlLineage = v
}

// GetCdlTags returns the CdlTags field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TrailEventDetailsV1) GetCdlTags() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.CdlTags
}

// GetCdlTagsOk returns a tuple with the CdlTags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailEventDetailsV1) GetCdlTagsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CdlTags) {
		return nil, false
	}
	return &o.CdlTags, true
}

// SetCdlTags sets field value
func (o *TrailEventDetailsV1) SetCdlTags(v interface{}) {
	o.CdlTags = v
}

// GetCdlVerification returns the CdlVerification field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TrailEventDetailsV1) GetCdlVerification() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.CdlVerification
}

// GetCdlVerificationOk returns a tuple with the CdlVerification field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailEventDetailsV1) GetCdlVerificationOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CdlVerification) {
		return nil, false
	}
	return &o.CdlVerification, true
}

// SetCdlVerification sets field value
func (o *TrailEventDetailsV1) SetCdlVerification(v interface{}) {
	o.CdlVerification = v
}

func (o TrailEventDetailsV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrailEventDetailsV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CdlEvent != nil {
		toSerialize["cdl:Event"] = o.CdlEvent
	}
	toSerialize["cdl:Lineage"] = o.CdlLineage
	if o.CdlTags != nil {
		toSerialize["cdl:Tags"] = o.CdlTags
	}
	if o.CdlVerification != nil {
		toSerialize["cdl:Verification"] = o.CdlVerification
	}
	return toSerialize, nil
}

type NullableTrailEventDetailsV1 struct {
	value *TrailEventDetailsV1
	isSet bool
}

func (v NullableTrailEventDetailsV1) Get() *TrailEventDetailsV1 {
	return v.value
}

func (v *NullableTrailEventDetailsV1) Set(val *TrailEventDetailsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableTrailEventDetailsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableTrailEventDetailsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrailEventDetailsV1(val *TrailEventDetailsV1) *NullableTrailEventDetailsV1 {
	return &NullableTrailEventDetailsV1{value: val, isSet: true}
}

func (v NullableTrailEventDetailsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrailEventDetailsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


