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

type AdvancedSettingsInitParameters struct {

	// If present, incoming audio is exported by Dialogflow to the configured Google Cloud Storage destination. Exposed at the following levels:
	AudioExportGcsDestination []AudioExportGcsDestinationInitParameters `json:"audioExportGcsDestination,omitempty" tf:"audio_export_gcs_destination,omitempty"`

	// Define behaviors for DTMF (dual tone multi frequency). DTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel. Exposed at the following levels:
	DtmfSettings []DtmfSettingsInitParameters `json:"dtmfSettings,omitempty" tf:"dtmf_settings,omitempty"`
}

type AdvancedSettingsObservation struct {

	// If present, incoming audio is exported by Dialogflow to the configured Google Cloud Storage destination. Exposed at the following levels:
	AudioExportGcsDestination []AudioExportGcsDestinationObservation `json:"audioExportGcsDestination,omitempty" tf:"audio_export_gcs_destination,omitempty"`

	// Define behaviors for DTMF (dual tone multi frequency). DTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel. Exposed at the following levels:
	DtmfSettings []DtmfSettingsObservation `json:"dtmfSettings,omitempty" tf:"dtmf_settings,omitempty"`
}

type AdvancedSettingsParameters struct {

	// If present, incoming audio is exported by Dialogflow to the configured Google Cloud Storage destination. Exposed at the following levels:
	// +kubebuilder:validation:Optional
	AudioExportGcsDestination []AudioExportGcsDestinationParameters `json:"audioExportGcsDestination,omitempty" tf:"audio_export_gcs_destination,omitempty"`

	// Define behaviors for DTMF (dual tone multi frequency). DTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel. Exposed at the following levels:
	// +kubebuilder:validation:Optional
	DtmfSettings []DtmfSettingsParameters `json:"dtmfSettings,omitempty" tf:"dtmf_settings,omitempty"`
}

