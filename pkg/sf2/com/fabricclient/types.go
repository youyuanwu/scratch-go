package fabricclient

import (
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// struct __MIDL___MIDL_itf_FabricClient_0004_0001_0001
type MIDL___MIDL_itf_FabricClient_0004_0001_0001__ struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// struct FABRIC_SERVICE_DESCRIPTION
type FABRIC_SERVICE_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_RESOLVED_SERVICE_PARTITION
type FABRIC_RESOLVED_SERVICE_PARTITION struct {
	Info          FABRIC_SERVICE_PARTITION_INFORMATION
	EndpointCount uint32
	Endpoints     *FABRIC_RESOLVED_SERVICE_ENDPOINT
	ServiceName   win32.PWSTR
	Reserved      unsafe.Pointer
}

// struct FABRIC_SERVICE_PARTITION_INFORMATION
type FABRIC_SERVICE_PARTITION_INFORMATION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_RESOLVED_SERVICE_ENDPOINT
type FABRIC_RESOLVED_SERVICE_ENDPOINT struct {
	Address  win32.PWSTR
	Role     int32
	Reserved unsafe.Pointer
}

// struct FABRIC_SERVICE_NOTIFICATION
type FABRIC_SERVICE_NOTIFICATION struct {
	ServiceName   win32.PWSTR
	PartitionId   syscall.GUID
	EndpointCount uint32
	Endpoints     *FABRIC_RESOLVED_SERVICE_ENDPOINT
	PartitionInfo *FABRIC_SERVICE_PARTITION_INFORMATION
	Reserved      unsafe.Pointer
}

// struct FABRIC_GATEWAY_INFORMATION
type FABRIC_GATEWAY_INFORMATION struct {
	NodeAddress    win32.PWSTR
	NodeId         FABRIC_NODE_ID
	NodeInstanceId uint64
	NodeName       win32.PWSTR
	Reserved       unsafe.Pointer
}

// struct FABRIC_NODE_ID
type FABRIC_NODE_ID struct {
	Low      uint64
	High     uint64
	Reserved unsafe.Pointer
}

// struct FABRIC_CLAIMS_RETRIEVAL_METADATA
type FABRIC_CLAIMS_RETRIEVAL_METADATA struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SECURITY_CREDENTIALS
type FABRIC_SECURITY_CREDENTIALS struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_CLIENT_SETTINGS
type FABRIC_CLIENT_SETTINGS struct {
	PartitionLocationCacheLimit              uint32
	ServiceChangePollIntervalInSeconds       uint32
	ConnectionInitializationTimeoutInSeconds uint32
	KeepAliveIntervalInSeconds               uint32
	HealthOperationTimeoutInSeconds          uint32
	HealthReportSendIntervalInSeconds        uint32
	Reserved                                 unsafe.Pointer
}

// struct FABRIC_NAMED_PROPERTY_METADATA
type FABRIC_NAMED_PROPERTY_METADATA struct {
	PropertyName    win32.PWSTR
	TypeId          int32
	ValueSize       int32
	SequenceNumber  int64
	LastModifiedUtc FILETIME_
	Name            win32.PWSTR
	Reserved        unsafe.Pointer
}

// struct _FILETIME
type FILETIME_ struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

// struct FABRIC_NAMED_PROPERTY
type FABRIC_NAMED_PROPERTY struct {
	Metadata *FABRIC_NAMED_PROPERTY_METADATA
	Value    *byte
	Reserved unsafe.Pointer
}

// struct FABRIC_PROPERTY_BATCH_OPERATION
type FABRIC_PROPERTY_BATCH_OPERATION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_PUT_CUSTOM_PROPERTY_OPERATION
type FABRIC_PUT_CUSTOM_PROPERTY_OPERATION struct {
	PropertyName         win32.PWSTR
	PropertyTypeId       int32
	PropertyValue        unsafe.Pointer
	PropertyCustomTypeId win32.PWSTR
	Reserved             unsafe.Pointer
}

// struct FABRIC_SERVICE_UPDATE_DESCRIPTION
type FABRIC_SERVICE_UPDATE_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_REMOVE_REPLICA_DESCRIPTION
type FABRIC_REMOVE_REPLICA_DESCRIPTION struct {
	NodeName            win32.PWSTR
	PartitionId         syscall.GUID
	ReplicaOrInstanceId int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_RESTART_REPLICA_DESCRIPTION
type FABRIC_RESTART_REPLICA_DESCRIPTION struct {
	NodeName            win32.PWSTR
	PartitionId         syscall.GUID
	ReplicaOrInstanceId int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_SERVICE_NOTIFICATION_FILTER_DESCRIPTION
type FABRIC_SERVICE_NOTIFICATION_FILTER_DESCRIPTION struct {
	Name     win32.PWSTR
	Flags    int32
	Reserved unsafe.Pointer
}

// struct FABRIC_DELETE_SERVICE_DESCRIPTION
type FABRIC_DELETE_SERVICE_DESCRIPTION struct {
	ServiceName win32.PWSTR
	ForceDelete int8
	Reserved    unsafe.Pointer
}

// struct FABRIC_SERVICE_FROM_TEMPLATE_DESCRIPTION
type FABRIC_SERVICE_FROM_TEMPLATE_DESCRIPTION struct {
	ApplicationName              win32.PWSTR
	ServiceName                  win32.PWSTR
	ServiceDnsName               win32.PWSTR
	ServiceTypeName              win32.PWSTR
	ServicePackageActivationMode int32
	InitializationDataSize       uint32
	InitializationData           *byte
	Reserved                     unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_DESCRIPTION
type FABRIC_SERVICE_GROUP_DESCRIPTION struct {
	Description        *FABRIC_SERVICE_DESCRIPTION
	MemberCount        uint32
	MemberDescriptions *FABRIC_SERVICE_GROUP_MEMBER_DESCRIPTION
	Reserved           unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_DESCRIPTION
type FABRIC_SERVICE_GROUP_MEMBER_DESCRIPTION struct {
	ServiceType            win32.PWSTR
	ServiceName            win32.PWSTR
	InitializationDataSize uint32
	InitializationData     *byte
	MetricCount            uint32
	Metrics                *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
	Reserved               unsafe.Pointer
}

// struct FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
type FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION struct {
	Name                 win32.PWSTR
	Weight               int32
	PrimaryDefaultLoad   uint32
	SecondaryDefaultLoad uint32
	Reserved             unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_UPDATE_DESCRIPTION
type FABRIC_SERVICE_GROUP_UPDATE_DESCRIPTION struct {
	Description *FABRIC_SERVICE_UPDATE_DESCRIPTION
	Reserved    unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_FROM_TEMPLATE_DESCRIPTION
type FABRIC_SERVICE_GROUP_FROM_TEMPLATE_DESCRIPTION struct {
	ApplicationName              win32.PWSTR
	ServiceName                  win32.PWSTR
	ServiceTypeName              win32.PWSTR
	ServicePackageActivationMode int32
	InitializationDataSize       uint32
	InitializationData           *byte
	Reserved                     unsafe.Pointer
}

// struct FABRIC_APPLICATION_DESCRIPTION
type FABRIC_APPLICATION_DESCRIPTION struct {
	ApplicationName        win32.PWSTR
	ApplicationTypeName    win32.PWSTR
	ApplicationTypeVersion win32.PWSTR
	ApplicationParameters  *FABRIC_APPLICATION_PARAMETER_LIST
	Reserved               unsafe.Pointer
}

// struct FABRIC_APPLICATION_PARAMETER_LIST
type FABRIC_APPLICATION_PARAMETER_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_PARAMETER
}

// struct FABRIC_APPLICATION_PARAMETER
type FABRIC_APPLICATION_PARAMETER struct {
	Name     win32.PWSTR
	Value    win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_APPLICATION_UPGRADE_DESCRIPTION
type FABRIC_APPLICATION_UPGRADE_DESCRIPTION struct {
	ApplicationName              win32.PWSTR
	TargetApplicationTypeVersion win32.PWSTR
	ApplicationParameters        *FABRIC_APPLICATION_PARAMETER_LIST
	UpgradeKind                  int32
	UpgradePolicyDescription     unsafe.Pointer
	Reserved                     unsafe.Pointer
}

// struct FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION
type FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION struct {
	Name     win32.PWSTR
	State    int32
	Reserved unsafe.Pointer
}

// struct FABRIC_APPLICATION_UPGRADE_UPDATE_DESCRIPTION
type FABRIC_APPLICATION_UPGRADE_UPDATE_DESCRIPTION struct {
	ApplicationName          win32.PWSTR
	UpgradeKind              int32
	UpdateFlags              uint32
	UpgradePolicyDescription unsafe.Pointer
	Reserved                 unsafe.Pointer
}

// struct FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION
type FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION struct {
	NodeName              win32.PWSTR
	ApplicationName       win32.PWSTR
	ServiceManifestName   win32.PWSTR
	CodePackageName       win32.PWSTR
	CodePackageInstanceId int64
	Reserved              unsafe.Pointer
}

// struct FABRIC_PACKAGE_SHARING_POLICY_LIST
type FABRIC_PACKAGE_SHARING_POLICY_LIST struct {
	Count uint32
	Items *FABRIC_PACKAGE_SHARING_POLICY
}

// struct FABRIC_PACKAGE_SHARING_POLICY
type FABRIC_PACKAGE_SHARING_POLICY struct {
	PackageName win32.PWSTR
	Scope       int32
	Reserved    unsafe.Pointer
}

// struct FABRIC_APPLICATION_UPDATE_DESCRIPTION
type FABRIC_APPLICATION_UPDATE_DESCRIPTION struct {
	Flags                     uint32
	ApplicationName           win32.PWSTR
	RemoveApplicationCapacity int8
	MaximumNodes              uint32
	MinimumNodes              uint32
	Metrics                   *FABRIC_APPLICATION_METRIC_LIST
	Reserved                  unsafe.Pointer
}

// struct FABRIC_APPLICATION_METRIC_LIST
type FABRIC_APPLICATION_METRIC_LIST struct {
	Count      uint32
	Capacities *FABRIC_APPLICATION_METRIC_DESCRIPTION
}

// struct FABRIC_APPLICATION_METRIC_DESCRIPTION
type FABRIC_APPLICATION_METRIC_DESCRIPTION struct {
	Name                     win32.PWSTR
	NodeReservationCapacity  uint32
	MaximumNodeCapacity      uint32
	TotalApplicationCapacity uint32
	Reserved                 unsafe.Pointer
}

// struct FABRIC_DELETE_APPLICATION_DESCRIPTION
type FABRIC_DELETE_APPLICATION_DESCRIPTION struct {
	ApplicationName win32.PWSTR
	ForceDelete     int8
	Reserved        unsafe.Pointer
}

// struct FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION
type FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION struct {
	BuildPath win32.PWSTR
	Async     int8
	Reserved  unsafe.Pointer
}

// struct FABRIC_UNPROVISION_APPLICATION_TYPE_DESCRIPTION
type FABRIC_UNPROVISION_APPLICATION_TYPE_DESCRIPTION struct {
	ApplicationTypeName    win32.PWSTR
	ApplicationTypeVersion win32.PWSTR
	Async                  int8
	Reserved               unsafe.Pointer
}

// struct FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION_BASE
type FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION_BASE struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_UPGRADE_DESCRIPTION
type FABRIC_UPGRADE_DESCRIPTION struct {
	CodeVersion              win32.PWSTR
	ConfigVersion            win32.PWSTR
	UpgradeKind              int32
	UpgradePolicyDescription unsafe.Pointer
	Reserved                 unsafe.Pointer
}

// struct FABRIC_UPGRADE_UPDATE_DESCRIPTION
type FABRIC_UPGRADE_UPDATE_DESCRIPTION struct {
	UpgradeKind              int32
	UpdateFlags              uint32
	UpgradePolicyDescription unsafe.Pointer
	Reserved                 unsafe.Pointer
}

// struct FABRIC_STOP_NODE_DESCRIPTION
type FABRIC_STOP_NODE_DESCRIPTION struct {
	NodeName       win32.PWSTR
	NodeInstanceId uint64
	Reserved       unsafe.Pointer
}

// struct FABRIC_RESTART_NODE_DESCRIPTION
type FABRIC_RESTART_NODE_DESCRIPTION struct {
	NodeName       win32.PWSTR
	NodeInstanceId uint64
	Reserved       unsafe.Pointer
}

// struct FABRIC_START_NODE_DESCRIPTION
type FABRIC_START_NODE_DESCRIPTION struct {
	NodeName              win32.PWSTR
	NodeInstanceId        uint64
	IPAddressOrFQDN       win32.PWSTR
	ClusterConnectionPort uint32
	Reserved              unsafe.Pointer
}

// struct FABRIC_START_UPGRADE_DESCRIPTION
type FABRIC_START_UPGRADE_DESCRIPTION struct {
	ClusterConfig                              win32.PWSTR
	HealthCheckRetryTimeoutInSeconds           uint32
	HealthCheckWaitDurationInSeconds           uint32
	HealthCheckStableDurationInSeconds         uint32
	UpgradeDomainTimeoutInSeconds              uint32
	UpgradeTimeoutInSeconds                    uint32
	MaxPercentUnhealthyApplications            byte
	MaxPercentUnhealthyNodes                   byte
	MaxPercentDeltaUnhealthyNodes              byte
	MaxPercentUpgradeDomainDeltaUnhealthyNodes byte
	Reserved                                   unsafe.Pointer
}

// struct FABRIC_ORCHESTRATION_UPGRADE_PROGRESS
type FABRIC_ORCHESTRATION_UPGRADE_PROGRESS struct {
	UpgradeState   int32
	ProgressStatus uint32
	Reserved       unsafe.Pointer
}

// struct FABRIC_HEALTH_REPORT
type FABRIC_HEALTH_REPORT struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH_POLICY
type FABRIC_CLUSTER_HEALTH_POLICY struct {
	ConsiderWarningAsError          int8
	MaxPercentUnhealthyNodes        byte
	MaxPercentUnhealthyApplications byte
	Reserved                        unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH
type FABRIC_CLUSTER_HEALTH struct {
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH
type FABRIC_NODE_HEALTH struct {
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	HealthEvents          *FABRIC_HEALTH_EVENT_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_HEALTH_EVENT_LIST
type FABRIC_HEALTH_EVENT_LIST struct {
	Count uint32
	Items *FABRIC_HEALTH_EVENT
}

// struct FABRIC_HEALTH_EVENT
type FABRIC_HEALTH_EVENT struct {
	HealthInformation        *FABRIC_HEALTH_INFORMATION
	SourceUtcTimestamp       FILETIME_
	LastModifiedUtcTimestamp FILETIME_
	IsExpired                int8
	Reserved                 unsafe.Pointer
}

// struct FABRIC_HEALTH_INFORMATION
type FABRIC_HEALTH_INFORMATION struct {
	SourceId          win32.PWSTR
	Property          win32.PWSTR
	TimeToLiveSeconds uint32
	State             int32
	Description       win32.PWSTR
	SequenceNumber    int64
	RemoveWhenExpired int8
	Reserved          unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_POLICY
type FABRIC_APPLICATION_HEALTH_POLICY struct {
	ConsiderWarningAsError                  int8
	MaxPercentUnhealthyDeployedApplications byte
	DefaultServiceTypeHealthPolicy          *FABRIC_SERVICE_TYPE_HEALTH_POLICY
	ServiceTypeHealthPolicyMap              *FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP
	Reserved                                unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_HEALTH_POLICY
type FABRIC_SERVICE_TYPE_HEALTH_POLICY struct {
	MaxPercentUnhealthyServices             byte
	MaxPercentUnhealthyPartitionsPerService byte
	MaxPercentUnhealthyReplicasPerPartition byte
	Reserved                                unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP
type FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP struct {
	Count uint32
	Items *FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP_ITEM
}

// struct FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP_ITEM
type FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP_ITEM struct {
	ServiceTypeName         win32.PWSTR
	ServiceTypeHealthPolicy *FABRIC_SERVICE_TYPE_HEALTH_POLICY
}

// struct FABRIC_APPLICATION_HEALTH
type FABRIC_APPLICATION_HEALTH struct {
	ApplicationName                 win32.PWSTR
	AggregatedHealthState           int32
	HealthEvents                    *FABRIC_HEALTH_EVENT_LIST
	DeployedApplicationHealthStates *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_LIST
	ServiceHealthStates             *FABRIC_SERVICE_HEALTH_STATE_LIST
	Reserved                        unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_LIST
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE struct {
	ApplicationName       win32.PWSTR
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATE_LIST
type FABRIC_SERVICE_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_HEALTH_STATE
}

// struct FABRIC_SERVICE_HEALTH_STATE
type FABRIC_SERVICE_HEALTH_STATE struct {
	ServiceName           win32.PWSTR
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH
type FABRIC_SERVICE_HEALTH struct {
	ServiceName           win32.PWSTR
	AggregatedHealthState int32
	HealthEvents          *FABRIC_HEALTH_EVENT_LIST
	PartitionHealthStates *FABRIC_PARTITION_HEALTH_STATE_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATE_LIST
type FABRIC_PARTITION_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_PARTITION_HEALTH_STATE
}

// struct FABRIC_PARTITION_HEALTH_STATE
type FABRIC_PARTITION_HEALTH_STATE struct {
	PartitionId           syscall.GUID
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH
type FABRIC_PARTITION_HEALTH struct {
	PartitionId           syscall.GUID
	AggregatedHealthState int32
	HealthEvents          *FABRIC_HEALTH_EVENT_LIST
	ReplicaHealthStates   *FABRIC_REPLICA_HEALTH_STATE_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATE_LIST
type FABRIC_REPLICA_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_REPLICA_HEALTH_STATE
}

// struct FABRIC_REPLICA_HEALTH_STATE
type FABRIC_REPLICA_HEALTH_STATE struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH
type FABRIC_REPLICA_HEALTH struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH
type FABRIC_DEPLOYED_APPLICATION_HEALTH struct {
	ApplicationName                    win32.PWSTR
	NodeName                           win32.PWSTR
	AggregatedHealthState              int32
	HealthEvents                       *FABRIC_HEALTH_EVENT_LIST
	DeployedServicePackageHealthStates *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_LIST
	Reserved                           unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_LIST
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE struct {
	ApplicationName       win32.PWSTR
	ServiceManifestName   win32.PWSTR
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH struct {
	ApplicationName       win32.PWSTR
	ServiceManifestName   win32.PWSTR
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	HealthEvents          *FABRIC_HEALTH_EVENT_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH_QUERY_DESCRIPTION
type FABRIC_CLUSTER_HEALTH_QUERY_DESCRIPTION struct {
	HealthPolicy               *FABRIC_CLUSTER_HEALTH_POLICY
	ApplicationHealthPolicyMap *FABRIC_APPLICATION_HEALTH_POLICY_MAP
	EventsFilter               *FABRIC_HEALTH_EVENTS_FILTER
	NodesFilter                *FABRIC_NODE_HEALTH_STATES_FILTER
	ApplicationsFilter         *FABRIC_APPLICATION_HEALTH_STATES_FILTER
	Reserved                   unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_POLICY_MAP
type FABRIC_APPLICATION_HEALTH_POLICY_MAP struct {
	Count uint32
	Items *FABRIC_APPLICATION_HEALTH_POLICY_MAP_ITEM
}

// struct FABRIC_APPLICATION_HEALTH_POLICY_MAP_ITEM
type FABRIC_APPLICATION_HEALTH_POLICY_MAP_ITEM struct {
	ApplicationName win32.PWSTR
	HealthPolicy    *FABRIC_APPLICATION_HEALTH_POLICY
}

// struct FABRIC_HEALTH_EVENTS_FILTER
type FABRIC_HEALTH_EVENTS_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATES_FILTER
type FABRIC_NODE_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_STATES_FILTER
type FABRIC_APPLICATION_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_QUERY_DESCRIPTION
type FABRIC_NODE_HEALTH_QUERY_DESCRIPTION struct {
	NodeName     win32.PWSTR
	HealthPolicy *FABRIC_CLUSTER_HEALTH_POLICY
	EventsFilter *FABRIC_HEALTH_EVENTS_FILTER
	Reserved     unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_QUERY_DESCRIPTION
type FABRIC_APPLICATION_HEALTH_QUERY_DESCRIPTION struct {
	ApplicationName            win32.PWSTR
	HealthPolicy               *FABRIC_APPLICATION_HEALTH_POLICY
	EventsFilter               *FABRIC_HEALTH_EVENTS_FILTER
	ServicesFilter             *FABRIC_SERVICE_HEALTH_STATES_FILTER
	DeployedApplicationsFilter *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATES_FILTER
	Reserved                   unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATES_FILTER
type FABRIC_SERVICE_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATES_FILTER
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_QUERY_DESCRIPTION
type FABRIC_SERVICE_HEALTH_QUERY_DESCRIPTION struct {
	ServiceName      win32.PWSTR
	HealthPolicy     *FABRIC_APPLICATION_HEALTH_POLICY
	EventsFilter     *FABRIC_HEALTH_EVENTS_FILTER
	PartitionsFilter *FABRIC_PARTITION_HEALTH_STATES_FILTER
	Reserved         unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATES_FILTER
type FABRIC_PARTITION_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_QUERY_DESCRIPTION
type FABRIC_PARTITION_HEALTH_QUERY_DESCRIPTION struct {
	PartitionId    syscall.GUID
	HealthPolicy   *FABRIC_APPLICATION_HEALTH_POLICY
	EventsFilter   *FABRIC_HEALTH_EVENTS_FILTER
	ReplicasFilter *FABRIC_REPLICA_HEALTH_STATES_FILTER
	Reserved       unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATES_FILTER
type FABRIC_REPLICA_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_QUERY_DESCRIPTION
type FABRIC_REPLICA_HEALTH_QUERY_DESCRIPTION struct {
	PartitionId         syscall.GUID
	ReplicaOrInstanceId int64
	HealthPolicy        *FABRIC_APPLICATION_HEALTH_POLICY
	EventsFilter        *FABRIC_HEALTH_EVENTS_FILTER
	Reserved            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_APPLICATION_HEALTH_QUERY_DESCRIPTION struct {
	ApplicationName               win32.PWSTR
	NodeName                      win32.PWSTR
	HealthPolicy                  *FABRIC_APPLICATION_HEALTH_POLICY
	EventsFilter                  *FABRIC_HEALTH_EVENTS_FILTER
	DeployedServicePackagesFilter *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATES_FILTER
	Reserved                      unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATES_FILTER
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_QUERY_DESCRIPTION struct {
	ApplicationName     win32.PWSTR
	NodeName            win32.PWSTR
	ServiceManifestName win32.PWSTR
	HealthPolicy        *FABRIC_APPLICATION_HEALTH_POLICY
	EventsFilter        *FABRIC_HEALTH_EVENTS_FILTER
	Reserved            unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH_CHUNK_QUERY_DESCRIPTION
type FABRIC_CLUSTER_HEALTH_CHUNK_QUERY_DESCRIPTION struct {
	ClusterHealthPolicy        *FABRIC_CLUSTER_HEALTH_POLICY
	ApplicationHealthPolicyMap *FABRIC_APPLICATION_HEALTH_POLICY_MAP
	ApplicationFilters         *FABRIC_APPLICATION_HEALTH_STATE_FILTER_LIST
	NodeFilters                *FABRIC_NODE_HEALTH_STATE_FILTER_LIST
	Reserved                   unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_STATE_FILTER_LIST
type FABRIC_APPLICATION_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_HEALTH_STATE_FILTER
}

// struct FABRIC_APPLICATION_HEALTH_STATE_FILTER
type FABRIC_APPLICATION_HEALTH_STATE_FILTER struct {
	HealthStateFilter          uint32
	ApplicationNameFilter      win32.PWSTR
	ServiceFilters             *FABRIC_SERVICE_HEALTH_STATE_FILTER_LIST
	DeployedApplicationFilters *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER_LIST
	Reserved                   unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATE_FILTER_LIST
type FABRIC_SERVICE_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_HEALTH_STATE_FILTER
}

// struct FABRIC_SERVICE_HEALTH_STATE_FILTER
type FABRIC_SERVICE_HEALTH_STATE_FILTER struct {
	HealthStateFilter uint32
	ServiceNameFilter win32.PWSTR
	PartitionFilters  *FABRIC_PARTITION_HEALTH_STATE_FILTER_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATE_FILTER_LIST
type FABRIC_PARTITION_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_PARTITION_HEALTH_STATE_FILTER
}

// struct FABRIC_PARTITION_HEALTH_STATE_FILTER
type FABRIC_PARTITION_HEALTH_STATE_FILTER struct {
	HealthStateFilter uint32
	PartitionIdFilter syscall.GUID
	ReplicaFilters    *FABRIC_REPLICA_HEALTH_STATE_FILTER_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATE_FILTER_LIST
type FABRIC_REPLICA_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_REPLICA_HEALTH_STATE_FILTER
}

// struct FABRIC_REPLICA_HEALTH_STATE_FILTER
type FABRIC_REPLICA_HEALTH_STATE_FILTER struct {
	HealthStateFilter         uint32
	ReplicaOrInstanceIdFilter int64
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER_LIST
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER struct {
	HealthStateFilter             uint32
	NodeNameFilter                win32.PWSTR
	DeployedServicePackageFilters *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER_LIST
	Reserved                      unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER_LIST
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER struct {
	HealthStateFilter         uint32
	ServiceManifestNameFilter win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATE_FILTER_LIST
type FABRIC_NODE_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_NODE_HEALTH_STATE_FILTER
}

// struct FABRIC_NODE_HEALTH_STATE_FILTER
type FABRIC_NODE_HEALTH_STATE_FILTER struct {
	HealthStateFilter uint32
	NodeNameFilter    win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH_CHUNK
type FABRIC_CLUSTER_HEALTH_CHUNK struct {
	HealthState                  int32
	NodeHealthStateChunks        *FABRIC_NODE_HEALTH_STATE_CHUNK_LIST
	ApplicationHealthStateChunks *FABRIC_APPLICATION_HEALTH_STATE_CHUNK_LIST
	Reserved                     unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATE_CHUNK_LIST
type FABRIC_NODE_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_NODE_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATE_CHUNK
type FABRIC_NODE_HEALTH_STATE_CHUNK struct {
	NodeName    win32.PWSTR
	HealthState int32
	Reserved    unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_STATE_CHUNK_LIST
type FABRIC_APPLICATION_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_APPLICATION_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_STATE_CHUNK
type FABRIC_APPLICATION_HEALTH_STATE_CHUNK struct {
	ApplicationName                      win32.PWSTR
	HealthState                          int32
	ServiceHealthStateChunks             *FABRIC_SERVICE_HEALTH_STATE_CHUNK_LIST
	DeployedApplicationHealthStateChunks *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_CHUNK_LIST
	Reserved                             unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATE_CHUNK_LIST
type FABRIC_SERVICE_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_SERVICE_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATE_CHUNK
type FABRIC_SERVICE_HEALTH_STATE_CHUNK struct {
	ServiceName                win32.PWSTR
	HealthState                int32
	PartitionHealthStateChunks *FABRIC_PARTITION_HEALTH_STATE_CHUNK_LIST
	Reserved                   unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATE_CHUNK_LIST
type FABRIC_PARTITION_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_PARTITION_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATE_CHUNK
type FABRIC_PARTITION_HEALTH_STATE_CHUNK struct {
	PartitionId              syscall.GUID
	HealthState              int32
	ReplicaHealthStateChunks *FABRIC_REPLICA_HEALTH_STATE_CHUNK_LIST
	Reserved                 unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATE_CHUNK_LIST
type FABRIC_REPLICA_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_REPLICA_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATE_CHUNK
type FABRIC_REPLICA_HEALTH_STATE_CHUNK struct {
	ReplicaOrInstanceId int64
	HealthState         int32
	Reserved            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_CHUNK_LIST
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_CHUNK
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_CHUNK struct {
	NodeName                                win32.PWSTR
	HealthState                             int32
	DeployedServicePackageHealthStateChunks *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK_LIST
	Reserved                                unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK_LIST
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK struct {
	ServiceManifestName win32.PWSTR
	HealthState         int32
	Reserved            unsafe.Pointer
}

// struct FABRIC_HEALTH_REPORT_SEND_OPTIONS
type FABRIC_HEALTH_REPORT_SEND_OPTIONS struct {
	Immediate int8
	Reserved  unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_DESCRIPTION
type FABRIC_NODE_QUERY_DESCRIPTION struct {
	NodeNameFilter win32.PWSTR
	Reserved       unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_RESULT_LIST
type FABRIC_NODE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_NODE_QUERY_RESULT_ITEM
}

// struct FABRIC_NODE_QUERY_RESULT_ITEM
type FABRIC_NODE_QUERY_RESULT_ITEM struct {
	NodeName              win32.PWSTR
	IPAddressOrFQDN       win32.PWSTR
	NodeType              win32.PWSTR
	CodeVersion           win32.PWSTR
	ConfigVersion         win32.PWSTR
	NodeStatus            int32
	NodeUpTimeInSeconds   int64
	AggregatedHealthState int32
	IsSeedNode            int8
	UpgradeDomain         win32.PWSTR
	FaultDomain           win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION
type FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION struct {
	ApplicationTypeNameFilter win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct FABRIC_APPLICATION_TYPE_QUERY_RESULT_LIST
type FABRIC_APPLICATION_TYPE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM
}

// struct FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM
type FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM struct {
	ApplicationTypeName    win32.PWSTR
	ApplicationTypeVersion win32.PWSTR
	DefaultParameters      *FABRIC_APPLICATION_PARAMETER_LIST
	Reserved               unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_QUERY_DESCRIPTION
type FABRIC_SERVICE_TYPE_QUERY_DESCRIPTION struct {
	ApplicationTypeName    win32.PWSTR
	ApplicationTypeVersion win32.PWSTR
	ServiceTypeNameFilter  win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_QUERY_RESULT_LIST
type FABRIC_SERVICE_TYPE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM
type FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM struct {
	ServiceTypeDescription *FABRIC_SERVICE_TYPE_DESCRIPTION
	ServiceManifestVersion win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_DESCRIPTION
type FABRIC_SERVICE_TYPE_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_APPLICATION_QUERY_DESCRIPTION
type FABRIC_APPLICATION_QUERY_DESCRIPTION struct {
	ApplicationNameFilter win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_APPLICATION_QUERY_RESULT_LIST
type FABRIC_APPLICATION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_QUERY_RESULT_ITEM
}

// struct FABRIC_APPLICATION_QUERY_RESULT_ITEM
type FABRIC_APPLICATION_QUERY_RESULT_ITEM struct {
	ApplicationName        win32.PWSTR
	ApplicationTypeName    win32.PWSTR
	ApplicationTypeVersion win32.PWSTR
	Status                 int32
	HealthState            int32
	ApplicationParameters  *FABRIC_APPLICATION_PARAMETER_LIST
	Reserved               unsafe.Pointer
}

// struct FABRIC_SERVICE_QUERY_DESCRIPTION
type FABRIC_SERVICE_QUERY_DESCRIPTION struct {
	ApplicationName   win32.PWSTR
	ServiceNameFilter win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_QUERY_RESULT_LIST
type FABRIC_SERVICE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_QUERY_RESULT_ITEM
type FABRIC_SERVICE_QUERY_RESULT_ITEM struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SERVICE_PARTITION_QUERY_DESCRIPTION
type FABRIC_SERVICE_PARTITION_QUERY_DESCRIPTION struct {
	ServiceName       win32.PWSTR
	PartitionIdFilter syscall.GUID
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_PARTITION_QUERY_RESULT_LIST
type FABRIC_SERVICE_PARTITION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_PARTITION_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_PARTITION_QUERY_RESULT_ITEM
type FABRIC_SERVICE_PARTITION_QUERY_RESULT_ITEM struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SERVICE_REPLICA_QUERY_DESCRIPTION
type FABRIC_SERVICE_REPLICA_QUERY_DESCRIPTION struct {
	PartitionId                 syscall.GUID
	ReplicaIdOrInstanceIdFilter int64
	Reserved                    unsafe.Pointer
}

// struct FABRIC_SERVICE_REPLICA_QUERY_RESULT_LIST
type FABRIC_SERVICE_REPLICA_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_REPLICA_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_REPLICA_QUERY_RESULT_ITEM
type FABRIC_SERVICE_REPLICA_QUERY_RESULT_ITEM struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_APPLICATION_QUERY_DESCRIPTION struct {
	NodeName              win32.PWSTR
	ApplicationNameFilter win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM struct {
	ApplicationName           win32.PWSTR
	ApplicationTypeName       win32.PWSTR
	DeployedApplicationStatus int32
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_DESCRIPTION struct {
	NodeName                  win32.PWSTR
	ApplicationName           win32.PWSTR
	ServiceManifestNameFilter win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_ITEM struct {
	ServiceManifestName          win32.PWSTR
	ServiceManifestVersion       win32.PWSTR
	DeployedServicePackageStatus int32
	Reserved                     unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_DESCRIPTION struct {
	NodeName                  win32.PWSTR
	ApplicationName           win32.PWSTR
	ServiceManifestNameFilter win32.PWSTR
	ServiceTypeNameFilter     win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_ITEM struct {
	ServiceTypeName     win32.PWSTR
	CodePackageName     win32.PWSTR
	ServiceManifestName win32.PWSTR
	Status              int32
	Reserved            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_DESCRIPTION struct {
	NodeName                  win32.PWSTR
	ApplicationName           win32.PWSTR
	ServiceManifestNameFilter win32.PWSTR
	CodePackageNameFilter     win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_ITEM struct {
	CodePackageName           win32.PWSTR
	CodePackageVersion        win32.PWSTR
	ServiceManifestName       win32.PWSTR
	RunFrequencyInterval      uint32
	DeployedCodePackageStatus int32
	SetupEntryPoint           *FABRIC_CODE_PACKAGE_ENTRY_POINT
	EntryPoint                *FABRIC_CODE_PACKAGE_ENTRY_POINT
	Reserved                  unsafe.Pointer
}

// struct FABRIC_CODE_PACKAGE_ENTRY_POINT
type FABRIC_CODE_PACKAGE_ENTRY_POINT struct {
	EntryPointLocation win32.PWSTR
	ProcessId          int64
	RunAsUserName      win32.PWSTR
	EntryPointStatus   int32
	NextActivationUtc  FILETIME_
	Statistics         *FABRIC_CODE_PACKAGE_ENTRY_POINT_STATISTICS
	Reserved           unsafe.Pointer
}

// struct FABRIC_CODE_PACKAGE_ENTRY_POINT_STATISTICS
type FABRIC_CODE_PACKAGE_ENTRY_POINT_STATISTICS struct {
	LastExitCode                     uint32
	LastActivationUtc                FILETIME_
	LastExitUtc                      FILETIME_
	LastSuccessfulActivationUtc      FILETIME_
	LastSuccessfulExitUtc            FILETIME_
	ActivationCount                  uint32
	ActivationFailureCount           uint32
	ContinuousActivationFailureCount uint32
	ExitCount                        uint32
	ExitFailureCount                 uint32
	ContinuousExitFailureCount       uint32
	Reserved                         unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_DESCRIPTION struct {
	NodeName                  win32.PWSTR
	ApplicationName           win32.PWSTR
	ServiceManifestNameFilter win32.PWSTR
	PartitionIdFilter         syscall.GUID
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_ITEM struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_REPLICA_DETAIL_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_SERVICE_REPLICA_DETAIL_QUERY_DESCRIPTION struct {
	NodeName    win32.PWSTR
	PartitionId syscall.GUID
	ReplicaId   int64
	Reserved    unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_CLUSTER_LOAD_INFORMATION
type FABRIC_CLUSTER_LOAD_INFORMATION struct {
	LastBalancingStartTimeUtc FILETIME_
	LastBalancingEndTimeUtc   FILETIME_
	LoadMetricInformation     *FABRIC_LOAD_METRIC_INFORMATION_LIST
	Reserved                  unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC_INFORMATION_LIST
type FABRIC_LOAD_METRIC_INFORMATION_LIST struct {
	Count uint32
	Items *FABRIC_LOAD_METRIC_INFORMATION
}

// struct FABRIC_LOAD_METRIC_INFORMATION
type FABRIC_LOAD_METRIC_INFORMATION struct {
	Name               win32.PWSTR
	IsBalancedBefore   int8
	IsBalancedAfter    int8
	DeviationBefore    float64
	DeviationAfter     float64
	BalancingThreshold float64
	Action             win32.PWSTR
	Reserved           unsafe.Pointer
}

// struct FABRIC_PARTITION_LOAD_INFORMATION_QUERY_DESCRIPTION
type FABRIC_PARTITION_LOAD_INFORMATION_QUERY_DESCRIPTION struct {
	PartitionId syscall.GUID
	Reserved    unsafe.Pointer
}

// struct FABRIC_PARTITION_LOAD_INFORMATION
type FABRIC_PARTITION_LOAD_INFORMATION struct {
	PartitionId                syscall.GUID
	PrimaryLoadMetricReports   *FABRIC_LOAD_METRIC_REPORT_LIST
	SecondaryLoadMetricReports *FABRIC_LOAD_METRIC_REPORT_LIST
	Reserved                   unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC_REPORT_LIST
type FABRIC_LOAD_METRIC_REPORT_LIST struct {
	Count uint32
	Items *FABRIC_LOAD_METRIC_REPORT
}

// struct FABRIC_LOAD_METRIC_REPORT
type FABRIC_LOAD_METRIC_REPORT struct {
	Name            win32.PWSTR
	Value           uint32
	LastReportedUtc FILETIME_
	Reserved        unsafe.Pointer
}

// struct FABRIC_PROVISIONED_CODE_VERSION_QUERY_DESCRIPTION
type FABRIC_PROVISIONED_CODE_VERSION_QUERY_DESCRIPTION struct {
	CodeVersionFilter win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_LIST
type FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_ITEM
}

// struct FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_ITEM
type FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_ITEM struct {
	CodeVersion win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_DESCRIPTION
type FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_DESCRIPTION struct {
	ConfigVersionFilter win32.PWSTR
	Reserved            unsafe.Pointer
}

// struct FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_LIST
type FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_ITEM
}

// struct FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_ITEM
type FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_ITEM struct {
	ConfigVersion win32.PWSTR
	Reserved      unsafe.Pointer
}

// struct FABRIC_NODE_LOAD_INFORMATION_QUERY_DESCRIPTION
type FABRIC_NODE_LOAD_INFORMATION_QUERY_DESCRIPTION struct {
	NodeName win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_NODE_LOAD_INFORMATION
type FABRIC_NODE_LOAD_INFORMATION struct {
	NodeName                  win32.PWSTR
	NodeLoadMetricInformation *FABRIC_NODE_LOAD_METRIC_INFORMATION_LIST
	Reserved                  unsafe.Pointer
}

// struct FABRIC_NODE_LOAD_METRIC_INFORMATION_LIST
type FABRIC_NODE_LOAD_METRIC_INFORMATION_LIST struct {
	Count uint32
	Items *FABRIC_NODE_LOAD_METRIC_INFORMATION
}

// struct FABRIC_NODE_LOAD_METRIC_INFORMATION
type FABRIC_NODE_LOAD_METRIC_INFORMATION struct {
	Name                  win32.PWSTR
	NodeCapacity          int64
	NodeLoad              int64
	NodeRemainingCapacity int64
	IsCapacityViolation   int8
	Reserved              unsafe.Pointer
}

// struct FABRIC_REPLICA_LOAD_INFORMATION_QUERY_DESCRIPTION
type FABRIC_REPLICA_LOAD_INFORMATION_QUERY_DESCRIPTION struct {
	PartitionId         syscall.GUID
	ReplicaOrInstanceId int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_REPLICA_LOAD_INFORMATION
type FABRIC_REPLICA_LOAD_INFORMATION struct {
	PartitionId         syscall.GUID
	ReplicaOrInstanceId int64
	LoadMetricReports   *FABRIC_LOAD_METRIC_REPORT_LIST
	Reserved            unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_QUERY_DESCRIPTION
type FABRIC_SERVICE_GROUP_MEMBER_QUERY_DESCRIPTION struct {
	ApplicationName   win32.PWSTR
	ServiceNameFilter win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_LIST
type FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_ITEM
type FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_ITEM struct {
	ServiceName win32.PWSTR
	Members     *FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_LIST
	Reserved    unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_LIST
type FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_ITEM
type FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_ITEM struct {
	ServiceType win32.PWSTR
	ServiceName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_DESCRIPTION
type FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_DESCRIPTION struct {
	ApplicationTypeName        win32.PWSTR
	ApplicationTypeVersion     win32.PWSTR
	ServiceGroupTypeNameFilter win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_LIST
type FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_ITEM
type FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_ITEM struct {
	ServiceGroupMemberTypeDescription *FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION_LIST
	ServiceManifestVersion            win32.PWSTR
	ServiceManifestName               win32.PWSTR
	Reserved                          unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION_LIST
type FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION
}

// struct FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION
type FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION struct {
	ServiceTypeName win32.PWSTR
	LoadMetrics     *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION_LIST
	Reserved        unsafe.Pointer
}

// struct FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION_LIST
type FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
}

// struct FABRIC_UNPLACED_REPLICA_INFORMATION_QUERY_DESCRIPTION
type FABRIC_UNPLACED_REPLICA_INFORMATION_QUERY_DESCRIPTION struct {
	ServiceName        win32.PWSTR
	PartitionId        syscall.GUID
	OnlyQueryPrimaries int8
	Reserved           unsafe.Pointer
}

// struct FABRIC_UNPLACED_REPLICA_INFORMATION
type FABRIC_UNPLACED_REPLICA_INFORMATION struct {
	ServiceName            win32.PWSTR
	PartitionId            syscall.GUID
	UnplacedReplicaReasons *FABRIC_STRING_LIST
	Reserved               unsafe.Pointer
}

// struct FABRIC_STRING_LIST
type FABRIC_STRING_LIST struct {
	Count uint32
	Items *win32.PWSTR
}

// struct FABRIC_PAGING_STATUS
type FABRIC_PAGING_STATUS struct {
	ContinuationToken win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_APPLICATION_LOAD_INFORMATION_QUERY_DESCRIPTION
type FABRIC_APPLICATION_LOAD_INFORMATION_QUERY_DESCRIPTION struct {
	ApplicationName win32.PWSTR
	Reserved        unsafe.Pointer
}

// struct FABRIC_APPLICATION_LOAD_INFORMATION
type FABRIC_APPLICATION_LOAD_INFORMATION struct {
	Name                             win32.PWSTR
	MinimumNodes                     uint32
	MaximumNodes                     uint32
	NodeCount                        uint32
	ApplicationLoadMetricInformation *FABRIC_APPLICATION_LOAD_METRIC_INFORMATION_LIST
	Reserved                         unsafe.Pointer
}

// struct FABRIC_APPLICATION_LOAD_METRIC_INFORMATION_LIST
type FABRIC_APPLICATION_LOAD_METRIC_INFORMATION_LIST struct {
	Count       uint32
	LoadMetrics *FABRIC_APPLICATION_LOAD_METRIC_INFORMATION
	Reserved    unsafe.Pointer
}

// struct FABRIC_APPLICATION_LOAD_METRIC_INFORMATION
type FABRIC_APPLICATION_LOAD_METRIC_INFORMATION struct {
	Name                win32.PWSTR
	ReservationCapacity int64
	ApplicationCapacity int64
	ApplicationLoad     int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_SERVICE_NAME_QUERY_DESCRIPTION
type FABRIC_SERVICE_NAME_QUERY_DESCRIPTION struct {
	PartitionId syscall.GUID
	Reserved    unsafe.Pointer
}

// struct FABRIC_SERVICE_NAME_QUERY_RESULT
type FABRIC_SERVICE_NAME_QUERY_RESULT struct {
	ServiceName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_APPLICATION_NAME_QUERY_DESCRIPTION
type FABRIC_APPLICATION_NAME_QUERY_DESCRIPTION struct {
	ServiceName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_APPLICATION_NAME_QUERY_RESULT
type FABRIC_APPLICATION_NAME_QUERY_RESULT struct {
	ApplicationName win32.PWSTR
	Reserved        unsafe.Pointer
}

// struct PAGED_FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION
type PAGED_FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION struct {
	ApplicationTypeNameFilter    win32.PWSTR
	MaxResults                   int32
	ContinuationToken            win32.PWSTR
	ExcludeApplicationParameters int8
	Reserved                     unsafe.Pointer
}

// struct FABRIC_PAGED_DEPLOYED_APPLICATION_QUERY_DESCRIPTION
type FABRIC_PAGED_DEPLOYED_APPLICATION_QUERY_DESCRIPTION struct {
	NodeName              win32.PWSTR
	ApplicationNameFilter win32.PWSTR
	IncludeHealthState    int8
	PagingDescription     *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved              unsafe.Pointer
}

// struct FABRIC_QUERY_PAGING_DESCRIPTION
type FABRIC_QUERY_PAGING_DESCRIPTION struct {
	ContinuationToken win32.PWSTR
	MaxResults        int32
	Reserved          unsafe.Pointer
}

// struct FABRIC_REPAIR_TASK
type FABRIC_REPAIR_TASK struct {
	Scope         *FABRIC_REPAIR_SCOPE_IDENTIFIER
	TaskId        win32.PWSTR
	Version       int64
	Description   win32.PWSTR
	State         int32
	Flags         uint32
	Action        win32.PWSTR
	Target        *FABRIC_REPAIR_TARGET_DESCRIPTION
	ExecutorState *FABRIC_REPAIR_EXECUTOR_STATE
	Impact        *FABRIC_REPAIR_IMPACT_DESCRIPTION
	Result        *FABRIC_REPAIR_RESULT_DESCRIPTION
	History       *FABRIC_REPAIR_TASK_HISTORY
	Reserved      unsafe.Pointer
}

// struct FABRIC_REPAIR_SCOPE_IDENTIFIER
type FABRIC_REPAIR_SCOPE_IDENTIFIER struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_REPAIR_TARGET_DESCRIPTION
type FABRIC_REPAIR_TARGET_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_REPAIR_EXECUTOR_STATE
type FABRIC_REPAIR_EXECUTOR_STATE struct {
	Executor     win32.PWSTR
	ExecutorData win32.PWSTR
	Reserved     unsafe.Pointer
}

// struct FABRIC_REPAIR_IMPACT_DESCRIPTION
type FABRIC_REPAIR_IMPACT_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_REPAIR_RESULT_DESCRIPTION
type FABRIC_REPAIR_RESULT_DESCRIPTION struct {
	ResultStatus  int32
	ResultCode    win32.HRESULT
	ResultDetails win32.PWSTR
	Reserved      unsafe.Pointer
}

// struct FABRIC_REPAIR_TASK_HISTORY
type FABRIC_REPAIR_TASK_HISTORY struct {
	CreatedUtcTimestamp   FILETIME_
	ClaimedUtcTimestamp   FILETIME_
	PreparingUtcTimestamp FILETIME_
	ApprovedUtcTimestamp  FILETIME_
	ExecutingUtcTimestamp FILETIME_
	RestoringUtcTimestamp FILETIME_
	CompletedUtcTimestamp FILETIME_
	Reserved              unsafe.Pointer
}

// struct FABRIC_REPAIR_CANCEL_DESCRIPTION
type FABRIC_REPAIR_CANCEL_DESCRIPTION struct {
	Scope        *FABRIC_REPAIR_SCOPE_IDENTIFIER
	RepairTaskId win32.PWSTR
	Version      int64
	RequestAbort int8
	Reserved     unsafe.Pointer
}

// struct FABRIC_REPAIR_APPROVE_DESCRIPTION
type FABRIC_REPAIR_APPROVE_DESCRIPTION struct {
	Scope        *FABRIC_REPAIR_SCOPE_IDENTIFIER
	RepairTaskId win32.PWSTR
	Version      int64
	Reserved     unsafe.Pointer
}

// struct FABRIC_REPAIR_DELETE_DESCRIPTION
type FABRIC_REPAIR_DELETE_DESCRIPTION struct {
	Scope        *FABRIC_REPAIR_SCOPE_IDENTIFIER
	RepairTaskId win32.PWSTR
	Version      int64
	Reserved     unsafe.Pointer
}

// struct FABRIC_REPAIR_TASK_QUERY_DESCRIPTION
type FABRIC_REPAIR_TASK_QUERY_DESCRIPTION struct {
	Scope          *FABRIC_REPAIR_SCOPE_IDENTIFIER
	TaskIdFilter   win32.PWSTR
	StateFilter    uint32
	ExecutorFilter win32.PWSTR
	Reserved       unsafe.Pointer
}

// struct FABRIC_REPAIR_TASK_LIST
type FABRIC_REPAIR_TASK_LIST struct {
	Count uint32
	Items *FABRIC_REPAIR_TASK
}

// struct FABRIC_REPAIR_TASK_HEALTH_POLICY_UPDATE_DESCRIPTION
type FABRIC_REPAIR_TASK_HEALTH_POLICY_UPDATE_DESCRIPTION struct {
	Scope                       *FABRIC_REPAIR_SCOPE_IDENTIFIER
	RepairTaskId                win32.PWSTR
	Version                     int64
	Flags                       uint32
	PerformPreparingHealthCheck int8
	PerformRestoringHealthCheck int8
	Reserved                    unsafe.Pointer
}

// struct FABRIC_START_PARTITION_DATA_LOSS_DESCRIPTION
type FABRIC_START_PARTITION_DATA_LOSS_DESCRIPTION struct {
	OperationId       syscall.GUID
	PartitionSelector *FABRIC_PARTITION_SELECTOR
	DataLossMode      int32
	Reserved          unsafe.Pointer
}

// struct FABRIC_PARTITION_SELECTOR
type FABRIC_PARTITION_SELECTOR struct {
	ServiceName           win32.PWSTR
	PartitionSelectorType int32
	PartitionKey          win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_PARTITION_DATA_LOSS_PROGRESS
type FABRIC_PARTITION_DATA_LOSS_PROGRESS struct {
	State    int32
	Result   *FABRIC_PARTITION_DATA_LOSS_RESULT
	Reserved unsafe.Pointer
}

// struct FABRIC_PARTITION_DATA_LOSS_RESULT
type FABRIC_PARTITION_DATA_LOSS_RESULT struct {
	SelectedPartition *FABRIC_SELECTED_PARTITION
	ErrorCode         win32.HRESULT
	Reserved          unsafe.Pointer
}

// struct FABRIC_SELECTED_PARTITION
type FABRIC_SELECTED_PARTITION struct {
	ServiceName win32.PWSTR
	PartitionId syscall.GUID
	Reserved    unsafe.Pointer
}

// struct FABRIC_START_PARTITION_QUORUM_LOSS_DESCRIPTION
type FABRIC_START_PARTITION_QUORUM_LOSS_DESCRIPTION struct {
	OperationId                      syscall.GUID
	PartitionSelector                *FABRIC_PARTITION_SELECTOR
	QuorumLossMode                   int32
	QuorumLossDurationInMilliSeconds int32
	Reserved                         unsafe.Pointer
}

// struct FABRIC_PARTITION_QUORUM_LOSS_PROGRESS
type FABRIC_PARTITION_QUORUM_LOSS_PROGRESS struct {
	State    int32
	Result   *FABRIC_PARTITION_QUORUM_LOSS_RESULT
	Reserved unsafe.Pointer
}

// struct FABRIC_PARTITION_QUORUM_LOSS_RESULT
type FABRIC_PARTITION_QUORUM_LOSS_RESULT struct {
	SelectedPartition *FABRIC_SELECTED_PARTITION
	ErrorCode         win32.HRESULT
	Reserved          unsafe.Pointer
}

// struct FABRIC_START_PARTITION_RESTART_DESCRIPTION
type FABRIC_START_PARTITION_RESTART_DESCRIPTION struct {
	OperationId          syscall.GUID
	PartitionSelector    *FABRIC_PARTITION_SELECTOR
	RestartPartitionMode int32
	Reserved             unsafe.Pointer
}

// struct FABRIC_PARTITION_RESTART_PROGRESS
type FABRIC_PARTITION_RESTART_PROGRESS struct {
	State    int32
	Result   *FABRIC_PARTITION_RESTART_RESULT
	Reserved unsafe.Pointer
}

// struct FABRIC_PARTITION_RESTART_RESULT
type FABRIC_PARTITION_RESTART_RESULT struct {
	SelectedPartition *FABRIC_SELECTED_PARTITION
	ErrorCode         win32.HRESULT
	Reserved          unsafe.Pointer
}

// struct FABRIC_TEST_COMMAND_LIST_DESCRIPTION
type FABRIC_TEST_COMMAND_LIST_DESCRIPTION struct {
	TestCommandStateFilter int32
	TestCommandTypeFilter  int32
	Reserved               unsafe.Pointer
}

// struct TEST_COMMAND_QUERY_RESULT_LIST
type TEST_COMMAND_QUERY_RESULT_LIST struct {
	Count uint32
	Items unsafe.Pointer
}

// struct FABRIC_CANCEL_TEST_COMMAND_DESCRIPTION
type FABRIC_CANCEL_TEST_COMMAND_DESCRIPTION struct {
	OperationId syscall.GUID
	Force       int8
	Reserved    unsafe.Pointer
}

// struct FABRIC_START_CHAOS_DESCRIPTION
type FABRIC_START_CHAOS_DESCRIPTION struct {
	ChaosParameters *FABRIC_CHAOS_PARAMETERS
	Reserved        unsafe.Pointer
}

// struct FABRIC_CHAOS_PARAMETERS
type FABRIC_CHAOS_PARAMETERS struct {
	MaxClusterStabilizationTimeoutInSeconds uint32
	MaxConcurrentFaults                     uint32
	EnableMoveReplicaFaults                 int8
	TimeToRunInSeconds                      uint64
	WaitTimeBetweenIterationsInSeconds      uint32
	WaitTimeBetweenFaultsInSeconds          uint32
	Context                                 *FABRIC_EVENT_CONTEXT_MAP
	Reserved                                unsafe.Pointer
}

// struct FABRIC_EVENT_CONTEXT_MAP
type FABRIC_EVENT_CONTEXT_MAP struct {
	Count uint32
	Items *FABRIC_EVENT_CONTEXT_MAP_ITEM
}

// struct FABRIC_EVENT_CONTEXT_MAP_ITEM
type FABRIC_EVENT_CONTEXT_MAP_ITEM struct {
	Key   win32.PWSTR
	Value win32.PWSTR
}

// struct FABRIC_GET_CHAOS_REPORT_DESCRIPTION
type FABRIC_GET_CHAOS_REPORT_DESCRIPTION struct {
	Filter            *FABRIC_CHAOS_REPORT_FILTER
	ContinuationToken win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_CHAOS_REPORT_FILTER
type FABRIC_CHAOS_REPORT_FILTER struct {
	StartTimeUtc FILETIME_
	EndTimeUtc   FILETIME_
	Reserved     unsafe.Pointer
}

// struct FABRIC_CHAOS_REPORT
type FABRIC_CHAOS_REPORT struct {
	ChaosParameters   *FABRIC_CHAOS_PARAMETERS
	Status            int32
	ContinuationToken win32.PWSTR
	History           *FABRIC_CHAOS_EVENT_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_CHAOS_EVENT_LIST
type FABRIC_CHAOS_EVENT_LIST struct {
	Count uint32
	Items *FABRIC_CHAOS_EVENT
}

// struct FABRIC_CHAOS_EVENT
type FABRIC_CHAOS_EVENT struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_NODE_TRANSITION_DESCRIPTION
type FABRIC_NODE_TRANSITION_DESCRIPTION struct {
	NodeTransitionType int32
	Value              unsafe.Pointer
}

// struct FABRIC_NODE_TRANSITION_PROGRESS
type FABRIC_NODE_TRANSITION_PROGRESS struct {
	State    int32
	Result   *FABRIC_NODE_TRANSITION_RESULT
	Reserved unsafe.Pointer
}

// struct FABRIC_NODE_TRANSITION_RESULT
type FABRIC_NODE_TRANSITION_RESULT struct {
	ErrorCode  win32.HRESULT
	NodeResult *FABRIC_NODE_RESULT
	Reserved   unsafe.Pointer
}

// struct FABRIC_NODE_RESULT
type FABRIC_NODE_RESULT struct {
	NodeName     win32.PWSTR
	NodeInstance uint64
	Reserved     unsafe.Pointer
}

// struct FABRIC_RESTART_NODE_DESCRIPTION2
type FABRIC_RESTART_NODE_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_START_NODE_DESCRIPTION2
type FABRIC_START_NODE_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_STOP_NODE_DESCRIPTION2
type FABRIC_STOP_NODE_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION2
type FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_DEPLOYED_CODE_PACKAGE_RESULT
type FABRIC_DEPLOYED_CODE_PACKAGE_RESULT struct {
	NodeName              win32.PWSTR
	ApplicationName       win32.PWSTR
	ServiceManifestName   win32.PWSTR
	CodePackageName       win32.PWSTR
	CodePackageInstanceId int64
	Reserved              unsafe.Pointer
}

// struct FABRIC_MOVE_PRIMARY_DESCRIPTION2
type FABRIC_MOVE_PRIMARY_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_MOVE_PRIMARY_RESULT
type FABRIC_MOVE_PRIMARY_RESULT struct {
	NodeName    win32.PWSTR
	ServiceName win32.PWSTR
	PartitionId syscall.GUID
	Reserved    unsafe.Pointer
}

// struct FABRIC_MOVE_SECONDARY_DESCRIPTION2
type FABRIC_MOVE_SECONDARY_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_MOVE_SECONDARY_RESULT
type FABRIC_MOVE_SECONDARY_RESULT struct {
	CurrentNodeName win32.PWSTR
	NewNodeName     win32.PWSTR
	ServiceName     win32.PWSTR
	PartitionId     syscall.GUID
	Reserved        unsafe.Pointer
}

// struct FABRIC_NETWORK_DESCRIPTION
type FABRIC_NETWORK_DESCRIPTION struct {
	NetworkType int32
	Value       unsafe.Pointer
}

// struct FABRIC_DELETE_NETWORK_DESCRIPTION
type FABRIC_DELETE_NETWORK_DESCRIPTION struct {
	NetworkName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_NETWORK_QUERY_DESCRIPTION
type FABRIC_NETWORK_QUERY_DESCRIPTION struct {
	NetworkNameFilter   win32.PWSTR
	NetworkStatusFilter uint32
	PagingDescription   *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved            unsafe.Pointer
}

// struct FABRIC_NETWORK_QUERY_RESULT_LIST
type FABRIC_NETWORK_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_NETWORK_INFORMATION
}

// struct FABRIC_NETWORK_INFORMATION
type FABRIC_NETWORK_INFORMATION struct {
	NetworkType int32
	Value       unsafe.Pointer
}

// struct FABRIC_NETWORK_APPLICATION_QUERY_DESCRIPTION
type FABRIC_NETWORK_APPLICATION_QUERY_DESCRIPTION struct {
	NetworkName           win32.PWSTR
	ApplicationNameFilter win32.PWSTR
	PagingDescription     *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved              unsafe.Pointer
}

// struct FABRIC_NETWORK_APPLICATION_QUERY_RESULT_LIST
type FABRIC_NETWORK_APPLICATION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_NETWORK_APPLICATION_QUERY_RESULT_ITEM
}

// struct FABRIC_NETWORK_APPLICATION_QUERY_RESULT_ITEM
type FABRIC_NETWORK_APPLICATION_QUERY_RESULT_ITEM struct {
	ApplicationName win32.PWSTR
	Reserved        unsafe.Pointer
}

// struct FABRIC_NETWORK_NODE_QUERY_DESCRIPTION
type FABRIC_NETWORK_NODE_QUERY_DESCRIPTION struct {
	NetworkName       win32.PWSTR
	NodeNameFilter    win32.PWSTR
	PagingDescription *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved          unsafe.Pointer
}

// struct FABRIC_NETWORK_NODE_QUERY_RESULT_LIST
type FABRIC_NETWORK_NODE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_NETWORK_NODE_QUERY_RESULT_ITEM
}

// struct FABRIC_NETWORK_NODE_QUERY_RESULT_ITEM
type FABRIC_NETWORK_NODE_QUERY_RESULT_ITEM struct {
	NodeName win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_APPLICATION_NETWORK_QUERY_DESCRIPTION
type FABRIC_APPLICATION_NETWORK_QUERY_DESCRIPTION struct {
	ApplicationName   win32.PWSTR
	PagingDescription *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved          unsafe.Pointer
}

// struct FABRIC_APPLICATION_NETWORK_QUERY_RESULT_LIST
type FABRIC_APPLICATION_NETWORK_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_NETWORK_QUERY_RESULT_ITEM
}

// struct FABRIC_APPLICATION_NETWORK_QUERY_RESULT_ITEM
type FABRIC_APPLICATION_NETWORK_QUERY_RESULT_ITEM struct {
	NetworkName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_DEPLOYED_NETWORK_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_NETWORK_QUERY_DESCRIPTION struct {
	NodeName          win32.PWSTR
	PagingDescription *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved          unsafe.Pointer
}

// struct FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_ITEM struct {
	NetworkName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_DESCRIPTION struct {
	NodeName                  win32.PWSTR
	NetworkName               win32.PWSTR
	ApplicationNameFilter     win32.PWSTR
	ServiceManifestNameFilter win32.PWSTR
	CodePackageNameFilter     win32.PWSTR
	PagingDescription         *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_ITEM struct {
	ApplicationName            win32.PWSTR
	NetworkName                win32.PWSTR
	CodePackageName            win32.PWSTR
	CodePackageVersion         win32.PWSTR
	ServiceManifestName        win32.PWSTR
	ServicePackageActivationId win32.PWSTR
	ContainerAddress           win32.PWSTR
	ContainerId                win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_APPLICATION_UPGRADE_PROGRESS
type FABRIC_APPLICATION_UPGRADE_PROGRESS struct {
	UpgradeDescription                    *FABRIC_APPLICATION_UPGRADE_DESCRIPTION
	UpgradeState                          int32
	UpgradeMode                           int32
	NextUpgradeDomain                     win32.PWSTR
	UpgradeDomains                        *FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION_LIST
	UpgradeDurationInSeconds              uint32
	CurrentUpgradeDomainDurationInSeconds uint32
	UnhealthyEvaluations                  *FABRIC_HEALTH_EVALUATION_LIST
	CurrentUpgradeDomainProgress          *FABRIC_UPGRADE_DOMAIN_PROGRESS
	Reserved                              unsafe.Pointer
}

// struct FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION_LIST
type FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION
}

// struct FABRIC_HEALTH_EVALUATION_LIST
type FABRIC_HEALTH_EVALUATION_LIST struct {
	Count uint32
	Items *FABRIC_HEALTH_EVALUATION
}

// struct FABRIC_HEALTH_EVALUATION
type FABRIC_HEALTH_EVALUATION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_UPGRADE_DOMAIN_PROGRESS
type FABRIC_UPGRADE_DOMAIN_PROGRESS struct {
	UpgradeDomainName win32.PWSTR
	NodeProgressList  *FABRIC_NODE_UPGRADE_PROGRESS_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_NODE_UPGRADE_PROGRESS_LIST
type FABRIC_NODE_UPGRADE_PROGRESS_LIST struct {
	Count uint32
	Items *FABRIC_NODE_UPGRADE_PROGRESS
}

// struct FABRIC_NODE_UPGRADE_PROGRESS
type FABRIC_NODE_UPGRADE_PROGRESS struct {
	NodeName            win32.PWSTR
	UpgradePhase        int32
	PendingSafetyChecks *FABRIC_UPGRADE_SAFETY_CHECK_LIST
	Reserved            unsafe.Pointer
}

// struct FABRIC_UPGRADE_SAFETY_CHECK_LIST
type FABRIC_UPGRADE_SAFETY_CHECK_LIST struct {
	Count uint32
	Items *FABRIC_UPGRADE_SAFETY_CHECK
}

// struct FABRIC_UPGRADE_SAFETY_CHECK
type FABRIC_UPGRADE_SAFETY_CHECK struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_UPGRADE_PROGRESS
type FABRIC_UPGRADE_PROGRESS struct {
	UpgradeDescription                    *FABRIC_UPGRADE_DESCRIPTION
	UpgradeState                          int32
	UpgradeMode                           int32
	NextUpgradeDomain                     win32.PWSTR
	UpgradeDomains                        *FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION_LIST
	UpgradeDurationInSeconds              uint32
	CurrentUpgradeDomainDurationInSeconds uint32
	UnhealthyEvaluations                  *FABRIC_HEALTH_EVALUATION_LIST
	CurrentUpgradeDomainProgress          *FABRIC_UPGRADE_DOMAIN_PROGRESS
	Reserved                              unsafe.Pointer
}

// struct FABRIC_ROLLING_UPGRADE_MONITORING_POLICY
type FABRIC_ROLLING_UPGRADE_MONITORING_POLICY struct {
	FailureAction                    int32
	HealthCheckWaitDurationInSeconds uint32
	HealthCheckRetryTimeoutInSeconds uint32
	UpgradeTimeoutInSeconds          uint32
	UpgradeDomainTimeoutInSeconds    uint32
	Reserved                         unsafe.Pointer
}

// struct FABRIC_CHAOS_DESCRIPTION
type FABRIC_CHAOS_DESCRIPTION struct {
	ChaosParameters *FABRIC_CHAOS_PARAMETERS
	Status          int32
	ScheduleStatus  int32
	Reserved        unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_DESCRIPTION
type FABRIC_CHAOS_SCHEDULE_DESCRIPTION struct {
	Version  uint32
	Schedule *FABRIC_CHAOS_SCHEDULE
	Reserved unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE
type FABRIC_CHAOS_SCHEDULE struct {
	StartDate          FILETIME_
	ExpiryDate         FILETIME_
	ChaosParametersMap *FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP
	Jobs               *FABRIC_CHAOS_SCHEDULE_JOB_LIST
	Reserved           unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP
type FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP struct {
	Count uint32
	Items *FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP_ITEM
}

// struct FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP_ITEM
type FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP_ITEM struct {
	Name       win32.PWSTR
	Parameters *FABRIC_CHAOS_PARAMETERS
}

// struct FABRIC_CHAOS_SCHEDULE_JOB_LIST
type FABRIC_CHAOS_SCHEDULE_JOB_LIST struct {
	Count uint32
	Items *FABRIC_CHAOS_SCHEDULE_JOB
}

// struct FABRIC_CHAOS_SCHEDULE_JOB
type FABRIC_CHAOS_SCHEDULE_JOB struct {
	ChaosParameters win32.PWSTR
	Days            *FABRIC_CHAOS_SCHEDULE_JOB_ACTIVE_DAYS
	Times           *FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC_LIST
	Reserved        unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_JOB_ACTIVE_DAYS
type FABRIC_CHAOS_SCHEDULE_JOB_ACTIVE_DAYS struct {
	Sunday    int8
	Monday    int8
	Tuesday   int8
	Wednesday int8
	Thursday  int8
	Friday    int8
	Saturday  int8
	Reserved  unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC_LIST
type FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC_LIST struct {
	Count uint32
	Items *FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC
}

// struct FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC
type FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC struct {
	StartTime *FABRIC_CHAOS_SCHEDULE_TIME_UTC
	EndTime   *FABRIC_CHAOS_SCHEDULE_TIME_UTC
	Reserved  unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_TIME_UTC
type FABRIC_CHAOS_SCHEDULE_TIME_UTC struct {
	Hour     uint32
	Minute   uint32
	Reserved unsafe.Pointer
}

// struct FABRIC_CHAOS_EVENTS_SEGMENT
type FABRIC_CHAOS_EVENTS_SEGMENT struct {
	ContinuationToken win32.PWSTR
	History           *FABRIC_CHAOS_EVENT_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_SECRET_REFERENCE_LIST
type FABRIC_SECRET_REFERENCE_LIST struct {
	Count uint32
	Items *FABRIC_SECRET_REFERENCE
}

// struct FABRIC_SECRET_REFERENCE
type FABRIC_SECRET_REFERENCE struct {
	Name    win32.PWSTR
	Version win32.PWSTR
}

// struct FABRIC_SECRET_LIST
type FABRIC_SECRET_LIST struct {
	Count uint32
	Items *FABRIC_SECRET
}

// struct FABRIC_SECRET
type FABRIC_SECRET struct {
	Name        win32.PWSTR
	Version     win32.PWSTR
	Value       win32.PWSTR
	Kind        win32.PWSTR
	ContentType win32.PWSTR
}
