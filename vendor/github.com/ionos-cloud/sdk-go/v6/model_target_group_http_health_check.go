/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// TargetGroupHttpHealthCheck struct for TargetGroupHttpHealthCheck
type TargetGroupHttpHealthCheck struct {
<<<<<<< HEAD
	// The path (destination URL) for the HTTP health check request; the default is /.
	Path *string `json:"path,omitempty"`
	// The method for the HTTP health check.
	Method *string `json:"method,omitempty"`
	//
	MatchType *string `json:"matchType"`
	// The response returned by the request, depending on the match type.
	Response *string `json:"response"`
	//
	Regex *bool `json:"regex,omitempty"`
	//
	Negate *bool `json:"negate,omitempty"`
=======
	// Specify the target's response type to match ALB's request.
	MatchType *string `json:"matchType"`
	// The method used for the health check request.
	Method *string `json:"method,omitempty"`
	// Specifies whether to negate an individual entry; the default value is 'FALSE'.
	Negate *bool `json:"negate,omitempty"`
	// The destination URL for HTTP the health check; the default is '/'.
	Path *string `json:"path,omitempty"`
	// Specifies whether to use a regular expression to parse the response body; the default value is 'FALSE'.  By using regular expressions, you can flexibly customize the expected response from a healthy server.
	Regex *bool `json:"regex,omitempty"`
	// The response returned by the request. It can be a status code or a response body depending on the definition of 'matchType'.
	Response *string `json:"response"`
>>>>>>> main
}

// NewTargetGroupHttpHealthCheck instantiates a new TargetGroupHttpHealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetGroupHttpHealthCheck(matchType string, response string) *TargetGroupHttpHealthCheck {
	this := TargetGroupHttpHealthCheck{}

	this.MatchType = &matchType
	this.Response = &response

	return &this
}

// NewTargetGroupHttpHealthCheckWithDefaults instantiates a new TargetGroupHttpHealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetGroupHttpHealthCheckWithDefaults() *TargetGroupHttpHealthCheck {
	this := TargetGroupHttpHealthCheck{}
	return &this
}

<<<<<<< HEAD
// GetPath returns the Path field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TargetGroupHttpHealthCheck) GetPath() *string {
	if o == nil {
		return nil
	}

	return o.Path

}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Path, true
}

// SetPath sets field value
func (o *TargetGroupHttpHealthCheck) SetPath(v string) {

	o.Path = &v

}

// HasPath returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// GetMethod returns the Method field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TargetGroupHttpHealthCheck) GetMethod() *string {
	if o == nil {
		return nil
	}

	return o.Method

}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Method, true
}

// SetMethod sets field value
func (o *TargetGroupHttpHealthCheck) SetMethod(v string) {

	o.Method = &v

}

// HasMethod returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// GetMatchType returns the MatchType field value
// If the value is explicit nil, the zero value for string will be returned
=======
// GetMatchType returns the MatchType field value
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *TargetGroupHttpHealthCheck) GetMatchType() *string {
	if o == nil {
		return nil
	}

	return o.MatchType

}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetMatchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.MatchType, true
}

// SetMatchType sets field value
func (o *TargetGroupHttpHealthCheck) SetMatchType(v string) {

	o.MatchType = &v

}

// HasMatchType returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasMatchType() bool {
	if o != nil && o.MatchType != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
// GetResponse returns the Response field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TargetGroupHttpHealthCheck) GetResponse() *string {
=======
// GetMethod returns the Method field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupHttpHealthCheck) GetMethod() *string {
>>>>>>> main
	if o == nil {
		return nil
	}

<<<<<<< HEAD
	return o.Response

}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetResponseOk() (*string, bool) {
=======
	return o.Method

}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetMethodOk() (*string, bool) {
>>>>>>> main
	if o == nil {
		return nil, false
	}

<<<<<<< HEAD
	return o.Response, true
}

// SetResponse sets field value
func (o *TargetGroupHttpHealthCheck) SetResponse(v string) {

	o.Response = &v

}

