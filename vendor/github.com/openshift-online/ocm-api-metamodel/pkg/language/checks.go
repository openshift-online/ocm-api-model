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

// This file contains the semantics checks.

package language

import (
	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
)

func (r *Reader) checkModel() {
	for _, service := range r.model.Services() {
		r.checkService(service)
	}
}

func (r *Reader) checkService(service *concepts.Service) {
	for _, version := range service.Versions() {
		r.checkVersion(version)
	}
}

func (r *Reader) checkVersion(version *concepts.Version) {
	// Check that there is a root resource:
	if version.Root() == nil {
		r.reporter.Errorf("Version '%s' doesn't have a root resource", version)
	}

	// Check the resources:
	for _, resource := range version.Resources() {
		r.checkResource(resource)
	}

	// Check that there are no loops in the tree of locators:
	r.checkLocatorLoops(version)
}

func (r *Reader) checkResource(resource *concepts.Resource) {
	for _, method := range resource.Methods() {
		r.checkMethod(method)
	}
}

func (r *Reader) checkMethod(method *concepts.Method) {
	// Run specific checks according tot he type of method:
	switch {
	case method.IsAdd():
		r.checkAdd(method)
	case method.IsDelete():
		r.checkDelete(method)
	case method.IsGet():
		r.checkGet(method)
	case method.IsList():
		r.checkList(method)
	case method.IsPost():
		r.checkPost(method)
	case method.IsSearch():
		r.checkSearch(method)
	case method.IsUpdate():
		r.checkUpdate(method)
	case method.IsAction():
		r.checkAction(method)
	default:
		r.reporter.Errorf(
			"Don't know how to check method '%s' because it doesn't belong to any "+
				"of the known categories",
			method,
		)
	}

	// Check the parameters:
	for _, parameter := range method.Parameters() {
		r.checkParameter(parameter)
	}
}

func (r *Reader) checkAdd(method *concepts.Method) {
	// Only scalar and struct parameters:
	for _, parameter := range method.Parameters() {
		if !parameter.Type().IsScalar() && !parameter.Type().IsStruct() {
			r.reporter.Errorf(
				"Type of parameter '%s' should be scalar or struct but it is %s",
				parameter, parameter.Type().Kind(),
			)
		}
	}

	// Exactly one struct parameter:
	var structs []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.Type().IsStruct() {
			structs = append(structs, parameter)
		}
	}
	count := len(structs)
	if count != 1 {
		r.reporter.Errorf(
			"Method '%s' should have exactly one struct parameter but it has %d",
			method, count,
		)
	}

	// Non scalar parameters should be input and output:
	for _, parameter := range structs {
		if !parameter.In() || !parameter.Out() {
			r.reporter.Errorf(
				"Direction of parameter '%s' should be 'in out'",
				parameter,
			)
		}
	}

	// Struct parameters should be named `body`:
	for _, parameter := range structs {
		if !nomenclator.Body.Equals(parameter.Name()) {
			r.reporter.Errorf(
				"Name of parameter '%s' should be '%s'",
				parameter, nomenclator.Body,
			)
		}
	}
}

func (r *Reader) checkDelete(method *concepts.Method) {
	// Only scalar parameters:
	for _, parameter := range method.Parameters() {
		if !parameter.Type().IsScalar() {
			r.reporter.Errorf(
				"Type of parameter '%s' should be scalar but it is %s",
				method, parameter.Type().Kind(),
			)
		}
	}

	// Only input parameters:
	for _, parameter := range method.Parameters() {
		if !parameter.In() || parameter.Out() {
			r.reporter.Errorf("Direction of parameter '%s' must be 'in'", parameter)
		}
	}
}

