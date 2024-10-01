/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.0.0-rc.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the FileBase64 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileBase64{}

// FileBase64 Represents a file-system file that has a name and a body which holds the file contents as a Base64 encoded string
type FileBase64 struct {
	// The file's contents encoded as a Base64 string.
	Body string `json:"body"`
	// The name as referred to on a file system
	Filename string `json:"filename"`
	// The relative path of the file, if it should be placed in a sub-directory
	Filepath *string `json:"filepath,omitempty"`
}

// NewFileBase64 instantiates a new FileBase64 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileBase64(body string, filename string) *FileBase64 {
	this := FileBase64{}
	this.Body = body
	this.Filename = filename
	return &this
}

// NewFileBase64WithDefaults instantiates a new FileBase64 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileBase64WithDefaults() *FileBase64 {
	this := FileBase64{}
	return &this
}

// GetBody returns the Body field value
func (o *FileBase64) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *FileBase64) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *FileBase64) SetBody(v string) {
	o.Body = v
}

// GetFilename returns the Filename field value
func (o *FileBase64) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *FileBase64) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *FileBase64) SetFilename(v string) {
	o.Filename = v
}

// GetFilepath returns the Filepath field value if set, zero value otherwise.
func (o *FileBase64) GetFilepath() string {
	if o == nil || IsNil(o.Filepath) {
		var ret string
		return ret
	}
	return *o.Filepath
}

// GetFilepathOk returns a tuple with the Filepath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBase64) GetFilepathOk() (*string, bool) {
	if o == nil || IsNil(o.Filepath) {
		return nil, false
	}
	return o.Filepath, true
}

// HasFilepath returns a boolean if a field has been set.
func (o *FileBase64) HasFilepath() bool {
	if o != nil && !IsNil(o.Filepath) {
		return true
	}

	return false
}

// SetFilepath gets a reference to the given string and assigns it to the Filepath field.
func (o *FileBase64) SetFilepath(v string) {
	o.Filepath = &v
}

func (o FileBase64) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileBase64) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["filename"] = o.Filename
	if !IsNil(o.Filepath) {
		toSerialize["filepath"] = o.Filepath
	}
	return toSerialize, nil
}

type NullableFileBase64 struct {
	value *FileBase64
	isSet bool
}

func (v NullableFileBase64) Get() *FileBase64 {
	return v.value
}

func (v *NullableFileBase64) Set(val *FileBase64) {
	v.value = val
	v.isSet = true
}

func (v NullableFileBase64) IsSet() bool {
	return v.isSet
}

func (v *NullableFileBase64) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileBase64(val *FileBase64) *NullableFileBase64 {
	return &NullableFileBase64{value: val, isSet: true}
}

func (v NullableFileBase64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileBase64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


