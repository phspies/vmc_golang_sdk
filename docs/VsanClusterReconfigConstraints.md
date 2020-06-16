# VsanClusterReconfigConstraints

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReconfigBiases** | [**[]VsanClusterReconfigBias**](VsanClusterReconfigBias.md) | Biases to reconfigure vSAN in an existing cluster. | [default to null]
**AvailableCapacities** | [**map[string][]VsanAvailableCapacity**](array.md) | A map of VsanClusterReconfigBias id to the list of VsanAvailableCapacity. It gives all of available vSAN capacities for each of reconfiguration biases.  | [default to null]
**DefaultCapacities** | [**map[string]VsanAvailableCapacity**](VsanAvailableCapacity.md) | A map of VsanClusterReconfigBias id to a VsanAvailableCapacity. It gives the default VsanAvailableCapacity for each of reconfiguration biases.  | [default to null]
**Hosts** | **int32** | The number of hosts in a cluster for the constraints. | [default to null]
**DefaultReconfigBiasId** | **string** | The id of default VsanClusterReconfigBias for this constraints. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

