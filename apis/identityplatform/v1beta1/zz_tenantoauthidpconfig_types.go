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

type TenantOAuthIdPConfigInitParameters struct {

	// Human friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// For OIDC Idps, the issuer identifier.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The name of the OauthIdpConfig. Must start with oidc..
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type TenantOAuthIdPConfigObservation struct {

	// Human friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// an identifier for the resource with format projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// For OIDC Idps, the issuer identifier.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The name of the OauthIdpConfig. Must start with oidc..
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the tenant where this OIDC IDP configuration resource exists
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`
}

type TenantOAuthIdPConfigParameters struct {

	// The client id of an OAuth client.
	// +kubebuilder:validation:Optional
	ClientIDSecretRef v1.SecretKeySelector `json:"clientIdSecretRef" tf:"-"`

	// The client secret of the OAuth client, to enable OIDC code flow.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef *v1.SecretKeySelector `json:"clientSecretSecretRef,omitempty" tf:"-"`

	// Human friendly display name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// For OIDC Idps, the issuer identifier.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The name of the OauthIdpConfig. Must start with oidc..
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the tenant where this OIDC IDP configuration resource exists
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/identityplatform/v1beta1.Tenant
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// Reference to a Tenant in identityplatform to populate tenant.
	// +kubebuilder:validation:Optional
	TenantRef *v1.Reference `json:"tenantRef,omitempty" tf:"-"`

	// Selector for a Tenant in identityplatform to populate tenant.
	// +kubebuilder:validation:Optional
	TenantSelector *v1.Selector `json:"tenantSelector,omitempty" tf:"-"`
}

// TenantOAuthIdPConfigSpec defines the desired state of TenantOAuthIdPConfig
type TenantOAuthIdPConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TenantOAuthIdPConfigParameters `json:"forProvider"`
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
	InitProvider TenantOAuthIdPConfigInitParameters `json:"initProvider,omitempty"`
}

// TenantOAuthIdPConfigStatus defines the observed state of TenantOAuthIdPConfig.
type TenantOAuthIdPConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TenantOAuthIdPConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TenantOAuthIdPConfig is the Schema for the TenantOAuthIdPConfigs API. OIDC IdP configuration for a Identity Toolkit project within a tenant.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TenantOAuthIdPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientIdSecretRef)",message="spec.forProvider.clientIdSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuer) || (has(self.initProvider) && has(self.initProvider.issuer))",message="spec.forProvider.issuer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   TenantOAuthIdPConfigSpec   `json:"spec"`
	Status TenantOAuthIdPConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantOAuthIdPConfigList contains a list of TenantOAuthIdPConfigs
type TenantOAuthIdPConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TenantOAuthIdPConfig `json:"items"`
}

// Repository type metadata.
var (
	TenantOAuthIdPConfig_Kind             = "TenantOAuthIdPConfig"
	TenantOAuthIdPConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TenantOAuthIdPConfig_Kind}.String()
	TenantOAuthIdPConfig_KindAPIVersion   = TenantOAuthIdPConfig_Kind + "." + CRDGroupVersion.String()
	TenantOAuthIdPConfig_GroupVersionKind = CRDGroupVersion.WithKind(TenantOAuthIdPConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&TenantOAuthIdPConfig{}, &TenantOAuthIdPConfigList{})
}
