# Vnic

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubInterfaces** | [***SubInterfaces**](SubInterfaces.md) |  | [optional] [default to null]
**AddressGroups** | [***EdgeVnicAddressGroups**](edgeVnicAddressGroups.md) |  | [optional] [default to null]
**IsConnected** | **bool** | Value is true if the vnic is connected to a logical switch, standard portgroup or distributed portgroup. | [optional] [default to null]
**EnableSendRedirects** | **bool** | Value is true if send redirects is enabled. Enable ICMP redirect to convey routing information to hosts. | [optional] [default to null]
**InShapingPolicy** | [***TrafficShapingPolicy**](TrafficShapingPolicy.md) |  | [optional] [default to null]
**Label** | **string** | Interface label of format vNic_{vnicIndex} provided by NSX Manager. Read only. | [optional] [default to null]
**EnableProxyArp** | **bool** | Value is true if proxy arp is enabled. Enable proxy ARP if you want to allow the NSX Edge of type &#x27;gatewayServices&#x27; to answer ARP requests intended for other machines. | [optional] [default to null]
**Index** | **int32** | Index of the vnic. Min value is 0 and max value is 9. | [default to null]
**Name** | **string** | Name of the interface. Optional. | [optional] [default to null]
**Mtu** | **int32** | MTU of the interface, with default as 1500. Min is 68, Max is 9000. Optional. | [optional] [default to null]
**FenceParameters** | [**[]KeyValueAttributes**](keyValueAttributes.md) |  | [optional] [default to null]
**MacAddresses** | [**[]MacAddress**](MacAddress.md) | Distinct MAC addresses configured for the vnic. Optional. | [optional] [default to null]
**OutShapingPolicy** | [***TrafficShapingPolicy**](TrafficShapingPolicy.md) |  | [optional] [default to null]
**PortgroupName** | **string** | Name of the port group or logical switch. | [optional] [default to null]
**EnableBridgeMode** | **bool** | Value is true if bridge mode is enabled. | [optional] [default to null]
**Type_** | **string** | Type of the vnic. Values are uplink, internal, trunk. At least one internal interface must be configured for NSX Edge HA to work. | [optional] [default to null]
**PortgroupId** | **string** | Value are port group ID (standard portgroup or distributed portgroup) or virtual wire ID (logical switch). Logical switch cannot be used for a TRUNK vnic. Portgroup cannot be shared among vnics/LIFs. Required when isConnected is specified as true. Example &#x27;network-17&#x27; (standard portgroup), &#x27;dvportgroup-34&#x27; (distributed portgroup) or &#x27;virtualwire-2&#x27; (logical switch). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

