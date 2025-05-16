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
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalControlPlaneOperatorIdentityList writes a list of values of the 'control_plane_operator_identity' type to
// the given writer.
func MarshalControlPlaneOperatorIdentityList(list []*ControlPlaneOperatorIdentity, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteControlPlaneOperatorIdentityList(list, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteControlPlaneOperatorIdentityList writes a list of value of the 'control_plane_operator_identity' type to
// the given stream.
func WriteControlPlaneOperatorIdentityList(list []*ControlPlaneOperatorIdentity, stream *jsoniter.Stream) {
	stream.WriteArrayStart()
	for i, value := range list {
		if i > 0 {
			stream.WriteMore()
		}
		WriteControlPlaneOperatorIdentity(value, stream)
	}
	stream.WriteArrayEnd()
}

// UnmarshalControlPlaneOperatorIdentityList reads a list of values of the 'control_plane_operator_identity' type
// from the given source, which can be a slice of bytes, a string or a reader.
func UnmarshalControlPlaneOperatorIdentityList(source interface{}) (items []*ControlPlaneOperatorIdentity, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	items = ReadControlPlaneOperatorIdentityList(iterator)
	err = iterator.Error
	return
}

// ReadControlPlaneOperatorIdentityList reads list of values of the ”control_plane_operator_identity' type from
// the given iterator.
func ReadControlPlaneOperatorIdentityList(iterator *jsoniter.Iterator) []*ControlPlaneOperatorIdentity {
	list := []*ControlPlaneOperatorIdentity{}
	for iterator.ReadArray() {
		item := ReadControlPlaneOperatorIdentity(iterator)
		list = append(list, item)
	}
	return list
}
