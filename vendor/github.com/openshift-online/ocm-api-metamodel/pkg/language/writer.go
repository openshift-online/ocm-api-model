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

// This file contains the implementation of the object that knows how to write a model to a set of
// source files.

package language

import (
	"fmt"

	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// Writer is used to write a model to a set of files. Don't create objects of this type directly,
// use the NewWriter function instead.
type Writer struct {
	// Object used to report errors and other relevant information.
	reporter *reporter.Reporter

	// The model that is being saved.
	model *concepts.Model

	// Paths of the output directory where the files will be created.
	output string
}

// NewWriter creates an object that knows how to writea a model to a set set of files. Typical usage
// is to create the object, configure it and then call the Write method. For example, to write a
// previously loaded model to the 'mydir' directory, do the following:
//
//	err := io.NewWriter().
//		Model(model).
//		Target("mydir").
//		Write()
//
// Error checking will only happen when the Save method is called.
func NewWriter() *Writer {
	return new(Writer)
}

// Reporter sets the object that will be used to report information about the saving process,
// including errors.
func (w *Writer) Reporter(reporter *reporter.Reporter) *Writer {
	w.reporter = reporter
	return w
}

// Model sets the model that will be written.
func (w *Writer) Model(value *concepts.Model) *Writer {
	w.model = value
	return w
}

// Target sets the output directory.
func (w *Writer) Output(value string) *Writer {
	w.output = value
	return w
}

// Write writes the model.
func (w *Writer) Write() error {
	// Check the parameters:
	if w.reporter == nil {
		return fmt.Errorf("parameter 'reporter' is mandatory")
	}
	if w.model == nil {
		return fmt.Errorf("parameter 'model' is mandatory")
	}
	if w.output == "" {
		return fmt.Errorf("parameter 'output' is mandatory")
	}

	// Check if there are errors:
	errors := w.reporter.Errors()
	if errors > 0 {
		if errors > 1 {
			return fmt.Errorf("there were %d errors", errors)
		} else {
			return fmt.Errorf("there was 1 error")
		}
	}

	// Write types and resources:
	for _, service := range w.model.Services() {
		for _, version := range service.Versions() {
			for _, typ := range version.Types() {
				w.writeType(typ)
			}
			for _, resource := range version.Resources() {
				w.writeResource(resource)
			}
		}
	}

	return nil
}

func (w *Writer) writeType(typ *concepts.Type) {
	w.reporter.Infof("Writing type '%s'", typ.Name())
}

func (w *Writer) writeResource(resource *concepts.Resource) {
	w.reporter.Infof("Writing resource '%s'", resource.Name())
}
