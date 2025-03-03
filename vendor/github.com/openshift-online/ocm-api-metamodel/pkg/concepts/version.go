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
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
)

// Version is the representation of a version of a service.
type Version struct {
	documentedSupport
	annotatedSupport
	namedSupport

	// Service that owns this version:
	owner *Service

	// All the types of the version, indexed by name:
	types map[string]*Type

	// All the resources of the version, indexed by name:
	resources map[string]*Resource

	// All the error catagories of the version, indexed by name:
	errors map[string]*Error
}

// NewVersion creates a new version containing only the built-in types.
func NewVersion() *Version {
	// Create an empty version:
	version := new(Version)
	version.types = make(map[string]*Type)
	version.resources = make(map[string]*Resource)
	version.errors = make(map[string]*Error)

	// Add the built-in scalar types:
	version.addScalarType(nomenclator.Boolean)
	version.addScalarType(nomenclator.Integer)
	version.addScalarType(nomenclator.Long)
	version.addScalarType(nomenclator.Float)
	version.addScalarType(nomenclator.String)
	version.addScalarType(nomenclator.Date)

	// Add the built-in interface{} type:
	// We treat the built-in interface type as scalar for most purposes,
	// but define it as a new type to differentiate in how pointers are
	// generated.
	version.addInterfaceType(nomenclator.Interface)

	return version
}

// Owner returns the service that owns this version.
func (v *Version) Owner() *Service {
	return v.owner
}

// SetOwner sets the service that owns this version.
func (v *Version) SetOwner(value *Service) {
	v.owner = value
}

// Types returns the list of types that are part of this version.
func (v *Version) Types() TypeSlice {
	count := len(v.types)
	types := make(TypeSlice, count)
	index := 0
	for _, typ := range v.types {
		types[index] = typ
		index++
	}
	sort.Sort(types)
	return types
}

// FindType returns the type with the given name, or nil of there is no such type.
func (v *Version) FindType(name *names.Name) *Type {
	if name == nil {
		return nil
	}
	return v.types[name.String()]
}

// AddType adds the given type to the version.
func (v *Version) AddType(typ *Type) {
	if typ != nil {
		v.types[typ.Name().String()] = typ
		typ.SetOwner(v)
	}
}

// AddTypes adds the given types to the version.
func (v *Version) AddTypes(types []*Type) {
	for _, typ := range types {
		v.AddType(typ)
	}
}

// Boolean returns the boolean type.
func (v *Version) Boolean() *Type {
	return v.FindType(nomenclator.Boolean)
}

// IntegerType returns the integer type.
func (v *Version) IntegerType() *Type {
	return v.FindType(nomenclator.Integer)
}

// LongType returns the long type.
func (v *Version) LongType() *Type {
	return v.FindType(nomenclator.Long)
}

// StringType returns the string type.
func (v *Version) StringType() *Type {
	return v.FindType(nomenclator.String)
}

// FloatType returns the floating point type.
func (v *Version) FloatType() *Type {
	return v.FindType(nomenclator.Float)
}

// DateType returns the date type.
func (v *Version) DateType() *Type {
	return v.FindType(nomenclator.Date)
}

// InterfaceType returns the interface{} type.
func (v *Version) InterfaceType() *Type {
	return v.FindType(nomenclator.Interface)
}

// Resources returns the list of resources that are part of this version.
func (v *Version) Resources() ResourceSlice {
	count := len(v.resources)
	resources := make(ResourceSlice, count)
	index := 0
	for _, resource := range v.resources {
		resources[index] = resource
		index++
	}
	sort.Sort(resources)
	return resources
}

// FindResource returns the resource with the given name, or nil if there is no such resource.
func (v *Version) FindResource(name *names.Name) *Resource {
	if name == nil {
		return nil
	}
	return v.resources[name.String()]
}

// AddResource adds the given resource to the version.
func (v *Version) AddResource(resource *Resource) {
	if resource != nil {
		v.resources[resource.Name().String()] = resource
		resource.SetOwner(v)
	}
}

// AddResources adds the given resources to the version.
func (v *Version) AddResources(resources []*Resource) {
	for _, resource := range resources {
		v.AddResource(resource)
	}
}

// Root returns the root resource, or nil if there is no such resource.
func (v *Version) Root() *Resource {
	return v.resources[nomenclator.Root.String()]
}

// Errors returns the list of errors that are part of this version.
func (v *Version) Errors() ErrorSlice {
	count := len(v.errors)
	errors := make(ErrorSlice, count)
	index := 0
	for _, err := range v.errors {
		errors[index] = err
		index++
	}
	sort.Sort(errors)
	return errors
}

// FindError returns the error with the given name, or nil if there is no such error.
func (v *Version) FindError(name *names.Name) *Error {
	if name == nil {
		return nil
	}
	return v.errors[name.String()]
}

// AddError adds the given error to the version.
func (v *Version) AddError(err *Error) {
	if err != nil {
		v.errors[err.Name().String()] = err
		err.SetOwner(v)
	}
}

// AddErrors adds the given errors to the version.
func (v *Version) AddErrors(errors []*Error) {
	for _, err := range errors {
		v.AddError(err)
	}
}

// Paths returns the list of paths of this version. A path is a sequence of locators that go from
// the root of the version to all reachable resources.
func (v *Version) Paths() [][]*Locator {
	// Start with a set of paths that contains the locators of the root resource and with an
	// empty set of results:
	var results [][]*Locator
	var pending [][]*Locator
	for _, locator := range v.Root().Locators() {
		path := []*Locator{locator}
		pending = append(pending, path)
	}

	// Iterate the list of pending paths. For each pending path we will copy it to the list of
	// results, and then we will check if it can be extended. If it can be extended then the
	// extensions will be added to the list of pending paths.
	for len(pending) > 0 {
		// Get the first pending path and remove it from the list of pending paths:
		current := pending[0]
		pending = pending[1:]

		// Copy the path to the list of results:
		size := len(current)
		result := make([]*Locator, size)
		copy(result, current)
		results = append(results, result)

		// Check if the path can be extended, and if so extend it and add it to the list of
		// pending paths:
		last := current[size-1]
		for _, locator := range last.Target().Locators() {
			path := make([]*Locator, size+1)
			copy(path, current)
			path[size] = locator
			pending = append(pending, path)
		}
	}

	return results
}

func (v *Version) addScalarType(name *names.Name) {
	// Add the scalar type:
	scalarType := NewType()
	scalarType.SetKind(ScalarType)
	scalarType.SetName(name)
	v.AddType(scalarType)

	// Add the list type:
	listType := NewType()
	listType.SetKind(ListType)
	listType.SetName(names.Cat(name, nomenclator.List))
	listType.SetElement(scalarType)
	v.AddType(listType)
}

func (v *Version) addInterfaceType(name *names.Name) {
	// Add the interface{} type:
	interfaceType := NewType()
	interfaceType.SetKind(InterfaceType)
	interfaceType.SetName(name)
	v.AddType(interfaceType)

	// Add the list type:
	listType := NewType()
	listType.SetKind(ListType)
	listType.SetName(names.Cat(name, nomenclator.List))
	listType.SetElement(interfaceType)
	v.AddType(listType)
}

// VersionSlice is used to simplify sorting of slices of versions by name.
type VersionSlice []*Version

func (s VersionSlice) Len() int {
	return len(s)
}

func (s VersionSlice) Less(i, j int) bool {
	return names.Compare(s[i].name, s[j].name) == -1
}

func (s VersionSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
