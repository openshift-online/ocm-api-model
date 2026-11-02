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

// AzureNodePoolIPForwarding represents the values of the 'azure_node_pool_IP_forwarding' enumerated type.
type AzureNodePoolIPForwarding string

const (
	// The network interfaces may send and receive traffic for any source or destination
	// address, including addresses that aren't assigned to the interface.
	AzureNodePoolIPForwardingAllowAllAddresses AzureNodePoolIPForwarding = "allow_all_addresses"
	// The network interfaces may only send and receive traffic for addresses assigned to
	// themselves. This is the Azure default.
	AzureNodePoolIPForwardingInterfaceAddressesOnly AzureNodePoolIPForwarding = "interface_addresses_only"
)
