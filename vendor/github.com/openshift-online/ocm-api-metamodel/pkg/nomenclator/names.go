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

package nomenclator

import (
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
)

var (
	// A:
	Adapt   = names.ParseUsingCase("Adapt")
	Adapter = names.ParseUsingCase("Adapter")
	Add     = names.ParseUsingCase("Add")

	// B:
	Body    = names.ParseUsingCase("Body")
	Boolean = names.ParseUsingCase("Boolean")
	Builder = names.ParseUsingCase("Builder")

	// C:
	Client  = names.ParseUsingCase("Client")
	Clients = names.ParseUsingCase("Clients")

	// D:
	Data     = names.ParseUsingCase("Data")
	Date     = names.ParseUsingCase("Date")
	Delete   = names.ParseUsingCase("Delete")
	Dispatch = names.ParseUsingCase("Dispatch")

	// E:
	Error  = names.ParseUsingCase("Error")
	Errors = names.ParseUsingCase("Errors")

	// F:
	Float = names.ParseUsingCase("Float")

	// G:
	Get = names.ParseUsingCase("Get")

	// H:
	HREF    = names.ParseUsingCase("HREF")
	Handler = names.ParseUsingCase("Handler")
	Helpers = names.ParseUsingCase("Helpers")

	// I:
	ID        = names.ParseUsingCase("ID")
	Index     = names.ParseUsingCase("Index")
	Integer   = names.ParseUsingCase("Integer")
	Interface = names.ParseUsingCase("Interface")
	Items     = names.ParseUsingCase("Items")

	// J:
	JSON = names.ParseUsingCase("JSON")

	// K:
	Kind = names.ParseUsingCase("Kind")

	// L:
	Link = names.ParseUsingCase("Link")
	List = names.ParseUsingCase("List")
	Long = names.ParseUsingCase("Long")

	// M:
	Map      = names.ParseUsingCase("Map")
	Marshal  = names.ParseUsingCase("Marshal")
	Metadata = names.ParseUsingCase("Metadata")
	Method   = names.ParseUsingCase("Method")
	Metrics  = names.ParseUsingCase("Metrics")

	// N:
	New = names.ParseUsingCase("New")

	// P:
	Page  = names.ParseUsingCase("Page")
	Parse = names.ParseUsingCase("Parse")
	Path  = names.ParseUsingCase("Path")
	Poll  = names.ParseUsingCase("Poll")
	Post  = names.ParseUsingCase("Post")

	// R:
	Read     = names.ParseUsingCase("Read")
	Reader   = names.ParseUsingCase("Reader")
	Readers  = names.ParseUsingCase("Readers")
	Request  = names.ParseUsingCase("Request")
	Resource = names.ParseUsingCase("Resource")
	Response = names.ParseUsingCase("Response")
	Root     = names.ParseUsingCase("Root")

	// S:
	Search  = names.ParseUsingCase("Search")
	Server  = names.ParseUsingCase("Server")
	Servers = names.ParseUsingCase("Servers")
	Service = names.ParseUsingCase("Service")
	Set     = names.ParseUsingCase("Set")
	Size    = names.ParseUsingCase("Size")
	Spec    = names.ParseUsingCase("Spec")
	Stream  = names.ParseUsingCase("Stream")
	String  = names.ParseUsingCase("String")

	// T:
	Total = names.ParseUsingCase("Total")
	Tree  = names.ParseUsingCase("Tree")
	Type  = names.ParseUsingCase("Type")

	// U:
	Unmarshal = names.ParseUsingCase("Unmarshal")
	Unwrap    = names.ParseUsingCase("Unwrap")
	Update    = names.ParseUsingCase("Update")

	// W:
	Wrap  = names.ParseUsingCase("Wrap")
	Write = names.ParseUsingCase("Write")
)
