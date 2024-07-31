// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BigqueryOptionsInitParameters struct {

	// Whether to use BigQuery's partition tables.
	// By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
	// tables, the date suffix is no longer present and special query syntax
	// has to be used instead. In both cases, tables are sharded based on UTC timezone.
	UsePartitionedTables *bool `json:"usePartitionedTables,omitempty" tf:"use_partitioned_tables,omitempty"`
}

type BigqueryOptionsObservation struct {

	// Whether to use BigQuery's partition tables.
	// By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
	// tables, the date suffix is no longer present and special query syntax
	// has to be used instead. In both cases, tables are sharded based on UTC timezone.
	UsePartitionedTables *bool `json:"usePartitionedTables,omitempty" tf:"use_partitioned_tables,omitempty"`
}

type BigqueryOptionsParameters struct {

	// Whether to use BigQuery's partition tables.
	// By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
	// tables, the date suffix is no longer present and special query syntax
	// has to be used instead. In both cases, tables are sharded based on UTC timezone.
	// +kubebuilder:validation:Optional
	UsePartitionedTables *bool `json:"usePartitionedTables" tf:"use_partitioned_tables,omitempty"`
}

type ExclusionsInitParameters struct {

	// A description of this exclusion.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// A client-assigned identifier, such as load-balancer-exclusion. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ExclusionsObservation struct {

	// A description of this exclusion.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// A client-assigned identifier, such as load-balancer-exclusion. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ExclusionsParameters struct {

	// A description of this exclusion.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See Advanced Log Filters for information on how to
	// write a filter.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// A client-assigned identifier, such as load-balancer-exclusion. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type FolderSinkInitParameters struct {

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions []BigqueryOptionsInitParameters `json:"bigqueryOptions,omitempty" tf:"bigquery_options,omitempty"`

	// A description of this sink. The maximum length of the description is 8000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, a BigQuery dataset, a Cloud Logging bucket, or a Google Cloud project. Examples:
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusions.filter, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []ExclusionsInitParameters `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren *bool `json:"includeChildren,omitempty" tf:"include_children,omitempty"`

	// Whether or not to intercept logs from child projects. If true, matching logs will not
	// match with sinks in child resources, except _Required sinks. This sink will be visible to child resources when listing sinks.
	InterceptChildren *bool `json:"interceptChildren,omitempty" tf:"intercept_children,omitempty"`
}

type FolderSinkObservation struct {

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions []BigqueryOptionsObservation `json:"bigqueryOptions,omitempty" tf:"bigquery_options,omitempty"`

	// A description of this sink. The maximum length of the description is 8000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, a BigQuery dataset, a Cloud Logging bucket, or a Google Cloud project. Examples:
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusions.filter, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []ExclusionsObservation `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The folder to be exported to the sink. Note that either [FOLDER_ID] or folders/[FOLDER_ID] is
	// accepted.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// an identifier for the resource with format folders/{{folder_id}}/sinks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren *bool `json:"includeChildren,omitempty" tf:"include_children,omitempty"`

	// Whether or not to intercept logs from child projects. If true, matching logs will not
	// match with sinks in child resources, except _Required sinks. This sink will be visible to child resources when listing sinks.
	InterceptChildren *bool `json:"interceptChildren,omitempty" tf:"intercept_children,omitempty"`

	// The identity associated with this sink. This identity must be granted write access to the
	// configured destination.
	WriterIdentity *string `json:"writerIdentity,omitempty" tf:"writer_identity,omitempty"`
}

type FolderSinkParameters struct {

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	// +kubebuilder:validation:Optional
	BigqueryOptions []BigqueryOptionsParameters `json:"bigqueryOptions,omitempty" tf:"bigquery_options,omitempty"`

	// A description of this sink. The maximum length of the description is 8000 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, a BigQuery dataset, a Cloud Logging bucket, or a Google Cloud project. Examples:
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// If set to True, then this sink is disabled and it does not export any log entries.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusions.filter, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	// +kubebuilder:validation:Optional
	Exclusions []ExclusionsParameters `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See Advanced Log Filters for information on how to
	// write a filter.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The folder to be exported to the sink. Note that either [FOLDER_ID] or folders/[FOLDER_ID] is
	// accepted.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Folder
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// Reference to a Folder in cloudplatform to populate folder.
	// +kubebuilder:validation:Optional
	FolderRef *v1.Reference `json:"folderRef,omitempty" tf:"-"`

	// Selector for a Folder in cloudplatform to populate folder.
	// +kubebuilder:validation:Optional
	FolderSelector *v1.Selector `json:"folderSelector,omitempty" tf:"-"`

	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	// +kubebuilder:validation:Optional
	IncludeChildren *bool `json:"includeChildren,omitempty" tf:"include_children,omitempty"`

	// Whether or not to intercept logs from child projects. If true, matching logs will not
	// match with sinks in child resources, except _Required sinks. This sink will be visible to child resources when listing sinks.
	// +kubebuilder:validation:Optional
	InterceptChildren *bool `json:"interceptChildren,omitempty" tf:"intercept_children,omitempty"`
}

// FolderSinkSpec defines the desired state of FolderSink
type FolderSinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderSinkParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FolderSinkInitParameters `json:"initProvider,omitempty"`
}

// FolderSinkStatus defines the observed state of FolderSink.
type FolderSinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderSinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FolderSink is the Schema for the FolderSinks API. Manages a folder-level logging sink.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type FolderSink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || (has(self.initProvider) && has(self.initProvider.destination))",message="spec.forProvider.destination is a required parameter"
	Spec   FolderSinkSpec   `json:"spec"`
	Status FolderSinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderSinkList contains a list of FolderSinks
type FolderSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderSink `json:"items"`
}

// Repository type metadata.
var (
	FolderSink_Kind             = "FolderSink"
	FolderSink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderSink_Kind}.String()
	FolderSink_KindAPIVersion   = FolderSink_Kind + "." + CRDGroupVersion.String()
	FolderSink_GroupVersionKind = CRDGroupVersion.WithKind(FolderSink_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderSink{}, &FolderSinkList{})
}
