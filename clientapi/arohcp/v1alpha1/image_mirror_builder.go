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

import (
	time "time"
)

// ImageMirrorBuilder contains the data and logic needed to build 'image_mirror' objects.
//
// ImageMirror represents a container image mirror configuration for a cluster.
// This enables Day 2 image mirroring configuration for ROSA HCP clusters using
// HyperShift's native imageContentSources mechanism.
type ImageMirrorBuilder struct {
	bitmap_             uint32
	id                  string
	href                string
	creationTimestamp   time.Time
	lastUpdateTimestamp time.Time
	mirrors             []string
	source              string
	type_               string
}

// NewImageMirror creates a new builder of 'image_mirror' objects.
func NewImageMirror() *ImageMirrorBuilder {
	return &ImageMirrorBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *ImageMirrorBuilder) Link(value bool) *ImageMirrorBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *ImageMirrorBuilder) ID(value string) *ImageMirrorBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *ImageMirrorBuilder) HREF(value string) *ImageMirrorBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ImageMirrorBuilder) Empty() bool {
	return b == nil || b.bitmap_&^1 == 0
}

// CreationTimestamp sets the value of the 'creation_timestamp' attribute to the given value.
func (b *ImageMirrorBuilder) CreationTimestamp(value time.Time) *ImageMirrorBuilder {
	b.creationTimestamp = value
	b.bitmap_ |= 8
	return b
}

// LastUpdateTimestamp sets the value of the 'last_update_timestamp' attribute to the given value.
func (b *ImageMirrorBuilder) LastUpdateTimestamp(value time.Time) *ImageMirrorBuilder {
	b.lastUpdateTimestamp = value
	b.bitmap_ |= 16
	return b
}

// Mirrors sets the value of the 'mirrors' attribute to the given values.
func (b *ImageMirrorBuilder) Mirrors(values ...string) *ImageMirrorBuilder {
	b.mirrors = make([]string, len(values))
	copy(b.mirrors, values)
	b.bitmap_ |= 32
	return b
}

// Source sets the value of the 'source' attribute to the given value.
func (b *ImageMirrorBuilder) Source(value string) *ImageMirrorBuilder {
	b.source = value
	b.bitmap_ |= 64
	return b
}

// Type sets the value of the 'type' attribute to the given value.
func (b *ImageMirrorBuilder) Type(value string) *ImageMirrorBuilder {
	b.type_ = value
	b.bitmap_ |= 128
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ImageMirrorBuilder) Copy(object *ImageMirror) *ImageMirrorBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	b.creationTimestamp = object.creationTimestamp
	b.lastUpdateTimestamp = object.lastUpdateTimestamp
	if object.mirrors != nil {
		b.mirrors = make([]string, len(object.mirrors))
		copy(b.mirrors, object.mirrors)
	} else {
		b.mirrors = nil
	}
	b.source = object.source
	b.type_ = object.type_
	return b
}

// Build creates a 'image_mirror' object using the configuration stored in the builder.
func (b *ImageMirrorBuilder) Build() (object *ImageMirror, err error) {
	object = new(ImageMirror)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	object.creationTimestamp = b.creationTimestamp
	object.lastUpdateTimestamp = b.lastUpdateTimestamp
	if b.mirrors != nil {
		object.mirrors = make([]string, len(b.mirrors))
		copy(object.mirrors, b.mirrors)
	}
	object.source = b.source
	object.type_ = b.type_
	return
}
