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

// This file contains the error listener that will be used to report the errors detected while
// scanning and parsing the model files.

package language

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// ErrorListenerBuilder is used to build error listeners.
type ErrorListenerBuilder struct {
	reporter *reporter.Reporter
}

// ErrorListener is used to report errors detected while scanning and parsing the model files.
type ErrorListener struct {
	reporter *reporter.Reporter
}

// NewErrorListener creates a new builder of error listeners.
func NewErrorListener() *ErrorListenerBuilder {
	return new(ErrorListenerBuilder)
}

// Reporter sets the object that will be used by the created error listeners to report errors.
func (b *ErrorListenerBuilder) Reporter(value *reporter.Reporter) *ErrorListenerBuilder {
	b.reporter = value
	return b
}

// Builde creates a new error listener using the configuration stored in the builder.
func (b *ErrorListenerBuilder) Build() (listener *ErrorListener, err error) {
	// Check mandatory parameters:
	if b.reporter == nil {
		err = fmt.Errorf("reporter is mandatory")
		return
	}

	// Create and populate the object:
	listener = new(ErrorListener)
	listener.reporter = b.reporter

	return
}

// Syntax error is callend when the ANTLR runtime needs to report an error detected by the scanner
// or by the parser of the model language.
func (l *ErrorListener) SyntaxError(parser antlr.Recognizer, symbol interface{}, line, column int,
	msg string, exception antlr.RecognitionException) {
	l.reporter.Errorf("Line %d, column %d: %s", line, column, msg)
}

func (l *ErrorListener) ReportAmbiguity(parser antlr.Parser, dfa *antlr.DFA, start, stop int,
	exact bool, alts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// This method is used to report potential issues in the grammar, and we don't want to
	// do that during runtime.
}

func (l *ErrorListener) ReportAttemptingFullContext(parser antlr.Parser, dfa *antlr.DFA, start,
	stop int, alts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// This method is used to report potential issues in the grammar, and we don't want to
	// do that during runtime.
}

func (l *ErrorListener) ReportContextSensitivity(parser antlr.Parser, dfa *antlr.DFA, start, stop,
	prediction int, configs antlr.ATNConfigSet) {
	// This method is used to report potential issues in the grammar, and we don't want to
	// do that during runtime.
}
