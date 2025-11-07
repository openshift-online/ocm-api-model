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

// AzureEtcdEncryptionBuilder contains the data and logic needed to build 'azure_etcd_encryption' objects.
//
// Contains the necessary attributes to support etcd encryption for Azure based clusters.
type AzureEtcdEncryptionBuilder struct {
	bitmap_        uint32
	dataEncryption *AzureEtcdDataEncryptionBuilder
}

// NewAzureEtcdEncryption creates a new builder of 'azure_etcd_encryption' objects.
func NewAzureEtcdEncryption() *AzureEtcdEncryptionBuilder {
	return &AzureEtcdEncryptionBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AzureEtcdEncryptionBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// DataEncryption sets the value of the 'data_encryption' attribute to the given value.
//
// Contains the necessary attributes to support data encryption for Azure based clusters.
func (b *AzureEtcdEncryptionBuilder) DataEncryption(value *AzureEtcdDataEncryptionBuilder) *AzureEtcdEncryptionBuilder {
	b.dataEncryption = value
	if value != nil {
		b.bitmap_ |= 1
	} else {
		b.bitmap_ &^= 1
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AzureEtcdEncryptionBuilder) Copy(object *AzureEtcdEncryption) *AzureEtcdEncryptionBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	if object.dataEncryption != nil {
		b.dataEncryption = NewAzureEtcdDataEncryption().Copy(object.dataEncryption)
	} else {
		b.dataEncryption = nil
	}
	return b
}

// Build creates a 'azure_etcd_encryption' object using the configuration stored in the builder.
func (b *AzureEtcdEncryptionBuilder) Build() (object *AzureEtcdEncryption, err error) {
	object = new(AzureEtcdEncryption)
	object.bitmap_ = b.bitmap_
	if b.dataEncryption != nil {
		object.dataEncryption, err = b.dataEncryption.Build()
		if err != nil {
			return
		}
	}
	return
}
