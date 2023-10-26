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

type NetworkEndpointInitParameters struct {

	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Port number of network endpoint.
	// Note port is required unless the Network Endpoint Group is created
	// with the type of GCE_VM_IP
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Zone where the containing network endpoint group is located.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NetworkEndpointObservation struct {

	// an identifier for the resource with format {{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name for a specific VM instance that the IP address belongs to.
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup *string `json:"networkEndpointGroup,omitempty" tf:"network_endpoint_group,omitempty"`

	// Port number of network endpoint.
	// Note port is required unless the Network Endpoint Group is created
	// with the type of GCE_VM_IP
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Zone where the containing network endpoint group is located.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NetworkEndpointParameters struct {

	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name for a specific VM instance that the IP address belongs to.
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// The network endpoint group this endpoint is part of.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.NetworkEndpointGroup
	// +kubebuilder:validation:Optional
	NetworkEndpointGroup *string `json:"networkEndpointGroup,omitempty" tf:"network_endpoint_group,omitempty"`

	// Reference to a NetworkEndpointGroup in compute to populate networkEndpointGroup.
	// +kubebuilder:validation:Optional
	NetworkEndpointGroupRef *v1.Reference `json:"networkEndpointGroupRef,omitempty" tf:"-"`

	// Selector for a NetworkEndpointGroup in compute to populate networkEndpointGroup.
	// +kubebuilder:validation:Optional
	NetworkEndpointGroupSelector *v1.Selector `json:"networkEndpointGroupSelector,omitempty" tf:"-"`

	// Port number of network endpoint.
	// Note port is required unless the Network Endpoint Group is created
	// with the type of GCE_VM_IP
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Zone where the containing network endpoint group is located.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// NetworkEndpointSpec defines the desired state of NetworkEndpoint
type NetworkEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkEndpointParameters `json:"forProvider"`
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
	InitProvider NetworkEndpointInitParameters `json:"initProvider,omitempty"`
}

// NetworkEndpointStatus defines the observed state of NetworkEndpoint.
type NetworkEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkEndpoint is the Schema for the NetworkEndpoints API. A Network endpoint represents a IP address and port combination that is part of a specific network endpoint group (NEG).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NetworkEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipAddress) || (has(self.initProvider) && has(self.initProvider.ipAddress))",message="spec.forProvider.ipAddress is a required parameter"
	Spec   NetworkEndpointSpec   `json:"spec"`
	Status NetworkEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkEndpointList contains a list of NetworkEndpoints
type NetworkEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkEndpoint `json:"items"`
}

// Repository type metadata.
var (
	NetworkEndpoint_Kind             = "NetworkEndpoint"
	NetworkEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkEndpoint_Kind}.String()
	NetworkEndpoint_KindAPIVersion   = NetworkEndpoint_Kind + "." + CRDGroupVersion.String()
	NetworkEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(NetworkEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkEndpoint{}, &NetworkEndpointList{})
}
