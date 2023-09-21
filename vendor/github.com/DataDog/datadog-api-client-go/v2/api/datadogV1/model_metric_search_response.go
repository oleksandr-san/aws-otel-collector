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

// MetricSearchResponse Object containing the list of metrics matching the search query.
type MetricSearchResponse struct {
	// Search result.
	Results *MetricSearchResponseResults `json:"results,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMetricSearchResponse instantiates a new MetricSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricSearchResponse() *MetricSearchResponse {
	this := MetricSearchResponse{}
	return &this
}

// NewMetricSearchResponseWithDefaults instantiates a new MetricSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricSearchResponseWithDefaults() *MetricSearchResponse {
	this := MetricSearchResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MetricSearchResponse) GetResults() MetricSearchResponseResults {
	if o == nil || o.Results == nil {
		var ret MetricSearchResponseResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSearchResponse) GetResultsOk() (*MetricSearchResponseResults, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MetricSearchResponse) HasResults() bool {
	return o != nil && o.Results != nil
}

// SetResults gets a reference to the given MetricSearchResponseResults and assigns it to the Results field.
func (o *MetricSearchResponse) SetResults(v MetricSearchResponseResults) {
	o.Results = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricSearchResponse) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
=======
>>>>>>> main
	all := struct {
		Results *MetricSearchResponseResults `json:"results,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"results"})
	} else {
		return err
	}
<<<<<<< HEAD
	if all.Results != nil && all.Results.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Results = all.Results
=======

	hasInvalidField := false
	if all.Results != nil && all.Results.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Results = all.Results

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
