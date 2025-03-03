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

package markdown

import (
	"fmt"
	"strings"

	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// NamesCalculatorBuilder is an object used to configure and build the Markdown names calculators.
// Don't create instances directly, use the NewNamesCalculator function instead.
type NamesCalculatorBuilder struct {
	reporter *reporter.Reporter
}

// NamesCalculator is an object used to calculate Markdown names. Don't create instances directly,
// use the builder instead.
type NamesCalculator struct {
	reporter *reporter.Reporter
}

// NewNamesCalculator creates an Markdown names calculator builder.
func NewNamesCalculator() *NamesCalculatorBuilder {
	builder := new(NamesCalculatorBuilder)
	return builder
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
		err = fmt.Errorf("parameter 'reporter' is mandatory")
		return
	}

	// Create the calculator:
	calculator = new(NamesCalculator)
	calculator.reporter = b.reporter

	return
}

// Display converts the given name into an string, following the rules for names displayed in the
// documentation.
func (c *NamesCalculator) Display(name *names.Name) string {
	words := name.Words()
	chunks := make([]string, len(words))
	for i, word := range words {
		chunks[i] = word.Capitalize()
	}
	return strings.Join(chunks, "")
}

// File converts the given name into an string, following the rules for Markdown files.
func (c *NamesCalculator) File(name *names.Name) string {
	words := name.Words()
	chunks := make([]string, len(words))
	for i, word := range words {
		chunks[i] = strings.ToLower(word.String())
	}
	return strings.Join(chunks, "_")
}

// Tag converts the given name into an string, following the rules for JSON field names.
func (c *NamesCalculator) Tag(name *names.Name) string {
	words := name.Words()
	chunks := make([]string, len(words))
	for i, word := range words {
		chunks[i] = strings.ToLower(word.String())
	}
	return "`" + strings.Join(chunks, "_") + "`"
}
