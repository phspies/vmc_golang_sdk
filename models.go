
package vmc

import (
	"time"
)

type AbstractEntity struct {
	Updated time.Time `json:"updated"`
	// User id that last updated this record
	UserId string `json:"user_id"`
	// User id that last updated this record
	UpdatedByUserId string `json:"updated_by_user_id"`
	Created time.Time `json:"created"`
	// Version of this entity
	Version int32 `json:"version"`
	// User name that last updated this record
	UpdatedByUserName string `json:"updated_by_user_name,omitempty"`
	// User name that last updated this record
	UserName string `json:"user_name"`
	// Unique ID for this entity
	Id string `json:"id"`
}



type AccountLinkConfig struct {
	// Boolean flag identifying whether account linking should be delayed or not for the SDDC.
	DelayAccountLink bool `json:"delay_account_link,omitempty"`
}



type AccountLinkSddcConfig struct {
	CustomerSubnetIds []string `json:"customer_subnet_ids,omitempty"`
	// The ID of the customer connected account to work with.
	ConnectedAccountId string `json:"connected_account_id,omitempty"`
}



// Source or Destination for firewall rule. Default is 'any'.
type AddressFwSourceDestination struct {
	// Exclude the specified source or destination.
	Exclude bool `json:"exclude,omitempty"`
	// List of string. Can specify single IP address, range of IP address, or in CIDR format. Can define multiple.
	IpAddress []string `json:"ipAddress,omitempty"`
	// List of string. Id of cluster, datacenter, distributedPortGroup, legacyPortGroup, VirtualMachine, vApp, resourcePool, logicalSwitch, IPSet, securityGroup. Can define multiple.
	GroupingObjectId []string `json:"groupingObjectId,omitempty"`
	// List of string. Possible values are vnic-index-[1-9], vse, external or internal. Can define multiple.
	VnicGroupId []string `json:"vnicGroupId,omitempty"`
}



type Agent struct {
	// The internal IP address of the agent which is provided by the underlying cloud provider
	InternalIp string `json:"internal_ip,omitempty"`
	// The accessible URL of the agent service, it is resolved to public IP address from the Internet and private IP address within SDDC
	AgentUrl string `json:"agent_url,omitempty"`
	// The internal management IP address of the agent exposed to the SDDC, which might be different from the internal IP
	ManagementIp string `json:"management_ip,omitempty"`
	// Boolean flag to indicate if the agent is using FQDN in the certificate
	HostnameVerifierEnabled bool `json:"hostname_verifier_enabled,omitempty"`
	// Boolean flag to indicate if the agent is the master, only the master Agent is accessible
	Master bool `json:"master,omitempty"`
	// Network netmask of the agent
	NetworkNetmask string `json:"network_netmask,omitempty"`
	// Network gateway of the agent
	NetworkGateway string `json:"network_gateway,omitempty"`
	// The cloud provider
	Provider string `json:"provider"`
	// Boolean flag to indicate if the agent is using CA signed certificate
	CertEnabled bool `json:"cert_enabled,omitempty"`
	// Agent state
	AgentState string `json:"agent_state,omitempty"`
}



// the AmiInfo used for deploying esx of the sddc
type AmiInfo struct {
	// instance type of the esx ami
	InstanceType string `json:"instance_type,omitempty"`
	// the region of the esx ami
	Region string `json:"region,omitempty"`
	// the ami id for the esx
	Id string `json:"id,omitempty"`
	// the name of the esx ami
	Name string `json:"name,omitempty"`
}



// NSX Edge appliance summary.
type AppliancesSummary struct {
	// vCenter MOID of the active NSX Edge appliance's data store.
	DataStoreMoidOfActiveVse string `json:"dataStoreMoidOfActiveVse,omitempty"`
	// Value is true if FIPS is enabled on NSX Edge appliance.
	EnableFips bool `json:"enableFips,omitempty"`
	// Host name of the active NSX Edge appliance.
	HostNameOfActiveVse string `json:"hostNameOfActiveVse,omitempty"`
	// NSX Edge appliance build version.
	VmBuildInfo string `json:"vmBuildInfo,omitempty"`
	// Value is true if NSX Edge appliances are to be deployed.
	DeployAppliances bool `json:"deployAppliances,omitempty"`
	// Communication channel used to communicate with NSX Edge appliance.
	CommunicationChannel string `json:"communicationChannel,omitempty"`
	// Name of the active NSX Edge appliance.
	VmNameOfActiveVse string `json:"vmNameOfActiveVse,omitempty"`
	// Number of deployed appliances of the NSX Edge.
	NumberOfDeployedVms int32 `json:"numberOfDeployedVms,omitempty"`
	// vCenter MOID of the active NSX Edge appliance's resource pool/cluster. Can be resource pool ID, e.g. resgroup-15 or cluster ID, e.g. domain-c41.
	ResourcePoolMoidOfActiveVse string `json:"resourcePoolMoidOfActiveVse,omitempty"`
	// Datastore name of the active NSX Edge appliance.
	DataStoreNameOfActiveVse string `json:"dataStoreNameOfActiveVse,omitempty"`
	// vCenter MOID of the active NSX Edge appliance.
	VmMoidOfActiveVse string `json:"vmMoidOfActiveVse,omitempty"`
	// Time stamp value when healthcheck status was last updated for the NSX Edge appliance.
	StatusFromVseUpdatedOn int64 `json:"statusFromVseUpdatedOn,omitempty"`
	// FQDN of the NSX Edge.
	Fqdn string `json:"fqdn,omitempty"`
	// NSX Edge appliance size.
	ApplianceSize string `json:"applianceSize,omitempty"`
	// Resource Pool/Cluster name of the active NSX Edge appliance.
	ResourcePoolNameOfActiveVse string `json:"resourcePoolNameOfActiveVse,omitempty"`
	// HA index of the active NSX Edge appliance.
	ActiveVseHaIndex int32 `json:"activeVseHaIndex,omitempty"`
	// NSX Edge appliance version.
	VmVersion string `json:"vmVersion,omitempty"`
	// vCenter MOID of the active NSX Edge appliance's host.
	HostMoidOfActiveVse string `json:"hostMoidOfActiveVse,omitempty"`
}



// Application for firewall rule
type Application struct {
	// List of string. Id of service or serviceGroup groupingObject. Can define multiple.
	ApplicationId []string `json:"applicationId,omitempty"`
	// List of protocol and ports. Can define multiple.
	Service []Nsxfirewallservice `json:"service,omitempty"`
}



type AvailableZoneInfo struct {
	Subnets []Subnet `json:"subnets,omitempty"`
	// available zone name
	Name string `json:"name,omitempty"`
}



type AwsAgent struct {
	// The internal IP address of the agent which is provided by the underlying cloud provider
	InternalIp string `json:"internal_ip,omitempty"`
	// The accessible URL of the agent service, it is resolved to public IP address from the Internet and private IP address within SDDC
	AgentUrl string `json:"agent_url,omitempty"`
	// The internal management IP address of the agent exposed to the SDDC, which might be different from the internal IP
	ManagementIp string `json:"management_ip,omitempty"`
	// Boolean flag to indicate if the agent is using FQDN in the certificate
	HostnameVerifierEnabled bool `json:"hostname_verifier_enabled,omitempty"`
	// Boolean flag to indicate if the agent is the master, only the master Agent is accessible
	Master bool `json:"master,omitempty"`
	// Network netmask of the agent
	NetworkNetmask string `json:"network_netmask,omitempty"`
	// Network gateway of the agent
	NetworkGateway string `json:"network_gateway,omitempty"`
	// The cloud provider
	Provider string `json:"provider"`
	// Boolean flag to indicate if the agent is using CA signed certificate
	CertEnabled bool `json:"cert_enabled,omitempty"`
	// Agent state
	AgentState string `json:"agent_state,omitempty"`
	InstanceId string `json:"instance_id,omitempty"`
	KeyPair *AwsKeyPair `json:"key_pair,omitempty"`
}



type AwsCloudProvider struct {
	// Name of the Cloud Provider
	Provider string `json:"provider"`
	Regions []string `json:"regions,omitempty"`
}



type AwsCompatibleSubnets struct {
	CustomerAvailableZones []string `json:"customer_available_zones,omitempty"`
	VpcMap map[string]VpcInfoSubnets `json:"vpc_map,omitempty"`
}


import (
	"time"
)

type AwsCustomerConnectedAccount struct {
	Updated time.Time `json:"updated"`
	// User id that last updated this record
	UserId string `json:"user_id"`
	// User id that last updated this record
	UpdatedByUserId string `json:"updated_by_user_id"`
	Created time.Time `json:"created"`
	// Version of this entity
	Version int32 `json:"version"`
	// User name that last updated this record
	UpdatedByUserName string `json:"updated_by_user_name,omitempty"`
	// User name that last updated this record
	UserName string `json:"user_name"`
	// Unique ID for this entity
	Id string `json:"id"`
	PolicyPayerArn string `json:"policy_payer_arn,omitempty"`
	// Provides a map of regions to availability zones from the shadow account's perspective
	RegionToAzToShadowMapping map[string]map[string]string `json:"region_to_az_to_shadow_mapping,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	CfStackName string `json:"cf_stack_name,omitempty"`
	State string `json:"state,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	PolicyServiceArn string `json:"policy_service_arn,omitempty"`
	PolicyExternalId string `json:"policy_external_id,omitempty"`
	PolicyPayerLinkedArn string `json:"policy_payer_linked_arn,omitempty"`
}



type AwsEsxHost struct {
	Name string `json:"name,omitempty"`
	// Availability zone where the host is provisioned.
	AvailabilityZone string `json:"availability_zone,omitempty"`
	EsxId string `json:"esx_id,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Provider string `json:"provider"`
	// Backing cloud provider instance type for host.
	InstanceType string `json:"instance_type,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	CustomProperties map[string]string `json:"custom_properties,omitempty"`
	EsxState string `json:"esx_state,omitempty"`
	InternalPublicIpPool []SddcPublicIp `json:"internal_public_ip_pool,omitempty"`
}



type AwsKeyPair struct {
	KeyName string `json:"key_name,omitempty"`
	KeyFingerprint string `json:"key_fingerprint,omitempty"`
	KeyMaterial string `json:"key_material,omitempty"`
}



type AwsKmsInfo struct {
	// The ARN associated with the customer master key for this cluster.
	AmazonResourceName string `json:"amazon_resource_name"`
}



type AwsSddcConfig struct {
	// Whether this sddc is reduced capacity 1NODE.
	OneNodeReducedCapacity bool `json:"one_node_reduced_capacity,omitempty"`
	// AWS VPC IP range. Only prefix of 16 or 20 is currently supported.
	VpcCidr string `json:"vpc_cidr,omitempty"`
	HostInstanceType string `json:"host_instance_type,omitempty"`
	// skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan bool `json:"skip_creating_vxlan,omitempty"`
	// VXLAN IP subnet in CIDR for compute gateway
	VxlanSubnet string `json:"vxlan_subnet,omitempty"`
	// The size of the vCenter and NSX appliances. \"large\" sddcSize corresponds to a 'large' vCenter appliance and 'large' NSX appliance. 'medium' sddcSize corresponds to 'medium' vCenter appliance and 'medium' NSX appliance. Value defaults to 'medium'. 
	Size string `json:"size,omitempty"`
	// The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of seperable storage will be used. 
	StorageCapacity int64 `json:"storage_capacity,omitempty"`
	Name string `json:"name"`
	// A list of the SDDC linking configurations to use.
	AccountLinkSddcConfig []AccountLinkSddcConfig `json:"account_link_sddc_config,omitempty"`
	// If provided, will be assigned as SDDC id of the provisioned SDDC.
	SddcId string `json:"sddc_id,omitempty"`
	NumHosts int32 `json:"num_hosts"`
	// Denotes the sddc type , if the value is null or empty, the type is considered as default.
	SddcType string `json:"sddc_type,omitempty"`
	AccountLinkConfig *AccountLinkConfig `json:"account_link_config,omitempty"`
	// Determines what additional properties are available based on cloud provider.
	Provider string `json:"provider"`
	// The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.
	SsoDomain string `json:"sso_domain,omitempty"`
	// If provided, configuration from the template will applied to the provisioned SDDC.
	SddcTemplateId string `json:"sddc_template_id,omitempty"`
	// Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ.
	DeploymentType string `json:"deployment_type,omitempty"`
	Region string `json:"region"`
}


import (
	"time"
)

type AwsSddcConnection struct {
	Updated time.Time `json:"updated"`
	// User id that last updated this record
	UserId string `json:"user_id"`
	// User id that last updated this record
	UpdatedByUserId string `json:"updated_by_user_id"`
	Created time.Time `json:"created"`
	// Version of this entity
	Version int32 `json:"version"`
	// User name that last updated this record
	UpdatedByUserName string `json:"updated_by_user_name,omitempty"`
	// User name that last updated this record
	UserName string `json:"user_name"`
	// Unique ID for this entity
	Id string `json:"id"`
	// The CIDR block of the customer's subnet this link is in.
	CidrBlockSubnet string `json:"cidr_block_subnet,omitempty"`
	// The corresponding connected (customer) account UUID this connection is attached to.
	ConnectedAccountId string `json:"connected_account_id,omitempty"`
	// Which group the ENIs belongs to. (deprecated)
	EniGroup string `json:"eni_group,omitempty"`
	// The ID of the subnet this link is to.
	SubnetId string `json:"subnet_id,omitempty"`
	// Determines whether the CGW is present in this connection set or not. Used for multi-az deployments.
	CgwPresent bool `json:"cgw_present,omitempty"`
	// The org this link belongs to.
	OrgId string `json:"org_id,omitempty"`
	// The SDDC this link is used for.
	SddcId string `json:"sddc_id,omitempty"`
	// The CIDR block of the customer's VPC.
	CidrBlockVpc string `json:"cidr_block_vpc,omitempty"`
	// The order of the connection
	ConnectionOrder int32 `json:"connection_order,omitempty"`
	// The state of the connection.
	State string `json:"state,omitempty"`
	// Which availability zone is this connection in?
	SubnetAvailabilityZone string `json:"subnet_availability_zone,omitempty"`
	// The VPC ID of the subnet this link is to.
	VpcId string `json:"vpc_id,omitempty"`
	// A list of all ENIs used for this connection.
	CustomerEniInfos []CustomerEniInfo `json:"customer_eni_infos,omitempty"`
	// The default routing table in the customer's VPC.
	DefaultRouteTable string `json:"default_route_table,omitempty"`
}