func (r *Reader) checkGet(method *concepts.Method) {
	// Only scalar and struct parameters:
	for _, parameter := range method.Parameters() {
		if !parameter.Type().IsScalar() && !parameter.Type().IsStruct() {
			r.reporter.Errorf(
				"Type of parameter '%s' must be scalar or struct but it is %s",
				parameter, parameter.Type().Kind(),
			)
		}
	}

	// Exactly one struct parameter:
	var structs []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.Type().IsStruct() {
			structs = append(structs, parameter)
		}
	}
	count := len(structs)
	if count != 1 {
		r.reporter.Errorf(
			"Method '%s' should have exactly one struct parameter but it has %d",
			method, count,
		)
	}

	// Scalar parameters should be input only:
	for _, parameter := range method.Parameters() {
		if parameter.Type().IsScalar() {
			if !parameter.In() || parameter.Out() {
				r.reporter.Errorf(
					"Direction of parameter '%s' should be 'in'",
					parameter,
				)
			}
		}
	}

	// Struct parameters should be output only:
	for _, parameter := range structs {
		if !parameter.Out() || parameter.In() {
			r.reporter.Errorf("Direction of parameter '%s' should be 'out'", parameter)
		}
	}

	// Struct parameters should be named `body`:
	for _, parameter := range structs {
		if !nomenclator.Body.Equals(parameter.Name()) {
			r.reporter.Errorf(
				"Name of parameter '%s' should be '%s'",
				parameter, nomenclator.Body,
			)
		}
	}
}

func (r *Reader) checkList(method *concepts.Method) {
	// Only scalar and list parameters:
	for _, parameter := range method.Parameters() {
		if !parameter.Type().IsScalar() && !parameter.Type().IsList() {
			r.reporter.Errorf(
				"Type of parameter '%s' should be scalar or list but it is '%s'",
				parameter, parameter.Type().Kind(),
			)
		}
	}

	// Exactly one list parameter:
	var lists []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.Type().IsList() {
			lists = append(lists, parameter)
		}
	}
	count := len(lists)
	if count != 1 {
		r.reporter.Errorf(
			"Method '%s' should have exactly one list parameter but it has %d",
			method, count,
		)
	}

	r.checkRequestParameter(method, nomenclator.Page, 1)
	r.checkRequestParameter(method, nomenclator.Size, 100)
	r.checkResponseParameter(method, nomenclator.Total)

	// Check the `items` parameter:
	items := method.GetParameter(nomenclator.Items)
	if items == nil {
		r.reporter.Errorf(
			"Method '%s' doesn't have a '%s' parameter",
			method, nomenclator.Items,
		)
	}

	// List parameters should be output only:
	for _, parameter := range lists {
		if parameter.In() || !parameter.Out() {
			r.reporter.Errorf(
				"Direction of parameter '%s' should be 'out'",
				parameter,
			)
		}
	}

	// List parameters should be named `items`:
	for _, parameter := range lists {
		if !nomenclator.Items.Equals(parameter.Name()) {
			r.reporter.Errorf(
				"Name of parameter '%s' should be '%s'",
				parameter, nomenclator.Items,
			)
		}
	}
}

func (r *Reader) checkPost(method *concepts.Method) {
	// Get the input and output parameters:
	var in []*concepts.Parameter
	var out []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.In() {
			in = append(in, parameter)
		}
		if parameter.Out() {
			out = append(out, parameter)
		}
	}

	// At most one input parameter that should be an struct:
	if len(in) > 1 {
		r.reporter.Errorf(
			"Method '%s' should have at most one input parameter but it has %d",
			method, len(in),
		)
	}
	for _, parameter := range in {
		if !parameter.Type().IsStruct() {
			r.reporter.Errorf(
				"Parameter '%s' should be a struct but it is %s",
				parameter, parameter.Type().Kind(),
			)
		}
	}

	// At most one output parameter that should be an struct:
	if len(out) != 1 {
		r.reporter.Errorf(
			"Method '%s' should have at most one output parameter but it has %d",
			method, len(out),
		)
	}
	for _, parameter := range out {
		if !parameter.Type().IsStruct() {
			r.reporter.Errorf(
				"Parameter '%s' should be a struct but it is %s",
				parameter, parameter.Type().Kind(),
			)
		}
	}
}

