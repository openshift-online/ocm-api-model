module github.com/openshift-online/ocm-api-model/clientapi

go 1.23.0

// this must not transitively depend on metamodel and must be as close to zero dep as possible.
// we can later eliminate json-iter

require github.com/json-iterator/go v1.1.12

require (
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
)
