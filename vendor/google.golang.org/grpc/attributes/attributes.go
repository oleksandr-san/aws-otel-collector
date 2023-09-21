/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package attributes defines a generic key/value store used in various gRPC
// components.
//
// # Experimental
//
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
// later release.
package attributes

import (
	"fmt"
	"strings"
)

// Attributes is an immutable struct for storing and retrieving generic
// key/value pairs.  Keys must be hashable, and users should define their own
// types for keys.  Values should not be modified after they are added to an
// Attributes or if they were received from one.  If values implement 'Equal(o
<<<<<<< HEAD
// interface{}) bool', it will be called by (*Attributes).Equal to determine
// whether two values with the same key should be considered equal.
type Attributes struct {
	m map[interface{}]interface{}
}

// New returns a new Attributes containing the key/value pair.
func New(key, value interface{}) *Attributes {
	return &Attributes{m: map[interface{}]interface{}{key: value}}
=======
// any) bool', it will be called by (*Attributes).Equal to determine whether
// two values with the same key should be considered equal.
type Attributes struct {
	m map[any]any
}

// New returns a new Attributes containing the key/value pair.
func New(key, value any) *Attributes {
	return &Attributes{m: map[any]any{key: value}}
>>>>>>> main
}

// WithValue returns a new Attributes containing the previous keys and values
// and the new key/value pair.  If the same key appears multiple times, the
// last value overwrites all previous values for that key.  To remove an
// existing key, use a nil value.  value should not be modified later.
<<<<<<< HEAD
func (a *Attributes) WithValue(key, value interface{}) *Attributes {
	if a == nil {
		return New(key, value)
	}
	n := &Attributes{m: make(map[interface{}]interface{}, len(a.m)+1)}
=======
func (a *Attributes) WithValue(key, value any) *Attributes {
	if a == nil {
		return New(key, value)
	}
	n := &Attributes{m: make(map[any]any, len(a.m)+1)}
>>>>>>> main
	for k, v := range a.m {
		n.m[k] = v
	}
	n.m[key] = value
	return n
}

// Value returns the value associated with these attributes for key, or nil if
// no value is associated with key.  The returned value should not be modified.
<<<<<<< HEAD
func (a *Attributes) Value(key interface{}) interface{} {
=======
func (a *Attributes) Value(key any) any {
>>>>>>> main
	if a == nil {
		return nil
	}
	return a.m[key]
}

<<<<<<< HEAD
// Equal returns whether a and o are equivalent.  If 'Equal(o interface{})
// bool' is implemented for a value in the attributes, it is called to
// determine if the value matches the one stored in the other attributes.  If
// Equal is not implemented, standard equality is used to determine if the two
// values are equal. Note that some types (e.g. maps) aren't comparable by
// default, so they must be wrapped in a struct, or in an alias type, with Equal
// defined.
=======
// Equal returns whether a and o are equivalent.  If 'Equal(o any) bool' is
// implemented for a value in the attributes, it is called to determine if the
// value matches the one stored in the other attributes.  If Equal is not
// implemented, standard equality is used to determine if the two values are
// equal. Note that some types (e.g. maps) aren't comparable by default, so
// they must be wrapped in a struct, or in an alias type, with Equal defined.
>>>>>>> main
func (a *Attributes) Equal(o *Attributes) bool {
	if a == nil && o == nil {
		return true
	}
	if a == nil || o == nil {
		return false
	}
	if len(a.m) != len(o.m) {
		return false
	}
	for k, v := range a.m {
		ov, ok := o.m[k]
		if !ok {
			// o missing element of a
			return false
		}
<<<<<<< HEAD
		if eq, ok := v.(interface{ Equal(o interface{}) bool }); ok {
=======
		if eq, ok := v.(interface{ Equal(o any) bool }); ok {
>>>>>>> main
			if !eq.Equal(ov) {
				return false
			}
		} else if v != ov {
			// Fallback to a standard equality check if Value is unimplemented.
			return false
		}
	}
	return true
}

// String prints the attribute map. If any key or values throughout the map
// implement fmt.Stringer, it calls that method and appends.
func (a *Attributes) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	first := true
	for k, v := range a.m {
<<<<<<< HEAD
		var key, val string
		if str, ok := k.(interface{ String() string }); ok {
			key = str.String()
		}
		if str, ok := v.(interface{ String() string }); ok {
			val = str.String()
		}
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%q: %q, ", key, val))
=======
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%q: %q ", str(k), str(v)))
>>>>>>> main
		first = false
	}
	sb.WriteString("}")
	return sb.String()
}
<<<<<<< HEAD
=======

func str(x any) string {
	if v, ok := x.(fmt.Stringer); ok {
		return v.String()
	} else if v, ok := x.(string); ok {
		return v
	}
	return fmt.Sprintf("<%p>", x)
}

// MarshalJSON helps implement the json.Marshaler interface, thereby rendering
// the Attributes correctly when printing (via pretty.JSON) structs containing
// Attributes as fields.
//
// Is it impossible to unmarshal attributes from a JSON representation and this
// method is meant only for debugging purposes.
func (a *Attributes) MarshalJSON() ([]byte, error) {
	return []byte(a.String()), nil
}
>>>>>>> main
