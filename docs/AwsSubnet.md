# AwsSubnet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedAccountId** | **string** | The connected account ID this subnet is accessible through. This is an internal ID formatted as a UUID specific to Skyscraper. | [optional] [default to null]
**RegionName** | **string** | The region this subnet is in, usually in the form of country code, general location, and a number (ex. us-west-2). | [optional] [default to null]
**AvailabilityZone** | **string** | The availability zone this subnet is in, which should be the region name plus one extra letter (ex. us-west-2a). | [optional] [default to null]
**SubnetId** | **string** | The subnet ID in AWS, provided in the form &#x27;subnet-######&#x27;. | [optional] [default to null]
**SubnetCidrBlock** | **string** | The CIDR block of the subnet, in the form of &#x27;#.#.#.#/#&#x27;. | [optional] [default to null]
**IsCompatible** | **bool** | Flag indicating whether this subnet is compatible. If true, this is a valid choice for the customer to deploy a SDDC in. | [optional] [default to null]
**VpcId** | **string** | The VPC ID the subnet resides in within AWS. Tends to be &#x27;vpc-#######&#x27;. | [optional] [default to null]
**VpcCidrBlock** | **string** | The CIDR block of the VPC, in the form of &#x27;#.#.#.#/#&#x27;. | [optional] [default to null]
**Name** | **string** | Optional field (may not be provided by AWS), indicates the found name tag for the subnet. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

