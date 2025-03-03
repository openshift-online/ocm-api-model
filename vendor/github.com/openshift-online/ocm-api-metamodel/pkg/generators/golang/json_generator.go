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
	"strconv"

	"github.com/openshift-online/ocm-api-metamodel/pkg/annotations"
	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/http"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// JSONSupportGeneratorBuilder is an object used to configure and build the JSON support generator.
// Don't create instances directly, use the NewJSONSupporgGenerator function instead.
type JSONSupportGeneratorBuilder struct {
	reporter *reporter.Reporter
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	types    *TypesCalculator
	binding  *http.BindingCalculator
}

// JSONSupportGenerator generates JSON support code. Don't create instances directly, use the
// builder instead.
type JSONSupportGenerator struct {
	reporter *reporter.Reporter
	errors   int
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	types    *TypesCalculator
	buffer   *Buffer
	binding  *http.BindingCalculator
}

// NewJSONSupportGenerator creates a new builder JSON support code generators.
func NewJSONSupportGenerator() *JSONSupportGeneratorBuilder {
	return &JSONSupportGeneratorBuilder{}
}

// Reporter sets the object that will be used to report information about the generation process,
// including errors.
func (b *JSONSupportGeneratorBuilder) Reporter(
	value *reporter.Reporter) *JSONSupportGeneratorBuilder {
	b.reporter = value
	return b
}

// Model sets the model that will be used by the types generator.
func (b *JSONSupportGeneratorBuilder) Model(value *concepts.Model) *JSONSupportGeneratorBuilder {
	b.model = value
	return b
}

// Output sets import path of the output package.
func (b *JSONSupportGeneratorBuilder) Output(value string) *JSONSupportGeneratorBuilder {
	b.output = value
	return b
}

// Package sets the object that will be used to calculate package names.
func (b *JSONSupportGeneratorBuilder) Packages(
	value *PackagesCalculator) *JSONSupportGeneratorBuilder {
	b.packages = value
	return b
}

// Names sets the object that will be used to calculate names.
func (b *JSONSupportGeneratorBuilder) Names(value *NamesCalculator) *JSONSupportGeneratorBuilder {
	b.names = value
	return b
}

// Types sets the object that will be used to calculate types.
func (b *JSONSupportGeneratorBuilder) Types(value *TypesCalculator) *JSONSupportGeneratorBuilder {
	b.types = value
	return b
}

