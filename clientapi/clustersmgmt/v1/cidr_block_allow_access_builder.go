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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1

// CIDRBlockAllowAccessBuilder contains the data and logic needed to build 'CIDR_block_allow_access' objects.
type CIDRBlockAllowAccessBuilder struct {
	bitmap_ uint32
	mode    string
	values  []string
}

// NewCIDRBlockAllowAccess creates a new builder of 'CIDR_block_allow_access' objects.
func NewCIDRBlockAllowAccess() *CIDRBlockAllowAccessBuilder {
	return &CIDRBlockAllowAccessBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *CIDRBlockAllowAccessBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Mode sets the value of the 'mode' attribute to the given value.
func (b *CIDRBlockAllowAccessBuilder) Mode(value string) *CIDRBlockAllowAccessBuilder {
	b.mode = value
	b.bitmap_ |= 1
	return b
}

// Values sets the value of the 'values' attribute to the given values.
func (b *CIDRBlockAllowAccessBuilder) Values(values ...string) *CIDRBlockAllowAccessBuilder {
	b.values = make([]string, len(values))
	copy(b.values, values)
	b.bitmap_ |= 2
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *CIDRBlockAllowAccessBuilder) Copy(object *CIDRBlockAllowAccess) *CIDRBlockAllowAccessBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.mode = object.mode
	if object.values != nil {
		b.values = make([]string, len(object.values))
		copy(b.values, object.values)
	} else {
		b.values = nil
	}
	return b
}

// Build creates a 'CIDR_block_allow_access' object using the configuration stored in the builder.
func (b *CIDRBlockAllowAccessBuilder) Build() (object *CIDRBlockAllowAccess, err error) {
	object = new(CIDRBlockAllowAccess)
	object.bitmap_ = b.bitmap_
	object.mode = b.mode
	if b.values != nil {
		object.values = make([]string, len(b.values))
		copy(object.values, b.values)
	}
	return
}
