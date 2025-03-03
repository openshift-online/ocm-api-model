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

package golang

import (
	"fmt"
	"path"

	"github.com/openshift-online/ocm-api-metamodel/pkg/annotations"
	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// TypesCalculatorBuilder is an object used to configure and build the Go type calculators. Don't
// create instances directly, use the NewTypesCalculator function instead.
type TypesCalculatorBuilder struct {
	reporter *reporter.Reporter
	packages *PackagesCalculator
	names    *NamesCalculator
}

// TypesCalculator is an object used to calculate Go types. Don't create instances directly, use the
// builder instead.
type TypesCalculator struct {
	reporter *reporter.Reporter
	packages *PackagesCalculator
	names    *NamesCalculator
}

// NewTypesCalculator creates a Go names calculator builder.
func NewTypesCalculator() *TypesCalculatorBuilder {
	return &TypesCalculatorBuilder{}
}

// Reporter sets the object that will be used to report information about the calculation processes,
// including errors.
func (b *TypesCalculatorBuilder) Reporter(value *reporter.Reporter) *TypesCalculatorBuilder {
	b.reporter = value
	return b
}

// Packages sets the object that will be used to calculate package names.
func (b *TypesCalculatorBuilder) Packages(value *PackagesCalculator) *TypesCalculatorBuilder {
	b.packages = value
	return b
}

// Names sets the names calculator object that will be used to calculate Go names.
func (b *TypesCalculatorBuilder) Names(value *NamesCalculator) *TypesCalculatorBuilder {
	b.names = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new
// calculator using it.
func (b *TypesCalculatorBuilder) Build() (calculator *TypesCalculator, err error) {
	// Check that the mandatory parameters have been provided:
	if b.reporter == nil {
		err = fmt.Errorf("reporter is mandatory")
		return
	}
	if b.packages == nil {
		err = fmt.Errorf("packages calculator is mandatory")
		return
	}
	if b.names == nil {
		err = fmt.Errorf("names is mandatory")
		return
	}

	// Create the calculator:
	calculator = &TypesCalculator{
		reporter: b.reporter,
		packages: b.packages,
		names:    b.names,
	}

	return
}

// Reference creates a new type reference with the given package and text.
func (c *TypesCalculator) Reference(imprt, selector, name, text string) *TypeReference {
	reference := &TypeReference{}
	reference.imprt = imprt
	reference.selector = selector
	reference.text = text
	return reference
}

// EnumName calculates the name of the given enumerated type.
func (c *TypesCalculator) EnumName(typ *concepts.Type) string {
	return c.EnumReference(typ).Name()
}

// EnumReference calculates a reference to the Go type used to represent the given model enumerated
// type.
func (c *TypesCalculator) EnumReference(typ *concepts.Type) *TypeReference {
	var ref *TypeReference
	if !typ.IsEnum() {
		c.reporter.Errorf(
			"Don't know how to calculate enumerated type reference for type '%s'",
			typ,
		)
		ref = &TypeReference{}
	} else {
		ref = c.ValueReference(typ)
	}
	return ref
}

// StructName calculates the name of the struct type used to represent objects of the given type.
func (c *TypesCalculator) StructName(typ *concepts.Type) string {
	return c.StructReference(typ).Name()
}

// StructReference calculates a reference to the struct type used to represent objects of the given
// type.
func (c *TypesCalculator) StructReference(typ *concepts.Type) *TypeReference {
	var ref *TypeReference
	switch {
	case typ.IsList():
		element := typ.Element()
		ref = &TypeReference{}
		ref.imprt, ref.selector = c.Package(element)
		ref.name = annotations.GoName(element)
		if ref.name == "" {
			ref.name = c.names.Public(element.Name())
		}
		ref.name += "List"
		ref.text = ref.name
	case typ.IsStruct():
		ref = &TypeReference{}
		ref.imprt, ref.selector = c.Package(typ)
		ref.name = annotations.GoName(typ)
		if ref.name == "" {
			ref.name = c.names.Public(typ.Name())
		}
		ref.text = ref.name
	default:
		c.reporter.Errorf(
			"Don't know how to calculate struct type reference for type '%s'",
			typ,
		)
		ref = &TypeReference{}
	}
	return ref
}

// ValueReference calculates a type reference to a value of the given type.
func (c *TypesCalculator) ValueReference(typ *concepts.Type) *TypeReference {
	var ref *TypeReference
	version := typ.Owner()
	switch {
	case typ == version.Boolean():
		ref = &TypeReference{}
		ref.name = "bool"
		ref.text = "bool"
	case typ == version.IntegerType():
		ref = &TypeReference{}
		ref.name = "int"
		ref.text = "int"
	case typ == version.LongType():
		ref = &TypeReference{}
		ref.name = "int64"
		ref.text = "int64"
	case typ == version.FloatType():
		ref = &TypeReference{}
		ref.name = "float64"
		ref.text = "float64"
	case typ == version.StringType():
		ref = &TypeReference{}
		ref.name = "string"
		ref.text = "string"
	case typ == version.InterfaceType():
		ref = &TypeReference{}
		ref.name = "interface{}"
		ref.text = "interface{}"
	case typ == version.DateType():
		ref = &TypeReference{}
		ref.imprt = "time"
		ref.selector = "time"
		ref.name = "Time"
		ref.text = "time.Time"
	case typ.IsEnum():
		ref = &TypeReference{}
		ref.imprt, ref.selector = c.Package(typ)
		ref.name = annotations.GoName(typ)
		if ref.name == "" {
			ref.name = c.names.Public(typ.Name())
		}
		ref.text = ref.name
	case typ.IsList():
		element := typ.Element()
		switch {
		case element.IsScalar():
			ref = c.ValueReference(element)
			ref.text = fmt.Sprintf("[]%s", ref.text)
		case element.IsStruct():
			ref = c.ValueReference(element)
			ref.text = fmt.Sprintf("[]*%s.%s", ref.selector, ref.text)
		}
	case typ.IsMap():
		element := typ.Element()
		switch {
		case element.IsScalar():
			ref = c.ValueReference(element)
			ref.text = fmt.Sprintf("map[string]%s", ref.text)
		case element.IsStruct():
			ref = c.ValueReference(element)
			ref.text = fmt.Sprintf("map[string]*%s.%s", ref.selector, ref.text)
		}
	case typ.IsStruct():
		ref = &TypeReference{}
		ref.imprt, ref.selector = c.Package(typ)
		ref.name = annotations.GoName(typ)
		if ref.name == "" {
			ref.name = c.names.Public(typ.Name())
		}
		ref.text = ref.name
	}
	if ref == nil {
		c.reporter.Errorf(
			"Don't know how to calculate value reference for type '%s'",
			typ,
		)

		ref = &TypeReference{}
	}
	return ref
}

// NullableReference calculates a type reference for a value of the given type that can be assigned
// the nil value.
func (c *TypesCalculator) NullableReference(typ *concepts.Type, referencedVersion string) *TypeReference {
	switch {
	case (typ.IsScalar() && !typ.IsInterface()):
		ref := c.ValueReference(typ)
		ref.text = fmt.Sprintf("*%s", ref.text)
		return ref
	case typ.IsStruct():
		ref := c.ValueReference(typ)
		if referencedVersion != "" {
			ref.text = fmt.Sprintf("*%s.%s", referencedVersion, ref.name)
		} else {
			ref.text = fmt.Sprintf("*%s.%s", ref.selector, ref.name)
		}
		return ref
	default:
		return c.ValueReference(typ)
	}
}

// ListName calculates the name of the type used to represent the given list type.
func (c *TypesCalculator) ListName(typ *concepts.Type) string {
	return c.ListReference(typ).Name()
}

// ListReference calculates a type reference for the given list type.
func (c *TypesCalculator) ListReference(typ *concepts.Type) *TypeReference {
	// Check that the given type is actually a list type:
	if !typ.IsList() {
		c.reporter.Errorf(
			"Don't know how to calculate list type reference for type '%s'",
			typ,
		)
		return &TypeReference{}
	}

	// Calculate the type reference:
	element := typ.Element()
	ref := &TypeReference{}
	ref.imprt, ref.selector = c.Package(element)
	ref.name = annotations.GoName(element)
	if ref.name == "" {
		ref.name = c.names.Public(element.Name())
	}
	ref.name += "List"
	ref.text = fmt.Sprintf("*%s.%s", ref.selector, ref.name)
	return ref
}

// JSONTypeReference calculates a reference for the type used to marshal and unmarshal JSON
// documents containing the given type.
func (c *TypesCalculator) JSONTypeReference(typ *concepts.Type) *TypeReference {
	var ref *TypeReference
	switch {
	case typ.IsScalar():
		ref = c.NullableReference(typ, "")
	case typ.IsList():
		element := typ.Element()
		switch {
		case element.IsScalar():
			ref = c.NullableReference(typ, "")
		case element.IsStruct():
			ref = c.JSONTypeReference(element)
			ref.text = fmt.Sprintf("[]%s", ref.text)
		}
	case typ.IsMap():
		element := typ.Element()
		switch {
		case element.IsScalar():
			ref = c.NullableReference(typ, "")
		case element.IsStruct():
			ref = c.JSONTypeReference(element)
			ref.text = fmt.Sprintf("map[string]%s", ref.text)
		}
	case typ.IsStruct():
		ref = &TypeReference{}
		ref.imprt, ref.selector = c.Package(typ)
		ref.name = c.names.Private(names.Cat(typ.Name(), nomenclator.JSON))
		ref.text = fmt.Sprintf("*%s.%s", ref.selector, ref.name)
	}
	if ref == nil {
		c.reporter.Errorf(
			"Don't know how to calculate JSON type reference for type '%s'",
			typ,
		)
		ref = &TypeReference{}
	}
	return ref
}

// JSONStructName calculates the name of the struct type used to marshal and unmarshal JSON
// documents containing the given type.
func (c *TypesCalculator) JSONStructName(typ *concepts.Type) string {
	return c.JSONStructReference(typ).Name()
}

// JSONStructReference calculates a reference for the struct type used to marshal and unmarshal JSON
// documents containing objects of the given type.
func (c *TypesCalculator) JSONStructReference(typ *concepts.Type) *TypeReference {
	var ref *TypeReference
	if typ.IsStruct() || typ.IsList() {
		ref = &TypeReference{}
		ref.imprt, ref.selector = c.Package(typ)
		ref.name = c.names.Private(names.Cat(typ.Name(), nomenclator.JSON))
		ref.text = fmt.Sprintf("*%s.%s", ref.selector, ref.name)
	}
	if ref == nil {
		c.reporter.Errorf(
			"Don't know how to calculate JSON struct type reference for type '%s'",
			typ,
		)
		ref = &TypeReference{}
	}
	return ref
}

// Builder reference calculates a reference for the type used build objects of the given type.
func (c *TypesCalculator) BuilderReference(typ *concepts.Type, refVersion string) *TypeReference {
	var ref *TypeReference
	switch {
	case typ.IsStruct():
		ref = &TypeReference{}
		ref.imprt, ref.selector = c.Package(typ)
		ref.name = annotations.GoName(typ)
		if ref.name == "" {
			ref.name = c.names.Public(typ.Name())
		}
		ref.name += "Builder"
		if refVersion == "" {
			ref.text = fmt.Sprintf("*%s.%s", ref.selector, ref.name)
		} else {
			ref.text = fmt.Sprintf("*%s.%s", refVersion, ref.name)
		}
	case typ.IsList():
		element := typ.Element()
		ref = &TypeReference{}
		ref.imprt, ref.selector = c.Package(element)
		ref.name = annotations.GoName(element)
		if ref.name == "" {
			ref.name = c.names.Public(element.Name())
		}
		ref.name += "ListBuilder"
		if refVersion == "" {
			ref.text = fmt.Sprintf("*%s.%s", ref.selector, ref.name)
		} else {
			ref.text = fmt.Sprintf("*%s.%s", refVersion, ref.name)
		}
	default:
		c.reporter.Errorf(
			"Don't know how to calculate builder reference for type '%s'",
			typ,
		)
		ref = &TypeReference{}
	}
	return ref
}

// Zero value calculates the zero value for the given type.
func (c *TypesCalculator) ZeroValue(typ *concepts.Type) string {
	version := typ.Owner()
	switch {
	case typ.IsStruct() || typ.IsList() || typ.IsMap() || typ.IsInterface():
		return `nil`
	case typ.IsEnum():
		ref := c.ValueReference(typ)
		return fmt.Sprintf(`%s("")`, ref.Name())
	case typ == version.Boolean():
		return `false`
	case typ == version.IntegerType():
		return `0`
	case typ == version.LongType():
		return `0`
	case typ == version.FloatType():
		return `0.0`
	case typ == version.DateType():
		return `time.Time{}`
	case typ == version.StringType():
		return `""`
	default:
		c.reporter.Errorf(
			"Don't know how to calculate zero value for type '%s'",
			typ.Name(),
		)
		return ""
	}
}

// Package calculates the name of the package for the given type. It returns the full import path of
// the package and the package selector.
func (c *TypesCalculator) Package(typ *concepts.Type) (imprt, selector string) {
	version := typ.Owner()
	switch {
	case typ == version.Boolean():
		return
	case typ == version.IntegerType():
		return
	case typ == version.LongType():
		return
	case typ == version.FloatType():
		return
	case typ == version.StringType():
		return
	case typ == version.InterfaceType():
		return
	case typ == version.DateType():
		imprt = "time"
		selector = "time"
		return
	case typ.IsEnum() || typ.IsStruct() || typ.IsList():
		imprt = c.packages.VersionImport(version)
		selector = path.Base(imprt)
		return
	default:
		c.reporter.Errorf(
			"Don't know how to make package name for type '%s' of kind '%s'",
			typ.Name(), typ.Kind(),
		)
		return
	}
}

// BitMask calculates the bit mask used to check presence of the given attribute.
func (c *TypesCalculator) BitMask(attribute *concepts.Attribute) string {
	// Calculate the index taking into account that the the builtin `link`, `id` and `href`
	// fields of class types:
	var index int
	typ := attribute.Owner()
	if typ.IsClass() {
		index += 3
	}
	for _, current := range attribute.Owner().Attributes() {
		if current == attribute {
			break
		}
		index++
	}

	// Calculate the mask:
	mask := 1 << index
	return fmt.Sprintf("%d", mask)
}

// BitmapType calculates the reference for the type that will be used to implement the attribute
// presence bitset for the given struct type.
func (c *TypesCalculator) BitmapType(typ *concepts.Type) *TypeReference {
	ref := &TypeReference{}

	// Check that the given type is a struct:
	if !typ.IsStruct() {
		c.reporter.Errorf(
			"Don't know how to make package name for type '%s' of kind '%s'",
			typ.Name(), typ.Kind(),
		)
		return ref
	}

	// Check that there are no more than the maximum number of attributes that can fit in the
	// bitmap:
	max := 64
	if typ.IsClass() {
		max -= 3
	}
	count := len(typ.Attributes())
	if count > max {
		c.reporter.Errorf(
			"Struct type '%s' has more %d attributes, but at most %d attributes "+
				"are supported",
			typ, count, max,
		)
		return ref
	}

	// Select the smaller usigned integer tha thas room for all the attributes:
	size := c.bitmapSize(typ)
	name := fmt.Sprintf("uint%d", size)
	ref.name = name
	ref.text = name
	return ref
}

// bitmapSize calculates the number of bits used for the presence bitmap of the given type, rounded
// up to a multiple of eight.
func (c *TypesCalculator) bitmapSize(typ *concepts.Type) int {
	// For classes the first bit is reserved to indicate if the object is a link, and the second
	// and third bits are used for the built-in `id` and `href` attributes:
	count := len(typ.Attributes())
	if typ.IsClass() {
		count += 3
	}

	// Use a normal unsigned integer if possible, or a long one if more than 32 bits are needed:
	if count <= 32 {
		return 32
	}
	return 64
}

// TypeReference represents a reference to a Go type.
type TypeReference struct {
	imprt    string
	selector string
	name     string
	text     string
}

// Import returns the import path of the package that contains the type.
func (r *TypeReference) Import() string {
	return r.imprt
}

// Selector returns the selector that should be used to import the package.
func (r *TypeReference) Selector() string {
	return r.selector
}

// Name returns the local name of the base type, without the selector.
func (r *TypeReference) Name() string {
	return r.name
}

// Text returns the text of the reference.
func (r *TypeReference) Text() string {
	return r.text
}
