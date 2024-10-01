/*
Hyperledger Cactus API

Interact with a Cactus deployment through HTTP.

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-cmd-api-server

import (
	"encoding/json"
)

// checks if the HealthCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckResponse{}

// HealthCheckResponse struct for HealthCheckResponse
type HealthCheckResponse struct {
	Success *bool `json:"success,omitempty"`
	CreatedAt string `json:"createdAt"`
	MemoryUsage MemoryUsage `json:"memoryUsage"`
}

// NewHealthCheckResponse instantiates a new HealthCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckResponse(createdAt string, memoryUsage MemoryUsage) *HealthCheckResponse {
	this := HealthCheckResponse{}
	this.CreatedAt = createdAt
	this.MemoryUsage = memoryUsage
	return &this
}

// NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckResponseWithDefaults() *HealthCheckResponse {
	this := HealthCheckResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *HealthCheckResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *HealthCheckResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *HealthCheckResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *HealthCheckResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetMemoryUsage returns the MemoryUsage field value
func (o *HealthCheckResponse) GetMemoryUsage() MemoryUsage {
	if o == nil {
		var ret MemoryUsage
		return ret
	}

	return o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetMemoryUsageOk() (*MemoryUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryUsage, true
}

// SetMemoryUsage sets field value
func (o *HealthCheckResponse) SetMemoryUsage(v MemoryUsage) {
	o.MemoryUsage = v
}

func (o HealthCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["memoryUsage"] = o.MemoryUsage
	return toSerialize, nil
}

type NullableHealthCheckResponse struct {
	value *HealthCheckResponse
	isSet bool
}

func (v NullableHealthCheckResponse) Get() *HealthCheckResponse {
	return v.value
}

func (v *NullableHealthCheckResponse) Set(val *HealthCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckResponse(val *HealthCheckResponse) *NullableHealthCheckResponse {
	return &NullableHealthCheckResponse{value: val, isSet: true}
}

func (v NullableHealthCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


