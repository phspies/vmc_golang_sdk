# AddressFwSourceDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | **bool** | Exclude the specified source or destination. | [optional] [default to null]
**IpAddress** | **[]string** | List of string. Can specify single IP address, range of IP address, or in CIDR format. Can define multiple. | [optional] [default to null]
**GroupingObjectId** | **[]string** | List of string. Id of cluster, datacenter, distributedPortGroup, legacyPortGroup, VirtualMachine, vApp, resourcePool, logicalSwitch, IPSet, securityGroup. Can define multiple. | [optional] [default to null]
**VnicGroupId** | **[]string** | List of string. Possible values are vnic-index-[1-9], vse, external or internal. Can define multiple. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

