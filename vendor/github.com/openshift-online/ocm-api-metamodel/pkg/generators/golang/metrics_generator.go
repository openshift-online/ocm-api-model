/*
Copyright (c) 2021 Red Hat, Inc.

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
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/http"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// MetricsSupportGeneratorBuilder is an object used to configure and build the Metrics support
// generator. Don't create instances directly, use the NewMetricsSupportGenerator function instead.
type MetricsSupportGeneratorBuilder struct {
	reporter *reporter.Reporter
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	binding  *http.BindingCalculator
}

// MetricsSupportGenerator generates metrics support code. Don't create instances directly, use the
// builder instead.
type MetricsSupportGenerator struct {
	reporter *reporter.Reporter
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	buffer   *Buffer
	binding  *http.BindingCalculator
}

// NewMetricsSupportGenerator creates a new builder metrics support code generator.
func NewMetricsSupportGenerator() *MetricsSupportGeneratorBuilder {
	return &MetricsSupportGeneratorBuilder{}
}

// Reporter sets the object that will be used to report information about the generation process,
// including errors.
func (b *MetricsSupportGeneratorBuilder) Reporter(
	value *reporter.Reporter) *MetricsSupportGeneratorBuilder {
	b.reporter = value
	return b
}

// Model sets the model that will be used by the generator.
func (b *MetricsSupportGeneratorBuilder) Model(value *concepts.Model) *MetricsSupportGeneratorBuilder {
	b.model = value
	return b
}

// Output sets import path of the output package.
func (b *MetricsSupportGeneratorBuilder) Output(value string) *MetricsSupportGeneratorBuilder {
	b.output = value
	return b
}

// Packages sets the object that will be used to calculate package names.
func (b *MetricsSupportGeneratorBuilder) Packages(
	value *PackagesCalculator) *MetricsSupportGeneratorBuilder {
	b.packages = value
	return b
}

// Names sets the object that will be used to calculate names.
func (b *MetricsSupportGeneratorBuilder) Names(value *NamesCalculator) *MetricsSupportGeneratorBuilder {
	b.names = value
	return b
}

// Binding sets the object that will by used to do HTTP binding calculations.
func (b *MetricsSupportGeneratorBuilder) Binding(
	value *http.BindingCalculator) *MetricsSupportGeneratorBuilder {
	b.binding = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new types
// generator using it.
func (b *MetricsSupportGeneratorBuilder) Build() (generator *MetricsSupportGenerator, err error) {
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

	// Create the generator:
	generator = &MetricsSupportGenerator{
		reporter: b.reporter,
		model:    b.model,
		output:   b.output,
		packages: b.packages,
		names:    b.names,
	}

	return
}

// Run executes the code generator.
func (g *MetricsSupportGenerator) Run() error {
	var err error

	// Calculate the package and file name:
	pkgName := g.packages.MetricsPackage()
	fileName := g.names.File(names.Cat(nomenclator.Path, nomenclator.Tree, nomenclator.Data))

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

	// Build the tree and serialize it to JSON:
	apiBranch := map[string]interface{}{}
	for _, service := range g.model.Services() {
		serviceBranch := map[string]interface{}{}
		for _, version := range service.Versions() {
			versionBranch := g.resourceBranch(version.Root())
			versionLabel := g.binding.VersionSegment(version)
			serviceBranch[versionLabel] = versionBranch
		}
		serviceLabel := g.binding.ServiceSegment(service)
		apiBranch[serviceLabel] = serviceBranch
	}
	mainBranch := map[string]interface{}{
		"api": apiBranch,
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(mainBranch)
	if err != nil {
		return err
	}
	data := buffer.String()

	// Generate the code:
	g.buffer.Emit(`
		// pathTreeData is the JSON representation of the tree of URL paths.
		var pathTreeData = {{ backtick }}{{ .Data }}{{ backtick }}
		`,
		"Data", data,
	)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *MetricsSupportGenerator) resourceBranch(resource *concepts.Resource) interface{} {
	// Check if the target resource has any locator of action method. If it hasn't it does't
	// need an additional branch.
	locators := resource.Locators()
	var actions []*concepts.Method
	for _, method := range resource.Methods() {
		if method.IsAction() {
			actions = append(actions, method)
		}
	}
	if len(locators) == 0 && len(actions) == 0 {
		return nil
	}

	// If the target has locators or action methods then populate the branch:
	branch := map[string]interface{}{}
	for _, action := range actions {
		actionLabel := g.binding.MethodSegment(action)
		branch[actionLabel] = nil
	}
	for _, locator := range locators {
		locatorLabel := g.binding.LocatorSegment(locator)
		if locator.Variable() {
			locatorLabel = "-"
		}
		locatorBranch := g.resourceBranch(locator.Target())
		branch[locatorLabel] = locatorBranch
	}
	return branch
}
