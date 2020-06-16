# Nsxfirewallrule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | **string** | Identifies the type of the rule. internal_high or user. | [optional] [default to null]
**Description** | **string** | Description for the rule | [optional] [default to null]
**RuleId** | **int64** | Identifier for the rule. | [optional] [default to null]
**MatchTranslated** | **bool** | Defines the order of NAT and Firewall pipeline. When false, firewall happens before NAT. Default : false | [optional] [default to null]
**InvalidApplication** | **bool** |  | [optional] [default to null]
**Direction** | **string** | Direction. Possible values in or out. Default is &#x27;any&#x27;. | [optional] [default to null]
**Statistics** | [***FirewallRuleStats**](firewallRuleStats.md) |  | [optional] [default to null]
**Name** | **string** | Name for the rule. | [optional] [default to null]
**InvalidSource** | **bool** |  | [optional] [default to null]
**LoggingEnabled** | **bool** | Enable logging for the rule. | [optional] [default to null]
**Destination** | [***AddressFwSourceDestination**](addressFWSourceDestination.md) |  | [optional] [default to null]
**Enabled** | **bool** | Enable rule. | [optional] [default to null]
**Application** | [***Application**](application.md) |  | [optional] [default to null]
**Source** | [***AddressFwSourceDestination**](addressFWSourceDestination.md) |  | [optional] [default to null]
**Action** | **string** | Action. Values : accept, deny | [optional] [default to null]
**InvalidDestination** | **bool** |  | [optional] [default to null]
**RuleTag** | **int64** | Rule tag. Used to specify user-defined ruleId. If not specified NSX Manager will generate ruleId. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

