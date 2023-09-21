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

// LanProperties struct for LanProperties
type LanProperties struct {
<<<<<<< HEAD
	// The name of the  resource.
	Name *string `json:"name,omitempty"`
	// IP failover configurations for lan
	IpFailover *[]IPFailover `json:"ipFailover,omitempty"`
=======
	// IP failover configurations for lan
	IpFailover *[]IPFailover `json:"ipFailover,omitempty"`
	// [The IPv6 feature is in beta phase and not ready for production usage.] For a GET request, this value is either 'null' or contains the LAN's /64 IPv6 CIDR block if this LAN is IPv6 enabled. For POST/PUT/PATCH requests, 'AUTO' will result in enabling this LAN for IPv6 and automatically assign a /64 IPv6 CIDR block to this LAN and /80 IPv6 CIDR blocks to the NICs and one /128 IPv6 address to each connected NIC. If you choose the IPv6 CIDR block for the LAN on your own, then you must provide a /64 block, which is inside the IPv6 CIDR block of the virtual datacenter and unique inside all LANs from this virtual datacenter. If you enable IPv6 on a LAN with NICs, those NICs will get a /80 IPv6 CIDR block and one IPv6 address assigned to each automatically, unless you specify them explicitly on the LAN and on the NICs. A virtual data center is limited to a maximum of 256 IPv6-enabled LANs.
	// to set this field to `nil` in order to be marshalled, the explicit nil address `Nilstring` can be used, or the setter `SetIpv6CidrBlockNil`
	Ipv6CidrBlock *string `json:"ipv6CidrBlock,omitempty"`
	// The name of the  resource.
	Name *string `json:"name,omitempty"`
>>>>>>> main
	// The unique identifier of the private Cross-Connect the LAN is connected to, if any.
	Pcc *string `json:"pcc,omitempty"`
	// This LAN faces the public Internet.
	Public *bool `json:"public,omitempty"`
}

// NewLanProperties instantiates a new LanProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanProperties() *LanProperties {
	this := LanProperties{}

	return &this
}

// NewLanPropertiesWithDefaults instantiates a new LanProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanPropertiesWithDefaults() *LanProperties {
	this := LanProperties{}
	return &this
}

<<<<<<< HEAD
// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LanProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *LanProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *LanProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetIpFailover returns the IpFailover field value
// If the value is explicit nil, the zero value for []IPFailover will be returned
=======
// GetIpFailover returns the IpFailover field value
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *LanProperties) GetIpFailover() *[]IPFailover {
	if o == nil {
		return nil
	}

	return o.IpFailover

}

// GetIpFailoverOk returns a tuple with the IpFailover field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetIpFailoverOk() (*[]IPFailover, bool) {
	if o == nil {
		return nil, false
	}

	return o.IpFailover, true
}

// SetIpFailover sets field value
func (o *LanProperties) SetIpFailover(v []IPFailover) {

	o.IpFailover = &v

}

// HasIpFailover returns a boolean if a field has been set.
func (o *LanProperties) HasIpFailover() bool {
	if o != nil && o.IpFailover != nil {
		return true
	}

	return false
}

<<<<<<< HEAD
// GetPcc returns the Pcc field value
// If the value is explicit nil, the zero value for string will be returned
=======
// GetIpv6CidrBlock returns the Ipv6CidrBlock field value
// If the value is explicit nil, nil is returned
func (o *LanProperties) GetIpv6CidrBlock() *string {
	if o == nil {
		return nil
	}

	return o.Ipv6CidrBlock

}

// GetIpv6CidrBlockOk returns a tuple with the Ipv6CidrBlock field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetIpv6CidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Ipv6CidrBlock, true
}

// SetIpv6CidrBlock sets field value
func (o *LanProperties) SetIpv6CidrBlock(v string) {

	o.Ipv6CidrBlock = &v

}

// sets Ipv6CidrBlock to the explicit address that will be encoded as nil when marshaled
func (o *LanProperties) SetIpv6CidrBlockNil() {
	o.Ipv6CidrBlock = &Nilstring
}

// HasIpv6CidrBlock returns a boolean if a field has been set.
func (o *LanProperties) HasIpv6CidrBlock() bool {
	if o != nil && o.Ipv6CidrBlock != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, nil is returned
func (o *LanProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *LanProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *LanProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetPcc returns the Pcc field value
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *LanProperties) GetPcc() *string {
	if o == nil {
		return nil
	}

	return o.Pcc

}

// GetPccOk returns a tuple with the Pcc field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetPccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Pcc, true
}

// SetPcc sets field value
func (o *LanProperties) SetPcc(v string) {

	o.Pcc = &v

}

// HasPcc returns a boolean if a field has been set.
func (o *LanProperties) HasPcc() bool {
	if o != nil && o.Pcc != nil {
		return true
	}

	return false
}

// GetPublic returns the Public field value
<<<<<<< HEAD
// If the value is explicit nil, the zero value for bool will be returned
=======
// If the value is explicit nil, nil is returned
>>>>>>> main
func (o *LanProperties) GetPublic() *bool {
	if o == nil {
		return nil
	}

	return o.Public

}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.Public, true
}

// SetPublic sets field value
func (o *LanProperties) SetPublic(v bool) {

	o.Public = &v

}

// HasPublic returns a boolean if a field has been set.
func (o *LanProperties) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

func (o LanProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
<<<<<<< HEAD
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.IpFailover != nil {
		toSerialize["ipFailover"] = o.IpFailover
	}
	if o.Pcc != nil {
		toSerialize["pcc"] = o.Pcc
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
=======
	if o.IpFailover != nil {
		toSerialize["ipFailover"] = o.IpFailover
	}

	if o.Ipv6CidrBlock == &Nilstring {
		toSerialize["ipv6CidrBlock"] = nil
	} else if o.Ipv6CidrBlock != nil {
		toSerialize["ipv6CidrBlock"] = o.Ipv6CidrBlock
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Pcc != nil {
		toSerialize["pcc"] = o.Pcc
	}

	if o.Public != nil {
		toSerialize["public"] = o.Public
	}

>>>>>>> main
	return json.Marshal(toSerialize)
}

type NullableLanProperties struct {
	value *LanProperties
	isSet bool
}

func (v NullableLanProperties) Get() *LanProperties {
	return v.value
}

func (v *NullableLanProperties) Set(val *LanProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLanProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLanProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanProperties(val *LanProperties) *NullableLanProperties {
	return &NullableLanProperties{value: val, isSet: true}
}

func (v NullableLanProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
