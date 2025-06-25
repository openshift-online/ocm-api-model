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

// AWSBackupConfigBuilder contains the data and logic needed to build 'AWS_backup_config' objects.
//
// Backup configuration for AWS clusters
type AWSBackupConfigBuilder struct {
	bitmap_             uint32
	s3Bucket            string
	accountId           string
	identityProviderArn string
	roleArn             string
}

// NewAWSBackupConfig creates a new builder of 'AWS_backup_config' objects.
func NewAWSBackupConfig() *AWSBackupConfigBuilder {
	return &AWSBackupConfigBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AWSBackupConfigBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// S3Bucket sets the value of the 'S3_bucket' attribute to the given value.
func (b *AWSBackupConfigBuilder) S3Bucket(value string) *AWSBackupConfigBuilder {
	b.s3Bucket = value
	b.bitmap_ |= 1
	return b
}

// AccountId sets the value of the 'account_id' attribute to the given value.
func (b *AWSBackupConfigBuilder) AccountId(value string) *AWSBackupConfigBuilder {
	b.accountId = value
	b.bitmap_ |= 2
	return b
}

// IdentityProviderArn sets the value of the 'identity_provider_arn' attribute to the given value.
func (b *AWSBackupConfigBuilder) IdentityProviderArn(value string) *AWSBackupConfigBuilder {
	b.identityProviderArn = value
	b.bitmap_ |= 4
	return b
}

// RoleArn sets the value of the 'role_arn' attribute to the given value.
func (b *AWSBackupConfigBuilder) RoleArn(value string) *AWSBackupConfigBuilder {
	b.roleArn = value
	b.bitmap_ |= 8
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AWSBackupConfigBuilder) Copy(object *AWSBackupConfig) *AWSBackupConfigBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.s3Bucket = object.s3Bucket
	b.accountId = object.accountId
	b.identityProviderArn = object.identityProviderArn
	b.roleArn = object.roleArn
	return b
}

// Build creates a 'AWS_backup_config' object using the configuration stored in the builder.
func (b *AWSBackupConfigBuilder) Build() (object *AWSBackupConfig, err error) {
	object = new(AWSBackupConfig)
	object.bitmap_ = b.bitmap_
	object.s3Bucket = b.s3Bucket
	object.accountId = b.accountId
	object.identityProviderArn = b.identityProviderArn
	object.roleArn = b.roleArn
	return
}
