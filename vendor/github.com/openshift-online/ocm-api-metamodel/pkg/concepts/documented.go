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

package concepts

// Documented is implemented by concepts that have documentation.
type Documented interface {
	// Doc returns the documentation attached to the object.
	Doc() string

	// SetDoc sets the documentation attached to the object.
	SetDoc(doc string)
}

// documentedSupport is an implementation of the Documented interface intended to be embedded in
// types that need to implement that interface.
type documentedSupport struct {
	doc string
}

// Make sure we implement the interface:
var _ Documented = &documentedSupport{}

func (s *documentedSupport) Doc() string {
	return s.doc
}

func (s *documentedSupport) SetDoc(doc string) {
	s.doc = doc
}
