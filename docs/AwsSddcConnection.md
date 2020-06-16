# AwsSddcConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | [**time.Time**](time.Time.md) |  | [default to null]
**UserId** | **string** | User id that last updated this record | [default to null]
**UpdatedByUserId** | **string** | User id that last updated this record | [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**Version** | **int32** | Version of this entity | [default to null]
**UpdatedByUserName** | **string** | User name that last updated this record | [optional] [default to null]
**UserName** | **string** | User name that last updated this record | [default to null]
**Id** | **string** | Unique ID for this entity | [default to null]
**CidrBlockSubnet** | **string** | The CIDR block of the customer&#x27;s subnet this link is in. | [optional] [default to null]
**ConnectedAccountId** | **string** | The corresponding connected (customer) account UUID this connection is attached to. | [optional] [default to null]
**EniGroup** | **string** | Which group the ENIs belongs to. (deprecated) | [optional] [default to null]
**SubnetId** | **string** | The ID of the subnet this link is to. | [optional] [default to null]
**CgwPresent** | **bool** | Determines whether the CGW is present in this connection set or not. Used for multi-az deployments. | [optional] [default to null]
**OrgId** | **string** | The org this link belongs to. | [optional] [default to null]
**SddcId** | **string** | The SDDC this link is used for. | [optional] [default to null]
**CidrBlockVpc** | **string** | The CIDR block of the customer&#x27;s VPC. | [optional] [default to null]
**ConnectionOrder** | **int32** | The order of the connection | [optional] [default to null]
**State** | **string** | The state of the connection. | [optional] [default to null]
**SubnetAvailabilityZone** | **string** | Which availability zone is this connection in? | [optional] [default to null]
**VpcId** | **string** | The VPC ID of the subnet this link is to. | [optional] [default to null]
**CustomerEniInfos** | [**[]CustomerEniInfo**](CustomerEniInfo.md) | A list of all ENIs used for this connection. | [optional] [default to null]
**DefaultRouteTable** | **string** | The default routing table in the customer&#x27;s VPC. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

