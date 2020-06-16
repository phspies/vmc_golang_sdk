# AwsSddcConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneNodeReducedCapacity** | **bool** | Whether this sddc is reduced capacity 1NODE. | [optional] [default to null]
**VpcCidr** | **string** | AWS VPC IP range. Only prefix of 16 or 20 is currently supported. | [optional] [default to 10.0.0.0/16]
**HostInstanceType** | **string** |  | [optional] [default to null]
**SkipCreatingVxlan** | **bool** | skip creating vxlan for compute gateway for SDDC provisioning | [optional] [default to null]
**VxlanSubnet** | **string** | VXLAN IP subnet in CIDR for compute gateway | [optional] [default to null]
**Size** | **string** | The size of the vCenter and NSX appliances. \&quot;large\&quot; sddcSize corresponds to a &#x27;large&#x27; vCenter appliance and &#x27;large&#x27; NSX appliance. &#x27;medium&#x27; sddcSize corresponds to &#x27;medium&#x27; vCenter appliance and &#x27;medium&#x27; NSX appliance. Value defaults to &#x27;medium&#x27;.  | [optional] [default to null]
**StorageCapacity** | **int64** | The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of seperable storage will be used.  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**AccountLinkSddcConfig** | [**[]AccountLinkSddcConfig**](AccountLinkSddcConfig.md) | A list of the SDDC linking configurations to use. | [optional] [default to null]
**SddcId** | **string** | If provided, will be assigned as SDDC id of the provisioned SDDC. | [optional] [default to null]
**NumHosts** | **int32** |  | [default to null]
**SddcType** | **string** | Denotes the sddc type , if the value is null or empty, the type is considered as default. | [optional] [default to null]
**AccountLinkConfig** | [***AccountLinkConfig**](AccountLinkConfig.md) |  | [optional] [default to null]
**Provider** | **string** | Determines what additional properties are available based on cloud provider. | [default to null]
**SsoDomain** | **string** | The SSO domain name to use for vSphere users. If not specified, vmc.local will be used. | [optional] [default to vmc.local]
**SddcTemplateId** | **string** | If provided, configuration from the template will applied to the provisioned SDDC. | [optional] [default to null]
**DeploymentType** | **string** | Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ. | [optional] [default to null]
**Region** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

