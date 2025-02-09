// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Namespaces_Topic_Spec. Use v1beta20210101preview.Namespaces_Topic_Spec instead
type Namespaces_Topic_Spec_ARM struct {
	Location   *string                `json:"location,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Properties *SBTopicProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string      `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespaces_Topic_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (topic Namespaces_Topic_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (topic *Namespaces_Topic_Spec_ARM) GetName() string {
	return topic.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics"
func (topic *Namespaces_Topic_Spec_ARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics"
}

// Deprecated version of SBTopicProperties. Use v1beta20210101preview.SBTopicProperties instead
type SBTopicProperties_ARM struct {
	AutoDeleteOnIdle                    *string `json:"autoDeleteOnIdle,omitempty"`
	DefaultMessageTimeToLive            *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool   `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool   `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool   `json:"enablePartitioning,omitempty"`
	MaxSizeInMegabytes                  *int    `json:"maxSizeInMegabytes,omitempty"`
	RequiresDuplicateDetection          *bool   `json:"requiresDuplicateDetection,omitempty"`
	SupportOrdering                     *bool   `json:"supportOrdering,omitempty"`
}
