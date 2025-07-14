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

// ControlPlaneKind is the name of the type used to represent objects
// of type 'control_plane'.
const ControlPlaneKind = "ControlPlane"

// ControlPlaneLinkKind is the name of the type used to represent links
// to objects of type 'control_plane'.
const ControlPlaneLinkKind = "ControlPlaneLink"

// ControlPlaneNilKind is the name of the type used to nil references
// to objects of type 'control_plane'.
const ControlPlaneNilKind = "ControlPlaneNil"

// ControlPlane represents the values of the 'control_plane' type.
//
// Representation of a Control Plane
type ControlPlane struct {
	bitmap_ uint32
	id      string
	href    string
	backup  *Backup
}

// Kind returns the name of the type of the object.
func (o *ControlPlane) Kind() string {
	if o == nil {
		return ControlPlaneNilKind
	}
	if o.bitmap_&1 != 0 {
		return ControlPlaneLinkKind
	}
	return ControlPlaneKind
}

// Link returns true if this is a link.
func (o *ControlPlane) Link() bool {
	return o != nil && o.bitmap_&1 != 0
}

// ID returns the identifier of the object.
func (o *ControlPlane) ID() string {
	if o != nil && o.bitmap_&2 != 0 {
		return o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *ControlPlane) GetID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.id
	}
	return
}

// HREF returns the link to the object.
func (o *ControlPlane) HREF() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *ControlPlane) GetHREF() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ControlPlane) Empty() bool {
	return o == nil || o.bitmap_&^1 == 0
}

// Backup returns the value of the 'backup' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Information about the backup of the control plane
func (o *ControlPlane) Backup() *Backup {
	if o != nil && o.bitmap_&8 != 0 {
		return o.backup
	}
	return nil
}

// GetBackup returns the value of the 'backup' attribute and
// a flag indicating if the attribute has a value.
//
// Information about the backup of the control plane
func (o *ControlPlane) GetBackup() (value *Backup, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.backup
	}
	return
}

// ControlPlaneListKind is the name of the type used to represent list of objects of
// type 'control_plane'.
const ControlPlaneListKind = "ControlPlaneList"

// ControlPlaneListLinkKind is the name of the type used to represent links to list
// of objects of type 'control_plane'.
const ControlPlaneListLinkKind = "ControlPlaneListLink"

// ControlPlaneNilKind is the name of the type used to nil lists of objects of
// type 'control_plane'.
const ControlPlaneListNilKind = "ControlPlaneListNil"

// ControlPlaneList is a list of values of the 'control_plane' type.
type ControlPlaneList struct {
	href  string
	link  bool
	items []*ControlPlane
}

// Kind returns the name of the type of the object.
func (l *ControlPlaneList) Kind() string {
	if l == nil {
		return ControlPlaneListNilKind
	}
	if l.link {
		return ControlPlaneListLinkKind
	}
	return ControlPlaneListKind
}

// Link returns true iif this is a link.
func (l *ControlPlaneList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *ControlPlaneList) HREF() string {
	if l != nil {
		return l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *ControlPlaneList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != ""
	if ok {
		value = l.href
	}
	return
}

// Len returns the length of the list.
func (l *ControlPlaneList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Items sets the items of the list.
func (l *ControlPlaneList) SetLink(link bool) {
	l.link = link
}

// Items sets the items of the list.
func (l *ControlPlaneList) SetHREF(href string) {
	l.href = href
}

// Items sets the items of the list.
func (l *ControlPlaneList) SetItems(items []*ControlPlane) {
	l.items = items
}

// Items returns the items of the list.
func (l *ControlPlaneList) Items() []*ControlPlane {
	if l == nil {
		return nil
	}
	return l.items
}

// Empty returns true if the list is empty.
func (l *ControlPlaneList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ControlPlaneList) Get(i int) *ControlPlane {
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
func (l *ControlPlaneList) Slice() []*ControlPlane {
	var slice []*ControlPlane
	if l == nil {
		slice = make([]*ControlPlane, 0)
	} else {
		slice = make([]*ControlPlane, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ControlPlaneList) Each(f func(item *ControlPlane) bool) {
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
func (l *ControlPlaneList) Range(f func(index int, item *ControlPlane) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
