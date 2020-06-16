# FirewallGlobalConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TcpAllowOutOfWindowPackets** | **bool** | Allow TCP out of window packets. | [optional] [default to null]
**UdpTimeout** | **int32** | UDP timeout close. | [optional] [default to null]
**IpGenericTimeout** | **int32** | IP generic timeout. | [optional] [default to null]
**TcpPickOngoingConnections** | **bool** | Pick TCP ongoing connections. | [optional] [default to null]
**TcpTimeoutOpen** | **int32** | TCP timeout open. | [optional] [default to null]
**TcpTimeoutClose** | **int32** | TCP timeout close. | [optional] [default to null]
**Icmp6Timeout** | **int32** | ICMP6 timeout. | [optional] [default to null]
**DropIcmpReplays** | **bool** | Drop icmp replays. | [optional] [default to null]
**LogIcmpErrors** | **bool** | Log icmp errors. | [optional] [default to null]
**TcpSendResetForClosedVsePorts** | **bool** | Send TCP reset for closed NSX Edge ports. | [optional] [default to null]
**DropInvalidTraffic** | **bool** | Drop invalid traffic. | [optional] [default to null]
**EnableSynFloodProtection** | **bool** | Protect against SYN flood attacks by detecting bogus TCP connections and terminating them without consuming firewall state tracking resources. Default : false | [optional] [default to null]
**IcmpTimeout** | **int32** | ICMP timeout. | [optional] [default to null]
**TcpTimeoutEstablished** | **int32** | TCP timeout established. | [optional] [default to null]
**LogInvalidTraffic** | **bool** | Log invalid traffic. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

