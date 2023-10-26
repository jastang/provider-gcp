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

type ClientInitParameters struct {

	// The Azure Active Directory Application ID.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Azure Active Directory Tenant ID.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ClientObservation struct {

	// The Azure Active Directory Application ID.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Output only. The PEM encoded x509 certificate.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Output only. The time at which this resource was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/azureClients/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Azure Active Directory Tenant ID.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Output only. A globally unique identifier for the client.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type ClientParameters struct {

	// The Azure Active Directory Application ID.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Azure Active Directory Tenant ID.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// ClientSpec defines the desired state of Client
type ClientSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClientParameters `json:"forProvider"`
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
	InitProvider ClientInitParameters `json:"initProvider,omitempty"`
}

// ClientStatus defines the observed state of Client.
type ClientStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClientObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Client is the Schema for the Clients API. AzureClient resources hold client authentication information needed by the Anthos Multi-Cloud API to manage Azure resources on your Azure subscription.When an AzureCluster is created, an AzureClient resource needs to be provided and all operations on Azure resources associated to that cluster will authenticate to Azure services using the given client.AzureClient resources are immutable and cannot be modified upon creation.Each AzureClient resource is bound to a single Azure Active Directory Application and tenant.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Client struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationId) || (has(self.initProvider) && has(self.initProvider.applicationId))",message="spec.forProvider.applicationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || (has(self.initProvider) && has(self.initProvider.tenantId))",message="spec.forProvider.tenantId is a required parameter"
	Spec   ClientSpec   `json:"spec"`
	Status ClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientList contains a list of Clients
type ClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Client `json:"items"`
}

// Repository type metadata.
var (
	Client_Kind             = "Client"
	Client_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Client_Kind}.String()
	Client_KindAPIVersion   = Client_Kind + "." + CRDGroupVersion.String()
	Client_GroupVersionKind = CRDGroupVersion.WithKind(Client_Kind)
)

func init() {
	SchemeBuilder.Register(&Client{}, &ClientList{})
}
