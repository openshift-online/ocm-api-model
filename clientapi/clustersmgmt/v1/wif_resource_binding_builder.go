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

// WifResourceBindingBuilder contains the data and logic needed to build 'wif_resource_binding' objects.
type WifResourceBindingBuilder struct {
	bitmap_ uint32
	name    string
	type_   string
}

// NewWifResourceBinding creates a new builder of 'wif_resource_binding' objects.
func NewWifResourceBinding() *WifResourceBindingBuilder {
	return &WifResourceBindingBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *WifResourceBindingBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Name sets the value of the 'name' attribute to the given value.
func (b *WifResourceBindingBuilder) Name(value string) *WifResourceBindingBuilder {
	b.name = value
	b.bitmap_ |= 1
	return b
}

// Type sets the value of the 'type' attribute to the given value.
func (b *WifResourceBindingBuilder) Type(value string) *WifResourceBindingBuilder {
	b.type_ = value
	b.bitmap_ |= 2
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *WifResourceBindingBuilder) Copy(object *WifResourceBinding) *WifResourceBindingBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.name = object.name
	b.type_ = object.type_
	return b
}

// Build creates a 'wif_resource_binding' object using the configuration stored in the builder.
func (b *WifResourceBindingBuilder) Build() (object *WifResourceBinding, err error) {
	object = new(WifResourceBinding)
	object.bitmap_ = b.bitmap_
	object.name = b.name
	object.type_ = b.type_
	return
}
