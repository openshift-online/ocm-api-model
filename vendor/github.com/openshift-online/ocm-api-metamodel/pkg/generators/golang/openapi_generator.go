/*
Copyright (c) 2020 Red Hat, Inc.

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

// This file contains the generator that creates the Go source files containing the OpenAPI
// specifications of the services.

package golang

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/generators/openapi"
	"github.com/openshift-online/ocm-api-metamodel/pkg/http"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// OpenAPIGeneratorBuilder is an object used to configure and build a OpenAPI specification
// generators. Don't create instances directly, use the NewOpenAPIGenerator function instead.
type OpenAPIGeneratorBuilder struct {
	reporter *reporter.Reporter
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	binding  *http.BindingCalculator
	names    *openapi.NamesCalculator
}

// OpenAPIGenerator generates helper code. Don't create instances directly, use the builder instead.
type OpenAPIGenerator struct {
	reporter *reporter.Reporter
	errors   int
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	binding  *http.BindingCalculator
	names    *openapi.NamesCalculator
	buffer   *Buffer
}

// NewOpenAPIGenerator creates a new builder for helpers generators.
func NewOpenAPIGenerator() *OpenAPIGeneratorBuilder {
	return new(OpenAPIGeneratorBuilder)
}

// Reporter sets the object that will be used to report information about the generation process,
// including errors.
func (b *OpenAPIGeneratorBuilder) Reporter(value *reporter.Reporter) *OpenAPIGeneratorBuilder {
	b.reporter = value
	return b
}

// Model sets the model that will be used by the client generator.
func (b *OpenAPIGeneratorBuilder) Model(value *concepts.Model) *OpenAPIGeneratorBuilder {
	b.model = value
	return b
}

// Output sets the output directory.
func (b *OpenAPIGeneratorBuilder) Output(value string) *OpenAPIGeneratorBuilder {
	b.output = value
	return b
}

// Packages sets the object that will be used to calculate package names.
func (b *OpenAPIGeneratorBuilder) Packages(
	value *PackagesCalculator) *OpenAPIGeneratorBuilder {
	b.packages = value
	return b
}

// Names sets the object that will be used to calculate names.
func (b *OpenAPIGeneratorBuilder) Names(value *openapi.NamesCalculator) *OpenAPIGeneratorBuilder {
	b.names = value
	return b
}

// Binding sets the object that will by used to do HTTP binding calculations.
func (b *OpenAPIGeneratorBuilder) Binding(value *http.BindingCalculator) *OpenAPIGeneratorBuilder {
	b.binding = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new client
// generator using it.
func (b *OpenAPIGeneratorBuilder) Build() (generator *OpenAPIGenerator, err error) {
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
		err = fmt.Errorf("output path is mandatory")
		return
	}
	if b.packages == nil {
		err = fmt.Errorf("packages calculator is mandatory")
		return
	}
	if b.binding == nil {
		err = fmt.Errorf("HTTP binding calculator is mandatory")
		return
	}
	if b.names == nil {
		err = fmt.Errorf("names calculator is mandatory")
		return
	}

	// Create the generator:
	generator = &OpenAPIGenerator{
		reporter: b.reporter,
		model:    b.model,
		output:   b.output,
		packages: b.packages,
		binding:  b.binding,
		names:    b.names,
	}

	return
}

// Run executes the code generator.
func (g *OpenAPIGenerator) Run() error {
	// Create a temporary directory to write the OpenAPI JSON to, and remember to delete it:
	jsonDir, err := os.MkdirTemp("", "openapi")
	if err != nil {
		return g.reporter.Errorf("Can't create temporary OpenAPI directory: %v", err)
	}
	defer func() {
		err := os.RemoveAll(jsonDir)
		if err != nil {
			g.reporter.Errorf(
				"Can't remove temporary OpenAPI directory '%s': %v",
				jsonDir, err,
			)
		}
	}()

	// Geenrate the OpenAPI specification in the temporary directory:
	generator, err := openapi.NewOpenAPIGenerator().
		Reporter(g.reporter).
		Model(g.model).
		Output(jsonDir).
		Names(g.names).
		Binding(g.binding).
		Build()
	if err != nil {
		return g.reporter.Errorf("Can't create OpenAPI generator: %v", err)
	}
	err = generator.Run()
	if err != nil {
		return g.reporter.Errorf("OpenAPI generation failed: %v", err)
	}

	// Generate the Go source file containing the OpenAPI specification for each version:
	for _, service := range g.model.Services() {
		for _, version := range service.Versions() {
			err = g.generateSpec(version, jsonDir)
			if err != nil {
				return err
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

func (g *OpenAPIGenerator) generateSpec(version *concepts.Version, jsonDir string) error {
	var err error

	// Read the JSON file that contains the previously generated OpenAPI specification:
	jsonFile := filepath.Join(jsonDir, g.names.FileName(version))
	jsonData, err := os.ReadFile(jsonFile)
	if err != nil {
		return g.reporter.Errorf(
			"Can't read JSON file '%s' for version '%s': %v",
			jsonFile, version, err,
		)
	}

	// Calculate the name of the Go package:
	pkgName := g.packages.VersionPackage(version)

	// Create the buffer:
	g.buffer, err = NewBuffer().
		Reporter(g.reporter).
		Output(g.output).
		Packages(g.packages).
		Package(pkgName).
		File("openapi").
		Build()
	if err != nil {
		return err
	}

	// Generate the Go source:
	g.buffer.Emit(`
		// OpenAPI contains the OpenAPI specification of the service in JSON.
		var OpenAPI = []byte{
			{{ byteArray .Data }}
		}
		`,
		"Data", jsonData,
	)

	// Write the generated code:
	return g.buffer.Write()
}
