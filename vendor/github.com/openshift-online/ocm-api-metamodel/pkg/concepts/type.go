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

// TypeKind specifies the kind of a type. It can be scalar, enum, struct, list or class.
type TypeKind int

// Values of the TypeKind type:
const (
	UndefinedType TypeKind = iota
	ClassType
	EnumType
	InterfaceType
	ListType
	MapType
	ScalarType
	StructType
)

// StringType generates the string representation of a type kind.
func (k TypeKind) String() string {
	switch k {
	case ClassType:
		return "class"
	case EnumType:
		return "enum"
	case InterfaceType:
		return "interface{}"
	case ListType:
		return "list"
	case MapType:
		return "map"
	case ScalarType:
		return "scalar"
	case StructType:
		return "struct"
	default:
		return "unknown"
	}
}

// NewType creates a new type.
func NewType() *Type {
	typ := new(Type)
	typ.kind = UndefinedType
	return typ
}

// Type specifies the data type of attributes of structs and method parameters.
type Type struct {
	documentedSupport
	annotatedSupport
	namedSupport

	owner            *Version
	kind             TypeKind
	attributes       AttributeSlice
	values           EnumValueSlice
	element          *Type
	index            *Type
	explicitDeclared bool
}

// ExplicitDeclared returns true if this type is explicitDeclared
// with a reference
func (t *Type) ExplicitDeclared() bool {
	return t.explicitDeclared
}

// SetExplicitDeclared sets ExplicitDeclared
func (t *Type) SetExplicitDeclared(value bool) {
	t.explicitDeclared = value
}

// Owner returns the version that owns this type.
func (t *Type) Owner() *Version {
	return t.owner
}

// SetOwner sets the version that owns this type.
func (t *Type) SetOwner(value *Version) {
	t.owner = value
}

// Kind returns the kind of this type.
func (t *Type) Kind() TypeKind {
	return t.kind
}

// SetKind sets the kind of this type.
func (t *Type) SetKind(value TypeKind) {
	t.kind = value
}

// IsBoolean returns true iff this type is the built-in boolean type.
func (t *Type) IsBoolean() bool {
	return t.owner != nil && t == t.owner.Boolean()
}

// IsInteger returns true iff this type is the built-in integer type.
func (t *Type) IsInteger() bool {
	return t.owner != nil && t == t.owner.IntegerType()
}

// IsLong returns true iff this type is the built-in long type.
func (t *Type) IsLong() bool {
	return t.owner != nil && t == t.owner.LongType()
}

// IsFloat returns true iff this type is the built-in float type.
func (t *Type) IsFloat() bool {
	return t.owner != nil && t == t.owner.FloatType()
}

// IsString returns true iff this type is the built-in string type.
func (t *Type) IsString() bool {
	return t.owner != nil && t == t.owner.StringType()
}

// IsDate returns true iff this type is the built-in date type.
func (t *Type) IsDate() bool {
	return t.owner != nil && t == t.owner.DateType()
}

// IsClass returns true iff this type is a class type.
func (t *Type) IsClass() bool {
	return t.kind == ClassType
}

// IsEnum returns true iff this type is an enum type.
func (t *Type) IsEnum() bool {
	return t.kind == EnumType
}

// IsInterface returns true iff this type is an interface{} type.
func (t *Type) IsInterface() bool {
	return t.kind == InterfaceType
}

// IsList returns true iff this type is a list type.
func (t *Type) IsList() bool {
	return t.kind == ListType
}

// IsMap returns true iff this type is a map type.
func (t *Type) IsMap() bool {
	return t.kind == MapType
}

// IsScalar returns true iff this type is an scalar type. Note that interface types are also considered
// scalar types due to their opaque nature in the SDK.
func (t *Type) IsScalar() bool {
	if t == nil {
		return false
	}
	return t.kind == ScalarType || t.kind == EnumType || t.kind == InterfaceType
}

func (t *Type) IsBasicType() bool {
	if t == nil {
		return false
	}

	return t.kind == ScalarType
}

// IsStruct returns true iff this type is an struct type. Note that class types are also considered
// struct types.
func (t *Type) IsStruct() bool {
	return t.kind == ClassType || t.kind == StructType
}

// Attributes returns the list of attributes of an struct type. If called for any other kind of type
// it will return nil.
func (t *Type) Attributes() AttributeSlice {
	return t.attributes
}

// AddAttribute adds an attribute to the type, assuming hat it is an structured type.
func (t *Type) AddAttribute(attribute *Attribute) {
	if attribute != nil {
		t.attributes = append(t.attributes, attribute)
		sort.Sort(t.attributes)
		attribute.SetOwner(t)
	}
}

func (t *Type) AddAttributes(attributes AttributeSlice) {
	if attributes.Len() > 0 {
		for _, att := range attributes {
			t.AddAttribute(att)
		}
	}
}

// FindAttribute returns the attribute with the given name, or nil if no such attribute exists.
func (t *Type) FindAttribute(name *names.Name) *Attribute {
	for _, attribute := range t.attributes {
		if attribute.Name().Equals(name) {
			return attribute
		}
	}
	return nil
}

// Values returns the list of values of an enumerated type. If called for any other kind of type it
// will return nil.
func (t *Type) Values() EnumValueSlice {
	return t.values
}

// AddValue adds an enumerated value to the type, assuming that it is an enumerated type.
func (t *Type) AddValue(value *EnumValue) {
	if value != nil {
		t.values = append(t.values, value)
		sort.Sort(t.values)
		value.SetType(t)
	}
}

// Element returns the element type for a list type. If called for any other kind of type it will
// return nil.
func (t *Type) Element() *Type {
	return t.element
}

// SetElement sets the element type for a list type.
func (t *Type) SetElement(value *Type) {
	t.element = value
}

// Index returns the index type for a list type. If called for any other kind of type it will
// return nil.
func (t *Type) Index() *Type {
	return t.index
}

// SetIndex sets the index type for a map type.
func (t *Type) SetIndex(value *Type) {
	t.index = value
}

// TypeSlice is used to simplify sorting of slices of types by name.
type TypeSlice []*Type

func (s TypeSlice) Len() int {
	return len(s)
}

func (s TypeSlice) Less(i, j int) bool {
	return names.Compare(s[i].name, s[j].name) == -1
}

func (s TypeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// EnumValue represents each of the values of an enum type.
type EnumValue struct {
	documentedSupport
	annotatedSupport
	namedSupport

	typ *Type
}

// NewEnumValue creates a new enumerated type value.
func NewEnumValue() *EnumValue {
	return new(EnumValue)
}

// Type returns the enum type that owns this value.
func (v *EnumValue) Type() *Type {
	return v.typ
}

// SetType sets the enum type that owns this value.
func (v *EnumValue) SetType(value *Type) {
	v.typ = value
}

// EnumValueSlice is used to simplify sorting of slices of enum values by name.
type EnumValueSlice []*EnumValue

func (s EnumValueSlice) Len() int {
	return len(s)
}

func (s EnumValueSlice) Less(i, j int) bool {
	return names.Compare(s[i].name, s[j].name) == -1
}

func (s EnumValueSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
