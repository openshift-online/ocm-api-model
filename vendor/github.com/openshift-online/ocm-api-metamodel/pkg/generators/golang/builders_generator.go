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
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// BuildersGeneratorBuilder is an object used to configure and build the builders generator. Don't
// create instances directly, use the NewBuildersGenerator function instead.
type BuildersGeneratorBuilder struct {
	reporter *reporter.Reporter
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	types    *TypesCalculator
}

// BuildersGenerator generates code for the builders of the model types. Don't create instances
// directly, use the builder instead.
type BuildersGenerator struct {
	reporter *reporter.Reporter
	errors   int
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	types    *TypesCalculator
	buffer   *Buffer
}

// NewBuildersGenerator creates a new builder for builders generators.
func NewBuildersGenerator() *BuildersGeneratorBuilder {
	return &BuildersGeneratorBuilder{}
}

// Reporter sets the object that will be used to report information about the generation process,
// including errors.
func (b *BuildersGeneratorBuilder) Reporter(value *reporter.Reporter) *BuildersGeneratorBuilder {
	b.reporter = value
	return b
}

// Model sets the model that will be used by the types generator.
func (b *BuildersGeneratorBuilder) Model(value *concepts.Model) *BuildersGeneratorBuilder {
	b.model = value
	return b
}

// Output sets the directory where the source will be generated.
func (b *BuildersGeneratorBuilder) Output(value string) *BuildersGeneratorBuilder {
	b.output = value
	return b
}

// Packages sets the object that will be used to calculate package names.
func (b *BuildersGeneratorBuilder) Packages(
	value *PackagesCalculator) *BuildersGeneratorBuilder {
	b.packages = value
	return b
}

// Names sets the object that will be used to calculate names.
func (b *BuildersGeneratorBuilder) Names(value *NamesCalculator) *BuildersGeneratorBuilder {
	b.names = value
	return b
}

// Types sets the object that will be used to calculate types.
func (b *BuildersGeneratorBuilder) Types(value *TypesCalculator) *BuildersGeneratorBuilder {
	b.types = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new
// builders generator using it.
func (b *BuildersGeneratorBuilder) Build() (generator *BuildersGenerator, err error) {
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
		err = fmt.Errorf("types is mandatory")
		return
	}

	// Create the generator:
	generator = &BuildersGenerator{
		reporter: b.reporter,
		model:    b.model,
		output:   b.output,
		packages: b.packages,
		names:    b.names,
		types:    b.types,
	}

	return
}

