# VsanConfigConstraints

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCapacity** | **int64** | Maximum capacity supported for cluster (GiB). | [default to null]
**RecommendedCapacities** | **[]int64** | List of supported capacities which may offer preferable performance (GiB). | [default to null]
**SupportedCapacityIncrement** | **int64** | Increment to be added to min_capacity to result in a supported capacity (GiB). | [optional] [default to null]
**MinCapacity** | **int64** | Minimum capacity supported for cluster (GiB). | [default to null]
**NumHosts** | **int64** | Number of hosts in cluster. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