type AwsSddcResourceConfig struct {
	// Name for management appliance network.
	MgmtApplianceNetworkName string `json:"mgmt_appliance_network_name,omitempty"`
	// if true, NSX-T UI is enabled.
	Nsxt bool `json:"nsxt,omitempty"`
	// Management Gateway Id
	MgwId string `json:"mgw_id,omitempty"`
	// URL of the NSX Manager
	NsxMgrUrl string `json:"nsx_mgr_url,omitempty"`
	// PSC internal management IP
	PscManagementIp string `json:"psc_management_ip,omitempty"`
	// URL of the PSC server
	PscUrl string `json:"psc_url,omitempty"`
	Cgws []string `json:"cgws,omitempty"`
	// Availability zones over which esx hosts are provisioned. MultiAZ SDDCs will have hosts provisioned over two availability zones while SingleAZ SDDCs will provision over one.
	AvailabilityZones []string `json:"availability_zones,omitempty"`
	// The ManagedObjectReference of the management Datastore
	ManagementDs string `json:"management_ds,omitempty"`
	// nsx api entire base url
	NsxApiPublicEndpointUrl string `json:"nsx_api_public_endpoint_url,omitempty"`
	CustomProperties map[string]string `json:"custom_properties,omitempty"`
	// Password for vCenter SDDC administrator
	CloudPassword string `json:"cloud_password,omitempty"`
	// Discriminator for additional properties
	Provider string `json:"provider"`
	// List of clusters in the SDDC.
	Clusters []Cluster `json:"clusters,omitempty"`
	// vCenter internal management IP
	VcManagementIp string `json:"vc_management_ip,omitempty"`
	SddcNetworks []string `json:"sddc_networks,omitempty"`
	// Username for vCenter SDDC administrator
	CloudUsername string `json:"cloud_username,omitempty"`
	EsxHosts []AwsEsxHost `json:"esx_hosts,omitempty"`
	// NSX Manager internal management IP
	NsxMgrManagementIp string `json:"nsx_mgr_management_ip,omitempty"`
	// unique id of the vCenter server
	VcInstanceId string `json:"vc_instance_id,omitempty"`
	// Cluster Id to add ESX workflow
	EsxClusterId string `json:"esx_cluster_id,omitempty"`
	// vCenter public IP
	VcPublicIp string `json:"vc_public_ip,omitempty"`
	// skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan bool `json:"skip_creating_vxlan,omitempty"`
	// URL of the vCenter server
	VcUrl string `json:"vc_url,omitempty"`
	SddcManifest *SddcManifest `json:"sddc_manifest,omitempty"`
	// VXLAN IP subnet
	VxlanSubnet string `json:"vxlan_subnet,omitempty"`
	// Group name for vCenter SDDC administrator
	CloudUserGroup string `json:"cloud_user_group,omitempty"`
	ManagementRp string `json:"management_rp,omitempty"`
	// region in which sddc is deployed
	Region string `json:"region,omitempty"`
	// Availability zone where the witness node is provisioned for a MultiAZ SDDC. This is null for a SingleAZ SDDC.
	WitnessAvailabilityZone string `json:"witness_availability_zone,omitempty"`
	// sddc identifier
	SddcId string `json:"sddc_id,omitempty"`
	PopAgentXeniConnection *PopAgentXeniConnection `json:"pop_agent_xeni_connection,omitempty"`
	// List of Controller IPs
	NsxControllerIps []string `json:"nsx_controller_ips,omitempty"`
	// ESX host subnet
	EsxHostSubnet string `json:"esx_host_subnet,omitempty"`
	// The SSO domain name to use for vSphere users
	SsoDomain string `json:"sso_domain,omitempty"`
	// Denotes if this is a SingleAZ SDDC or a MultiAZ SDDC.
	DeploymentType string `json:"deployment_type,omitempty"`
	NsxtAddons *NsxtAddons `json:"nsxt_addons,omitempty"`
	// if true, use the private IP addresses to register DNS records for the management VMs
	DnsWithManagementVmPrivateIp bool `json:"dns_with_management_vm_private_ip,omitempty"`
	BackupRestoreBucket string `json:"backup_restore_bucket,omitempty"`
	PublicIpPool []SddcPublicIp `json:"public_ip_pool,omitempty"`
	VpcInfo *VpcInfo `json:"vpc_info,omitempty"`
	KmsVpcEndpoint *KmsVpcEndpoint `json:"kms_vpc_endpoint,omitempty"`
	// maximum number of public IP that user can allocate.
	MaxNumPublicIp int32 `json:"max_num_public_ip,omitempty"`
	AccountLinkSddcConfig []SddcLinkConfig `json:"account_link_sddc_config,omitempty"`
	VsanEncryptionConfig *VsanEncryptionConfig `json:"vsan_encryption_config,omitempty"`
	VpcInfoPeeredAgent *VpcInfo `json:"vpc_info_peered_agent,omitempty"`
}



type AwsSubnet struct {
	// The connected account ID this subnet is accessible through. This is an internal ID formatted as a UUID specific to Skyscraper.
	ConnectedAccountId string `json:"connected_account_id,omitempty"`
	// The region this subnet is in, usually in the form of country code, general location, and a number (ex. us-west-2).
	RegionName string `json:"region_name,omitempty"`
	// The availability zone this subnet is in, which should be the region name plus one extra letter (ex. us-west-2a).
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// The subnet ID in AWS, provided in the form 'subnet-######'.
	SubnetId string `json:"subnet_id,omitempty"`
	// The CIDR block of the subnet, in the form of '#.#.#.#/#'.
	SubnetCidrBlock string `json:"subnet_cidr_block,omitempty"`
	// Flag indicating whether this subnet is compatible. If true, this is a valid choice for the customer to deploy a SDDC in.
	IsCompatible bool `json:"is_compatible,omitempty"`
	// The VPC ID the subnet resides in within AWS. Tends to be 'vpc-#######'.
	VpcId string `json:"vpc_id,omitempty"`
	// The CIDR block of the VPC, in the form of '#.#.#.#/#'.
	VpcCidrBlock string `json:"vpc_cidr_block,omitempty"`
	// Optional field (may not be provided by AWS), indicates the found name tag for the subnet.
	Name string `json:"name,omitempty"`
}



// CA certificate list. Optional.
type CaCertificates struct {
	CaCertificate []string `json:"caCertificate,omitempty"`
}



// Statistics data for each vnic.
type CbmStatistic struct {
	// Vnic index.
	Vnic int32 `json:"vnic,omitempty"`
	// Timestamp value.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Tx rate (Kilobits per second - kbps)
	Out float64 `json:"out,omitempty"`
	// Rx rate (Kilobits per second - kbps)
	In float64 `json:"in,omitempty"`
}



// NSX Edge Interface Statistics.
type CbmStatistics struct {
	DataDto *CbmStatsData `json:"dataDto,omitempty"`
	MetaDto *MetaDashboardStats `json:"metaDto,omitempty"`
}



// Statistics data.
type CbmStatsData struct {
	Vnic9 []CbmStatistic `json:"vnic_9,omitempty"`
	Vnic8 []CbmStatistic `json:"vnic_8,omitempty"`
	Vnic7 []CbmStatistic `json:"vnic_7,omitempty"`
	Vnic6 []CbmStatistic `json:"vnic_6,omitempty"`
	Vnic5 []CbmStatistic `json:"vnic_5,omitempty"`
	Vnic4 []CbmStatistic `json:"vnic_4,omitempty"`
	Vnic3 []CbmStatistic `json:"vnic_3,omitempty"`
	Vnic2 []CbmStatistic `json:"vnic_2,omitempty"`
	Vnic1 []CbmStatistic `json:"vnic_1,omitempty"`
	Vnic0 []CbmStatistic `json:"vnic_0,omitempty"`
}



type CloudProvider struct {
	// Name of the Cloud Provider
	Provider string `json:"provider"`
}



type Cluster struct {
	EsxHostList []AwsEsxHost `json:"esx_host_list,omitempty"`
	ClusterState string `json:"cluster_state,omitempty"`
	AwsKmsInfo *AwsKmsInfo `json:"aws_kms_info,omitempty"`
	EsxHostInfo *EsxHostInfo `json:"esx_host_info,omitempty"`
	// Number of cores enabled on ESX hosts added to this cluster
	HostCpuCoresCount int32 `json:"host_cpu_cores_count,omitempty"`
	ClusterCapacity *EntityCapacity `json:"cluster_capacity,omitempty"`
	ClusterId string `json:"cluster_id"`
	ClusterName string `json:"cluster_name,omitempty"`
}



type ClusterConfig struct {
	// Customize CPU cores on hosts in a cluster. Specify number of cores to be enabled on hosts in a cluster.
	HostCpuCoresCount int32 `json:"host_cpu_cores_count,omitempty"`
	HostInstanceType string `json:"host_instance_type,omitempty"`
	// For EBS-backed instances only, the requested storage capacity in GiB.
	StorageCapacity int64 `json:"storage_capacity,omitempty"`
	NumHosts int32 `json:"num_hosts"`
}



type ClusterReconfigureParams struct {
	// The final desired storage capacity after reconfiguring the cluster in GiB.
	StorageCapacity int64 `json:"storage_capacity,omitempty"`
	// Bias value as obtained from the storage constraints call.
	Bias string `json:"bias,omitempty"`
	// Number of hosts in the cluster after reconfiguring.
	NumHosts int32 `json:"num_hosts"`
}



type ComputeGatewayTemplate struct {
	PublicIp *SddcPublicIp `json:"public_ip,omitempty"`
	PrimaryDns string `json:"primary_dns,omitempty"`
	SecondaryDns string `json:"secondary_dns,omitempty"`
	FirewallRules []FirewallRule `json:"firewall_rules,omitempty"`
	Vpns []Vpn `json:"vpns,omitempty"`
	LogicalNetworks []LogicalNetwork `json:"logical_networks,omitempty"`
	NatRules []NatRule `json:"nat_rules,omitempty"`
	L2Vpn *interface{} `json:"l2_vpn,omitempty"`
}



// Represents a configuration spec for any sddc provision operation.
type ConfigSpec struct {
	// Indicates after how many days the sddc should expire
	ExpiryInDays int32 `json:"expiry_in_days,omitempty"`
	// Map of region to instance types available in that region
	Availability map[string][]InstanceTypeConfig `json:"availability,omitempty"`
}



type ConnectivityAgentValidation struct {
	// source appliance of connectivity test, i.e. VCENTER, SRM, VR.
	Source string `json:"source,omitempty"`
	// type of connectivity test, i.e. PING, TRACEROUTE, DNS, CONNECTIVITY, CURL. For CONNECTIVITY and CURL tests only, please specify the ports to be tested against.
	Type_ string `json:"type,omitempty"`
	// TCP ports ONLY for CONNECTIVITY and CURL tests.
	Ports []string `json:"ports,omitempty"`
	// URL path ONLY for CURL tests.
	Path string `json:"path,omitempty"`
}



type ConnectivityValidationGroup struct {
	// test group id, currently, only HLM.
	Id string `json:"id,omitempty"`
	// Name of the test group.
	Name string `json:"name,omitempty"`
	// List of sub groups.
	SubGroups []ConnectivityValidationSubGroup `json:"sub_groups,omitempty"`
}



type ConnectivityValidationGroups struct {
	// List of groups.
	Groups []ConnectivityValidationGroup `json:"groups,omitempty"`
}



type ConnectivityValidationInput struct {
	// input value type, i.e. HOSTNAME_OR_IP, HOST_IP, HOSTNAME. Accept FQDN or IP address as input value when id = HOSTNAME_OR_IP, accept FQDN ONLY when id = HOSTNAME, accept IP address ONLY when id = HOST_IP.
	Id string `json:"id,omitempty"`
	// the FQDN or IP address to run the test against, use \\#primary-dns or \\#secondary-dns as the on-prem primary/secondary DNS server IP.
	Value string `json:"value,omitempty"`
	// (Optional, for UI display only) input value label.
	Label string `json:"label,omitempty"`
}



type ConnectivityValidationSubGroup struct {
	// List of user inputs for the sub group.
	Inputs []ConnectivityValidationInput `json:"inputs,omitempty"`
	// List of connectivity tests.
	Tests []ConnectivityAgentValidation `json:"tests,omitempty"`
	// Name of the sub-group.
	Label string `json:"label,omitempty"`
	// Help text.
	Help string `json:"help,omitempty"`
	// subGroup id, i.e. PRIMARY_DNS, SECONDARY_DNS, ONPREM_VCENTER, ONPREM_PSC, ACTIVE_DIRECTORY, ONPREM_ESX, DRAAS_ONPREM_VCENTER, DRAAS_ONPREM_PSC, DRAAS_ONPREM_SRM and DRAAS_ONPREM_VR.
	Id string `json:"id,omitempty"`
}



// CRL certificate list. Optional.
type CrlCertificates struct {
	CrlCertificate []string `json:"crlCertificate,omitempty"`
}



// Indicates a single cross-account ENI and its characteristics.
type CustomerEniInfo struct {
	// Indicates list of secondary IP created for this ENI.
	SecondaryIpAddresses []string `json:"secondary_ip_addresses,omitempty"`
	// Interface ID on customer account.
	EniId string `json:"eni_id,omitempty"`
	// Indicates primary address of the ENI.
	PrimaryIpAddress string `json:"primary_ip_address,omitempty"`
}



// Dashboard Statistics data.
type DashboardData struct {
	Firewall *FirewallDashboardStats `json:"firewall,omitempty"`
	Sslvpn *SslvpnDashboardStats `json:"sslvpn,omitempty"`
	Interfaces *InterfacesDashboardStats `json:"interfaces,omitempty"`
	LoadBalancer *LoadBalancerDashboardStats `json:"loadBalancer,omitempty"`
	Ipsec *IpsecDashboardStats `json:"ipsec,omitempty"`
}



type DashboardStat struct {
	Timestamp int64 `json:"timestamp,omitempty"`
	Value float64 `json:"value,omitempty"`
}



// Dashboard Statistics data.
type DashboardStatistics struct {
	DataDto *DashboardData `json:"dataDto,omitempty"`
	MetaDto *MetaDashboardStats `json:"metaDto,omitempty"`
}



type DataPageEdgeSummary struct {
	PagingInfo *PagingInfo `json:"pagingInfo,omitempty"`
	Data []EdgeSummary `json:"data,omitempty"`
}



type DataPageSddcNetwork struct {
	PagingInfo *PagingInfo `json:"pagingInfo,omitempty"`
	Data []SddcNetwork `json:"data,omitempty"`
}



type DataPermissions struct {
	SavePermission bool `json:"savePermission,omitempty"`
	PublishPermission bool `json:"publishPermission,omitempty"`
}



// DHCP lease information.
type DhcpLeaseInfo struct {
	// List of DHCP leases.
	HostLeaseInfoDtos []HostLeaseInfo `json:"hostLeaseInfoDtos,omitempty"`
}



// DHCP leases information
type DhcpLeases struct {
	// The timestamp of the DHCP lease.
	TimeStamp int64 `json:"timeStamp,omitempty"`
	HostLeaseInfosDto *DhcpLeaseInfo `json:"hostLeaseInfosDto,omitempty"`
}



