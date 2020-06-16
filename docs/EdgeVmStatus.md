# EdgeVmStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | High Availability index of the appliance. Values are 0 and 1. | [optional] [default to null]
**HaState** | **string** | High Availability state of the appliance. Values are active and standby. | [optional] [default to null]
**Name** | **string** | Name of the NSX Edge appliance. | [optional] [default to null]
**Id** | **string** | vCenter MOID of the NSX Edge appliance. | [optional] [default to null]
**EdgeVMStatus** | **string** | NSX Edge appliance health status identified by GREY (unknown status), GREEN (health checks are successful), YELLOW (intermittent health check failure), RED (appliance not in serving state). | [optional] [default to null]
**PreRulesGenerationNumber** | **int64** | Value of the last published pre rules generation number. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