type AgentInitParameters struct {

	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.
	// Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	// Structure is documented below.
	AdvancedSettings []AdvancedSettingsInitParameters `json:"advancedSettings,omitempty" tf:"advanced_settings,omitempty"`

	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	AvatarURI *string `json:"avatarUri,omitempty" tf:"avatar_uri,omitempty"`

	// The default language of the agent as a language tag. See Language Support
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty" tf:"default_language_code,omitempty"`

	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the agent, unique within the location.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty" tf:"enable_spell_correction,omitempty"`

	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging,omitempty"`

	// Git integration settings for this agent.
	// Structure is documented below.
	GitIntegrationSettings []GitIntegrationSettingsInitParameters `json:"gitIntegrationSettings,omitempty" tf:"git_integration_settings,omitempty"`

	// The name of the location this agent is located in.
	// ~> Note: The first time you are deploying an Agent in your project you must configure location settings.
	// This is a one time step but at the moment you can only configure location settings via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the SecuritySettings reference for the agent. Format: projects//locations//securitySettings/.
	SecuritySettings *string `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Settings related to speech recognition.
	// Structure is documented below.
	SpeechToTextSettings []SpeechToTextSettingsInitParameters `json:"speechToTextSettings,omitempty" tf:"speech_to_text_settings,omitempty"`

	// The list of all languages supported by this agent (except for the default_language_code).
	SupportedLanguageCodes []*string `json:"supportedLanguageCodes,omitempty" tf:"supported_language_codes,omitempty"`

	// Settings related to speech synthesizing.
	// Structure is documented below.
	TextToSpeechSettings []TextToSpeechSettingsInitParameters `json:"textToSpeechSettings,omitempty" tf:"text_to_speech_settings,omitempty"`

	// The time zone of this agent from the time zone database, e.g., America/New_York,
	// Europe/Paris.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type AgentObservation struct {

	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.
	// Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	// Structure is documented below.
	AdvancedSettings []AdvancedSettingsObservation `json:"advancedSettings,omitempty" tf:"advanced_settings,omitempty"`

	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	AvatarURI *string `json:"avatarUri,omitempty" tf:"avatar_uri,omitempty"`

	// The default language of the agent as a language tag. See Language Support
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty" tf:"default_language_code,omitempty"`

	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the agent, unique within the location.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty" tf:"enable_spell_correction,omitempty"`

	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging,omitempty"`

	// Git integration settings for this agent.
	// Structure is documented below.
	GitIntegrationSettings []GitIntegrationSettingsObservation `json:"gitIntegrationSettings,omitempty" tf:"git_integration_settings,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/agents/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the location this agent is located in.
	// ~> Note: The first time you are deploying an Agent in your project you must configure location settings.
	// This is a one time step but at the moment you can only configure location settings via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The unique identifier of the agent.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the SecuritySettings reference for the agent. Format: projects//locations//securitySettings/.
	SecuritySettings *string `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Settings related to speech recognition.
	// Structure is documented below.
	SpeechToTextSettings []SpeechToTextSettingsObservation `json:"speechToTextSettings,omitempty" tf:"speech_to_text_settings,omitempty"`

	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects//locations//agents//flows/.
	StartFlow *string `json:"startFlow,omitempty" tf:"start_flow,omitempty"`

	// The list of all languages supported by this agent (except for the default_language_code).
	SupportedLanguageCodes []*string `json:"supportedLanguageCodes,omitempty" tf:"supported_language_codes,omitempty"`

	// Settings related to speech synthesizing.
	// Structure is documented below.
	TextToSpeechSettings []TextToSpeechSettingsObservation `json:"textToSpeechSettings,omitempty" tf:"text_to_speech_settings,omitempty"`

	// The time zone of this agent from the time zone database, e.g., America/New_York,
	// Europe/Paris.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type AgentParameters struct {

	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.
	// Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AdvancedSettings []AdvancedSettingsParameters `json:"advancedSettings,omitempty" tf:"advanced_settings,omitempty"`

	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	// +kubebuilder:validation:Optional
	AvatarURI *string `json:"avatarUri,omitempty" tf:"avatar_uri,omitempty"`

	// The default language of the agent as a language tag. See Language Support
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	// +kubebuilder:validation:Optional
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty" tf:"default_language_code,omitempty"`

	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the agent, unique within the location.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	// +kubebuilder:validation:Optional
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty" tf:"enable_spell_correction,omitempty"`

	// Determines whether this agent should log conversation queries.
	// +kubebuilder:validation:Optional
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging,omitempty"`

	// Git integration settings for this agent.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GitIntegrationSettings []GitIntegrationSettingsParameters `json:"gitIntegrationSettings,omitempty" tf:"git_integration_settings,omitempty"`

	// The name of the location this agent is located in.
	// ~> Note: The first time you are deploying an Agent in your project you must configure location settings.
	// This is a one time step but at the moment you can only configure location settings via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the SecuritySettings reference for the agent. Format: projects//locations//securitySettings/.
	// +kubebuilder:validation:Optional
	SecuritySettings *string `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Settings related to speech recognition.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SpeechToTextSettings []SpeechToTextSettingsParameters `json:"speechToTextSettings,omitempty" tf:"speech_to_text_settings,omitempty"`

	// The list of all languages supported by this agent (except for the default_language_code).
	// +kubebuilder:validation:Optional
	SupportedLanguageCodes []*string `json:"supportedLanguageCodes,omitempty" tf:"supported_language_codes,omitempty"`

	// Settings related to speech synthesizing.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TextToSpeechSettings []TextToSpeechSettingsParameters `json:"textToSpeechSettings,omitempty" tf:"text_to_speech_settings,omitempty"`

	// The time zone of this agent from the time zone database, e.g., America/New_York,
	// Europe/Paris.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type AudioExportGcsDestinationInitParameters struct {

	// The Google Cloud Storage URI for the exported objects. Whether a full object name, or just a prefix, its usage depends on the Dialogflow operation.
	// Format: gs://bucket/object-name-or-prefix
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type AudioExportGcsDestinationObservation struct {

	// The Google Cloud Storage URI for the exported objects. Whether a full object name, or just a prefix, its usage depends on the Dialogflow operation.
	// Format: gs://bucket/object-name-or-prefix
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type AudioExportGcsDestinationParameters struct {

	// The Google Cloud Storage URI for the exported objects. Whether a full object name, or just a prefix, its usage depends on the Dialogflow operation.
	// Format: gs://bucket/object-name-or-prefix
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type DtmfSettingsInitParameters struct {

	// If true, incoming audio is processed for DTMF (dual tone multi frequency) events. For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will detect the event (e.g. a "3" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The digit that terminates a DTMF digit sequence.
	FinishDigit *string `json:"finishDigit,omitempty" tf:"finish_digit,omitempty"`

	// Max length of DTMF digits.
	MaxDigits *float64 `json:"maxDigits,omitempty" tf:"max_digits,omitempty"`
}

type DtmfSettingsObservation struct {

	// If true, incoming audio is processed for DTMF (dual tone multi frequency) events. For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will detect the event (e.g. a "3" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The digit that terminates a DTMF digit sequence.
	FinishDigit *string `json:"finishDigit,omitempty" tf:"finish_digit,omitempty"`

	// Max length of DTMF digits.
	MaxDigits *float64 `json:"maxDigits,omitempty" tf:"max_digits,omitempty"`
}

type DtmfSettingsParameters struct {

	// If true, incoming audio is processed for DTMF (dual tone multi frequency) events. For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will detect the event (e.g. a "3" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The digit that terminates a DTMF digit sequence.
	// +kubebuilder:validation:Optional
	FinishDigit *string `json:"finishDigit,omitempty" tf:"finish_digit,omitempty"`

	// Max length of DTMF digits.
	// +kubebuilder:validation:Optional
	MaxDigits *float64 `json:"maxDigits,omitempty" tf:"max_digits,omitempty"`
}

type GitIntegrationSettingsInitParameters struct {

	// Settings of integration with GitHub.
	// Structure is documented below.
	GithubSettings []GithubSettingsInitParameters `json:"githubSettings,omitempty" tf:"github_settings,omitempty"`
}

type GitIntegrationSettingsObservation struct {

	// Settings of integration with GitHub.
	// Structure is documented below.
	GithubSettings []GithubSettingsObservation `json:"githubSettings,omitempty" tf:"github_settings,omitempty"`
}

type GitIntegrationSettingsParameters struct {

	// Settings of integration with GitHub.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GithubSettings []GithubSettingsParameters `json:"githubSettings,omitempty" tf:"github_settings,omitempty"`
}

type GithubSettingsInitParameters struct {

	// The access token used to authenticate the access to the GitHub repository.
	// Note: This property is sensitive and will not be displayed in the plan.
	AccessTokenSecretRef *v1.SecretKeySelector `json:"accessTokenSecretRef,omitempty" tf:"-"`

	// A list of branches configured to be used from Dialogflow.
	Branches []*string `json:"branches,omitempty" tf:"branches,omitempty"`

	// The unique repository display name for the GitHub repository.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The GitHub repository URI related to the agent.
	RepositoryURI *string `json:"repositoryUri,omitempty" tf:"repository_uri,omitempty"`

	// The branch of the GitHub repository tracked for this agent.
	TrackingBranch *string `json:"trackingBranch,omitempty" tf:"tracking_branch,omitempty"`
}

type GithubSettingsObservation struct {

	// A list of branches configured to be used from Dialogflow.
	Branches []*string `json:"branches,omitempty" tf:"branches,omitempty"`

	// The unique repository display name for the GitHub repository.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The GitHub repository URI related to the agent.
	RepositoryURI *string `json:"repositoryUri,omitempty" tf:"repository_uri,omitempty"`

	// The branch of the GitHub repository tracked for this agent.
	TrackingBranch *string `json:"trackingBranch,omitempty" tf:"tracking_branch,omitempty"`
}

type GithubSettingsParameters struct {

	// The access token used to authenticate the access to the GitHub repository.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	AccessTokenSecretRef *v1.SecretKeySelector `json:"accessTokenSecretRef,omitempty" tf:"-"`

	// A list of branches configured to be used from Dialogflow.
	// +kubebuilder:validation:Optional
	Branches []*string `json:"branches,omitempty" tf:"branches,omitempty"`

	// The unique repository display name for the GitHub repository.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The GitHub repository URI related to the agent.
	// +kubebuilder:validation:Optional
	RepositoryURI *string `json:"repositoryUri,omitempty" tf:"repository_uri,omitempty"`

	// The branch of the GitHub repository tracked for this agent.
	// +kubebuilder:validation:Optional
	TrackingBranch *string `json:"trackingBranch,omitempty" tf:"tracking_branch,omitempty"`
}

type SpeechToTextSettingsInitParameters struct {

	// Whether to use speech adaptation for speech recognition.
	EnableSpeechAdaptation *bool `json:"enableSpeechAdaptation,omitempty" tf:"enable_speech_adaptation,omitempty"`
}

type SpeechToTextSettingsObservation struct {

	// Whether to use speech adaptation for speech recognition.
	EnableSpeechAdaptation *bool `json:"enableSpeechAdaptation,omitempty" tf:"enable_speech_adaptation,omitempty"`
}

type SpeechToTextSettingsParameters struct {

	// Whether to use speech adaptation for speech recognition.
	// +kubebuilder:validation:Optional
	EnableSpeechAdaptation *bool `json:"enableSpeechAdaptation,omitempty" tf:"enable_speech_adaptation,omitempty"`
}

type TextToSpeechSettingsInitParameters struct {

	// Configuration of how speech should be synthesized, mapping from language to SynthesizeSpeechConfig.
	// These settings affect:
	SynthesizeSpeechConfigs *string `json:"synthesizeSpeechConfigs,omitempty" tf:"synthesize_speech_configs,omitempty"`
}

type TextToSpeechSettingsObservation struct {

	// Configuration of how speech should be synthesized, mapping from language to SynthesizeSpeechConfig.
	// These settings affect:
	SynthesizeSpeechConfigs *string `json:"synthesizeSpeechConfigs,omitempty" tf:"synthesize_speech_configs,omitempty"`
}

type TextToSpeechSettingsParameters struct {

	// Configuration of how speech should be synthesized, mapping from language to SynthesizeSpeechConfig.
	// These settings affect:
	// +kubebuilder:validation:Optional
	SynthesizeSpeechConfigs *string `json:"synthesizeSpeechConfigs,omitempty" tf:"synthesize_speech_configs,omitempty"`
}

// AgentSpec defines the desired state of Agent
type AgentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgentParameters `json:"forProvider"`
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
	InitProvider AgentInitParameters `json:"initProvider,omitempty"`
}

// AgentStatus defines the observed state of Agent.
type AgentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Agent is the Schema for the Agents API. Agents are best described as Natural Language Understanding (NLU) modules that transform user requests into actionable data.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Agent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultLanguageCode) || (has(self.initProvider) && has(self.initProvider.defaultLanguageCode))",message="spec.forProvider.defaultLanguageCode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timeZone) || (has(self.initProvider) && has(self.initProvider.timeZone))",message="spec.forProvider.timeZone is a required parameter"
	Spec   AgentSpec   `json:"spec"`
	Status AgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentList contains a list of Agents
type AgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Agent `json:"items"`
}

// Repository type metadata.
var (
	Agent_Kind             = "Agent"
	Agent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Agent_Kind}.String()
	Agent_KindAPIVersion   = Agent_Kind + "." + CRDGroupVersion.String()
	Agent_GroupVersionKind = CRDGroupVersion.WithKind(Agent_Kind)
)

func init() {
	SchemeBuilder.Register(&Agent{}, &AgentList{})
}