// Run executes the code generator.
func (g *BuildersGenerator) Run() error {
	var err error

	// Generate the code for each type:
	for _, service := range g.model.Services() {
		for _, version := range service.Versions() {
			for _, typ := range version.Types() {
				switch {
				case typ.IsStruct():
					err = g.generateStructBuilderFile(typ)
				case typ.IsList() && typ.Element().IsStruct():
					err = g.generateListBuilderFile(typ)
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

func (g *BuildersGenerator) generateStructBuilderFile(typ *concepts.Type) error {
	var err error

	// Calculate the package and file names:
	pkgName := g.packages.VersionPackage(typ.Owner())
	fileName := g.fileName(typ)

	// Create the buffer for the generated code:
	g.buffer, err = NewBuffer().
		Reporter(g.reporter).
		Output(g.output).
		Packages(g.packages).
		Package(pkgName).
		File(fileName).
		Function("bitMask", g.types.BitMask).
		Function("bitmapType", g.types.BitmapType).
		Function("builderCtor", g.builderCtor).
		Function("builderName", g.builderName).
		Function("fieldName", g.fieldName).
		Function("fieldType", g.fieldType).
		Function("selectorType", g.selectorType).
		Function("objectName", g.objectName).
		Function("setterName", g.setterName).
		Function("setterType", g.setterType).
		Function("valueType", g.valueType).
		Build()
	if err != nil {
		return err
	}

	// Generate the code:
	g.generateStructBuilderSource(typ)

	// Write the generated code:
	return g.buffer.Write()

}

func (g *BuildersGenerator) generateStructBuilderSource(typ *concepts.Type) {
	g.buffer.Emit(`
		{{ $builderName := builderName .Type }}
		{{ $builderCtor := builderCtor .Type }}
		{{ $objectName := objectName .Type }}

		// {{ $builderName }} contains the data and logic needed to build '{{ .Type.Name }}' objects.
		//
		{{ lineComment .Type.Doc }}
		type  {{ $builderName }} struct {
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

		// {{ $builderCtor }} creates a new builder of '{{ .Type.Name }}' objects.
		func {{ $builderCtor }}() *{{ $builderName }} {
			return &{{ $builderName }}{}
		}

		{{ if .Type.IsClass }}
			// Link sets the flag that indicates if this is a link.
			func (b *{{ $builderName }}) Link(value bool) *{{ $builderName }} {
				b.bitmap_ |= 1
				return b
			}

			// ID sets the identifier of the object.
			func (b *{{ $builderName }}) ID(value string) *{{ $builderName }} {
				b.id = value
				b.bitmap_ |= 2
				return b
			}

			// HREF sets the link to the object.
			func (b *{{ $builderName }}) HREF(value string) *{{ $builderName }} {
				b.href = value
				b.bitmap_ |= 4
				return b
			}
		{{ end }}

		// Empty returns true if the builder is empty, i.e. no attribute has a value.
		func (b *{{ $builderName }}) Empty() bool {
			{{ if .Type.IsClass }}
				return b == nil || b.bitmap_&^1 == 0
			{{ else }}
				return b == nil || b.bitmap_ == 0
			{{ end }}
		}

		{{ range .Type.Attributes }}
			{{ $fieldName := fieldName . }}
			{{ $setterName := setterName . }}
			{{ $setterType := setterType . }}
			{{ $fieldMask := bitMask . }}
			{{ $selectorType := selectorType . }}

			{{ if .Type.IsList }}
				// {{ $setterName }} sets the value of the '{{ .Name }}' attribute to the given values.
				//
				{{ lineComment .Type.Doc }}
				{{ if .Link }}
					func (b *{{ $builderName }}) {{ $setterName }}(value {{ $setterType }}) *{{ $builderName }} {
						b.{{ $fieldName }} = value
						b.bitmap_ |= {{ $fieldMask }}
						return b
					}
				{{ else }}
					{{ $elementType := valueType .Type.Element }}
					{{ if .Type.Element.IsScalar }}
						func (b *{{ $builderName }}) {{ $setterName }}(values ...{{ $elementType }}) *{{ $builderName }} {
							b.{{ $fieldName }} = make([]{{ $elementType }}, len(values))
							copy(b.{{ $fieldName }}, values)
							b.bitmap_ |= {{ $fieldMask }}
							return b
						}
					{{ else }}
						{{ $elementBuilderName := builderName .Type.Element }}
						func (b *{{ $builderName }}) {{ $setterName }}(values ...*{{ selectorType . }}{{ $elementBuilderName }}) *{{ $builderName }} {
							b.{{ $fieldName }} = make([]*{{ selectorType . }}{{ $elementBuilderName }}, len(values))
							copy(b.{{ $fieldName }}, values)
							b.bitmap_ |= {{ $fieldMask }}
							return b
						}
					{{ end }}
				{{ end }}
			{{ else }}
				// {{ $setterName }} sets the value of the '{{ .Name }}' attribute to the given value.
				//
				{{ lineComment .Type.Doc }}
				func (b *{{ $builderName }}) {{ $setterName }}(value {{ $setterType }}) *{{ $builderName }} {
					b.{{ $fieldName }} = value
					{{ if .Type.IsScalar }}
						b.bitmap_ |= {{ $fieldMask }}
					{{ else }}
						if value != nil {
							b.bitmap_ |= {{ $fieldMask }}
						} else {
							b.bitmap_ &^= {{ $fieldMask }}
						}
					{{ end }}
					return b
				}
			{{ end }}
		{{ end }}

		// Copy copies the attributes of the given object into this builder, discarding any previous values.
		func (b *{{ $builderName }}) Copy(object *{{ $objectName }}) *{{ $builderName }} {
			if object == nil {
				return b
			}
			b.bitmap_ = object.bitmap_
			{{ if .Type.IsClass }}
				b.id = object.id
				b.href = object.href
			{{ end }}
			{{ range .Type.Attributes }}
				{{ $fieldName := fieldName . }}
				{{ $fieldType := fieldType . }}
				{{ $selectorType := selectorType . }}
				{{ if .Type.IsScalar }}
					b.{{ $fieldName }} = object.{{ $fieldName }}
				{{ else if .Type.IsStruct }}
					if object.{{ $fieldName }} != nil {
						b.{{ $fieldName }} = {{ selectorType . }}{{ builderCtor .Type }}().Copy(object.{{ $fieldName }})
					} else {
						b.{{ $fieldName }} = nil
					}
				{{ else if .Type.IsList }}
					if object.{{ $fieldName }} != nil {
						{{ if .Link }}
							b.{{ $fieldName }} = {{ selectorType . }}{{ builderCtor .Type }}().Copy(object.{{ $fieldName }})
						{{ else }}
							{{ if .Type.Element.IsScalar }}
								b.{{ $fieldName }} = make({{ $fieldType }}, len(object.{{ $fieldName }}))
								copy(b.{{ $fieldName }}, object.{{ $fieldName }})
							{{ else if .Type.Element.IsStruct }}
								b.{{ $fieldName }} = make([]*{{ selectorType . }}{{ builderName .Type.Element }}, len(object.{{ $fieldName }}))
								for i, v := range object.{{ $fieldName }} {
									b.{{ $fieldName }}[i] = {{ selectorType . }}{{ builderCtor .Type.Element }}().Copy(v)
								}
							{{ end }}
						{{ end }}
					} else {
						b.{{ $fieldName }} = nil
					}
				{{ else if .Type.IsMap }}
					if len(object.{{ $fieldName }}) > 0 {
						{{ if .Type.Element.IsScalar }}
							b.{{ $fieldName }} = {{ $fieldType }}{}
							for k, v := range object.{{ $fieldName }} {
								b.{{ $fieldName }}[k] = v
							}
						{{ else if .Type.Element.IsStruct }}
							b.{{ $fieldName }} = map[string]*{{ selectorType . }}{{ builderName .Type.Element }}{}
							for k, v := range object.{{ $fieldName }} {
								b.{{ $fieldName }}[k] = {{ selectorType . }}{{ builderCtor .Type.Element }}().Copy(v)
							}
						{{ end }}
					} else {
						b.{{ $fieldName }} = nil
					}
				{{ end }}
			{{ end }}
			return b
		}

		// Build creates a '{{ .Type.Name }}' object using the configuration stored in the builder.
		func (b *{{ $builderName }}) Build() (object *{{ $objectName }}, err error) {
			object = new({{ $objectName }})
			{{ if .Type.IsClass }}
				object.id = b.id
				object.href = b.href
			{{ end }}
			object.bitmap_ = b.bitmap_
			{{ range .Type.Attributes }}
				{{ $fieldName := fieldName . }}
				{{ $fieldType := fieldType . }}
				{{ if .Type.IsScalar }}
					object.{{ $fieldName }} = b.{{ $fieldName }}
				{{ else if .Type.IsStruct }}
					if b.{{ $fieldName }} != nil {
						object.{{ $fieldName }}, err = b.{{ $fieldName }}.Build()
						if err != nil {
							return
						}
					}
				{{ else if .Type.IsList}}
					if b.{{ $fieldName }} != nil {
						{{ if .Link }}
							object.{{ $fieldName }}, err = b.{{ $fieldName }}.Build()
							if err != nil {
								return
							}
						{{ else }}
							{{ if .Type.Element.IsScalar }}
								object.{{ $fieldName }} = make({{ $fieldType }}, len(b.{{ $fieldName }}))
								copy(object.{{ $fieldName }}, b.{{ $fieldName }})
							{{ else if .Type.Element.IsStruct }}
								object.{{ $fieldName }} = make([]*{{ selectorType . }}{{ objectName .Type.Element }}, len(b.{{ $fieldName }}))
								for i, v := range b.{{ $fieldName }} {
									object.{{ $fieldName }}[i], err = v.Build()
									if err != nil {
										return
									}
								}
							{{ end }}
						{{ end }}
					}
				{{ else if .Type.IsMap}}
					if b.{{ $fieldName }} != nil {
						{{ if .Type.Element.IsScalar }}
							object.{{ $fieldName }} = make({{ $fieldType }})
							for k, v := range b.{{ $fieldName }} {
								object.{{ $fieldName }}[k] = v
							}
						{{ else if .Type.Element.IsStruct }}
							object.{{ $fieldName }}  = make(map[string]*{{ selectorType . }}{{ objectName .Type.Element }})
							for k, v := range b.{{ $fieldName }} {
								object.{{ $fieldName }}[k], err = v.Build()
								if err != nil {
									return
								}
							}
						{{ end }}
					}
				{{ end }}
			{{ end }}
			return
		}
		`,
		"Type", typ,
	)
}

func (g *BuildersGenerator) generateListBuilderFile(typ *concepts.Type) error {
	var err error

	// Calculate the package and file name:
	pkgName := g.packages.VersionPackage(typ.Owner())
	fileName := g.fileName(typ)

	// Create the buffer for the generated code:
	g.buffer, err = NewBuffer().
		Reporter(g.reporter).
		Output(g.output).
		Packages(g.packages).
		Package(pkgName).
		File(fileName).
		Function("builderCtor", g.builderCtor).
		Function("builderName", g.builderName).
		Function("objectName", g.objectName).
		Build()
	if err != nil {
		return err
	}

	// Generate the code:
	g.generateListBuilderSource(typ)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *BuildersGenerator) generateListBuilderSource(typ *concepts.Type) {
	if typ.Element().IsStruct() {
		g.generateStructListBuilderSource(typ)
	}
}

func (g *BuildersGenerator) generateStructListBuilderSource(typ *concepts.Type) {
	g.buffer.Emit(`
		{{ $objectName := objectName .Type }}
		{{ $builderName := builderName .Type }}
		{{ $builderCtor := builderCtor .Type }}
		{{ $elementObjectName := objectName .Type.Element }}
		{{ $elementBuilderName := builderName .Type.Element }}
		{{ $elementBuilderCtor := builderCtor .Type.Element }}

		// {{ $builderName }} contains the data and logic needed to build
		// '{{ .Type.Element.Name }}' objects.
		type  {{ $builderName }} struct {
			items []*{{ $elementBuilderName }}
		}

		// {{ $builderCtor }} creates a new builder of '{{ .Type.Element.Name }}' objects.
		func {{ $builderCtor }}() *{{ $builderName }} {
			return new({{ $builderName }})
		}

		// Items sets the items of the list.
		func (b *{{ $builderName }}) Items(values ...*{{ $elementBuilderName }}) *{{ $builderName }} {
                        b.items = make([]*{{ $elementBuilderName }}, len(values))
                        copy(b.items, values)
			return b
		}

		// Empty returns true if the list is empty.
		func (b *{{ $builderName }}) Empty() bool {
			return b == nil || len(b.items) == 0
		}

		// Copy copies the items of the given list into this builder, discarding any previous items.
		func (b *{{ $builderName }}) Copy(list *{{ $objectName }}) *{{ $builderName }} {
			if list == nil || list.items == nil {
				b.items = nil
			} else {
				b.items = make([]*{{ $elementBuilderName }}, len(list.items))
				for i, v := range list.items {
					b.items[i] = {{ $elementBuilderCtor }}().Copy(v)
				}
			}
			return b
		}

		// Build creates a list of '{{ .Type.Element.Name }}' objects using the
		// configuration stored in the builder.
		func (b *{{ $builderName }}) Build() (list *{{ $objectName }}, err error) {
			items := make([]*{{ $elementObjectName }}, len(b.items))
			for i, item := range b.items {
				items[i], err = item.Build()
				if err != nil {
					return
				}
			}
			list = new({{ $objectName }})
			list.items = items
			return
		}
		`,
		"Type", typ,
	)
}

func (g *BuildersGenerator) fileName(typ *concepts.Type) string {
	return g.names.File(names.Cat(typ.Name(), nomenclator.Builder))
}

func (g *BuildersGenerator) objectName(typ *concepts.Type) string {
	var name string
	switch {
	case typ.IsStruct():
		name = annotations.GoName(typ)
		if name == "" {
			name = g.names.Public(typ.Name())
		}
	case typ.IsList():
		element := typ.Element()
		name = annotations.GoName(element)
		if name == "" {
			name = g.names.Public(element.Name())
		}
		name += "List"
	default:
		g.reporter.Errorf(
			"Don't know how to calculate object type name for type '%s'",
			typ,
		)
	}
	return name
}

func (g *BuildersGenerator) builderName(typ *concepts.Type) string {
	var name string
	switch {
	case typ.IsStruct():
		name = annotations.GoName(typ)
		if name == "" {
			name = g.names.Public(typ.Name())
		}
		name += "Builder"
	case typ.IsList():
		element := typ.Element()
		name = annotations.GoName(element)
		if name == "" {
			name = g.names.Public(element.Name())
		}
		name += "ListBuilder"
	default:
		g.reporter.Errorf(
			"Don't know how to calculate builder type name for type '%s'",
			typ,
		)
		return ""

	}
	return name
}

func (g *BuildersGenerator) builderCtor(typ *concepts.Type) string {
	var name string
	switch {
	case typ.IsList():
		element := typ.Element()
		name = annotations.GoName(element)
		if name == "" {
			name = g.names.Public(element.Name())
		}
		name += "List"
	case typ.IsStruct():
		name = annotations.GoName(typ)
		if name == "" {
			name = g.names.Public(typ.Name())
		}
	default:
		g.reporter.Errorf(
			"Don't know how to calculate builder constructor name for type '%s'",
			typ,
		)
		return ""
	}
	name = "New" + name
	return name
}

func (g *BuildersGenerator) fieldName(attribute *concepts.Attribute) string {
	return g.names.Private(attribute.Name())
}

func (g *BuildersGenerator) fieldType(attribute *concepts.Attribute) *TypeReference {
	typ := attribute.Type()
	var ref *TypeReference
	referencedVersion := ""
	if attribute.LinkOwner() != nil {
		referencedVersion = attribute.LinkOwner().Name().String()
	}
	switch {
	case typ.IsScalar():
		ref = g.types.ValueReference(typ)
	case typ.IsStruct():
		ref = g.types.BuilderReference(typ, referencedVersion)
	case typ.IsList():
		if attribute.Link() {
			ref = g.types.BuilderReference(typ, referencedVersion)
		} else {
			element := typ.Element()
			if element.IsScalar() {
				ref = g.types.NullableReference(typ, referencedVersion)
			} else {
				ref = g.types.BuilderReference(element, referencedVersion)
				ref = g.types.Reference(
					ref.Import(),
					ref.Selector(),
					ref.Name(),
					fmt.Sprintf("[]%s", ref.Text()),
				)
			}
		}
	case typ.IsMap():
		element := typ.Element()
		if element.IsScalar() {
			ref = g.types.NullableReference(typ, referencedVersion)
		} else {
			ref = g.types.BuilderReference(element, referencedVersion)
			ref = g.types.Reference(
				ref.Import(),
				ref.Selector(),
				ref.Name(),
				fmt.Sprintf("map[string]%s", ref.Text()),
			)
		}
	}
	if ref == nil {
		g.reporter.Errorf(
			"Don't know how to calculate builder field type for attribute '%s'",
			attribute,
		)
		ref = &TypeReference{}
	}
	return ref
}

func (g *BuildersGenerator) selectorType(attribute *concepts.Attribute) string {
	if attribute.LinkOwner() == nil {
		return ""
	}
	return fmt.Sprintf("%s.", g.packages.VersionSelector(attribute.LinkOwner()))
}

func (g *BuildersGenerator) setterName(attribute *concepts.Attribute) string {
	name := annotations.GoName(attribute)
	if name == "" {
		name = g.names.Public(attribute.Name())
	}
	return name
}

func (g *BuildersGenerator) setterType(attribute *concepts.Attribute) *TypeReference {
	typ := attribute.Type()
	var ref *TypeReference
	referencedVersion := ""
	if attribute.LinkOwner() != nil {
		referencedVersion = attribute.LinkOwner().Name().String()
	}
	switch {
	case typ.IsScalar():
		ref = g.types.ValueReference(typ)
	case typ.IsStruct():
		ref = g.types.BuilderReference(typ, referencedVersion)
	case typ.IsList():
		if attribute.Link() {
			ref = g.types.BuilderReference(typ, referencedVersion)
		} else {
			element := typ.Element()
			if element.IsScalar() {
				ref = g.types.ValueReference(typ)
			} else {
				ref = g.types.BuilderReference(element, referencedVersion)
				ref = g.types.Reference(
					ref.Import(),
					ref.Selector(),
					ref.Name(),
					fmt.Sprintf("[]%s", ref.Text()),
				)
			}
		}
	case typ.IsMap():
		element := typ.Element()
		if element.IsScalar() {
			ref = g.types.ValueReference(typ)
		} else {
			ref = g.types.BuilderReference(element, referencedVersion)
			ref = g.types.Reference(
				ref.Import(),
				ref.Selector(),
				ref.Name(),
				fmt.Sprintf("map[string]%s", ref.Text()),
			)
		}
	}
	if ref == nil {
		g.reporter.Errorf(
			"Don't know how to calculate builder setter type for attribute '%s'",
			attribute,
		)
		ref = &TypeReference{}
	}
	return ref
}

func (g *BuildersGenerator) valueType(typ *concepts.Type) *TypeReference {
	return g.types.ValueReference(typ)
}
