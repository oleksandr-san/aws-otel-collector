// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsGetAPITestLatestResultsResponse Object with the latest Synthetic API test run.
type SyntheticsGetAPITestLatestResultsResponse struct {
	// Timestamp of the latest API test run.
	LastTimestampFetched *int64 `json:"last_timestamp_fetched,omitempty"`
	// Result of the latest API test run.
	Results []SyntheticsAPITestResultShort `json:"results,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsGetAPITestLatestResultsResponse instantiates a new SyntheticsGetAPITestLatestResultsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsGetAPITestLatestResultsResponse() *SyntheticsGetAPITestLatestResultsResponse {
	this := SyntheticsGetAPITestLatestResultsResponse{}
	return &this
}

// NewSyntheticsGetAPITestLatestResultsResponseWithDefaults instantiates a new SyntheticsGetAPITestLatestResultsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsGetAPITestLatestResultsResponseWithDefaults() *SyntheticsGetAPITestLatestResultsResponse {
	this := SyntheticsGetAPITestLatestResultsResponse{}
	return &this
}

// GetLastTimestampFetched returns the LastTimestampFetched field value if set, zero value otherwise.
func (o *SyntheticsGetAPITestLatestResultsResponse) GetLastTimestampFetched() int64 {
	if o == nil || o.LastTimestampFetched == nil {
		var ret int64
		return ret
	}
	return *o.LastTimestampFetched
}

// GetLastTimestampFetchedOk returns a tuple with the LastTimestampFetched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGetAPITestLatestResultsResponse) GetLastTimestampFetchedOk() (*int64, bool) {
	if o == nil || o.LastTimestampFetched == nil {
		return nil, false
	}
	return o.LastTimestampFetched, true
}

// HasLastTimestampFetched returns a boolean if a field has been set.
func (o *SyntheticsGetAPITestLatestResultsResponse) HasLastTimestampFetched() bool {
	return o != nil && o.LastTimestampFetched != nil
}

// SetLastTimestampFetched gets a reference to the given int64 and assigns it to the LastTimestampFetched field.
func (o *SyntheticsGetAPITestLatestResultsResponse) SetLastTimestampFetched(v int64) {
	o.LastTimestampFetched = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SyntheticsGetAPITestLatestResultsResponse) GetResults() []SyntheticsAPITestResultShort {
	if o == nil || o.Results == nil {
		var ret []SyntheticsAPITestResultShort
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGetAPITestLatestResultsResponse) GetResultsOk() (*[]SyntheticsAPITestResultShort, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SyntheticsGetAPITestLatestResultsResponse) HasResults() bool {
	return o != nil && o.Results != nil
}

// SetResults gets a reference to the given []SyntheticsAPITestResultShort and assigns it to the Results field.
func (o *SyntheticsGetAPITestLatestResultsResponse) SetResults(v []SyntheticsAPITestResultShort) {
	o.Results = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsGetAPITestLatestResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LastTimestampFetched != nil {
		toSerialize["last_timestamp_fetched"] = o.LastTimestampFetched
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsGetAPITestLatestResultsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LastTimestampFetched *int64                         `json:"last_timestamp_fetched,omitempty"`
		Results              []SyntheticsAPITestResultShort `json:"results,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"last_timestamp_fetched", "results"})
	} else {
		return err
	}
	o.LastTimestampFetched = all.LastTimestampFetched
	o.Results = all.Results

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
