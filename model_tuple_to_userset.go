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
)

// TupleToUserset struct for TupleToUserset
type TupleToUserset struct {
	Tupleset        ObjectRelation `json:"tupleset"yaml:"tupleset"`
	ComputedUserset ObjectRelation `json:"computedUserset"yaml:"computedUserset"`
}

// NewTupleToUserset instantiates a new TupleToUserset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTupleToUserset(tupleset ObjectRelation, computedUserset ObjectRelation) *TupleToUserset {
	this := TupleToUserset{}
	this.Tupleset = tupleset
	this.ComputedUserset = computedUserset
	return &this
}

// NewTupleToUsersetWithDefaults instantiates a new TupleToUserset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleToUsersetWithDefaults() *TupleToUserset {
	this := TupleToUserset{}
	return &this
}

// GetTupleset returns the Tupleset field value
func (o *TupleToUserset) GetTupleset() ObjectRelation {
	if o == nil {
		var ret ObjectRelation
		return ret
	}

	return o.Tupleset
}

// GetTuplesetOk returns a tuple with the Tupleset field value
// and a boolean to check if the value has been set.
func (o *TupleToUserset) GetTuplesetOk() (*ObjectRelation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tupleset, true
}

// SetTupleset sets field value
func (o *TupleToUserset) SetTupleset(v ObjectRelation) {
	o.Tupleset = v
}

// GetComputedUserset returns the ComputedUserset field value
func (o *TupleToUserset) GetComputedUserset() ObjectRelation {
	if o == nil {
		var ret ObjectRelation
		return ret
	}

	return o.ComputedUserset
}

// GetComputedUsersetOk returns a tuple with the ComputedUserset field value
// and a boolean to check if the value has been set.
func (o *TupleToUserset) GetComputedUsersetOk() (*ObjectRelation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputedUserset, true
}

// SetComputedUserset sets field value
func (o *TupleToUserset) SetComputedUserset(v ObjectRelation) {
	o.ComputedUserset = v
}

func (o TupleToUserset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tupleset"] = o.Tupleset
	toSerialize["computedUserset"] = o.ComputedUserset
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableTupleToUserset struct {
	value *TupleToUserset
	isSet bool
}

func (v NullableTupleToUserset) Get() *TupleToUserset {
	return v.value
}

func (v *NullableTupleToUserset) Set(val *TupleToUserset) {
	v.value = val
	v.isSet = true
}

func (v NullableTupleToUserset) IsSet() bool {
	return v.isSet
}

func (v *NullableTupleToUserset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTupleToUserset(val *TupleToUserset) *NullableTupleToUserset {
	return &NullableTupleToUserset{value: val, isSet: true}
}

func (v NullableTupleToUserset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTupleToUserset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
