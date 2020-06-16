# EdgeStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreRulesExists** | **bool** | Value is true if pre rules publish is enabled. | [optional] [default to null]
**FeatureStatuses** | [**[]FeatureStatus**](featureStatus.md) | Individual feature status. | [optional] [default to null]
**Timestamp** | **int64** | Timestamp value at which the NSX Edge healthcheck was done. | [optional] [default to null]
**PublishStatus** | **string** | Status of the latest configuration change for the NSX Edge. Values are APPLIED or PERSISTED (not published to NSX Edge Appliance yet). | [optional] [default to null]
**LastPublishedPreRulesGenerationNumber** | **int64** | Value of the last published pre rules generation number. | [optional] [default to null]
**Version** | **int64** | Version number of the current configuration. | [optional] [default to null]
**EdgeVmStatus** | [**[]EdgeVmStatus**](edgeVmStatus.md) | Detailed status of each of the deployed NSX Edge appliances. | [optional] [default to null]
**ActiveVseHaIndex** | **int32** | Index of the active NSX Edge appliance. Values are 0 and 1. | [optional] [default to null]
**SystemStatus** | **string** | System status of the active NSX Edge appliance. | [optional] [default to null]
**HaVnicInUse** | **int32** | Index of the vnic consumed for NSX Edge HA. | [optional] [default to null]
**EdgeStatus** | **string** | NSX Edge appliance health status identified by GREY (unknown status), GREEN (health checks are successful), YELLOW (intermittent health check failure), RED (none of the appliances are in serving state). If health check fails for 5 consecutive times for all appliance (2 for HA else 1) then status will turn from YELLOW to RED. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

