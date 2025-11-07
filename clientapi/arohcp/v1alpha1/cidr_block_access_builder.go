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

// CIDRBlockAccessBuilder contains the data and logic needed to build 'CIDR_block_access' objects.
//
// Describes the CIDR Block access policy to the Kubernetes API server.
// Currently, only supported for ARO-HCP based clusters.
// The default policy mode is "allow_all" that is, all access is allowed.
type CIDRBlockAccessBuilder struct {
	bitmap_ uint32
	allow   *CIDRBlockAllowAccessBuilder
}

// NewCIDRBlockAccess creates a new builder of 'CIDR_block_access' objects.
func NewCIDRBlockAccess() *CIDRBlockAccessBuilder {
	return &CIDRBlockAccessBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *CIDRBlockAccessBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Allow sets the value of the 'allow' attribute to the given value.
func (b *CIDRBlockAccessBuilder) Allow(value *CIDRBlockAllowAccessBuilder) *CIDRBlockAccessBuilder {
	b.allow = value
	if value != nil {
		b.bitmap_ |= 1
	} else {
		b.bitmap_ &^= 1
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *CIDRBlockAccessBuilder) Copy(object *CIDRBlockAccess) *CIDRBlockAccessBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	if object.allow != nil {
		b.allow = NewCIDRBlockAllowAccess().Copy(object.allow)
	} else {
		b.allow = nil
	}
	return b
}

// Build creates a 'CIDR_block_access' object using the configuration stored in the builder.
func (b *CIDRBlockAccessBuilder) Build() (object *CIDRBlockAccess, err error) {
	object = new(CIDRBlockAccess)
	object.bitmap_ = b.bitmap_
	if b.allow != nil {
		object.allow, err = b.allow.Build()
		if err != nil {
			return
		}
	}
	return
}
