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
	"sort"

	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
)

// Resource represents an API resource.
type Resource struct {
	documentedSupport
	annotatedSupport
	namedSupport

	owner    *Version
	methods  MethodSlice
	locators LocatorSlice
}

// NewResource creates a new resource.
func NewResource() *Resource {
	return new(Resource)
}

// Owner returns the version that owns the resource.
func (r *Resource) Owner() *Version {
	return r.owner
}

// SetOwner sets the version that owns the resource.
func (r *Resource) SetOwner(value *Version) {
	r.owner = value
}

// Methods returns the methods of the resource.
func (r *Resource) Methods() MethodSlice {
	return r.methods
}

// AddMethod adds a method to the resource.
func (r *Resource) AddMethod(method *Method) {
	if method != nil {
		r.methods = append(r.methods, method)
		sort.Sort(r.methods)
		method.SetOwner(r)
	}
}

// FindMethod returns the method with the given name, or nil of there is no such method.
func (r *Resource) FindMethod(name *names.Name) *Method {
	for _, method := range r.methods {
		if method.Name().Equals(name) {
			return method
		}
	}
	return nil
}

// Locators returns the locators of the resource.
func (r *Resource) Locators() LocatorSlice {
	return r.locators
}

// VariableLocator returns the variable locator of the resource. If there is no such resource it
// returns nil.
func (r *Resource) VariableLocator() *Locator {
	for _, locator := range r.locators {
		if locator.Variable() {
			return locator
		}
	}
	return nil
}

// ConstantLocators returns the constant (non variable) locators of the resuurce.
func (r *Resource) ConstantLocators() LocatorSlice {
	locators := LocatorSlice{}
	for _, locator := range r.locators {
		if !locator.Variable() {
			locators = append(locators, locator)
		}
	}
	return locators
}

// AddLocator adds a locator to the resource.
func (r *Resource) AddLocator(locator *Locator) {
	if locator != nil {
		r.locators = append(r.locators, locator)
		sort.Sort(r.locators)
		locator.SetOwner(r)
	}
}

// IsRoot returns `true` if this is the root of the tree of resources of the version.
func (r *Resource) IsRoot() bool {
	return r.owner != nil && r == r.owner.Root()
}

// ResourceSlice is used to simplify sorting of slices of resources by name.
type ResourceSlice []*Resource

func (s ResourceSlice) Len() int {
	return len(s)
}

func (s ResourceSlice) Less(i, j int) bool {
	return names.Compare(s[i].name, s[j].name) == -1
}

func (s ResourceSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
