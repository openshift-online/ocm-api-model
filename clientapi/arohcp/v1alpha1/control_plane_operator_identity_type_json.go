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

// MarshalControlPlaneOperatorIdentity writes a value of the 'control_plane_operator_identity' type to the given writer.
func MarshalControlPlaneOperatorIdentity(object *ControlPlaneOperatorIdentity, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteControlPlaneOperatorIdentity(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteControlPlaneOperatorIdentity writes a value of the 'control_plane_operator_identity' type to the given stream.
func WriteControlPlaneOperatorIdentity(object *ControlPlaneOperatorIdentity, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = object.bitmap_&1 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("max_openshift_version")
		stream.WriteString(object.maxOpenShiftVersion)
		count++
	}
	present_ = object.bitmap_&2 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("min_openshift_version")
		stream.WriteString(object.minOpenShiftVersion)
		count++
	}
	present_ = object.bitmap_&4 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("operator_name")
		stream.WriteString(object.operatorName)
		count++
	}
	present_ = object.bitmap_&8 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("required")
		stream.WriteString(object.required)
		count++
	}
	present_ = object.bitmap_&16 != 0 && object.roleDefinitions != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("role_definitions")
		WriteRoleDefinitionList(object.roleDefinitions, stream)
	}
	stream.WriteObjectEnd()
}

// UnmarshalControlPlaneOperatorIdentity reads a value of the 'control_plane_operator_identity' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalControlPlaneOperatorIdentity(source interface{}) (object *ControlPlaneOperatorIdentity, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadControlPlaneOperatorIdentity(iterator)
	err = iterator.Error
	return
}

// ReadControlPlaneOperatorIdentity reads a value of the 'control_plane_operator_identity' type from the given iterator.
func ReadControlPlaneOperatorIdentity(iterator *jsoniter.Iterator) *ControlPlaneOperatorIdentity {
	object := &ControlPlaneOperatorIdentity{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "max_openshift_version":
			value := iterator.ReadString()
			object.maxOpenShiftVersion = value
			object.bitmap_ |= 1
		case "min_openshift_version":
			value := iterator.ReadString()
			object.minOpenShiftVersion = value
			object.bitmap_ |= 2
		case "operator_name":
			value := iterator.ReadString()
			object.operatorName = value
			object.bitmap_ |= 4
		case "required":
			value := iterator.ReadString()
			object.required = value
			object.bitmap_ |= 8
		case "role_definitions":
			value := ReadRoleDefinitionList(iterator)
			object.roleDefinitions = value
			object.bitmap_ |= 16
		default:
			iterator.ReadAny()
		}
	}
	return object
}
