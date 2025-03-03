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

package openapi

import (
	"fmt"
	"path/filepath"

	"github.com/openshift-online/ocm-api-metamodel/pkg/annotations"
	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// NamesCalculatorBuilder is an object used to configure and build the OpenAPI names calculators.
// Don't create instances directly, use the NewNamesCalculator function instead.
type NamesCalculatorBuilder struct {
	reporter *reporter.Reporter
}

// NamesCalculator is an object used to calculate OpenAPI names. Don't create instances directly,
// use the builder instead.
type NamesCalculator struct {
	reporter *reporter.Reporter
}

// NewNamesCalculator creates an OpenAPI names calculator builder.
func NewNamesCalculator() *NamesCalculatorBuilder {
	return &NamesCalculatorBuilder{}
}

// Reporter sets the object that will be used to report information about the calculation processes,
// including errors.
func (b *NamesCalculatorBuilder) Reporter(value *reporter.Reporter) *NamesCalculatorBuilder {
	b.reporter = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new
// calculator using it.
func (b *NamesCalculatorBuilder) Build() (calculator *NamesCalculator, err error) {
	// Check that the mandatory parameters have been provided:
	if b.reporter == nil {
		err = fmt.Errorf("reporter is mandatory")
		return
	}

	// Create the calculator:
	calculator = &NamesCalculator{
		reporter: b.reporter,
	}

	return
}

// FileName calculates the relative file path where the specification for the given service version
// should be written. For example if the service identifier is `clusters_mgmt` and the version
// identifiers is `v1` then the result will be `clusters_mgmt/v1/openapi.json`.
func (c *NamesCalculator) FileName(version *concepts.Version) string {
	service := version.Owner()
	return filepath.Join(service.Name().Snake(), version.Name().Snake(), "openapi.json")
}

// SchemaName calculates the schema name for the given type.
func (c *NamesCalculator) SchemaName(typ *concepts.Type) string {
	return typ.Name().Camel()
}

// AttributePropertyName calculates the property name for an attribute of a struct type.
func (c *NamesCalculator) AttributePropertyName(attribute *concepts.Attribute) string {
	name := annotations.JSONName(attribute)
	if name == "" {
		name = attribute.Name().Snake()
	}
	return name
}

// ParameterPropertyName calculates the property name for an parameter of a method.
func (c *NamesCalculator) ParameterPropertyName(parameter *concepts.Parameter) string {
	return parameter.Name().Snake()
}
