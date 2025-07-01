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

// CIDRBlocksAccess represents the values of the 'CIDR_blocks_access' type.
//
// Describes the CIDR Blocks access to the Kubernetes API server.
// Currently, only supported for ARO-HCP based clusters.
type CIDRBlocksAccess struct {
	fieldSet_ []bool
	allow     *CIDRBlocksAllowAccess
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *CIDRBlocksAccess) Empty() bool {
	if o == nil || len(o.fieldSet_) == 0 {
		return true
	}
	for _, set := range o.fieldSet_ {
		if set {
			return false
		}
	}
	return true
}

// Allow returns the value of the 'allow' attribute, or
// the zero value of the type if the attribute doesn't have a value.
func (o *CIDRBlocksAccess) Allow() *CIDRBlocksAllowAccess {
	if o != nil && len(o.fieldSet_) > 0 && o.fieldSet_[0] {
		return o.allow
	}
	return nil
}

// GetAllow returns the value of the 'allow' attribute and
// a flag indicating if the attribute has a value.
func (o *CIDRBlocksAccess) GetAllow() (value *CIDRBlocksAllowAccess, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 0 && o.fieldSet_[0]
	if ok {
		value = o.allow
	}
	return
}

// CIDRBlocksAccessListKind is the name of the type used to represent list of objects of
// type 'CIDR_blocks_access'.
const CIDRBlocksAccessListKind = "CIDRBlocksAccessList"

// CIDRBlocksAccessListLinkKind is the name of the type used to represent links to list
// of objects of type 'CIDR_blocks_access'.
const CIDRBlocksAccessListLinkKind = "CIDRBlocksAccessListLink"

// CIDRBlocksAccessNilKind is the name of the type used to nil lists of objects of
// type 'CIDR_blocks_access'.
const CIDRBlocksAccessListNilKind = "CIDRBlocksAccessListNil"

// CIDRBlocksAccessList is a list of values of the 'CIDR_blocks_access' type.
type CIDRBlocksAccessList struct {
	href  string
	link  bool
	items []*CIDRBlocksAccess
}

// Len returns the length of the list.
func (l *CIDRBlocksAccessList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Items sets the items of the list.
func (l *CIDRBlocksAccessList) SetLink(link bool) {
	l.link = link
}

// Items sets the items of the list.
func (l *CIDRBlocksAccessList) SetHREF(href string) {
	l.href = href
}

// Items sets the items of the list.
func (l *CIDRBlocksAccessList) SetItems(items []*CIDRBlocksAccess) {
	l.items = items
}

// Items returns the items of the list.
func (l *CIDRBlocksAccessList) Items() []*CIDRBlocksAccess {
	if l == nil {
		return nil
	}
	return l.items
}

// Empty returns true if the list is empty.
func (l *CIDRBlocksAccessList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *CIDRBlocksAccessList) Get(i int) *CIDRBlocksAccess {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *CIDRBlocksAccessList) Slice() []*CIDRBlocksAccess {
	var slice []*CIDRBlocksAccess
	if l == nil {
		slice = make([]*CIDRBlocksAccess, 0)
	} else {
		slice = make([]*CIDRBlocksAccess, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *CIDRBlocksAccessList) Each(f func(item *CIDRBlocksAccess) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *CIDRBlocksAccessList) Range(f func(index int, item *CIDRBlocksAccess) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
