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

// Lan struct for Lan
type Lan struct {
<<<<<<< HEAD
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href       *string                    `json:"href,omitempty"`
	Metadata   *DatacenterElementMetadata `json:"metadata,omitempty"`
	Properties *LanProperties             `json:"properties"`
	Entities   *LanEntities               `json:"entities,omitempty"`
=======
	Entities *LanEntities `json:"entities,omitempty"`
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// The resource's unique identifier.
	Id         *string                    `json:"id,omitempty"`
	Metadata   *DatacenterElementMetadata `json:"metadata,omitempty"`
	Properties *LanProperties             `json:"properties"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
>>>>>>> main
}

// NewLan instantiates a new Lan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLan(properties LanProperties) *Lan {
	this := Lan{}

	this.Properties = &properties

	return &this
}

// NewLanWithDefaults instantiates a new Lan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanWithDefaults() *Lan {
	this := Lan{}
	return &this
}

<<<<<<< HEAD
// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Lan) GetId() *string {
=======
// GetEntities returns the Entities field value
// If the value is explicit nil, nil is returned
func (o *Lan) GetEntities() *LanEntities {
>>>>>>> main
	if o == nil {
		return nil
	}

<<<<<<< HEAD
	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lan) GetIdOk() (*string, bool) {
=======
	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lan) GetEntitiesOk() (*LanEntities, bool) {
>>>>>>> main
	if o == nil {
		return nil, false
	}

<<<<<<< HEAD
	return o.Id, true
}

// SetId sets field value
func (o *Lan) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Lan) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *Lan) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lan) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *Lan) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *Lan) HasType() bool {
	if o != nil && o.Type != nil {
=======
	return o.Entities, true
}

// SetEntities sets field value
func (o *Lan) SetEntities(v LanEntities) {

	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *Lan) HasEntities() bool {
	if o != nil && o.Entities != nil {
>>>>>>> main
		return true
	}

	return false
}

// GetHref returns the Href field value
<<<<<<< HEAD
// If the value is explicit nil, the zero value for string will be returned
=======
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *Lan) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lan) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *Lan) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *Lan) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for DatacenterElementMetadata will be returned
=======
// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *Lan) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lan) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *Lan) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Lan) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *Lan) GetMetadata() *DatacenterElementMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lan) GetMetadataOk() (*DatacenterElementMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Lan) SetMetadata(v DatacenterElementMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *Lan) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
<<<<<<< HEAD
// If the value is explicit nil, the zero value for LanProperties will be returned
=======
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *Lan) GetProperties() *LanProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lan) GetPropertiesOk() (*LanProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *Lan) SetProperties(v LanProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *Lan) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
// GetEntities returns the Entities field value
// If the value is explicit nil, the zero value for LanEntities will be returned
func (o *Lan) GetEntities() *LanEntities {
=======
// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *Lan) GetType() *Type {
>>>>>>> main
	if o == nil {
		return nil
	}

<<<<<<< HEAD
	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lan) GetEntitiesOk() (*LanEntities, bool) {
=======
	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lan) GetTypeOk() (*Type, bool) {
>>>>>>> main
	if o == nil {
		return nil, false
	}

<<<<<<< HEAD
	return o.Entities, true
}

// SetEntities sets field value
func (o *Lan) SetEntities(v LanEntities) {

	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *Lan) HasEntities() bool {
	if o != nil && o.Entities != nil {
=======
	return o.Type, true
}

// SetType sets field value
func (o *Lan) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *Lan) HasType() bool {
	if o != nil && o.Type != nil {
>>>>>>> main
		return true
	}

	return false
}

func (o Lan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
<<<<<<< HEAD
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
=======
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

>>>>>>> main
	return json.Marshal(toSerialize)
}

type NullableLan struct {
	value *Lan
	isSet bool
}

func (v NullableLan) Get() *Lan {
	return v.value
}

func (v *NullableLan) Set(val *Lan) {
	v.value = val
	v.isSet = true
}

func (v NullableLan) IsSet() bool {
	return v.isSet
}

func (v *NullableLan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLan(val *Lan) *NullableLan {
	return &NullableLan{value: val, isSet: true}
}

func (v NullableLan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
