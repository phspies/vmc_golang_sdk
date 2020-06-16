# FirewallRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | **string** |  | [optional] [default to null]
**ApplicationIds** | **[]string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**RuleInterface** | **string** | Deprecated, left for backwards compatibility. Remove once UI stops using it. | [optional] [default to null]
**Destination** | **string** | Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to &#x27;any&#x27;. | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**DestinationScope** | [***FirewallRuleScope**](FirewallRuleScope.md) |  | [optional] [default to null]
**Source** | **string** | Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to &#x27;any&#x27;. | [optional] [default to null]
**SourceScope** | [***FirewallRuleScope**](FirewallRuleScope.md) |  | [optional] [default to null]
**Services** | [**[]FirewallService**](FirewallService.md) | list of protocols and ports for this firewall rule | [optional] [default to null]
**Action** | **string** |  | [optional] [default to null]
**Revision** | **int32** | current revision of the list of firewall rules, used to protect against concurrent modification (first writer wins) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

