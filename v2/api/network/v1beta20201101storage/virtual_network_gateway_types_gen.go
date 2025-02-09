// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=virtualnetworkgateways,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={virtualnetworkgateways/status,virtualnetworkgateways/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.VirtualNetworkGateway
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworkGateways
type VirtualNetworkGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGateway_Spec   `json:"spec,omitempty"`
	Status            VirtualNetworkGateway_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetworkGateway{}

// GetConditions returns the conditions of the resource
func (gateway *VirtualNetworkGateway) GetConditions() conditions.Conditions {
	return gateway.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (gateway *VirtualNetworkGateway) SetConditions(conditions conditions.Conditions) {
	gateway.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualNetworkGateway{}

// AzureName returns the Azure name of the resource
func (gateway *VirtualNetworkGateway) AzureName() string {
	return gateway.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (gateway VirtualNetworkGateway) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (gateway *VirtualNetworkGateway) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (gateway *VirtualNetworkGateway) GetSpec() genruntime.ConvertibleSpec {
	return &gateway.Spec
}

// GetStatus returns the status of this resource
func (gateway *VirtualNetworkGateway) GetStatus() genruntime.ConvertibleStatus {
	return &gateway.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworkGateways"
func (gateway *VirtualNetworkGateway) GetType() string {
	return "Microsoft.Network/virtualNetworkGateways"
}

// NewEmptyStatus returns a new empty (blank) status
func (gateway *VirtualNetworkGateway) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &VirtualNetworkGateway_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (gateway *VirtualNetworkGateway) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(gateway.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  gateway.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (gateway *VirtualNetworkGateway) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*VirtualNetworkGateway_STATUS); ok {
		gateway.Status = *st
		return nil
	}

	// Convert status to required version
	var st VirtualNetworkGateway_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	gateway.Status = st
	return nil
}

// Hub marks that this VirtualNetworkGateway is the hub type for conversion
func (gateway *VirtualNetworkGateway) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (gateway *VirtualNetworkGateway) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: gateway.Spec.OriginalVersion,
		Kind:    "VirtualNetworkGateway",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.VirtualNetworkGateway
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworkGateways
type VirtualNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkGateway `json:"items"`
}

// Storage version of v1beta20201101.VirtualNetworkGateway_Spec
type VirtualNetworkGateway_Spec struct {
	ActiveActive *bool `json:"activeActive,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName              string                                                   `json:"azureName,omitempty"`
	BgpSettings            *BgpSettings                                             `json:"bgpSettings,omitempty"`
	CustomRoutes           *AddressSpace                                            `json:"customRoutes,omitempty"`
	EnableBgp              *bool                                                    `json:"enableBgp,omitempty"`
	EnableDnsForwarding    *bool                                                    `json:"enableDnsForwarding,omitempty"`
	EnablePrivateIpAddress *bool                                                    `json:"enablePrivateIpAddress,omitempty"`
	GatewayDefaultSite     *SubResource                                             `json:"gatewayDefaultSite,omitempty"`
	GatewayType            *string                                                  `json:"gatewayType,omitempty"`
	IpConfigurations       []VirtualNetworkGateway_Properties_IpConfigurations_Spec `json:"ipConfigurations,omitempty"`
	Location               *string                                                  `json:"location,omitempty"`
	OriginalVersion        string                                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *VirtualNetworkGatewaySku          `json:"sku,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`

	// VNetExtendedLocationResourceReference: MAS FIJI customer vnet resource id. VirtualNetworkGateway of type local gateway
	// is associated with the customer vnet.
	VNetExtendedLocationResourceReference *genruntime.ResourceReference                                 `armReference:"VNetExtendedLocationResourceId" json:"vNetExtendedLocationResourceReference,omitempty"`
	VirtualNetworkExtendedLocation        *ExtendedLocation                                             `json:"virtualNetworkExtendedLocation,omitempty"`
	VpnClientConfiguration                *VirtualNetworkGateway_Properties_VpnClientConfiguration_Spec `json:"vpnClientConfiguration,omitempty"`
	VpnGatewayGeneration                  *string                                                       `json:"vpnGatewayGeneration,omitempty"`
	VpnType                               *string                                                       `json:"vpnType,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetworkGateway_Spec{}

// ConvertSpecFrom populates our VirtualNetworkGateway_Spec from the provided source
func (gateway *VirtualNetworkGateway_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == gateway {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(gateway)
}

// ConvertSpecTo populates the provided destination from our VirtualNetworkGateway_Spec
func (gateway *VirtualNetworkGateway_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == gateway {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(gateway)
}

// Storage version of v1beta20201101.VirtualNetworkGateway_STATUS
type VirtualNetworkGateway_STATUS struct {
	ActiveActive                   *bool                                         `json:"activeActive,omitempty"`
	BgpSettings                    *BgpSettings_STATUS                           `json:"bgpSettings,omitempty"`
	Conditions                     []conditions.Condition                        `json:"conditions,omitempty"`
	CustomRoutes                   *AddressSpace_STATUS                          `json:"customRoutes,omitempty"`
	EnableBgp                      *bool                                         `json:"enableBgp,omitempty"`
	EnableDnsForwarding            *bool                                         `json:"enableDnsForwarding,omitempty"`
	EnablePrivateIpAddress         *bool                                         `json:"enablePrivateIpAddress,omitempty"`
	Etag                           *string                                       `json:"etag,omitempty"`
	ExtendedLocation               *ExtendedLocation_STATUS                      `json:"extendedLocation,omitempty"`
	GatewayDefaultSite             *SubResource_STATUS                           `json:"gatewayDefaultSite,omitempty"`
	GatewayType                    *string                                       `json:"gatewayType,omitempty"`
	Id                             *string                                       `json:"id,omitempty"`
	InboundDnsForwardingEndpoint   *string                                       `json:"inboundDnsForwardingEndpoint,omitempty"`
	IpConfigurations               []VirtualNetworkGatewayIPConfiguration_STATUS `json:"ipConfigurations,omitempty"`
	Location                       *string                                       `json:"location,omitempty"`
	Name                           *string                                       `json:"name,omitempty"`
	PropertyBag                    genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	ProvisioningState              *string                                       `json:"provisioningState,omitempty"`
	ResourceGuid                   *string                                       `json:"resourceGuid,omitempty"`
	Sku                            *VirtualNetworkGatewaySku_STATUS              `json:"sku,omitempty"`
	Tags                           map[string]string                             `json:"tags,omitempty"`
	Type                           *string                                       `json:"type,omitempty"`
	VNetExtendedLocationResourceId *string                                       `json:"vNetExtendedLocationResourceId,omitempty"`
	VpnClientConfiguration         *VpnClientConfiguration_STATUS                `json:"vpnClientConfiguration,omitempty"`
	VpnGatewayGeneration           *string                                       `json:"vpnGatewayGeneration,omitempty"`
	VpnType                        *string                                       `json:"vpnType,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualNetworkGateway_STATUS{}

// ConvertStatusFrom populates our VirtualNetworkGateway_STATUS from the provided source
func (gateway *VirtualNetworkGateway_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == gateway {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(gateway)
}

// ConvertStatusTo populates the provided destination from our VirtualNetworkGateway_STATUS
func (gateway *VirtualNetworkGateway_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == gateway {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(gateway)
}

// Storage version of v1beta20201101.BgpSettings
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/BgpSettings
type BgpSettings struct {
	Asn                 *uint32                            `json:"asn,omitempty"`
	BgpPeeringAddress   *string                            `json:"bgpPeeringAddress,omitempty"`
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddress `json:"bgpPeeringAddresses,omitempty"`
	PeerWeight          *int                               `json:"peerWeight,omitempty"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.BgpSettings_STATUS
type BgpSettings_STATUS struct {
	Asn                 *uint32                                   `json:"asn,omitempty"`
	BgpPeeringAddress   *string                                   `json:"bgpPeeringAddress,omitempty"`
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddress_STATUS `json:"bgpPeeringAddresses,omitempty"`
	PeerWeight          *int                                      `json:"peerWeight,omitempty"`
	PropertyBag         genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.VirtualNetworkGateway_Properties_IpConfigurations_Spec
type VirtualNetworkGateway_Properties_IpConfigurations_Spec struct {
	Name                      *string                `json:"name,omitempty"`
	PrivateIPAllocationMethod *string                `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicIPAddress           *SubResource           `json:"publicIPAddress,omitempty"`
	Subnet                    *SubResource           `json:"subnet,omitempty"`
}

// Storage version of v1beta20201101.VirtualNetworkGateway_Properties_VpnClientConfiguration_Spec
type VirtualNetworkGateway_Properties_VpnClientConfiguration_Spec struct {
	AadAudience                  *string                                                                                     `json:"aadAudience,omitempty"`
	AadIssuer                    *string                                                                                     `json:"aadIssuer,omitempty"`
	AadTenant                    *string                                                                                     `json:"aadTenant,omitempty"`
	PropertyBag                  genruntime.PropertyBag                                                                      `json:"$propertyBag,omitempty"`
	RadiusServerAddress          *string                                                                                     `json:"radiusServerAddress,omitempty"`
	RadiusServerSecret           *string                                                                                     `json:"radiusServerSecret,omitempty"`
	RadiusServers                []RadiusServer                                                                              `json:"radiusServers,omitempty"`
	VpnAuthenticationTypes       []string                                                                                    `json:"vpnAuthenticationTypes,omitempty"`
	VpnClientAddressPool         *AddressSpace                                                                               `json:"vpnClientAddressPool,omitempty"`
	VpnClientIpsecPolicies       []IpsecPolicy                                                                               `json:"vpnClientIpsecPolicies,omitempty"`
	VpnClientProtocols           []string                                                                                    `json:"vpnClientProtocols,omitempty"`
	VpnClientRevokedCertificates []VirtualNetworkGateway_Properties_VpnClientConfiguration_VpnClientRevokedCertificates_Spec `json:"vpnClientRevokedCertificates,omitempty"`
	VpnClientRootCertificates    []VirtualNetworkGateway_Properties_VpnClientConfiguration_VpnClientRootCertificates_Spec    `json:"vpnClientRootCertificates,omitempty"`
}

// Storage version of v1beta20201101.VirtualNetworkGatewayIPConfiguration_STATUS
type VirtualNetworkGatewayIPConfiguration_STATUS struct {
	Etag                      *string                `json:"etag,omitempty"`
	Id                        *string                `json:"id,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PrivateIPAddress          *string                `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *string                `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                `json:"provisioningState,omitempty"`
	PublicIPAddress           *SubResource_STATUS    `json:"publicIPAddress,omitempty"`
	Subnet                    *SubResource_STATUS    `json:"subnet,omitempty"`
}

// Storage version of v1beta20201101.VirtualNetworkGatewaySku
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkGatewaySku
type VirtualNetworkGatewaySku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20201101.VirtualNetworkGatewaySku_STATUS
type VirtualNetworkGatewaySku_STATUS struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20201101.VpnClientConfiguration_STATUS
type VpnClientConfiguration_STATUS struct {
	AadAudience                  *string                              `json:"aadAudience,omitempty"`
	AadIssuer                    *string                              `json:"aadIssuer,omitempty"`
	AadTenant                    *string                              `json:"aadTenant,omitempty"`
	PropertyBag                  genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	RadiusServerAddress          *string                              `json:"radiusServerAddress,omitempty"`
	RadiusServerSecret           *string                              `json:"radiusServerSecret,omitempty"`
	RadiusServers                []RadiusServer_STATUS                `json:"radiusServers,omitempty"`
	VpnAuthenticationTypes       []string                             `json:"vpnAuthenticationTypes,omitempty"`
	VpnClientAddressPool         *AddressSpace_STATUS                 `json:"vpnClientAddressPool,omitempty"`
	VpnClientIpsecPolicies       []IpsecPolicy_STATUS                 `json:"vpnClientIpsecPolicies,omitempty"`
	VpnClientProtocols           []string                             `json:"vpnClientProtocols,omitempty"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificate_STATUS `json:"vpnClientRevokedCertificates,omitempty"`
	VpnClientRootCertificates    []VpnClientRootCertificate_STATUS    `json:"vpnClientRootCertificates,omitempty"`
}

// Storage version of v1beta20201101.IPConfigurationBgpPeeringAddress
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IPConfigurationBgpPeeringAddress
type IPConfigurationBgpPeeringAddress struct {
	CustomBgpIpAddresses []string               `json:"customBgpIpAddresses,omitempty"`
	IpconfigurationId    *string                `json:"ipconfigurationId,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.IPConfigurationBgpPeeringAddress_STATUS
type IPConfigurationBgpPeeringAddress_STATUS struct {
	CustomBgpIpAddresses  []string               `json:"customBgpIpAddresses,omitempty"`
	DefaultBgpIpAddresses []string               `json:"defaultBgpIpAddresses,omitempty"`
	IpconfigurationId     *string                `json:"ipconfigurationId,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TunnelIpAddresses     []string               `json:"tunnelIpAddresses,omitempty"`
}

// Storage version of v1beta20201101.IpsecPolicy
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpsecPolicy
type IpsecPolicy struct {
	DhGroup             *string                `json:"dhGroup,omitempty"`
	IkeEncryption       *string                `json:"ikeEncryption,omitempty"`
	IkeIntegrity        *string                `json:"ikeIntegrity,omitempty"`
	IpsecEncryption     *string                `json:"ipsecEncryption,omitempty"`
	IpsecIntegrity      *string                `json:"ipsecIntegrity,omitempty"`
	PfsGroup            *string                `json:"pfsGroup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SaDataSizeKilobytes *int                   `json:"saDataSizeKilobytes,omitempty"`
	SaLifeTimeSeconds   *int                   `json:"saLifeTimeSeconds,omitempty"`
}

// Storage version of v1beta20201101.IpsecPolicy_STATUS
type IpsecPolicy_STATUS struct {
	DhGroup             *string                `json:"dhGroup,omitempty"`
	IkeEncryption       *string                `json:"ikeEncryption,omitempty"`
	IkeIntegrity        *string                `json:"ikeIntegrity,omitempty"`
	IpsecEncryption     *string                `json:"ipsecEncryption,omitempty"`
	IpsecIntegrity      *string                `json:"ipsecIntegrity,omitempty"`
	PfsGroup            *string                `json:"pfsGroup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SaDataSizeKilobytes *int                   `json:"saDataSizeKilobytes,omitempty"`
	SaLifeTimeSeconds   *int                   `json:"saLifeTimeSeconds,omitempty"`
}

// Storage version of v1beta20201101.RadiusServer
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/RadiusServer
type RadiusServer struct {
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RadiusServerAddress *string                `json:"radiusServerAddress,omitempty"`
	RadiusServerScore   *int                   `json:"radiusServerScore,omitempty"`
	RadiusServerSecret  *string                `json:"radiusServerSecret,omitempty"`
}

// Storage version of v1beta20201101.RadiusServer_STATUS
type RadiusServer_STATUS struct {
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RadiusServerAddress *string                `json:"radiusServerAddress,omitempty"`
	RadiusServerScore   *int                   `json:"radiusServerScore,omitempty"`
	RadiusServerSecret  *string                `json:"radiusServerSecret,omitempty"`
}

// Storage version of v1beta20201101.VirtualNetworkGateway_Properties_VpnClientConfiguration_VpnClientRevokedCertificates_Spec
type VirtualNetworkGateway_Properties_VpnClientConfiguration_VpnClientRevokedCertificates_Spec struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Thumbprint  *string                `json:"thumbprint,omitempty"`
}

// Storage version of v1beta20201101.VirtualNetworkGateway_Properties_VpnClientConfiguration_VpnClientRootCertificates_Spec
type VirtualNetworkGateway_Properties_VpnClientConfiguration_VpnClientRootCertificates_Spec struct {
	Name           *string                `json:"name,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicCertData *string                `json:"publicCertData,omitempty"`
}

// Storage version of v1beta20201101.VpnClientRevokedCertificate_STATUS
type VpnClientRevokedCertificate_STATUS struct {
	Etag              *string                `json:"etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Thumbprint        *string                `json:"thumbprint,omitempty"`
}

// Storage version of v1beta20201101.VpnClientRootCertificate_STATUS
type VpnClientRootCertificate_STATUS struct {
	Etag              *string                `json:"etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	PublicCertData    *string                `json:"publicCertData,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VirtualNetworkGateway{}, &VirtualNetworkGatewayList{})
}
