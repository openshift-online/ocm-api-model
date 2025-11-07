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

// ControlPlaneBuilder contains the data and logic needed to build 'control_plane' objects.
//
// Representation of a Control Plane
type ControlPlaneBuilder struct {
	bitmap_ uint32
	backup  *BackupBuilder
}

// NewControlPlane creates a new builder of 'control_plane' objects.
func NewControlPlane() *ControlPlaneBuilder {
	return &ControlPlaneBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ControlPlaneBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Backup sets the value of the 'backup' attribute to the given value.
//
// Representation of a Backup.
func (b *ControlPlaneBuilder) Backup(value *BackupBuilder) *ControlPlaneBuilder {
	b.backup = value
	if value != nil {
		b.bitmap_ |= 1
	} else {
		b.bitmap_ &^= 1
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ControlPlaneBuilder) Copy(object *ControlPlane) *ControlPlaneBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	if object.backup != nil {
		b.backup = NewBackup().Copy(object.backup)
	} else {
		b.backup = nil
	}
	return b
}

// Build creates a 'control_plane' object using the configuration stored in the builder.
func (b *ControlPlaneBuilder) Build() (object *ControlPlane, err error) {
	object = new(ControlPlane)
	object.bitmap_ = b.bitmap_
	if b.backup != nil {
		object.backup, err = b.backup.Build()
		if err != nil {
			return
		}
	}
	return
}
