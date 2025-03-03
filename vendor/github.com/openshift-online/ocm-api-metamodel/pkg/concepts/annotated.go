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

// Annotated is implemented by concepts that have annotations.
type Annotated interface {
	// Annotations returns the annotations attached to the object.
	Annotations() []*Annotation

	// AddAnnotation adds an annotation to the object.
	AddAnnotation(annotation *Annotation)

	// FindAnnotation returns the annotation attached to the object that has the given name, or
	// nil if there is no such annotation.
	GetAnnotation(name string) *Annotation
}

// annotatedSupport is an implementation of the Annotated interface intended to be embedded in types
// that need to implement that interface.
type annotatedSupport struct {
	annotations []*Annotation
}

// Make sure we implement the interface:
var _ Annotated = &annotatedSupport{}

func (s *annotatedSupport) Annotations() []*Annotation {
	return s.annotations
}

func (s *annotatedSupport) AddAnnotation(annotation *Annotation) {
	s.annotations = append(s.annotations, annotation)
}

func (s *annotatedSupport) GetAnnotation(name string) *Annotation {
	for _, annotation := range s.annotations {
		if annotation.name == name {
			return annotation
		}
	}
	return nil
}
