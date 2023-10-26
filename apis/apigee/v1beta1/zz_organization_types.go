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

type OrganizationInitParameters struct {

	// Primary GCP region for analytics data storage. For valid values, see Create an Apigee organization.
	AnalyticsRegion *string `json:"analyticsRegion,omitempty" tf:"analytics_region,omitempty"`

	// Billing type of the Apigee organization. See Apigee pricing.
	BillingType *string `json:"billingType,omitempty" tf:"billing_type,omitempty"`

	// Description of the Apigee organization.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Flag that specifies whether the VPC Peering through Private Google Access should be
	// disabled between the consumer network and Apigee. Required if an authorizedNetwork
	// on the consumer project is not provided, in which case the flag should be set to true.
	// Valid only when RuntimeType is set to CLOUD. The value must be set before the creation
	// of any Apigee runtime instance and can be updated only when there are no runtime instances.
	DisableVPCPeering *bool `json:"disableVpcPeering,omitempty" tf:"disable_vpc_peering,omitempty"`

	// The display name of the Apigee organization.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The project ID associated with the Apigee organization.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Properties defined in the Apigee organization profile.
	// Structure is documented below.
	Properties []PropertiesInitParameters `json:"properties,omitempty" tf:"properties,omitempty"`

	// Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	// operation completes. During this period, the Organization may be restored to its last known state.
	// After this period, the Organization will no longer be able to be restored.
	// Default value is DELETION_RETENTION_UNSPECIFIED.
	// Possible values are: DELETION_RETENTION_UNSPECIFIED, MINIMUM.
	Retention *string `json:"retention,omitempty" tf:"retention,omitempty"`

	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is CLOUD.
	// Possible values are: CLOUD, HYBRID.
	RuntimeType *string `json:"runtimeType,omitempty" tf:"runtime_type,omitempty"`
}

type OrganizationObservation struct {

	// Primary GCP region for analytics data storage. For valid values, see Create an Apigee organization.
	AnalyticsRegion *string `json:"analyticsRegion,omitempty" tf:"analytics_region,omitempty"`

	// Output only. Project ID of the Apigee Tenant Project.
	ApigeeProjectID *string `json:"apigeeProjectId,omitempty" tf:"apigee_project_id,omitempty"`

	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See Getting started with the Service Networking API.
	// Valid only when RuntimeType is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`

	// Billing type of the Apigee organization. See Apigee pricing.
	BillingType *string `json:"billingType,omitempty" tf:"billing_type,omitempty"`

	// Output only. Base64-encoded public certificate for the root CA of the Apigee organization.
	// Valid only when RuntimeType is CLOUD. A base64-encoded string.
	CACertificate *string `json:"caCertificate,omitempty" tf:"ca_certificate,omitempty"`

	// Description of the Apigee organization.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Flag that specifies whether the VPC Peering through Private Google Access should be
	// disabled between the consumer network and Apigee. Required if an authorizedNetwork
	// on the consumer project is not provided, in which case the flag should be set to true.
	// Valid only when RuntimeType is set to CLOUD. The value must be set before the creation
	// of any Apigee runtime instance and can be updated only when there are no runtime instances.
	DisableVPCPeering *bool `json:"disableVpcPeering,omitempty" tf:"disable_vpc_peering,omitempty"`

	// The display name of the Apigee organization.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format organizations/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Output only. Name of the Apigee organization.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project ID associated with the Apigee organization.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Properties defined in the Apigee organization profile.
	// Structure is documented below.
	Properties []PropertiesObservation `json:"properties,omitempty" tf:"properties,omitempty"`

	// Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	// operation completes. During this period, the Organization may be restored to its last known state.
	// After this period, the Organization will no longer be able to be restored.
	// Default value is DELETION_RETENTION_UNSPECIFIED.
	// Possible values are: DELETION_RETENTION_UNSPECIFIED, MINIMUM.
	Retention *string `json:"retention,omitempty" tf:"retention,omitempty"`

	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when RuntimeType is CLOUD. For example: projects/foo/locations/us/keyRings/bar/cryptoKeys/baz.
	RuntimeDatabaseEncryptionKeyName *string `json:"runtimeDatabaseEncryptionKeyName,omitempty" tf:"runtime_database_encryption_key_name,omitempty"`

	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is CLOUD.
	// Possible values are: CLOUD, HYBRID.
	RuntimeType *string `json:"runtimeType,omitempty" tf:"runtime_type,omitempty"`

	// Output only. Subscription type of the Apigee organization.
	// Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased).
	SubscriptionType *string `json:"subscriptionType,omitempty" tf:"subscription_type,omitempty"`
}

