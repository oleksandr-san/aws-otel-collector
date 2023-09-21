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

// NetworkLoadBalancerForwardingRulePut struct for NetworkLoadBalancerForwardingRulePut
type NetworkLoadBalancerForwardingRulePut struct {
<<<<<<< HEAD
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href       *string                                      `json:"href,omitempty"`
	Properties *NetworkLoadBalancerForwardingRuleProperties `json:"properties"`
=======
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// The resource's unique identifier.
	Id         *string                                      `json:"id,omitempty"`
	Properties *NetworkLoadBalancerForwardingRuleProperties `json:"properties"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
>>>>>>> main
}

// NewNetworkLoadBalancerForwardingRulePut instantiates a new NetworkLoadBalancerForwardingRulePut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLoadBalancerForwardingRulePut(properties NetworkLoadBalancerForwardingRuleProperties) *NetworkLoadBalancerForwardingRulePut {
	this := NetworkLoadBalancerForwardingRulePut{}

	this.Properties = &properties

	return &this
}

// NewNetworkLoadBalancerForwardingRulePutWithDefaults instantiates a new NetworkLoadBalancerForwardingRulePut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLoadBalancerForwardingRulePutWithDefaults() *NetworkLoadBalancerForwardingRulePut {
	this := NetworkLoadBalancerForwardingRulePut{}
	return &this
}

<<<<<<< HEAD
// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkLoadBalancerForwardingRulePut) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRulePut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *NetworkLoadBalancerForwardingRulePut) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRulePut) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *NetworkLoadBalancerForwardingRulePut) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRulePut) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *NetworkLoadBalancerForwardingRulePut) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRulePut) HasType() bool {
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
func (o *NetworkLoadBalancerForwardingRulePut) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRulePut) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *NetworkLoadBalancerForwardingRulePut) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRulePut) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for NetworkLoadBalancerForwardingRuleProperties will be returned
=======
// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *NetworkLoadBalancerForwardingRulePut) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRulePut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *NetworkLoadBalancerForwardingRulePut) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRulePut) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *NetworkLoadBalancerForwardingRulePut) GetProperties() *NetworkLoadBalancerForwardingRuleProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRulePut) GetPropertiesOk() (*NetworkLoadBalancerForwardingRuleProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *NetworkLoadBalancerForwardingRulePut) SetProperties(v NetworkLoadBalancerForwardingRuleProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRulePut) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
func (o NetworkLoadBalancerForwardingRulePut) MarshalJSON() ([]byte, error) {
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
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
=======
// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *NetworkLoadBalancerForwardingRulePut) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRulePut) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *NetworkLoadBalancerForwardingRulePut) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRulePut) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o NetworkLoadBalancerForwardingRulePut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
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

type NullableNetworkLoadBalancerForwardingRulePut struct {
	value *NetworkLoadBalancerForwardingRulePut
	isSet bool
}

func (v NullableNetworkLoadBalancerForwardingRulePut) Get() *NetworkLoadBalancerForwardingRulePut {
	return v.value
}

func (v *NullableNetworkLoadBalancerForwardingRulePut) Set(val *NetworkLoadBalancerForwardingRulePut) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancerForwardingRulePut) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancerForwardingRulePut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancerForwardingRulePut(val *NetworkLoadBalancerForwardingRulePut) *NullableNetworkLoadBalancerForwardingRulePut {
	return &NullableNetworkLoadBalancerForwardingRulePut{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancerForwardingRulePut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancerForwardingRulePut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