func (r *Reader) checkSearch(method *concepts.Method) {
	// Get the input and output parameters:
	var in []*concepts.Parameter
	var out []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.In() {
			in = append(in, parameter)
		}
		if parameter.Out() {
			out = append(out, parameter)
		}
	}

	// Only scalar or list parameters:
	for _, parameter := range out {
		if !parameter.Type().IsScalar() && !parameter.Type().IsList() {
			r.reporter.Errorf(
				"Parameter '%s' should be scalar or list but it is %s",
				parameter, parameter.Type().Kind(),
			)
		}
	}

	// Only scalar or struct parameters:
	for _, parameter := range in {
		if !parameter.Type().IsScalar() && !parameter.Type().IsStruct() {
			r.reporter.Errorf(
				"Parameter '%s' should be scalar or struct but it is '%s'",
				parameter, parameter.Type().Kind(),
			)
		}
	}

	// Exactly one list parameter:
	var lists []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.Type().IsList() {
			lists = append(lists, parameter)
		}
	}
	count := len(lists)
	if count != 1 {
		r.reporter.Errorf(
			"Method '%s' should have exactly one list parameter but it has %d",
			method, count,
		)
	}

	r.checkRequestParameter(method, nomenclator.Page, 1)
	r.checkRequestParameter(method, nomenclator.Size, 100)
	r.checkResponseParameter(method, nomenclator.Total)

	// Check the `items` parameter:
	items := method.GetParameter(nomenclator.Items)
	if items == nil {
		r.reporter.Errorf(
			"Method '%s' doesn't have a '%s' parameter",
			method, nomenclator.Items,
		)
	}

	// List parameters should be output only:
	for _, parameter := range lists {
		if parameter.In() || !parameter.Out() {
			r.reporter.Errorf(
				"Direction of parameter '%s' should be 'out'",
				parameter,
			)
		}
	}

	// List parameters should be named `items`:
	for _, parameter := range lists {
		if !nomenclator.Items.Equals(parameter.Name()) {
			r.reporter.Errorf(
				"Name of parameter '%s' should be '%s'",
				parameter, nomenclator.Items,
			)
		}
	}

	// Exactly one body parameter:
	var bodies []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.Type().IsStruct() && parameter.In() {
			bodies = append(bodies, parameter)
		}
	}
	count = len(bodies)
	if count != 1 {
		r.reporter.Errorf(
			"Method '%s' should have exactly one body parameter but it has %d",
			method, count,
		)
	}

	for _, parameter := range bodies {
		// Struct parameters should be input only:
		if parameter.Out() || !parameter.In() {
			r.reporter.Errorf(
				"Direction of parameter '%s' should be 'in'",
				parameter,
			)
		}

		// Struct parameter should be named `body`:
		if !nomenclator.Body.Equals(parameter.Name()) {
			r.reporter.Errorf(
				"Name of parameter '%s' should be be '%s'",
				parameter, nomenclator.Body,
			)
		}
	}
}

func (r *Reader) checkUpdate(method *concepts.Method) {
	// Only scalar, struct and list parameters:
	for _, parameter := range method.Parameters() {
		if !parameter.Type().IsScalar() && !parameter.Type().IsStruct() && !parameter.Type().IsList() {
			r.reporter.Errorf(
				"Type of parameter '%s' should be scalar, struct or list but it is '%s'",
				parameter, parameter.Type().Kind(),
			)
		}
	}

	// Exactly one struct parameter:
	var parameters []*concepts.Parameter
	for _, parameter := range method.Parameters() {
		if parameter.Type().IsStruct() || parameter.Type().IsList() {
			parameters = append(parameters, parameter)
		}
	}
	count := len(parameters)
	if count != 1 {
		r.reporter.Errorf(
			"Method '%s' should have exactly one struct or list parameter but it has %d",
			method, count,
		)
	}

	// Scalar parameters should be input only:
	for _, parameter := range method.Parameters() {
		if parameter.Type().IsScalar() {
			if !parameter.In() || parameter.Out() {
				r.reporter.Errorf(
					"Direction of parameter '%s' should be 'in'",
					parameter,
				)
			}
		}
	}

	// Struct and list parameters should be input and output:
	for _, parameter := range parameters {
		if !parameter.In() || !parameter.Out() {
			r.reporter.Errorf(
				"Direction of parameter '%s' should be 'in out'",
				parameter,
			)
		}
	}

	// Struct and list parameters should be named `body`:
	for _, parameter := range parameters {
		if !nomenclator.Body.Equals(parameter.Name()) {
			r.reporter.Errorf(
				"Name of parameter '%s' should be be '%s'",
				parameter, nomenclator.Body,
			)
		}
	}
}

func (r *Reader) checkAction(method *concepts.Method) {
	// Empty on purpose.
}

