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

// RUMApplicationType RUM application response type.
type RUMApplicationType string

// List of RUMApplicationType.
const (
	RUMAPPLICATIONTYPE_RUM_APPLICATION RUMApplicationType = "rum_application"
)

var allowedRUMApplicationTypeEnumValues = []RUMApplicationType{
	RUMAPPLICATIONTYPE_RUM_APPLICATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMApplicationType) GetAllowedValues() []RUMApplicationType {
	return allowedRUMApplicationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMApplicationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMApplicationType(value)
	return nil
}

// NewRUMApplicationTypeFromValue returns a pointer to a valid RUMApplicationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMApplicationTypeFromValue(v string) (*RUMApplicationType, error) {
	ev := RUMApplicationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMApplicationType: valid values are %v", v, allowedRUMApplicationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMApplicationType) IsValid() bool {
	for _, existing := range allowedRUMApplicationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMApplicationType value.
func (v RUMApplicationType) Ptr() *RUMApplicationType {
	return &v
}
<<<<<<< HEAD

// NullableRUMApplicationType handles when a null is used for RUMApplicationType.
type NullableRUMApplicationType struct {
	value *RUMApplicationType
	isSet bool
}

// Get returns the associated value.
func (v NullableRUMApplicationType) Get() *RUMApplicationType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableRUMApplicationType) Set(val *RUMApplicationType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableRUMApplicationType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableRUMApplicationType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRUMApplicationType initializes the struct as if Set has been called.
func NewNullableRUMApplicationType(val *RUMApplicationType) *NullableRUMApplicationType {
	return &NullableRUMApplicationType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableRUMApplicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableRUMApplicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
