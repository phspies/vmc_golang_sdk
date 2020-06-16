# Agent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalIp** | **string** | The internal IP address of the agent which is provided by the underlying cloud provider | [optional] [default to null]
**AgentUrl** | **string** | The accessible URL of the agent service, it is resolved to public IP address from the Internet and private IP address within SDDC | [optional] [default to null]
**ManagementIp** | **string** | The internal management IP address of the agent exposed to the SDDC, which might be different from the internal IP | [optional] [default to null]
**HostnameVerifierEnabled** | **bool** | Boolean flag to indicate if the agent is using FQDN in the certificate | [optional] [default to null]
**Master** | **bool** | Boolean flag to indicate if the agent is the master, only the master Agent is accessible | [optional] [default to null]
**NetworkNetmask** | **string** | Network netmask of the agent | [optional] [default to null]
**NetworkGateway** | **string** | Network gateway of the agent | [optional] [default to null]
**Provider** | **string** | The cloud provider | [default to null]
**CertEnabled** | **bool** | Boolean flag to indicate if the agent is using CA signed certificate | [optional] [default to null]
**AgentState** | **string** | Agent state | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