func (r *Reader) checkParameter(parameter *concepts.Parameter) {
	// Get the version:
	version := parameter.Owner().Owner().Owner()

	// Check that the default value is of a type compatible with the type of the parameter:
	value := parameter.Default()
	if value != nil {
		var typ *concepts.Type
		switch value.(type) {
		case bool:
			typ = version.Boolean()
		case int:
			typ = version.IntegerType()
		case string:
			typ = version.StringType()
		default:
			r.reporter.Errorf(
				"Don't know how to check if default value '%v' is compatible "+
					"with type of parameter '%s'",
				value, parameter,
			)
		}
		if typ != nil && typ != parameter.Type() {
			r.reporter.Errorf(
				"Type of default value of parameter '%s' should be '%s'",
				parameter, parameter.Type(),
			)
		}
	}
}

func (r *Reader) checkRequestParameter(method *concepts.Method, name *names.Name, dflt int) {
	// Get the reference to the resource and to the version:
	resource := method.Owner()
	version := resource.Owner()

	// Check the parameter:
	param := method.GetParameter(name)
	if param == nil {
		r.reporter.Warnf(
			"Method '%s' doesn't have a '%s' parameter",
			method, name,
		)
	} else {
		if param.Type() != version.IntegerType() {
			r.reporter.Errorf(
				"Type of parameter '%s' should be integer but it is '%s'",
				param, param.Type(),
			)
		}
		if !param.In() || !param.Out() {
			r.reporter.Errorf(
				"Direction of parameter '%s' should be 'in out'",
				param,
			)
		}
		if param.Default() != dflt {
			r.reporter.Errorf(
				"Default value of parameter `%s` should be %d",
				param, dflt,
			)
		}
	}
}

func (r *Reader) checkResponseParameter(method *concepts.Method, name *names.Name) {
	// Get the reference to the resource and to the version:
	resource := method.Owner()
	version := resource.Owner()

	// Check the parameter:
	param := method.GetParameter(name)
	if param == nil {
		r.reporter.Warnf(
			"Method '%s' doesn't have a '%s' parameter",
			method, name,
		)
	} else {
		if param.Type() != version.IntegerType() {
			r.reporter.Errorf(
				"Type of parameter '%s' should be integer but it is '%s'",
				param, param.Type(),
			)
		}
		if param.In() || !param.Out() {
			r.reporter.Errorf(
				"Direction of parameter '%s' should be 'out'",
				param,
			)
		}
	}
}

func (r *Reader) checkLocatorLoops(version *concepts.Version) {
	for _, locator := range version.Root().Locators() {
		path := []*concepts.Locator{locator}
		r.checkLocatorPathLoops(path)
	}
}

func (r *Reader) checkLocatorPathLoops(path []*concepts.Locator) {
	// Try to find a previous locator that is defined in the same resource that is the target of
	// the last locator. If there is such locator then there is a loop and we should not
	// continue checking this path.
	last := len(path) - 1
	first := -1
	for i := last; i >= 0; i-- {
		if path[i].Owner() == path[last].Target() {
			first = i
			break
		}
	}
	if first >= 0 {
		r.reportLocatorPathLoop(path[first:])
		return
	}

	// If we are here then there is no loop in this path, so we can continue recusively:
	for _, next := range path[last].Target().Locators() {
		r.checkLocatorPathLoops(append(path, next))
	}
}

func (r *Reader) reportLocatorPathLoop(path []*concepts.Locator) {
	if len(path) == 1 {
		locator := path[0]
		r.reporter.Errorf(
			"Locator '%s' introduces a loop, it is defined in resource '%s' and "+
				"targets that same resource",
			locator.Name(), locator.Owner(),
		)
	} else {
		first := path[0]
		last := path[len(path)-1]
		r.reporter.Errorf(
			"Locator '%s' defined in resource '%s' introduces a loop, see the "+
				"details below",
			last.Name(), last.Owner(),
		)
		r.reporter.Infof(
			"First locator of the loop is '%s' defined in resource '%s' and "+
				"targeting resource '%s'",
			first.Name(), first.Owner(), first.Target(),
		)
		for i := 1; i < len(path)-1; i++ {
			next := path[i]
			r.reporter.Infof(
				"Next locator of the loop is '%s' defined in resource '%s' and "+
					"targeting resource '%s'",
				next.Name(), next.Owner(), next.Target(),
			)
		}
		r.reporter.Infof(
			"Last locator of the loop is '%s' defined in resource '%s' and "+
				"targeting resource '%s'",
			last.Name(), last.Owner(), last.Target(),
		)
	}
}
