# Nsxnatrule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vnic** | **string** | Interface on which the NAT rule is applied. | [optional] [default to null]
**RuleType** | **string** | Identifies the type of the rule. internal_high or user. | [optional] [default to null]
**Protocol** | **string** | Protocol. Default is &#x27;any&#x27; | [optional] [default to null]
**Description** | **string** | Description for the rule. | [optional] [default to null]
**RuleId** | **int64** | Identifier for the rule. | [optional] [default to null]
**SnatMatchDestinationPort** | **string** | Apply SNAT rule only if traffic has this destination port. Default is &#x27;any&#x27;. | [optional] [default to null]
**OriginalAddress** | **string** | Original address or address range. This is the original source address for SNAT rules and the original destination address for DNAT rules. | [optional] [default to null]
**DnatMatchSourceAddress** | **string** | Apply DNAT rule only if traffic has this source address. Default is &#x27;any&#x27;. | [optional] [default to null]
**DnatMatchSourcePort** | **string** | Apply DNAT rule only if traffic has this source port. Default is &#x27;any&#x27;. | [optional] [default to null]
**SnatMatchDestinationAddress** | **string** | Apply SNAT rule only if traffic has this destination address. Default is &#x27;any&#x27;. | [optional] [default to null]
**OriginalPort** | **string** | Original port. This is the original source port for SNAT rules, and the original destination port for DNAT rules. | [optional] [default to null]
**LoggingEnabled** | **bool** | Enable logging for the rule. | [optional] [default to null]
**TranslatedAddress** | **string** | Translated address or address range. | [optional] [default to null]
**Enabled** | **bool** | Enable rule. | [optional] [default to null]
**IcmpType** | **string** | ICMP type. Only supported when protocol is icmp. Default is &#x27;any&#x27;. | [optional] [default to null]
**TranslatedPort** | **string** | Translated port. Supported in DNAT rules only. | [optional] [default to null]
**Action** | **string** | Action for the rule. SNAT or DNAT. | [optional] [default to null]
**RuleTag** | **int64** | Rule tag. Used to specify user-defined ruleId. If not specified NSX Manager will generate ruleId. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

