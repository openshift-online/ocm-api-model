module github.com/openshift-online/ocm-api-model/clientapi

go 1.23.0

// this must not transitively depend on metamodel

replace github.com/openshift-online/ocm-api-model/model => ../model

require (
	github.com/json-iterator/go v1.1.12
	github.com/openshift-online/ocm-api-model/model v0.0.0-00010101000000-000000000000
)

require (
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
)
