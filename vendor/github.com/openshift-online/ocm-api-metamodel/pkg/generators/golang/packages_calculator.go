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

	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// PackagesCalculatorBuilder is an object used to configure and build the Go packages calculators.
// Don't create instances directly, use the NewPackagesCalculator function instead.
type PackagesCalculatorBuilder struct {
	reporter *reporter.Reporter
	base     string
}

// PackagesCalculator is an object used to calculate Go packages. Don't create instances directly,
// use the builder instead.
type PackagesCalculator struct {
	reporter *reporter.Reporter
	base     string
}

// NewPackagesCalculator creates a Go packages calculator builder.
func NewPackagesCalculator() *PackagesCalculatorBuilder {
	return &PackagesCalculatorBuilder{}
}

// Reporter sets the object that will be used to report information about the calculation processes,
// including errors.
func (b *PackagesCalculatorBuilder) Reporter(value *reporter.Reporter) *PackagesCalculatorBuilder {
	b.reporter = value
	return b
}

// Base sets the import path of the base package were the code will be generated.
func (b *PackagesCalculatorBuilder) Base(value string) *PackagesCalculatorBuilder {
	b.base = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new
// calculator using it.
func (b *PackagesCalculatorBuilder) Build() (calculator *PackagesCalculator, err error) {
	// Check that the mandatory parameters have been provided:
	if b.reporter == nil {
		err = fmt.Errorf("reporter is mandatory")
		return
	}
	if b.base == "" {
		err = fmt.Errorf("base package is mandatory")
		return
	}

	// Create the calculator:
	calculator = &PackagesCalculator{
		reporter: b.reporter,
		base:     b.base,
	}

	return
}

// ServicePackage returns the name of the package for the given service.
func (g *PackagesCalculator) ServicePackage(service *concepts.Service) string {
	return service.Name().LowerJoined("")
}

// ServiceImport returns the complete import path of the package for the given service.
func (g *PackagesCalculator) ServiceImport(service *concepts.Service) string {
	return path.Join(g.base, g.ServicePackage(service))
}

// ServiceSelector returns the selector of the package for the given service.
func (g *PackagesCalculator) ServiceSelector(service *concepts.Service) string {
	return path.Base(g.ServicePackage(service))
}

// VersionPackage returns the name of the package for the given version.
func (g *PackagesCalculator) VersionPackage(version *concepts.Version) string {
	return path.Join(
		version.Owner().Name().LowerJoined(""),
		version.Name().LowerJoined(""),
	)
}

// VersionImport returns the complete import path of the package for the given version.
func (g *PackagesCalculator) VersionImport(version *concepts.Version) string {
	return path.Join(g.base, g.VersionPackage(version))
}

// VersionSelector returns the selector of the package for the given service.
func (g *PackagesCalculator) VersionSelector(version *concepts.Version) string {
	return path.Base(g.VersionPackage(version))
}

// HelpersPackage returns the name of the helpers package.
func (g *PackagesCalculator) HelpersPackage() string {
	return nomenclator.Helpers.LowerJoined("")
}

// HelpersImport returns complete import path of the helpers package.
func (g *PackagesCalculator) HelpersImport() string {
	return path.Join(g.base, g.HelpersPackage())
}

// MetricsPackage returns the name of the metrics package.
func (g *PackagesCalculator) MetricsPackage() string {
	return nomenclator.Metrics.LowerJoined("")
}

// MetricsImport returns complete import path of the metrics package.
func (g *PackagesCalculator) MetricsImport() string {
	return path.Join(g.base, g.HelpersPackage())
}

// ErrorsPackage returns the name of the errors package.
func (g *PackagesCalculator) ErrorsPackage() string {
	return nomenclator.Errors.LowerJoined("")
}

// ErrorsImport returns complete import path of the errors package.
func (g *PackagesCalculator) ErrorsImport() string {
	return path.Join(g.base, g.ErrorsPackage())
}

// BasePackage returns the import path of the base package.
func (g *PackagesCalculator) BasePackage() string {
	return g.base
}
