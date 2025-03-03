/*
Copyright (c) 2019 Red Hat, Inc.

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

package concepts

import (
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
)

// Attribute is the representation of an attribute of an structured type.
type Attribute struct {
	documentedSupport
	annotatedSupport
	namedSupport
	typedSupport

	owner     *Type
	link      bool
	linkOwner *Version
}

func (a *Attribute) SetLinkOwner(v *Version) {
	a.linkOwner = v
}

func (a *Attribute) LinkOwner() *Version {
	return a.linkOwner
}

// NewAttribute creates a new attribute.
func NewAttribute() *Attribute {
	return new(Attribute)
}

// Owner returns the type that owns this attribute.
func (a *Attribute) Owner() *Type {
	return a.owner
}

// SetOwner sets the type that owns this attribute.
func (a *Attribute) SetOwner(value *Type) {
	a.owner = value
}

// Link returns true if the attribute is a link, false otherwise.
func (a *Attribute) Link() bool {
	return a.link
}

// SetLink sets the flag that indicates if this attribute is a link.
func (a *Attribute) SetLink(value bool) {
	a.link = value
}

// AttributeSlice is used to simplify sorting of slices of attributes by name.
type AttributeSlice []*Attribute

func (s AttributeSlice) Len() int {
	return len(s)
}

func (s AttributeSlice) Less(i, j int) bool {
	return names.Compare(s[i].name, s[j].name) == -1
}

func (s AttributeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