// DNS configuration
type DnsConfig struct {
	FeatureType string `json:"featureType,omitempty"`
	Logging *Logging `json:"logging,omitempty"`
	// Value is true if feature is enabled. Default value is true. Optional.
	Enabled bool `json:"enabled,omitempty"`
	DnsViews *DnsViews `json:"dnsViews,omitempty"`
	Listeners *DnsListeners `json:"listeners,omitempty"`
	// Version number tracking each configuration change. To avoid problems with overwriting changes, always retrieve and modify the latest configuration to include the current version number in your request. If you provide a version number which is not current, the request is rejected. If you omit the version number, the request is accepted but may overwrite any current changes if your change is not in sync with the latest change.
	Version int64 `json:"version,omitempty"`
	Template string `json:"template,omitempty"`
	// The cache size of the DNS service.
	CacheSize int64 `json:"cacheSize,omitempty"`
	DnsServers *IpAddresses `json:"dnsServers,omitempty"`
}



// DNS forwarders.
type DnsForwarders struct {
	// IP addresses of the DNS servers.
	IpAddress []string `json:"ipAddress,omitempty"`
}



type DnsListeners struct {
	// List of IP addresses.
	IpAddress []string `json:"ipAddress,omitempty"`
	// Vnic for DNS listener.
	Vnic []string `json:"vnic,omitempty"`
	Type_ string `json:"type,omitempty"`
}



// DNS response statistics.
type DnsResponseStats struct {
	Total int64 `json:"total,omitempty"`
	FormErr int64 `json:"formErr,omitempty"`
	NxDomain int64 `json:"nxDomain,omitempty"`
	Success int64 `json:"success,omitempty"`
	ServerFail int64 `json:"serverFail,omitempty"`
	Nxrrset int64 `json:"nxrrset,omitempty"`
	Others int64 `json:"others,omitempty"`
}



// DNS statistics.
type DnsStatusAndStats struct {
	TimeStamp int64 `json:"timeStamp,omitempty"`
	Requests *Requests `json:"requests,omitempty"`
	Responses *DnsResponseStats `json:"responses,omitempty"`
	CachedDBRRSet int64 `json:"cachedDBRRSet,omitempty"`
}



// DNS View
type DnsView struct {
	// Name of the DNS view.
	Name string `json:"name"`
	ViewMatch *DnsViewMatch `json:"viewMatch,omitempty"`
	// Recursion enabled on DNS view.
	Recursion bool `json:"recursion,omitempty"`
	// Identifier for the DNS view.
	ViewId string `json:"viewId,omitempty"`
	Forwarders *DnsForwarders `json:"forwarders,omitempty"`
	// DNS view is enabled.
	Enabled bool `json:"enabled,omitempty"`
}



// Dns view match
type DnsViewMatch struct {
	Vnic []string `json:"vnic,omitempty"`
	IpSet []string `json:"ipSet,omitempty"`
	IpAddress []string `json:"ipAddress,omitempty"`
}



// DNS views.
type DnsViews struct {
	// List of DNS views.
	DnsView []DnsView `json:"dnsView,omitempty"`
}



// information for EBS-backed VSAN configuration
type EbsBackedVsanConfig struct {
	// instance type for EBS-backed VSAN configuration
	InstanceType string `json:"instance_type,omitempty"`
}


import (
	"time"
)

// Job status information for the configuration change carried out on NSX Edge.
type EdgeJob struct {
	// Job status.
	Status string `json:"status,omitempty"`
	// NSX Edge ID.
	EdgeId string `json:"edgeId,omitempty"`
	// Module information.
	Module string `json:"module,omitempty"`
	// Job ID.
	JobId string `json:"jobId,omitempty"`
	// Error code identifying the failure of the configuration change.
	ErrorCode string `json:"errorCode,omitempty"`
	// Job result information.
	Result []Result `json:"result,omitempty"`
	// Job start time.
	StartTime time.Time `json:"startTime,omitempty"`
	// Job message.
	Message string `json:"message,omitempty"`
	// Job end time.
	EndTime time.Time `json:"endTime,omitempty"`
}



// NSX Edge Appliance status.
type EdgeStatus struct {
	// Value is true if pre rules publish is enabled.
	PreRulesExists bool `json:"preRulesExists,omitempty"`
	// Individual feature status.
	FeatureStatuses []FeatureStatus `json:"featureStatuses,omitempty"`
	// Timestamp value at which the NSX Edge healthcheck was done.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Status of the latest configuration change for the NSX Edge. Values are APPLIED or PERSISTED (not published to NSX Edge Appliance yet).
	PublishStatus string `json:"publishStatus,omitempty"`
	// Value of the last published pre rules generation number.
	LastPublishedPreRulesGenerationNumber int64 `json:"lastPublishedPreRulesGenerationNumber,omitempty"`
	// Version number of the current configuration.
	Version int64 `json:"version,omitempty"`
	// Detailed status of each of the deployed NSX Edge appliances.
	EdgeVmStatus []EdgeVmStatus `json:"edgeVmStatus,omitempty"`
	// Index of the active NSX Edge appliance. Values are 0 and 1.
	ActiveVseHaIndex int32 `json:"activeVseHaIndex,omitempty"`
	// System status of the active NSX Edge appliance.
	SystemStatus string `json:"systemStatus,omitempty"`
	// Index of the vnic consumed for NSX Edge HA.
	HaVnicInUse int32 `json:"haVnicInUse,omitempty"`
	// NSX Edge appliance health status identified by GREY (unknown status), GREEN (health checks are successful), YELLOW (intermittent health check failure), RED (none of the appliances are in serving state). If health check fails for 5 consecutive times for all appliance (2 for HA else 1) then status will turn from YELLOW to RED.
	EdgeStatus string `json:"edgeStatus,omitempty"`
}



// NSX Edge summary. Read only.
type EdgeSummary struct {
	FeatureCapabilities *FeatureCapabilities `json:"featureCapabilities,omitempty"`
	// NSX Edge type, whether 'gatewayServices' or 'distributedRouter'.
	EdgeType string `json:"edgeType,omitempty"`
	LogicalRouterScopes *LogicalRouterScopes `json:"logicalRouterScopes,omitempty"`
	RecentJobInfo *EdgeJob `json:"recentJobInfo,omitempty"`
	HypervisorAssist bool `json:"hypervisorAssist,omitempty"`
	// ID generated by NSX Manager for Distributed Logical Router only.
	EdgeAssistId int64 `json:"edgeAssistId,omitempty"`
	// NSX Edge appliance health status identified by GREY (unknown status), GREEN (health checks are successful), YELLOW (intermittent health check failure), RED (none of the appliances are in serving state). If health check fails for 5 consecutive times for all appliance (2 for HA else 1) then status will turn from YELLOW to RED.
	EdgeStatus string `json:"edgeStatus,omitempty"`
	// Name derived by NSX Manager only for Distributed Logical Router.
	EdgeAssistInstanceName string `json:"edgeAssistInstanceName,omitempty"`
	ObjectId string `json:"objectId,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	// NSX Edge ID.
	Id string `json:"id,omitempty"`
	// Datacenter name where the NSX Edge is deployed.
	DatacenterName string `json:"datacenterName,omitempty"`
	// Deployment state of the NSX Edge appliance. Values are 'deployed' when VMs have been deployed, 'undeployed' when no VMs are deployed and 'active' when Edge type is Distributed Logical Router and has no appliance deployed but is serving data path.
	State string `json:"state,omitempty"`
	ClientHandle string `json:"clientHandle,omitempty"`
	Scope *ScopeInfo `json:"scope,omitempty"`
	Type_ *ObjectType `json:"type,omitempty"`
	Revision int64 `json:"revision,omitempty"`
	VsmUuid string `json:"vsmUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ExtendedAttributes []ExtendedAttribute `json:"extendedAttributes,omitempty"`
	// Value is true if local egress is enabled for UDLR traffic. Applicable only for Universal Distributed Logical Router.
	LocalEgressEnabled bool `json:"localEgressEnabled,omitempty"`
	UniversalRevision int64 `json:"universalRevision,omitempty"`
	AllowedActions []string `json:"allowedActions,omitempty"`
	ObjectTypeName string `json:"objectTypeName,omitempty"`
	// Value is true if NSX Edge upgrade is available.
	IsUpgradeAvailable bool `json:"isUpgradeAvailable,omitempty"`
	IsUniversal bool `json:"isUniversal,omitempty"`
	Name string `json:"name,omitempty"`
	// Distributed Logical Router UUID provided by the NSX Controller.
	LrouterUuid string `json:"lrouterUuid,omitempty"`
	AppliancesSummary *AppliancesSummary `json:"appliancesSummary,omitempty"`
	// REST API version applicable for the NSX Edge.
	ApiVersion string `json:"apiVersion,omitempty"`
	// Tenant ID for the NSX Edge.
	TenantId string `json:"tenantId,omitempty"`
	// vCenter MOID of the datacenter where the NSX Edge is deployed.
	DatacenterMoid string `json:"datacenterMoid,omitempty"`
	// Number of connected vnics that are configured on the NSX Edge.
	NumberOfConnectedVnics int32 `json:"numberOfConnectedVnics,omitempty"`
}



// Status of each of the deployed NSX Edge appliances.
type EdgeVmStatus struct {
	// High Availability index of the appliance. Values are 0 and 1.
	Index int32 `json:"index,omitempty"`
	// High Availability state of the appliance. Values are active and standby.
	HaState string `json:"haState,omitempty"`
	// Name of the NSX Edge appliance.
	Name string `json:"name,omitempty"`
	// vCenter MOID of the NSX Edge appliance.
	Id string `json:"id,omitempty"`
	// NSX Edge appliance health status identified by GREY (unknown status), GREEN (health checks are successful), YELLOW (intermittent health check failure), RED (appliance not in serving state).
	EdgeVMStatus string `json:"edgeVMStatus,omitempty"`
	// Value of the last published pre rules generation number.
	PreRulesGenerationNumber int64 `json:"preRulesGenerationNumber,omitempty"`
}



// Address group configuration of the NSX Edge vnic. An interface can have one primary and multiple secondary IP addresses.
type EdgeVnicAddressGroup struct {
	// Subnet prefix length of the primary IP address.
	SubnetPrefixLength string `json:"subnetPrefixLength,omitempty"`
	SecondaryAddresses *SecondaryAddresses `json:"secondaryAddresses,omitempty"`
	// Primary IP address of the vnic interface. Required.
	PrimaryAddress string `json:"primaryAddress,omitempty"`
	SubnetMask string `json:"subnetMask,omitempty"`
}



// NSX Edge vnic address group configuration details.
type EdgeVnicAddressGroups struct {
	// Address group configuration of the NSX Edge vnic. Vnic can be configured to have more than one address group/subnets.
	AddressGroups []EdgeVnicAddressGroup `json:"addressGroups,omitempty"`
}



// Information of the x-eni created.
type EniInfo struct {
	// Subnet it belongs to.
	SubnetId string `json:"subnet_id,omitempty"`
	// Interface ID.
	Id string `json:"id,omitempty"`
	// Security Group of Eni.
	SecurityGroupId string `json:"security_group_id,omitempty"`
	// Private ip of eni.
	PrivateIp string `json:"private_ip,omitempty"`
	// Mac address of nic.
	MacAddress string `json:"mac_address,omitempty"`
}



// Decribes the capacity of a given entity.
type EntityCapacity struct {
	StorageCapacityGib int32 `json:"storage_capacity_gib,omitempty"`
	MemoryCapacityGib int32 `json:"memory_capacity_gib,omitempty"`
	TotalNumberOfCores int32 `json:"total_number_of_cores,omitempty"`
	NumberOfSsds int32 `json:"number_of_ssds,omitempty"`
	CpuCapacityGhz float64 `json:"cpu_capacity_ghz,omitempty"`
	NumberOfSockets int32 `json:"number_of_sockets,omitempty"`
}



type ErrorResponse struct {
	// HTTP status code
	Status int32 `json:"status"`
	// Originating request URI
	Path string `json:"path"`
	// If true, client should retry operation
	Retryable bool `json:"retryable"`
	// unique error code
	ErrorCode string `json:"error_code"`
	// localized error messages
	ErrorMessages []string `json:"error_messages"`
}



type EsxConfig struct {
	// Availability zone where the hosts should be provisioned. (Can be specified only for privileged host operations).
	AvailabilityZone string `json:"availability_zone,omitempty"`
	Esxs []string `json:"esxs,omitempty"`
	// An optional cluster id if the esxs operation has to be on a specific cluster.
	ClusterId string `json:"cluster_id,omitempty"`
	NumHosts int32 `json:"num_hosts"`
}



type EsxHost struct {
	Name string `json:"name,omitempty"`
	// Availability zone where the host is provisioned.
	AvailabilityZone string `json:"availability_zone,omitempty"`
	EsxId string `json:"esx_id,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Provider string `json:"provider"`
	// Backing cloud provider instance type for host.
	InstanceType string `json:"instance_type,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	CustomProperties map[string]string `json:"custom_properties,omitempty"`
	EsxState string `json:"esx_state,omitempty"`
}



type EsxHostInfo struct {
	// Backing cloud provider instance type for cluster.
	InstanceType string `json:"instance_type,omitempty"`
}



type ExtendedAttribute struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}



// List of features and their capability details based on NSX Edge appliance form factor.
type FeatureCapabilities struct {
	// Time stamp value at which the feature capabilities were retrieved.
	Timestamp int64 `json:"timestamp,omitempty"`
	// List of feature capability information.
	FeatureCapabilities []FeatureCapability `json:"featureCapabilities,omitempty"`
}



// Feature capability information.
type FeatureCapability struct {
	// List of key value pairs describing the feature configuration limits.
	ConfigurationLimits []KeyValueAttributes `json:"configurationLimits,omitempty"`
	// Value is true if feature is supported on NSX Edge.
	IsSupported bool `json:"isSupported,omitempty"`
	// Name of the feature or service.
	Service string `json:"service,omitempty"`
	Permission *LicenceAclPermissions `json:"permission,omitempty"`
}



// Individual feature status.
type FeatureStatus struct {
	// Status of the feature or service.
	Status string `json:"status,omitempty"`
	// Value is true if feature is configured.
	Configured bool `json:"configured,omitempty"`
	// Server status of the feature or service. Values are up and down.
	ServerStatus string `json:"serverStatus,omitempty"`
	// Publish status of the feature, whether APPLIED or PERSISTED.
	PublishStatus string `json:"publishStatus,omitempty"`
	// Name of the feature or service.
	Service string `json:"service,omitempty"`
}



// Firewall Configuration
type FirewallConfig struct {
	FirewallRules *FirewallRules `json:"firewallRules,omitempty"`
	FeatureType string `json:"featureType,omitempty"`
	// Version number tracking each configuration change. To avoid problems with overwriting changes, always retrieve and modify the latest configuration to include the current version number in your request. If you provide a version number which is not current, the request is rejected. If you omit the version number, the request is accepted but may overwrite any current changes if your change is not in sync with the latest change.
	Version int64 `json:"version,omitempty"`
	Template string `json:"template,omitempty"`
	GlobalConfig *FirewallGlobalConfig `json:"globalConfig,omitempty"`
	// Value is true if feature is enabled. Default value is true. Optional.
	Enabled bool `json:"enabled,omitempty"`
	DefaultPolicy *FirewallDefaultPolicy `json:"defaultPolicy,omitempty"`
}