// Binding sets the object that will by used to do HTTP binding calculations.
func (b *JSONSupportGeneratorBuilder) Binding(
	value *http.BindingCalculator) *JSONSupportGeneratorBuilder {
	b.binding = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new types
// generator using it.
func (b *JSONSupportGeneratorBuilder) Build() (generator *JSONSupportGenerator, err error) {
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

	// Create the generator:
	generator = &JSONSupportGenerator{
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
func (g *JSONSupportGenerator) Run() error {
	var err error

	// Generate the helpers:
	err = g.generateHelpers()
	if err != nil {
		return err
	}

	// Generate the code for each type:
	for _, service := range g.model.Services() {
		for _, version := range service.Versions() {
			// Generate the code for the version metadata type:
			err := g.generateVersionMetadataSupport(version)
			if err != nil {
				return err
			}

			// Generate the code for the model types:
			var importRefs []struct {
				path     string
				selector string
			}
			for _, typ := range version.Types() {
				for _, att := range typ.Attributes() {
					if att.LinkOwner() != nil {
						importRefs = append(importRefs,
							struct {
								path     string
								selector string
							}{
								path:     g.packages.VersionImport(att.LinkOwner()),
								selector: g.packages.VersionSelector(att.LinkOwner()),
							})
					}
				}
				switch {
				case typ.IsStruct():
					err = g.generateStructTypeSupport(typ, importRefs)
				case typ.IsList():
					element := typ.Element()
					if element.IsScalar() || element.IsStruct() {
						err = g.generateListTypeSupport(typ, importRefs)
					}
				}
				if err != nil {
					return err
				}
			}

			// Generate the code for the model methods:
			for _, resource := range version.Resources() {
				err = g.generateResourceSupport(resource)
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

func (g *JSONSupportGenerator) generateHelpers() error {
	var err error

	// Calculate the package and file name:
	pkgName := g.packages.HelpersPackage()
	fileName := g.helpersFile()

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

	// Generate the code:
	g.buffer.Import("bytes", "")
	g.buffer.Import("fmt", "")
	g.buffer.Import("io", "")
	g.buffer.Import("net/url", "")
	g.buffer.Import("strconv", "")
	g.buffer.Import("github.com/json-iterator/go", "jsoniter")
	g.buffer.Emit(`
		// NewIterator creates a new JSON iterator that will read to the given source, which
		// can be a slice of bytes, a string, a reader or an existing iterator.
		func NewIterator(source interface{}) (iterator *jsoniter.Iterator, err error) {
			config := jsoniter.Config{}
			api := config.Froze()
			switch typed := source.(type) {
			case []byte:
				iterator = jsoniter.ParseBytes(api, typed)
			case string:
				iterator = jsoniter.ParseString(api, typed)
			case io.Reader:
				iterator = jsoniter.Parse(api, typed, 4096)
			case *jsoniter.Iterator:
				iterator = typed
			default:
				err = fmt.Errorf(
					"expected slice of bytes, string, reader or iterator but got '%T'",
					source,
				)
			}
			return
		}

		// NewStream creates a new JSON stream that will write to the given writer.
		func NewStream(writer io.Writer) *jsoniter.Stream {
			config := jsoniter.Config{
				IndentionStep: 2,
			}
			api := config.Froze()
			return jsoniter.NewStream(api, writer, 0)
		}

		// NewBoolean allocates a new bool in the heap and returns a pointer to it.
		func NewBoolean(value bool) *bool {
			return &value
		}

		// NewInteger allocates a new integer in the heap and returns a pointer to it.
		func NewInteger(value int) *int {
			return &value
		}

		// NewFloat allocates a new floating point value in the heap and returns an pointer
		// to it.
		func NewFloat(value float64) *float64 {
			return &value
		}

		// NewString allocates a new string in the heap and returns a pointer to it.
		func NewString(value string) *string {
			return &value
		}

		// NewDate allocates a new date in the heap and returns a pointer to it.
		func NewDate(value time.Time) *time.Time {
			return &value
		}

		// ParseInteger reads a string and parses it to integer,
		// if an error occurred it returns a non-nil error.
		func ParseInteger(query url.Values, parameterName string) (*int, error) {
			values := query[parameterName]
			count := len(values)
			if count == 0 {
				return nil, nil
			}
		   	if count > 1 {
				err := fmt.Errorf(
					"expected at most one value for parameter '%s' but got %d",
					parameterName, count,
				)
				return nil, err
			}
			value := values[0]
			parsedInt64, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf(
					"value '%s' isn't valid for the '%s' parameter because it isn't an integer: %v",
					value, parameterName, err,
				)
			}
			parsedInt := int(parsedInt64)
			return &parsedInt, nil
		}

		// ParseFloat reads a string and parses it to float,
		// if an error occurred it returns a non-nil error.
		func ParseFloat(query url.Values, parameterName string) (*float64, error) {
			values := query[parameterName]
			count := len(values)
			if count == 0 {
				return nil, nil
			}
		   	if count > 1 {
				err := fmt.Errorf(
					"expected at most one value for parameter '%s' but got %d",
					parameterName, count,
				)
				return nil, err
			}
			value := values[0]
			parsedFloat, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf(
					"value '%s' isn't valid for the '%s' parameter because it isn't a float: %v",
					value, parameterName, err,
				)
			}
			return &parsedFloat, nil
		}

		// ParseString returns a pointer to the string and nil error.
		func ParseString(query url.Values, parameterName string) (*string, error) {
			values := query[parameterName]
			count := len(values)
			if count == 0 {
				return nil, nil
			}
		   	if count > 1 {
				err := fmt.Errorf(
					"expected at most one value for parameter '%s' but got %d",
					parameterName, count,
				)
				return nil, err
			}
			return &values[0], nil
		}

		// ParseBoolean reads a string and parses it to boolean,
		// if an error occurred it returns a non-nil error.
		func ParseBoolean(query url.Values, parameterName string) (*bool, error) {
			values := query[parameterName]
			count := len(values)
			if count == 0 {
				return nil, nil
			}
		   	if count > 1 {
				err := fmt.Errorf(
					"expected at most one value for parameter '%s' but got %d",
					parameterName, count,
				)
				return nil, err
			}
			value := values[0]
			parsedBool, err := strconv.ParseBool(value)
			if err != nil {
				return nil, fmt.Errorf(
					"value '%s' isn't valid for the '%s' parameter because it isn't a boolean: %v",
					value, parameterName, err,
				)
			}
			return &parsedBool, nil
		}

		// ParseDate reads a string and parses it to a time.Time,
		// if an error occurred it returns a non-nil error.
		func ParseDate(query url.Values, parameterName string) (*time.Time, error) {
			values := query[parameterName]
			count := len(values)
			if count == 0 {
				return nil, nil
			}
		   	if count > 1 {
				err := fmt.Errorf(
					"expected at most one value for parameter '%s' but got %d",
					parameterName, count,
				)
				return nil, err
			}
			value := values[0]
			parsedTime, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return nil, fmt.Errorf(
					"value '%s' isn't valid for the '%s' parameter because it isn't a date: %v",
					value, parameterName, err,
				)
			}
			return &parsedTime, nil
		}
	`)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *JSONSupportGenerator) generateVersionMetadataSupport(version *concepts.Version) error {
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
		Function("enumName", g.types.EnumName).
		Function("structName", g.types.StructReference).
		Function("generateReadValue", g.generateReadValue).
		Function("selectorFromLinkOwner", g.selectorFromLinkOwner).
		Function("generateWriteValue", g.generateWriteValue).
		Function("writeRefTypeFunc", g.writeRefTypeFunc).
		Function("readTypeFunc", g.readTypeFunc).
		Function("readRefTypeFunc", g.readRefTypeFunc).
		Function("writeTypeFunc", g.writeTypeFunc).
		Function("valueReference", g.types.ValueReference).
		Build()
	if err != nil {
		return err
	}

	// Generate the code:
	g.generateVersionMetadataSource(version)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *JSONSupportGenerator) generateVersionMetadataSource(version *concepts.Version) {
	g.buffer.Import("fmt", "")
	g.buffer.Import("io", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Import("github.com/json-iterator/go", "jsoniter")
	g.buffer.Emit(`
		// MarshalMetadata writes a value of the metadata type to the given target, which
		// can be a writer or a JSON encoder.
		func MarshalMetadata(object *Metadata, writer io.Writer) error {
			stream := helpers.NewStream(writer)
			writeMetadata(object, stream)
			err := stream.Flush()
			if err != nil {
				return err
			}
			return stream.Error
		}

		func writeMetadata(object *Metadata, stream *jsoniter.Stream) {
			stream.WriteObjectStart()
			if object.bitmap_&1 != 0 {
				stream.WriteObjectField("server_version")
				stream.WriteString(object.serverVersion)
			}
			stream.WriteObjectEnd()
		}

		// UnmarshalMetadata reads a value of the metadata type from the given source, which
		// which can be a reader, a slice of byte or a string.
		func UnmarshalMetadata(source interface{}) (object *Metadata, err error) {
			iterator, err := helpers.NewIterator(source)
			if err != nil {
				return
			}
			object = readMetadata(iterator)
			err = iterator.Error
			return
		}

		func readMetadata(iterator *jsoniter.Iterator) *Metadata {
			object := &Metadata{}
			for {
				field := iterator.ReadObject()
				if field == "" {
					break
				}
				switch field {
				case "server_version":
					object.serverVersion = iterator.ReadString()
					object.bitmap_ |= 1
				default:
					iterator.ReadAny()
				}
			}
			return object
		}
		`,
		"Version", version,
	)
}

func (g *JSONSupportGenerator) generateStructTypeSupport(typ *concepts.Type,
	importRefs []struct {
		path     string
		selector string
	}) error {
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
		Function("enumName", g.types.EnumName).
		Function("fieldName", g.fieldName).
		Function("fieldTag", g.binding.AttributeName).
		Function("generateReadValue", g.generateReadValue).
		Function("selectorFromLinkOwner", g.selectorFromLinkOwner).
		Function("generateWriteValue", g.generateWriteValue).
		Function("writeRefTypeFunc", g.writeRefTypeFunc).
		Function("marshalTypeFunc", g.marshalTypeFunc).
		Function("readTypeFunc", g.readTypeFunc).
		Function("readRefTypeFunc", g.readRefTypeFunc).
		Function("structName", g.types.StructReference).
		Function("unmarshalTypeFunc", g.unmarshalTypeFunc).
		Function("valueReference", g.types.ValueReference).
		Function("writeTypeFunc", g.writeTypeFunc).
		Build()
	if err != nil {
		return err
	}

	for _, ref := range importRefs {
		g.buffer.Import(ref.path, ref.selector)
	}

	// Generate the code:
	g.generateStructTypeSource(typ)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *JSONSupportGenerator) generateStructTypeSource(typ *concepts.Type) {
	g.buffer.Import("fmt", "")
	g.buffer.Import("io", "")
	g.buffer.Import("time", "")
	g.buffer.Import("github.com/json-iterator/go", "jsoniter")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		{{ $structName := structName .Type }}
		{{ $marshalTypeFunc := marshalTypeFunc .Type }}
		{{ $writeTypeFunc := writeTypeFunc .Type }}
		{{ $unmarshalTypeFunc := unmarshalTypeFunc .Type }}
		{{ $readTypeFunc := readTypeFunc .Type }}

		// {{ $marshalTypeFunc }} writes a value of the '{{ .Type.Name }}' type to the given writer.
		func {{ $marshalTypeFunc }}(object *{{ $structName }}, writer io.Writer) error {
			stream := helpers.NewStream(writer)
			{{ $writeTypeFunc }}(object, stream)
			err := stream.Flush()
			if err != nil {
				return err
			}
			return stream.Error
		}

		// {{ $writeTypeFunc }} writes a value of the '{{ .Type.Name }}' type to the given stream.
		func {{ $writeTypeFunc }}(object *{{ $structName }}, stream *jsoniter.Stream) {
			count := 0
			stream.WriteObjectStart()
			{{ if .Type.IsClass }}
				stream.WriteObjectField("kind")
				if object.bitmap_&1 != 0 {
					stream.WriteString({{ $structName }}LinkKind)
				} else {
					stream.WriteString({{ $structName }}Kind)
				}
				count++
				if object.bitmap_&2 != 0 {
					if count > 0 {
						stream.WriteMore()
					}
					stream.WriteObjectField("id")
					stream.WriteString(object.id)
					count++
				}
				if object.bitmap_&4 != 0 {
					if count > 0 {
						stream.WriteMore()
					}
					stream.WriteObjectField("href")
					stream.WriteString(object.href)
					{{ if .Type.Attributes }}
						count++
					{{ end }}
				}
			{{ end }}
			{{ if .Type.Attributes }}
				var present_ bool
				{{ $n := len .Type.Attributes }}
				{{ range $i, $v := .Type.Attributes }}
					{{ $fieldName := fieldName $v }}
					{{ $fieldTag := fieldTag $v }}
					{{ $fieldMask := bitMask $v }}
					{{ if .Type.IsScalar }}
						present_ = object.bitmap_&{{ $fieldMask }} != 0
					{{ else }}
						present_ = object.bitmap_&{{ $fieldMask }} != 0 && object.{{ $fieldName }} != nil
					{{ end }}
					if present_ {
						if count > 0 {
							stream.WriteMore()
						}
						stream.WriteObjectField("{{ $fieldTag }}")
						{{ generateWriteValue (print "object." $fieldName) $v.Type $v.Link $v.LinkOwner }}
						{{ if lt $i (sub $n 1) }}
							count++
						{{ end }}
					}
				{{ end }}
			{{ end }}
			stream.WriteObjectEnd()
		}

		// {{ $unmarshalTypeFunc }} reads a value of the '{{ .Type.Name }}' type from the given
		// source, which can be an slice of bytes, a string or a reader.
		func {{ $unmarshalTypeFunc }}(source interface{}) (object *{{ $structName }}, err error) {
			iterator, err := helpers.NewIterator(source)
			if err != nil {
				return
			}
			object = {{ $readTypeFunc }}(iterator)
			err = iterator.Error
			return
		}

		// {{ $readTypeFunc }} reads a value of the '{{ .Type.Name }}' type from the given iterator.
		func {{ $readTypeFunc }}(iterator *jsoniter.Iterator) *{{ $structName }} {
			object := &{{ $structName }}{}
			for {
				field := iterator.ReadObject()
				if field == "" {
					break
				}
				switch field {
				{{ if .Type.IsClass }}
					case "kind":
						value := iterator.ReadString()
						if value == {{ $structName }}LinkKind {
							object.bitmap_ |= 1
						}
					case "id":
						object.id = iterator.ReadString()
						object.bitmap_ |= 2
					case "href":
						object.href = iterator.ReadString()
						object.bitmap_ |= 4
				{{ end }}
				{{ range .Type.Attributes }}
					{{ $fieldName := fieldName . }}
					{{ $fieldTag := fieldTag . }}
					{{ $fieldMask := bitMask . }}
					case "{{ $fieldTag }}":
						{{ generateReadValue "value" .Type .Link .LinkOwner }}
						object.{{ $fieldName }} = value
						object.bitmap_ |= {{ $fieldMask }}
				{{ end }}
				default:
					iterator.ReadAny()
				}
			}
			return object
		}
		`,
		"Type", typ,
	)
}

func (g *JSONSupportGenerator) generateListTypeSupport(typ *concepts.Type,
	importRefs []struct {
		path     string
		selector string
	}) error {
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
		Function("enumName", g.types.EnumName).
		Function("generateReadValue", g.generateReadValue).
		Function("selectorFromLinkOwner", g.selectorFromLinkOwner).
		Function("generateWriteValue", g.generateWriteValue).
		Function("writeRefTypeFunc", g.writeRefTypeFunc).
		Function("marshalTypeFunc", g.marshalTypeFunc).
		Function("readTypeFunc", g.readTypeFunc).
		Function("readRefTypeFunc", g.readRefTypeFunc).
		Function("structName", g.types.StructReference).
		Function("unmarshalTypeFunc", g.unmarshalTypeFunc).
		Function("valueReference", g.types.ValueReference).
		Function("writeTypeFunc", g.writeTypeFunc).
		Build()
	if err != nil {
		return err
	}

	for _, ref := range importRefs {
		g.buffer.Import(ref.path, ref.selector)
	}

	// Generate the code:
	g.generateListTypeSource(typ)
	// Write the generated code:
	return g.buffer.Write()
}

func (g *JSONSupportGenerator) generateListTypeSource(
	typ *concepts.Type,
) {
	var linkOwner *concepts.Version
	g.buffer.Import("fmt", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Import("github.com/json-iterator/go", "jsoniter")
	g.buffer.Emit(`
		{{ $structName := structName .Type }}
		{{ $sliceType := valueReference .Type }}
		{{ $marshalTypeFunc := marshalTypeFunc .Type }}
		{{ $writeTypeFunc := writeTypeFunc .Type }}
		{{ $unmarshalTypeFunc := unmarshalTypeFunc .Type }}
		{{ $readTypeFunc := readTypeFunc .Type }}

		// {{ $marshalTypeFunc }} writes a list of values of the '{{ .Type.Element.Name }}' type to
		// the given writer.
		func {{ $marshalTypeFunc }}(list {{ $sliceType }}, writer io.Writer) error {
			stream := helpers.NewStream(writer)
			{{ $writeTypeFunc }}(list, stream)
			err := stream.Flush()
			if err != nil {
				return err
			}
			return stream.Error
		}

		// {{ $writeTypeFunc }} writes a list of value of the '{{ .Type.Element.Name }}' type to
		// the given stream.
		func {{ $writeTypeFunc }}(list {{ $sliceType }}, stream *jsoniter.Stream) {
			stream.WriteArrayStart()
			for i, value := range list {
				if i > 0 {
					stream.WriteMore()
				}
				{{ generateWriteValue "value" .Type.Element false .linkOwner }}
			}
			stream.WriteArrayEnd()
		}

		// {{ $unmarshalTypeFunc }} reads a list of values of the '{{ .Type.Element.Name }}' type
		// from the given source, which can be a slice of bytes, a string or a reader.
		func {{ $unmarshalTypeFunc }}(source interface{}) (items {{ $sliceType }}, err error) {
			iterator, err := helpers.NewIterator(source)
			if err != nil {
				return
			}
			items = {{ $readTypeFunc }}(iterator)
			err = iterator.Error
			return
		}

		// {{ $readTypeFunc }} reads list of values of the ''{{ .Type.Element.Name }}' type from
		// the given iterator.
		func {{ $readTypeFunc }}(iterator *jsoniter.Iterator) {{ $sliceType }} {
			list := {{ valueReference .Type }}{}
			for iterator.ReadArray() {
				{{ generateReadValue "item" .Type.Element false .linkOwner }}
				list = append(list, item)
			}
			return list
		}
		`,
		"Type", typ,
		"linkOwner", linkOwner,
	)
}

func (g *JSONSupportGenerator) generateResourceSupport(resource *concepts.Resource) error {
	var err error

	// Calculate the package and file name:
	pkgName := g.packages.VersionPackage(resource.Owner())
	fileName := g.resourceFile(resource)

	// Create the buffer for the generated code:
	g.buffer, err = NewBuffer().
		Reporter(g.reporter).
		Output(g.output).
		Packages(g.packages).
		Package(pkgName).
		File(fileName).
		Function("clientRequestName", g.clientRequestName).
		Function("clientResponseName", g.clientResponseName).
		Function("defaultValue", g.defaultValue).
		Function("enumName", g.types.EnumName).
		Function("generateReadBodyParameter", g.generateReadBodyParameter).
		Function("generateReadQueryParameter", g.generateReadQueryParameter).
		Function("generateReadValue", g.generateReadValue).
		Function("selectorFromLinkOwner", g.selectorFromLinkOwner).
		Function("generateWriteBodyParameter", g.generateWriteBodyParameter).
		Function("generateWriteValue", g.generateWriteValue).
		Function("writeRefTypeFunc", g.writeRefTypeFunc).
		Function("marshalTypeFunc", g.marshalTypeFunc).
		Function("parameterFieldName", g.parameterFieldName).
		Function("parameterQueryName", g.binding.QueryParameterName).
		Function("readResponseFunc", g.readResponseFunc).
		Function("readTypeFunc", g.readTypeFunc).
		Function("readRefTypeFunc", g.readRefTypeFunc).
		Function("requestBodyParameters", g.binding.RequestBodyParameters).
		Function("requestQueryParameters", g.binding.RequestQueryParameters).
		Function("responseBodyParameters", g.binding.ResponseParameters).
		Function("structName", g.types.StructReference).
		Function("unmarshalTypeFunc", g.unmarshalTypeFunc).
		Function("valueReference", g.types.ValueReference).
		Function("writeRequestFunc", g.writeRequestFunc).
		Function("writeTypeFunc", g.writeTypeFunc).
		Build()
	if err != nil {
		return err
	}

	// Generate the code:
	for _, method := range resource.Methods() {
		g.generateMethodSource(method)
	}

	// Write the generated code:
	return g.buffer.Write()
}

func (g *JSONSupportGenerator) generateMethodSource(method *concepts.Method) {
	switch {
	case method.IsAdd():
		g.generateAddMethodSource(method)
	case method.IsDelete():
		g.generateDeleteMethodSource(method)
	case method.IsGet():
		g.generateGetMethodSource(method)
	case method.IsList():
		g.generateListMethodSource(method)
	case method.IsPost():
		g.generatePostMethodSource(method)
	case method.IsSearch():
		g.generateSearchMethodSource(method)
	case method.IsUpdate():
		g.generateUpdateMethodSource(method)
	case method.IsAction():
		g.generateActionMethodSource(method)
	default:
		g.reporter.Errorf(
			"Don't know how to generate encoding/decoding code for method '%s'",
			method,
		)
	}
}

func (g *JSONSupportGenerator) generateAddMethodSource(method *concepts.Method) {
	// For `Add` methods we need to put in the request and response the `Body` parameter:
	body := method.GetParameter(nomenclator.Body)

	// Generate the code:
	g.buffer.Import("net/http", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		func {{ writeRequestFunc .Method }}(request *{{ clientRequestName .Method }}, writer io.Writer) error {
			return {{ marshalTypeFunc .Body.Type }}(request.body, writer)
		}

		func {{ readResponseFunc .Method }}(response *{{ clientResponseName .Method }}, reader io.Reader) error {
			var err error
			response.body, err = {{ unmarshalTypeFunc .Body.Type }}(reader)
			return err
		}
		`,
		"Method", method,
		"Body", body,
	)
}

func (g *JSONSupportGenerator) generateDeleteMethodSource(method *concepts.Method) {
	g.buffer.Import("net/http", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		func {{ writeRequestFunc .Method }}(request *{{ clientRequestName .Method }}, writer io.Writer) error {
			return nil
		}

		func {{ readResponseFunc .Method }}(response *{{ clientResponseName .Method }}, reader io.Reader) error {
			return nil
		}
		`,
		"Method", method,
	)
}

func (g *JSONSupportGenerator) generateGetMethodSource(method *concepts.Method) {
	// For `Get` methods we need to put in the response the `Body` parameter:
	body := method.GetParameter(nomenclator.Body)

	// Generate the code:
	g.buffer.Import("net/http", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		func {{ writeRequestFunc .Method }}(request *{{ clientRequestName .Method }}, writer io.Writer) error {
			return nil
		}

		func {{ readResponseFunc .Method }}(response *{{ clientResponseName .Method }}, reader io.Reader) error {
			var err error
			response.body, err = {{ unmarshalTypeFunc .Body.Type }}(reader)
			return err
		}
		`,
		"Method", method,
		"Body", body,
	)
}

func (g *JSONSupportGenerator) generateListMethodSource(method *concepts.Method) {
	// For list methods we want to put first the paging parameters and last the items, so we
	// need to classify the parameters.
	var page *concepts.Parameter
	var size *concepts.Parameter
	var total *concepts.Parameter
	var items *concepts.Parameter
	var other []*concepts.Parameter
	var linkOwner *concepts.Version
	for _, parameter := range method.Parameters() {
		switch {
		case parameter.Name().Equals(nomenclator.Page):
			page = parameter
		case parameter.Name().Equals(nomenclator.Size):
			size = parameter
		case parameter.Name().Equals(nomenclator.Total):
			total = parameter
		case parameter.Name().Equals(nomenclator.Items):
			items = parameter
		default:
			other = append(other, parameter)
		}
	}

	// Generate the code:
	g.buffer.Import("net/http", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		func {{ writeRequestFunc .Method }}(request *{{ clientRequestName .Method }}, writer io.Writer) error {
			return nil
		}

		func {{ readResponseFunc .Method }}(response *{{ clientResponseName .Method }}, reader io.Reader) error {
			iterator, err := helpers.NewIterator(reader)
			if err != nil {
				return err
			}
			for {
				field := iterator.ReadObject()
				if field == "" {
					break
				}
				switch field {
				{{ if .Page }}
					{{ generateReadBodyParameter "response" .Page }}
				{{ end }}
				{{ if .Size }}
					{{ generateReadBodyParameter "response" .Size }}
				{{ end }}
				{{ if .Total }}
					{{ generateReadBodyParameter "response" .Total }}
				{{ end }}
				{{ range .Other }}
					{{ if .Out }}
						{{ generateReadBodyParameter "response" . }}
					{{ end }}
				{{ end }}
				case "items":
				{{ generateReadValue "items" .Items.Type false .LinkOwner }}
				{{ if and .Items.Type.IsList .Items.Type.Element.IsScalar }}
					response.items = items
				{{ else }}
					response.items = &{{ structName .Items.Type }}{
						items: items,
					}
				{{ end }}
				default:
					iterator.ReadAny()
				}
			}
			return iterator.Error
		}
		`,
		"Method", method,
		"Page", page,
		"Size", size,
		"Total", total,
		"Items", items,
		"Other", other,
		"LinkOwner", linkOwner,
	)
}

func (g *JSONSupportGenerator) generatePostMethodSource(method *concepts.Method) {
	// Post methods have at most one input parameter (the request) and one output parameter
	// (the response):
	var request *concepts.Parameter
	var response *concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.In() {
			request = parameter
		}
		if parameter.Out() {
			response = parameter
		}
	}

	// Generate the code:
	g.buffer.Import("net/http", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		func {{ writeRequestFunc .Method }}(request *{{ clientRequestName .Method }}, writer io.Writer) error {
			{{ if .Request }}
				return {{ marshalTypeFunc .Request.Type }}(request.{{ parameterFieldName .Request }}, writer)
			{{ else }}
				return nil
			{{ end }}
		}

		func {{ readResponseFunc .Method }}(response *{{ clientResponseName .Method }}, reader io.Reader) error {
			{{ if .Response }}
				var err error
				response.{{ parameterFieldName .Response }}, err = {{ unmarshalTypeFunc .Response.Type }}(reader)
				return err
			{{ else }}
				return nil
			{{ end }}
		}
		`,
		"Method", method,
		"Request", request,
		"Response", response,
	)
}

func (g *JSONSupportGenerator) generateSearchMethodSource(method *concepts.Method) {
	// For list methods we want to put first the paging parameters and last the items, so we
	// need to classify the parameters.
	var page *concepts.Parameter
	var size *concepts.Parameter
	var total *concepts.Parameter
	var body *concepts.Parameter
	var items *concepts.Parameter
	var other []*concepts.Parameter
	var linkOwner *concepts.Version
	for _, parameter := range method.Parameters() {
		switch {
		case parameter.Name().Equals(nomenclator.Page):
			page = parameter
		case parameter.Name().Equals(nomenclator.Size):
			size = parameter
		case parameter.Name().Equals(nomenclator.Total):
			total = parameter
		case parameter.Name().Equals(nomenclator.Body):
			body = parameter
		case parameter.Name().Equals(nomenclator.Items):
			items = parameter
		default:
			other = append(other, parameter)
		}
	}

	// Generate the code:
	g.buffer.Import("net/http", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		func {{ writeRequestFunc .Method }}(request *{{ clientRequestName .Method }}, writer io.Writer) error {
			{{ if .Body }}
				return {{ marshalTypeFunc .Body.Type }}(request.{{ parameterFieldName .Body }}, writer)
			{{ else }}
				return nil
			{{ end }}
		}

		func {{ readResponseFunc .Method }}(response *{{ clientResponseName .Method }}, reader io.Reader) error {
			iterator, err := helpers.NewIterator(reader)
			if err != nil {
				return err
			}
			for {
				field := iterator.ReadObject()
				if field == "" {
					break
				}
				switch field {
				{{ if .Page }}
					{{ generateReadBodyParameter "response" .Page }}
				{{ end }}
				{{ if .Size }}
					{{ generateReadBodyParameter "response" .Size }}
				{{ end }}
				{{ if .Total }}
					{{ generateReadBodyParameter "response" .Total }}
				{{ end }}
				{{ range .Other }}
					{{ if .Out }}
						{{ generateReadBodyParameter "response" . }}
					{{ end }}
				{{ end }}
				case "items":
					{{ generateReadValue "items" .Items.Type false .LinkOwner }}
					{{ if and .Items.Type.IsList .Items.Type.Element.IsScalar }}
						response.items = items
					{{ else }}
						response.items = &{{ structName .Items.Type }}{
							items: items,
						}
					{{ end }}
				default:
					iterator.ReadAny()
				}
			}
			return iterator.Error
		}
		`,
		"Version", method.Owner().Owner(),
		"Method", method,
		"Page", page,
		"Size", size,
		"Total", total,
		"Body", body,
		"Items", items,
		"Other", other,
		"LinkOwner", linkOwner,
	)
}

func (g *JSONSupportGenerator) generateUpdateMethodSource(method *concepts.Method) {
	// For `Update` methods we need to put in the request and response the `Body` parameter:
	body := method.GetParameter(nomenclator.Body)

	// Generate the code:
	g.buffer.Import("net/http", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		func {{ writeRequestFunc .Method }}(request *{{ clientRequestName .Method }}, writer io.Writer) error {
			return {{ marshalTypeFunc .Body.Type }}(request.body, writer)
		}

		func {{ readResponseFunc .Method }}(response *{{ clientResponseName .Method }}, reader io.Reader) error {
			var err error
			response.body, err = {{ unmarshalTypeFunc .Body.Type }}(reader)
			return err
		}
		`,
		"Method", method,
		"Body", body,
	)
}

func (g *JSONSupportGenerator) generateActionMethodSource(method *concepts.Method) {
	g.buffer.Import("net/http", "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		{{ $requestBodyParameters := requestBodyParameters .Method }}
		{{ $responseBodyParameters := responseBodyParameters .Method }}
		func {{ writeRequestFunc .Method }}(request *{{ clientRequestName .Method }}, writer io.Writer) error {
			{{ if $requestBodyParameters }} 
				count := 0
				stream := helpers.NewStream(writer)
				stream.WriteObjectStart()
				{{ range $requestBodyParameters }}
					{{ generateWriteBodyParameter "request" . }}
				{{ end }}
				stream.WriteObjectEnd()
				err := stream.Flush()
				if err != nil {
					return err
				}
				return stream.Error
			{{ else }}
				return nil
			{{ end }}
		}

		func {{ readResponseFunc .Method }}(response *{{ clientResponseName .Method }}, reader io.Reader) error {
			{{ if $responseBodyParameters }} 
				iterator, err := helpers.NewIterator(reader)
				if err != nil {
					return err
				}
				for {
					field := iterator.ReadObject()
					if field == "" {
						break
					}
					switch field {
					{{ range $responseBodyParameters }}
						{{ generateReadBodyParameter "response" . }}
					{{ end }}
					default:
						iterator.ReadAny()
					}
				}
				return iterator.Error
			{{ else }}
				return nil
			{{ end }}
		}
		`,
		"Method", method,
	)
}

func (g *JSONSupportGenerator) generateReadQueryParameter(parameter *concepts.Parameter) string {
	typ := parameter.Type()
	if !typ.IsScalar() {
		g.reporter.Errorf(
			"Don't know how to generate code to read query parameter '%s'",
			parameter,
		)
		return ""
	}
	field := g.parameterFieldName(parameter)
	tag := g.binding.QueryParameterName(parameter)
	kind := typ.Name().Camel()
	return g.buffer.Eval(`
		{{ $field := parameterFieldName .Parameter }}
		{{ $tag := parameterQueryName .Parameter }}
		{{ $kind := .Parameter.Type.Name.Camel }}

		request.{{ $field }}, err = helpers.Parse{{ $kind }}(query, "{{ $tag }}")
		if err != nil {
			return err
		}
		{{ if .Parameter.Default }}
			if request.{{ $field }} == nil {
				request.{{ $field }} = helpers.New{{ $kind }}({{ defaultValue .Parameter }})
			}
		{{ end }}
		`,
		"Field", field,
		"Tag", tag,
		"Kind", kind,
		"Parameter", parameter,
	)
}

func (g *JSONSupportGenerator) generateReadBodyParameter(object string,
	parameter *concepts.Parameter) string {
	field := g.parameterFieldName(parameter)
	tag := g.binding.BodyParameterName(parameter)
	typ := parameter.Type()
	var linkOwner *concepts.Version
	return g.buffer.Eval(`
		case "{{ .Tag }}":
			{{ generateReadValue "value" .Type false .LinkOwner }}
			{{ if .Type.IsScalar }}
				{{ .Object }}.{{ .Field }} = &value
			{{ else }}
				{{ .Object }}.{{ .Field }} = value
			{{ end }}
		`,
		"Object", object,
		"Field", field,
		"Tag", tag,
		"Type", typ,
		"LinkOwner", linkOwner,
	)
}

func (g *JSONSupportGenerator) generateReadValue(variable string, typ *concepts.Type,
	link bool,
	linkOwner *concepts.Version) string {
	g.buffer.Import("time", "")
	return g.buffer.Eval(`
		{{ if .Type.IsBoolean }}
			{{ .Variable }} := iterator.ReadBool()
		{{ else if .Type.IsInteger }}
			{{ .Variable }} := iterator.ReadInt()
		{{ else if .Type.IsLong }}
			{{ .Variable }} := iterator.ReadInt64()
		{{ else if .Type.IsFloat }}
			{{ .Variable }} := iterator.ReadFloat64()
		{{ else if .Type.IsString }}
			{{ .Variable }} := iterator.ReadString()
		{{ else if .Type.IsInterface }}
			var {{ .Variable }} interface{}
			iterator.ReadVal(&{{ .Variable }})
		{{ else if .Type.IsDate }}
			text := iterator.ReadString()
			{{ .Variable }}, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
		{{ else if .Type.IsEnum }}
			text := iterator.ReadString()
			{{ .Variable }} := {{ enumName .Type }}(text)
		{{ else if .Type.IsStruct }}
			{{ .Variable }} := {{ readRefTypeFunc .Type .LinkOwner }}(iterator)
		{{ else if .Type.IsList }}
			{{ if .Link }}
			 	{{ $selectorFromLinkOwner := selectorFromLinkOwner .LinkOwner }}
				{{ $structName := structName .Type }}
				{{ .Variable }} := &{{ $selectorFromLinkOwner }}{{ $structName }}{}
				for {
					field := iterator.ReadObject()
					if field == "" {
						break
					}
					switch field {
					case "kind":
						text := iterator.ReadString()
						{{ .Variable }}.SetLink(text == {{ $selectorFromLinkOwner }}{{ $structName }}LinkKind)
					case "href":
						{{ .Variable }}.SetHREF(iterator.ReadString())
					case "items":
						{{ .Variable }}.SetItems({{ readRefTypeFunc .Type .LinkOwner }}(iterator))
					default:
						iterator.ReadAny()
					}
				}
			{{ else }}
				{{ .Variable }} := {{ readTypeFunc .Type }}(iterator)
			{{ end }}
		{{ else if .Type.IsMap }}
			{{ .Variable }} := {{ valueReference .Type }}{}
			for {
				key := iterator.ReadObject()
				if key == "" {
					break
				}
				{{ generateReadValue "item" .Type.Element false .LinkOwner }}
				{{ .Variable }}[key] = item
			}
		{{ else }}
			iterator.ReadAny()
		{{ end }}
		`,
		"Variable", variable,
		"Type", typ,
		"Link", link,
		"LinkOwner", linkOwner,
	)
}

func (g *JSONSupportGenerator) generateWriteBodyParameter(object string,
	parameter *concepts.Parameter,
) string {
	typ := parameter.Type()
	field := g.parameterFieldName(parameter)
	tag := g.binding.BodyParameterName(parameter)
	var value string
	var linkOwner *concepts.Version
	if typ.IsScalar() && !typ.IsInterface() {
		value = g.buffer.Eval(
			`*{{ .Object }}.{{ .Field }}`,
			"Object", object,
			"Field", field,
		)
	} else {
		value = g.buffer.Eval(
			`{{ .Object }}.{{ .Field }}`,
			"Object", object,
			"Field", field,
		)
	}
	return g.buffer.Eval(`
		if {{ .Object }}.{{ .Field }} != nil {
			if count > 0 {
				stream.WriteMore()
			}
			stream.WriteObjectField("{{ .Tag }}")
			{{ generateWriteValue .Value .Type false .LinkOwner }}
			count++
		}
		`,
		"Object", object,
		"Field", field,
		"Tag", tag,
		"Value", value,
		"Type", typ,
		"LinkOwner", linkOwner,
	)
}

func (g *JSONSupportGenerator) generateWriteValue(value string,
	typ *concepts.Type,
	link bool,
	linkOwner *concepts.Version) string {
	g.buffer.Import("sort", "")
	g.buffer.Import("time", "")
	return g.buffer.Eval(`
		{{ if .Type.IsBoolean }}
			stream.WriteBool({{ .Value }})
		{{ else if .Type.IsInteger }}
			stream.WriteInt({{ .Value }})
		{{ else if .Type.IsLong }}
			stream.WriteInt64({{ .Value }})
		{{ else if .Type.IsFloat }}
			stream.WriteFloat64({{ .Value }})
		{{ else if .Type.IsString }}
			stream.WriteString({{ .Value }})
		{{ else if .Type.IsDate }}
			stream.WriteString(({{ .Value }}).Format(time.RFC3339))
		{{ else if .Type.IsEnum }}
			stream.WriteString(string({{ .Value }}))
		{{ else if .Type.IsInterface }}
			stream.WriteVal({{ .Value }})
		{{ else if .Type.IsStruct }}
			{{ writeRefTypeFunc .Type .LinkOwner }}({{ .Value}}, stream)
		{{ else if .Type.IsList }}
			{{ if .Link }}
				stream.WriteObjectStart()
				stream.WriteObjectField("items")
				{{ writeRefTypeFunc .Type .LinkOwner }}({{ .Value }}.Items(), stream)
				stream.WriteObjectEnd()
			{{ else }}
				{{ writeTypeFunc .Type }}({{ .Value }}, stream)
			{{ end }}
		{{ else if .Type.IsMap }}
			if {{ .Value }} != nil {
				stream.WriteObjectStart()
				keys := make([]string, len({{ .Value }}))
				i := 0;
				for key := range {{ .Value }} {
					keys[i] = key
					i++
				}
				sort.Strings(keys)
				for i, key := range keys {
					if i > 0 {
						stream.WriteMore()
					}
					item := {{ .Value }}[key]
					stream.WriteObjectField(key)
					{{ generateWriteValue "item" .Type.Element false .LinkOwner }}
				}
				stream.WriteObjectEnd()
			} else {
				stream.WriteNil()
			}
		{{ end }}
		`,
		"Value", value,
		"Type", typ,
		"Link", link,
		"LinkOwner", linkOwner,
	)
}

func (g *JSONSupportGenerator) helpersFile() string {
	return g.names.File(names.Cat(nomenclator.JSON, nomenclator.Helpers))
}

func (g *JSONSupportGenerator) metadataFile() string {
	return g.names.File(names.Cat(nomenclator.Metadata, nomenclator.Reader))
}

func (g *JSONSupportGenerator) typeFile(typ *concepts.Type) string {
	return g.names.File(names.Cat(typ.Name(), nomenclator.Type, nomenclator.JSON))
}

func (g *JSONSupportGenerator) resourceFile(resource *concepts.Resource) string {
	return g.names.File(names.Cat(resource.Name(), nomenclator.Resource, nomenclator.JSON))
}

func (g *JSONSupportGenerator) marshalTypeFunc(typ *concepts.Type) string {
	name := annotations.GoName(typ)
	if name == "" {
		name = g.names.Public(typ.Name())
	}
	name = "Marshal" + name
	return name
}

func (g *JSONSupportGenerator) writeTypeFunc(typ *concepts.Type) string {
	name := names.Cat(nomenclator.Write, typ.Name())
	return g.names.Public(name)
}

func (g *JSONSupportGenerator) writeRefTypeFunc(typ *concepts.Type, refVersion *concepts.Version) string {
	name := names.Cat(nomenclator.Write, typ.Name())
	if refVersion != nil {
		version := g.packages.VersionSelector(refVersion)
		return fmt.Sprintf("%s.%s", version, g.names.Public(name))
	}
	return g.names.Public(name)
}

func (g *JSONSupportGenerator) unmarshalTypeFunc(typ *concepts.Type) string {
	name := annotations.GoName(typ)
	if name == "" {
		name = g.names.Public(typ.Name())
	}
	name = "Unmarshal" + name
	return name
}

func (g *JSONSupportGenerator) readTypeFunc(typ *concepts.Type) string {
	name := names.Cat(nomenclator.Read, typ.Name())
	return g.names.Public(name)
}

func (g *JSONSupportGenerator) readRefTypeFunc(typ *concepts.Type, refVersion *concepts.Version) string {
	name := names.Cat(nomenclator.Read, typ.Name())
	if refVersion != nil {
		version := g.packages.VersionSelector(refVersion)
		return fmt.Sprintf("%s.%s", version, g.names.Public(name))
	}
	return g.names.Public(name)
}

func (g *JSONSupportGenerator) fieldName(attribute *concepts.Attribute) string {
	return g.names.Private(attribute.Name())
}

func (g *JSONSupportGenerator) parameterFieldName(parameter *concepts.Parameter) string {
	return g.names.Private(parameter.Name())
}

func (g *JSONSupportGenerator) clientRequestName(method *concepts.Method) string {
	resource := method.Owner()
	var name string
	if resource.IsRoot() {
		name = annotations.GoName(method)
		if name == "" {
			name = g.names.Public(method.Name())
		}
	} else {
		resourceName := annotations.GoName(resource)
		if resourceName == "" {
			resourceName = g.names.Public(resource.Name())
		}
		methodName := annotations.GoName(method)
		if methodName == "" {
			methodName = g.names.Public(method.Name())
		}
		name = resourceName + methodName
	}
	name += "Request"
	return name
}

func (g *JSONSupportGenerator) clientResponseName(method *concepts.Method) string {
	resource := method.Owner()
	var name string
	if resource.IsRoot() {
		name = annotations.GoName(method)
		if name == "" {
			name = g.names.Public(method.Name())
		}
	} else {
		resourceName := annotations.GoName(resource)
		if resourceName == "" {
			resourceName = g.names.Public(resource.Name())
		}
		methodName := annotations.GoName(method)
		if methodName == "" {
			methodName = g.names.Public(method.Name())
		}
		name = resourceName + methodName
	}
	name += "Response"
	return name
}

func (g *JSONSupportGenerator) writeRequestFunc(method *concepts.Method) string {
	resource := method.Owner()
	var name *names.Name
	if resource.IsRoot() {
		name = names.Cat(
			nomenclator.Write,
			method.Name(),
			nomenclator.Request,
		)
	} else {
		name = names.Cat(
			nomenclator.Write,
			resource.Name(),
			method.Name(),
			nomenclator.Request,
		)
	}
	return g.names.Private(name)
}

func (g *JSONSupportGenerator) readResponseFunc(method *concepts.Method) string {
	resource := method.Owner()
	var name *names.Name
	if resource.IsRoot() {
		name = names.Cat(
			nomenclator.Read,
			method.Name(),
			nomenclator.Response,
		)
	} else {
		name = names.Cat(
			nomenclator.Read,
			resource.Name(),
			method.Name(),
			nomenclator.Response,
		)
	}
	return g.names.Private(name)
}

func (g *JSONSupportGenerator) defaultValue(parameter *concepts.Parameter) string {
	switch value := parameter.Default().(type) {
	case nil:
		return "nil"
	case bool:
		return strconv.FormatBool(value)
	case int:
		return strconv.FormatInt(int64(value), 10)
	case string:
		return strconv.Quote(value)
	default:
		g.reporter.Errorf(
			"Don't know how to render default value '%v' for parameter '%s'",
			value, parameter,
		)
		return ""
	}
}

func (g JSONSupportGenerator) selectorFromLinkOwner(linkOwner *concepts.Version) string {
	if linkOwner == nil {
		return ""
	}
	return fmt.Sprintf("%s.", g.packages.VersionSelector(linkOwner))
}
