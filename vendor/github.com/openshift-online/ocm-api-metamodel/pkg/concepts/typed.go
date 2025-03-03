/*
Copyright (c) 2021 Red Hat, Inc.

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

// Typed is implemented by concepts that have a type.
type Typed interface {
	// Type returns the type of the object.
	Type() *Type

	// SetType sets the type of the object.
	SetType(typ *Type)
}

// typedSupport is an implementation of the Typed interface intended to be embedded in types that
// need to implement that interface.
type typedSupport struct {
	typ *Type
}

// Make sure we implement the interface:
var _ Typed = &typedSupport{}

func (s *typedSupport) Type() *Type {
	return s.typ
}

func (s *typedSupport) SetType(typ *Type) {
	s.typ = typ
}
