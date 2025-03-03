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

	"github.com/openshift-online/ocm-api-metamodel/pkg/annotations"
	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/http"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// TypesGeneratorBuilder is an object used to configure and build the types generator. Don't create
// instances directly, use the NewTypesGenerator function instead.
type TypesGeneratorBuilder struct {
	reporter *reporter.Reporter
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	types    *TypesCalculator
	binding  *http.BindingCalculator
}

// TypesGenerator Go types for the model types. Don't create instances directly, use the builder
// instead.
type TypesGenerator struct {
	reporter *reporter.Reporter
	errors   int
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	types    *TypesCalculator
	binding  *http.BindingCalculator
	buffer   *Buffer
}

// NewTypesGenerator creates a new builder for types generators.
func NewTypesGenerator() *TypesGeneratorBuilder {
	return &TypesGeneratorBuilder{}
}

// Reporter sets the object that will be used to report information about the generation process,
// including errors.
func (b *TypesGeneratorBuilder) Reporter(value *reporter.Reporter) *TypesGeneratorBuilder {
	b.reporter = value
	return b
}

// Model sets the model that will be used by the types generator.
func (b *TypesGeneratorBuilder) Model(value *concepts.Model) *TypesGeneratorBuilder {
	b.model = value
	return b
}

// Output sets import path of the output package.
func (b *TypesGeneratorBuilder) Output(value string) *TypesGeneratorBuilder {
	b.output = value
	return b
}

// Packages sets the object that will be used to calculate package names.
func (b *TypesGeneratorBuilder) Packages(value *PackagesCalculator) *TypesGeneratorBuilder {
	b.packages = value
	return b
}

// Names sets the object that will be used to calculate names.
func (b *TypesGeneratorBuilder) Names(value *NamesCalculator) *TypesGeneratorBuilder {
	b.names = value
	return b
}

// Types sets the object that will be used to calculate types.
func (b *TypesGeneratorBuilder) Types(value *TypesCalculator) *TypesGeneratorBuilder {
	b.types = value
	return b
}

