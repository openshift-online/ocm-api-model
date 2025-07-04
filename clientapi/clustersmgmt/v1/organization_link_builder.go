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

// OrganizationLinkBuilder contains the data and logic needed to build 'organization_link' objects.
//
// Definition of an organization link.
type OrganizationLinkBuilder struct {
	bitmap_ uint32
	href    string
	id      string
}

// NewOrganizationLink creates a new builder of 'organization_link' objects.
func NewOrganizationLink() *OrganizationLinkBuilder {
	return &OrganizationLinkBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *OrganizationLinkBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// HREF sets the value of the 'HREF' attribute to the given value.
func (b *OrganizationLinkBuilder) HREF(value string) *OrganizationLinkBuilder {
	b.href = value
	b.bitmap_ |= 1
	return b
}

// ID sets the value of the 'ID' attribute to the given value.
func (b *OrganizationLinkBuilder) ID(value string) *OrganizationLinkBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *OrganizationLinkBuilder) Copy(object *OrganizationLink) *OrganizationLinkBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.href = object.href
	b.id = object.id
	return b
}

// Build creates a 'organization_link' object using the configuration stored in the builder.
func (b *OrganizationLinkBuilder) Build() (object *OrganizationLink, err error) {
	object = new(OrganizationLink)
	object.bitmap_ = b.bitmap_
	object.href = b.href
	object.id = b.id
	return
}
