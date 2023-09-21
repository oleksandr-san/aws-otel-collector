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
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsDateRemapper As Datadog receives logs, it timestamps them using the value(s) from any of these default attributes.
//
//   - `timestamp`
<<<<<<< HEAD
//   - `date`
//   - `_timestamp`
//   - `Timestamp`
//   - `eventTime`
//   - `published_date`
//
//   If your logs put their dates in an attribute not in this list,
//   use the log date Remapper Processor to define their date attribute as the official log timestamp.
//   The recognized date formats are ISO8601, UNIX (the milliseconds EPOCH format), and RFC3164.
//
//   **Note:** If your logs don’t contain any of the default attributes
//   and you haven’t defined your own date attribute, Datadog timestamps
//   the logs with the date it received them.
//
//   If multiple log date remapper processors can be applied to a given log,
//   only the first one (according to the pipelines order) is taken into account.
=======
//
//   - `date`
//
//   - `_timestamp`
//
//   - `Timestamp`
//
//   - `eventTime`
//
//   - `published_date`
//
//     If your logs put their dates in an attribute not in this list,
//     use the log date Remapper Processor to define their date attribute as the official log timestamp.
//     The recognized date formats are ISO8601, UNIX (the milliseconds EPOCH format), and RFC3164.
//
//     **Note:** If your logs don’t contain any of the default attributes
//     and you haven’t defined your own date attribute, Datadog timestamps
//     the logs with the date it received them.
//
//     If multiple log date remapper processors can be applied to a given log,
//     only the first one (according to the pipelines order) is taken into account.
>>>>>>> main
type LogsDateRemapper struct {
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
	// Array of source attributes.
	Sources []string `json:"sources"`
	// Type of logs date remapper.
	Type LogsDateRemapperType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewLogsDateRemapper instantiates a new LogsDateRemapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsDateRemapper(sources []string, typeVar LogsDateRemapperType) *LogsDateRemapper {
	this := LogsDateRemapper{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	this.Sources = sources
	this.Type = typeVar
	return &this
}

// NewLogsDateRemapperWithDefaults instantiates a new LogsDateRemapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsDateRemapperWithDefaults() *LogsDateRemapper {
	this := LogsDateRemapper{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var typeVar LogsDateRemapperType = LOGSDATEREMAPPERTYPE_DATE_REMAPPER
	this.Type = typeVar
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsDateRemapper) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsDateRemapper) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsDateRemapper) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsDateRemapper) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsDateRemapper) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsDateRemapper) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsDateRemapper) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsDateRemapper) SetName(v string) {
	o.Name = &v
}

// GetSources returns the Sources field value.
func (o *LogsDateRemapper) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *LogsDateRemapper) GetSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *LogsDateRemapper) SetSources(v []string) {
	o.Sources = v
}

// GetType returns the Type field value.
func (o *LogsDateRemapper) GetType() LogsDateRemapperType {
	if o == nil {
		var ret LogsDateRemapperType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsDateRemapper) GetTypeOk() (*LogsDateRemapperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsDateRemapper) SetType(v LogsDateRemapperType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsDateRemapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["sources"] = o.Sources
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsDateRemapper) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
=======
>>>>>>> main
	all := struct {
		IsEnabled *bool                 `json:"is_enabled,omitempty"`
		Name      *string               `json:"name,omitempty"`
		Sources   *[]string             `json:"sources"`
		Type      *LogsDateRemapperType `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
<<<<<<< HEAD
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
=======
		return json.Unmarshal(bytes, &o.UnparsedObject)
>>>>>>> main
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_enabled", "name", "sources", "type"})
	} else {
		return err
	}
<<<<<<< HEAD
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.Sources = *all.Sources
	o.Type = *all.Type
=======

	hasInvalidField := false
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.Sources = *all.Sources
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

>>>>>>> main
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

<<<<<<< HEAD
=======
	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

>>>>>>> main
	return nil
}
