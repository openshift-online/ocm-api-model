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

import "github.com/openshift-online/ocm-api-metamodel/pkg/names"

// namedSupport is an implementation of the names.Named interface intended to be embeded in other
// types that need to implement that interface.
type namedSupport struct {
	name *names.Name
}

// Make sure we implement the interface:
var _ names.Named = &namedSupport{}

func (s *namedSupport) Name() *names.Name {
	return s.name
}

func (s *namedSupport) SetName(value *names.Name) {
	s.name = value
}