type OrganizationParameters struct {

	// Primary GCP region for analytics data storage. For valid values, see Create an Apigee organization.
	// +kubebuilder:validation:Optional
	AnalyticsRegion *string `json:"analyticsRegion,omitempty" tf:"analytics_region,omitempty"`

	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	// See Getting started with the Service Networking API.
	// Valid only when RuntimeType is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`

	// Reference to a Network in compute to populate authorizedNetwork.
	// +kubebuilder:validation:Optional
	AuthorizedNetworkRef *v1.Reference `json:"authorizedNetworkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate authorizedNetwork.
	// +kubebuilder:validation:Optional
	AuthorizedNetworkSelector *v1.Selector `json:"authorizedNetworkSelector,omitempty" tf:"-"`

	// Billing type of the Apigee organization. See Apigee pricing.
	// +kubebuilder:validation:Optional
	BillingType *string `json:"billingType,omitempty" tf:"billing_type,omitempty"`

	// Description of the Apigee organization.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Flag that specifies whether the VPC Peering through Private Google Access should be
	// disabled between the consumer network and Apigee. Required if an authorizedNetwork
	// on the consumer project is not provided, in which case the flag should be set to true.
	// Valid only when RuntimeType is set to CLOUD. The value must be set before the creation
	// of any Apigee runtime instance and can be updated only when there are no runtime instances.
	// +kubebuilder:validation:Optional
	DisableVPCPeering *bool `json:"disableVpcPeering,omitempty" tf:"disable_vpc_peering,omitempty"`

	// The display name of the Apigee organization.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The project ID associated with the Apigee organization.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Properties defined in the Apigee organization profile.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Properties []PropertiesParameters `json:"properties,omitempty" tf:"properties,omitempty"`

	// Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	// is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	// operation completes. During this period, the Organization may be restored to its last known state.
	// After this period, the Organization will no longer be able to be restored.
	// Default value is DELETION_RETENTION_UNSPECIFIED.
	// Possible values are: DELETION_RETENTION_UNSPECIFIED, MINIMUM.
	// +kubebuilder:validation:Optional
	Retention *string `json:"retention,omitempty" tf:"retention,omitempty"`

	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	// Update is not allowed after the organization is created.
	// If not specified, a Google-Managed encryption key will be used.
	// Valid only when RuntimeType is CLOUD. For example: projects/foo/locations/us/keyRings/bar/cryptoKeys/baz.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RuntimeDatabaseEncryptionKeyName *string `json:"runtimeDatabaseEncryptionKeyName,omitempty" tf:"runtime_database_encryption_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate runtimeDatabaseEncryptionKeyName.
	// +kubebuilder:validation:Optional
	RuntimeDatabaseEncryptionKeyNameRef *v1.Reference `json:"runtimeDatabaseEncryptionKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate runtimeDatabaseEncryptionKeyName.
	// +kubebuilder:validation:Optional
	RuntimeDatabaseEncryptionKeyNameSelector *v1.Selector `json:"runtimeDatabaseEncryptionKeyNameSelector,omitempty" tf:"-"`

	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	// Default value is CLOUD.
	// Possible values are: CLOUD, HYBRID.
	// +kubebuilder:validation:Optional
	RuntimeType *string `json:"runtimeType,omitempty" tf:"runtime_type,omitempty"`
}

type PropertiesInitParameters struct {

	// List of all properties in the object.
	// Structure is documented below.
	Property []PropertyInitParameters `json:"property,omitempty" tf:"property,omitempty"`
}

type PropertiesObservation struct {

	// List of all properties in the object.
	// Structure is documented below.
	Property []PropertyObservation `json:"property,omitempty" tf:"property,omitempty"`
}

type PropertiesParameters struct {

	// List of all properties in the object.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Property []PropertyParameters `json:"property,omitempty" tf:"property,omitempty"`
}

type PropertyInitParameters struct {

	// Name of the property.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Value of the property.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PropertyObservation struct {

	// Name of the property.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Value of the property.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PropertyParameters struct {

	// Name of the property.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Value of the property.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// OrganizationSpec defines the desired state of Organization
type OrganizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationParameters `json:"forProvider"`
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
	InitProvider OrganizationInitParameters `json:"initProvider,omitempty"`
}

// OrganizationStatus defines the observed state of Organization.
type OrganizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Organization is the Schema for the Organizations API. An
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	Spec   OrganizationSpec   `json:"spec"`
	Status OrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationList contains a list of Organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Organization `json:"items"`
}

// Repository type metadata.
var (
	Organization_Kind             = "Organization"
	Organization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Organization_Kind}.String()
	Organization_KindAPIVersion   = Organization_Kind + "." + CRDGroupVersion.String()
	Organization_GroupVersionKind = CRDGroupVersion.WithKind(Organization_Kind)
)

func init() {
	SchemeBuilder.Register(&Organization{}, &OrganizationList{})
}
