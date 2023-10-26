// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkEndpointsInitParameters struct {
}

type NetworkEndpointsObservation struct {

	// (Output)
	// The IP address of this network endpoint.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// (Output)
	// The port of this network endpoint.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type NetworkEndpointsParameters struct {
}

type NodeInitParameters struct {

	// The type of hardware accelerators associated with this node.
	AcceleratorType *string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`

	// The CIDR block that the TPU node will use when selecting an IP
	// address. This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Resource labels to represent user provided metadata.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Sets the scheduling options for this TPU instance.
	// Structure is documented below.
	SchedulingConfig []SchedulingConfigInitParameters `json:"schedulingConfig,omitempty" tf:"scheduling_config,omitempty"`

	// The version of Tensorflow running in the Node.
	TensorflowVersion *string `json:"tensorflowVersion,omitempty" tf:"tensorflow_version,omitempty"`

	// Whether the VPC peering for the node is set up through Service Networking API.
	// The VPC Peering should be set up before provisioning the node. If this field is set,
	// cidr_block field should not be specified. If the network that you want to peer the
	// TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	UseServiceNetworking *bool `json:"useServiceNetworking,omitempty" tf:"use_service_networking,omitempty"`
}

type NodeObservation struct {

	// The type of hardware accelerators associated with this node.
	AcceleratorType *string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`

	// The CIDR block that the TPU node will use when selecting an IP
	// address. This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{zone}}/nodes/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource labels to represent user provided metadata.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of a network to peer the TPU node to. It must be a
	// preexisting Compute Engine network inside of the project on which
	// this API has been activated. If none is provided, "default" will be
	// used.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The network endpoints where TPU workers can be accessed and sent work.
	// It is recommended that Tensorflow clients of the node first reach out
	// to the first (index 0) entry.
	// Structure is documented below.
	NetworkEndpoints []NetworkEndpointsObservation `json:"networkEndpoints,omitempty" tf:"network_endpoints,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Sets the scheduling options for this TPU instance.
	// Structure is documented below.
	SchedulingConfig []SchedulingConfigObservation `json:"schedulingConfig,omitempty" tf:"scheduling_config,omitempty"`

	// The service account used to run the tensor flow services within the
	// node. To share resources, including Google Cloud Storage data, with
	// the Tensorflow job running in the Node, this account must have
	// permissions to that data.
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// The version of Tensorflow running in the Node.
	TensorflowVersion *string `json:"tensorflowVersion,omitempty" tf:"tensorflow_version,omitempty"`

	// Whether the VPC peering for the node is set up through Service Networking API.
	// The VPC Peering should be set up before provisioning the node. If this field is set,
	// cidr_block field should not be specified. If the network that you want to peer the
	// TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	UseServiceNetworking *bool `json:"useServiceNetworking,omitempty" tf:"use_service_networking,omitempty"`

	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NodeParameters struct {

	// The type of hardware accelerators associated with this node.
	// +kubebuilder:validation:Optional
	AcceleratorType *string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`

	// The CIDR block that the TPU node will use when selecting an IP
	// address. This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The user-supplied description of the TPU. Maximum of 512 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of a network to peer the TPU node to. It must be a
	// preexisting Compute Engine network inside of the project on which
	// this API has been activated. If none is provided, "default" will be
	// used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/servicenetworking/v1beta1.Connection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("network",false)
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Connection in servicenetworking to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Connection in servicenetworking to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Sets the scheduling options for this TPU instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SchedulingConfig []SchedulingConfigParameters `json:"schedulingConfig,omitempty" tf:"scheduling_config,omitempty"`

	// The version of Tensorflow running in the Node.
	// +kubebuilder:validation:Optional
	TensorflowVersion *string `json:"tensorflowVersion,omitempty" tf:"tensorflow_version,omitempty"`

	// Whether the VPC peering for the node is set up through Service Networking API.
	// The VPC Peering should be set up before provisioning the node. If this field is set,
	// cidr_block field should not be specified. If the network that you want to peer the
	// TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	// +kubebuilder:validation:Optional
	UseServiceNetworking *bool `json:"useServiceNetworking,omitempty" tf:"use_service_networking,omitempty"`

	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type SchedulingConfigInitParameters struct {

	// Defines whether the TPU instance is preemptible.
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
}

type SchedulingConfigObservation struct {

	// Defines whether the TPU instance is preemptible.
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
}

type SchedulingConfigParameters struct {

	// Defines whether the TPU instance is preemptible.
	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible" tf:"preemptible,omitempty"`
}

// NodeSpec defines the desired state of Node
type NodeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NodeInitParameters `json:"initProvider,omitempty"`
}

// NodeStatus defines the observed state of Node.
type NodeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Node is the Schema for the Nodes API. A Cloud TPU instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Node struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.acceleratorType) || (has(self.initProvider) && has(self.initProvider.acceleratorType))",message="spec.forProvider.acceleratorType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tensorflowVersion) || (has(self.initProvider) && has(self.initProvider.tensorflowVersion))",message="spec.forProvider.tensorflowVersion is a required parameter"
	Spec   NodeSpec   `json:"spec"`
	Status NodeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeList contains a list of Nodes
type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Node `json:"items"`
}

// Repository type metadata.
var (
	Node_Kind             = "Node"
	Node_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Node_Kind}.String()
	Node_KindAPIVersion   = Node_Kind + "." + CRDGroupVersion.String()
	Node_GroupVersionKind = CRDGroupVersion.WithKind(Node_Kind)
)

func init() {
	SchemeBuilder.Register(&Node{}, &NodeList{})
}