// HasResponse returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// GetRegex returns the Regex field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *TargetGroupHttpHealthCheck) GetRegex() *bool {
	if o == nil {
		return nil
	}

	return o.Regex

}

// GetRegexOk returns a tuple with the Regex field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetRegexOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.Regex, true
}

// SetRegex sets field value
func (o *TargetGroupHttpHealthCheck) SetRegex(v bool) {

	o.Regex = &v

}

// HasRegex returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasRegex() bool {
	if o != nil && o.Regex != nil {
=======
	return o.Method, true
}

// SetMethod sets field value
func (o *TargetGroupHttpHealthCheck) SetMethod(v string) {

	o.Method = &v

}

// HasMethod returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasMethod() bool {
	if o != nil && o.Method != nil {
>>>>>>> main
		return true
	}

	return false
}

// GetNegate returns the Negate field value
<<<<<<< HEAD
// If the value is explicit nil, the zero value for bool will be returned
=======
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *TargetGroupHttpHealthCheck) GetNegate() *bool {
	if o == nil {
		return nil
	}

	return o.Negate

}

// GetNegateOk returns a tuple with the Negate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetNegateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.Negate, true
}

// SetNegate sets field value
func (o *TargetGroupHttpHealthCheck) SetNegate(v bool) {

	o.Negate = &v

}

// HasNegate returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasNegate() bool {
	if o != nil && o.Negate != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
func (o TargetGroupHttpHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.MatchType != nil {
		toSerialize["matchType"] = o.MatchType
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.Regex != nil {
		toSerialize["regex"] = o.Regex
	}
	if o.Negate != nil {
		toSerialize["negate"] = o.Negate
	}
=======
// GetPath returns the Path field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupHttpHealthCheck) GetPath() *string {
	if o == nil {
		return nil
	}

	return o.Path

}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Path, true
}

// SetPath sets field value
func (o *TargetGroupHttpHealthCheck) SetPath(v string) {

	o.Path = &v

}

// HasPath returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// GetRegex returns the Regex field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupHttpHealthCheck) GetRegex() *bool {
	if o == nil {
		return nil
	}

	return o.Regex

}

// GetRegexOk returns a tuple with the Regex field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetRegexOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.Regex, true
}

// SetRegex sets field value
func (o *TargetGroupHttpHealthCheck) SetRegex(v bool) {

	o.Regex = &v

}

// HasRegex returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasRegex() bool {
	if o != nil && o.Regex != nil {
		return true
	}

	return false
}

// GetResponse returns the Response field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupHttpHealthCheck) GetResponse() *string {
	if o == nil {
		return nil
	}

	return o.Response

}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupHttpHealthCheck) GetResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Response, true
}

// SetResponse sets field value
func (o *TargetGroupHttpHealthCheck) SetResponse(v string) {

	o.Response = &v

}

// HasResponse returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

func (o TargetGroupHttpHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MatchType != nil {
		toSerialize["matchType"] = o.MatchType
	}

	if o.Method != nil {
		toSerialize["method"] = o.Method
	}

	if o.Negate != nil {
		toSerialize["negate"] = o.Negate
	}

	if o.Path != nil {
		toSerialize["path"] = o.Path
	}

	if o.Regex != nil {
		toSerialize["regex"] = o.Regex
	}

	if o.Response != nil {
		toSerialize["response"] = o.Response
	}

>>>>>>> main
	return json.Marshal(toSerialize)
}

type NullableTargetGroupHttpHealthCheck struct {
	value *TargetGroupHttpHealthCheck
	isSet bool
}

func (v NullableTargetGroupHttpHealthCheck) Get() *TargetGroupHttpHealthCheck {
	return v.value
}

func (v *NullableTargetGroupHttpHealthCheck) Set(val *TargetGroupHttpHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetGroupHttpHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetGroupHttpHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetGroupHttpHealthCheck(val *TargetGroupHttpHealthCheck) *NullableTargetGroupHttpHealthCheck {
	return &NullableTargetGroupHttpHealthCheck{value: val, isSet: true}
}

func (v NullableTargetGroupHttpHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetGroupHttpHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