// Dashboard Statistics data for Firewall.
type FirewallDashboardStats struct {
	// Number of NSX Edge firewall connections and rules.
	Connections []DashboardStat `json:"connections,omitempty"`
}



// Firewall default policy. Default is deny.
type FirewallDefaultPolicy struct {
	// Action. Default is deny. Supported values accept, deny
	Action string `json:"action,omitempty"`
	// Enable logging for the rule.
	LoggingEnabled bool `json:"loggingEnabled,omitempty"`
}



// Global configuration applicable to all rules.
type FirewallGlobalConfig struct {
	// Allow TCP out of window packets.
	TcpAllowOutOfWindowPackets bool `json:"tcpAllowOutOfWindowPackets,omitempty"`
	// UDP timeout close.
	UdpTimeout int32 `json:"udpTimeout,omitempty"`
	// IP generic timeout.
	IpGenericTimeout int32 `json:"ipGenericTimeout,omitempty"`
	// Pick TCP ongoing connections.
	TcpPickOngoingConnections bool `json:"tcpPickOngoingConnections,omitempty"`
	// TCP timeout open.
	TcpTimeoutOpen int32 `json:"tcpTimeoutOpen,omitempty"`
	// TCP timeout close.
	TcpTimeoutClose int32 `json:"tcpTimeoutClose,omitempty"`
	// ICMP6 timeout.
	Icmp6Timeout int32 `json:"icmp6Timeout,omitempty"`
	// Drop icmp replays.
	DropIcmpReplays bool `json:"dropIcmpReplays,omitempty"`
	// Log icmp errors.
	LogIcmpErrors bool `json:"logIcmpErrors,omitempty"`
	// Send TCP reset for closed NSX Edge ports.
	TcpSendResetForClosedVsePorts bool `json:"tcpSendResetForClosedVsePorts,omitempty"`
	// Drop invalid traffic.
	DropInvalidTraffic bool `json:"dropInvalidTraffic,omitempty"`
	// Protect against SYN flood attacks by detecting bogus TCP connections and terminating them without consuming firewall state tracking resources. Default : false
	EnableSynFloodProtection bool `json:"enableSynFloodProtection,omitempty"`
	// ICMP timeout.
	IcmpTimeout int32 `json:"icmpTimeout,omitempty"`
	// TCP timeout established.
	TcpTimeoutEstablished int32 `json:"tcpTimeoutEstablished,omitempty"`
	// Log invalid traffic.
	LogInvalidTraffic bool `json:"logInvalidTraffic,omitempty"`
}



type FirewallRule struct {
	RuleType string `json:"rule_type,omitempty"`
	ApplicationIds []string `json:"application_ids,omitempty"`
	Name string `json:"name,omitempty"`
	// Deprecated, left for backwards compatibility. Remove once UI stops using it.
	RuleInterface string `json:"rule_interface,omitempty"`
	// Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to 'any'.
	Destination string `json:"destination,omitempty"`
	Id string `json:"id,omitempty"`
	DestinationScope *FirewallRuleScope `json:"destination_scope,omitempty"`
	// Optional. Possible formats are IP, IP1-IPn, CIDR or comma separated list of those entries. If not specified, defaults to 'any'.
	Source string `json:"source,omitempty"`
	SourceScope *FirewallRuleScope `json:"source_scope,omitempty"`
	// list of protocols and ports for this firewall rule
	Services []FirewallService `json:"services,omitempty"`
	Action string `json:"action,omitempty"`
	// current revision of the list of firewall rules, used to protect against concurrent modification (first writer wins)
	Revision int32 `json:"revision,omitempty"`
}



// Optional for FirewallRule. If not specified, defaults to 'any'.
type FirewallRuleScope struct {
	GroupingObjectIds []string `json:"grouping_object_ids,omitempty"`
	VnicGroupIds []string `json:"vnic_group_ids,omitempty"`
}



// Statistics for firewall rule
type FirewallRuleStats struct {
	// Timestamp of statistics collection.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Connection count.
	ConnectionCount int64 `json:"connectionCount,omitempty"`
	// Byte count.
	ByteCount int64 `json:"byteCount,omitempty"`
	// Packet count.
	PacketCount int64 `json:"packetCount,omitempty"`
}



// Ordered list of firewall rules.
type FirewallRules struct {
	// Ordered list of firewall rules.
	FirewallRules []Nsxfirewallrule `json:"firewallRules,omitempty"`
}



type FirewallService struct {
	// protocol name, such as 'tcp', 'udp' etc.
	Protocol string `json:"protocol,omitempty"`
	// a list of port numbers and port ranges, such as {80, 91-95, 99}. If not specified, defaults to 'any'.
	Ports []string `json:"ports,omitempty"`
}



// Describes common properties for MGW and CGW configuration templates
type GatewayTemplate struct {
	PublicIp *SddcPublicIp `json:"public_ip,omitempty"`
	PrimaryDns string `json:"primary_dns,omitempty"`
	SecondaryDns string `json:"secondary_dns,omitempty"`
	FirewallRules []FirewallRule `json:"firewall_rules,omitempty"`
	Vpns []Vpn `json:"vpns,omitempty"`
}



// the GlcmBundle used for deploying the sddc
type GlcmBundle struct {
	// the glcmbundle's s3 bucket
	S3Bucket string `json:"s3Bucket,omitempty"`
	// the glcmbundle's id
	Id string `json:"id,omitempty"`
}



// DHCP lease information.
type HostLeaseInfo struct {
	// MAC address of the client.
	MacAddress string `json:"macAddress,omitempty"`
	// End time of the lease.
	Ends string `json:"ends,omitempty"`
	// Time stamp of when IP address was marked as abandoned.
	Abandoned string `json:"abandoned,omitempty"`
	// Client Last Transaction Time of the lease info.
	Cltt string `json:"cltt,omitempty"`
	// Name of the client.
	ClientHostname string `json:"clientHostname,omitempty"`
	// Start time of the lease.
	Starts string `json:"starts,omitempty"`
	// Lease's binding state.
	BindingState string `json:"bindingState,omitempty"`
	// The hardware type on which the lease will be used.
	HardwareType string `json:"hardwareType,omitempty"`
	// Time Sent From Partner of the lease info.
	Tsfp string `json:"tsfp,omitempty"`
	// Uid to identify the DHCP lease.
	Uid string `json:"uid,omitempty"`
	// Indicates what state the lease will move to when the current state expires.
	NextBindingState string `json:"nextBindingState,omitempty"`
	// IP address of the client.
	IpAddress string `json:"ipAddress,omitempty"`
	// Time Sent To Partner of the lease info.
	Tstp string `json:"tstp,omitempty"`
}



// Represents a structure for instance type config
type InstanceTypeConfig struct {
	// Instance type name.
	InstanceType string `json:"instance_type,omitempty"`
	// Array of number of hosts allowed for this operation. Range of hosts user can select for sddc provision
	Hosts []int32 `json:"hosts,omitempty"`
	// Display name of instance_type.
	DisplayName string `json:"display_name,omitempty"`
	EntityCapacity *EntityCapacity `json:"entity_capacity,omitempty"`
}



type InteractionPermissions struct {
	ManagePermission bool `json:"managePermission,omitempty"`
	ViewPermission bool `json:"viewPermission,omitempty"`
}



// Dashboard Statistics data for Interfaces.
type InterfacesDashboardStats struct {
	Vnic7InPkt []DashboardStat `json:"vnic_7_in_pkt,omitempty"`
	Vnic0InByte []DashboardStat `json:"vnic_0_in_byte,omitempty"`
	Vnic8OutPkt []DashboardStat `json:"vnic_8_out_pkt,omitempty"`
	Vnic5InByte []DashboardStat `json:"vnic_5_in_byte,omitempty"`
	Vnic2InPkt []DashboardStat `json:"vnic_2_in_pkt,omitempty"`
	Vnic3InPkt []DashboardStat `json:"vnic_3_in_pkt,omitempty"`
	Vnic6OutByte []DashboardStat `json:"vnic_6_out_byte,omitempty"`
	Vnic3InByte []DashboardStat `json:"vnic_3_in_byte,omitempty"`
	Vnic8InPkt []DashboardStat `json:"vnic_8_in_pkt,omitempty"`
	Vnic1InByte []DashboardStat `json:"vnic_1_in_byte,omitempty"`
	Vnic1OutPkt []DashboardStat `json:"vnic_1_out_pkt,omitempty"`
	Vnic5OutByte []DashboardStat `json:"vnic_5_out_byte,omitempty"`
	Vnic0OutPkt []DashboardStat `json:"vnic_0_out_pkt,omitempty"`
	Vnic0OutByte []DashboardStat `json:"vnic_0_out_byte,omitempty"`
	Vnic6OutPkt []DashboardStat `json:"vnic_6_out_pkt,omitempty"`
	Vnic3OutByte []DashboardStat `json:"vnic_3_out_byte,omitempty"`
	Vnic7InByte []DashboardStat `json:"vnic_7_in_byte,omitempty"`
	Vnic1OutByte []DashboardStat `json:"vnic_1_out_byte,omitempty"`
	Vnic9OutPkt []DashboardStat `json:"vnic_9_out_pkt,omitempty"`
	Vnic9InPkt []DashboardStat `json:"vnic_9_in_pkt,omitempty"`
	Vnic4InByte []DashboardStat `json:"vnic_4_in_byte,omitempty"`
	Vnic5OutPkt []DashboardStat `json:"vnic_5_out_pkt,omitempty"`
	Vnic2OutPkt []DashboardStat `json:"vnic_2_out_pkt,omitempty"`
	Vnic2InByte []DashboardStat `json:"vnic_2_in_byte,omitempty"`
	Vnic5InPkt []DashboardStat `json:"vnic_5_in_pkt,omitempty"`
	Vnic7OutPkt []DashboardStat `json:"vnic_7_out_pkt,omitempty"`
	Vnic3OutPkt []DashboardStat `json:"vnic_3_out_pkt,omitempty"`
	Vnic4OutPkt []DashboardStat `json:"vnic_4_out_pkt,omitempty"`
	Vnic4OutByte []DashboardStat `json:"vnic_4_out_byte,omitempty"`
	Vnic1InPkt []DashboardStat `json:"vnic_1_in_pkt,omitempty"`
	Vnic2OutByte []DashboardStat `json:"vnic_2_out_byte,omitempty"`
	Vnic6InByte []DashboardStat `json:"vnic_6_in_byte,omitempty"`
	Vnic0InPkt []DashboardStat `json:"vnic_0_in_pkt,omitempty"`
	Vnic9InByte []DashboardStat `json:"vnic_9_in_byte,omitempty"`
	Vnic7OutByte []DashboardStat `json:"vnic_7_out_byte,omitempty"`
	Vnic4InPkt []DashboardStat `json:"vnic_4_in_pkt,omitempty"`
	Vnic9OutByte []DashboardStat `json:"vnic_9_out_byte,omitempty"`
	Vnic8OutByte []DashboardStat `json:"vnic_8_out_byte,omitempty"`
	Vnic8InByte []DashboardStat `json:"vnic_8_in_byte,omitempty"`
	Vnic6InPkt []DashboardStat `json:"vnic_6_in_pkt,omitempty"`
}



// IP address
type IpAddresses struct {
	// List of IP addresses.
	IpAddress []string `json:"ipAddress,omitempty"`
}



// NSX Edge IPsec configuration details.
type Ipsec struct {
	FeatureType string `json:"featureType,omitempty"`
	Logging *Logging `json:"logging,omitempty"`
	Global *IpsecGlobalConfig `json:"global,omitempty"`
	// Value is true if feature is enabled. Default value is true. Optional.
	Enabled bool `json:"enabled,omitempty"`
	Sites *IpsecSites `json:"sites,omitempty"`
	// Enable/disable event generation on NSX Edge appliance for IPsec.
	DisableEvent bool `json:"disableEvent,omitempty"`
	// Version number tracking each configuration change. To avoid problems with overwriting changes, always retrieve and modify the latest configuration to include the current version number in your request. If you provide a version number which is not current, the request is rejected. If you omit the version number, the request is accepted but may overwrite any current changes if your change is not in sync with the latest change.
	Version int64 `json:"version,omitempty"`
	Template string `json:"template,omitempty"`
}



// Dashboard Statistics data for Ipsec.
type IpsecDashboardStats struct {
	// Tx transmitted bytes.
	IpsecBytesOut []DashboardStat `json:"ipsecBytesOut,omitempty"`
	// Rx received bytes.
	IpsecBytesIn []DashboardStat `json:"ipsecBytesIn,omitempty"`
	// Number of Ipsec tunnels.
	IpsecTunnels []DashboardStat `json:"ipsecTunnels,omitempty"`
}



// IPsec Global configuration details.
type IpsecGlobalConfig struct {
	// IPsec Global Pre Shared Key. Maximum characters is 128. Required when peerIp is configured as 'any' in NSX Edge IPsec Site configuration.
	Psk string `json:"psk,omitempty"`
	CaCertificates *CaCertificates `json:"caCertificates,omitempty"`
	// Certificate name or identifier. Required when x.509 is selected as the authentication mode.
	ServiceCertificate string `json:"serviceCertificate,omitempty"`
	CrlCertificates *CrlCertificates `json:"crlCertificates,omitempty"`
	Extension string `json:"extension,omitempty"`
}



// NSX Edge IPsec Site configuration details.
type IpsecSite struct {
	// Pre Shared Key for the IPsec Site. Required if Site peerIp is not 'any'. Global PSK is used when Authentication mode is PSK and Site peerIp is 'any'.
	Psk string `json:"psk,omitempty"`
	// Local ID of the IPsec Site. Defaults to the local IP.
	LocalId string `json:"localId,omitempty"`
	// Enable/disable Perfect Forward Secrecy. Default is true.
	EnablePfs bool `json:"enablePfs,omitempty"`
	// Authentication mode for the IPsec Site. Valid values are psk and x.509, with psk as default.
	AuthenticationMode string `json:"authenticationMode,omitempty"`
	PeerSubnets *Subnets `json:"peerSubnets,omitempty"`
	// Diffie-Hellman algorithm group. Defaults to DH14 for FIPS enabled NSX Edge. DH2 and DH5 are not supported when FIPS is enabled on NSX Edge. Valid values are DH2, DH5, DH14, DH15, DH16.
	DhGroup string `json:"dhGroup,omitempty"`
	// ID of the IPsec Site configuration provided by NSX Manager.
	SiteId string `json:"siteId,omitempty"`
	// Description of the IPsec Site.
	Description string `json:"description,omitempty"`
	// IP (IPv4) address or FQDN of the Peer. Can also be specified as 'any'. Required.
	PeerIp string `json:"peerIp,omitempty"`
	// Name of the IPsec Site.
	Name string `json:"name,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	// Local IP of the IPsec Site. Should be one of the IP addresses configured on the uplink interfaces of the NSX Edge. Required.
	LocalIp string `json:"localIp,omitempty"`
	// IPsec encryption algorithm with default as aes256. Valid values are 'aes', 'aes256', '3des', 'aes-gcm'.
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`
	// Enable/disable IPsec Site.
	Enabled bool `json:"enabled,omitempty"`
	// MTU for the IPsec site. Defaults to the mtu of the NSX Edge vnic specified by the localIp. Optional.
	Mtu int32 `json:"mtu,omitempty"`
	Extension string `json:"extension,omitempty"`
	// Peer ID. Should be unique for all IPsec Site's configured for an NSX Edge.
	PeerId string `json:"peerId,omitempty"`
	LocalSubnets *Subnets `json:"localSubnets,omitempty"`
}



