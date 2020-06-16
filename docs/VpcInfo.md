# VpcInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcCidr** | **string** |  | [optional] [default to null]
**VgwId** | **string** |  | [optional] [default to null]
**EsxPublicSecurityGroupId** | **string** |  | [optional] [default to null]
**VifIds** | **[]string** | set of virtual interfaces attached to the sddc | [optional] [default to null]
**VmSecurityGroupId** | **string** |  | [optional] [default to null]
**TgwIps** | [**map[string][]string**](array.md) | Mapping from AZ to a list of IP addresses assigned to TGW ENI that&#x27;s connected with Vpc | [optional] [default to null]
**RouteTableId** | **string** | (deprecated) | [optional] [default to null]
**EdgeSubnetId** | **string** | Id of the NSX edge associated with this VPC (deprecated) | [optional] [default to null]
**Id** | **string** | vpc id | [optional] [default to null]
**ApiAssociationId** | **string** | Id of the association between subnet and route-table (deprecated) | [optional] [default to null]
**ApiSubnetId** | **string** | Id associated with this VPC (deprecated) | [optional] [default to null]
**PrivateSubnetId** | **string** | (deprecated) | [optional] [default to null]
**PrivateAssociationId** | **string** | (deprecated) | [optional] [default to null]
**EsxSecurityGroupId** | **string** |  | [optional] [default to null]
**SubnetId** | **string** | (deprecated) | [optional] [default to null]
**InternetGatewayId** | **string** |  | [optional] [default to null]
**SecurityGroupId** | **string** |  | [optional] [default to null]
**AssociationId** | **string** | (deprecated) | [optional] [default to null]
**VgwRouteTableId** | **string** | Route table which contains the route to VGW (deprecated) | [optional] [default to null]
**EdgeAssociationId** | **string** | Id of the association between edge subnet and route-table (deprecated) | [optional] [default to null]
**Provider** | **string** |  | [optional] [default to null]
**PeeringConnectionId** | **string** | (deprecated) | [optional] [default to null]
**NetworkType** | **string** |  | [optional] [default to null]
**AvailableZones** | [**[]AvailableZoneInfo**](AvailableZoneInfo.md) |  | [optional] [default to null]
**Routetables** | [**map[string]RouteTableInfo**](RouteTableInfo.md) | map from routeTableName to routeTableInfo | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

