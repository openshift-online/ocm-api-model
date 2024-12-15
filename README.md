# OpenShift cluster manager API model

## Introduction

This project contains the specification of the OpenShift cluster manager API,
also known as the _model_.

## Concepts

The specification of the API is written using a DSL (the model language)
similar to _Go_.

The specification is split into multiple _model_ files with extension `.model`
inside the `model` directory of this project. The location of the model files
within sub-directories is important as it indicates which service and version a
type or resource belongs to. For example, the `clusters_mgmt/v1` sub-directory
corresponds to version 1 of the clusters management service.

Data types are represented by `class`, `struct` and `enum` blocks.

For example, the `Cluster` type defining the _cluster_ concept should be located
in a file within the `clusters_mgmt/v1` directory, should be named
`cluster_type.model` and should contain something like this:

```
// Definition of an _OpenShift_ cluster.
class Cluster {
	// Name of the cluster.
	Name String

	// Link to the _flavour_ that was used to create the cluster.
	link Flavour Flavour

	...
}
```

Classes and structs contain _attributes_ defined by an attribute name followed
by the type of that attribute. In the above example there are two attributes
defined: `Name` of type string and `Flavour` of type `Flavour`.

In the model language attributes are written using _CamelCase_ because that is
what is typically used in Go, which is the language that the model language
tries to be close to. But the actual JSON representation uses _snake_case_.
These are some examples of what is the correspondence between attribute names
in the model and in JSON:

| Model         | JSON            |
| ------------- | --------------- |
| `ID`          | `id`            |
| `HREF`        | `href`          |
| `Cluster`     | `cluster`       |
| `AccessKeyID` | `access_key_id` |

The main difference between classes and structs is that classes are intended to
represent objects that have an _identity_. In practice that means that classes
have built-in `ID` and `HREF` attributes. The `ID` attribute will contain the
unique identifier of the object, and the `HREF` attribute will contain the
location of the object in the server. For example, when retrieving a particular
cluster from the server the returned JSON document will be like this:

```json
{
    "id": "123",
    "href": "/api/clusters_mgmt/v1/clusters/123",
    ...
}
```

Resources are represented by `resource` blocks.

For example, the `Clusters` resource defining the collection of clusters should
be located in a file within the `clusters_mgmt/v1` directory, should be named
`clusters_resource.model` and should contain something like this:

```
// Manages the collection of clusters.
resource Clusters {
	// Retrieves the list of clusters.
	method List {
		...
	}

	// Provision a new cluster and add it to the collection of clusters.
	method Add {
		...
	}

	...
}
```

Resource _methods_ are represented as nested `method` blocks.

For example, the method that retrieves the list of clusters is defined in a
nested `List` method block inside the `resource` block for the clusters
resource:

```
// Retrieves the list of clusters.
method List {
	// Index of the requested page, where one corresponds to the first page.
	in out Page Integer

	// Maximum number of items that will be contained in the returned page.
	in out Size Integer

	// Search criteria.
	in Search String

	// Order criteria.
	in Order String

	// Total number of items of the collection that match the search criteria,
	// regardless of the size of the page.
	out Total Integer

	// Retrieved list of clusters.
	out Items []Cluster
}
```

Methods have _parameters_ defined by their direction (_in_ or _out_), their
name and their type. In the above example there are four input parameters
(named `Page`, `Size`, `Search` and `Order`) and four output parameter (named
`Page`, `Size`, `Total` and `Items`).

- "List" conceptual method is implemented by HTTP GET method, with
  _in_ parameters represented in the URL as HTTP query parameters
  (with names converted from _CamelCase_ to _snake_case_).

  `cluster_mgmt` API supports an alternative way to List — as HTTP
  POST method, with a `method=get` query parameter and _in_ paramaters
  sent as fields of a JSON request body (again, with _snake_case_ names).
 
  _Out_ parameters become top-level fields in the JSON response body
  (again, with _snake_case_ names).
  An additional `kind` field is added automatically (set to the type's
  name + a "List" suffix), it should not be included in the method block.

- "Get" is regular HTTP GET method, declared with single _out_ parameter
  representing the JSON response body.

- "Add" is performed by HTTP POST method, declared with single _in
  out_ parameter representing the JSON request body — as well as the
  response body.

- "Delete" is performed by HTTP DELETE method, with no parameters.

In addition to methods resources also have _locators_, defined by nested
`locator` blocks. Locators represent the relationships between resources. For
example, the resource that manages the collection of clusters knows how to
_locate_ the resource that manages a specific cluster. That is represented in
the model language with a `locator` block like this:

```
// Returns a reference to the service that manages an specific cluster.
locator Cluster {
	target Cluster
	variable ID
}
```

All resource locators have a name and a _target_. The target is defined using
the `target` keyword and the name of the resource.

There are two kinds of resource locators: with and without variable.

