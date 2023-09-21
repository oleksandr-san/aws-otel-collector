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

// BackupUnits struct for BackupUnits
type BackupUnits struct {
<<<<<<< HEAD
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *string `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// Array of items in the collection.
	Items *[]BackupUnit `json:"items,omitempty"`
=======
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// Array of items in the collection.
	Items *[]BackupUnit `json:"items,omitempty"`
	// The type of object that has been created.
	Type *string `json:"type,omitempty"`
>>>>>>> main
}

// NewBackupUnits instantiates a new BackupUnits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupUnits() *BackupUnits {
	this := BackupUnits{}

	return &this
}

// NewBackupUnitsWithDefaults instantiates a new BackupUnits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupUnitsWithDefaults() *BackupUnits {
	this := BackupUnits{}
	return &this
}

<<<<<<< HEAD
// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupUnits) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnits) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *BackupUnits) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *BackupUnits) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupUnits) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnits) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *BackupUnits) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *BackupUnits) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
=======
// GetHref returns the Href field value
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *BackupUnits) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnits) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *BackupUnits) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *BackupUnits) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []BackupUnit will be returned
=======
// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *BackupUnits) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnits) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *BackupUnits) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *BackupUnits) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *BackupUnits) GetItems() *[]BackupUnit {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnits) GetItemsOk() (*[]BackupUnit, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *BackupUnits) SetItems(v []BackupUnit) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *BackupUnits) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
func (o BackupUnits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
=======
// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *BackupUnits) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnits) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *BackupUnits) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *BackupUnits) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o BackupUnits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

>>>>>>> main
	return json.Marshal(toSerialize)
}

type NullableBackupUnits struct {
	value *BackupUnits
	isSet bool
}

func (v NullableBackupUnits) Get() *BackupUnits {
	return v.value
}

func (v *NullableBackupUnits) Set(val *BackupUnits) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupUnits) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupUnits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupUnits(val *BackupUnits) *NullableBackupUnits {
	return &NullableBackupUnits{value: val, isSet: true}
}

func (v NullableBackupUnits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupUnits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
