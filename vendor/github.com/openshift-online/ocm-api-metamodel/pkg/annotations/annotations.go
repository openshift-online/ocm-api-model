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

package annotations

import (
	"fmt"

	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
)

// HTTPName checks if the given concept has a `http` annotation. If it does then it returns the
// value of the `name` parameter. If it doesn't, it returns an empty string.
func HTTPName(concept concepts.Annotated) string {
	annotation := concept.GetAnnotation("http")
	if annotation == nil {
		return ""
	}
	name := annotation.FindParameter("name")
	if name == nil {
		return ""
	}
	return fmt.Sprintf("%s", name)
}

// JSONName checks if the given concept has a `json` annotation. If it does then it returns the
// value of the `name` parameter. If it doesn't, it returns an empty string.
func JSONName(concept concepts.Annotated) string {
	annotation := concept.GetAnnotation("json")
	if annotation == nil {
		return ""
	}
	name := annotation.FindParameter("name")
	if name == nil {
		return ""
	}
	return fmt.Sprintf("%s", name)
}

// GoName checks if the given concept as a `go` annotation. If it has it then it returns the value
// of the `name` parameter. It returns an empty string if there is no such annotation or parameter.
func GoName(concept concepts.Annotated) string {
	annotation := concept.GetAnnotation("go")
	if annotation == nil {
		return ""
	}
	name := annotation.FindParameter("name")
	if name == nil {
		return ""
	}
	return fmt.Sprintf("%s", name)
}

// Reference checks if the given concept has a `reference` annotation. If it has it then it returns the value
// of the `path` parameter. It returns an empty string if there is no such annotation or parameter.
func ReferencePath(concept concepts.Annotated) string {
	annotation := concept.GetAnnotation("ref")
	if annotation == nil {
		return ""
	}
	name := annotation.FindParameter("path")
	if name == nil {
		return ""
	}
	return fmt.Sprintf("%s", name)
}
