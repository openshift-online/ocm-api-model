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

// AWSCapacityReservationKind is the name of the type used to represent objects
// of type 'AWS_capacity_reservation'.
const AWSCapacityReservationKind = "AWSCapacityReservation"

// AWSCapacityReservationLinkKind is the name of the type used to represent links
// to objects of type 'AWS_capacity_reservation'.
const AWSCapacityReservationLinkKind = "AWSCapacityReservationLink"

// AWSCapacityReservationNilKind is the name of the type used to nil references
// to objects of type 'AWS_capacity_reservation'.
const AWSCapacityReservationNilKind = "AWSCapacityReservationNil"

// AWSCapacityReservation represents the values of the 'AWS_capacity_reservation' type.
//
// AWS Capacity Reservation specification.
type AWSCapacityReservation struct {
	fieldSet_  []bool
	id         string
	href       string
	id         string
	marketType MarketType
}

// Kind returns the name of the type of the object.
func (o *AWSCapacityReservation) Kind() string {
	if o == nil {
		return AWSCapacityReservationNilKind
	}
	if len(o.fieldSet_) > 0 && o.fieldSet_[0] {
		return AWSCapacityReservationLinkKind
	}
	return AWSCapacityReservationKind
}

// Link returns true if this is a link.
func (o *AWSCapacityReservation) Link() bool {
	return o != nil && len(o.fieldSet_) > 0 && o.fieldSet_[0]
}

// ID returns the identifier of the object.
func (o *AWSCapacityReservation) ID() string {
	if o != nil && len(o.fieldSet_) > 1 && o.fieldSet_[1] {
		return o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *AWSCapacityReservation) GetID() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 1 && o.fieldSet_[1]
	if ok {
		value = o.id
	}
	return
}

// HREF returns the link to the object.
func (o *AWSCapacityReservation) HREF() string {
	if o != nil && len(o.fieldSet_) > 2 && o.fieldSet_[2] {
		return o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *AWSCapacityReservation) GetHREF() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 2 && o.fieldSet_[2]
	if ok {
		value = o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *AWSCapacityReservation) Empty() bool {
	if o == nil || len(o.fieldSet_) == 0 {
		return true
	}

	// Check all fields except the link flag (index 0)
	for i := 1; i < len(o.fieldSet_); i++ {
		if o.fieldSet_[i] {
			return false
		}
	}
	return true
}

// Id returns the value of the 'id' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Specify the target Capacity Reservation in which the EC2 instances will be launched.
func (o *AWSCapacityReservation) Id() string {
	if o != nil && len(o.fieldSet_) > 3 && o.fieldSet_[3] {
		return o.id
	}
	return ""
}

// GetId returns the value of the 'id' attribute and
// a flag indicating if the attribute has a value.
//
// Specify the target Capacity Reservation in which the EC2 instances will be launched.
func (o *AWSCapacityReservation) GetId() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 3 && o.fieldSet_[3]
	if ok {
		value = o.id
	}
	return
}

// MarketType returns the value of the 'market_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// marketType specifies the market type of the CapacityReservation for the EC2
// instances. Valid values are OnDemand, CapacityBlocks.
// "OnDemand": EC2 instances run as standard On-Demand instances.
// "CapacityBlocks": scheduled pre-purchased compute capacity.
func (o *AWSCapacityReservation) MarketType() MarketType {
	if o != nil && len(o.fieldSet_) > 4 && o.fieldSet_[4] {
		return o.marketType
	}
	return MarketType("")
}

// GetMarketType returns the value of the 'market_type' attribute and
// a flag indicating if the attribute has a value.
//
// marketType specifies the market type of the CapacityReservation for the EC2
// instances. Valid values are OnDemand, CapacityBlocks.
// "OnDemand": EC2 instances run as standard On-Demand instances.
// "CapacityBlocks": scheduled pre-purchased compute capacity.
func (o *AWSCapacityReservation) GetMarketType() (value MarketType, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 4 && o.fieldSet_[4]
	if ok {
		value = o.marketType
	}
	return
}

// AWSCapacityReservationListKind is the name of the type used to represent list of objects of
// type 'AWS_capacity_reservation'.
const AWSCapacityReservationListKind = "AWSCapacityReservationList"

// AWSCapacityReservationListLinkKind is the name of the type used to represent links to list
// of objects of type 'AWS_capacity_reservation'.
const AWSCapacityReservationListLinkKind = "AWSCapacityReservationListLink"

// AWSCapacityReservationNilKind is the name of the type used to nil lists of objects of
// type 'AWS_capacity_reservation'.
const AWSCapacityReservationListNilKind = "AWSCapacityReservationListNil"

// AWSCapacityReservationList is a list of values of the 'AWS_capacity_reservation' type.
type AWSCapacityReservationList struct {
	href  string
	link  bool
	items []*AWSCapacityReservation
}

// Kind returns the name of the type of the object.
func (l *AWSCapacityReservationList) Kind() string {
	if l == nil {
		return AWSCapacityReservationListNilKind
	}
	if l.link {
		return AWSCapacityReservationListLinkKind
	}
	return AWSCapacityReservationListKind
}

// Link returns true iif this is a link.
func (l *AWSCapacityReservationList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *AWSCapacityReservationList) HREF() string {
	if l != nil {
		return l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *AWSCapacityReservationList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != ""
	if ok {
		value = l.href
	}
	return
}

// Len returns the length of the list.
func (l *AWSCapacityReservationList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Items sets the items of the list.
func (l *AWSCapacityReservationList) SetLink(link bool) {
	l.link = link
}

// Items sets the items of the list.
func (l *AWSCapacityReservationList) SetHREF(href string) {
	l.href = href
}

// Items sets the items of the list.
func (l *AWSCapacityReservationList) SetItems(items []*AWSCapacityReservation) {
	l.items = items
}

// Items returns the items of the list.
func (l *AWSCapacityReservationList) Items() []*AWSCapacityReservation {
	if l == nil {
		return nil
	}
	return l.items
}

// Empty returns true if the list is empty.
func (l *AWSCapacityReservationList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *AWSCapacityReservationList) Get(i int) *AWSCapacityReservation {
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
func (l *AWSCapacityReservationList) Slice() []*AWSCapacityReservation {
	var slice []*AWSCapacityReservation
	if l == nil {
		slice = make([]*AWSCapacityReservation, 0)
	} else {
		slice = make([]*AWSCapacityReservation, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *AWSCapacityReservationList) Each(f func(item *AWSCapacityReservation) bool) {
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
func (l *AWSCapacityReservationList) Range(f func(index int, item *AWSCapacityReservation) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