type IpsecSiteIkeStatus struct {
	ChannelStatus string `json:"channelStatus,omitempty"`
	ChannelState string `json:"channelState,omitempty"`
	PeerIpAddress string `json:"peerIpAddress,omitempty"`
	LocalIpAddress string `json:"localIpAddress,omitempty"`
	PeerSubnets []string `json:"peerSubnets,omitempty"`
	PeerId string `json:"peerId,omitempty"`
	LastInformationalMessage string `json:"lastInformationalMessage,omitempty"`
	LocalSubnets []string `json:"localSubnets,omitempty"`
}



type IpsecSiteStats struct {
	RxBytesOnSite int32 `json:"rxBytesOnSite,omitempty"`
	TunnelStats []IpsecTunnelStats `json:"tunnelStats,omitempty"`
	IkeStatus *IpsecSiteIkeStatus `json:"ikeStatus,omitempty"`
	SiteStatus string `json:"siteStatus,omitempty"`
	TxBytesFromSite int32 `json:"txBytesFromSite,omitempty"`
}



// List of IPsec sites for NSX Edge.
type IpsecSites struct {
	Sites []IpsecSite `json:"sites,omitempty"`
}



type IpsecStatusAndStats struct {
	TimeStamp int64 `json:"timeStamp,omitempty"`
	ServerStatus string `json:"serverStatus,omitempty"`
	SiteStatistics []IpsecSiteStats `json:"siteStatistics,omitempty"`
}



type IpsecTunnelStats struct {
	TunnelStatus string `json:"tunnelStatus,omitempty"`
	PeerSPI string `json:"peerSPI,omitempty"`
	RxBytesOnLocalSubnet int32 `json:"rxBytesOnLocalSubnet,omitempty"`
	EstablishedDate string `json:"establishedDate,omitempty"`
	PeerSubnet string `json:"peerSubnet,omitempty"`
	AuthenticationAlgorithm string `json:"authenticationAlgorithm,omitempty"`
	TunnelState string `json:"tunnelState,omitempty"`
	TxBytesFromLocalSubnet int32 `json:"txBytesFromLocalSubnet,omitempty"`
	LastInformationalMessage string `json:"lastInformationalMessage,omitempty"`
	LocalSPI string `json:"localSPI,omitempty"`
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`
	LocalSubnet string `json:"localSubnet,omitempty"`
}



// Key value pair describing the feature configuration limit.
type KeyValueAttributes struct {
	// Value corresponding to the key of the configuration limit parameter.
	Value string `json:"value,omitempty"`
	// Key name of the configuration limit parameter.
	Key string `json:"key,omitempty"`
}



type KmsVpcEndpoint struct {
	// The identifier of the VPC endpoint created to AWS Key Management Service
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`
	NetworkInterfaceIds []string `json:"network_interface_ids,omitempty"`
}



// Layer 2 extension.
type L2Extension struct {
	// Identifier for layer 2 extension tunnel. Valid range: 1-4093.
	TunnelId int32 `json:"tunnelId"`
}



type L2Vpn struct {
	// Enable (true) or disable (false) L2 VPN.
	Enabled bool `json:"enabled,omitempty"`
	// Array of L2 vpn site config.
	Sites []Site `json:"sites"`
	// Public uplink ip address. IP of external interface on which L2VPN service listens to.
	ListenerIp string `json:"listener_ip,omitempty"`
}



// L2 VPN status and statistics of a single L2 VPN site.
type L2vpnStats struct {
	// Status of the tunnel (UP/DOWN).
	TunnelStatus string `json:"tunnelStatus,omitempty"`
	// Tunnel established date.
	EstablishedDate int64 `json:"establishedDate,omitempty"`
	// User defined name of the site.
	Name string `json:"name,omitempty"`
	// Number of received packets dropped.
	DroppedRxPackets int32 `json:"droppedRxPackets,omitempty"`
	// Cipher used in encryption.
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`
	// Reason for the tunnel down.
	FailureMessage string `json:"failureMessage,omitempty"`
	// Number of bytes transferred from local subnet.
	TxBytesFromLocalSubnet int32 `json:"txBytesFromLocalSubnet,omitempty"`
	// Number of bytes received on the local subnet.
	RxBytesOnLocalSubnet int32 `json:"rxBytesOnLocalSubnet,omitempty"`
	// Number of transferred packets dropped.
	DroppedTxPackets int32 `json:"droppedTxPackets,omitempty"`
	// Time stamp of the statistics collection.
	LastUpdatedTime int64 `json:"lastUpdatedTime,omitempty"`
}



// L2 VPN status and statistics.
type L2vpnStatusAndStats struct {
	// Time stamp of statistics collection.
	TimeStamp int64 `json:"timeStamp,omitempty"`
	ServerStatus string `json:"serverStatus,omitempty"`
	// List of statistics for each Site.
	SiteStats []L2vpnStats `json:"siteStats,omitempty"`
}



// Licence and access control information for the feature.
type LicenceAclPermissions struct {
	DataPermission *DataPermissions `json:"dataPermission,omitempty"`
	// Value is true if feature is licenced.
	IsLicensed bool `json:"isLicensed,omitempty"`
	AccessPermission *InteractionPermissions `json:"accessPermission,omitempty"`
}



// Dashboard Statistics data for Load Balancer.
type LoadBalancerDashboardStats struct {
	// Number of bytes in.
	LbBpsIn []DashboardStat `json:"lbBpsIn,omitempty"`
	// Number of HTTP requests received by Load Balancer.
	LbHttpReqs []DashboardStat `json:"lbHttpReqs,omitempty"`
	// Number of bytes out.
	LbBpsOut []DashboardStat `json:"lbBpsOut,omitempty"`
	// Number of Load Balancer sessions.
	LbSessions []DashboardStat `json:"lbSessions,omitempty"`
}



// logging.
type Logging struct {
	// Log level. Valid values: emergency, alert, critical, error, warning, notice, info, debug.
	LogLevel string `json:"logLevel,omitempty"`
	// Logging enabled.
	Enable bool `json:"enable,omitempty"`
}



type LogicalNetwork struct {
	// the subnet cidr
	SubnetCidr string `json:"subnet_cidr,omitempty"`
	// name of the network
	Name string `json:"name,omitempty"`
	// gateway ip of the logical network
	GatewayIp string `json:"gatewayIp,omitempty"`
	// if 'true' - enabled; if 'false' - disabled
	DhcpEnabled string `json:"dhcp_enabled,omitempty"`
	// ip range within the subnet mask, range delimiter is '-' (example 10.118.10.130-10.118.10.140)
	DhcpIpRange string `json:"dhcp_ip_range,omitempty"`
	// tunnel id of extended network
	TunnelId int32 `json:"tunnel_id,omitempty"`
	Id string `json:"id,omitempty"`
	NetworkType string `json:"network_type,omitempty"`
}



type LogicalRouterScope struct {
	Type_ string `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
}



type LogicalRouterScopes struct {
	LogicalRouterScope []LogicalRouterScope `json:"logicalRouterScope,omitempty"`
}


type MacAddress struct {
	EdgeVmHaIndex int32  `json:"edgeVmHaIndex,omitempty"`
	Value         string `json:"value,omitempty"`
}



type MaintenanceWindow struct {
	DayOfWeek string `json:"day_of_week,omitempty"`
	HourOfDay int32 `json:"hour_of_day,omitempty"`
}



type MaintenanceWindowEntry struct {
	// true if the SDDC is in the defined Mainentance Window
	InMaintenanceWindow bool `json:"in_maintenance_window,omitempty"`
	ReservationSchedule *ReservationSchedule `json:"reservation_schedule,omitempty"`
	// ID for reservation
	ReservationId string `json:"reservation_id,omitempty"`
	// true if the SDDC is currently undergoing maintenance
	InMaintenanceMode bool `json:"in_maintenance_mode,omitempty"`
	// SDDC ID for this reservation
	SddcId string `json:"sddc_id,omitempty"`
}



type MaintenanceWindowGet struct {
	DayOfWeek string `json:"day_of_week,omitempty"`
	HourOfDay int32 `json:"hour_of_day,omitempty"`
	DurationMin int64 `json:"duration_min,omitempty"`
	Version int64 `json:"version,omitempty"`
}



type ManagementGatewayTemplate struct {
	PublicIp *SddcPublicIp `json:"public_ip,omitempty"`
	PrimaryDns string `json:"primary_dns,omitempty"`
	SecondaryDns string `json:"secondary_dns,omitempty"`
	FirewallRules []FirewallRule `json:"firewall_rules,omitempty"`
	Vpns []Vpn `json:"vpns,omitempty"`
	// mgw network subnet cidr
	SubnetCidr string `json:"subnet_cidr,omitempty"`
}



type MapZonesRequest struct {
	// The connected account ID to remap. This is a standard UUID.
	ConnectedAccountId string `json:"connected_account_id,omitempty"`
	// The org ID to remap in. This is a standard UUID.
	OrgId string `json:"org_id,omitempty"`
	// A list of Petronas regions to map.
	PetronasRegionsToMap []string `json:"petronas_regions_to_map,omitempty"`
}



// Start time, end time and interval details.
type MetaDashboardStats struct {
	// Statistics data is collected for these vNICs.
	Vnics []Vnic `json:"vnics,omitempty"`
	// End time in seconds.
	EndTime int64 `json:"endTime,omitempty"`
	// Start time in seconds.
	StartTime int64 `json:"startTime,omitempty"`
	// Time interval in seconds.
	Interval int32 `json:"interval,omitempty"`
}



// metadata of the sddc manifest
type Metadata struct {
	// the timestamp for the bundle
	Timestamp string `json:"timestamp,omitempty"`
	// the cycle id
	CycleId string `json:"cycle_id,omitempty"`
}



// NAT configuration
type Nat struct {
	Rules *NatRules `json:"rules,omitempty"`
	FeatureType string `json:"featureType,omitempty"`
	// Version number tracking each configuration change. To avoid problems with overwriting changes, always retrieve and modify the latest configuration to include the current version number in your request. If you provide a version number which is not current, the request is rejected. If you omit the version number, the request is accepted but may overwrite any current changes if your change is not in sync with the latest change.
	Version int64 `json:"version,omitempty"`
	// Value is true if feature is enabled. Default value is true. Optional.
	Enabled bool `json:"enabled,omitempty"`
	Template string `json:"template,omitempty"`
}



