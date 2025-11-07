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

// AccountGroupAssignmentBuilder contains the data and logic needed to build 'account_group_assignment' objects.
type AccountGroupAssignmentBuilder struct {
	bitmap_        uint32
	id             string
	href           string
	accountID      string
	accountGroup   *AccountGroupBuilder
	accountGroupID string
	createdAt      time.Time
	managedBy      AccountGroupAssignmentManagedBy
	updatedAt      time.Time
}

// NewAccountGroupAssignment creates a new builder of 'account_group_assignment' objects.
func NewAccountGroupAssignment() *AccountGroupAssignmentBuilder {
	return &AccountGroupAssignmentBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *AccountGroupAssignmentBuilder) Link(value bool) *AccountGroupAssignmentBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *AccountGroupAssignmentBuilder) ID(value string) *AccountGroupAssignmentBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *AccountGroupAssignmentBuilder) HREF(value string) *AccountGroupAssignmentBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AccountGroupAssignmentBuilder) Empty() bool {
	return b == nil || b.bitmap_&^1 == 0
}

// AccountID sets the value of the 'account_ID' attribute to the given value.
func (b *AccountGroupAssignmentBuilder) AccountID(value string) *AccountGroupAssignmentBuilder {
	b.accountID = value
	b.bitmap_ |= 8
	return b
}

// AccountGroup sets the value of the 'account_group' attribute to the given value.
func (b *AccountGroupAssignmentBuilder) AccountGroup(value *AccountGroupBuilder) *AccountGroupAssignmentBuilder {
	b.accountGroup = value
	if value != nil {
		b.bitmap_ |= 16
	} else {
		b.bitmap_ &^= 16
	}
	return b
}

// AccountGroupID sets the value of the 'account_group_ID' attribute to the given value.
func (b *AccountGroupAssignmentBuilder) AccountGroupID(value string) *AccountGroupAssignmentBuilder {
	b.accountGroupID = value
	b.bitmap_ |= 32
	return b
}

// CreatedAt sets the value of the 'created_at' attribute to the given value.
func (b *AccountGroupAssignmentBuilder) CreatedAt(value time.Time) *AccountGroupAssignmentBuilder {
	b.createdAt = value
	b.bitmap_ |= 64
	return b
}

// ManagedBy sets the value of the 'managed_by' attribute to the given value.
func (b *AccountGroupAssignmentBuilder) ManagedBy(value AccountGroupAssignmentManagedBy) *AccountGroupAssignmentBuilder {
	b.managedBy = value
	b.bitmap_ |= 128
	return b
}

// UpdatedAt sets the value of the 'updated_at' attribute to the given value.
func (b *AccountGroupAssignmentBuilder) UpdatedAt(value time.Time) *AccountGroupAssignmentBuilder {
	b.updatedAt = value
	b.bitmap_ |= 256
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AccountGroupAssignmentBuilder) Copy(object *AccountGroupAssignment) *AccountGroupAssignmentBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	b.accountID = object.accountID
	if object.accountGroup != nil {
		b.accountGroup = NewAccountGroup().Copy(object.accountGroup)
	} else {
		b.accountGroup = nil
	}
	b.accountGroupID = object.accountGroupID
	b.createdAt = object.createdAt
	b.managedBy = object.managedBy
	b.updatedAt = object.updatedAt
	return b
}

// Build creates a 'account_group_assignment' object using the configuration stored in the builder.
func (b *AccountGroupAssignmentBuilder) Build() (object *AccountGroupAssignment, err error) {
	object = new(AccountGroupAssignment)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	object.accountID = b.accountID
	if b.accountGroup != nil {
		object.accountGroup, err = b.accountGroup.Build()
		if err != nil {
			return
		}
	}
	object.accountGroupID = b.accountGroupID
	object.createdAt = b.createdAt
	object.managedBy = b.managedBy
	object.updatedAt = b.updatedAt
	return
}
