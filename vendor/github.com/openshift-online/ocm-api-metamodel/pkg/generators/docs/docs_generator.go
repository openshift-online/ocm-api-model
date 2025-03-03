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

package docs

import (
	"fmt"
	"strings"

	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/markdown"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// DocsGeneratorBuilder is an object used to configure and build the builders generator. Don't
// create instances directly, use the NewDocsGenerator function instead.
type DocsGeneratorBuilder struct {
	reporter *reporter.Reporter
	model    *concepts.Model
	output   string
	names    *markdown.NamesCalculator
}

// DocsGenerator generates code for the builders of the model types. Don't create instances
// directly, use the builder instead.
type DocsGenerator struct {
	reporter *reporter.Reporter
	errors   int
	model    *concepts.Model
	output   string
	names    *markdown.NamesCalculator
	buffer   *markdown.Buffer
}

// NewDocsGenerator creates a new builder for builders generators.
func NewDocsGenerator() *DocsGeneratorBuilder {
	return new(DocsGeneratorBuilder)
}

// Reporter sets the object that will be used to report information about the generation process,
// including errors.
func (b *DocsGeneratorBuilder) Reporter(value *reporter.Reporter) *DocsGeneratorBuilder {
	b.reporter = value
	return b
}

// Model sets the model that will be used by the types generator.
func (b *DocsGeneratorBuilder) Model(value *concepts.Model) *DocsGeneratorBuilder {
	b.model = value
	return b
}

// Output sets path of the output directory.
func (b *DocsGeneratorBuilder) Output(value string) *DocsGeneratorBuilder {
	b.output = value
	return b
}

// Names sets the object that will be used to calculate names.
func (b *DocsGeneratorBuilder) Names(value *markdown.NamesCalculator) *DocsGeneratorBuilder {
	b.names = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new
// documetation generator using it.
func (b *DocsGeneratorBuilder) Build() (generator *DocsGenerator, err error) {
	// Check that the mandatory parameters have been provided:
	if b.reporter == nil {
		err = fmt.Errorf("parameter 'reporter' is mandatory")
		return
	}
	if b.model == nil {
		err = fmt.Errorf("parameter 'model' is mandatory")
		return
	}
	if b.output == "" {
		err = fmt.Errorf("parameter 'output' is mandatory")
		return
	}

	// Create the generator:
	generator = new(DocsGenerator)
	generator.reporter = b.reporter
	generator.model = b.model
	generator.output = b.output
	generator.names = b.names

	return
}

// Run executes the documentation generator.
func (g *DocsGenerator) Run() error {
	var err error

	// Generate the index:
	err = g.generateIndex()
	if err != nil {
		return err
	}

	for _, service := range g.model.Services() {
		for _, version := range service.Versions() {
			// Generate documentation for each resource:
			for _, resource := range version.Resources() {
				err = g.generateResource(resource)
				if err != nil {
					return err
				}
			}

			// Generate documentation for each type:
			for _, typ := range version.Types() {
				switch {
				case typ.IsEnum() || typ.IsStruct():
					err = g.generateType(typ)
					if err != nil {
						return err
					}
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

func (g *DocsGenerator) generateIndex() error {
	var err error

	// Calculate the file name:
	fileName := g.names.File(nomenclator.Index)

	// Create the buffer for the generated documentation:
	g.buffer, err = markdown.NewBufferBuilder().
		Reporter(g.reporter).
		Output(g.output).
		File(fileName).
		Function("displayName", g.displayName).
		Function("resourceFile", g.fileName).
		Build()
	if err != nil {
		return err
	}

	// Generate the documentation:
	g.documentIndex()

	// Write the generated documentation:
	return g.buffer.Write()
}

func (g *DocsGenerator) documentIndex() {
	g.buffer.Emit(`
		# Model

		## Resources

		{{ range .Model.Services }}
			{{ range .Versions }}
				{{ range .Resources }}
					[{{ displayName . }}]({{ resourceFile . }}.md)
				{{ end }}
			{{ end }}
		{{ end }}

		## Types

		{{ range .Model.Services }}
			{{ range .Versions }}
				{{ range .Types }}
					[{{ displayName . }}]({{ resourceFile . }}.md)
				{{ end }}
			{{ end }}
		{{ end }}
		`,
		"Model", g.model,
	)
}

func (g *DocsGenerator) generateResource(resource *concepts.Resource) error {
	var err error

	// Calculate the file name:
	fileName := g.fileName(resource)

	// Create the buffer for the generated documentation:
	g.buffer, err = markdown.NewBufferBuilder().
		Reporter(g.reporter).
		Output(g.output).
		File(fileName).
		Function("displayName", g.displayName).
		Function("docDetail", g.docDetail).
		Function("docSummary", g.docSummary).
		Function("httpMethod", g.httpMethod).
		Function("tagName", g.tagName).
		Build()
	if err != nil {
		return err
	}

	// Generate the documentation:
	g.documentResource(resource)

	// Write the generated documentation:
	return g.buffer.Write()
}

func (g *DocsGenerator) documentResource(resource *concepts.Resource) {
	g.buffer.Emit(`
		# {{ displayName .Resource }}

		{{ docDetail .Resource }}

		{{ if .Resource.Methods }}
			| Name | Summary |
			| ---- | ------- |
			{{ range .Resource.Methods -}}
				| {{ displayName . }} | {{ docSummary . }} |
			{{ end -}}
		{{ end }}

		{{ range .Resource.Methods }}
			## {{ displayName . }} - {{ httpMethod . }}

			{{ docDetail . }}

			{{ if .Parameters }}
				| Name | Type | Direction | Summary |
				| ---- | ---- | --------- | ------- |
				{{ range .Parameters -}}
					| {{ tagName . }} | - | - | {{ docSummary . }} |
				{{ end -}}
			{{ end }}

			{{ range .Parameters }}
				### {{ displayName . }}

				{{ docDetail . }}
			{{ end }}
		{{ end }}
		`,
		"Resource", resource,
	)
}

func (g *DocsGenerator) generateType(typ *concepts.Type) error {
	var err error

	// Calculate the file name:
	fileName := g.fileName(typ)

	// Create the buffer for the generated documentation:
	g.buffer, err = markdown.NewBufferBuilder().
		Reporter(g.reporter).
		Output(g.output).
		File(fileName).
		Function("displayName", g.displayName).
		Function("docDetail", g.docDetail).
		Function("docSummary", g.docSummary).
		Function("tagName", g.tagName).
		Build()
	if err != nil {
		return err
	}

	// Generate the documentation:
	g.documentType(typ)

	// Write the generated documentation:
	return g.buffer.Write()
}

func (g *DocsGenerator) documentType(typ *concepts.Type) {
	switch {
	case typ.IsEnum():
		g.documentEnum(typ)
	case typ.IsStruct():
		g.documentStruct(typ)
	}
}

func (g *DocsGenerator) documentEnum(typ *concepts.Type) {
	g.buffer.Emit(`
		# {{ displayName .Type }}

		{{ docDetail .Type }}

		{{ if .Type.Values }}
			| Name | Summary |
			| ---- | ------- |
			{{ range .Type.Values -}}
				| {{ tagName . }} | {{ docSummary . }} |
			{{ end -}}
		{{ end }}

		{{ range .Type.Values }}
			### {{ tagName . }}

			{{ docDetail . }}
		{{ end }}
		`,
		"Type", typ,
	)
}

func (g *DocsGenerator) documentStruct(typ *concepts.Type) {
	g.buffer.Emit(`
		# {{ displayName .Type }}

		{{ docDetail .Type }}

		{{ if .Type.Attributes }}
			| Name | Type | Summary |
			| ---- | ---- | ------- |
			{{ if .Type.IsClass -}}
				| {{ backTicks "kind" }} | String | Name of the type of the object. |
				| {{ backTicks "id" }} | String | Unique identifier of the object. |
				| {{ backTicks "href" }} | String | Link to the object. |
			{{ end -}}
			{{ range .Type.Attributes -}}
				| {{ tagName . }} | - | {{ docSummary . }} |
			{{ end -}}
		{{ end }}

		{{ range .Type.Attributes }}
			## {{ tagName . }}

			{{ docDetail . }}
		{{ end }}
		`,
		"Type", typ,
	)
}

func (g *DocsGenerator) displayName(object names.Named) string {
	return g.names.Display(object.Name())
}

func (g *DocsGenerator) tagName(object names.Named) string {
	return g.names.Tag(object.Name())
}

func (g *DocsGenerator) fileName(object names.Named) string {
	name := object.Name()
	switch object.(type) {
	case *concepts.Type:
		name = names.Cat(name, nomenclator.Type)
	case *concepts.Resource:
		name = names.Cat(name, nomenclator.Resource)
	}
	return g.names.File(name)
}

func (g *DocsGenerator) httpMethod(method *concepts.Method) string {
	name := method.Name()
	if name.Equals(nomenclator.Get) || name.Equals(nomenclator.List) {
		return "GET"
	}
	if name.Equals(nomenclator.Add) {
		return "POST"
	}
	if name.Equals(nomenclator.Update) {
		return "PATCH"
	}
	if name.Equals(nomenclator.Delete) {
		return "DELETE"
	}
	return "POST"
}

func (g *DocsGenerator) docDetail(object concepts.Documented) string {
	return object.Doc()
}

func (g *DocsGenerator) docSummary(object concepts.Documented) string {
	// If there is no documentation then consider it empty:
	doc := object.Doc()
	if doc == "" {
		return ""
	}

	// The summary is the first sentence of the documentation, or the complete documentation if
	// there is no dot to end the first sentence.
	index := strings.IndexRune(doc, '.')
	if index != -1 {
		doc = doc[0 : index+1]
		doc = strings.ReplaceAll(doc, "\n", "")
	}

	return doc
}
