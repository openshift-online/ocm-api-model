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

// AzureNodePoolOsDiskBuilder contains the data and logic needed to build 'azure_node_pool_os_disk' objects.
//
// Defines the configuration of a Node Pool's OS disk.
type AzureNodePoolOsDiskBuilder struct {
	bitmap_                    uint32
	persistence                string
	sizeGibibytes              int
	sseEncryptionSetResourceId string
	storageAccountType         string
}

// NewAzureNodePoolOsDisk creates a new builder of 'azure_node_pool_os_disk' objects.
func NewAzureNodePoolOsDisk() *AzureNodePoolOsDiskBuilder {
	return &AzureNodePoolOsDiskBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AzureNodePoolOsDiskBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Persistence sets the value of the 'persistence' attribute to the given value.
func (b *AzureNodePoolOsDiskBuilder) Persistence(value string) *AzureNodePoolOsDiskBuilder {
	b.persistence = value
	b.bitmap_ |= 1
	return b
}

// SizeGibibytes sets the value of the 'size_gibibytes' attribute to the given value.
func (b *AzureNodePoolOsDiskBuilder) SizeGibibytes(value int) *AzureNodePoolOsDiskBuilder {
	b.sizeGibibytes = value
	b.bitmap_ |= 2
	return b
}

// SseEncryptionSetResourceId sets the value of the 'sse_encryption_set_resource_id' attribute to the given value.
func (b *AzureNodePoolOsDiskBuilder) SseEncryptionSetResourceId(value string) *AzureNodePoolOsDiskBuilder {
	b.sseEncryptionSetResourceId = value
	b.bitmap_ |= 4
	return b
}

// StorageAccountType sets the value of the 'storage_account_type' attribute to the given value.
func (b *AzureNodePoolOsDiskBuilder) StorageAccountType(value string) *AzureNodePoolOsDiskBuilder {
	b.storageAccountType = value
	b.bitmap_ |= 8
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AzureNodePoolOsDiskBuilder) Copy(object *AzureNodePoolOsDisk) *AzureNodePoolOsDiskBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.persistence = object.persistence
	b.sizeGibibytes = object.sizeGibibytes
	b.sseEncryptionSetResourceId = object.sseEncryptionSetResourceId
	b.storageAccountType = object.storageAccountType
	return b
}

// Build creates a 'azure_node_pool_os_disk' object using the configuration stored in the builder.
func (b *AzureNodePoolOsDiskBuilder) Build() (object *AzureNodePoolOsDisk, err error) {
	object = new(AzureNodePoolOsDisk)
	object.bitmap_ = b.bitmap_
	object.persistence = b.persistence
	object.sizeGibibytes = b.sizeGibibytes
	object.sseEncryptionSetResourceId = b.sseEncryptionSetResourceId
	object.storageAccountType = b.storageAccountType
	return
}
