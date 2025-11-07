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

// ClusterImageRegistryBuilder contains the data and logic needed to build 'cluster_image_registry' objects.
//
// ClusterImageRegistry represents the configuration for the cluster's internal image registry.
type ClusterImageRegistryBuilder struct {
	bitmap_ uint32
	state   string
}

// NewClusterImageRegistry creates a new builder of 'cluster_image_registry' objects.
func NewClusterImageRegistry() *ClusterImageRegistryBuilder {
	return &ClusterImageRegistryBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ClusterImageRegistryBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// State sets the value of the 'state' attribute to the given value.
func (b *ClusterImageRegistryBuilder) State(value string) *ClusterImageRegistryBuilder {
	b.state = value
	b.bitmap_ |= 1
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterImageRegistryBuilder) Copy(object *ClusterImageRegistry) *ClusterImageRegistryBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.state = object.state
	return b
}

// Build creates a 'cluster_image_registry' object using the configuration stored in the builder.
func (b *ClusterImageRegistryBuilder) Build() (object *ClusterImageRegistry, err error) {
	object = new(ClusterImageRegistry)
	object.bitmap_ = b.bitmap_
	object.state = b.state
	return
}
