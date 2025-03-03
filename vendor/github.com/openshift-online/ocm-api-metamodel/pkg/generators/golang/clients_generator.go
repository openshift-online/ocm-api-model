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

package golang

import (
	"fmt"

	"github.com/openshift-online/ocm-api-metamodel/pkg/annotations"
	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/http"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// ClientsGeneratorBuilder is an object used to configure and build a client generator. Don't create
// instances directly, use the NewClientsGenerator function instead.
type ClientsGeneratorBuilder struct {
	reporter *reporter.Reporter
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	types    *TypesCalculator
	binding  *http.BindingCalculator
}

// ClientsGenerator generates client code. Don't create instances directly, use the builder instead.
type ClientsGenerator struct {
	reporter *reporter.Reporter
	errors   int
	model    *concepts.Model
	output   string
	packages *PackagesCalculator
	names    *NamesCalculator
	types    *TypesCalculator
	binding  *http.BindingCalculator
	buffer   *Buffer
}

// NewClientsGenerator creates a new builder for client generators.
func NewClientsGenerator() *ClientsGeneratorBuilder {
	return new(ClientsGeneratorBuilder)
}

// Reporter sets the object that will be used to report information about the generation process,
// including errors.
func (b *ClientsGeneratorBuilder) Reporter(value *reporter.Reporter) *ClientsGeneratorBuilder {
	b.reporter = value
	return b
}

// Model sets the model that will be used by the client generator.
func (b *ClientsGeneratorBuilder) Model(value *concepts.Model) *ClientsGeneratorBuilder {
	b.model = value
	return b
}

// Output sets the output directory.
func (b *ClientsGeneratorBuilder) Output(value string) *ClientsGeneratorBuilder {
	b.output = value
	return b
}

// Packages sets the object that will be used to calculate package names.
func (b *ClientsGeneratorBuilder) Packages(
	value *PackagesCalculator) *ClientsGeneratorBuilder {
	b.packages = value
	return b
}

// Names sets the object that will be used to calculate names.
func (b *ClientsGeneratorBuilder) Names(value *NamesCalculator) *ClientsGeneratorBuilder {
	b.names = value
	return b
}

// Types sets the object that will be used to calculate types.
func (b *ClientsGeneratorBuilder) Types(value *TypesCalculator) *ClientsGeneratorBuilder {
	b.types = value
	return b
}

// Binding sets the object that will by used to do HTTP binding calculations.
func (b *ClientsGeneratorBuilder) Binding(value *http.BindingCalculator) *ClientsGeneratorBuilder {
	b.binding = value
	return b
}

// Build checks the configuration stored in the builder and, if it is correct, creates a new client
// generator using it.
func (b *ClientsGeneratorBuilder) Build() (generator *ClientsGenerator, err error) {
	// Check that the mandatory parameters have been provided:
	if b.reporter == nil {
		err = fmt.Errorf("reporter is mandatory")
		return
	}
	if b.model == nil {
		err = fmt.Errorf("model is mandatory")
		return
	}
	if b.output == "" {
		err = fmt.Errorf("path is mandatory")
		return
	}
	if b.packages == nil {
		err = fmt.Errorf("packages calculator is mandatory")
		return
	}
	if b.names == nil {
		err = fmt.Errorf("names calculator is mandatory")
		return
	}
	if b.types == nil {
		err = fmt.Errorf("types calculator is mandatory")
		return
	}
	if b.binding == nil {
		err = fmt.Errorf("binding calculator is mandatory")
		return
	}

	// Create the generator:
	generator = &ClientsGenerator{
		reporter: b.reporter,
		model:    b.model,
		output:   b.output,
		packages: b.packages,
		names:    b.names,
		types:    b.types,
		binding:  b.binding,
	}

	return
}

// Run executes the code generator.
func (g *ClientsGenerator) Run() error {
	var err error

	// Generate the client for each service:
	for _, service := range g.model.Services() {
		g.reporter.Infof("Generating client for service '%s'", service.Name())
		err = g.generateServiceClient(service)
		if err != nil {
			return err
		}
	}

	// Check if there were errors:
	if g.errors > 0 {
		if g.errors > 1 {
			err = fmt.Errorf("there were %d errors", g.errors)
		} else {
			err = fmt.Errorf("there was 1 error")
		}
		return err
	}

	return nil
}

func (g *ClientsGenerator) generateServiceClient(service *concepts.Service) error {
	var err error

	// Calculate the package and file name:
	pkgName := g.packages.ServicePackage(service)
	fileName := g.names.File(nomenclator.Client)

	// Create the buffer for the service:
	g.buffer, err = NewBuffer().
		Reporter(g.reporter).
		Output(g.output).
		Packages(g.packages).
		Package(pkgName).
		File(fileName).
		Function("clientName", g.clientName).
		Function("versionSegment", g.binding.VersionSegment).
		Function("versionName", g.versionName).
		Function("versionSelector", g.packages.VersionSelector).
		Build()
	if err != nil {
		return err
	}

	// Generate the source for the service:
	err = g.generateServiceClientSource(service)
	if err != nil {
		return err
	}
	err = g.buffer.Write()
	if err != nil {
		return err
	}

	// Generate the clients for the versions:
	for _, version := range service.Versions() {
		err = g.generateVersionClient(version)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *ClientsGenerator) generateServiceClientSource(service *concepts.Service) error {
	g.buffer.Import("net/http", "")
	g.buffer.Import("path", "")
	for _, version := range service.Versions() {
		g.buffer.Import(g.packages.VersionImport(version), "")
	}
	g.buffer.Emit(`
		// Client is the client for service '{{ .Service.Name }}'.
		type Client struct {
			transport http.RoundTripper
			path string
		}

		// NewClient creates a new client for the service '{{ .Service.Name }}' using the
		// given transport to send the requests and receive the responses.
		func NewClient(transport http.RoundTripper, path string) *Client {
			client := new(Client)
			client.transport = transport
			client.path = path
			return client
		}

		{{ range .Service.Versions }}
			{{ $versionName := versionName . }}
			{{ $versionSelector := versionSelector . }}
			{{ $versionSegment := versionSegment . }}
			{{ $rootName := clientName .Root }}

			// {{ $versionName }} returns a reference to a client for version '{{ .Name }}'.
			func (c *Client) {{ $versionName }}() *{{ $versionSelector }}.{{ $rootName }} {
				return {{ $versionSelector }}.New{{ $rootName }}(
					c.transport,
					path.Join(c.path, "{{ $versionSegment }}"),
				)
			}
		{{ end }}
		`,
		"Service", service,
	)

	return nil
}

func (g *ClientsGenerator) generateVersionClient(version *concepts.Version) error {
	// Generate the metadata client:
	err := g.generateVersionMetadataClient(version)
	if err != nil {
		return err
	}

	// Generate the resource clients:
	for _, resource := range version.Resources() {
		err := g.generateResourceClient(resource)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *ClientsGenerator) generateVersionMetadataClient(version *concepts.Version) error {
	var err error

	// Calculate the package and file name:
	pkgName := g.packages.VersionPackage(version)
	fileName := g.metadataFile()

	// Create the buffer for the generated code:
	g.buffer, err = NewBuffer().
		Reporter(g.reporter).
		Output(g.output).
		Packages(g.packages).
		Package(pkgName).
		File(fileName).
		Build()
	if err != nil {
		return err
	}

	// Generate the code:
	g.generateVersionMetadataClientSource(version)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *ClientsGenerator) generateVersionMetadataClientSource(version *concepts.Version) {
	g.buffer.Import("bufio", "")
	g.buffer.Import("context", "")
	g.buffer.Import("io", "")
	g.buffer.Import("net/http", "")
	g.buffer.Import("net/url", "")
	g.buffer.Import(g.packages.ErrorsImport(), "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		// MetadataRequest is the request to retrieve the metadata.
		type MetadataRequest struct {
			transport http.RoundTripper
			path      string
			query     url.Values
			header    http.Header
		}

		// MetadataResponse is the response for the metadata request.
		type MetadataResponse struct {
			status int
			header http.Header
			err    *errors.Error
			body   *Metadata
		}

		// Parameter adds a query parameter.
		func (r *MetadataRequest) Parameter(name string, value interface{}) *MetadataRequest {
			helpers.AddValue(&r.query, name, value)
			return r
		}

		// Header adds a request header.
		func (r *MetadataRequest) Header(name string, value interface{}) *MetadataRequest {
			helpers.AddHeader(&r.header, name, value)
			return r
		}

		// Send sends the metadata request, waits for the response, and returns it.
		//
		// This is a potentially lengthy operation, as it requires network communication.
		// Consider using a context and the SendContext method.
		func (r *MetadataRequest) Send() (result *MetadataResponse, err error) {
			return r.SendContext(context.Background())
		}

		// SendContext sends the metadata request, waits for the response, and returns it.
		func (r *MetadataRequest) SendContext(ctx context.Context) (result *MetadataResponse, err error) {
			query := helpers.CopyQuery(r.query)
			header := helpers.CopyHeader(r.header)
			uri := &url.URL{
				Path: r.path,
				RawQuery: query.Encode(),
			}
			request := &http.Request{
				Method: http.MethodGet,
				URL:    uri,
				Header: header,
			}
			if ctx != nil {
				request = request.WithContext(ctx)
			}
			response, err := r.transport.RoundTrip(request)
			if err != nil {
				return
			}
			defer response.Body.Close()
			result = &MetadataResponse{
				status: response.StatusCode,
				header: response.Header,
			}
			reader := bufio.NewReader(response.Body)
			_, err = reader.Peek(1)
			if err == io.EOF {
				return
			}
			if result.status >= 400 {
				result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
				if err != nil {
					return
				}
				err = result.err
				return
			}
			result.body, err = UnmarshalMetadata(reader)
			if err != nil {
				return
			}
			return
		}

		// Status returns the response status code.
		func (r *MetadataResponse) Status() int {
			if r == nil {
				return 0
			}
			return r.status
		}

		// Header returns header of the response.
		func (r *MetadataResponse) Header() http.Header {
			if r == nil {
				return nil
			}
			return r.header
		}

		// Error returns the response error.
		func (r *MetadataResponse) Error() *errors.Error {
			if r == nil {
				return nil
			}
			return r.err
		}

		// Body returns the response body.
		func (r *MetadataResponse) Body() *Metadata {
			return r.body
		}
		`,
	)
}

func (g *ClientsGenerator) generateResourceClient(resource *concepts.Resource) error {
	var err error

	// Calculate the package and file name:
	pkgName := g.packages.VersionPackage(resource.Owner())
	fileName := g.resourceFile(resource)

	// Create the buffer for the generated code:
	g.buffer, err = NewBuffer().
		Reporter(g.reporter).
		Output(g.output).
		Packages(g.packages).
		Package(pkgName).
		File(fileName).
		Function("clientName", g.clientName).
		Function("enumName", g.enumName).
		Function("fieldName", g.fieldName).
		Function("fieldType", g.fieldType).
		Function("getterName", g.getterName).
		Function("getterType", g.getterType).
		Function("httpMethod", g.binding.Method).
		Function("locatorName", g.locatorName).
		Function("locatorSegment", g.binding.LocatorSegment).
		Function("methodName", g.methodName).
		Function("methodSegment", g.binding.MethodSegment).
		Function("parameterName", g.binding.QueryParameterName).
		Function("pollRequestName", g.pollRequestName).
		Function("pollResponseName", g.pollResponseName).
		Function("readResponseFunc", g.readResponseFunc).
		Function("requestBodyParameters", g.binding.RequestBodyParameters).
		Function("requestName", g.requestName).
		Function("requestParameters", g.binding.RequestParameters).
		Function("requestQueryParameters", g.binding.RequestQueryParameters).
		Function("responseName", g.responseName).
		Function("responseParameters", g.binding.ResponseParameters).
		Function("setterName", g.setterName).
		Function("setterType", g.setterType).
		Function("structName", g.types.StructName).
		Function("valueType", g.types.ValueReference).
		Function("writeRequestFunc", g.writeRequestFunc).
		Function("zeroValue", g.types.ZeroValue).
		Build()
	if err != nil {
		return err
	}

	// Generate the code:
	g.generateResourceClientSource(resource)

	// Write the generated code:
	return g.buffer.Write()
}

func (g *ClientsGenerator) generateResourceClientSource(resource *concepts.Resource) {
	// We need to know if this is the root resource in order to add the metadata methods:
	root := resource == resource.Owner().Root()

	// Generate the source:
	g.buffer.Import("net/http", "")
	g.buffer.Import("path", "")
	g.buffer.Emit(`
		{{ $clientName := clientName .Resource }}

		// {{ $clientName }} is the client of the '{{ .Resource.Name }}' resource.
		//
		{{ lineComment .Resource.Doc }}
		type {{ $clientName }} struct {
			transport http.RoundTripper
			path string
		}

		// New{{ $clientName }} creates a new client for the '{{ .Resource.Name }}'
		// resource using the given transport to send the requests and receive the
		// responses.
		func New{{ $clientName }}(transport http.RoundTripper, path string) *{{ $clientName }} {
			return &{{ $clientName }}{
				transport: transport,
				path:      path,
			}
		}

		{{ range .Resource.Methods }}
			{{ $methodName := methodName . }}
			{{ $methodSegment := methodSegment . }}
			{{ $requestName := requestName . }}

			// {{ $methodName }} creates a request for the '{{ .Name }}' method.
			//
			{{ lineComment .Doc }}
			func (c *{{ $clientName }}) {{ $methodName }}() *{{ $requestName }} {
				return &{{ $requestName }}{
					transport: c.transport,
					{{ if $methodSegment }}
						path: path.Join(c.path, "{{ $methodSegment }}"),
					{{ else }}
						path: c.path,
					{{ end }}
				}
			}
		{{ end }}

		{{ if .Root }}
			// Creates a new request for the method that retrieves the metadata.
			func (c *{{ $clientName }}) Get() *MetadataRequest {
				return &MetadataRequest{
					transport: c.transport,
					path:      c.path,
				}
			}
		{{ end }}

		{{ range .Resource.Locators }}
			{{ $locatorName := locatorName . }}
			{{ $locatorSegment := locatorSegment . }}
			{{ $targetName := clientName .Target }}

			{{ if .Variable }}
				// {{ $locatorName }} returns the target '{{ .Target.Name }}' resource for the given identifier.
				//
				{{ lineComment .Doc }}
				func (c *{{ $clientName }}) {{ $locatorName }}(id string) *{{ $targetName }} {
					return New{{ $targetName }}(
						c.transport,
						path.Join(c.path, id),
					)
				}
			{{ else }}
				// {{ $locatorName }} returns the target '{{ .Target.Name }}' resource.
				//
				{{ lineComment .Doc }}
				func (c *{{ $clientName }}) {{ $locatorName }}() *{{ $targetName }} {
					return New{{ $targetName }}(
						c.transport,
						path.Join(c.path, "{{ $locatorSegment }}"),
					)
				}
			{{ end }}
		{{ end }}
		`,
		"Resource", resource,
		"Root", root,
	)

	// If the resource has a `Get` method then generate the `Poll` method:
	method := resource.FindMethod(nomenclator.Get)
	if method != nil {
		g.generatePollMethodSource(resource, method)
	}

	// Generate the request and response types:
	for _, method := range resource.Methods() {
		g.generateRequestSource(method)
		g.generateResponseSource(method)
	}
}

func (g *ClientsGenerator) generatePollMethodSource(resource *concepts.Resource, method *concepts.Method) {
	// Find the response body parameter:
	body := g.binding.ResponseBodyParameters(method)[0]

	// Generate the code:
	g.buffer.Import("net/http", "")
	g.buffer.Import("time", "")
	g.buffer.Import(g.packages.ErrorsImport(), "")
	g.buffer.Emit(`
		{{ $clientName := clientName .Resource }}
		{{ $requestName := pollRequestName .Resource }}
		{{ $responseName := pollResponseName .Resource }}
		{{ $methodRequestName := requestName .Method }}
		{{ $methodResponseName := responseName .Method }}
		{{ $methodRequestParameters := requestParameters .Method }}
		{{ $methodResponseParameters := responseParameters .Method }}

		// {{ $requestName }} is the request for the Poll method.
		type {{ $requestName }} struct {
			request    *{{ $methodRequestName }}
			interval   time.Duration
			statuses   []int
			predicates []func (interface{}) bool
		}

		// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
		func (r *{{ $requestName }}) Parameter(name string, value interface{}) *{{ $requestName }} {
			r.request.Parameter(name, value)
			return r
		}

		// Header adds a request header to all the requests that will be used to retrieve the object.
		func (r *{{ $requestName }}) Header(name string, value interface{}) *{{ $requestName }} {
			r.request.Header(name, value)
			return r
		}

		{{ range $methodRequestParameters }}
			{{ $setterName := setterName . }}
			{{ $setterType := setterType . }}

			// {{ $setterName }} sets the value of the '{{ .Name }}' parameter for all the requests that
			// will be used to retrieve the object.
			//
			{{ lineComment .Doc }}
			func (r *{{ $requestName }}) {{ $setterName }}(value {{ $setterType }}) *{{ $requestName }} {
				r.request.{{ $setterName }}(value)
				return r
			}
		{{ end }}

		// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
		func (r *{{ $requestName }}) Interval(value time.Duration) *{{ $requestName }} {
			r.interval = value
			return r
		}

		// Status set the expected status of the response. Multiple values can be set calling this method
		// multiple times. The response will be considered successful if the status is any of those values.
		func (r *{{ $requestName }}) Status(value int) *{{ $requestName }} {
			r.statuses = append(r.statuses, value)
			return r
		}

		// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
		// predicates can be set calling this method multiple times. The response will be considered successful
		// if all the predicates are satisfied.
		func (r *{{ $requestName }}) Predicate(value func (*{{ $methodResponseName }}) bool) *{{ $requestName }} {
			r.predicates = append(r.predicates, func(response interface{}) bool {
				return value(response.(*{{ $methodResponseName }}))
			})
			return r
		}

		// StartContext starts the polling loop. Responses will be considered successful if the status is one of
		// the values specified with the Status method and if all the predicates specified with the Predicate
		// method return nil.
		//
		// The context must have a timeout or deadline, otherwise this method will immediately return an error.
		func (r *{{ $requestName }}) StartContext(ctx context.Context) (response *{{ $responseName }}, err error) {
			result, err := helpers.PollContext(ctx, r.interval, r.statuses, r.predicates, r.task)
			if result != nil {
				response = &{{ $responseName }}{
					response: result.(*{{ $methodResponseName }}),
				}
			}
			return
		}

		// task adapts the types of the request/response types so that they can be used with the generic
		// polling function from the helpers package.
		func (r *{{ $requestName }}) task(ctx context.Context) (status int, result interface{}, err error) {
			response, err := r.request.SendContext(ctx)
			if response != nil {
				status = response.Status()
				result = response
			}
			return
		}

		// {{ $responseName }} is the response for the Poll method.
		type {{ $responseName }} struct {
			response *{{ $methodResponseName }}
		}

		// Status returns the response status code.
		func (r *{{ $responseName }}) Status() int {
			if r == nil {
				return 0
			}
			return r.response.Status()
		}

		// Header returns header of the response.
		func (r *{{ $responseName }}) Header() http.Header {
			if r == nil {
				return nil
			}
			return r.response.Header()
		}

		// Error returns the response error.
		func (r *{{ $responseName }}) Error() *errors.Error {
			if r == nil {
				return nil
			}
			return r.response.Error()
		}

		{{ range $methodResponseParameters }}
			{{ $parameterType := .Type.Name.String }}
			{{ $fieldName := fieldName . }}
			{{ $getterName := getterName . }}
			{{ $getterType := getterType . }}

			// {{ $getterName }} returns the value of the '{{ .Name }}' parameter.
			//
			{{ lineComment .Doc }}
			func (r *{{ $responseName }}) {{ $getterName }}() {{ $getterType }} {
				return r.response.{{ $getterName }}()
			}

			// Get{{ $getterName }} returns the value of the '{{ .Name }}' parameter and
			// a flag indicating if the parameter has a value.
			//
			{{ lineComment .Doc }}
			func (r *{{ $responseName }}) Get{{ $getterName }}() (value {{ $getterType }}, ok bool) {
				return r.response.Get{{ $getterName }}()
			}
		{{ end }}

		// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
		// of states and satisfies a set of predicates.
		func (c *{{ $clientName }}) Poll() *{{ $requestName }} {
			return &{{ $requestName }}{
				request: c.{{ methodName .Method }}(),
			}
		}
		`,
		"Resource", resource,
		"Method", method,
		"Body", body,
	)
}

func (g *ClientsGenerator) generateRequestSource(method *concepts.Method) {
	// Classify the parameters:
	all := g.binding.RequestBodyParameters(method)
	var main *concepts.Parameter
	var others []*concepts.Parameter
	for _, parameter := range all {
		if parameter.IsItems() || parameter.IsBody() {
			main = parameter
		} else {
			others = append(others, parameter)
		}
	}

	g.buffer.Import("bufio", "")
	g.buffer.Import("bytes", "")
	g.buffer.Import("context", "")
	g.buffer.Import("io", "")
	g.buffer.Import("os", "")
	g.buffer.Import("net/http", "")
	g.buffer.Import("net/url", "")
	g.buffer.Import(g.packages.ErrorsImport(), "")
	g.buffer.Import(g.packages.HelpersImport(), "")
	g.buffer.Emit(`
		{{ $requestName := requestName .Method }}
		{{ $requestParameters := requestParameters .Method }}
		{{ $requestQueryParameters := requestQueryParameters .Method }}
		{{ $requestBodyParameters := requestBodyParameters .Method }}
		{{ $requestBodyLen := len $requestBodyParameters }}
		{{ $responseName := responseName .Method }}
		{{ $responseParameters := responseParameters .Method }}
		{{ $isAction := .Method.IsAction }}

		// {{ $requestName }} is the request for the '{{ .Method.Name }}' method.
		type {{ $requestName }} struct {
			transport http.RoundTripper
			path      string
			query     url.Values
			header    http.Header
			{{ range $requestParameters }}
				{{ fieldName . }} {{ fieldType . }}
			{{ end }}
		}

		// Parameter adds a query parameter.
		func (r *{{ $requestName }}) Parameter(name string, value interface{}) *{{ $requestName }} {
			helpers.AddValue(&r.query, name, value)
			return r
		}

		// Header adds a request header.
		func (r *{{ $requestName }}) Header(name string, value interface{}) *{{ $requestName }} {
			helpers.AddHeader(&r.header, name, value)
			return r
		}

		// Impersonate wraps requests on behalf of another user.
		// Note: Services that do not support this feature may silently ignore this call. 
		func (r *{{ $requestName }}) Impersonate(user string) *{{ $requestName }} {
			helpers.AddImpersonationHeader(&r.header, user)
			return r
		}

		{{ range $requestParameters }}
			{{ $fieldName := fieldName . }}
			{{ $setterName := setterName . }}
			{{ $setterType := setterType . }}

			// {{ $setterName }} sets the value of the '{{ .Name }}' parameter.
			//
			{{ lineComment .Doc }}
			func (r *{{ $requestName }}) {{ $setterName }}(value {{ $setterType }}) *{{ $requestName }} {
				{{ if or .Type.IsStruct .Type.IsList }}
					r.{{ $fieldName }} = value
				{{ else }}
					r.{{ $fieldName }} = &value
				{{ end }}
				return r
			}
		{{ end }}

		// Send sends this request, waits for the response, and returns it.
		//
		// This is a potentially lengthy operation, as it requires network communication.
		// Consider using a context and the SendContext method.
		func (r *{{ $requestName }}) Send() (result *{{ $responseName }}, err error) {
			return r.SendContext(context.Background())
		}

		// SendContext sends this request, waits for the response, and returns it.
		func (r *{{ $requestName }}) SendContext(ctx context.Context) (result *{{ $responseName }}, err error) {
			query := helpers.CopyQuery(r.query)
			{{ range $requestQueryParameters }}
				{{ $fieldName := fieldName . }}
				{{ $parameterName := parameterName . }}
				if r.{{ $fieldName }} != nil {
					helpers.AddValue(&query, "{{ $parameterName }}", *r.{{ $fieldName }})
				}
			{{ end }}
			header := helpers.CopyHeader(r.header)
			{{ if $requestBodyParameters }}
				buffer := &bytes.Buffer{}
				err = {{ writeRequestFunc .Method }}(r, buffer)
				if err != nil {
					return
				}
			{{ end }}
			uri := &url.URL{
				Path: r.path,
				RawQuery: query.Encode(),
			}
			request := &http.Request{
				Method: "{{ httpMethod .Method }}",
				URL:    uri,
				Header: header,
				{{ if $requestBodyParameters }}
					Body: io.NopCloser(buffer),
				{{ end }}
			}
			if ctx != nil {
				request = request.WithContext(ctx)
			}
			response, err := r.transport.RoundTrip(request)
			if err != nil {
				return
			}
			defer response.Body.Close()
			result = &{{ $responseName }}{}
			result.status = response.StatusCode
			result.header = response.Header
			reader := bufio.NewReader(response.Body)
			_, err = reader.Peek(1)
			if err == io.EOF {
				err = nil
				return
			}
			if result.status >= 400 {
				result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
				if err != nil {
					return
				}
				err = result.err
				return
			}
			{{ if $responseParameters }}
				err = {{ readResponseFunc .Method }}(result, reader)
				if err != nil {
					return
				}
			{{ end }}
			return
		}
		`,
		"Method", method,
		"Main", main,
		"Others", others,
	)
}

func (g *ClientsGenerator) generateResponseSource(method *concepts.Method) {
	// Classify the parameters:
	all := g.binding.ResponseBodyParameters(method)
	var main *concepts.Parameter
	var others []*concepts.Parameter
	for _, parameter := range all {
		if parameter.IsItems() || parameter.IsBody() {
			main = parameter
		} else {
			others = append(others, parameter)
		}
	}

	// Generate the code:
	g.buffer.Import("io", "")
	g.buffer.Import("net/http", "")
	g.buffer.Import(g.packages.ErrorsImport(), "")
	g.buffer.Emit(`
		{{ $responseName := responseName .Method }}
		{{ $responseParameters := responseParameters .Method }}
		{{ $responseBodyLen := len $responseParameters }}
		{{ $isAction := .Method.IsAction }}

		// {{ $responseName }} is the response for the '{{ .Method.Name }}' method.
		type  {{ $responseName }} struct {
			status int
			header http.Header
			err    *errors.Error
			{{ range $responseParameters }}
				{{ if and .Type.IsList .Type.Element.IsScalar }}
					{{ fieldName . }} []{{ valueType .Type.Element }}
				{{ else }}
					{{ fieldName . }} {{ fieldType . }}
				{{ end }}
			{{ end }}
		}

		// Status returns the response status code.
		func (r *{{ $responseName }}) Status() int {
			if r == nil {
				return 0
			}
			return r.status
		}

		// Header returns header of the response.
		func (r *{{ $responseName }}) Header() http.Header {
			if r == nil {
				return nil
			}
			return r.header
		}

		// Error returns the response error.
		func (r *{{ $responseName }}) Error() *errors.Error {
			if r == nil {
				return nil
			}
			return r.err
		}

		{{ range $responseParameters }}
			{{ $fieldName := fieldName . }}
			{{ $getterName := getterName . }}
			{{ $getterType := getterType . }}

			// {{ $getterName }} returns the value of the '{{ .Name }}' parameter.
			//
			{{ lineComment .Doc }}
			func (r *{{ $responseName }}) {{ $getterName }}() {{ $getterType }} {
				{{ if .Type.IsScalar }}
					if r != nil && r.{{ $fieldName }} != nil {
						return *r.{{ $fieldName }}
					}
					return {{ zeroValue .Type }}
				{{ else if or .Type.IsStruct .Type.IsList .Type.IsMap }}
					if r == nil {
						return nil
					}
					return r.{{ $fieldName }}
				{{ end }}
			}

			// Get{{ $getterName }} returns the value of the '{{ .Name }}' parameter and
			// a flag indicating if the parameter has a value.
			//
			{{ lineComment .Doc }}
			func (r *{{ $responseName }}) Get{{ $getterName }}() (value {{ $getterType }}, ok bool) {
				ok = r != nil && r.{{ $fieldName }} != nil
				if ok {
					{{ if .Type.IsScalar }}
						value = *r.{{ $fieldName }}
					{{ else if or .Type.IsStruct .Type.IsList .Type.IsMap }}
						value = r.{{ $fieldName }}
					{{ end }}
				}
				return
			}
		{{ end }}
		`,
		"Method", method,
		"Main", main,
		"Others", others,
	)
}

func (g *ClientsGenerator) versionName(version *concepts.Version) string {
	name := annotations.GoName(version)
	if name == "" {
		name = g.names.Public(version.Name())
	}
	return name
}

func (g *ClientsGenerator) metadataFile() string {
	return g.names.File(names.Cat(nomenclator.Metadata, nomenclator.Client))
}

func (g *ClientsGenerator) resourceFile(resource *concepts.Resource) string {
	return g.names.File(names.Cat(resource.Name(), nomenclator.Client))
}

func (g *ClientsGenerator) enumName(typ *concepts.Type) string {
	name := annotations.GoName(typ)
	if name == "" {
		name = g.names.Public(typ.Name())
	}
	return name
}

func (g *ClientsGenerator) fieldName(parameter *concepts.Parameter) string {
	name := g.names.Private(parameter.Name())
	name = g.avoidBuiltin(name, builtinFields)
	return name
}

func (g *ClientsGenerator) fieldType(parameter *concepts.Parameter) *TypeReference {
	var ref *TypeReference
	typ := parameter.Type()
	switch {
	case parameter.IsItems():
		ref = g.types.ListReference(typ)
	case typ.IsScalar() || typ.IsStruct() || typ.IsList() || typ.IsMap():
		ref = g.types.NullableReference(typ, "")
	}
	if ref == nil {
		g.reporter.Errorf(
			"Don't know how to calculate field type for parameter '%s'",
			parameter,
		)
	}
	return ref
}

func (g *ClientsGenerator) getterName(parameter *concepts.Parameter) string {
	name := annotations.GoName(parameter)
	if name == "" {
		name = g.names.Public(parameter.Name())
	}
	name = g.avoidBuiltin(name, builtinGetters)
	return name
}

func (g *ClientsGenerator) getterType(parameter *concepts.Parameter) *TypeReference {
	return g.accessorType(parameter)
}

func (g *ClientsGenerator) setterName(parameter *concepts.Parameter) string {
	name := annotations.GoName(parameter)
	if name == "" {
		name = g.names.Public(parameter.Name())
	}
	name = g.avoidBuiltin(name, builtinSetters)
	return name
}

func (g *ClientsGenerator) setterType(parameter *concepts.Parameter) *TypeReference {
	return g.accessorType(parameter)
}

func (g *ClientsGenerator) accessorType(parameter *concepts.Parameter) *TypeReference {
	var ref *TypeReference
	typ := parameter.Type()
	switch {
	case typ.IsList() && typ.Element().IsScalar():
		ref = g.types.NullableReference(typ, "")
	case parameter.IsItems():
		ref = g.types.ListReference(typ)
	case typ.IsScalar():
		ref = g.types.ValueReference(typ)
	case typ.IsStruct() || typ.IsList() || typ.IsMap():
		ref = g.types.NullableReference(typ, "")
	}
	if ref == nil {
		g.reporter.Errorf(
			"Don't know how to calculate accessor type for parameter '%s'",
			parameter,
		)
	}
	return ref
}

func (g *ClientsGenerator) locatorName(locator *concepts.Locator) string {
	name := annotations.GoName(locator)
	if name == "" {
		name = g.names.Public(locator.Name())
	}
	return name
}

func (g *ClientsGenerator) methodName(method *concepts.Method) string {
	name := annotations.GoName(method)
	if name == "" {
		name = g.names.Public(method.Name())
	}
	return name
}

func (g *ClientsGenerator) clientName(resource *concepts.Resource) string {
	var name string
	if !resource.IsRoot() {
		name = annotations.GoName(resource)
		if name == "" {
			name = g.names.Public(resource.Name())
		}
	}
	name += "Client"
	return name
}

func (g *ClientsGenerator) requestName(method *concepts.Method) string {
	resource := method.Owner()
	var name string
	if resource.IsRoot() {
		name = annotations.GoName(method)
		if name == "" {
			name = g.names.Public(method.Name())
		}
	} else {
		resourceName := annotations.GoName(resource)
		if resourceName == "" {
			resourceName = g.names.Public(resource.Name())
		}
		methodName := annotations.GoName(method)
		if methodName == "" {
			methodName = g.names.Public(method.Name())
		}
		name = resourceName + methodName
	}
	name += "Request"
	return name
}

func (g *ClientsGenerator) responseName(method *concepts.Method) string {
	resource := method.Owner()
	var name string
	if resource.IsRoot() {
		name = annotations.GoName(method)
		if name == "" {
			name = g.names.Public(method.Name())
		}
	} else {
		resourceName := annotations.GoName(resource)
		if resourceName == "" {
			resourceName = g.names.Public(resource.Name())
		}
		methodName := annotations.GoName(method)
		if methodName == "" {
			methodName = g.names.Public(method.Name())
		}
		name = resourceName + methodName
	}
	name += "Response"
	return name
}

func (g *ClientsGenerator) pollRequestName(resource *concepts.Resource) string {
	name := annotations.GoName(resource)
	if name == "" {
		name = g.names.Public(resource.Name())
	}
	name += "PollRequest"
	return name
}

func (g *ClientsGenerator) pollResponseName(resource *concepts.Resource) string {
	name := annotations.GoName(resource)
	if name == "" {
		name = g.names.Public(resource.Name())
	}
	name += "PollResponse"
	return name
}

func (g *ClientsGenerator) writeRequestFunc(method *concepts.Method) string {
	resource := method.Owner()
	var name *names.Name
	if resource.IsRoot() {
		name = names.Cat(
			nomenclator.Write,
			method.Name(),
			nomenclator.Request,
		)
	} else {
		name = names.Cat(
			nomenclator.Write,
			resource.Name(),
			method.Name(),
			nomenclator.Request,
		)
	}
	return g.names.Private(name)
}

func (g *ClientsGenerator) readResponseFunc(method *concepts.Method) string {
	resource := method.Owner()
	var name *names.Name
	if resource.IsRoot() {
		name = names.Cat(
			nomenclator.Read,
			method.Name(),
			nomenclator.Response,
		)
	} else {
		name = names.Cat(
			nomenclator.Read,
			resource.Name(),
			method.Name(),
			nomenclator.Response,
		)
	}
	return g.names.Private(name)
}

func (g *ClientsGenerator) avoidBuiltin(name string, builtins map[string]interface{}) string {
	_, ok := builtins[name]
	if ok {
		name = name + "_"
	}
	return name
}

var builtinFields = map[string]interface{}{
	"err":    nil,
	"status": nil,
}

var builtinGetters = map[string]interface{}{
	"Error":  nil,
	"Status": nil,
}

var builtinSetters = map[string]interface{}{}