type NatRule struct {
	RuleType string `json:"rule_type,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Name string `json:"name,omitempty"`
	InternalPorts string `json:"internal_ports,omitempty"`
	PublicPorts string `json:"public_ports,omitempty"`
	PublicIp string `json:"public_ip,omitempty"`
	InternalIp string `json:"internal_ip,omitempty"`
	Action string `json:"action,omitempty"`
	Id string `json:"id,omitempty"`
	// current revision of the list of nat rules, used to protect against concurrent modification (first writer wins)
	Revision int32 `json:"revision,omitempty"`
}



// Ordered list of NAT rules.
type NatRules struct {
	// Ordered list of NAT rules.
	NatRulesDtos []Nsxnatrule `json:"natRulesDtos,omitempty"`
}



type NetworkTemplate struct {
	ManagementGatewayTemplates []ManagementGatewayTemplate `json:"management_gateway_templates,omitempty"`
	ComputeGatewayTemplates []ComputeGatewayTemplate `json:"compute_gateway_templates,omitempty"`
}



type NewCredentials struct {
	// Username of the credentials
	Username string `json:"username"`
	// Password associated with the credentials
	Password string `json:"password"`
	// Name of the credentials
	Name string `json:"name"`
}



// Firewall Rule
type Nsxfirewallrule struct {
	// Identifies the type of the rule. internal_high or user.
	RuleType string `json:"ruleType,omitempty"`
	// Description for the rule
	Description string `json:"description,omitempty"`
	// Identifier for the rule.
	RuleId int64 `json:"ruleId,omitempty"`
	// Defines the order of NAT and Firewall pipeline. When false, firewall happens before NAT. Default : false
	MatchTranslated bool `json:"matchTranslated,omitempty"`
	InvalidApplication bool `json:"invalidApplication,omitempty"`
	// Direction. Possible values in or out. Default is 'any'.
	Direction string `json:"direction,omitempty"`
	Statistics *FirewallRuleStats `json:"statistics,omitempty"`
	// Name for the rule.
	Name string `json:"name,omitempty"`
	InvalidSource bool `json:"invalidSource,omitempty"`
	// Enable logging for the rule.
	LoggingEnabled bool `json:"loggingEnabled,omitempty"`
	Destination *AddressFwSourceDestination `json:"destination,omitempty"`
	// Enable rule.
	Enabled bool `json:"enabled,omitempty"`
	Application *Application `json:"application,omitempty"`
	Source *AddressFwSourceDestination `json:"source,omitempty"`
	// Action. Values : accept, deny
	Action string `json:"action,omitempty"`
	InvalidDestination bool `json:"invalidDestination,omitempty"`
	// Rule tag. Used to specify user-defined ruleId. If not specified NSX Manager will generate ruleId.
	RuleTag int64 `json:"ruleTag,omitempty"`
}



// Application (service) for firewall rule.
type Nsxfirewallservice struct {
	// List of source ports.
	SourcePort []string `json:"sourcePort,omitempty"`
	// Protocol.
	Protocol string `json:"protocol,omitempty"`
	// List of destination ports.
	Port []string `json:"port,omitempty"`
	// IcmpType. Only supported when protocol is icmp. Default is 'any'.
	IcmpType string `json:"icmpType,omitempty"`
}



// L2 VPN server configuration.
type Nsxl2vpn struct {
	// Listener IP addresses.
	ListenerIps []string `json:"listenerIps"`
	// Enabled state of L2 VPN service.
	Enabled bool `json:"enabled,omitempty"`
	Sites *Sites `json:"sites"`
}



// NAT rule
type Nsxnatrule struct {
	// Interface on which the NAT rule is applied.
	Vnic string `json:"vnic,omitempty"`
	// Identifies the type of the rule. internal_high or user.
	RuleType string `json:"ruleType,omitempty"`
	// Protocol. Default is 'any'
	Protocol string `json:"protocol,omitempty"`
	// Description for the rule.
	Description string `json:"description,omitempty"`
	// Identifier for the rule.
	RuleId int64 `json:"ruleId,omitempty"`
	// Apply SNAT rule only if traffic has this destination port. Default is 'any'.
	SnatMatchDestinationPort string `json:"snatMatchDestinationPort,omitempty"`
	// Original address or address range. This is the original source address for SNAT rules and the original destination address for DNAT rules.
	OriginalAddress string `json:"originalAddress,omitempty"`
	// Apply DNAT rule only if traffic has this source address. Default is 'any'.
	DnatMatchSourceAddress string `json:"dnatMatchSourceAddress,omitempty"`
	// Apply DNAT rule only if traffic has this source port. Default is 'any'.
	DnatMatchSourcePort string `json:"dnatMatchSourcePort,omitempty"`
	// Apply SNAT rule only if traffic has this destination address. Default is 'any'.
	SnatMatchDestinationAddress string `json:"snatMatchDestinationAddress,omitempty"`
	// Original port. This is the original source port for SNAT rules, and the original destination port for DNAT rules.
	OriginalPort string `json:"originalPort,omitempty"`
	// Enable logging for the rule.
	LoggingEnabled bool `json:"loggingEnabled,omitempty"`
	// Translated address or address range.
	TranslatedAddress string `json:"translatedAddress,omitempty"`
	// Enable rule.
	Enabled bool `json:"enabled,omitempty"`
	// ICMP type. Only supported when protocol is icmp. Default is 'any'.
	IcmpType string `json:"icmpType,omitempty"`
	// Translated port. Supported in DNAT rules only.
	TranslatedPort string `json:"translatedPort,omitempty"`
	// Action for the rule. SNAT or DNAT.
	Action string `json:"action,omitempty"`
	// Rule tag. Used to specify user-defined ruleId. If not specified NSX Manager will generate ruleId.
	RuleTag int64 `json:"ruleTag,omitempty"`
}



// L2 VPN site.
type Nsxsite struct {
	// Secure L2VPN traffic.
	SecureTraffic bool `json:"secureTraffic,omitempty"`
	// Identifier for L2 VPN site.
	SiteId string `json:"siteId,omitempty"`
	// Name of L2 VPN site. Length: 1-255 characters.
	Name string `json:"name,omitempty"`
	// Password for L2 VPN user. Passwords must contain the following: 12-63 characters, a mix of upper case letters, lower case letters, numbers, and at least one special character. Password must not contain the username as a substring. Do not repeat a character 3 or more times.
	Password string `json:"password,omitempty"`
	// L2 VPN user ID. Valid user names: 1-63 characters, letters and numbers only. No white space or special characters.
	UserId string `json:"userId,omitempty"`
	// Description of L2 VPN site.
	Description string `json:"description,omitempty"`
}



// Details the state of different NSX add-ons.
type NsxtAddons struct {
	// Indicates whether NSX Advanced addon is enabled or disabled.
	EnableNsxAdvancedAddon bool `json:"enable_nsx_advanced_addon,omitempty"`
}



type ObjectType struct {
	Name string `json:"name,omitempty"`
}



// Holder for the offer instances.
type OfferInstancesHolder struct {
	OnDemand *OnDemandOfferInstance `json:"on_demand"`
	Offers []TermOfferInstance `json:"offers"`
}



// Holder for the on-demand offer instance.
type OnDemandOfferInstance struct {
	Product string `json:"product"`
	// Deprecated. Please use product and type fields instead.
	ProductType string `json:"product_type,omitempty"`
	Name string `json:"name"`
	Currency string `json:"currency"`
	Region string `json:"region"`
	UnitPrice string `json:"unit_price"`
	MonthlyCost string `json:"monthly_cost"`
	Version string `json:"version"`
	Type_ string `json:"type"`
	Description string `json:"description"`
}



type OrgProperties struct {
	// A map of string properties to values.
	Values map[string]string `json:"values,omitempty"`
}


import (
	"time"
)

type Organization struct {
	Updated time.Time `json:"updated"`
	// User id that last updated this record
	UserId string `json:"user_id"`
	// User id that last updated this record
	UpdatedByUserId string `json:"updated_by_user_id"`
	Created time.Time `json:"created"`
	// Version of this entity
	Version int32 `json:"version"`
	// User name that last updated this record
	UpdatedByUserName string `json:"updated_by_user_name,omitempty"`
	// User name that last updated this record
	UserName string `json:"user_name"`
	// Unique ID for this entity
	Id string `json:"id"`
	// ORG_TYPE to be associated with the org
	OrgType string `json:"org_type,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Name string `json:"name,omitempty"`
	ProjectState string `json:"project_state,omitempty"`
	Properties *OrgProperties `json:"properties,omitempty"`
}



// NSX Edges listed by pages.
type PagedEdgeList struct {
	EdgePage *DataPageEdgeSummary `json:"edgePage,omitempty"`
}



type PagingInfo struct {
	SortOrderAscending bool `json:"sortOrderAscending,omitempty"`
	TotalCount int64 `json:"totalCount,omitempty"`
	StartIndex int32 `json:"startIndex,omitempty"`
	SortBy string `json:"sortBy,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
}



type PaymentMethodInfo struct {
	Type_ string `json:"type,omitempty"`
	DefaultFlag bool `json:"default_flag,omitempty"`
	PaymentMethodId string `json:"payment_method_id,omitempty"`
}



type PopAgentXeniConnection struct {
	// The gateway route ip fo the subnet.
	DefaultSubnetRoute string `json:"default_subnet_route,omitempty"`
	EniInfo *EniInfo `json:"eni_info,omitempty"`
}



type PopAmiInfo struct {
	// instance type of the esx ami
	InstanceType string `json:"instance_type,omitempty"`
	// the region of the esx ami
	Region string `json:"region,omitempty"`
	// the ami id for the esx
	Id string `json:"id,omitempty"`
	// the name of the esx ami
	Name string `json:"name,omitempty"`
	// PoP AMI type. CENTOS: a Centos AMI; POP: a PoP AMI.
	Type_ string `json:"type,omitempty"`
}


import (
	"time"
)

// Present a SDDC PoP information.
type PopInfo struct {
	// A map of [region name of PoP / PoP-AMI]:[PopAmiInfo].
	AmiInfos map[string]PopAmiInfo `json:"ami_infos"`
	// The PopInfo (or PoP AMI) created time. Using ISO 8601 date-time pattern.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// A map of [service type]:[PopServiceInfo]
	ServiceInfos map[string]PopServiceInfo `json:"service_infos,omitempty"`
	// UUID of the PopInfo
	Id string `json:"id,omitempty"`
	// version of the manifest.
	ManifestVersion string `json:"manifest_version,omitempty"`
}



type PopServiceInfo struct {
	// The service change set number.
	Cln string `json:"cln,omitempty"`
	// The service API version.
	Version string `json:"version,omitempty"`
	// The service build number.
	Build string `json:"build,omitempty"`
	// An enum of PoP related services (including os platform and JRE).
	Service string `json:"service,omitempty"`
}



// Represents a provisioning spec for a sddc
type ProvisionSpec struct {
	// Map of provider to sddc config spec
	Provider map[string]SddcConfigSpec `json:"provider,omitempty"`
}



// DNS request statistics.
type Requests struct {
	Total int64 `json:"total,omitempty"`
	Queries int64 `json:"queries,omitempty"`
}


import (
	"time"
)

type Reservation struct {
	// Duration - required for reservation in maintenance window
	Duration int64 `json:"duration,omitempty"`
	// Reservation ID
	Rid string `json:"rid,omitempty"`
	// Optional
	CreateTime string `json:"create_time,omitempty"`
	// Start time of a reservation
	StartTime time.Time `json:"start_time,omitempty"`
	// Optional
	Metadata map[string]string `json:"metadata,omitempty"`
}


import (
	"time"
)

type ReservationInMw struct {
	// Reservation ID
	Rid string `json:"rid,omitempty"`
	// SUNDAY of the week that maintenance is scheduled, ISO format date
	WeekOf string `json:"week_of,omitempty"`
	// Optional
	CreateTime time.Time `json:"create_time,omitempty"`
	// Optional
	Metadata map[string]string `json:"metadata,omitempty"`
}



type ReservationSchedule struct {
	DurationMin int64 `json:"duration_min,omitempty"`
	Version int64 `json:"version,omitempty"`
	Reservations []Reservation `json:"reservations,omitempty"`
	ReservationsMw []ReservationInMw `json:"reservations_mw,omitempty"`
}



type ReservationWindow struct {
	ReservationState string `json:"reservation_state,omitempty"`
	Emergency bool `json:"emergency,omitempty"`
	MaintenanceProperties *ReservationWindowMaintenanceProperties `json:"maintenance_properties,omitempty"`
	ReserveId string `json:"reserve_id,omitempty"`
	StartHour int32 `json:"start_hour,omitempty"`
	SddcId string `json:"sddc_id,omitempty"`
	ManifestId string `json:"manifest_id,omitempty"`
	DurationHours int64 `json:"duration_hours,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	// Metadata for reservation window, in key-value form
	Metadata map[string]string `json:"metadata,omitempty"`
}



type ReservationWindowMaintenanceProperties struct {
	// Status of upgrade, if any
	Status string `json:"status,omitempty"`
}



// Job result information for the configuration change carried out on NSX Edge.
type Result struct {
	// Job Result value associated with key ID.
	Value string `json:"value,omitempty"`
	// Job Result key ID.
	Key string `json:"key,omitempty"`
}



type RouteTableInfo struct {
	// route table name
	Name string `json:"name,omitempty"`
	// route table id
	Id string `json:"id,omitempty"`
}



type ScopeInfo struct {
	ObjectTypeName string `json:"objectTypeName,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}


import (
	"time"
)

type Sddc struct {
	Updated time.Time `json:"updated"`
	// User id that last updated this record
	UserId string `json:"user_id"`
	// User id that last updated this record
	UpdatedByUserId string `json:"updated_by_user_id"`
	Created time.Time `json:"created"`
	// Version of this entity
	Version int32 `json:"version"`
	// User name that last updated this record
	UpdatedByUserName string `json:"updated_by_user_name,omitempty"`
	// User name that last updated this record
	UserName string `json:"user_name"`
	// Unique ID for this entity
	Id string `json:"id"`
	// name for SDDC
	Name string `json:"name,omitempty"`
	SddcState string `json:"sddc_state,omitempty"`
	// Expiration date of a sddc in UTC (will be set if its applicable)
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// Type of the sddc
	SddcType string `json:"sddc_type,omitempty"`
	// Whether this sddc is reduced capacity 1NODE
	OneNodeReducedCapacity bool `json:"one_node_reduced_capacity,omitempty"`
	Provider string `json:"provider,omitempty"`
	// Account linking state of the sddc
	AccountLinkState string `json:"account_link_state,omitempty"`
	// Describes the access state of sddc, valid state is DISABLED or ENABLED
	SddcAccessState string `json:"sddc_access_state,omitempty"`
	ResourceConfig *AwsSddcResourceConfig `json:"resource_config,omitempty"`
}



type SddcAllocatePublicIpSpec struct {
	Count int32 `json:"count"`
	// List of workload VM private IPs to be assigned the public IP just allocated.
	PrivateIps []string `json:"private_ips,omitempty"`
	// List of names for the workload VM public IP assignment.
	Names []string `json:"names,omitempty"`
}



type SddcConfig struct {
	// Whether this sddc is reduced capacity 1NODE.
	OneNodeReducedCapacity bool `json:"one_node_reduced_capacity,omitempty"`
	// AWS VPC IP range. Only prefix of 16 or 20 is currently supported.
	VpcCidr string `json:"vpc_cidr,omitempty"`
	HostInstanceType string `json:"host_instance_type,omitempty"`
	// skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan bool `json:"skip_creating_vxlan,omitempty"`
	// VXLAN IP subnet in CIDR for compute gateway
	VxlanSubnet string `json:"vxlan_subnet,omitempty"`
	// The size of the vCenter and NSX appliances. \"large\" sddcSize corresponds to a 'large' vCenter appliance and 'large' NSX appliance. 'medium' sddcSize corresponds to 'medium' vCenter appliance and 'medium' NSX appliance. Value defaults to 'medium'. 
	Size string `json:"size,omitempty"`
	// The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of seperable storage will be used. 
	StorageCapacity int64 `json:"storage_capacity,omitempty"`
	Name string `json:"name"`
	// A list of the SDDC linking configurations to use.
	AccountLinkSddcConfig []AccountLinkSddcConfig `json:"account_link_sddc_config,omitempty"`
	// If provided, will be assigned as SDDC id of the provisioned SDDC.
	SddcId string `json:"sddc_id,omitempty"`
	NumHosts int32 `json:"num_hosts"`
	// Denotes the sddc type , if the value is null or empty, the type is considered as default.
	SddcType string `json:"sddc_type,omitempty"`
	AccountLinkConfig *AccountLinkConfig `json:"account_link_config,omitempty"`
	// Determines what additional properties are available based on cloud provider.
	Provider string `json:"provider"`
	// The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.
	SsoDomain string `json:"sso_domain,omitempty"`
	// If provided, configuration from the template will applied to the provisioned SDDC.
	SddcTemplateId string `json:"sddc_template_id,omitempty"`
	// Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ.
	DeploymentType string `json:"deployment_type,omitempty"`
}



// Represents a configuration spec for a sddc
type SddcConfigSpec struct {
	// Map of sddc type to config spec
	SddcTypeConfigSpec map[string]ConfigSpec `json:"sddc_type_config_spec,omitempty"`
	// The region name to display names mapping
	RegionDisplayNames map[string]string `json:"region_display_names,omitempty"`
}



type SddcId struct {
	// Sddc ID
	SddcId string `json:"sddc_id,omitempty"`
}



type SddcLinkConfig struct {
	CustomerSubnetIds []string `json:"customer_subnet_ids,omitempty"`
	// Determines which connected customer account to link to
	ConnectedAccountId string `json:"connected_account_id,omitempty"`
}



// Describes software components of the installed SDDC
type SddcManifest struct {
	// the vmcVersion of the sddc for display
	VmcVersion string `json:"vmc_version,omitempty"`
	GlcmBundle *GlcmBundle `json:"glcm_bundle,omitempty"`
	PopInfo *PopInfo `json:"pop_info,omitempty"`
	// the vmcInternalVersion of the sddc for internal use
	VmcInternalVersion string `json:"vmc_internal_version,omitempty"`
	EbsBackedVsanConfig *EbsBackedVsanConfig `json:"ebs_backed_vsan_config,omitempty"`
	VsanWitnessAmi *AmiInfo `json:"vsan_witness_ami,omitempty"`
	EsxAmi *AmiInfo `json:"esx_ami,omitempty"`
	EsxNsxtAmi *AmiInfo `json:"esx_nsxt_ami,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
}



