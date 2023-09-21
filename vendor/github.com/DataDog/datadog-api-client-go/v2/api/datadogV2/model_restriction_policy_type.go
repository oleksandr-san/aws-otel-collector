// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// RestrictionPolicyType Restriction policy type.
type RestrictionPolicyType string

// List of RestrictionPolicyType.
const (
	RESTRICTIONPOLICYTYPE_RESTRICTION_POLICY RestrictionPolicyType = "restriction_policy"
)

var allowedRestrictionPolicyTypeEnumValues = []RestrictionPolicyType{
	RESTRICTIONPOLICYTYPE_RESTRICTION_POLICY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RestrictionPolicyType) GetAllowedValues() []RestrictionPolicyType {
	return allowedRestrictionPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RestrictionPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RestrictionPolicyType(value)
	return nil
}

// NewRestrictionPolicyTypeFromValue returns a pointer to a valid RestrictionPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRestrictionPolicyTypeFromValue(v string) (*RestrictionPolicyType, error) {
	ev := RestrictionPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RestrictionPolicyType: valid values are %v", v, allowedRestrictionPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RestrictionPolicyType) IsValid() bool {
	for _, existing := range allowedRestrictionPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RestrictionPolicyType value.
func (v RestrictionPolicyType) Ptr() *RestrictionPolicyType {
	return &v
}
<<<<<<< HEAD

// NullableRestrictionPolicyType handles when a null is used for RestrictionPolicyType.
type NullableRestrictionPolicyType struct {
	value *RestrictionPolicyType
	isSet bool
}

// Get returns the associated value.
func (v NullableRestrictionPolicyType) Get() *RestrictionPolicyType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableRestrictionPolicyType) Set(val *RestrictionPolicyType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableRestrictionPolicyType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableRestrictionPolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRestrictionPolicyType initializes the struct as if Set has been called.
func NewNullableRestrictionPolicyType(val *RestrictionPolicyType) *NullableRestrictionPolicyType {
	return &NullableRestrictionPolicyType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableRestrictionPolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableRestrictionPolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
