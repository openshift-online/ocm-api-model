/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1alpha1 // github.com/openshift-online/ocm-api-model/clientapi/arohcp/v1alpha1

// ExternalAuthStatusBuilder contains the data and logic needed to build 'external_auth_status' objects.
//
// Representation of the status of an external authentication provider.
type ExternalAuthStatusBuilder struct {
	bitmap_ uint32
	id      string
	href    string
	message string
	state   *ExternalAuthStateBuilder
}

// NewExternalAuthStatus creates a new builder of 'external_auth_status' objects.
func NewExternalAuthStatus() *ExternalAuthStatusBuilder {
	return &ExternalAuthStatusBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *ExternalAuthStatusBuilder) Link(value bool) *ExternalAuthStatusBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *ExternalAuthStatusBuilder) ID(value string) *ExternalAuthStatusBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *ExternalAuthStatusBuilder) HREF(value string) *ExternalAuthStatusBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ExternalAuthStatusBuilder) Empty() bool {
	return b == nil || b.bitmap_&^1 == 0
}

// Message sets the value of the 'message' attribute to the given value.
func (b *ExternalAuthStatusBuilder) Message(value string) *ExternalAuthStatusBuilder {
	b.message = value
	b.bitmap_ |= 8
	return b
}

// State sets the value of the 'state' attribute to the given value.
//
// Representation of the state of an external authentication provider.
func (b *ExternalAuthStatusBuilder) State(value *ExternalAuthStateBuilder) *ExternalAuthStatusBuilder {
	b.state = value
	if value != nil {
		b.bitmap_ |= 16
	} else {
		b.bitmap_ &^= 16
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ExternalAuthStatusBuilder) Copy(object *ExternalAuthStatus) *ExternalAuthStatusBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	b.message = object.message
	if object.state != nil {
		b.state = NewExternalAuthState().Copy(object.state)
	} else {
		b.state = nil
	}
	return b
}

// Build creates a 'external_auth_status' object using the configuration stored in the builder.
func (b *ExternalAuthStatusBuilder) Build() (object *ExternalAuthStatus, err error) {
	object = new(ExternalAuthStatus)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	object.message = b.message
	if b.state != nil {
		object.state, err = b.state.Build()
		if err != nil {
			return
		}
	}
	return
}
