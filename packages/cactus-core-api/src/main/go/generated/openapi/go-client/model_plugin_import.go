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

// checks if the PluginImport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginImport{}

// PluginImport struct for PluginImport
type PluginImport struct {
	PackageName string `json:"packageName"`
	Type PluginImportType `json:"type"`
	Action PluginImportAction `json:"action"`
	Options interface{} `json:"options,omitempty"`
}

// NewPluginImport instantiates a new PluginImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginImport(packageName string, type_ PluginImportType, action PluginImportAction) *PluginImport {
	this := PluginImport{}
	this.PackageName = packageName
	this.Type = type_
	this.Action = action
	return &this
}

// NewPluginImportWithDefaults instantiates a new PluginImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginImportWithDefaults() *PluginImport {
	this := PluginImport{}
	return &this
}

// GetPackageName returns the PackageName field value
func (o *PluginImport) GetPackageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value
// and a boolean to check if the value has been set.
func (o *PluginImport) GetPackageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageName, true
}

// SetPackageName sets field value
func (o *PluginImport) SetPackageName(v string) {
	o.PackageName = v
}

// GetType returns the Type field value
func (o *PluginImport) GetType() PluginImportType {
	if o == nil {
		var ret PluginImportType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PluginImport) GetTypeOk() (*PluginImportType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PluginImport) SetType(v PluginImportType) {
	o.Type = v
}

// GetAction returns the Action field value
func (o *PluginImport) GetAction() PluginImportAction {
	if o == nil {
		var ret PluginImportAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PluginImport) GetActionOk() (*PluginImportAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PluginImport) SetAction(v PluginImportAction) {
	o.Action = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginImport) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginImport) GetOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PluginImport) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *PluginImport) SetOptions(v interface{}) {
	o.Options = v
}

func (o PluginImport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginImport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packageName"] = o.PackageName
	toSerialize["type"] = o.Type
	toSerialize["action"] = o.Action
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullablePluginImport struct {
	value *PluginImport
	isSet bool
}

func (v NullablePluginImport) Get() *PluginImport {
	return v.value
}

func (v *NullablePluginImport) Set(val *PluginImport) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginImport) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginImport(val *PluginImport) *NullablePluginImport {
	return &NullablePluginImport{value: val, isSet: true}
}

func (v NullablePluginImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


