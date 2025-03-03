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

import (
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
)

// Error is the representation of a catagery of errors.
type Error struct {
	documentedSupport
	annotatedSupport
	namedSupport

	owner *Version
	name  *names.Name
	code  int
}

// NewError creates a new error.
func NewError() *Error {
	return new(Error)
}

// Owner returns the version that owns this error.
func (e *Error) Owner() *Version {
	return e.owner
}

// SetOwner sets the version that owns this error.
func (e *Error) SetOwner(version *Version) {
	e.owner = version
}

// Code returns the numeric code of this error.
func (e *Error) Code() int {
	return e.code
}

// SetCode sets the numeric code of this error.
func (e *Error) SetCode(value int) {
	e.code = value
}

// ErrorSlice is used to simplify sorting of slices of errors by name.
type ErrorSlice []*Error

func (s ErrorSlice) Len() int {
	return len(s)
}

func (s ErrorSlice) Less(i, j int) bool {
	return names.Compare(s[i].name, s[j].name) == -1
}

func (s ErrorSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
