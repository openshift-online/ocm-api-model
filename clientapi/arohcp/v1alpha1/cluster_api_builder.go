/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1alpha1 // github.com/openshift-online/ocm-api-model/clientapi/arohcp/v1alpha1

// ClusterAPIBuilder contains the data and logic needed to build 'cluster_API' objects.
//
// Information about the API of a cluster.
type ClusterAPIBuilder struct {
	bitmap_         uint32
	cidrBlockAccess *CIDRBlockAccessBuilder
	url             string
	listening       ListeningMethod
}

// NewClusterAPI creates a new builder of 'cluster_API' objects.
func NewClusterAPI() *ClusterAPIBuilder {
	return &ClusterAPIBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ClusterAPIBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// CIDRBlockAccess sets the value of the 'CIDR_block_access' attribute to the given value.
//
// Describes the CIDR Block access policy to the Kubernetes API server.
// Currently, only supported for ARO-HCP based clusters.
// The default policy mode is "allow_all" that is, all access is allowed.
func (b *ClusterAPIBuilder) CIDRBlockAccess(value *CIDRBlockAccessBuilder) *ClusterAPIBuilder {
	b.cidrBlockAccess = value
	if value != nil {
		b.bitmap_ |= 1
	} else {
		b.bitmap_ &^= 1
	}
	return b
}

// URL sets the value of the 'URL' attribute to the given value.
func (b *ClusterAPIBuilder) URL(value string) *ClusterAPIBuilder {
	b.url = value
	b.bitmap_ |= 2
	return b
}

// Listening sets the value of the 'listening' attribute to the given value.
//
// Cluster components listening method.
func (b *ClusterAPIBuilder) Listening(value ListeningMethod) *ClusterAPIBuilder {
	b.listening = value
	b.bitmap_ |= 4
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterAPIBuilder) Copy(object *ClusterAPI) *ClusterAPIBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	if object.cidrBlockAccess != nil {
		b.cidrBlockAccess = NewCIDRBlockAccess().Copy(object.cidrBlockAccess)
	} else {
		b.cidrBlockAccess = nil
	}
	b.url = object.url
	b.listening = object.listening
	return b
}

// Build creates a 'cluster_API' object using the configuration stored in the builder.
func (b *ClusterAPIBuilder) Build() (object *ClusterAPI, err error) {
	object = new(ClusterAPI)
	object.bitmap_ = b.bitmap_
	if b.cidrBlockAccess != nil {
		object.cidrBlockAccess, err = b.cidrBlockAccess.Build()
		if err != nil {
			return
		}
	}
	object.url = b.url
	object.listening = b.listening
	return
}
