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

type HaVPNGatewayInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The stack type for this VPN gateway to identify the IP protocols that are enabled.
	// If not specified, IPV4_ONLY will be used.
	// Default value is IPV4_ONLY.
	// Possible values are: IPV4_ONLY, IPV4_IPV6.
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`

	// A list of interfaces on this VPN gateway.
	// Structure is documented below.
	VPNInterfaces []VPNInterfacesInitParameters `json:"vpnInterfaces,omitempty" tf:"vpn_interfaces,omitempty"`
}

type HaVPNGatewayObservation struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The network this VPN gateway is accepting traffic for.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region this gateway should sit in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The stack type for this VPN gateway to identify the IP protocols that are enabled.
	// If not specified, IPV4_ONLY will be used.
	// Default value is IPV4_ONLY.
	// Possible values are: IPV4_ONLY, IPV4_IPV6.
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`

	// A list of interfaces on this VPN gateway.
	// Structure is documented below.
	VPNInterfaces []VPNInterfacesObservation `json:"vpnInterfaces,omitempty" tf:"vpn_interfaces,omitempty"`
}

type HaVPNGatewayParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The network this VPN gateway is accepting traffic for.
	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region this gateway should sit in.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The stack type for this VPN gateway to identify the IP protocols that are enabled.
	// If not specified, IPV4_ONLY will be used.
	// Default value is IPV4_ONLY.
	// Possible values are: IPV4_ONLY, IPV4_IPV6.
	// +kubebuilder:validation:Optional
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`

	// A list of interfaces on this VPN gateway.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VPNInterfaces []VPNInterfacesParameters `json:"vpnInterfaces,omitempty" tf:"vpn_interfaces,omitempty"`
}

type VPNInterfacesInitParameters struct {

	// The numeric ID of this VPN gateway interface.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`
}

type VPNInterfacesObservation struct {

	// The numeric ID of this VPN gateway interface.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// (Output)
	// The external IP address for this VPN gateway interface.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// URL of the interconnect attachment resource. When the value
	// of this field is present, the VPN Gateway will be used for
	// IPsec-encrypted Cloud Interconnect; all Egress or Ingress
	// traffic for this VPN Gateway interface will go through the
	// specified interconnect attachment resource.
	// Not currently available publicly.
	InterconnectAttachment *string `json:"interconnectAttachment,omitempty" tf:"interconnect_attachment,omitempty"`
}

type VPNInterfacesParameters struct {

	// The numeric ID of this VPN gateway interface.
	// +kubebuilder:validation:Optional
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// URL of the interconnect attachment resource. When the value
	// of this field is present, the VPN Gateway will be used for
	// IPsec-encrypted Cloud Interconnect; all Egress or Ingress
	// traffic for this VPN Gateway interface will go through the
	// specified interconnect attachment resource.
	// Not currently available publicly.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.InterconnectAttachment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	InterconnectAttachment *string `json:"interconnectAttachment,omitempty" tf:"interconnect_attachment,omitempty"`

	// Reference to a InterconnectAttachment in compute to populate interconnectAttachment.
	// +kubebuilder:validation:Optional
	InterconnectAttachmentRef *v1.Reference `json:"interconnectAttachmentRef,omitempty" tf:"-"`

	// Selector for a InterconnectAttachment in compute to populate interconnectAttachment.
	// +kubebuilder:validation:Optional
	InterconnectAttachmentSelector *v1.Selector `json:"interconnectAttachmentSelector,omitempty" tf:"-"`
}

// HaVPNGatewaySpec defines the desired state of HaVPNGateway
type HaVPNGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HaVPNGatewayParameters `json:"forProvider"`
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
	InitProvider HaVPNGatewayInitParameters `json:"initProvider,omitempty"`
}

// HaVPNGatewayStatus defines the observed state of HaVPNGateway.
type HaVPNGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HaVPNGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HaVPNGateway is the Schema for the HaVPNGateways API. Represents a VPN gateway running in GCP.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type HaVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HaVPNGatewaySpec   `json:"spec"`
	Status            HaVPNGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HaVPNGatewayList contains a list of HaVPNGateways
type HaVPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HaVPNGateway `json:"items"`
}

// Repository type metadata.
var (
	HaVPNGateway_Kind             = "HaVPNGateway"
	HaVPNGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HaVPNGateway_Kind}.String()
	HaVPNGateway_KindAPIVersion   = HaVPNGateway_Kind + "." + CRDGroupVersion.String()
	HaVPNGateway_GroupVersionKind = CRDGroupVersion.WithKind(HaVPNGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&HaVPNGateway{}, &HaVPNGatewayList{})
}
