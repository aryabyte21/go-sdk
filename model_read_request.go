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

// ReadRequest struct for ReadRequest
type ReadRequest struct {
	TupleKey          *ReadRequestTupleKey `json:"tuple_key,omitempty"yaml:"tuple_key,omitempty"`
	PageSize          *int32               `json:"page_size,omitempty"yaml:"page_size,omitempty"`
	ContinuationToken *string              `json:"continuation_token,omitempty"yaml:"continuation_token,omitempty"`
}

// NewReadRequest instantiates a new ReadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadRequest() *ReadRequest {
	this := ReadRequest{}
	return &this
}

// NewReadRequestWithDefaults instantiates a new ReadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadRequestWithDefaults() *ReadRequest {
	this := ReadRequest{}
	return &this
}

// GetTupleKey returns the TupleKey field value if set, zero value otherwise.
func (o *ReadRequest) GetTupleKey() ReadRequestTupleKey {
	if o == nil || o.TupleKey == nil {
		var ret ReadRequestTupleKey
		return ret
	}
	return *o.TupleKey
}

// GetTupleKeyOk returns a tuple with the TupleKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRequest) GetTupleKeyOk() (*ReadRequestTupleKey, bool) {
	if o == nil || o.TupleKey == nil {
		return nil, false
	}
	return o.TupleKey, true
}

// HasTupleKey returns a boolean if a field has been set.
func (o *ReadRequest) HasTupleKey() bool {
	if o != nil && o.TupleKey != nil {
		return true
	}

	return false
}

// SetTupleKey gets a reference to the given ReadRequestTupleKey and assigns it to the TupleKey field.
func (o *ReadRequest) SetTupleKey(v ReadRequestTupleKey) {
	o.TupleKey = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ReadRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ReadRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ReadRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *ReadRequest) GetContinuationToken() string {
	if o == nil || o.ContinuationToken == nil {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRequest) GetContinuationTokenOk() (*string, bool) {
	if o == nil || o.ContinuationToken == nil {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *ReadRequest) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken != nil {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *ReadRequest) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

func (o ReadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TupleKey != nil {
		toSerialize["tuple_key"] = o.TupleKey
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.ContinuationToken != nil {
		toSerialize["continuation_token"] = o.ContinuationToken
	}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableReadRequest struct {
	value *ReadRequest
	isSet bool
}

func (v NullableReadRequest) Get() *ReadRequest {
	return v.value
}

func (v *NullableReadRequest) Set(val *ReadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadRequest(val *ReadRequest) *NullableReadRequest {
	return &NullableReadRequest{value: val, isSet: true}
}

func (v NullableReadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