// Binding sets the object that will be used to calculate HTTP names.
func (b *TypesGeneratorBuilder) Binding(value *http.BindingCalculator) *TypesGeneratorBuilder {
	b.binding = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new
// types generator using it.
func (b *TypesGeneratorBuilder) Build() (generator *TypesGenerator, err error) {
	// Check that the mandatory parameters have been provided:
	if b.reporter == nil {
		err = fmt.Errorf("reporter is mandatory")
		return
	}
	if b.model == nil {
		err = fmt.Errorf("model is mandatory")
		return
	}
	if b.output == "" {
		err = fmt.Errorf("output is mandatory")
		return
	}
	if b.packages == nil {
		err = fmt.Errorf("packages calculator is mandatory")
		return
	}
	if b.names == nil {
		err = fmt.Errorf("names calculator is mandatory")
		return
	}
	if b.types == nil {
		err = fmt.Errorf("types calculator is mandatory")
		return
	}
	if b.binding == nil {
		err = fmt.Errorf("binding calculator is mandatory")
		return
	}

	// Create the generator:
	generator = &TypesGenerator{
		reporter: b.reporter,
		model:    b.model,
		output:   b.output,
		packages: b.packages,
		names:    b.names,
		types:    b.types,
		binding:  b.binding,
	}

	return
}

// Run executes the code generator.
func (g *TypesGenerator) Run() error {
	var err error

	// Generate the go types:
	for _, service := range g.model.Services() {
		for _, version := range service.Versions() {
			// Generate the version metadata type:
			err := g.generateVersionMetadataTypeFile(version)
			if err != nil {
				return err
			}

			// Generate the Go types that correspond to model types:
			for _, typ := range version.Types() {
				switch {
				case typ.IsEnum() || typ.IsStruct():
					err = g.generateTypeFile(typ)
				}
				if err != nil {
					return err
				}
			}
		}
	}

	// Check if there were errors:
	if g.errors > 0 {
		if g.errors > 1 {
			err = fmt.Errorf("there were %d errors", g.errors)
		} else {
			err = fmt.Errorf("there was 1 error")
		}
		return err
	}

	return nil
}

func (g *TypesGenerator) generateVersionMetadataTypeFile(version *concepts.Version) error {
	var err error

	// Calculate the package and file name:
	pkgName := g.packages.VersionPackage(version)
	fileName := g.metadataFile()

	// Create the buffer for the generated code:
	g.buffer, err = NewBuffer().
		Reporter(g.reporter).
		Output(g.output).
		Packages(g.packages).
		Package(pkgName).
		File(fileName).
		Build()
	if err != nil {
		return err
	}

	// Generate the source:
	g.generateVersionMetadataTypeSource(version)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *TypesGenerator) generateVersionMetadataTypeSource(version *concepts.Version) {
	g.buffer.Emit(`
		// Metadata contains the version metadata.
		type Metadata struct {
			bitmap_       uint32
			serverVersion string
		}

		// ServerVersion returns the version of the server.
		func (m *Metadata) ServerVersion() string {
			if m != nil && m.bitmap_&1 != 0 {
				return m.serverVersion
			}
			return ""
		}

		// GetServerVersion returns the value of the server version and a flag indicating if
		// the attribute has a value.
		func (m *Metadata) GetServerVersion() (value string, ok bool) {
			ok = m != nil && m.bitmap_&1 != 0
			if ok {
				value = m.serverVersion
			}
			return
		}
		`,
	)
}

func (g *TypesGenerator) generateTypeFile(typ *concepts.Type) error {
	var err error

	// Calculate the package and file name:
	pkgName := g.packages.VersionPackage(typ.Owner())
	fileName := g.typeFile(typ)

	// Create the buffer for the generated code:
	g.buffer, err = NewBuffer().
		Reporter(g.reporter).
		Output(g.output).
		Packages(g.packages).
		Package(pkgName).
		File(fileName).
		Function("bitMask", g.types.BitMask).
		Function("bitmapType", g.types.BitmapType).
		Function("enumName", g.types.EnumName).
		Function("fieldName", g.fieldName).
		Function("fieldType", g.fieldType).
		Function("getterName", g.getterName).
		Function("getterType", g.getterType).
		Function("listName", g.listName).
		Function("objectName", g.objectName).
		Function("valueName", g.valueName).
		Function("valueTag", g.valueTag).
		Function("zeroValue", g.types.ZeroValue).
		Build()
	if err != nil {
		return err
	}

	// Generate the source:
	g.generateTypeSource(typ)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *TypesGenerator) generateTypeSource(typ *concepts.Type) {
	switch {
	case typ.IsEnum():
		g.generateEnumTypeSource(typ)
	case typ.IsStruct():
		g.generateStructTypeSource(typ)
	}
}

func (g *TypesGenerator) generateEnumTypeSource(typ *concepts.Type) {
	g.buffer.Emit(`
		{{ $enumName := enumName .Type }}

		// {{ $enumName }} represents the values of the '{{ .Type.Name }}' enumerated type.
		type {{ $enumName }} string

		const (
			{{ range .Type.Values }}
				{{ lineComment .Doc }}
				{{ valueName . }} {{ $enumName }} = "{{ valueTag . }}"
			{{ end }}
		)
		`,
		"Type", typ,
	)
}

func (g *TypesGenerator) generateStructTypeSource(typ *concepts.Type) {
	g.buffer.Import("time", "")
	g.buffer.Emit(`
		{{ $objectName := objectName .Type }}
		{{ $listName := listName .Type }}

		{{ if .Type.IsClass }}
			// {{ $objectName }}Kind is the name of the type used to represent objects
			// of type '{{ .Type.Name }}'.
			const {{ $objectName }}Kind = "{{ $objectName }}"

			// {{ $objectName }}LinkKind is the name of the type used to represent links
			// to objects of type '{{ .Type.Name }}'.
			const {{ $objectName }}LinkKind = "{{ $objectName }}Link"

			// {{ $objectName }}NilKind is the name of the type used to nil references
			// to objects of type '{{ .Type.Name }}'.
			const {{ $objectName }}NilKind = "{{ $objectName }}Nil"
		{{ end }}

		// {{ $objectName }} represents the values of the '{{ .Type.Name }}' type.
		//
		{{ lineComment .Type.Doc }}
		type  {{ $objectName }} struct {
			bitmap_ {{ bitmapType .Type }}
			{{ if .Type.IsClass }}
				id   string
				href string
			{{ end }}
			{{ range .Type.Attributes }}
				{{ if not .Type.IsBoolean }}
					{{ fieldName . }} {{ fieldType . }}
				{{ end }}
			{{ end }}
			{{ range .Type.Attributes }}
				{{ if .Type.IsBoolean }}
					{{ fieldName . }} {{ fieldType . }}
				{{ end }}
			{{ end }}
		}

		{{ if .Type.IsClass }}
			// Kind returns the name of the type of the object.
			func (o *{{ $objectName }}) Kind() string {
				if o == nil {
					return {{ $objectName }}NilKind
				}
				if o.bitmap_&1 != 0 {
					return {{ $objectName }}LinkKind
				}
				return {{ $objectName }}Kind
			}

			// Link returns true if this is a link.
			func (o *{{ $objectName }}) Link() bool {
				return o != nil && o.bitmap_&1 != 0
			}

			// ID returns the identifier of the object.
			func (o *{{ $objectName }}) ID() string {
				if o != nil && o.bitmap_&2 != 0 {
					return o.id
				}
				return ""
			}

			// GetID returns the identifier of the object and a flag indicating if the
			// identifier has a value.
			func (o *{{ $objectName }}) GetID() (value string, ok bool) {
				ok = o != nil && o.bitmap_&2 != 0
				if ok {
					value = o.id
				}
				return
			}

			// HREF returns the link to the object.
			func (o *{{ $objectName }}) HREF() string {
				if o != nil && o.bitmap_&4 != 0 {
					return o.href
				}
				return ""
			}

			// GetHREF returns the link of the object and a flag indicating if the
			// link has a value.
			func (o *{{ $objectName }}) GetHREF() (value string, ok bool) {
				ok = o != nil && o.bitmap_&4 != 0
				if ok {
					value = o.href
				}
				return
			}
		{{ end }}

		// Empty returns true if the object is empty, i.e. no attribute has a value.
		func (o *{{ $objectName }}) Empty() bool {
			{{ if .Type.IsClass }}
				return o == nil || o.bitmap_&^1 == 0
			{{ else }}
				return o == nil || o.bitmap_ == 0
			{{ end }}
		}

		{{ range .Type.Attributes }}
			{{ $attributeType := .Type.Name.String }}
			{{ $fieldName := fieldName . }}
			{{ $fieldMask := bitMask . }}
			{{ $getterName := getterName . }}
			{{ $getterType := getterType . }}

			// {{ $getterName }} returns the value of the '{{ .Name }}' attribute, or
			// the zero value of the type if the attribute doesn't have a value.
			//
			{{ lineComment .Doc }}
			func (o *{{ $objectName }}) {{ $getterName }}() {{ $getterType }} {
				if o != nil && o.bitmap_&{{ $fieldMask }} != 0 {
					return o.{{ $fieldName }}
				}
				return {{ zeroValue .Type }}
			}

			// Get{{ $getterName }} returns the value of the '{{ .Name }}' attribute and
			// a flag indicating if the attribute has a value.
			//
			{{ lineComment .Doc }}
			func (o *{{ $objectName }}) Get{{ $getterName }}() (value {{ $getterType }}, ok bool) {
				ok = o != nil && o.bitmap_&{{ $fieldMask }} != 0
				if ok {
					value = o.{{ $fieldName }}
				}
				return
			}
		{{ end }}

		// {{ $listName }}Kind is the name of the type used to represent list of objects of
		// type '{{ .Type.Name }}'.
		const {{ $listName }}Kind = "{{ $listName }}"

		// {{ $listName }}LinkKind is the name of the type used to represent links to list
		// of objects of type '{{ .Type.Name }}'.
		const {{ $listName }}LinkKind = "{{ $listName }}Link"

		// {{ $objectName }}NilKind is the name of the type used to nil lists of objects of
		// type '{{ .Type.Name }}'.
		const {{ $listName }}NilKind = "{{ $listName }}Nil"

		// {{ $listName }} is a list of values of the '{{ .Type.Name }}' type.
		type {{ $listName }} struct {
			href  string
			link  bool
			items []*{{ $objectName }}
		}

		{{ if .Type.IsClass }}
			// Kind returns the name of the type of the object.
			func (l *{{ $listName }}) Kind() string {
				if l == nil {
					return {{ $listName }}NilKind
				}
				if l.link {
					return {{ $listName }}LinkKind
				}
				return {{ $listName }}Kind
			}

			// Link returns true iif this is a link.
			func (l *{{ $listName }}) Link() bool {
				return l != nil && l.link
			}

			// HREF returns the link to the list.
			func (l *{{ $listName }}) HREF() string {
				if l != nil {
					return l.href
				}
				return ""
			}

			// GetHREF returns the link of the list and a flag indicating if the
			// link has a value.
			func (l *{{ $listName }}) GetHREF() (value string, ok bool) {
				ok = l != nil && l.href != ""
				if ok {
					value = l.href
				}
				return
			}
		{{ end }}

		// Len returns the length of the list.
		func (l *{{ $listName }}) Len() int {
			if l == nil {
				return 0
			}
			return len(l.items)
		}

		// Items sets the items of the list.
		func (l *{{ $listName }}) SetLink(link bool) {
			l.link = link
		}

		// Items sets the items of the list.
		func (l *{{ $listName }}) SetHREF(href string) {
			l.href = href
		}

		// Items sets the items of the list.
		func (l *{{ $listName }}) SetItems(items []*{{ $objectName }}) {
			l.items = items
		}

		// Items returns the items of the list.
		func (l *{{ $listName }}) Items() []*{{ $objectName }} {
			if l == nil {
				return nil
			}
			return l.items
		}

		// Empty returns true if the list is empty.
		func (l *{{ $listName }}) Empty() bool {
			return l == nil || len(l.items) == 0
		}

		// Get returns the item of the list with the given index. If there is no item with
		// that index it returns nil.
		func (l *{{ $listName }}) Get(i int) *{{ $objectName }} {
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
		func (l *{{ $listName }}) Slice() []*{{ $objectName }} {
			var slice []*{{ $objectName }}
			if l == nil {
				slice = make([]*{{ $objectName}}, 0)
			} else {
				slice = make([]*{{ $objectName}}, len(l.items))
				copy(slice, l.items)
			}
			return slice
		}

		// Each runs the given function for each item of the list, in order. If the function
		// returns false the iteration stops, otherwise it continues till all the elements
		// of the list have been processed.
		func (l *{{ $listName }}) Each(f func(item *{{ $objectName }}) bool) {
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
		func (l *{{ $listName }}) Range(f func(index int, item *{{ $objectName }}) bool) {
			if l == nil {
				return
			}
			for index, item := range l.items {
				if !f(index, item) {
					break
				}
			}
		}
		`,
		"Type", typ,
	)
}

func (g *TypesGenerator) metadataFile() string {
	return g.names.File(names.Cat(nomenclator.Metadata, nomenclator.Type))
}

func (g *TypesGenerator) typeFile(typ *concepts.Type) string {
	return g.names.File(names.Cat(typ.Name(), nomenclator.Type))
}

func (g *TypesGenerator) fieldName(attribute *concepts.Attribute) string {
	return g.names.Private(attribute.Name())
}

func (g *TypesGenerator) getterType(attribute *concepts.Attribute) *TypeReference {
	var ref *TypeReference
	typ := attribute.Type()
	referencedVersion := ""
	if attribute.LinkOwner() != nil {
		referencedVersion = attribute.LinkOwner().Name().String()
	}
	switch {
	case typ.IsScalar():
		ref = g.types.ValueReference(typ)
	case typ.IsStruct():
		ref = g.types.NullableReference(typ, referencedVersion)
	case typ.IsList():
		if attribute.Link() {
			ref = g.types.ListReference(typ)
		} else {
			ref = g.types.NullableReference(typ, referencedVersion)
		}
	case typ.IsMap():
		ref = g.types.NullableReference(typ, referencedVersion)
	}
	if ref == nil {
		g.reporter.Errorf(
			"Don't know how to calculate getter type for attribute '%s'",
			attribute,
		)
		ref = &TypeReference{}
	}
	return ref
}

func (g *TypesGenerator) objectName(typ *concepts.Type) string {
	name := annotations.GoName(typ)
	if name == "" {
		name = g.names.Public(typ.Name())
	}
	return name
}

func (g *TypesGenerator) valueName(value *concepts.EnumValue) string {
	typeName := annotations.GoName(value.Type())
	if typeName == "" {
		typeName = g.names.Public(value.Type().Name())
	}
	valueName := annotations.GoName(value)
	if valueName == "" {
		valueName = g.names.Public(value.Name())
	}
	return typeName + valueName
}

func (g *TypesGenerator) valueTag(value *concepts.EnumValue) string {
	return g.binding.EnumValueName(value)
}

func (g *TypesGenerator) getterName(attribute *concepts.Attribute) string {
	name := annotations.GoName(attribute)
	if name == "" {
		name = g.names.Public(attribute.Name())
	}
	return name
}

func (g *TypesGenerator) fieldType(attribute *concepts.Attribute) *TypeReference {
	var ref *TypeReference
	typ := attribute.Type()
	referencedVersion := ""
	if attribute.LinkOwner() != nil {
		referencedVersion = attribute.LinkOwner().Name().String()
	}
	switch {
	case typ.IsScalar():
		ref = g.types.ValueReference(typ)
	case typ.IsStruct():
		ref = g.types.NullableReference(typ, referencedVersion)
	case typ.IsList():
		if attribute.Link() {
			ref = g.types.ListReference(typ)
		} else {
			ref = g.types.NullableReference(typ, referencedVersion)
		}
	case typ.IsMap():
		ref = g.types.NullableReference(typ, referencedVersion)
	}
	if ref == nil {
		g.reporter.Errorf(
			"Don't know how to calculate field type for attribute '%s'",
			attribute,
		)
		ref = &TypeReference{}
	}
	return ref
}

func (g *TypesGenerator) listName(typ *concepts.Type) string {
	typeName := annotations.GoName(typ)
	if typeName == "" {
		typeName = g.names.Public(typ.Name())
	}
	return typeName + "List"
}
