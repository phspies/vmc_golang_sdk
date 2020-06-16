# Sddc

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
**Name** | **string** | name for SDDC | [optional] [default to null]
**SddcState** | **string** |  | [optional] [default to null]
**ExpirationDate** | [**time.Time**](time.Time.md) | Expiration date of a sddc in UTC (will be set if its applicable) | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**SddcType** | **string** | Type of the sddc | [optional] [default to null]
**OneNodeReducedCapacity** | **bool** | Whether this sddc is reduced capacity 1NODE | [optional] [default to null]
**Provider** | **string** |  | [optional] [default to null]
**AccountLinkState** | **string** | Account linking state of the sddc | [optional] [default to null]
**SddcAccessState** | **string** | Describes the access state of sddc, valid state is DISABLED or ENABLED | [optional] [default to null]
**ResourceConfig** | [***AwsSddcResourceConfig**](AwsSddcResourceConfig.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

