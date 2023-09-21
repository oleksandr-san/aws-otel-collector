// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

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

// IncidentUpdateData Incident data for an update request.
type IncidentUpdateData struct {
	// The incident's attributes for an update request.
	Attributes *IncidentUpdateAttributes `json:"attributes,omitempty"`
	// The incident's ID.
	Id string `json:"id"`
	// The incident's relationships for an update request.
	Relationships *IncidentUpdateRelationships `json:"relationships,omitempty"`
	// Incident resource type.
	Type IncidentType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIncidentUpdateData instantiates a new IncidentUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUpdateData(id string, typeVar IncidentType) *IncidentUpdateData {
	this := IncidentUpdateData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentUpdateDataWithDefaults instantiates a new IncidentUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUpdateDataWithDefaults() *IncidentUpdateData {
	this := IncidentUpdateData{}
	var typeVar IncidentType = INCIDENTTYPE_INCIDENTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentUpdateData) GetAttributes() IncidentUpdateAttributes {
	if o == nil || o.Attributes == nil {
		var ret IncidentUpdateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateData) GetAttributesOk() (*IncidentUpdateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentUpdateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given IncidentUpdateAttributes and assigns it to the Attributes field.
func (o *IncidentUpdateData) SetAttributes(v IncidentUpdateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *IncidentUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentUpdateData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentUpdateData) GetRelationships() IncidentUpdateRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentUpdateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateData) GetRelationshipsOk() (*IncidentUpdateRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentUpdateData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given IncidentUpdateRelationships and assigns it to the Relationships field.
func (o *IncidentUpdateData) SetRelationships(v IncidentUpdateRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *IncidentUpdateData) GetType() IncidentType {
	if o == nil {
		var ret IncidentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentUpdateData) GetTypeOk() (*IncidentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentUpdateData) SetType(v IncidentType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUpdateData) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
=======
>>>>>>> main
	all := struct {
		Attributes    *IncidentUpdateAttributes    `json:"attributes,omitempty"`
		Id            *string                      `json:"id"`
		Relationships *IncidentUpdateRelationships `json:"relationships,omitempty"`
		Type          *IncidentType                `json:"type"`
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
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
=======

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
>>>>>>> main
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
<<<<<<< HEAD
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Relationships = all.Relationships
	o.Type = *all.Type
=======
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
