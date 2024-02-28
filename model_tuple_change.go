/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://openfga.dev/community
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"bytes"

	"encoding/json"
	"time"
)

// TupleChange struct for TupleChange
type TupleChange struct {
	TupleKey  TupleKey       `json:"tuple_key"yaml:"tuple_key"`
	Operation TupleOperation `json:"operation"yaml:"operation"`
	Timestamp time.Time      `json:"timestamp"yaml:"timestamp"`
}

// NewTupleChange instantiates a new TupleChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTupleChange(tupleKey TupleKey, operation TupleOperation, timestamp time.Time) *TupleChange {
	this := TupleChange{}
	this.TupleKey = tupleKey
	this.Operation = operation
	this.Timestamp = timestamp
	return &this
}

// NewTupleChangeWithDefaults instantiates a new TupleChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleChangeWithDefaults() *TupleChange {
	this := TupleChange{}
	var operation TupleOperation = WRITE
	this.Operation = operation
	return &this
}

// GetTupleKey returns the TupleKey field value
func (o *TupleChange) GetTupleKey() TupleKey {
	if o == nil {
		var ret TupleKey
		return ret
	}

	return o.TupleKey
}

// GetTupleKeyOk returns a tuple with the TupleKey field value
// and a boolean to check if the value has been set.
func (o *TupleChange) GetTupleKeyOk() (*TupleKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TupleKey, true
}

// SetTupleKey sets field value
func (o *TupleChange) SetTupleKey(v TupleKey) {
	o.TupleKey = v
}

// GetOperation returns the Operation field value
func (o *TupleChange) GetOperation() TupleOperation {
	if o == nil {
		var ret TupleOperation
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *TupleChange) GetOperationOk() (*TupleOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *TupleChange) SetOperation(v TupleOperation) {
	o.Operation = v
}

// GetTimestamp returns the Timestamp field value
func (o *TupleChange) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *TupleChange) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *TupleChange) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o TupleChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuple_key"] = o.TupleKey
	toSerialize["operation"] = o.Operation
	toSerialize["timestamp"] = o.Timestamp
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableTupleChange struct {
	value *TupleChange
	isSet bool
}

func (v NullableTupleChange) Get() *TupleChange {
	return v.value
}

func (v *NullableTupleChange) Set(val *TupleChange) {
	v.value = val
	v.isSet = true
}

func (v NullableTupleChange) IsSet() bool {
	return v.isSet
}

func (v *NullableTupleChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTupleChange(val *TupleChange) *NullableTupleChange {
	return &NullableTupleChange{value: val, isSet: true}
}

func (v NullableTupleChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTupleChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
