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

// NatGatewayRules struct for NatGatewayRules
type NatGatewayRules struct {
<<<<<<< HEAD
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// Array of items in the collection.
	Items *[]NatGatewayRule `json:"items,omitempty"`
=======
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// Array of items in the collection.
	Items *[]NatGatewayRule `json:"items,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
>>>>>>> main
}

// NewNatGatewayRules instantiates a new NatGatewayRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNatGatewayRules() *NatGatewayRules {
	this := NatGatewayRules{}

	return &this
}

// NewNatGatewayRulesWithDefaults instantiates a new NatGatewayRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNatGatewayRulesWithDefaults() *NatGatewayRules {
	this := NatGatewayRules{}
	return &this
}

<<<<<<< HEAD
// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NatGatewayRules) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRules) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *NatGatewayRules) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *NatGatewayRules) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *NatGatewayRules) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRules) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *NatGatewayRules) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *NatGatewayRules) HasType() bool {
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
func (o *NatGatewayRules) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRules) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *NatGatewayRules) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *NatGatewayRules) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []NatGatewayRule will be returned
=======
// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayRules) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRules) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *NatGatewayRules) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *NatGatewayRules) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *NatGatewayRules) GetItems() *[]NatGatewayRule {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRules) GetItemsOk() (*[]NatGatewayRule, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *NatGatewayRules) SetItems(v []NatGatewayRule) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *NatGatewayRules) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
func (o NatGatewayRules) MarshalJSON() ([]byte, error) {
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
func (o *NatGatewayRules) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayRules) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *NatGatewayRules) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *NatGatewayRules) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o NatGatewayRules) MarshalJSON() ([]byte, error) {
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

type NullableNatGatewayRules struct {
	value *NatGatewayRules
	isSet bool
}

func (v NullableNatGatewayRules) Get() *NatGatewayRules {
	return v.value
}

func (v *NullableNatGatewayRules) Set(val *NatGatewayRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNatGatewayRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNatGatewayRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatGatewayRules(val *NatGatewayRules) *NullableNatGatewayRules {
	return &NullableNatGatewayRules{value: val, isSet: true}
}

func (v NullableNatGatewayRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatGatewayRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
