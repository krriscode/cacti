/*
Hyperledger Cactus Plugin - Connector Polkadot

Can perform basic tasks on a Polkadot parachain

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-polkadot

import (
	"encoding/json"
)

// checks if the ErrorExceptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorExceptionResponse{}

// ErrorExceptionResponse Error response from the connector.
type ErrorExceptionResponse struct {
	// Short error description message.
	Message string `json:"message"`
	// Detailed error information.
	Error string `json:"error"`
}

// NewErrorExceptionResponse instantiates a new ErrorExceptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorExceptionResponse(message string, error_ string) *ErrorExceptionResponse {
	this := ErrorExceptionResponse{}
	this.Message = message
	this.Error = error_
	return &this
}

// NewErrorExceptionResponseWithDefaults instantiates a new ErrorExceptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorExceptionResponseWithDefaults() *ErrorExceptionResponse {
	this := ErrorExceptionResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *ErrorExceptionResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorExceptionResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorExceptionResponse) SetMessage(v string) {
	o.Message = v
}

// GetError returns the Error field value
func (o *ErrorExceptionResponse) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ErrorExceptionResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ErrorExceptionResponse) SetError(v string) {
	o.Error = v
}

func (o ErrorExceptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorExceptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableErrorExceptionResponse struct {
	value *ErrorExceptionResponse
	isSet bool
}

func (v NullableErrorExceptionResponse) Get() *ErrorExceptionResponse {
	return v.value
}

func (v *NullableErrorExceptionResponse) Set(val *ErrorExceptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorExceptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorExceptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorExceptionResponse(val *ErrorExceptionResponse) *NullableErrorExceptionResponse {
	return &NullableErrorExceptionResponse{value: val, isSet: true}
}

func (v NullableErrorExceptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorExceptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


