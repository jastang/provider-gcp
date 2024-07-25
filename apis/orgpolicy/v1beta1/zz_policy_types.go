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

type ConditionInitParameters struct {

	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionObservation struct {

	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionParameters struct {

	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type DryRunSpecInitParameters struct {

	// Determines the inheritance behavior for this policy. If inherit_from_parent is true, policy rules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this policy becomes the new root for evaluation. This field can be set only for policies which configure list constraints.
	InheritFromParent *bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent,omitempty"`

	// Ignores policies set above this resource and restores the constraint_default enforcement behavior of the specific constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, rules must be empty and inherit_from_parent must be set to false.
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`

	// In policies for boolean constraints, the following requirements apply: - There must be one and only one policy rule where condition is unset. - Boolean policy rules with conditions must set enforced to the opposite of the policy rule without a condition. - During policy evaluation, policy rules with conditions that are true for a target resource take precedence.
	Rules []RulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type DryRunSpecObservation struct {

	// An opaque tag indicating the current version of the policy, used for concurrency control. This field is ignored if used in a CreatePolicy request. When the policyis returned from either aGetPolicyor aListPoliciesrequest, thisetagindicates the version of the current policy to use when executing a read-modify-write loop. When the policy is returned from aGetEffectivePolicyrequest, theetag` will be unset.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Determines the inheritance behavior for this policy. If inherit_from_parent is true, policy rules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this policy becomes the new root for evaluation. This field can be set only for policies which configure list constraints.
	InheritFromParent *bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent,omitempty"`

	// Ignores policies set above this resource and restores the constraint_default enforcement behavior of the specific constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, rules must be empty and inherit_from_parent must be set to false.
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`

	// In policies for boolean constraints, the following requirements apply: - There must be one and only one policy rule where condition is unset. - Boolean policy rules with conditions must set enforced to the opposite of the policy rule without a condition. - During policy evaluation, policy rules with conditions that are true for a target resource take precedence.
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// Output only. The time stamp this was previously updated. This represents the last time a call to CreatePolicy or UpdatePolicy was made for that policy.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type DryRunSpecParameters struct {

	// Determines the inheritance behavior for this policy. If inherit_from_parent is true, policy rules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this policy becomes the new root for evaluation. This field can be set only for policies which configure list constraints.
	// +kubebuilder:validation:Optional
	InheritFromParent *bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent,omitempty"`

	// Ignores policies set above this resource and restores the constraint_default enforcement behavior of the specific constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, rules must be empty and inherit_from_parent must be set to false.
	// +kubebuilder:validation:Optional
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`

	// In policies for boolean constraints, the following requirements apply: - There must be one and only one policy rule where condition is unset. - Boolean policy rules with conditions must set enforced to the opposite of the policy rule without a condition. - During policy evaluation, policy rules with conditions that are true for a target resource take precedence.
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type PolicyInitParameters struct {

	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced.
	DryRunSpec *DryRunSpecInitParameters `json:"dryRunSpec,omitempty" tf:"dry_run_spec,omitempty"`

	// Basic information about the Organization Policy.
	Spec *SpecInitParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type PolicyObservation struct {

	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced.
	DryRunSpec *DryRunSpecObservation `json:"dryRunSpec,omitempty" tf:"dry_run_spec,omitempty"`

	// Optional. An opaque tag indicating the current state of the policy, used for concurrency control. This 'etag' is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// an identifier for the resource with format {{parent}}/policies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The parent of the resource.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Basic information about the Organization Policy.
	Spec *SpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`
}

type PolicyParameters struct {

	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced.
	// +kubebuilder:validation:Optional
	DryRunSpec *DryRunSpecParameters `json:"dryRunSpec,omitempty" tf:"dry_run_spec,omitempty"`

	// The parent of the resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Folder
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Folder in cloudplatform to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Folder in cloudplatform to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// Basic information about the Organization Policy.
	// +kubebuilder:validation:Optional
	Spec *SpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type RulesConditionInitParameters struct {

	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type RulesConditionObservation struct {

	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type RulesConditionParameters struct {

	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type RulesInitParameters struct {

	// Setting this to "TRUE" means that all values are allowed. This field can be set only in policies for list constraints.
	AllowAll *string `json:"allowAll,omitempty" tf:"allow_all,omitempty"`

	// A condition which determines whether this rule is used in the evaluation of the policy. When set, the expression field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
	Condition *ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Setting this to "TRUE" means that all values are denied. This field can be set only in policies for list constraints.
	DenyAll *string `json:"denyAll,omitempty" tf:"deny_all,omitempty"`

	// If "TRUE", then the policy is enforced. If "FALSE", then any configuration is acceptable. This field can be set only in policies for boolean constraints.
	Enforce *string `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// List of values to be used for this policy rule. This field can be set only in policies for list constraints.
	Values *ValuesInitParameters `json:"values,omitempty" tf:"values,omitempty"`
}

type RulesObservation struct {

	// Setting this to "TRUE" means that all values are allowed. This field can be set only in policies for list constraints.
	AllowAll *string `json:"allowAll,omitempty" tf:"allow_all,omitempty"`

	// A condition which determines whether this rule is used in the evaluation of the policy. When set, the expression field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
	Condition *ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// Setting this to "TRUE" means that all values are denied. This field can be set only in policies for list constraints.
	DenyAll *string `json:"denyAll,omitempty" tf:"deny_all,omitempty"`

	// If "TRUE", then the policy is enforced. If "FALSE", then any configuration is acceptable. This field can be set only in policies for boolean constraints.
	Enforce *string `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// List of values to be used for this policy rule. This field can be set only in policies for list constraints.
	Values *ValuesObservation `json:"values,omitempty" tf:"values,omitempty"`
}

type RulesParameters struct {

	// Setting this to "TRUE" means that all values are allowed. This field can be set only in policies for list constraints.
	// +kubebuilder:validation:Optional
	AllowAll *string `json:"allowAll,omitempty" tf:"allow_all,omitempty"`

	// A condition which determines whether this rule is used in the evaluation of the policy. When set, the expression field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
	// +kubebuilder:validation:Optional
	Condition *ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Setting this to "TRUE" means that all values are denied. This field can be set only in policies for list constraints.
	// +kubebuilder:validation:Optional
	DenyAll *string `json:"denyAll,omitempty" tf:"deny_all,omitempty"`

	// If "TRUE", then the policy is enforced. If "FALSE", then any configuration is acceptable. This field can be set only in policies for boolean constraints.
	// +kubebuilder:validation:Optional
	Enforce *string `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// List of values to be used for this policy rule. This field can be set only in policies for list constraints.
	// +kubebuilder:validation:Optional
	Values *ValuesParameters `json:"values,omitempty" tf:"values,omitempty"`
}

type RulesValuesInitParameters struct {

	// List of values allowed at this resource.
	AllowedValues []*string `json:"allowedValues,omitempty" tf:"allowed_values,omitempty"`

	// List of values denied at this resource.
	DeniedValues []*string `json:"deniedValues,omitempty" tf:"denied_values,omitempty"`
}

type RulesValuesObservation struct {

	// List of values allowed at this resource.
	AllowedValues []*string `json:"allowedValues,omitempty" tf:"allowed_values,omitempty"`

	// List of values denied at this resource.
	DeniedValues []*string `json:"deniedValues,omitempty" tf:"denied_values,omitempty"`
}

type RulesValuesParameters struct {

	// List of values allowed at this resource.
	// +kubebuilder:validation:Optional
	AllowedValues []*string `json:"allowedValues,omitempty" tf:"allowed_values,omitempty"`

	// List of values denied at this resource.
	// +kubebuilder:validation:Optional
	DeniedValues []*string `json:"deniedValues,omitempty" tf:"denied_values,omitempty"`
}

type SpecInitParameters struct {

	// Determines the inheritance behavior for this Policy. If inherit_from_parent is true, PolicyRules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this Policy becomes the new root for evaluation. This field can be set only for Policies which configure list constraints.
	InheritFromParent *bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent,omitempty"`

	// Ignores policies set above this resource and restores the constraint_default enforcement behavior of the specific Constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, rules must be empty and inherit_from_parent must be set to false.
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`

	// Up to 10 PolicyRules are allowed. In Policies for boolean constraints, the following requirements apply: - There must be one and only one PolicyRule where condition is unset. - BooleanPolicyRules with conditions must set enforced to the opposite of the PolicyRule without a condition. - During policy evaluation, PolicyRules with conditions that are true for a target resource take precedence.
	Rules []SpecRulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type SpecObservation struct {

	// An opaque tag indicating the current version of the Policy, used for concurrency control. This field is ignored if used in a CreatePolicy request. When the Policy is returned from either a GetPolicy or a ListPolicies request, this etag indicates the version of the current Policy to use when executing a read-modify-write loop. When the Policy is returned from a GetEffectivePolicy request, the etag will be unset.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Determines the inheritance behavior for this Policy. If inherit_from_parent is true, PolicyRules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this Policy becomes the new root for evaluation. This field can be set only for Policies which configure list constraints.
	InheritFromParent *bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent,omitempty"`

	// Ignores policies set above this resource and restores the constraint_default enforcement behavior of the specific Constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, rules must be empty and inherit_from_parent must be set to false.
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`

	// Up to 10 PolicyRules are allowed. In Policies for boolean constraints, the following requirements apply: - There must be one and only one PolicyRule where condition is unset. - BooleanPolicyRules with conditions must set enforced to the opposite of the PolicyRule without a condition. - During policy evaluation, PolicyRules with conditions that are true for a target resource take precedence.
	Rules []SpecRulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// Output only. The time stamp this was previously updated. This represents the last time a call to CreatePolicy or UpdatePolicy was made for that Policy.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type SpecParameters struct {

	// Determines the inheritance behavior for this Policy. If inherit_from_parent is true, PolicyRules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this Policy becomes the new root for evaluation. This field can be set only for Policies which configure list constraints.
	// +kubebuilder:validation:Optional
	InheritFromParent *bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent,omitempty"`

	// Ignores policies set above this resource and restores the constraint_default enforcement behavior of the specific Constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, rules must be empty and inherit_from_parent must be set to false.
	// +kubebuilder:validation:Optional
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`

	// Up to 10 PolicyRules are allowed. In Policies for boolean constraints, the following requirements apply: - There must be one and only one PolicyRule where condition is unset. - BooleanPolicyRules with conditions must set enforced to the opposite of the PolicyRule without a condition. - During policy evaluation, PolicyRules with conditions that are true for a target resource take precedence.
	// +kubebuilder:validation:Optional
	Rules []SpecRulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type SpecRulesInitParameters struct {

	// Setting this to "TRUE" means that all values are allowed. This field can be set only in policies for list constraints.
	AllowAll *string `json:"allowAll,omitempty" tf:"allow_all,omitempty"`

	// A condition which determines whether this rule is used in the evaluation of the policy. When set, the expression field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
	Condition *RulesConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Setting this to "TRUE" means that all values are denied. This field can be set only in policies for list constraints.
	DenyAll *string `json:"denyAll,omitempty" tf:"deny_all,omitempty"`

	// If "TRUE", then the policy is enforced. If "FALSE", then any configuration is acceptable. This field can be set only in policies for boolean constraints.
	Enforce *string `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// List of values to be used for this policy rule. This field can be set only in policies for list constraints.
	Values *RulesValuesInitParameters `json:"values,omitempty" tf:"values,omitempty"`
}

type SpecRulesObservation struct {

	// Setting this to "TRUE" means that all values are allowed. This field can be set only in policies for list constraints.
	AllowAll *string `json:"allowAll,omitempty" tf:"allow_all,omitempty"`

	// A condition which determines whether this rule is used in the evaluation of the policy. When set, the expression field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
	Condition *RulesConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// Setting this to "TRUE" means that all values are denied. This field can be set only in policies for list constraints.
	DenyAll *string `json:"denyAll,omitempty" tf:"deny_all,omitempty"`

	// If "TRUE", then the policy is enforced. If "FALSE", then any configuration is acceptable. This field can be set only in policies for boolean constraints.
	Enforce *string `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// List of values to be used for this policy rule. This field can be set only in policies for list constraints.
	Values *RulesValuesObservation `json:"values,omitempty" tf:"values,omitempty"`
}

type SpecRulesParameters struct {

	// Setting this to "TRUE" means that all values are allowed. This field can be set only in policies for list constraints.
	// +kubebuilder:validation:Optional
	AllowAll *string `json:"allowAll,omitempty" tf:"allow_all,omitempty"`

	// A condition which determines whether this rule is used in the evaluation of the policy. When set, the expression field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
	// +kubebuilder:validation:Optional
	Condition *RulesConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Setting this to "TRUE" means that all values are denied. This field can be set only in policies for list constraints.
	// +kubebuilder:validation:Optional
	DenyAll *string `json:"denyAll,omitempty" tf:"deny_all,omitempty"`

	// If "TRUE", then the policy is enforced. If "FALSE", then any configuration is acceptable. This field can be set only in policies for boolean constraints.
	// +kubebuilder:validation:Optional
	Enforce *string `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// List of values to be used for this policy rule. This field can be set only in policies for list constraints.
	// +kubebuilder:validation:Optional
	Values *RulesValuesParameters `json:"values,omitempty" tf:"values,omitempty"`
}

type ValuesInitParameters struct {

	// List of values allowed at this resource.
	AllowedValues []*string `json:"allowedValues,omitempty" tf:"allowed_values,omitempty"`

	// List of values denied at this resource.
	DeniedValues []*string `json:"deniedValues,omitempty" tf:"denied_values,omitempty"`
}

type ValuesObservation struct {

	// List of values allowed at this resource.
	AllowedValues []*string `json:"allowedValues,omitempty" tf:"allowed_values,omitempty"`

	// List of values denied at this resource.
	DeniedValues []*string `json:"deniedValues,omitempty" tf:"denied_values,omitempty"`
}

type ValuesParameters struct {

	// List of values allowed at this resource.
	// +kubebuilder:validation:Optional
	AllowedValues []*string `json:"allowedValues,omitempty" tf:"allowed_values,omitempty"`

	// List of values denied at this resource.
	// +kubebuilder:validation:Optional
	DeniedValues []*string `json:"deniedValues,omitempty" tf:"denied_values,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
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
	InitProvider PolicyInitParameters `json:"initProvider,omitempty"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Policy is the Schema for the Policys API. An organization policy gives you programmatic control over your organization's cloud resources.  Using Organization Policies, you will be able to configure constraints across your entire resource hierarchy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
