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

// MachineTypeFeaturesBuilder contains the data and logic needed to build 'machine_type_features' objects.
//
// Defines the features enabled or disabled on a particular machine type
type MachineTypeFeaturesBuilder struct {
	bitmap_ uint32
	winLI   bool
}

// NewMachineTypeFeatures creates a new builder of 'machine_type_features' objects.
func NewMachineTypeFeatures() *MachineTypeFeaturesBuilder {
	return &MachineTypeFeaturesBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *MachineTypeFeaturesBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// WinLI sets the value of the 'win_LI' attribute to the given value.
func (b *MachineTypeFeaturesBuilder) WinLI(value bool) *MachineTypeFeaturesBuilder {
	b.winLI = value
	b.bitmap_ |= 1
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *MachineTypeFeaturesBuilder) Copy(object *MachineTypeFeatures) *MachineTypeFeaturesBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.winLI = object.winLI
	return b
}

// Build creates a 'machine_type_features' object using the configuration stored in the builder.
func (b *MachineTypeFeaturesBuilder) Build() (object *MachineTypeFeatures, err error) {
	object = new(MachineTypeFeatures)
	object.bitmap_ = b.bitmap_
	object.winLI = b.winLI
	return
}
