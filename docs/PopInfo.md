# PopInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmiInfos** | [**map[string]PopAmiInfo**](PopAmiInfo.md) | A map of [region name of PoP / PoP-AMI]:[PopAmiInfo]. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The PopInfo (or PoP AMI) created time. Using ISO 8601 date-time pattern. | [optional] [default to null]
**ServiceInfos** | [**map[string]PopServiceInfo**](PopServiceInfo.md) | A map of [service type]:[PopServiceInfo] | [optional] [default to null]
**Id** | **string** | UUID of the PopInfo | [optional] [default to null]
**ManifestVersion** | **string** | version of the manifest. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