Locators with variable are intended for collections, where location a
sub-resource resource requires specifying the identifier of that object that is
managed by that sub-resource. For example, to locate the sub-resource that
manages a specific cluster it is necessary to provide the identifier of that
cluster.  That identifier is the _variable_. These kind of locators are defined
using the `variable` keyword and the name of the variable, like in the previous
example.

Locators without variable are intended for cases where no additional
information is needed to identify the sub-resource. For example, the locator
for the credentials sub-resource of a cluster can be defined like this:

```
// Reference to the resource that manages the credentials of the cluster.
locator Credentials {
	target Credentials
}
```

Locators also define the URL structure of the API: the path component of the
URL of a particular resource is constructed concatenating the names/variables
of the locators that are in the chain of locators from the root to that
resource. For example, to get to the `Credentials` resource of cluster with
identifier `123` the chain starts at the root resource of the clusters
management service, it continues with the `Clusters` locator to find the
clusters collection, then the `Cluster` locator with variable `123` to get the
cluster resource and finally the `Credentials` locator to get to the
credentials resource:

```
Root -> Clusters -> Cluster(123) -> Credentials
```

Each link in that chain of locators is translated into an URL path segment
using the following rules:

- The root resource corresponds to the root of the service/version. For
  example, for version 1 of the clusters management service that would be
  `/api/clusters_mgmt/v1`.

- Locators without variables correspond to URL segments named like the locator,
  but using _snake_case_ instead of _CamelCase_. For example, for the first
  link in the above example the URL path segment would be `clusters`.

- Locators with variables correspond to URL segments that contain the value of
  the variable. For example, for the second link in the above example the URL
  path segment would be `123`.

Taking these rules into account the complete URL path for the above example
would be the following:

```
/api/clusters_mgmt/v1/clusters/123/credentials
```

### Class references using @ref annotation
One can define a class reference inside a service such that it will inherit its content from
another service using the special `@ref` annotation, for example:

```
// In /aro_hcp/v1_alpha1/cluster_type.model
@ref(path = "/clusters_mgmt/v1/cluster")
class Cluster {
}
```

The above decelaration will inherit its content from the `Cluster` class under the `/clusters_mgmt/v1` service.
This means that any changes made in `Cluster` class will be reflected under this derived type as well.

Links to other resources are preserved as they are.

If one wishes to **override** a type she can create a class inside `/aro_hcp/v1_alpha1/` in order to override any nested types defined.
For example the following type declaration will override the `NodePools` link inside `Cluster` under `/aro_hcp/v1`:

```
// In /aro_hcp/v1_alpha1/node_pool_type.model
@ref(path = "/clusters_mgmt/v1/node_pool")
class NodePool {
}
```

This means that now under `Cluster` the `NodePools` field type is linked to the one defined in `/aro_hcp/v1_alpha1` (i.e. `v1alpha.NodePool`) which itself
references and derived from the `NodePool` type under `/clusters_mgmt/v1`.

## Documentation

The Go language supports adding documentation in the code itself, using the
documentation comments. These comments start with `//` and appear immediately
before the documented item. The model language uses the same kind of
documentation comments. For example, the `Cluster` type can be documented
like this:

```
// Definition of an _OpenShift_ cluster.
//
// The `cloud_provider` attribute is a reference to the cloud provider. When a
// cluster is retrieved it will be a link to the cloud provider, containing only
// the kind, id and href attributes:
//
// ```json
// {
//   "cloud_provider": {
//     "kind": "CloudProviderLink",
//     "id": "123",
//     "href": "/api/clusters_mgmt/v1/cloud_providers/123"
//   }
// }
// ```
//
// When a cluster is created this is optional, and if used it should contain the
// identifier of the cloud provider to use:
//
// ```
// {
//   "cloud_provider": {
//     "id": "123",
//   }
// }
// ```
//
// If not included, then the cluster will be created using the default cloud
// provider, which is currently Amazon Web Services.
//
// The region attribute is mandatory when a cluster is created.
//
// The `aws.access_key_id`, `aws.secret_access_key` and `dns.base_domain`
// attributes are mandatory when creation a cluster with your own Amazon Web
// Services account.
class Cluster {
	...
}
```

Unlike Go the format of this documentation isn't plain text, but
[Markdown](https://daringfireball.net/projects/markdown/syntax).

Attributes of types, methods of resources and parameters of methods can all be
documented in a similar way, just placing documentation comment before the
definition of the item. For example, to document the `Search` parameter of the
`List` method of the `Clusters` resource the following documentation comment
could be used:

```
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause of a
// SQL statement, but using the names of the attributes of the cluster instead of
// the names of the columns of a table. For example, in order to retrieve all the
// clusters with a name starting with `my` in the `us-east-1` region the value
// should be:
//
// ```sql
// name like 'my%' and region.id = 'us-east-1'
// ```
//
// If the parameter isn't provided, or if the value is empty, then all the
// clusters that the user has permission to see will be returned.
in Search String
```

This documentation is used to automatically generate OpenAPI reference
documentation.
