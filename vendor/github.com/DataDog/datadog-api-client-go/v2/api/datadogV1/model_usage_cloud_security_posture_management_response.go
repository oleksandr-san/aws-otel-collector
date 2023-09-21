// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
<<<<<<< HEAD
	"encoding/json"
=======
	"github.com/goccy/go-json"
>>>>>>> main

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageCloudSecurityPostureManagementResponse The response containing the Cloud Security Posture Management usage for each hour for a given organization.
type UsageCloudSecurityPostureManagementResponse struct {
	// Get hourly usage for Cloud Security Posture Management.
	Usage []UsageCloudSecurityPostureManagementHour `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageCloudSecurityPostureManagementResponse instantiates a new UsageCloudSecurityPostureManagementResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageCloudSecurityPostureManagementResponse() *UsageCloudSecurityPostureManagementResponse {
	this := UsageCloudSecurityPostureManagementResponse{}
	return &this
}

// NewUsageCloudSecurityPostureManagementResponseWithDefaults instantiates a new UsageCloudSecurityPostureManagementResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageCloudSecurityPostureManagementResponseWithDefaults() *UsageCloudSecurityPostureManagementResponse {
	this := UsageCloudSecurityPostureManagementResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageCloudSecurityPostureManagementResponse) GetUsage() []UsageCloudSecurityPostureManagementHour {
	if o == nil || o.Usage == nil {
		var ret []UsageCloudSecurityPostureManagementHour
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCloudSecurityPostureManagementResponse) GetUsageOk() (*[]UsageCloudSecurityPostureManagementHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageCloudSecurityPostureManagementResponse) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given []UsageCloudSecurityPostureManagementHour and assigns it to the Usage field.
func (o *UsageCloudSecurityPostureManagementResponse) SetUsage(v []UsageCloudSecurityPostureManagementHour) {
	o.Usage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageCloudSecurityPostureManagementResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageCloudSecurityPostureManagementResponse) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
=======
>>>>>>> main
	all := struct {
		Usage []UsageCloudSecurityPostureManagementHour `json:"usage,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"usage"})
	} else {
		return err
	}
	o.Usage = all.Usage
<<<<<<< HEAD
=======

>>>>>>> main
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
