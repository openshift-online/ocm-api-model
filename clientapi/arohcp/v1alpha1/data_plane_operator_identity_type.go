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

// DataPlaneOperatorIdentity represents the values of the 'data_plane_operator_identity' type.
type DataPlaneOperatorIdentity struct {
	bitmap_             uint32
	maxOpenShiftVersion string
	minOpenShiftVersion string
	operatorName        string
	required            string
	roleDefinitions     []*RoleDefinition
	serviceAccounts     []*IdentityServiceAccount
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *DataPlaneOperatorIdentity) Empty() bool {
	return o == nil || o.bitmap_ == 0
}

// MaxOpenShiftVersion returns the value of the 'max_open_shift_version' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The field is a string and it is of format X.Y (e.g 4.18) where X and Y are major and
// minor segments of the OpenShift version respectively.
// Not specifying it indicates support for this operator in all Openshift versions,
// starting from min_openshift_version if min_openshift_version is defined.
func (o *DataPlaneOperatorIdentity) MaxOpenShiftVersion() string {
	if o != nil && o.bitmap_&1 != 0 {
		return o.maxOpenShiftVersion
	}
	return ""
}

// GetMaxOpenShiftVersion returns the value of the 'max_open_shift_version' attribute and
// a flag indicating if the attribute has a value.
//
// The field is a string and it is of format X.Y (e.g 4.18) where X and Y are major and
// minor segments of the OpenShift version respectively.
// Not specifying it indicates support for this operator in all Openshift versions,
// starting from min_openshift_version if min_openshift_version is defined.
func (o *DataPlaneOperatorIdentity) GetMaxOpenShiftVersion() (value string, ok bool) {
	ok = o != nil && o.bitmap_&1 != 0
	if ok {
		value = o.maxOpenShiftVersion
	}
	return
}

// MinOpenShiftVersion returns the value of the 'min_open_shift_version' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The field is a string and it is of format X.Y (e.g 4.18) where X and Y are major and
// minor segments of the OpenShift version respectively.
// Not specifying it indicates support for this operator in all Openshift versions,
// or up to max_openshift_version, if defined.
func (o *DataPlaneOperatorIdentity) MinOpenShiftVersion() string {
	if o != nil && o.bitmap_&2 != 0 {
		return o.minOpenShiftVersion
	}
	return ""
}

// GetMinOpenShiftVersion returns the value of the 'min_open_shift_version' attribute and
// a flag indicating if the attribute has a value.
//
// The field is a string and it is of format X.Y (e.g 4.18) where X and Y are major and
// minor segments of the OpenShift version respectively.
// Not specifying it indicates support for this operator in all Openshift versions,
// or up to max_openshift_version, if defined.
func (o *DataPlaneOperatorIdentity) GetMinOpenShiftVersion() (value string, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.minOpenShiftVersion
	}
	return
}

// OperatorName returns the value of the 'operator_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The name of the data plane operator that needs the identity
func (o *DataPlaneOperatorIdentity) OperatorName() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.operatorName
	}
	return ""
}

// GetOperatorName returns the value of the 'operator_name' attribute and
// a flag indicating if the attribute has a value.
//
// The name of the data plane operator that needs the identity
func (o *DataPlaneOperatorIdentity) GetOperatorName() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.operatorName
	}
	return
}

// Required returns the value of the 'required' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates the whether the identity is always required or not
// "always" means that the identity is always required
// "on_enablement" means that the identity is only required when a functionality
// that leverages the operator is enabled.
// Possible values are ("always", "on_enablement")
func (o *DataPlaneOperatorIdentity) Required() string {
	if o != nil && o.bitmap_&8 != 0 {
		return o.required
	}
	return ""
}

// GetRequired returns the value of the 'required' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates the whether the identity is always required or not
// "always" means that the identity is always required
// "on_enablement" means that the identity is only required when a functionality
// that leverages the operator is enabled.
// Possible values are ("always", "on_enablement")
func (o *DataPlaneOperatorIdentity) GetRequired() (value string, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.required
	}
	return
}