// Logical network.
type SddcNetwork struct {
	Subnets *SddcNetworkAddressGroups `json:"subnets,omitempty"`
	// Name of the compute gateway to which the logical network is attached.
	CgwName string `json:"cgwName,omitempty"`
	// Name of logical network. Length needs to be between 1-35 characters.
	Name string `json:"name"`
	L2Extension *L2Extension `json:"l2Extension,omitempty"`
	// ID of the compute gateway edge to which the logical network is attached.
	CgwId string `json:"cgwId"`
	DhcpConfigs *SddcNetworkDhcpConfig `json:"dhcpConfigs,omitempty"`
	// ID of logical network.
	Id string `json:"id,omitempty"`
}



// Logical Network address group.
type SddcNetworkAddressGroup struct {
	// Prefix length of logical network.
	PrefixLength string `json:"prefixLength,omitempty"`
	// Primary address for logical network.
	PrimaryAddress string `json:"primaryAddress,omitempty"`
}



// Logical network address groups.
type SddcNetworkAddressGroups struct {
	// List of logical network address groups.
	AddressGroups []SddcNetworkAddressGroup `json:"addressGroups,omitempty"`
}



// DHCP configuration for the logical network.
type SddcNetworkDhcpConfig struct {
	// List of IP pools in DHCP configuration.
	IpPools []SddcNetworkDhcpIpPool `json:"ipPools,omitempty"`
}



// DHCP IP pool for logical network.
type SddcNetworkDhcpIpPool struct {
	// IP range for DHCP IP pool.
	IpRange string `json:"ipRange,omitempty"`
	// DNS domain name.
	DomainName string `json:"domainName,omitempty"`
}



// Patch request body for SDDC
type SddcPatchRequest struct {
	// The new name of the SDDC to be changed to.
	Name string `json:"name,omitempty"`
}



type SddcPublicIp struct {
	PublicIp string `json:"public_ip"`
	Name string `json:"name,omitempty"`
	AllocationId string `json:"allocation_id,omitempty"`
	DnatRuleId string `json:"dnat_rule_id,omitempty"`
	AssociatedPrivateIp string `json:"associated_private_ip,omitempty"`
	SnatRuleId string `json:"snat_rule_id,omitempty"`
}



type SddcResourceConfig struct {
	// Name for management appliance network.
	MgmtApplianceNetworkName string `json:"mgmt_appliance_network_name,omitempty"`
	// if true, NSX-T UI is enabled.
	Nsxt bool `json:"nsxt,omitempty"`
	// Management Gateway Id
	MgwId string `json:"mgw_id,omitempty"`
	// URL of the NSX Manager
	NsxMgrUrl string `json:"nsx_mgr_url,omitempty"`
	// PSC internal management IP
	PscManagementIp string `json:"psc_management_ip,omitempty"`
	// URL of the PSC server
	PscUrl string `json:"psc_url,omitempty"`
	Cgws []string `json:"cgws,omitempty"`
	// Availability zones over which esx hosts are provisioned. MultiAZ SDDCs will have hosts provisioned over two availability zones while SingleAZ SDDCs will provision over one.
	AvailabilityZones []string `json:"availability_zones,omitempty"`
	// The ManagedObjectReference of the management Datastore
	ManagementDs string `json:"management_ds,omitempty"`
	// nsx api entire base url
	NsxApiPublicEndpointUrl string `json:"nsx_api_public_endpoint_url,omitempty"`
	CustomProperties map[string]string `json:"custom_properties,omitempty"`
	// Password for vCenter SDDC administrator
	CloudPassword string `json:"cloud_password,omitempty"`
	// Discriminator for additional properties
	Provider string `json:"provider"`
	// List of clusters in the SDDC.
	Clusters []Cluster `json:"clusters,omitempty"`
	// vCenter internal management IP
	VcManagementIp string `json:"vc_management_ip,omitempty"`
	SddcNetworks []string `json:"sddc_networks,omitempty"`
	// Username for vCenter SDDC administrator
	CloudUsername string `json:"cloud_username,omitempty"`
	EsxHosts []AwsEsxHost `json:"esx_hosts,omitempty"`
	// NSX Manager internal management IP
	NsxMgrManagementIp string `json:"nsx_mgr_management_ip,omitempty"`
	// unique id of the vCenter server
	VcInstanceId string `json:"vc_instance_id,omitempty"`
	// Cluster Id to add ESX workflow
	EsxClusterId string `json:"esx_cluster_id,omitempty"`
	// vCenter public IP
	VcPublicIp string `json:"vc_public_ip,omitempty"`
	// skip creating vxlan for compute gateway for SDDC provisioning
	SkipCreatingVxlan bool `json:"skip_creating_vxlan,omitempty"`
	// URL of the vCenter server
	VcUrl string `json:"vc_url,omitempty"`
	SddcManifest *SddcManifest `json:"sddc_manifest,omitempty"`
	// VXLAN IP subnet
	VxlanSubnet string `json:"vxlan_subnet,omitempty"`
	// Group name for vCenter SDDC administrator
	CloudUserGroup string `json:"cloud_user_group,omitempty"`
	ManagementRp string `json:"management_rp,omitempty"`
	// region in which sddc is deployed
	Region string `json:"region,omitempty"`
	// Availability zone where the witness node is provisioned for a MultiAZ SDDC. This is null for a SingleAZ SDDC.
	WitnessAvailabilityZone string `json:"witness_availability_zone,omitempty"`
	// sddc identifier
	SddcId string `json:"sddc_id,omitempty"`
	PopAgentXeniConnection *PopAgentXeniConnection `json:"pop_agent_xeni_connection,omitempty"`
	// List of Controller IPs
	NsxControllerIps []string `json:"nsx_controller_ips,omitempty"`
	// ESX host subnet
	EsxHostSubnet string `json:"esx_host_subnet,omitempty"`
	// The SSO domain name to use for vSphere users
	SsoDomain string `json:"sso_domain,omitempty"`
	// Denotes if this is a SingleAZ SDDC or a MultiAZ SDDC.
	DeploymentType string `json:"deployment_type,omitempty"`
	NsxtAddons *NsxtAddons `json:"nsxt_addons,omitempty"`
	// if true, use the private IP addresses to register DNS records for the management VMs
	DnsWithManagementVmPrivateIp bool `json:"dns_with_management_vm_private_ip,omitempty"`
}



type SddcStateRequest struct {
	Sddcs []string `json:"sddcs,omitempty"`
	States []string `json:"states,omitempty"`
}


import (
	"time"
)

type SddcTemplate struct {
	Updated time.Time `json:"updated"`
	// User id that last updated this record
	UserId string `json:"user_id"`
	// User id that last updated this record
	UpdatedByUserId string `json:"updated_by_user_id"`
	Created time.Time `json:"created"`
	// Version of this entity
	Version int32 `json:"version"`
	// User name that last updated this record
	UpdatedByUserName string `json:"updated_by_user_name,omitempty"`
	// User name that last updated this record
	UserName string `json:"user_name"`
	// Unique ID for this entity
	Id string `json:"id"`
	// A list of the SDDC linking configurations to use.
	AccountLinkSddcConfigs []AccountLinkSddcConfig `json:"account_link_sddc_configs,omitempty"`
	State string `json:"state,omitempty"`
	NetworkTemplate *NetworkTemplate `json:"network_template,omitempty"`
	// name for SDDC configuration template
	Name string `json:"name,omitempty"`
	SourceSddcId string `json:"source_sddc_id,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	Sddc *Sddc `json:"sddc,omitempty"`
}



// Secondary IP addresses of the NSX Edge vnic address group. These are used for NAT, LB, VPN etc. Optional.
type SecondaryAddresses struct {
	Type_ string `json:"type,omitempty"`
	// List of IP addresses.
	IpAddress []string `json:"ipAddress,omitempty"`
}



// Detailed service errors associated with a task.
type ServiceError struct {
	// Error message in English.
	DefaultMessage string `json:"default_message,omitempty"`
	// The original service name of the error.
	OriginalService string `json:"original_service,omitempty"`
	// The localized message.
	LocalizedMessage string `json:"localized_message,omitempty"`
	// The original error code of the service.
	OriginalServiceErrorCode string `json:"original_service_error_code,omitempty"`
}



type Site struct {
	// Site password.
	Password string `json:"password,omitempty"`
	// Site user id.
	UserId string `json:"user_id,omitempty"`
	// Unique name for the site getting configured.
	Name string `json:"name,omitempty"`
	// Bytes received on local network.
	RxBytesOnLocalSubnet int64 `json:"rx_bytes_on_local_subnet,omitempty"`
	// Enable/disable encription.
	SecureTraffic bool `json:"secure_traffic,omitempty"`
	// Date tunnel was established.
	EstablishedDate string `json:"established_date,omitempty"`
	// failure message.
	FailureMessage string `json:"failure_message,omitempty"`
	// Number of transmitted packets dropped.
	DroppedTxPackets string `json:"dropped_tx_packets,omitempty"`
	// Number of received packets dropped.
	DroppedRxPackets string `json:"dropped_rx_packets,omitempty"`
	// Site tunnel status.
	TunnelStatus string `json:"tunnel_status,omitempty"`
	// Bytes transmitted from local subnet.
	TxBytesFromLocalSubnet int64 `json:"tx_bytes_from_local_subnet,omitempty"`
}



// L2 VPN sites.
type Sites struct {
	Sites []Nsxsite `json:"sites,omitempty"`
}



// Dashboard Statistics data for SSL VPN.
type SslvpnDashboardStats struct {
	// Number of active clients.
	ActiveClients []DashboardStat `json:"activeClients,omitempty"`
	// Rx bytes received for SSL VPN.
	SslvpnBytesIn []DashboardStat `json:"sslvpnBytesIn,omitempty"`
	// Number of authentication failures.
	AuthFailures []DashboardStat `json:"authFailures,omitempty"`
	// Number of SSL VPN sessions created.
	SessionsCreated []DashboardStat `json:"sessionsCreated,omitempty"`
	// Tx bytes transmitted for SSL VPN.
	SslvpnBytesOut []DashboardStat `json:"sslvpnBytesOut,omitempty"`
}



// NSX Edge sub interface configuration details. Sub interfaces are created on a trunk interface.
type SubInterface struct {
	// Index of the sub interface assigned by NSX Manager. Min value is 10 and max value is 4103.
	Index int32 `json:"index,omitempty"`
	// Valid values for tunnel ID are min 1 to max 4093. Required.
	TunnelId int32 `json:"tunnelId"`
	// Name of the sub interface. Required.
	Name string `json:"name,omitempty"`
	AddressGroups *EdgeVnicAddressGroups `json:"addressGroups,omitempty"`
	// VLAN ID of the virtual LAN used by this sub interface. VLAN IDs can range from 0 to 4094.
	VlanId int32 `json:"vlanId,omitempty"`
	// Sub interface label of format vNic_{index} provided by NSX Manager. Read only.
	Label string `json:"label,omitempty"`
	// Name of the logical switch connected to this sub interface.
	LogicalSwitchName string `json:"logicalSwitchName,omitempty"`
	// Value is true if the sub interface is connected to a logical switch, standard portgroup or distributed portgroup.
	IsConnected bool `json:"isConnected,omitempty"`
	// MTU value of the sub interface. This value would be the least mtu for all the trunk interfaces of the NSX Edge. Default is 1500.
	Mtu int32 `json:"mtu,omitempty"`
	// ID of the logical switch connected to this sub interface.
	LogicalSwitchId string `json:"logicalSwitchId,omitempty"`
	// Value is true if send redirects is enabled. Enable ICMP redirect to convey routing information to hosts.
	EnableSendRedirects bool `json:"enableSendRedirects,omitempty"`
}



type SubInterfaces struct {
	// List of sub interfaces.
	SubInterfaces []SubInterface `json:"subInterfaces,omitempty"`
}



// (as there's already one SubnetInfo, use Subnet instead)
type Subnet struct {
	// subnet id
	SubnetId string `json:"subnet_id,omitempty"`
	// subnet name
	Name string `json:"name,omitempty"`
	RouteTables []SubnetRouteTableInfo `json:"route_tables,omitempty"`
}



type SubnetInfo struct {
}



type SubnetRouteTableInfo struct {
	// subnet id
	SubnetId string `json:"subnet_id,omitempty"`
	// subnet - route table association id
	AssociationId string `json:"association_id,omitempty"`
	// route table id
	RoutetableId string `json:"routetable_id,omitempty"`
}



type Subnets struct {
	// List of subnets for which IPsec VPN is configured. Subnets should be network address specified in CIDR format and can accept '0.0.0.0/0' (any)
	Subnets []string `json:"subnets,omitempty"`
}



// details of a subscription
type SubscriptionDetails struct {
	Status string `json:"status,omitempty"`
	AnniversaryBillingDate string `json:"anniversary_billing_date,omitempty"`
	EndDate string `json:"end_date,omitempty"`
	// The frequency at which the customer is billed. Currently supported values are \"Upfront\" and \"Monthly\"
	BillingFrequency string `json:"billing_frequency,omitempty"`
	AutoRenewedAllowed string `json:"auto_renewed_allowed,omitempty"`
	CommitmentTerm string `json:"commitment_term,omitempty"`
	CspSubscriptionId string `json:"csp_subscription_id,omitempty"`
	BillingSubscriptionId string `json:"billing_subscription_id,omitempty"`
	OfferVersion string `json:"offer_version,omitempty"`
	OfferType string `json:"offer_type,omitempty"`
	Description string `json:"description,omitempty"`
	ProductId string `json:"product_id,omitempty"`
	Region string `json:"region,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	OfferName string `json:"offer_name,omitempty"`
	// unit of measurment for commitment term
	CommitmentTermUom string `json:"commitment_term_uom,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}



// Details of products that are available for purchase.
type SubscriptionProducts struct {
	// The name of the product
	Product string `json:"product,omitempty"`
	// A list of different types/version for the product.
	Types []string `json:"types,omitempty"`
}



type SubscriptionRequest struct {
	// The product for which subscription needs to be created. Refer /vmc/api/orgs/{orgId}/products.
	Product string `json:"product,omitempty"`
	// Old identifier for product. *Deprecarted*. See product and type
	ProductType string `json:"product_type"`
	ProductId string `json:"product_id,omitempty"`
	// Frequency of the billing.
	BillingFrequency string `json:"billing_frequency,omitempty"`
	Region string `json:"region"`
	CommitmentTerm string `json:"commitment_term"`
	OfferContextId string `json:"offer_context_id,omitempty"`
	OfferVersion string `json:"offer_version"`
	OfferName string `json:"offer_name"`
	Quantity int32 `json:"quantity"`
	// The type of the product for which the subscription needs to be created.
	Type_ string `json:"type,omitempty"`
	Price int32 `json:"price,omitempty"`
	ProductChargeId string `json:"product_charge_id,omitempty"`
}



type SupportWindow struct {
	StartDay string `json:"start_day,omitempty"`
	Seats int64 `json:"seats,omitempty"`
	// SDDCs in this window
	Sddcs []string `json:"sddcs,omitempty"`
	DurationHours int64 `json:"duration_hours,omitempty"`
	StartHour int32 `json:"start_hour,omitempty"`
	SupportWindowId string `json:"support_window_id,omitempty"`
	Metadata *interface{} `json:"metadata,omitempty"`
}



type SupportWindowId struct {
	// Support Window ID
	WindowId string `json:"window_id,omitempty"`
}


import (
	"time"
)

type Task struct {
	Updated time.Time `json:"updated"`
	// User id that last updated this record
	UserId string `json:"user_id"`
	// User id that last updated this record
	UpdatedByUserId string `json:"updated_by_user_id"`
	Created time.Time `json:"created"`
	// Version of this entity
	Version int32 `json:"version"`
	// User name that last updated this record
	UpdatedByUserName string `json:"updated_by_user_name,omitempty"`
	// User name that last updated this record
	UserName string `json:"user_name"`
	// Unique ID for this entity
	Id string `json:"id"`
	Status string `json:"status,omitempty"`
	LocalizedErrorMessage string `json:"localized_error_message,omitempty"`
	// UUID of the resource the task is acting upon
	ResourceId string `json:"resource_id,omitempty"`
	// If this task was created by another task - this provides the linkage. Mostly for debugging.
	ParentTaskId string `json:"parent_task_id,omitempty"`
	TaskVersion string `json:"task_version,omitempty"`
	// (Optional) Client provided uniqifier to make task creation idempotent. Be aware not all tasks support this. For tasks that do - supplying the same correlation Id, for the same task type, within a predefined window will ensure the operation happens at most once.
	CorrelationId string `json:"correlation_id,omitempty"`
	// Entity version of the resource at the start of the task. This is only set for some task types.
	StartResourceEntityVersion int32 `json:"start_resource_entity_version,omitempty"`
	SubStatus string `json:"sub_status,omitempty"`
	TaskType string `json:"task_type,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	// Task progress phases involved in current task execution
	TaskProgressPhases []TaskProgressPhase `json:"task_progress_phases,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// Entity version of the resource at the end of the task. This is only set for some task types.
	EndResourceEntityVersion int32 `json:"end_resource_entity_version,omitempty"`
	// Service errors returned from SDDC services.
	ServiceErrors []ServiceError `json:"service_errors,omitempty"`
	OrgType string `json:"org_type,omitempty"`
	// Estimated remaining time in minute of the task execution, < 0 means no estimation for the task.
	EstimatedRemainingMinutes int32 `json:"estimated_remaining_minutes,omitempty"`
	Params *interface{} `json:"params,omitempty"`
	// Estimated progress percentage the task executed
	ProgressPercent int32 `json:"progress_percent,omitempty"`
	// The current in progress phase ID in the task execution, if none in progress, empty string returned.
	PhaseInProgress string `json:"phase_in_progress,omitempty"`
	// Type of resource being acted upon
	ResourceType string `json:"resource_type,omitempty"`
	EndTime time.Time `json:"end_time,omitempty"`
}



// A task progress can be (but does NOT have to be) divided to more meaningful progress phases.
type TaskProgressPhase struct {
	// The identifier of the task progress phase
	Id string `json:"id"`
	// The display name of the task progress phase
	Name string `json:"name"`
	// The percentage of the phase that has completed
	ProgressPercent int32 `json:"progress_percent"`
}



// Holder for term billing options.
type TermBillingOptions struct {
	UnitPrice string `json:"unit_price,omitempty"`
	BillingFrequency string `json:"billing_frequency,omitempty"`
}



// Holder for the term offer instances.
type TermOfferInstance struct {
	Description string `json:"description"`
	Product string `json:"product"`
	// Deprecated. Please use product and type fields instead.
	ProductType string `json:"product_type,omitempty"`
	Name string `json:"name"`
	Currency string `json:"currency"`
	Region string `json:"region"`
	CommitmentTerm int32 `json:"commitment_term"`
	// (deprecated. unit_price is moved into TermBillingOptions. For backward compatibility, this field reflect \"Prepaid\" price at the offer level.)
	UnitPrice string `json:"unit_price"`
	BillingOptions []TermBillingOptions `json:"billing_options,omitempty"`
	Version string `json:"version"`
	OfferContextId string `json:"offer_context_id,omitempty"`
	ProductChargeId string `json:"product_charge_id,omitempty"`
	Type_ string `json:"type"`
	ProductId string `json:"product_id,omitempty"`
}



type TrafficShapingPolicy struct {
	BurstSize int64 `json:"burstSize,omitempty"`
	AverageBandwidth int64 `json:"averageBandwidth,omitempty"`
	PeakBandwidth int64 `json:"peakBandwidth,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Inherited bool `json:"inherited,omitempty"`
}



