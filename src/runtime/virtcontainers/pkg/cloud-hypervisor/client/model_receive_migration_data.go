/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ReceiveMigrationData struct for ReceiveMigrationData
type ReceiveMigrationData struct {
	ReceiverUrl string `json:"receiver_url"`
}

// NewReceiveMigrationData instantiates a new ReceiveMigrationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiveMigrationData(receiverUrl string) *ReceiveMigrationData {
	this := ReceiveMigrationData{}
	this.ReceiverUrl = receiverUrl
	return &this
}

// NewReceiveMigrationDataWithDefaults instantiates a new ReceiveMigrationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiveMigrationDataWithDefaults() *ReceiveMigrationData {
	this := ReceiveMigrationData{}
	return &this
}

// GetReceiverUrl returns the ReceiverUrl field value
func (o *ReceiveMigrationData) GetReceiverUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiverUrl
}

// GetReceiverUrlOk returns a tuple with the ReceiverUrl field value
// and a boolean to check if the value has been set.
func (o *ReceiveMigrationData) GetReceiverUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverUrl, true
}

// SetReceiverUrl sets field value
func (o *ReceiveMigrationData) SetReceiverUrl(v string) {
	o.ReceiverUrl = v
}

func (o ReceiveMigrationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["receiver_url"] = o.ReceiverUrl
	}
	return json.Marshal(toSerialize)
}

type NullableReceiveMigrationData struct {
	value *ReceiveMigrationData
	isSet bool
}

func (v NullableReceiveMigrationData) Get() *ReceiveMigrationData {
	return v.value
}

func (v *NullableReceiveMigrationData) Set(val *ReceiveMigrationData) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiveMigrationData) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiveMigrationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiveMigrationData(val *ReceiveMigrationData) *NullableReceiveMigrationData {
	return &NullableReceiveMigrationData{value: val, isSet: true}
}

func (v NullableReceiveMigrationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiveMigrationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