// RoleDefinitions returns the value of the 'role_definitions' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// A list of roles that are required by the operator
func (o *DataPlaneOperatorIdentity) RoleDefinitions() []*RoleDefinition {
	if o != nil && o.bitmap_&16 != 0 {
		return o.roleDefinitions
	}
	return nil
}

// GetRoleDefinitions returns the value of the 'role_definitions' attribute and
// a flag indicating if the attribute has a value.
//
// A list of roles that are required by the operator
func (o *DataPlaneOperatorIdentity) GetRoleDefinitions() (value []*RoleDefinition, ok bool) {
	ok = o != nil && o.bitmap_&16 != 0
	if ok {
		value = o.roleDefinitions
	}
	return
}

// ServiceAccounts returns the value of the 'service_accounts' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// It is a list of K8s ServiceAccounts leveraged by the operator.
// There must be at least a single service account specified.
// This information is needed to federate a managed identity to a k8s subject.
// There should be no duplicated "name:namespace" entries within this field.
func (o *DataPlaneOperatorIdentity) ServiceAccounts() []*IdentityServiceAccount {
	if o != nil && o.bitmap_&32 != 0 {
		return o.serviceAccounts
	}
	return nil
}

// GetServiceAccounts returns the value of the 'service_accounts' attribute and
// a flag indicating if the attribute has a value.
//
// It is a list of K8s ServiceAccounts leveraged by the operator.
// There must be at least a single service account specified.
// This information is needed to federate a managed identity to a k8s subject.
// There should be no duplicated "name:namespace" entries within this field.
func (o *DataPlaneOperatorIdentity) GetServiceAccounts() (value []*IdentityServiceAccount, ok bool) {
	ok = o != nil && o.bitmap_&32 != 0
	if ok {
		value = o.serviceAccounts
	}
	return
}

// DataPlaneOperatorIdentityListKind is the name of the type used to represent list of objects of
// type 'data_plane_operator_identity'.
const DataPlaneOperatorIdentityListKind = "DataPlaneOperatorIdentityList"

// DataPlaneOperatorIdentityListLinkKind is the name of the type used to represent links to list
// of objects of type 'data_plane_operator_identity'.
const DataPlaneOperatorIdentityListLinkKind = "DataPlaneOperatorIdentityListLink"

// DataPlaneOperatorIdentityNilKind is the name of the type used to nil lists of objects of
// type 'data_plane_operator_identity'.
const DataPlaneOperatorIdentityListNilKind = "DataPlaneOperatorIdentityListNil"

// DataPlaneOperatorIdentityList is a list of values of the 'data_plane_operator_identity' type.
type DataPlaneOperatorIdentityList struct {
	href  string
	link  bool
	items []*DataPlaneOperatorIdentity
}

// Len returns the length of the list.
func (l *DataPlaneOperatorIdentityList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Items sets the items of the list.
func (l *DataPlaneOperatorIdentityList) SetLink(link bool) {
	l.link = link
}

// Items sets the items of the list.
func (l *DataPlaneOperatorIdentityList) SetHREF(href string) {
	l.href = href
}

// Items sets the items of the list.
func (l *DataPlaneOperatorIdentityList) SetItems(items []*DataPlaneOperatorIdentity) {
	l.items = items
}

// Items returns the items of the list.
func (l *DataPlaneOperatorIdentityList) Items() []*DataPlaneOperatorIdentity {
	if l == nil {
		return nil
	}
	return l.items
}

// Empty returns true if the list is empty.
func (l *DataPlaneOperatorIdentityList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *DataPlaneOperatorIdentityList) Get(i int) *DataPlaneOperatorIdentity {
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
func (l *DataPlaneOperatorIdentityList) Slice() []*DataPlaneOperatorIdentity {
	var slice []*DataPlaneOperatorIdentity
	if l == nil {
		slice = make([]*DataPlaneOperatorIdentity, 0)
	} else {
		slice = make([]*DataPlaneOperatorIdentity, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *DataPlaneOperatorIdentityList) Each(f func(item *DataPlaneOperatorIdentity) bool) {
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
func (l *DataPlaneOperatorIdentityList) Range(f func(index int, item *DataPlaneOperatorIdentity) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