type UpdateCredentials struct {
	// Username of the credentials
	Username string `json:"username"`
	// Password associated with the credentials
	Password string `json:"password"`
}



type VmcLocale struct {
	// The locale to be used for translating responses for the session
	Locale string `json:"locale,omitempty"`
}



// NSX Edge vnic configuration details.
type Vnic struct {
	SubInterfaces *SubInterfaces `json:"subInterfaces,omitempty"`
	AddressGroups *EdgeVnicAddressGroups `json:"addressGroups,omitempty"`
	// Value is true if the vnic is connected to a logical switch, standard portgroup or distributed portgroup.
	IsConnected bool `json:"isConnected,omitempty"`
	// Value is true if send redirects is enabled. Enable ICMP redirect to convey routing information to hosts.
	EnableSendRedirects bool `json:"enableSendRedirects,omitempty"`
	InShapingPolicy *TrafficShapingPolicy `json:"inShapingPolicy,omitempty"`
	// Interface label of format vNic_{vnicIndex} provided by NSX Manager. Read only.
	Label string `json:"label,omitempty"`
	// Value is true if proxy arp is enabled. Enable proxy ARP if you want to allow the NSX Edge of type 'gatewayServices' to answer ARP requests intended for other machines.
	EnableProxyArp bool `json:"enableProxyArp,omitempty"`
	// Index of the vnic. Min value is 0 and max value is 9.
	Index int32 `json:"index"`
	// Name of the interface. Optional.
	Name string `json:"name,omitempty"`
	// MTU of the interface, with default as 1500. Min is 68, Max is 9000. Optional.
	Mtu int32 `json:"mtu,omitempty"`
	FenceParameters []KeyValueAttributes `json:"fenceParameters,omitempty"`
	// Distinct MAC addresses configured for the vnic. Optional.
	MacAddresses []MacAddress `json:"macAddresses,omitempty"`
	OutShapingPolicy *TrafficShapingPolicy `json:"outShapingPolicy,omitempty"`
	// Name of the port group or logical switch.
	PortgroupName string `json:"portgroupName,omitempty"`
	// Value is true if bridge mode is enabled.
	EnableBridgeMode bool `json:"enableBridgeMode,omitempty"`
	// Type of the vnic. Values are uplink, internal, trunk. At least one internal interface must be configured for NSX Edge HA to work.
	Type_ string `json:"type,omitempty"`
	// Value are port group ID (standard portgroup or distributed portgroup) or virtual wire ID (logical switch). Logical switch cannot be used for a TRUNK vnic. Portgroup cannot be shared among vnics/LIFs. Required when isConnected is specified as true. Example 'network-17' (standard portgroup), 'dvportgroup-34' (distributed portgroup) or 'virtualwire-2' (logical switch).
	PortgroupId string `json:"portgroupId,omitempty"`
}



// Ordered list of NSX Edge vnics. Until one connected vnic is configured, none of the configured features will serve the network.
type Vnics struct {
	// Ordered list of NSX Edge vnics.
	Vnics []Vnic `json:"vnics,omitempty"`
}



type VpcInfo struct {
	VpcCidr string `json:"vpc_cidr,omitempty"`
	VgwId string `json:"vgw_id,omitempty"`
	EsxPublicSecurityGroupId string `json:"esx_public_security_group_id,omitempty"`
	// set of virtual interfaces attached to the sddc
	VifIds []string `json:"vif_ids,omitempty"`
	VmSecurityGroupId string `json:"vm_security_group_id,omitempty"`
	// Mapping from AZ to a list of IP addresses assigned to TGW ENI that's connected with Vpc
	TgwIps map[string][]string `json:"tgwIps,omitempty"`
	// (deprecated)
	RouteTableId string `json:"route_table_id,omitempty"`
	// Id of the NSX edge associated with this VPC (deprecated)
	EdgeSubnetId string `json:"edge_subnet_id,omitempty"`
	// vpc id
	Id string `json:"id,omitempty"`
	// Id of the association between subnet and route-table (deprecated)
	ApiAssociationId string `json:"api_association_id,omitempty"`
	// Id associated with this VPC (deprecated)
	ApiSubnetId string `json:"api_subnet_id,omitempty"`
	// (deprecated)
	PrivateSubnetId string `json:"private_subnet_id,omitempty"`
	// (deprecated)
	PrivateAssociationId string `json:"private_association_id,omitempty"`
	EsxSecurityGroupId string `json:"esx_security_group_id,omitempty"`
	// (deprecated)
	SubnetId string `json:"subnet_id,omitempty"`
	InternetGatewayId string `json:"internet_gateway_id,omitempty"`
	SecurityGroupId string `json:"security_group_id,omitempty"`
	// (deprecated)
	AssociationId string `json:"association_id,omitempty"`
	// Route table which contains the route to VGW (deprecated)
	VgwRouteTableId string `json:"vgw_route_table_id,omitempty"`
	// Id of the association between edge subnet and route-table (deprecated)
	EdgeAssociationId string `json:"edge_association_id,omitempty"`
	Provider string `json:"provider,omitempty"`
	// (deprecated)
	PeeringConnectionId string `json:"peering_connection_id,omitempty"`
	NetworkType string `json:"network_type,omitempty"`
	AvailableZones []AvailableZoneInfo `json:"available_zones,omitempty"`
	// map from routeTableName to routeTableInfo
	Routetables map[string]RouteTableInfo `json:"routetables,omitempty"`
}



type VpcInfoSubnets struct {
}



type Vpn struct {
	Version string `json:"version,omitempty"`
	OnPremGatewayIp string `json:"on_prem_gateway_ip,omitempty"`
	OnPremNetworkCidr string `json:"on_prem_network_cidr,omitempty"`
	PfsEnabled bool `json:"pfs_enabled,omitempty"`
	Id string `json:"id,omitempty"`
	ChannelStatus *VpnChannelStatus `json:"channel_status,omitempty"`
	OnPremNatIp string `json:"on_prem_nat_ip,omitempty"`
	Name string `json:"name,omitempty"`
	InternalNetworkIds []string `json:"internal_network_ids,omitempty"`
	TunnelStatuses []VpnTunnelStatus `json:"tunnel_statuses,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	State string `json:"state,omitempty"`
	DhGroup string `json:"dh_group,omitempty"`
	Authentication string `json:"authentication,omitempty"`
	PreSharedKey string `json:"pre_shared_key,omitempty"`
	IkeOption string `json:"ike_option,omitempty"`
	DigestAlgorithm string `json:"digest_algorithm,omitempty"`
}



type VpnChannelStatus struct {
	ChannelStatus string `json:"channel_status,omitempty"`
	ChannelState string `json:"channel_state,omitempty"`
	LastInfoMessage string `json:"last_info_message,omitempty"`
	FailureMessage string `json:"failure_message,omitempty"`
}



type VpnTunnelStatus struct {
	OnPremSubnet string `json:"on_prem_subnet,omitempty"`
	TrafficStats *VpnTunnelTrafficStats `json:"traffic_stats,omitempty"`
	LastInfoMessage string `json:"last_info_message,omitempty"`
	LocalSubnet string `json:"local_subnet,omitempty"`
	TunnelState string `json:"tunnel_state,omitempty"`
	FailureMessage string `json:"failure_message,omitempty"`
	TunnelStatus string `json:"tunnel_status,omitempty"`
}



type VpnTunnelTrafficStats struct {
	PacketsOut string `json:"packets_out,omitempty"`
	PacketReceivedErrors string `json:"packet_received_errors,omitempty"`
	RxBytesOnLocalSubnet string `json:"rx_bytes_on_local_subnet,omitempty"`
	ReplayErrors string `json:"replay_errors,omitempty"`
	SequenceNumberOverFlowErrors string `json:"sequence_number_over_flow_errors,omitempty"`
	EncryptionFailures string `json:"encryption_failures,omitempty"`
	IntegrityErrors string `json:"integrity_errors,omitempty"`
	PacketSentErrors string `json:"packet_sent_errors,omitempty"`
	DecryptionFailures string `json:"decryption_failures,omitempty"`
	PacketsIn string `json:"packets_in,omitempty"`
	TxBytesFromLocalSubnet string `json:"tx_bytes_from_local_subnet,omitempty"`
}



// Infomation about an available vSAN capacity in a cluster.
type VsanAvailableCapacity struct {
	Cost string `json:"cost"`
	Quality string `json:"quality"`
	Size int64 `json:"size"`
}



// Bias for reconfiguring vSAN in a cluster.
type VsanClusterReconfigBias struct {
	ShortDescription string `json:"short_description"`
	FullDescription string `json:"full_description"`
	Id string `json:"id"`
}



// Storage constraint information for reconfiguring vSAN in existing cluster.
type VsanClusterReconfigConstraints struct {
	// Biases to reconfigure vSAN in an existing cluster.
	ReconfigBiases []VsanClusterReconfigBias `json:"reconfig_biases"`
	// A map of VsanClusterReconfigBias id to the list of VsanAvailableCapacity. It gives all of available vSAN capacities for each of reconfiguration biases. 
	AvailableCapacities map[string][]VsanAvailableCapacity `json:"available_capacities"`
	// A map of VsanClusterReconfigBias id to a VsanAvailableCapacity. It gives the default VsanAvailableCapacity for each of reconfiguration biases. 
	DefaultCapacities map[string]VsanAvailableCapacity `json:"default_capacities"`
	// The number of hosts in a cluster for the constraints.
	Hosts int32 `json:"hosts"`
	// The id of default VsanClusterReconfigBias for this constraints.
	DefaultReconfigBiasId string `json:"default_reconfig_bias_id"`
}



// This describes the possible physical storage capacity choices for use with a given VsanStorageDesigner implementation.  These choices are specific to a customer-defined number of hosts per cluster. 
type VsanConfigConstraints struct {
	// Maximum capacity supported for cluster (GiB).
	MaxCapacity int64 `json:"max_capacity"`
	// List of supported capacities which may offer preferable performance (GiB).
	RecommendedCapacities []int64 `json:"recommended_capacities"`
	// Increment to be added to min_capacity to result in a supported capacity (GiB).
	SupportedCapacityIncrement int64 `json:"supported_capacity_increment,omitempty"`
	// Minimum capacity supported for cluster (GiB).
	MinCapacity int64 `json:"min_capacity"`
	// Number of hosts in cluster.
	NumHosts int64 `json:"num_hosts"`
}



type VsanEncryptionConfig struct {
	// Port to connect to AWS Key Management Service
	Port int32 `json:"port,omitempty"`
	// Public certificate used to connect to AWS Key Management Service
	Certificate string `json:"certificate,omitempty"`
}
