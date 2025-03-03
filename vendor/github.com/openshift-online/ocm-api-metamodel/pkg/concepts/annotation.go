/*
Copyright (c) 2022 Red Hat, Inc.

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

package concepts

import (
	"fmt"
)

// Annotation is the representation of the application of an annotation.
type Annotation struct {
	name       string
	parameters map[string]interface{}
}

// NewAnnotation creates a new annotation.
func NewAnnotation() *Annotation {
	return &Annotation{
		parameters: map[string]interface{}{},
	}
}

// Name returns the name of this annotation.
func (a *Annotation) Name() string {
	return a.name
}

// SetName sets the name of this annotation.
func (s *Annotation) SetName(value string) {
	s.name = value
}

// AddParameter adds a new parameter to the annotation with the given name and value.
func (a *Annotation) AddParameter(name string, value interface{}) {
	a.parameters[name] = value
}

// FindParameter returns the value of the parameter with the given name, or nil if no such parameter
// exists.
func (a *Annotation) FindParameter(name string) interface{} {
	return a.parameters[name]
}

// GetString returns the value of the given parameter converted to a string, or the empty string if
// no such parameter exists.
func (a *Annotation) GetString(name string) string {
	value, ok := a.parameters[name]
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s", value)
}
