// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSDSResponse Response containing the Sensitive Data Scanner usage for each hour for a given organization.
type UsageSDSResponse struct {
	// Get hourly usage for Sensitive Data Scanner.
	Usage []UsageSDSHour `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageSDSResponse instantiates a new UsageSDSResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSDSResponse() *UsageSDSResponse {
	this := UsageSDSResponse{}
	return &this
}

// NewUsageSDSResponseWithDefaults instantiates a new UsageSDSResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSDSResponseWithDefaults() *UsageSDSResponse {
	this := UsageSDSResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageSDSResponse) GetUsage() []UsageSDSHour {
	if o == nil || o.Usage == nil {
		var ret []UsageSDSHour
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSDSResponse) GetUsageOk() (*[]UsageSDSHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageSDSResponse) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given []UsageSDSHour and assigns it to the Usage field.
func (o *UsageSDSResponse) SetUsage(v []UsageSDSHour) {
	o.Usage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSDSResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageSDSResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Usage []UsageSDSHour `json:"usage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"usage"})
	} else {
		return err
	}
	o.Usage = all.Usage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
