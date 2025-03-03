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

package names

// Named is implemented by concepts that have a name.
type Named interface {
	Name() *Name
}

// namedSupport is an implementation of the Named interface intended to be embeded in other types
// that need to implement that interface.
type namedSupport struct {
	name *Name
}

func (s *namedSupport) Name() *Name {
	return s.name
}

func (s *namedSupport) SetName(value *Name) {
	s.name = value
}
