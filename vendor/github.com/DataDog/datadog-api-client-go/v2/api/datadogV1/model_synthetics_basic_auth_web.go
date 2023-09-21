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

// SyntheticsBasicAuthWeb Object to handle basic authentication when performing the test.
type SyntheticsBasicAuthWeb struct {
	// Password to use for the basic authentication.
	Password string `json:"password"`
	// The type of basic authentication to use when performing the test.
	Type *SyntheticsBasicAuthWebType `json:"type,omitempty"`
	// Username to use for the basic authentication.
	Username string `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsBasicAuthWeb instantiates a new SyntheticsBasicAuthWeb object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsBasicAuthWeb(password string, username string) *SyntheticsBasicAuthWeb {
	this := SyntheticsBasicAuthWeb{}
	this.Password = password
	var typeVar SyntheticsBasicAuthWebType = SYNTHETICSBASICAUTHWEBTYPE_WEB
	this.Type = &typeVar
	this.Username = username
	return &this
}

// NewSyntheticsBasicAuthWebWithDefaults instantiates a new SyntheticsBasicAuthWeb object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsBasicAuthWebWithDefaults() *SyntheticsBasicAuthWeb {
	this := SyntheticsBasicAuthWeb{}
	var typeVar SyntheticsBasicAuthWebType = SYNTHETICSBASICAUTHWEBTYPE_WEB
	this.Type = &typeVar
	return &this
}

// GetPassword returns the Password field value.
func (o *SyntheticsBasicAuthWeb) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthWeb) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *SyntheticsBasicAuthWeb) SetPassword(v string) {
	o.Password = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsBasicAuthWeb) GetType() SyntheticsBasicAuthWebType {
	if o == nil || o.Type == nil {
		var ret SyntheticsBasicAuthWebType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthWeb) GetTypeOk() (*SyntheticsBasicAuthWebType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsBasicAuthWeb) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SyntheticsBasicAuthWebType and assigns it to the Type field.
func (o *SyntheticsBasicAuthWeb) SetType(v SyntheticsBasicAuthWebType) {
	o.Type = &v
}

// GetUsername returns the Username field value.
func (o *SyntheticsBasicAuthWeb) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthWeb) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *SyntheticsBasicAuthWeb) SetUsername(v string) {
	o.Username = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsBasicAuthWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["password"] = o.Password
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsBasicAuthWeb) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
=======
>>>>>>> main
	all := struct {
		Password *string                     `json:"password"`
		Type     *SyntheticsBasicAuthWebType `json:"type,omitempty"`
		Username *string                     `json:"username"`
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
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"password", "type", "username"})
	} else {
		return err
	}
<<<<<<< HEAD
	if v := all.Type; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Password = *all.Password
	o.Type = all.Type
	o.Username = *all.Username
=======

	hasInvalidField := false
	o.Password = *all.Password
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Username = *all.Username

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
