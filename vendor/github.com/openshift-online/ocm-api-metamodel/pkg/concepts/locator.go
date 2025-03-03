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

// Locator represents a resource locator, the reference from a resource to another resource.
type Locator struct {
	documentedSupport
	annotatedSupport
	namedSupport

	owner    *Resource
	variable bool
	target   *Resource
}

// NewLocator creates a new locator.
func NewLocator() *Locator {
	return new(Locator)
}

// Owner returns the resource that owns this locator.
func (l *Locator) Owner() *Resource {
	return l.owner
}

// SetOwner sets the resource that owns this locator.
func (l *Locator) SetOwner(value *Resource) {
	l.owner = value
}

// Variable returns the flag that indicates if the name of the referenced resource is a variable
// instead of a fixed URL segment. For example, in the reference from the clusters resource to the
// cluster resource the name is a variable, which contains the identifier of the cluster.
func (l *Locator) Variable() bool {
	return l.variable
}

// SetVariable sets the flag that indicates if the name of the referenced resource is a variable
// instead of a fixed URL segment.
func (l *Locator) SetVariable(value bool) {
	l.variable = true
}

// Target returns the resource that is referenced by the locator.
func (l *Locator) Target() *Resource {
	return l.target
}

// SetTarget sets the resource that is referenced by the locator.
func (l *Locator) SetTarget(value *Resource) {
	l.target = value
}

// LocatorSlice is used to simplify sorting of slices of locators by name.
type LocatorSlice []*Locator

func (s LocatorSlice) Len() int {
	return len(s)
}

func (s LocatorSlice) Less(i, j int) bool {
	return names.Compare(s[i].name, s[j].name) == -1
}

func (s LocatorSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
