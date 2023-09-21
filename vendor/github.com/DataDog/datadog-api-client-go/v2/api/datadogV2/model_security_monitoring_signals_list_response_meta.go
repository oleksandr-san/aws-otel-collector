// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
=======
	"github.com/goccy/go-json"
>>>>>>> main

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalsListResponseMeta Meta attributes.
type SecurityMonitoringSignalsListResponseMeta struct {
	// Paging attributes.
	Page *SecurityMonitoringSignalsListResponseMetaPage `json:"page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringSignalsListResponseMeta instantiates a new SecurityMonitoringSignalsListResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalsListResponseMeta() *SecurityMonitoringSignalsListResponseMeta {
	this := SecurityMonitoringSignalsListResponseMeta{}
	return &this
}

// NewSecurityMonitoringSignalsListResponseMetaWithDefaults instantiates a new SecurityMonitoringSignalsListResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalsListResponseMetaWithDefaults() *SecurityMonitoringSignalsListResponseMeta {
	this := SecurityMonitoringSignalsListResponseMeta{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalsListResponseMeta) GetPage() SecurityMonitoringSignalsListResponseMetaPage {
	if o == nil || o.Page == nil {
		var ret SecurityMonitoringSignalsListResponseMetaPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsListResponseMeta) GetPageOk() (*SecurityMonitoringSignalsListResponseMetaPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalsListResponseMeta) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given SecurityMonitoringSignalsListResponseMetaPage and assigns it to the Page field.
func (o *SecurityMonitoringSignalsListResponseMeta) SetPage(v SecurityMonitoringSignalsListResponseMetaPage) {
	o.Page = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalsListResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalsListResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
=======
>>>>>>> main
	all := struct {
		Page *SecurityMonitoringSignalsListResponseMetaPage `json:"page,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"page"})
	} else {
		return err
	}
<<<<<<< HEAD
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Page = all.Page
=======

	hasInvalidField := false
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page

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
