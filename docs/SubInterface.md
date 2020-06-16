# SubInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Index of the sub interface assigned by NSX Manager. Min value is 10 and max value is 4103. | [optional] [default to null]
**TunnelId** | **int32** | Valid values for tunnel ID are min 1 to max 4093. Required. | [default to null]
**Name** | **string** | Name of the sub interface. Required. | [optional] [default to null]
**AddressGroups** | [***EdgeVnicAddressGroups**](edgeVnicAddressGroups.md) |  | [optional] [default to null]
**VlanId** | **int32** | VLAN ID of the virtual LAN used by this sub interface. VLAN IDs can range from 0 to 4094. | [optional] [default to null]
**Label** | **string** | Sub interface label of format vNic_{index} provided by NSX Manager. Read only. | [optional] [default to null]
**LogicalSwitchName** | **string** | Name of the logical switch connected to this sub interface. | [optional] [default to null]
**IsConnected** | **bool** | Value is true if the sub interface is connected to a logical switch, standard portgroup or distributed portgroup. | [optional] [default to null]
**Mtu** | **int32** | MTU value of the sub interface. This value would be the least mtu for all the trunk interfaces of the NSX Edge. Default is 1500. | [optional] [default to null]
**LogicalSwitchId** | **string** | ID of the logical switch connected to this sub interface. | [optional] [default to null]
**EnableSendRedirects** | **bool** | Value is true if send redirects is enabled. Enable ICMP redirect to convey routing information to hosts. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

