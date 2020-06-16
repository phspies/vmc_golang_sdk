# FirewallConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallRules** | [***FirewallRules**](firewallRules.md) |  | [optional] [default to null]
**FeatureType** | **string** |  | [optional] [default to null]
**Version** | **int64** | Version number tracking each configuration change. To avoid problems with overwriting changes, always retrieve and modify the latest configuration to include the current version number in your request. If you provide a version number which is not current, the request is rejected. If you omit the version number, the request is accepted but may overwrite any current changes if your change is not in sync with the latest change. | [optional] [default to null]
**Template** | **string** |  | [optional] [default to null]
**GlobalConfig** | [***FirewallGlobalConfig**](firewallGlobalConfig.md) |  | [optional] [default to null]
**Enabled** | **bool** | Value is true if feature is enabled. Default value is true. Optional. | [optional] [default to null]
**DefaultPolicy** | [***FirewallDefaultPolicy**](firewallDefaultPolicy.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

