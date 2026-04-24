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

// Configuration for zero egress clusters.
type ZeroEgressConfigurationBuilder struct {
	fieldSet_      []bool
	defaultDomains []string
}

// NewZeroEgressConfiguration creates a new builder of 'zero_egress_configuration' objects.
func NewZeroEgressConfiguration() *ZeroEgressConfigurationBuilder {
	return &ZeroEgressConfigurationBuilder{
		fieldSet_: make([]bool, 1),
	}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ZeroEgressConfigurationBuilder) Empty() bool {
	if b == nil || len(b.fieldSet_) == 0 {
		return true
	}
	for _, set := range b.fieldSet_ {
		if set {
			return false
		}
	}
	return true
}

// DefaultDomains sets the value of the 'default_domains' attribute to the given values.
func (b *ZeroEgressConfigurationBuilder) DefaultDomains(values ...string) *ZeroEgressConfigurationBuilder {
	if len(b.fieldSet_) == 0 {
		b.fieldSet_ = make([]bool, 1)
	}
	b.defaultDomains = make([]string, len(values))
	copy(b.defaultDomains, values)
	b.fieldSet_[0] = true
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ZeroEgressConfigurationBuilder) Copy(object *ZeroEgressConfiguration) *ZeroEgressConfigurationBuilder {
	if object == nil {
		return b
	}
	if len(object.fieldSet_) > 0 {
		b.fieldSet_ = make([]bool, len(object.fieldSet_))
		copy(b.fieldSet_, object.fieldSet_)
	}
	if object.defaultDomains != nil {
		b.defaultDomains = make([]string, len(object.defaultDomains))
		copy(b.defaultDomains, object.defaultDomains)
	} else {
		b.defaultDomains = nil
	}
	return b
}

// Build creates a 'zero_egress_configuration' object using the configuration stored in the builder.
func (b *ZeroEgressConfigurationBuilder) Build() (object *ZeroEgressConfiguration, err error) {
	object = new(ZeroEgressConfiguration)
	if len(b.fieldSet_) > 0 {
		object.fieldSet_ = make([]bool, len(b.fieldSet_))
		copy(object.fieldSet_, b.fieldSet_)
	}
	if b.defaultDomains != nil {
		object.defaultDomains = make([]string, len(b.defaultDomains))
		copy(object.defaultDomains, b.defaultDomains)
	}
	return
}
