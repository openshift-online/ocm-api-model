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

// AzureEtcdDataEncryptionCustomerManagedBuilder contains the data and logic needed to build 'azure_etcd_data_encryption_customer_managed' objects.
//
// Contains the necessary attributes to support etcd data encryption with customer managed keys
// for Azure based clusters.
type AzureEtcdDataEncryptionCustomerManagedBuilder struct {
	bitmap_        uint32
	encryptionType string
	kms            *AzureKmsEncryptionBuilder
}

// NewAzureEtcdDataEncryptionCustomerManaged creates a new builder of 'azure_etcd_data_encryption_customer_managed' objects.
func NewAzureEtcdDataEncryptionCustomerManaged() *AzureEtcdDataEncryptionCustomerManagedBuilder {
	return &AzureEtcdDataEncryptionCustomerManagedBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AzureEtcdDataEncryptionCustomerManagedBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// EncryptionType sets the value of the 'encryption_type' attribute to the given value.
func (b *AzureEtcdDataEncryptionCustomerManagedBuilder) EncryptionType(value string) *AzureEtcdDataEncryptionCustomerManagedBuilder {
	b.encryptionType = value
	b.bitmap_ |= 1
	return b
}

// Kms sets the value of the 'kms' attribute to the given value.
//
// Contains the necessary attributes to support KMS encryption for Azure based clusters.
func (b *AzureEtcdDataEncryptionCustomerManagedBuilder) Kms(value *AzureKmsEncryptionBuilder) *AzureEtcdDataEncryptionCustomerManagedBuilder {
	b.kms = value
	if value != nil {
		b.bitmap_ |= 2
	} else {
		b.bitmap_ &^= 2
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AzureEtcdDataEncryptionCustomerManagedBuilder) Copy(object *AzureEtcdDataEncryptionCustomerManaged) *AzureEtcdDataEncryptionCustomerManagedBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.encryptionType = object.encryptionType
	if object.kms != nil {
		b.kms = NewAzureKmsEncryption().Copy(object.kms)
	} else {
		b.kms = nil
	}
	return b
}

// Build creates a 'azure_etcd_data_encryption_customer_managed' object using the configuration stored in the builder.
func (b *AzureEtcdDataEncryptionCustomerManagedBuilder) Build() (object *AzureEtcdDataEncryptionCustomerManaged, err error) {
	object = new(AzureEtcdDataEncryptionCustomerManaged)
	object.bitmap_ = b.bitmap_
	object.encryptionType = b.encryptionType
	if b.kms != nil {
		object.kms, err = b.kms.Build()
		if err != nil {
			return
		}
	}
	return
}
