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

	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// HelpersGeneratorBuilder is an object used to configure and build a helpers generator. Don't
// create instances directly, use the NewHelpersGenerator function instead.
type HelpersGeneratorBuilder struct {
	reporter *reporter.Reporter
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
}

// HelpersGenerator generates helper code. Don't create instances directly, use the builder instead.
type HelpersGenerator struct {
	reporter *reporter.Reporter
	errors   int
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	buffer   *Buffer
}

// NewHelpersGenerator creates a new builder for helpers generators.
func NewHelpersGenerator() *HelpersGeneratorBuilder {
	return new(HelpersGeneratorBuilder)
}

// Reporter sets the object that will be used to report information about the generation process,
// including errors.
func (b *HelpersGeneratorBuilder) Reporter(value *reporter.Reporter) *HelpersGeneratorBuilder {
	b.reporter = value
	return b
}

// Model sets the model that will be used by the client generator.
func (b *HelpersGeneratorBuilder) Model(value *concepts.Model) *HelpersGeneratorBuilder {
	b.model = value
	return b
}

// Output sets the output directory.
func (b *HelpersGeneratorBuilder) Output(value string) *HelpersGeneratorBuilder {
	b.output = value
	return b
}

// Packages sets the object that will be used to calculate package names.
func (b *HelpersGeneratorBuilder) Packages(
	value *PackagesCalculator) *HelpersGeneratorBuilder {
	b.packages = value
	return b
}

// Names sets the object that will be used to calculate names.
func (b *HelpersGeneratorBuilder) Names(value *NamesCalculator) *HelpersGeneratorBuilder {
	b.names = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new client
// generator using it.
func (b *HelpersGeneratorBuilder) Build() (generator *HelpersGenerator, err error) {
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
		err = fmt.Errorf("path is mandatory")
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
	generator = &HelpersGenerator{
		reporter: b.reporter,
		model:    b.model,
		output:   b.output,
		packages: b.packages,
		names:    b.names,
	}

	return
}

// Run executes the code generator.
func (g *HelpersGenerator) Run() error {
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
	g.buffer.Import("context", "")
	g.buffer.Import("fmt", "")
	g.buffer.Import("net/http", "")
	g.buffer.Import("net/url", "")
	g.buffer.Import("strings", "")
	g.buffer.Import("time", "")
	g.buffer.Emit(`
		// AddValue creates the given set of query parameters if needed, an then adds
		// the given parameter.
		func AddValue(query *url.Values, name string, value interface{}) {
			if *query == nil {
				*query = make(url.Values)
			}
			var text string
			switch typed := value.(type) {
			case time.Time:
				text = typed.UTC().Format(time.RFC3339)
			default:
				text = fmt.Sprintf("%v", value)
			}
			query.Add(name, text)
		}

		// CopyQuery creates a copy of the given set of query parameters.
		func CopyQuery(query url.Values) url.Values {
			if query == nil {
				return nil
			}
			result := make(url.Values)
			for name, values := range query {
				result[name] = CopyValues(values)
			}
			return result
		}

		// AddHeader creates the given set of headers if needed, and then adds the given
		// header:
		func AddHeader(header *http.Header, name string, value interface{}) {
			if *header == nil {
				*header = make(http.Header)
			}
			header.Add(name, fmt.Sprintf("%v", value))
		}

		// CopyHeader creates a copy of the given set of headers.
		func CopyHeader(header http.Header) http.Header {
			result := make(http.Header)
			for name, values := range header {
				result[name] = CopyValues(values)
			}
			return result
		}

		const impersonateUserHeader = "Impersonate-User"
		func AddImpersonationHeader(header *http.Header, user string) {
			AddHeader(header, impersonateUserHeader, user)
		}

		// CopyValues copies a slice of strings.
		func CopyValues(values []string) []string {
			if values == nil {
				return nil
			}
			result := make([]string, len(values))
			copy(result, values)
			return result
		}

		// Segments calculates the path segments for the given path.
		func Segments(path string) []string {
			for strings.HasPrefix(path, "/") {
				path = path[1:]
			}
			for strings.HasSuffix(path, "/") {
				path = path[0:len(path)-1]
			}
			return strings.Split(path, "/")
		}

		// PollContext repeatedly executes a task till it returns one of the given statuses and till the result
		// satisfies all the given predicates.
		func PollContext(
			ctx context.Context,
			interval time.Duration,
			statuses []int,
			predicates []func(interface{}) bool,
			task func(context.Context) (int, interface{}, error),
		) (result interface{}, err error) {
			// Check the deadline:
			deadline, ok := ctx.Deadline()
			if !ok {
				err = fmt.Errorf("context deadline is mandatory")
				return
			}

			// Check the interval:
			if interval <= 0 {
				err = fmt.Errorf("interval must be greater than zero")
				return
			}

			// Create a cancellable context so that we can explicitly cancel it when we know that the next
			// iteration of the loop will be after the deadline:
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			// If no expected status has been explicitly specified then add the default:
			if len(statuses) == 0 {
				statuses = []int{http.StatusOK}
			}

			for {
				// Execute the task. If this produces an error and the status code is zero it means that
				// there was an error like a timeout, or a low level communications problem. In that
				// case we want to immediately stop waiting.
				var status int
				status, result, err = task(ctx)
				if err != nil && status == 0 {
					break
				}

				// Evaluate the status and the predicates:
				statusOK := evalStatus(statuses, status)
				predicatesOK := evalPredicates(predicates, result)
				if statusOK && predicatesOK {
					break
				}

				// If either the status or the predicates aren't acceptable then we need to check if we
				// have enough time for another iteration before the deadline:
				if time.Now().Add(interval).After(deadline) {
					cancel()
					break
				}
				time.Sleep(interval)
			}

			return
		}

		// evalStatus checks if the actual status is one of the expected ones.
		func evalStatus(expected []int, actual int) bool {
			for _, current := range expected {
				if actual == current {
					return true
				}
			}
			return false
		}

		// evalPredicates checks if the object satisfies all the predicates.
		func evalPredicates(predicates []func(interface{}) bool, object interface{}) bool {
			if len(predicates) > 0 && object == nil {
				return false
			}
			for _, predicate := range predicates {
				if !predicate(object) {
					return false
				}
			}
			return true
		}
        `)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *HelpersGenerator) helpersFile() string {
	return g.names.File(nomenclator.Helpers)
}
