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

// S3ObjectStorageSSO struct for S3ObjectStorageSSO
type S3ObjectStorageSSO struct {
	// The S3 object storage single sign on url
	SsoUrl *string `json:"ssoUrl,omitempty"`
}

// NewS3ObjectStorageSSO instantiates a new S3ObjectStorageSSO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3ObjectStorageSSO() *S3ObjectStorageSSO {
	this := S3ObjectStorageSSO{}

	return &this
}

// NewS3ObjectStorageSSOWithDefaults instantiates a new S3ObjectStorageSSO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ObjectStorageSSOWithDefaults() *S3ObjectStorageSSO {
	this := S3ObjectStorageSSO{}
	return &this
}

// GetSsoUrl returns the SsoUrl field value
<<<<<<< HEAD
// If the value is explicit nil, the zero value for string will be returned
=======
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *S3ObjectStorageSSO) GetSsoUrl() *string {
	if o == nil {
		return nil
	}

	return o.SsoUrl

}

// GetSsoUrlOk returns a tuple with the SsoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3ObjectStorageSSO) GetSsoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.SsoUrl, true
}

// SetSsoUrl sets field value
func (o *S3ObjectStorageSSO) SetSsoUrl(v string) {

	o.SsoUrl = &v

}

// HasSsoUrl returns a boolean if a field has been set.
func (o *S3ObjectStorageSSO) HasSsoUrl() bool {
	if o != nil && o.SsoUrl != nil {
		return true
	}

	return false
}

func (o S3ObjectStorageSSO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SsoUrl != nil {
		toSerialize["ssoUrl"] = o.SsoUrl
	}
<<<<<<< HEAD
=======

>>>>>>> main
	return json.Marshal(toSerialize)
}

type NullableS3ObjectStorageSSO struct {
	value *S3ObjectStorageSSO
	isSet bool
}

func (v NullableS3ObjectStorageSSO) Get() *S3ObjectStorageSSO {
	return v.value
}

func (v *NullableS3ObjectStorageSSO) Set(val *S3ObjectStorageSSO) {
	v.value = val
	v.isSet = true
}

func (v NullableS3ObjectStorageSSO) IsSet() bool {
	return v.isSet
}

func (v *NullableS3ObjectStorageSSO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3ObjectStorageSSO(val *S3ObjectStorageSSO) *NullableS3ObjectStorageSSO {
	return &NullableS3ObjectStorageSSO{value: val, isSet: true}
}

func (v NullableS3ObjectStorageSSO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3ObjectStorageSSO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
