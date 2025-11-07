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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1

// ClusterAutoNodeStatusBuilder contains the data and logic needed to build 'cluster_auto_node_status' objects.
//
// Additional status information on the AutoNode configuration on this Cluster
type ClusterAutoNodeStatusBuilder struct {
	bitmap_ uint32
	message string
}

// NewClusterAutoNodeStatus creates a new builder of 'cluster_auto_node_status' objects.
func NewClusterAutoNodeStatus() *ClusterAutoNodeStatusBuilder {
	return &ClusterAutoNodeStatusBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ClusterAutoNodeStatusBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Message sets the value of the 'message' attribute to the given value.
func (b *ClusterAutoNodeStatusBuilder) Message(value string) *ClusterAutoNodeStatusBuilder {
	b.message = value
	b.bitmap_ |= 1
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterAutoNodeStatusBuilder) Copy(object *ClusterAutoNodeStatus) *ClusterAutoNodeStatusBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.message = object.message
	return b
}

// Build creates a 'cluster_auto_node_status' object using the configuration stored in the builder.
func (b *ClusterAutoNodeStatusBuilder) Build() (object *ClusterAutoNodeStatus, err error) {
	object = new(ClusterAutoNodeStatus)
	object.bitmap_ = b.bitmap_
	object.message = b.message
	return
}
