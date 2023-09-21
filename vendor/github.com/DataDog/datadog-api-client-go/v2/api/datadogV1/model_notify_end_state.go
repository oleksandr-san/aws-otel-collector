// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// NotifyEndState A notification end state.
type NotifyEndState string

// List of NotifyEndState.
const (
	NOTIFYENDSTATE_ALERT   NotifyEndState = "alert"
	NOTIFYENDSTATE_NO_DATA NotifyEndState = "no data"
	NOTIFYENDSTATE_WARN    NotifyEndState = "warn"
)

var allowedNotifyEndStateEnumValues = []NotifyEndState{
	NOTIFYENDSTATE_ALERT,
	NOTIFYENDSTATE_NO_DATA,
	NOTIFYENDSTATE_WARN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotifyEndState) GetAllowedValues() []NotifyEndState {
	return allowedNotifyEndStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotifyEndState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotifyEndState(value)
	return nil
}

// NewNotifyEndStateFromValue returns a pointer to a valid NotifyEndState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotifyEndStateFromValue(v string) (*NotifyEndState, error) {
	ev := NotifyEndState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotifyEndState: valid values are %v", v, allowedNotifyEndStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotifyEndState) IsValid() bool {
	for _, existing := range allowedNotifyEndStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotifyEndState value.
func (v NotifyEndState) Ptr() *NotifyEndState {
	return &v
}
<<<<<<< HEAD

// NullableNotifyEndState handles when a null is used for NotifyEndState.
type NullableNotifyEndState struct {
	value *NotifyEndState
	isSet bool
}

// Get returns the associated value.
func (v NullableNotifyEndState) Get() *NotifyEndState {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableNotifyEndState) Set(val *NotifyEndState) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableNotifyEndState) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableNotifyEndState) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNotifyEndState initializes the struct as if Set has been called.
func NewNullableNotifyEndState(val *NotifyEndState) *NullableNotifyEndState {
	return &NullableNotifyEndState{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableNotifyEndState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableNotifyEndState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
