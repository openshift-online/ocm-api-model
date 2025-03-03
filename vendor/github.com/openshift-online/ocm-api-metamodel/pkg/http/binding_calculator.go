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

// This file contains an object that implements the calculations related to the HTTP binding of the
// model. For example, it calculates which parameters of a method go in the query and which go in
// the request body.

package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/openshift-online/ocm-api-metamodel/pkg/annotations"
	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// BindingCalculatorBuilder is an object used to configure and build the HTTP bindings calculator.
// Don't create instances directly, use the NewBindingCalculator function instead.
type BindingCalculatorBuilder struct {
	reporter *reporter.Reporter
}

// BindingCalculator is an object used to do calculations related to the HTTP binding. Don't create
// instances directly, use the builder instead.
type BindingCalculator struct {
	reporter *reporter.Reporter
}

// NewBindingCalculator creates new object used to do calculations related to the HTTP binding.
func NewBindingCalculator() *BindingCalculatorBuilder {
	return &BindingCalculatorBuilder{}
}

// Reporter sets the object that will be used to report information about the calculation processes,
// including errors.
func (b *BindingCalculatorBuilder) Reporter(value *reporter.Reporter) *BindingCalculatorBuilder {
	b.reporter = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new
// calculator using it.
func (b *BindingCalculatorBuilder) Build() (calculator *BindingCalculator, err error) {
	// Check that the mandatory parameters have been provided:
	if b.reporter == nil {
		err = fmt.Errorf("reporter is mandatory")
		return
	}

	// Create the calculator:
	calculator = &BindingCalculator{
		reporter: b.reporter,
	}

	return
}

// RequestParameter returns the parameters of the given method that will be part of the HTTP
// request.
func (c *BindingCalculator) RequestParameters(method *concepts.Method) []*concepts.Parameter {
	var result []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.In() {
			result = append(result, parameter)
		}
	}
	return result
}

// RequestQueryParameters returns the parameters of the given method that should be placed in the
// HTTP request query.
func (c *BindingCalculator) RequestQueryParameters(method *concepts.Method) []*concepts.Parameter {
	var result []*concepts.Parameter
	if !method.IsAction() {
		for _, parameter := range method.Parameters() {
			if parameter.In() && parameter.Type().IsScalar() {
				result = append(result, parameter)
			}
		}
	}
	return result
}

// RequestBodyParameters returns the parameters of the given method that should be placed in the
// HTTP request body.
func (c *BindingCalculator) RequestBodyParameters(method *concepts.Method) []*concepts.Parameter {
	var result []*concepts.Parameter
	if method.IsAction() {
		for _, parameter := range method.Parameters() {
			if parameter.In() {
				result = append(result, parameter)
			}
		}
	} else {
		for _, parameter := range method.Parameters() {
			if parameter.In() && !parameter.Type().IsScalar() {
				result = append(result, parameter)
			}
		}
	}
	return result
}

// ResponseParameters returns the parameters of the given method that should be placed in the HTTP
// response.
func (c *BindingCalculator) ResponseParameters(method *concepts.Method) []*concepts.Parameter {
	var result []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.Out() {
			result = append(result, parameter)
		}
	}
	return result
}

// ResponseBodyParameters returns the parameters of the given method that should be placed in the
// HTTP response body.
func (c *BindingCalculator) ResponseBodyParameters(method *concepts.Method) []*concepts.Parameter {
	var result []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.Out() {
			result = append(result, parameter)
		}
	}
	return result
}

// Method returns the HTTP method corresponding to the given model method.
func (c *BindingCalculator) Method(method *concepts.Method) string {
	name := method.Name()
	switch {
	case name.Equals(nomenclator.Add):
		return http.MethodPost
	case name.Equals(nomenclator.Delete):
		return http.MethodDelete
	case name.Equals(nomenclator.Get):
		return http.MethodGet
	case name.Equals(nomenclator.List):
		return http.MethodGet
	case name.Equals(nomenclator.Post):
		return http.MethodPost
	case name.Equals(nomenclator.Update):
		return http.MethodPatch
	default:
		return http.MethodPost
	}
}

// Default status returns the HTTP status code that should be returned by default by the given
// method when there are no errors.
func (c *BindingCalculator) DefaultStatus(method *concepts.Method) string {
	name := method.Name()
	status := http.StatusOK
	switch {
	case name.Equals(nomenclator.Post):
		status = http.StatusCreated
	case name.Equals(nomenclator.Add):
		status = http.StatusCreated
	case name.Equals(nomenclator.Update):
		status = http.StatusOK
	case name.Equals(nomenclator.Delete):
		status = http.StatusNoContent
	}
	return strconv.Itoa(status)
}

// AttributeName returns the field name corresponding to the given model attribute.
func (c *BindingCalculator) AttributeName(attribute *concepts.Attribute) string {
	name := annotations.JSONName(attribute)
	if name == "" {
		name = attribute.Name().Snake()
	}
	return name
}

// QueryParameterName returns the name of the query parameter corresponding to the given model
// method parameter.
func (c *BindingCalculator) QueryParameterName(parameter *concepts.Parameter) string {
	name := annotations.HTTPName(parameter)
	if name == "" {
		// We check also the JSON annotation here for "bug compatibility". In the past we
		// used to check only the JSON name, but that is incorrect because this is about
		// HTTP query parameters, not JSON field names. But some models (the 'dryRun'
		// parameter of the accounts management service, for example) are already using that
		// bug as it was a feature.
		name = annotations.JSONName(parameter)
		if name == "" {
			name = parameter.Name().Snake()
		}
	}
	return name
}

// BodyParameterName returns the name of the body field corresponding to the given  model method
// parameter.
func (c *BindingCalculator) BodyParameterName(parameter *concepts.Parameter) string {
	name := annotations.JSONName(parameter)
	if name == "" {
		name = parameter.Name().Snake()
	}
	return name
}

// ServiceSegment calculates the URL segment corresponding to the given service.
func (c *BindingCalculator) ServiceSegment(service *concepts.Service) string {
	name := annotations.HTTPName(service)
	if name == "" {
		name = service.Name().Snake()
	}
	return name
}

// VersionSegment calculates the URL segment corresponding to the given version.
func (c *BindingCalculator) VersionSegment(version *concepts.Version) string {
	name := annotations.HTTPName(version)
	if name == "" {
		name = version.Name().Snake()
	}
	return name
}

// LocatorSegment calculates the URL segment corresponding to the given method.
func (c *BindingCalculator) MethodSegment(method *concepts.Method) string {
	if method.IsAction() {
		name := annotations.HTTPName(method)
		if name == "" {
			name = method.Name().Snake()
		}
		return name
	}
	return ""
}

// LocatorSegment calculates the URL segment corresponding to the given locator.
func (c *BindingCalculator) LocatorSegment(locator *concepts.Locator) string {
	name := annotations.HTTPName(locator)
	if name == "" {
		name = locator.Name().Snake()
	}
	return name
}

// EnumValueName returns the name corresponding to a value of an enumerated type.
func (c *BindingCalculator) EnumValueName(value *concepts.EnumValue) string {
	name := annotations.JSONName(value)
	if name == "" {
		name = value.Name().Snake()
	}
	return name
}
