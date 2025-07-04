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

import (
	time "time"
)

// ExternalAuthStateBuilder contains the data and logic needed to build 'external_auth_state' objects.
//
// Representation of the state of an external authentication provider.
type ExternalAuthStateBuilder struct {
	bitmap_              uint32
	lastUpdatedTimestamp time.Time
	value                string
}

// NewExternalAuthState creates a new builder of 'external_auth_state' objects.
func NewExternalAuthState() *ExternalAuthStateBuilder {
	return &ExternalAuthStateBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ExternalAuthStateBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// LastUpdatedTimestamp sets the value of the 'last_updated_timestamp' attribute to the given value.
func (b *ExternalAuthStateBuilder) LastUpdatedTimestamp(value time.Time) *ExternalAuthStateBuilder {
	b.lastUpdatedTimestamp = value
	b.bitmap_ |= 1
	return b
}

// Value sets the value of the 'value' attribute to the given value.
func (b *ExternalAuthStateBuilder) Value(value string) *ExternalAuthStateBuilder {
	b.value = value
	b.bitmap_ |= 2
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ExternalAuthStateBuilder) Copy(object *ExternalAuthState) *ExternalAuthStateBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.lastUpdatedTimestamp = object.lastUpdatedTimestamp
	b.value = object.value
	return b
}

// Build creates a 'external_auth_state' object using the configuration stored in the builder.
func (b *ExternalAuthStateBuilder) Build() (object *ExternalAuthState, err error) {
	object = new(ExternalAuthState)
	object.bitmap_ = b.bitmap_
	object.lastUpdatedTimestamp = b.lastUpdatedTimestamp
	object.value = b.value
	return
}
