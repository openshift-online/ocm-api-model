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

package reporter

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
)

// Builder contains the data and logic needed to create a new reporter. Don't create instances of
// this type directly, use the New function instead.
type Builder struct {
	out io.Writer
	err io.Writer
}

// Reporter is the reported used by the metamodel tools. It prints the messages to the standard
// output stream. Don't create instances directly, use the New function instead.
type Reporter struct {
	out    io.Writer
	err    io.Writer
	errors int
}

// New creates a builder that can then be used to configurer and create a reporter.
func New() *Builder {
	return &Builder{
		out: os.Stdout,
		err: os.Stderr,
	}
}

// Streams sets the streams that the reporter will use to write messages. If not specified it will
// use os.Stdout for information messages and os.Stderr for error messages.
func (b *Builder) Streams(out, err io.Writer) *Builder {
	b.out = out
	b.err = err
	return b
}

// Build uses the configuration stored in the builder to create a new reporter.
func (b *Builder) Build() (result *Reporter, err error) {
	// Check the parameters:
	if b.out == nil {
		err = fmt.Errorf("output stream is mandatory")
		return
	}
	if b.err == nil {
		err = fmt.Errorf("error stream is mandatory")
		return
	}

	// Create and populate the object:
	result = &Reporter{
		out: b.out,
		err: b.err,
	}

	return
}

// Infof prints an informative message with the given format and arguments.
func (r *Reporter) Infof(format string, args ...interface{}) {
	message := r.printf(format, args)
	fmt.Fprintf(r.out, "%s%s\n", infoPrefix, message)
}

// Warnf prints an warning message with the given format and arguments.
func (r *Reporter) Warnf(format string, args ...interface{}) {
	message := r.printf(format, args)
	fmt.Fprintf(r.out, "%s%s\n", warnPrefix, message)
}

// Errorf prints an error message with the given format and arguments. It also return an error
// containing the same information, which will be usually discarded, except when the caller needs to
// report the error and also return it.
func (r *Reporter) Errorf(format string, args ...interface{}) error {
	message := r.printf(format, args)
	fmt.Fprintf(r.err, "%s%s\n", errorPrefix, message)
	r.errors++
	return errors.New(message)
}

func (r *Reporter) printf(format string, args []interface{}) string {
	// Replace arguments that are named objects or names with their camel case equivalent,
	// as that is what users type in the source files:
	for i, arg := range args {
		switch typed := arg.(type) {
		case *concepts.Service:
			args[i] = r.fqn(typed)
		case *concepts.Version:
			service := typed.Owner()
			args[i] = r.fqn(service, typed)
		case *concepts.Resource:
			version := typed.Owner()
			service := version.Owner()
			args[i] = r.fqn(service, version, typed)
		case *concepts.Type:
			if typed.IsScalar() || !typed.IsEnum() {
				args[i] = typed.Name().Camel()
			} else {
				version := typed.Owner()
				service := version.Owner()
				args[i] = r.fqn(service, version, typed)
			}
		case *concepts.Attribute:
			typ := typed.Owner()
			version := typ.Owner()
			service := version.Owner()
			args[i] = r.fqn(service, version, typ, typed)
		case *concepts.Method:
			resource := typed.Owner()
			version := resource.Owner()
			service := version.Owner()
			args[i] = r.fqn(service, version, resource, typed)
		case *concepts.Locator:
			resource := typed.Owner()
			version := resource.Owner()
			service := version.Owner()
			args[i] = r.fqn(service, version, resource, typed)
		case *concepts.Parameter:
			method := typed.Owner()
			resource := method.Owner()
			version := resource.Owner()
			service := version.Owner()
			args[i] = r.fqn(service, version, resource, method, typed)
		case names.Named:
			args[i] = r.fqn(typed)
		case *names.Name:
			args[i] = typed.Camel()
		}
	}

	// Format the message:
	return fmt.Sprintf(format, args...)
}

func (r *Reporter) fqn(nameds ...names.Named) string {
	segments := make([]string, len(nameds))
	for i, named := range nameds {
		name := named.Name()
		segments[i] = name.Camel()
	}
	return strings.Join(segments, ".")
}

// Errors returns the number of errors that have been reported via this reporter.
func (r *Reporter) Errors() int {
	return r.errors
}

// Message prefix using ANSI scape seequences to set colors:
const (
	infoPrefix  = "\033[0;32mI:\033[m "
	warnPrefix  = "\033[0;33mW:\033[m "
	errorPrefix = "\033[0;31mE:\033[m "
)
