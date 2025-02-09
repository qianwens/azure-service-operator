// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	alpha20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101"
	alpha20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101storage"
	v20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type VirtualNetworkExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *VirtualNetworkExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&alpha20201101.VirtualNetwork{},
		&alpha20201101s.VirtualNetwork{},
		&v20201101.VirtualNetwork{},
		&v20201101s.VirtualNetwork{}}
}
