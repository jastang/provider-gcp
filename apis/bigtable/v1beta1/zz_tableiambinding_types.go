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

type TableIAMBindingConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TableIAMBindingConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TableIAMBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type TableIAMBindingInitParameters struct {
	Condition []TableIAMBindingConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type TableIAMBindingObservation struct {
	Condition []TableIAMBindingConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Table *string `json:"table,omitempty" tf:"table,omitempty"`
}

type TableIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []TableIAMBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// +kubebuilder:validation:Optional
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	Table *string `json:"table" tf:"table,omitempty"`
}

// TableIAMBindingSpec defines the desired state of TableIAMBinding
type TableIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableIAMBindingParameters `json:"forProvider"`
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
	InitProvider TableIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// TableIAMBindingStatus defines the observed state of TableIAMBinding.
type TableIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMBinding is the Schema for the TableIAMBindings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TableIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instance) || (has(self.initProvider) && has(self.initProvider.instance))",message="spec.forProvider.instance is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
	Spec   TableIAMBindingSpec   `json:"spec"`
	Status TableIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMBindingList contains a list of TableIAMBindings
type TableIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	TableIAMBinding_Kind             = "TableIAMBinding"
	TableIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableIAMBinding_Kind}.String()
	TableIAMBinding_KindAPIVersion   = TableIAMBinding_Kind + "." + CRDGroupVersion.String()
	TableIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(TableIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&TableIAMBinding{}, &TableIAMBindingList{})
}
