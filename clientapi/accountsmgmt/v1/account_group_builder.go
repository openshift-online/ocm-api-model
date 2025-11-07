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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/accountsmgmt/v1

import (
	time "time"
)

// AccountGroupBuilder contains the data and logic needed to build 'account_group' objects.
type AccountGroupBuilder struct {
	bitmap_        uint32
	id             string
	href           string
	createdAt      time.Time
	description    string
	externalID     string
	managedBy      AccountGroupManagedBy
	name           string
	organizationID string
	updatedAt      time.Time
}

// NewAccountGroup creates a new builder of 'account_group' objects.
func NewAccountGroup() *AccountGroupBuilder {
	return &AccountGroupBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *AccountGroupBuilder) Link(value bool) *AccountGroupBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *AccountGroupBuilder) ID(value string) *AccountGroupBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *AccountGroupBuilder) HREF(value string) *AccountGroupBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AccountGroupBuilder) Empty() bool {
	return b == nil || b.bitmap_&^1 == 0
}

// CreatedAt sets the value of the 'created_at' attribute to the given value.
func (b *AccountGroupBuilder) CreatedAt(value time.Time) *AccountGroupBuilder {
	b.createdAt = value
	b.bitmap_ |= 8
	return b
}

// Description sets the value of the 'description' attribute to the given value.
func (b *AccountGroupBuilder) Description(value string) *AccountGroupBuilder {
	b.description = value
	b.bitmap_ |= 16
	return b
}

// ExternalID sets the value of the 'external_ID' attribute to the given value.
func (b *AccountGroupBuilder) ExternalID(value string) *AccountGroupBuilder {
	b.externalID = value
	b.bitmap_ |= 32
	return b
}

// ManagedBy sets the value of the 'managed_by' attribute to the given value.
func (b *AccountGroupBuilder) ManagedBy(value AccountGroupManagedBy) *AccountGroupBuilder {
	b.managedBy = value
	b.bitmap_ |= 64
	return b
}

// Name sets the value of the 'name' attribute to the given value.
func (b *AccountGroupBuilder) Name(value string) *AccountGroupBuilder {
	b.name = value
	b.bitmap_ |= 128
	return b
}

// OrganizationID sets the value of the 'organization_ID' attribute to the given value.
func (b *AccountGroupBuilder) OrganizationID(value string) *AccountGroupBuilder {
	b.organizationID = value
	b.bitmap_ |= 256
	return b
}

// UpdatedAt sets the value of the 'updated_at' attribute to the given value.
func (b *AccountGroupBuilder) UpdatedAt(value time.Time) *AccountGroupBuilder {
	b.updatedAt = value
	b.bitmap_ |= 512
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AccountGroupBuilder) Copy(object *AccountGroup) *AccountGroupBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	b.createdAt = object.createdAt
	b.description = object.description
	b.externalID = object.externalID
	b.managedBy = object.managedBy
	b.name = object.name
	b.organizationID = object.organizationID
	b.updatedAt = object.updatedAt
	return b
}

// Build creates a 'account_group' object using the configuration stored in the builder.
func (b *AccountGroupBuilder) Build() (object *AccountGroup, err error) {
	object = new(AccountGroup)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	object.createdAt = b.createdAt
	object.description = b.description
	object.externalID = b.externalID
	object.managedBy = b.managedBy
	object.name = b.name
	object.organizationID = b.organizationID
	object.updatedAt = b.updatedAt
	return
}
