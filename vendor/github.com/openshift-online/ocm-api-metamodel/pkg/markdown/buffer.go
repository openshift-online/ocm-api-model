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
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// BufferBuilder is used to create a new Markdown buffer. Don't create it directly, use the
// NewBuffer function instead.
type BufferBuilder struct {
	reporter  *reporter.Reporter
	output    string
	file      string
	functions map[string]interface{}
}

// Buffer is a type that simplifies the generation of Markdown code.
type Buffer struct {
	reporter  *reporter.Reporter
	output    string
	pkg       string
	file      string
	functions map[string]interface{}
	imports   map[string]interface{}
	code      *bytes.Buffer
}

// NewBuffer creates a builder for Markdown buffers.
func NewBufferBuilder() *BufferBuilder {
	return new(BufferBuilder)
}

// Reporter sets the object that will be used to report errors and other relevant information.
func (b *BufferBuilder) Reporter(reporter *reporter.Reporter) *BufferBuilder {
	b.reporter = reporter
	return b
}

// Output sets the output directory where the source code will be generated.
func (b *BufferBuilder) Output(value string) *BufferBuilder {
	b.output = value
	return b
}

// File sets the name of the file, without the .md extension.
func (b *BufferBuilder) File(value string) *BufferBuilder {
	b.file = value
	return b
}

// Function adds a function that can then be used in the templates.
func (b *BufferBuilder) Function(name string, function interface{}) *BufferBuilder {
	if b.functions == nil {
		b.functions = make(map[string]interface{})
	}
	b.functions[name] = function
	return b
}

// Build creates a new buffer using the configuration stored in the builder.
func (b *BufferBuilder) Build() (buffer *Buffer, err error) {
	// Check that the mandatory parameters have been provided:
	if b.reporter == nil {
		err = fmt.Errorf("reporter is mandatory")
		return
	}
	if b.output == "" {
		err = fmt.Errorf("output is mandatory")
		return
	}
	if b.file == "" {
		err = fmt.Errorf("file is mandatory")
		return
	}

	// Allocate and populate the buffer:
	buffer = new(Buffer)
	buffer.reporter = b.reporter
	buffer.file = filepath.Join(b.output, b.file+".md")
	buffer.functions = make(map[string]interface{})
	buffer.functions["backTicks"] = buffer.backTicks
	for name, function := range b.functions {
		buffer.functions[name] = function
	}
	buffer.code = new(bytes.Buffer)

	return
}

// Emit writes to the code buffer, using the given template and arguments. The syntax of the
// template is the one used by the text/template package. The arguments should be a set nave/value
// pairs. Names should be strings, and values can be anything. These names and values will be put in
// a map that will then be the data object used to execute the template. For example, a template
// that generates the code of a Go method could be used like this:
//
//	// Calculate the name of the method and of the type:
//	typeName := ...
//	methodName := ...
//
//	// Generate the code of the method:
//	buffer.Emit(`
//		func (c *{{ .type }}) {{ .method }} {
//			...
//		}
//		`,
//		"type", typeName,
//		"method", methodName,
//	)
//
// All the facilities defined in the text/template package will also be available in these
// templates.
func (b *Buffer) Emit(tmpl string, args ...interface{}) {
	// Check that there is an even number of args, and that the first of each pair is
	// an string:
	count := len(args)
	if count%2 != 0 {
		b.reporter.Errorf(
			"Template '%s' should have an even number of arguments, but it has %d",
			tmpl, count,
		)
		return
	}
	for i := 0; i < count; i = i + 2 {
		name := args[i]
		_, ok := name.(string)
		if !ok {
			b.reporter.Errorf(
				"Argument %d of template '%s' is a key, so it should be a string, "+
					"but its type is %T",
				i, tmpl, name,
			)
		}
	}
	if b.reporter.Errors() > 0 {
		return
	}

	// Put the variables in the map that will be passed as the data object for the execution of
	// the template:
	data := make(map[string]interface{})
	for i := 0; i < count; i = i + 2 {
		name := args[i].(string)
		value := args[i+1]
		data[name] = value
	}

	// Parse the template:
	obj, err := template.New("").
		Funcs(b.functions).
		Parse(tmpl)
	if err != nil {
		b.reporter.Errorf("Can't parse template '%s': %v", tmpl, err)
		return
	}

	// Execute the template:
	buffer := new(bytes.Buffer)
	err = obj.Execute(buffer, data)
	if err != nil {
		b.reporter.Errorf("Can't execute template '%s': %v", tmpl, err)
		return
	}

	// Add the generated text to the code:
	text := buffer.String()
	_, err = b.code.WriteString(text)
	if err != nil {
		b.reporter.Errorf("Can't add generated text '%s' to the buffer: %v", text, err)
		return
	}

	// Add a blank line:
	b.code.WriteString("\n")
}

// Write creates the output file and writes the generated content.
func (b *Buffer) Write() error {
	var err error

	// Inform that we are writing the file:
	b.reporter.Infof("Writing file '%s'", b.file)

	// Check if there were errors:
	errors := b.reporter.Errors()
	if errors > 0 {
		if errors > 1 {
			err = fmt.Errorf("there were %d errors", errors)
		} else {
			err = fmt.Errorf("there was 1 error")
		}
		return err
	}

	// Make sure that the output directory exists:
	outputDir := filepath.Dir(b.file)
	err = os.MkdirAll(outputDir, 0777)
	if err != nil {
		return fmt.Errorf("can't create output directory '%s': %w", outputDir, err)
	}

	// Open the file:
	outputFd, err := os.OpenFile(b.file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("can't open output file '%s': %w", b.file, err)
	}

	// Remove the leading tabs:
	b.removeTabs()

	// Write the code:
	_, err = b.code.WriteTo(outputFd)
	if err != nil {
		return fmt.Errorf(
			"can't write generated documentation to file '%s': %v",
			b.file, err,
		)
	}

	// Close the file:
	err = outputFd.Close()
	if err != nil {
		return fmt.Errorf("can't close output file '%s': %w", b.file, err)
	}

	return nil
}

// removeTabbs removes the tabs that are used in the templates to align them.
func (b *Buffer) removeTabs() error {
	// Split into lines:
	var lines []string
	scanner := bufio.NewScanner(b.code)
	for scanner.Scan() {
		line := strings.TrimRightFunc(scanner.Text(), unicode.IsSpace)
		lines = append(lines, line)
	}
	err := scanner.Err()
	if err != nil {
		return err
	}

	// Remove the leading tabs from each line:
	for i := range lines {
		lines[i] = strings.TrimLeftFunc(lines[i], func(r rune) bool {
			return r == '\t'
		})
	}

	// Save the modified lines to the buffer:
	b.code.Reset()
	for _, line := range lines {
		fmt.Fprintln(b.code, line)
	}

	return nil
}

// backTicks is a function that will be available in templates in order to make usage of back ticks
// simpler together with the special meaning that they have in the Go language.
func (b *Buffer) backTicks(value interface{}) string {
	return fmt.Sprintf("`%s`", value)
}
