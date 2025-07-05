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

// AzureKmsKeyBuilder contains the data and logic needed to build 'azure_kms_key' objects.
//
// Contains the necessary attributes to support KMS encryption key for Azure based clusters
type AzureKmsKeyBuilder struct {
	bitmap_      uint32
	keyName      string
	keyVaultName string
	keyVersion   string
}

// NewAzureKmsKey creates a new builder of 'azure_kms_key' objects.
func NewAzureKmsKey() *AzureKmsKeyBuilder {
	return &AzureKmsKeyBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AzureKmsKeyBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// KeyName sets the value of the 'key_name' attribute to the given value.
func (b *AzureKmsKeyBuilder) KeyName(value string) *AzureKmsKeyBuilder {
	b.keyName = value
	b.bitmap_ |= 1
	return b
}

// KeyVaultName sets the value of the 'key_vault_name' attribute to the given value.
func (b *AzureKmsKeyBuilder) KeyVaultName(value string) *AzureKmsKeyBuilder {
	b.keyVaultName = value
	b.bitmap_ |= 2
	return b
}

// KeyVersion sets the value of the 'key_version' attribute to the given value.
func (b *AzureKmsKeyBuilder) KeyVersion(value string) *AzureKmsKeyBuilder {
	b.keyVersion = value
	b.bitmap_ |= 4
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AzureKmsKeyBuilder) Copy(object *AzureKmsKey) *AzureKmsKeyBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.keyName = object.keyName
	b.keyVaultName = object.keyVaultName
	b.keyVersion = object.keyVersion
	return b
}

// Build creates a 'azure_kms_key' object using the configuration stored in the builder.
func (b *AzureKmsKeyBuilder) Build() (object *AzureKmsKey, err error) {
	object = new(AzureKmsKey)
	object.bitmap_ = b.bitmap_
	object.keyName = b.keyName
	object.keyVaultName = b.keyVaultName
	object.keyVersion = b.keyVersion
	return
}
