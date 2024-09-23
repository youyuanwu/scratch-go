package fabrictypes

import (
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// struct __MIDL___MIDL_itf_FabricTypes_0003_0001_0001
type MIDL___MIDL_itf_FabricTypes_0003_0001_0001__ struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// struct FABRIC_NODE_ID
type FABRIC_NODE_ID struct {
	Low      uint64
	High     uint64
	Reserved unsafe.Pointer
}

// struct FABRIC_OPERATION_DATA_BUFFER
type FABRIC_OPERATION_DATA_BUFFER struct {
	BufferSize uint32
	Buffer     *byte
}

// struct FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
type FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION struct {
	Name                 win32.PWSTR
	Weight               int32
	PrimaryDefaultLoad   uint32
	SecondaryDefaultLoad uint32
	Reserved             unsafe.Pointer
}

// struct FABRIC_PARTITION_SELECTOR
type FABRIC_PARTITION_SELECTOR struct {
	ServiceName           win32.PWSTR
	PartitionSelectorType int32
	PartitionKey          win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_START_PARTITION_DATA_LOSS_DESCRIPTION
type FABRIC_START_PARTITION_DATA_LOSS_DESCRIPTION struct {
	OperationId       syscall.GUID
	PartitionSelector *FABRIC_PARTITION_SELECTOR
	DataLossMode      int32
	Reserved          unsafe.Pointer
}

// struct FABRIC_SELECTED_PARTITION
type FABRIC_SELECTED_PARTITION struct {
	ServiceName win32.PWSTR
	PartitionId syscall.GUID
	Reserved    unsafe.Pointer
}

// struct FABRIC_PARTITION_DATA_LOSS_RESULT
type FABRIC_PARTITION_DATA_LOSS_RESULT struct {
	SelectedPartition *FABRIC_SELECTED_PARTITION
	ErrorCode         win32.HRESULT
	Reserved          unsafe.Pointer
}

// struct FABRIC_PARTITION_DATA_LOSS_PROGRESS
type FABRIC_PARTITION_DATA_LOSS_PROGRESS struct {
	State    int32
	Result   *FABRIC_PARTITION_DATA_LOSS_RESULT
	Reserved unsafe.Pointer
}

// struct FABRIC_START_PARTITION_QUORUM_LOSS_DESCRIPTION
type FABRIC_START_PARTITION_QUORUM_LOSS_DESCRIPTION struct {
	OperationId                      syscall.GUID
	PartitionSelector                *FABRIC_PARTITION_SELECTOR
	QuorumLossMode                   int32
	QuorumLossDurationInMilliSeconds int32
	Reserved                         unsafe.Pointer
}

// struct FABRIC_PARTITION_QUORUM_LOSS_RESULT
type FABRIC_PARTITION_QUORUM_LOSS_RESULT struct {
	SelectedPartition *FABRIC_SELECTED_PARTITION
	ErrorCode         win32.HRESULT
	Reserved          unsafe.Pointer
}

// struct FABRIC_PARTITION_QUORUM_LOSS_PROGRESS
type FABRIC_PARTITION_QUORUM_LOSS_PROGRESS struct {
	State    int32
	Result   *FABRIC_PARTITION_QUORUM_LOSS_RESULT
	Reserved unsafe.Pointer
}

// struct FABRIC_START_PARTITION_RESTART_DESCRIPTION
type FABRIC_START_PARTITION_RESTART_DESCRIPTION struct {
	OperationId          syscall.GUID
	PartitionSelector    *FABRIC_PARTITION_SELECTOR
	RestartPartitionMode int32
	Reserved             unsafe.Pointer
}

// struct FABRIC_ORCHESTRATION_UPGRADE_PROGRESS
type FABRIC_ORCHESTRATION_UPGRADE_PROGRESS struct {
	UpgradeState   int32
	ProgressStatus uint32
	Reserved       unsafe.Pointer
}

// struct FABRIC_ORCHESTRATION_UPGRADE_PROGRESS_EX1
type FABRIC_ORCHESTRATION_UPGRADE_PROGRESS_EX1 struct {
	ConfigVersion win32.PWSTR
	Reserved      unsafe.Pointer
}

// struct FABRIC_ORCHESTRATION_UPGRADE_PROGRESS_EX2
type FABRIC_ORCHESTRATION_UPGRADE_PROGRESS_EX2 struct {
	Details  win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_PARTITION_RESTART_RESULT
type FABRIC_PARTITION_RESTART_RESULT struct {
	SelectedPartition *FABRIC_SELECTED_PARTITION
	ErrorCode         win32.HRESULT
	Reserved          unsafe.Pointer
}

// struct FABRIC_PARTITION_RESTART_PROGRESS
type FABRIC_PARTITION_RESTART_PROGRESS struct {
	State    int32
	Result   *FABRIC_PARTITION_RESTART_RESULT
	Reserved unsafe.Pointer
}

// struct FABRIC_CANCEL_TEST_COMMAND_DESCRIPTION
type FABRIC_CANCEL_TEST_COMMAND_DESCRIPTION struct {
	OperationId syscall.GUID
	Force       int8
	Reserved    unsafe.Pointer
}

// struct FABRIC_OPERATION_ID
type FABRIC_OPERATION_ID struct {
	PartitionId syscall.GUID
	Reserved    unsafe.Pointer
}

// struct FABRIC_TEST_COMMAND_LIST_DESCRIPTION
type FABRIC_TEST_COMMAND_LIST_DESCRIPTION struct {
	TestCommandStateFilter int32
	TestCommandTypeFilter  int32
	Reserved               unsafe.Pointer
}

// struct TEST_COMMAND_QUERY_RESULT_ITEM
type TEST_COMMAND_QUERY_RESULT_ITEM struct {
	OperationId      syscall.GUID
	TestCommandState int32
	TestCommandType  int32
	Reserved         unsafe.Pointer
}

// struct TEST_COMMAND_QUERY_RESULT_LIST
type TEST_COMMAND_QUERY_RESULT_LIST struct {
	Count uint32
	Items unsafe.Pointer
}

// struct FABRIC_NODE_RESULT
type FABRIC_NODE_RESULT struct {
	NodeName     win32.PWSTR
	NodeInstance uint64
	Reserved     unsafe.Pointer
}

// struct FABRIC_RESTART_NODE_STATUS
type FABRIC_RESTART_NODE_STATUS struct {
	NodeResult *FABRIC_NODE_RESULT
	Reserved   unsafe.Pointer
}

// struct FABRIC_START_NODE_STATUS
type FABRIC_START_NODE_STATUS struct {
	NodeResult *FABRIC_NODE_RESULT
	Reserved   unsafe.Pointer
}

// struct FABRIC_STOP_NODE_STATUS
type FABRIC_STOP_NODE_STATUS struct {
	NodeResult *FABRIC_NODE_RESULT
	Reserved   unsafe.Pointer
}

// struct FABRIC_NODE_TRANSITION_RESULT
type FABRIC_NODE_TRANSITION_RESULT struct {
	ErrorCode  win32.HRESULT
	NodeResult *FABRIC_NODE_RESULT
	Reserved   unsafe.Pointer
}

// struct FABRIC_NODE_TRANSITION_PROGRESS
type FABRIC_NODE_TRANSITION_PROGRESS struct {
	State    int32
	Result   *FABRIC_NODE_TRANSITION_RESULT
	Reserved unsafe.Pointer
}

// struct FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION_LIST
type FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
}

// struct FABRIC_SERVICE_CORRELATION_DESCRIPTION
type FABRIC_SERVICE_CORRELATION_DESCRIPTION struct {
	ServiceName win32.PWSTR
	Scheme      int32
	Reserved    unsafe.Pointer
}

// struct FABRIC_SCALING_TRIGGER
type FABRIC_SCALING_TRIGGER struct {
	ScalingTriggerKind        int32
	ScalingTriggerDescription unsafe.Pointer
}

// struct FABRIC_SCALING_MECHANISM
type FABRIC_SCALING_MECHANISM struct {
	ScalingMechanismKind        int32
	ScalingMechanismDescription unsafe.Pointer
}

// struct FABRIC_SERVICE_SCALING_POLICY
type FABRIC_SERVICE_SCALING_POLICY struct {
	ServiceScalingPolicyTrigger   FABRIC_SCALING_TRIGGER
	ServiceScalingPolicyMechanism FABRIC_SCALING_MECHANISM
	Reserved                      unsafe.Pointer
}

// struct FABRIC_SCALING_TRIGGER_AVERAGE_PARTITION_LOAD
type FABRIC_SCALING_TRIGGER_AVERAGE_PARTITION_LOAD struct {
	MetricName             win32.PWSTR
	LowerLoadThreshold     float64
	UpperLoadThreshold     float64
	ScaleIntervalInSeconds uint32
	Reserved               unsafe.Pointer
}

// struct FABRIC_SCALING_TRIGGER_AVERAGE_SERVICE_LOAD
type FABRIC_SCALING_TRIGGER_AVERAGE_SERVICE_LOAD struct {
	MetricName             win32.PWSTR
	LowerLoadThreshold     float64
	UpperLoadThreshold     float64
	ScaleIntervalInSeconds uint32
	Reserved               unsafe.Pointer
}

// struct FABRIC_SCALING_TRIGGER_AVERAGE_SERVICE_LOAD_EX1
type FABRIC_SCALING_TRIGGER_AVERAGE_SERVICE_LOAD_EX1 struct {
	UseOnlyPrimaryLoad int8
	Reserved           unsafe.Pointer
}

// struct FABRIC_SCALING_MECHANISM_PARTITION_INSTANCE_COUNT
type FABRIC_SCALING_MECHANISM_PARTITION_INSTANCE_COUNT struct {
	MaximumInstanceCount int32
	MinimumInstanceCount int32
	ScaleIncrement       int32
	Reserved             unsafe.Pointer
}

// struct FABRIC_SCALING_MECHANISM_ADD_REMOVE_INCREMENTAL_NAMED_PARTITION
type FABRIC_SCALING_MECHANISM_ADD_REMOVE_INCREMENTAL_NAMED_PARTITION struct {
	MaximumPartitionCount int32
	MinimumPartitionCount int32
	ScaleIncrement        int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_PLACEMENT_POLICY_INVALID_DOMAIN_DESCRIPTION
type FABRIC_PLACEMENT_POLICY_INVALID_DOMAIN_DESCRIPTION struct {
	InvalidFaultDomain win32.PWSTR
	Reserved           unsafe.Pointer
}

// struct FABRIC_PLACEMENT_POLICY_REQUIRED_DOMAIN_DESCRIPTION
type FABRIC_PLACEMENT_POLICY_REQUIRED_DOMAIN_DESCRIPTION struct {
	RequiredFaultDomain win32.PWSTR
	Reserved            unsafe.Pointer
}

// struct FABRIC_PLACEMENT_POLICY_PREFERRED_PRIMARY_DOMAIN_DESCRIPTION
type FABRIC_PLACEMENT_POLICY_PREFERRED_PRIMARY_DOMAIN_DESCRIPTION struct {
	PreferredPrimaryFaultDomain win32.PWSTR
	Reserved                    unsafe.Pointer
}

// struct FABRIC_PLACEMENT_POLICY_REQUIRED_DOMAIN_DISTRIBUTION_DESCRIPTION
type FABRIC_PLACEMENT_POLICY_REQUIRED_DOMAIN_DISTRIBUTION_DESCRIPTION struct {
	Reserved unsafe.Pointer
}

// struct FABRIC_PLACEMENT_POLICY_NONPARTIALLY_PLACE_SERVICE_DESCRIPTION
type FABRIC_PLACEMENT_POLICY_NONPARTIALLY_PLACE_SERVICE_DESCRIPTION struct {
	Reserved unsafe.Pointer
}

// struct FABRIC_SERVICE_PLACEMENT_POLICY_DESCRIPTION
type FABRIC_SERVICE_PLACEMENT_POLICY_DESCRIPTION struct {
	Type  int32
	Value unsafe.Pointer
}

// struct FABRIC_SERVICE_PLACEMENT_POLICY_LIST
type FABRIC_SERVICE_PLACEMENT_POLICY_LIST struct {
	PolicyCount uint32
	Policies    *FABRIC_SERVICE_PLACEMENT_POLICY_DESCRIPTION
}

// struct FABRIC_SERVICE_PARTITION_INFORMATION
type FABRIC_SERVICE_PARTITION_INFORMATION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SINGLETON_PARTITION_INFORMATION
type FABRIC_SINGLETON_PARTITION_INFORMATION struct {
	Id       syscall.GUID
	Reserved unsafe.Pointer
}

// struct FABRIC_INT64_RANGE_PARTITION_INFORMATION
type FABRIC_INT64_RANGE_PARTITION_INFORMATION struct {
	Id       syscall.GUID
	LowKey   int64
	HighKey  int64
	Reserved unsafe.Pointer
}

// struct FABRIC_NAMED_PARTITION_INFORMATION
type FABRIC_NAMED_PARTITION_INFORMATION struct {
	Id       syscall.GUID
	Name     win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_CODE_PACKAGE_ENTRY_POINT_DESCRIPTION
type FABRIC_CODE_PACKAGE_ENTRY_POINT_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_CONTAINERHOST_ENTRY_POINT_DESCRIPTION
type FABRIC_CONTAINERHOST_ENTRY_POINT_DESCRIPTION struct {
	ImageName  win32.PWSTR
	Commands   win32.PWSTR
	EntryPoint win32.PWSTR
	Reserved   unsafe.Pointer
}

// struct FABRIC_EXEHOST_ENTRY_POINT_DESCRIPTION
type FABRIC_EXEHOST_ENTRY_POINT_DESCRIPTION struct {
	Program       win32.PWSTR
	Arguments     win32.PWSTR
	WorkingFolder int32
	Reserved      unsafe.Pointer
}

// struct FABRIC_EXEHOST_ENTRY_POINT_DESCRIPTION_EX1
type FABRIC_EXEHOST_ENTRY_POINT_DESCRIPTION_EX1 struct {
	PeriodicIntervalInSeconds            uint32
	ConsoleRedirectionEnabled            int8
	ConsoleRedirectionFileRetentionCount uint32
	ConsoleRedirectionFileMaxSizeInKb    uint32
	Reserved                             unsafe.Pointer
}

// struct FABRIC_EXEHOST_ENTRY_POINT_DESCRIPTION_EX2
type FABRIC_EXEHOST_ENTRY_POINT_DESCRIPTION_EX2 struct {
	IsExternalExecutable int8
	Reserved             unsafe.Pointer
}

// struct FABRIC_DLLHOST_HOSTED_DLL_DESCRIPTION
type FABRIC_DLLHOST_HOSTED_DLL_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_DLLHOST_HOSTED_UNMANAGED_DLL_DESCRIPTION
type FABRIC_DLLHOST_HOSTED_UNMANAGED_DLL_DESCRIPTION struct {
	DllName  win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_DLLHOST_HOSTED_MANAGED_DLL_DESCRIPTION
type FABRIC_DLLHOST_HOSTED_MANAGED_DLL_DESCRIPTION struct {
	AssemblyName win32.PWSTR
	Reserved     unsafe.Pointer
}

// struct FABRIC_DLLHOST_HOSTED_DLL_DESCRIPTION_LIST
type FABRIC_DLLHOST_HOSTED_DLL_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_DLLHOST_HOSTED_DLL_DESCRIPTION
}

// struct FABRIC_DLLHOST_ENTRY_POINT_DESCRIPTION
type FABRIC_DLLHOST_ENTRY_POINT_DESCRIPTION struct {
	IsolationPolicyType int32
	HostedDlls          *FABRIC_DLLHOST_HOSTED_DLL_DESCRIPTION_LIST
	Reserved            unsafe.Pointer
}

// struct FABRIC_CODE_PACKAGE_DESCRIPTION
type FABRIC_CODE_PACKAGE_DESCRIPTION struct {
	Name                   win32.PWSTR
	Version                win32.PWSTR
	ServiceManifestName    win32.PWSTR
	ServiceManifestVersion win32.PWSTR
	IsShared               int8
	SetupEntryPoint        *FABRIC_EXEHOST_ENTRY_POINT_DESCRIPTION
	EntryPoint             *FABRIC_CODE_PACKAGE_ENTRY_POINT_DESCRIPTION
	Reserved               unsafe.Pointer
}

// struct FABRIC_RUNAS_POLICY_DESCRIPTION
type FABRIC_RUNAS_POLICY_DESCRIPTION struct {
	UserName win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_CONFIGURATION_PACKAGE_DESCRIPTION
type FABRIC_CONFIGURATION_PACKAGE_DESCRIPTION struct {
	Name                   win32.PWSTR
	Version                win32.PWSTR
	ServiceManifestName    win32.PWSTR
	ServiceManifestVersion win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_DATA_PACKAGE_DESCRIPTION
type FABRIC_DATA_PACKAGE_DESCRIPTION struct {
	Name                   win32.PWSTR
	Version                win32.PWSTR
	ServiceManifestName    win32.PWSTR
	ServiceManifestVersion win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_DESCRIPTION
type FABRIC_SERVICE_TYPE_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_DESCRIPTION_LIST
type FABRIC_SERVICE_TYPE_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_TYPE_DESCRIPTION
}

// struct FABRIC_SERVICE_TYPE_DESCRIPTION_EXTENSION
type FABRIC_SERVICE_TYPE_DESCRIPTION_EXTENSION struct {
	Name     win32.PWSTR
	Value    win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_DESCRIPTION_EXTENSION_LIST
type FABRIC_SERVICE_TYPE_DESCRIPTION_EXTENSION_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_TYPE_DESCRIPTION_EXTENSION
}

// struct FABRIC_STATELESS_SERVICE_TYPE_DESCRIPTION
type FABRIC_STATELESS_SERVICE_TYPE_DESCRIPTION struct {
	ServiceTypeName      win32.PWSTR
	PlacementConstraints win32.PWSTR
	LoadMetrics          *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION_LIST
	Extensions           *FABRIC_SERVICE_TYPE_DESCRIPTION_EXTENSION_LIST
	UseImplicitHost      int8
	Reserved             unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_TYPE_DESCRIPTION_EX1
type FABRIC_STATELESS_SERVICE_TYPE_DESCRIPTION_EX1 struct {
	PolicyList *FABRIC_SERVICE_PLACEMENT_POLICY_LIST
	Reserved   unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_TYPE_DESCRIPTION
type FABRIC_STATEFUL_SERVICE_TYPE_DESCRIPTION struct {
	ServiceTypeName      win32.PWSTR
	PlacementConstraints win32.PWSTR
	LoadMetrics          *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION_LIST
	Extensions           *FABRIC_SERVICE_TYPE_DESCRIPTION_EXTENSION_LIST
	HasPersistedState    int8
	Reserved             unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_TYPE_DESCRIPTION_EX1
type FABRIC_STATEFUL_SERVICE_TYPE_DESCRIPTION_EX1 struct {
	PolicyList *FABRIC_SERVICE_PLACEMENT_POLICY_LIST
	Reserved   unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION
type FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION struct {
	ServiceTypeName win32.PWSTR
	LoadMetrics     *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION_LIST
	Reserved        unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION_LIST
type FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION
}

// struct FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION
type FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION struct {
	Description        *FABRIC_SERVICE_TYPE_DESCRIPTION
	Members            *FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION_LIST
	UseImplicitFactory int8
	Reserved           unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION_LIST
type FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION
}

// struct FABRIC_ENDPOINT_RESOURCE_DESCRIPTION
type FABRIC_ENDPOINT_RESOURCE_DESCRIPTION struct {
	Name            win32.PWSTR
	Protocol        win32.PWSTR
	Type            win32.PWSTR
	Port            uint32
	CertificateName win32.PWSTR
	Reserved        unsafe.Pointer
}

// struct FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_EX1
type FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_EX1 struct {
	UriScheme  win32.PWSTR
	PathSuffix win32.PWSTR
	Reserved   unsafe.Pointer
}

// struct FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_EX2
type FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_EX2 struct {
	CodePackageName win32.PWSTR
	IpAddressOrFqdn win32.PWSTR
	Reserved        unsafe.Pointer
}

// struct FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_LIST
type FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_ENDPOINT_RESOURCE_DESCRIPTION
}

// struct FABRIC_CODE_PACKAGE_DESCRIPTION_LIST
type FABRIC_CODE_PACKAGE_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_CODE_PACKAGE_DESCRIPTION
}

// struct FABRIC_CONFIGURATION_PACKAGE_DESCRIPTION_LIST
type FABRIC_CONFIGURATION_PACKAGE_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_CONFIGURATION_PACKAGE_DESCRIPTION
}

// struct FABRIC_DATA_PACKAGE_DESCRIPTION_LIST
type FABRIC_DATA_PACKAGE_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_DATA_PACKAGE_DESCRIPTION
}

// struct FABRIC_CONFIGURATION_PARAMETER
type FABRIC_CONFIGURATION_PARAMETER struct {
	Name         win32.PWSTR
	Value        win32.PWSTR
	MustOverride int8
	IsEncrypted  int8
	Reserved     unsafe.Pointer
}

// struct FABRIC_CONFIGURATION_PARAMETER_EX1
type FABRIC_CONFIGURATION_PARAMETER_EX1 struct {
	Type     win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_CONFIGURATION_PARAMETER_LIST
type FABRIC_CONFIGURATION_PARAMETER_LIST struct {
	Count uint32
	Items *FABRIC_CONFIGURATION_PARAMETER
}

// struct FABRIC_CONFIGURATION_SECTION
type FABRIC_CONFIGURATION_SECTION struct {
	Name       win32.PWSTR
	Parameters *FABRIC_CONFIGURATION_PARAMETER_LIST
	Reserved   unsafe.Pointer
}

// struct FABRIC_CONFIGURATION_SECTION_LIST
type FABRIC_CONFIGURATION_SECTION_LIST struct {
	Count uint32
	Items *FABRIC_CONFIGURATION_SECTION
}

// struct FABRIC_CONFIGURATION_SETTINGS
type FABRIC_CONFIGURATION_SETTINGS struct {
	Sections *FABRIC_CONFIGURATION_SECTION_LIST
	Reserved unsafe.Pointer
}

// struct FABRIC_STRING_LIST
type FABRIC_STRING_LIST struct {
	Count uint32
	Items *win32.PWSTR
}

// struct FABRIC_EPOCH
type FABRIC_EPOCH struct {
	DataLossNumber      int64
	ConfigurationNumber int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_OPERATION_METADATA
type FABRIC_OPERATION_METADATA struct {
	Type           int32
	SequenceNumber int64
	AtomicGroupId  int64
	Reserved       unsafe.Pointer
}

// struct FABRIC_REPLICA_INFORMATION
type FABRIC_REPLICA_INFORMATION struct {
	Id                int64
	Role              int32
	Status            int32
	ReplicatorAddress win32.PWSTR
	CurrentProgress   int64
	CatchUpCapability int64
	Reserved          unsafe.Pointer
}

// struct FABRIC_REPLICA_INFORMATION_EX1
type FABRIC_REPLICA_INFORMATION_EX1 struct {
	MustCatchup int8
	Reserved    unsafe.Pointer
}

// struct FABRIC_REPLICA_SET_CONFIGURATION
type FABRIC_REPLICA_SET_CONFIGURATION struct {
	ReplicaCount uint32
	Replicas     *FABRIC_REPLICA_INFORMATION
	WriteQuorum  uint32
	Reserved     unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC
type FABRIC_LOAD_METRIC struct {
	Name     win32.PWSTR
	Value    uint32
	Reserved unsafe.Pointer
}

// struct FABRIC_SECURITY_CREDENTIALS
type FABRIC_SECURITY_CREDENTIALS struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_REPLICATOR_SETTINGS
type FABRIC_REPLICATOR_SETTINGS struct {
	Flags                                    uint32
	RetryIntervalMilliseconds                uint32
	BatchAcknowledgementIntervalMilliseconds uint32
	ReplicatorAddress                        win32.PWSTR
	RequireServiceAck                        int8
	InitialReplicationQueueSize              uint32
	MaxReplicationQueueSize                  uint32
	InitialCopyQueueSize                     uint32
	MaxCopyQueueSize                         uint32
	SecurityCredentials                      *FABRIC_SECURITY_CREDENTIALS
	Reserved                                 unsafe.Pointer
}

// struct FABRIC_REPLICATOR_SETTINGS_EX1
type FABRIC_REPLICATOR_SETTINGS_EX1 struct {
	MaxReplicationQueueMemorySize        uint32
	SecondaryClearAcknowledgedOperations int8
	MaxReplicationMessageSize            uint32
	Reserved                             unsafe.Pointer
}

// struct FABRIC_REPLICATOR_SETTINGS_EX2
type FABRIC_REPLICATOR_SETTINGS_EX2 struct {
	UseStreamFaultsAndEndOfStreamOperationAck int8
	Reserved                                  unsafe.Pointer
}

// struct FABRIC_REPLICATOR_SETTINGS_EX3
type FABRIC_REPLICATOR_SETTINGS_EX3 struct {
	InitialPrimaryReplicationQueueSize              uint32
	MaxPrimaryReplicationQueueSize                  uint32
	MaxPrimaryReplicationQueueMemorySize            uint32
	InitialSecondaryReplicationQueueSize            uint32
	MaxSecondaryReplicationQueueSize                uint32
	MaxSecondaryReplicationQueueMemorySize          uint32
	PrimaryWaitForPendingQuorumsTimeoutMilliseconds uint32
	Reserved                                        unsafe.Pointer
}

// struct FABRIC_REPLICATOR_SETTINGS_EX4
type FABRIC_REPLICATOR_SETTINGS_EX4 struct {
	ReplicatorListenAddress  win32.PWSTR
	ReplicatorPublishAddress win32.PWSTR
	Reserved                 unsafe.Pointer
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

// struct FABRIC_NAMED_PROPERTY_METADATA_EX1
type FABRIC_NAMED_PROPERTY_METADATA_EX1 struct {
	CustomTypeId win32.PWSTR
	Reserved     unsafe.Pointer
}

// struct FABRIC_NAMED_PROPERTY
type FABRIC_NAMED_PROPERTY struct {
	Metadata *FABRIC_NAMED_PROPERTY_METADATA
	Value    *byte
	Reserved unsafe.Pointer
}

// struct FABRIC_PUT_PROPERTY_OPERATION
type FABRIC_PUT_PROPERTY_OPERATION struct {
	PropertyName   win32.PWSTR
	PropertyTypeId int32
	PropertyValue  unsafe.Pointer
	Reserved       unsafe.Pointer
}

// struct FABRIC_PUT_CUSTOM_PROPERTY_OPERATION
type FABRIC_PUT_CUSTOM_PROPERTY_OPERATION struct {
	PropertyName         win32.PWSTR
	PropertyTypeId       int32
	PropertyValue        unsafe.Pointer
	PropertyCustomTypeId win32.PWSTR
	Reserved             unsafe.Pointer
}

// struct FABRIC_GET_PROPERTY_OPERATION
type FABRIC_GET_PROPERTY_OPERATION struct {
	PropertyName win32.PWSTR
	IncludeValue int8
	Reserved     unsafe.Pointer
}

// struct FABRIC_DELETE_PROPERTY_OPERATION
type FABRIC_DELETE_PROPERTY_OPERATION struct {
	PropertyName win32.PWSTR
	Reserved     unsafe.Pointer
}

// struct FABRIC_CHECK_SEQUENCE_PROPERTY_OPERATION
type FABRIC_CHECK_SEQUENCE_PROPERTY_OPERATION struct {
	PropertyName   win32.PWSTR
	SequenceNumber int64
	Reserved       unsafe.Pointer
}

// struct FABRIC_CHECK_EXISTS_PROPERTY_OPERATION
type FABRIC_CHECK_EXISTS_PROPERTY_OPERATION struct {
	PropertyName   win32.PWSTR
	ExistenceCheck int8
	Reserved       unsafe.Pointer
}

// struct FABRIC_CHECK_VALUE_PROPERTY_OPERATION
type FABRIC_CHECK_VALUE_PROPERTY_OPERATION struct {
	PropertyName   win32.PWSTR
	PropertyTypeId int32
	PropertyValue  unsafe.Pointer
	Reserved       unsafe.Pointer
}

// struct FABRIC_PROPERTY_BATCH_OPERATION
type FABRIC_PROPERTY_BATCH_OPERATION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SERVICE_DESCRIPTION
type FABRIC_SERVICE_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_NODE_TRANSITION_DESCRIPTION
type FABRIC_NODE_TRANSITION_DESCRIPTION struct {
	NodeTransitionType int32
	Value              unsafe.Pointer
}

// struct FABRIC_NODE_STOP_DESCRIPTION
type FABRIC_NODE_STOP_DESCRIPTION struct {
	OperationId           syscall.GUID
	NodeName              win32.PWSTR
	NodeInstanceId        uint64
	StopDurationInSeconds uint32
	Reserved              unsafe.Pointer
}

// struct FABRIC_NODE_START_DESCRIPTION
type FABRIC_NODE_START_DESCRIPTION struct {
	OperationId    syscall.GUID
	NodeName       win32.PWSTR
	NodeInstanceId uint64
	Reserved       unsafe.Pointer
}

// struct FABRIC_SERVICE_UPDATE_DESCRIPTION
type FABRIC_SERVICE_UPDATE_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
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

// struct FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION
type FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION struct {
	Name     win32.PWSTR
	State    int32
	Reserved unsafe.Pointer
}

// struct FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION_LIST
type FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION
}

// struct FABRIC_UPGRADE_SAFETY_CHECK
type FABRIC_UPGRADE_SAFETY_CHECK struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_UPGRADE_PARTITION_SAFETY_CHECK
type FABRIC_UPGRADE_PARTITION_SAFETY_CHECK struct {
	PartitionId syscall.GUID
	Reserved    unsafe.Pointer
}

// struct FABRIC_UPGRADE_SEED_NODE_SAFETY_CHECK
type FABRIC_UPGRADE_SEED_NODE_SAFETY_CHECK struct {
	Reserved unsafe.Pointer
}

// struct FABRIC_UPGRADE_SAFETY_CHECK_LIST
type FABRIC_UPGRADE_SAFETY_CHECK_LIST struct {
	Count uint32
	Items *FABRIC_UPGRADE_SAFETY_CHECK
}

// struct FABRIC_NODE_UPGRADE_PROGRESS
type FABRIC_NODE_UPGRADE_PROGRESS struct {
	NodeName            win32.PWSTR
	UpgradePhase        int32
	PendingSafetyChecks *FABRIC_UPGRADE_SAFETY_CHECK_LIST
	Reserved            unsafe.Pointer
}

// struct FABRIC_NODE_UPGRADE_PROGRESS_LIST
type FABRIC_NODE_UPGRADE_PROGRESS_LIST struct {
	Count uint32
	Items *FABRIC_NODE_UPGRADE_PROGRESS
}

// struct FABRIC_UPGRADE_DOMAIN_PROGRESS
type FABRIC_UPGRADE_DOMAIN_PROGRESS struct {
	UpgradeDomainName win32.PWSTR
	NodeProgressList  *FABRIC_NODE_UPGRADE_PROGRESS_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_DESCRIPTION
type FABRIC_STATELESS_SERVICE_DESCRIPTION struct {
	ApplicationName            win32.PWSTR
	ServiceName                win32.PWSTR
	ServiceTypeName            win32.PWSTR
	InitializationDataSize     uint32
	InitializationData         *byte
	PartitionScheme            int32
	PartitionSchemeDescription unsafe.Pointer
	InstanceCount              int32
	PlacementConstraints       win32.PWSTR
	CorrelationCount           uint32
	Correlations               *FABRIC_SERVICE_CORRELATION_DESCRIPTION
	MetricCount                uint32
	Metrics                    *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
	Reserved                   unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_DESCRIPTION_EX1
type FABRIC_STATELESS_SERVICE_DESCRIPTION_EX1 struct {
	PolicyList *FABRIC_SERVICE_PLACEMENT_POLICY_LIST
	Reserved   unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_DESCRIPTION_EX2
type FABRIC_STATELESS_SERVICE_DESCRIPTION_EX2 struct {
	IsDefaultMoveCostSpecified int8
	DefaultMoveCost            int32
	Reserved                   unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_DESCRIPTION_EX3
type FABRIC_STATELESS_SERVICE_DESCRIPTION_EX3 struct {
	ServicePackageActivationMode int32
	ServiceDnsName               win32.PWSTR
	Reserved                     unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_DESCRIPTION_EX4
type FABRIC_STATELESS_SERVICE_DESCRIPTION_EX4 struct {
	ScalingPolicyCount     uint32
	ServiceScalingPolicies *FABRIC_SERVICE_SCALING_POLICY
	Reserved               unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_UPDATE_DESCRIPTION
type FABRIC_STATELESS_SERVICE_UPDATE_DESCRIPTION struct {
	Flags         uint32
	InstanceCount int32
	Reserved      unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_UPDATE_DESCRIPTION_EX1
type FABRIC_STATELESS_SERVICE_UPDATE_DESCRIPTION_EX1 struct {
	PlacementConstraints win32.PWSTR
	PolicyList           *FABRIC_SERVICE_PLACEMENT_POLICY_LIST
	CorrelationCount     uint32
	Correlations         *FABRIC_SERVICE_CORRELATION_DESCRIPTION
	MetricCount          uint32
	Metrics              *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
	Reserved             unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_UPDATE_DESCRIPTION_EX2
type FABRIC_STATELESS_SERVICE_UPDATE_DESCRIPTION_EX2 struct {
	DefaultMoveCost int32
	Reserved        unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_UPDATE_DESCRIPTION_EX3
type FABRIC_STATELESS_SERVICE_UPDATE_DESCRIPTION_EX3 struct {
	RepartitionKind        int32
	RepartitionDescription unsafe.Pointer
	ScalingPolicyCount     uint32
	ServiceScalingPolicies *FABRIC_SERVICE_SCALING_POLICY
	Reserved               unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_DESCRIPTION
type FABRIC_STATEFUL_SERVICE_DESCRIPTION struct {
	ApplicationName            win32.PWSTR
	ServiceName                win32.PWSTR
	ServiceTypeName            win32.PWSTR
	InitializationDataSize     uint32
	InitializationData         *byte
	PartitionScheme            int32
	PartitionSchemeDescription unsafe.Pointer
	TargetReplicaSetSize       int32
	MinReplicaSetSize          int32
	PlacementConstraints       win32.PWSTR
	CorrelationCount           uint32
	Correlations               *FABRIC_SERVICE_CORRELATION_DESCRIPTION
	MetricCount                uint32
	Metrics                    *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
	HasPersistedState          int8
	Reserved                   unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_FAILOVER_SETTINGS
type FABRIC_STATEFUL_SERVICE_FAILOVER_SETTINGS struct {
	Flags                             uint32
	ReplicaRestartWaitDurationSeconds uint32
	QuorumLossWaitDurationSeconds     uint32
	Reserved                          unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_FAILOVER_SETTINGS_EX1
type FABRIC_STATEFUL_SERVICE_FAILOVER_SETTINGS_EX1 struct {
	StandByReplicaKeepDurationSeconds uint32
	Reserved                          unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_DESCRIPTION_EX1
type FABRIC_STATEFUL_SERVICE_DESCRIPTION_EX1 struct {
	PolicyList       *FABRIC_SERVICE_PLACEMENT_POLICY_LIST
	FailoverSettings *FABRIC_STATEFUL_SERVICE_FAILOVER_SETTINGS
	Reserved         unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_DESCRIPTION_EX2
type FABRIC_STATEFUL_SERVICE_DESCRIPTION_EX2 struct {
	IsDefaultMoveCostSpecified int8
	DefaultMoveCost            int32
	Reserved                   unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_DESCRIPTION_EX3
type FABRIC_STATEFUL_SERVICE_DESCRIPTION_EX3 struct {
	ServicePackageActivationMode int32
	ServiceDnsName               win32.PWSTR
	Reserved                     unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_DESCRIPTION_EX4
type FABRIC_STATEFUL_SERVICE_DESCRIPTION_EX4 struct {
	ScalingPolicyCount     uint32
	ServiceScalingPolicies *FABRIC_SERVICE_SCALING_POLICY
	Reserved               unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION
type FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION struct {
	Flags                             uint32
	TargetReplicaSetSize              int32
	ReplicaRestartWaitDurationSeconds uint32
	QuorumLossWaitDurationSeconds     uint32
	Reserved                          unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX1
type FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX1 struct {
	StandByReplicaKeepDurationSeconds uint32
	Reserved                          unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX2
type FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX2 struct {
	MinReplicaSetSize int32
	Reserved          unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX3
type FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX3 struct {
	PlacementConstraints win32.PWSTR
	PolicyList           *FABRIC_SERVICE_PLACEMENT_POLICY_LIST
	CorrelationCount     uint32
	Correlations         *FABRIC_SERVICE_CORRELATION_DESCRIPTION
	MetricCount          uint32
	Metrics              *FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
	Reserved             unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX4
type FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX4 struct {
	DefaultMoveCost int32
	Reserved        unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX5
type FABRIC_STATEFUL_SERVICE_UPDATE_DESCRIPTION_EX5 struct {
	RepartitionKind        int32
	RepartitionDescription unsafe.Pointer
	ScalingPolicyCount     uint32
	ServiceScalingPolicies *FABRIC_SERVICE_SCALING_POLICY
	Reserved               unsafe.Pointer
}

// struct FABRIC_NAMED_REPARTITION_DESCRIPTION
type FABRIC_NAMED_REPARTITION_DESCRIPTION struct {
	NamesToAddCount    uint32
	NamesToAdd         *win32.PWSTR
	NamesToRemoveCount uint32
	NamesToRemove      *win32.PWSTR
	Reserved           unsafe.Pointer
}

// struct FABRIC_UNIFORM_INT64_RANGE_PARTITION_SCHEME_DESCRIPTION
type FABRIC_UNIFORM_INT64_RANGE_PARTITION_SCHEME_DESCRIPTION struct {
	PartitionCount int32
	LowKey         int64
	HighKey        int64
	Reserved       unsafe.Pointer
}

// struct FABRIC_NAMED_PARTITION_SCHEME_DESCRIPTION
type FABRIC_NAMED_PARTITION_SCHEME_DESCRIPTION struct {
	PartitionCount int32
	Names          *win32.PWSTR
	Reserved       unsafe.Pointer
}

// struct FABRIC_RESOLVED_SERVICE_ENDPOINT
type FABRIC_RESOLVED_SERVICE_ENDPOINT struct {
	Address  win32.PWSTR
	Role     int32
	Reserved unsafe.Pointer
}

// struct FABRIC_RESOLVED_SERVICE_PARTITION
type FABRIC_RESOLVED_SERVICE_PARTITION struct {
	Info          FABRIC_SERVICE_PARTITION_INFORMATION
	EndpointCount uint32
	Endpoints     *FABRIC_RESOLVED_SERVICE_ENDPOINT
	ServiceName   win32.PWSTR
	Reserved      unsafe.Pointer
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

// struct FABRIC_SERVICE_NOTIFICATION_FILTER_DESCRIPTION
type FABRIC_SERVICE_NOTIFICATION_FILTER_DESCRIPTION struct {
	Name     win32.PWSTR
	Flags    int32
	Reserved unsafe.Pointer
}

// struct FABRIC_APPLICATION_PARAMETER
type FABRIC_APPLICATION_PARAMETER struct {
	Name     win32.PWSTR
	Value    win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_APPLICATION_PARAMETER_LIST
type FABRIC_APPLICATION_PARAMETER_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_PARAMETER
}

// struct FABRIC_STRING_MAP
type FABRIC_STRING_MAP struct {
	Count uint32
	Items *FABRIC_APPLICATION_PARAMETER
}

// struct FABRIC_APPLICATION_DESCRIPTION
type FABRIC_APPLICATION_DESCRIPTION struct {
	ApplicationName        win32.PWSTR
	ApplicationTypeName    win32.PWSTR
	ApplicationTypeVersion win32.PWSTR
	ApplicationParameters  *FABRIC_APPLICATION_PARAMETER_LIST
	Reserved               unsafe.Pointer
}

// struct FABRIC_APPLICATION_METRIC_DESCRIPTION
type FABRIC_APPLICATION_METRIC_DESCRIPTION struct {
	Name                     win32.PWSTR
	NodeReservationCapacity  uint32
	MaximumNodeCapacity      uint32
	TotalApplicationCapacity uint32
	Reserved                 unsafe.Pointer
}

// struct FABRIC_APPLICATION_METRIC_LIST
type FABRIC_APPLICATION_METRIC_LIST struct {
	Count      uint32
	Capacities *FABRIC_APPLICATION_METRIC_DESCRIPTION
}

// struct FABRIC_APPLICATION_CAPACITY_DESCRIPTION
type FABRIC_APPLICATION_CAPACITY_DESCRIPTION struct {
	MaximumNodes uint32
	MinimumNodes uint32
	Metrics      *FABRIC_APPLICATION_METRIC_LIST
	Reserved     unsafe.Pointer
}

// struct FABRIC_APPLICATION_DESCRIPTION_EX1
type FABRIC_APPLICATION_DESCRIPTION_EX1 struct {
	ApplicationCapacity *FABRIC_APPLICATION_CAPACITY_DESCRIPTION
	Reserved            unsafe.Pointer
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

// struct FABRIC_ROLLING_UPGRADE_POLICY_DESCRIPTION
type FABRIC_ROLLING_UPGRADE_POLICY_DESCRIPTION struct {
	RollingUpgradeMode                     int32
	ForceRestart                           int8
	UpgradeReplicaSetCheckTimeoutInSeconds uint32
	Reserved                               unsafe.Pointer
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

// struct FABRIC_ROLLING_UPGRADE_MONITORING_POLICY_EX1
type FABRIC_ROLLING_UPGRADE_MONITORING_POLICY_EX1 struct {
	HealthCheckStableDurationInSeconds uint32
	Reserved                           unsafe.Pointer
}

// struct FABRIC_ROLLING_UPGRADE_POLICY_DESCRIPTION_EX1
type FABRIC_ROLLING_UPGRADE_POLICY_DESCRIPTION_EX1 struct {
	MonitoringPolicy *FABRIC_ROLLING_UPGRADE_MONITORING_POLICY
	HealthPolicy     unsafe.Pointer
	Reserved         unsafe.Pointer
}

// struct FABRIC_ROLLING_UPGRADE_POLICY_DESCRIPTION_EX2
type FABRIC_ROLLING_UPGRADE_POLICY_DESCRIPTION_EX2 struct {
	EnableDeltaHealthEvaluation int8
	UpgradeHealthPolicy         unsafe.Pointer
	Reserved                    unsafe.Pointer
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

// struct FABRIC_APPLICATION_UPGRADE_UPDATE_DESCRIPTION
type FABRIC_APPLICATION_UPGRADE_UPDATE_DESCRIPTION struct {
	ApplicationName          win32.PWSTR
	UpgradeKind              int32
	UpdateFlags              uint32
	UpgradePolicyDescription unsafe.Pointer
	Reserved                 unsafe.Pointer
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

// struct FABRIC_SERVICE_GROUP_DESCRIPTION
type FABRIC_SERVICE_GROUP_DESCRIPTION struct {
	Description        *FABRIC_SERVICE_DESCRIPTION
	MemberCount        uint32
	MemberDescriptions *FABRIC_SERVICE_GROUP_MEMBER_DESCRIPTION
	Reserved           unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_UPDATE_DESCRIPTION
type FABRIC_SERVICE_GROUP_UPDATE_DESCRIPTION struct {
	Description *FABRIC_SERVICE_UPDATE_DESCRIPTION
	Reserved    unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_ITEM_METADATA
type FABRIC_KEY_VALUE_STORE_ITEM_METADATA struct {
	Key              win32.PWSTR
	ValueSizeInBytes int32
	SequenceNumber   int64
	LastModifiedUtc  FILETIME_
	Reserved         unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_ITEM_METADATA_EX1
type FABRIC_KEY_VALUE_STORE_ITEM_METADATA_EX1 struct {
	LastModifiedOnPrimaryUtc FILETIME_
	Reserved                 unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_ITEM
type FABRIC_KEY_VALUE_STORE_ITEM struct {
	Metadata *FABRIC_KEY_VALUE_STORE_ITEM_METADATA
	Value    *byte
	Reserved unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_TRANSACTION_SETTINGS
type FABRIC_KEY_VALUE_STORE_TRANSACTION_SETTINGS struct {
	SerializationBlockSize uint32
	Reserved               unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_RESTORE_SETTINGS
type FABRIC_KEY_VALUE_STORE_RESTORE_SETTINGS struct {
	InlineReopen int8
	Reserved     unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_RESTORE_SETTINGS_EX1
type FABRIC_KEY_VALUE_STORE_RESTORE_SETTINGS_EX1 struct {
	EnableLsnCheck int8
	Reserved       unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS
type FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS struct {
	TransactionDrainTimeoutInSeconds uint32
	SecondaryNotificationMode        int32
	Reserved                         unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS_EX1
type FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS_EX1 struct {
	EnableCopyNotificationPrefetch int8
	Reserved                       unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS_EX2
type FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS_EX2 struct {
	FullCopyMode int32
	Reserved     unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS_EX3
type FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS_EX3 struct {
	LogTruncationIntervalInMinutes int32
	Reserved                       unsafe.Pointer
}

// struct FABRIC_ESE_LOCAL_STORE_SETTINGS
type FABRIC_ESE_LOCAL_STORE_SETTINGS struct {
	DbFolderPath                      win32.PWSTR
	LogFileSizeInKB                   int32
	LogBufferSizeInKB                 int32
	MaxCursors                        int32
	MaxVerPages                       int32
	MaxAsyncCommitDelayInMilliseconds int32
	Reserved                          unsafe.Pointer
}

// struct FABRIC_ESE_LOCAL_STORE_SETTINGS_EX1
type FABRIC_ESE_LOCAL_STORE_SETTINGS_EX1 struct {
	EnableIncrementalBackup int8
	Reserved                unsafe.Pointer
}

// struct FABRIC_ESE_LOCAL_STORE_SETTINGS_EX2
type FABRIC_ESE_LOCAL_STORE_SETTINGS_EX2 struct {
	MaxCacheSizeInMB int32
	Reserved         unsafe.Pointer
}

// struct FABRIC_ESE_LOCAL_STORE_SETTINGS_EX3
type FABRIC_ESE_LOCAL_STORE_SETTINGS_EX3 struct {
	MaxDefragFrequencyInMinutes int32
	DefragThresholdInMB         int32
	DatabasePageSizeInKB        int32
	Reserved                    unsafe.Pointer
}

// struct FABRIC_ESE_LOCAL_STORE_SETTINGS_EX4
type FABRIC_ESE_LOCAL_STORE_SETTINGS_EX4 struct {
	CompactionThresholdInMB int32
	Reserved                unsafe.Pointer
}

// struct FABRIC_ESE_LOCAL_STORE_SETTINGS_EX5
type FABRIC_ESE_LOCAL_STORE_SETTINGS_EX5 struct {
	IntrinsicValueThresholdInBytes int32
	Reserved                       unsafe.Pointer
}

// struct FABRIC_ESE_LOCAL_STORE_SETTINGS_EX6
type FABRIC_ESE_LOCAL_STORE_SETTINGS_EX6 struct {
	EnableOverwriteOnUpdate int8
	Reserved                unsafe.Pointer
}

// struct FABRIC_NODE_CONTEXT
type FABRIC_NODE_CONTEXT struct {
	NodeName        win32.PWSTR
	NodeType        win32.PWSTR
	IpAddressOrFqdn win32.PWSTR
	NodeInstanceId  uint64
	NodeId          FABRIC_NODE_ID
	Reserved        unsafe.Pointer
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

// struct FABRIC_CLIENT_SETTINGS_EX1
type FABRIC_CLIENT_SETTINGS_EX1 struct {
	ClientFriendlyName                     win32.PWSTR
	PartitionLocationCacheBucketCount      uint32
	HealthReportRetrySendIntervalInSeconds uint32
	Reserved                               unsafe.Pointer
}

// struct FABRIC_CLIENT_SETTINGS_EX2
type FABRIC_CLIENT_SETTINGS_EX2 struct {
	NotificationGatewayConnectionTimeoutInSeconds uint32
	NotificationCacheUpdateTimeoutInSeconds       uint32
	Reserved                                      unsafe.Pointer
}

// struct FABRIC_CLIENT_SETTINGS_EX3
type FABRIC_CLIENT_SETTINGS_EX3 struct {
	AuthTokenBufferSize uint32
	Reserved            unsafe.Pointer
}

// struct FABRIC_CLIENT_SETTINGS_EX4
type FABRIC_CLIENT_SETTINGS_EX4 struct {
	ConnectionIdleTimeoutInSeconds uint32
	Reserved                       unsafe.Pointer
}

// struct FABRIC_QUERY_PAGING_DESCRIPTION
type FABRIC_QUERY_PAGING_DESCRIPTION struct {
	ContinuationToken win32.PWSTR
	MaxResults        int32
	Reserved          unsafe.Pointer
}

// struct FABRIC_CLUSTER_MANIFEST_QUERY_DESCRIPTION
type FABRIC_CLUSTER_MANIFEST_QUERY_DESCRIPTION struct {
	ClusterManifestVersion win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_DESCRIPTION
type FABRIC_NODE_QUERY_DESCRIPTION struct {
	NodeNameFilter win32.PWSTR
	Reserved       unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_RESULT_ITEM
type FABRIC_NODE_QUERY_RESULT_ITEM struct {
	NodeName              win32.PWSTR
	IpAddressOrFqdn       win32.PWSTR
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

// struct FABRIC_NODE_QUERY_RESULT_ITEM_EX1
type FABRIC_NODE_QUERY_RESULT_ITEM_EX1 struct {
	NodeId   FABRIC_NODE_ID
	Reserved unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_RESULT_ITEM_EX2
type FABRIC_NODE_QUERY_RESULT_ITEM_EX2 struct {
	NodeInstanceId uint64
	Reserved       unsafe.Pointer
}

// struct FABRIC_NODE_DEACTIVATION_TASK_ID
type FABRIC_NODE_DEACTIVATION_TASK_ID struct {
	Id       win32.PWSTR
	Type     int32
	Reserved unsafe.Pointer
}

// struct FABRIC_NODE_DEACTIVATION_TASK
type FABRIC_NODE_DEACTIVATION_TASK struct {
	TaskId   *FABRIC_NODE_DEACTIVATION_TASK_ID
	Intent   int32
	Reserved unsafe.Pointer
}

// struct FABRIC_NODE_DEACTIVATION_TASK_LIST
type FABRIC_NODE_DEACTIVATION_TASK_LIST struct {
	Count uint32
	Items *FABRIC_NODE_DEACTIVATION_TASK
}

// struct FABRIC_NODE_DEACTIVATION_QUERY_RESULT_ITEM
type FABRIC_NODE_DEACTIVATION_QUERY_RESULT_ITEM struct {
	EffectiveIntent int32
	Status          int32
	Tasks           *FABRIC_NODE_DEACTIVATION_TASK_LIST
	Reserved        unsafe.Pointer
}

// struct FABRIC_SAFETY_CHECK
type FABRIC_SAFETY_CHECK struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_PARTITION_SAFETY_CHECK
type FABRIC_PARTITION_SAFETY_CHECK struct {
	PartitionId syscall.GUID
	Reserved    unsafe.Pointer
}

// struct FABRIC_SEED_NODE_SAFETY_CHECK
type FABRIC_SEED_NODE_SAFETY_CHECK struct {
	Reserved unsafe.Pointer
}

// struct FABRIC_SAFETY_CHECK_LIST
type FABRIC_SAFETY_CHECK_LIST struct {
	Count uint32
	Items *FABRIC_SAFETY_CHECK
}

// struct FABRIC_NODE_DEACTIVATION_QUERY_RESULT_ITEM_EX1
type FABRIC_NODE_DEACTIVATION_QUERY_RESULT_ITEM_EX1 struct {
	PendingSafetyChecks *FABRIC_SAFETY_CHECK_LIST
	Reserved            unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_RESULT_ITEM_EX3
type FABRIC_NODE_QUERY_RESULT_ITEM_EX3 struct {
	NodeDeactivationInfo *FABRIC_NODE_DEACTIVATION_QUERY_RESULT_ITEM
	Reserved             unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_RESULT_ITEM_EX4
type FABRIC_NODE_QUERY_RESULT_ITEM_EX4 struct {
	IsStopped int8
	Reserved  unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_RESULT_ITEM_EX5
type FABRIC_NODE_QUERY_RESULT_ITEM_EX5 struct {
	NodeDownTimeInSeconds int64
	Reserved              unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_RESULT_ITEM_EX6
type FABRIC_NODE_QUERY_RESULT_ITEM_EX6 struct {
	NodeUpAt   FILETIME_
	NodeDownAt FILETIME_
	Reserved   unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_RESULT_LIST
type FABRIC_NODE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_NODE_QUERY_RESULT_ITEM
}

// struct FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION
type FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION struct {
	ApplicationTypeNameFilter win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct PAGED_FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION
type PAGED_FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION struct {
	ApplicationTypeNameFilter    win32.PWSTR
	MaxResults                   int32
	ContinuationToken            win32.PWSTR
	ExcludeApplicationParameters int8
	Reserved                     unsafe.Pointer
}

// struct PAGED_FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION_EX1
type PAGED_FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION_EX1 struct {
	ApplicationTypeVersionFilter win32.PWSTR
	Reserved                     unsafe.Pointer
}

// struct PAGED_FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION_EX2
type PAGED_FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION_EX2 struct {
	ApplicationTypeDefinitionKindFilter uint32
	Reserved                            unsafe.Pointer
}

// struct FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM
type FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM struct {
	ApplicationTypeName    win32.PWSTR
	ApplicationTypeVersion win32.PWSTR
	DefaultParameters      *FABRIC_APPLICATION_PARAMETER_LIST
	Reserved               unsafe.Pointer
}

// struct FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM_EX1
type FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM_EX1 struct {
	Status        int32
	StatusDetails win32.PWSTR
	Reserved      unsafe.Pointer
}

// struct FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM_EX2
type FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM_EX2 struct {
	ApplicationTypeDefinitionKind int32
	Reserved                      unsafe.Pointer
}

// struct FABRIC_APPLICATION_TYPE_QUERY_RESULT_LIST
type FABRIC_APPLICATION_TYPE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_TYPE_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_TYPE_QUERY_DESCRIPTION
type FABRIC_SERVICE_TYPE_QUERY_DESCRIPTION struct {
	ApplicationTypeName    win32.PWSTR
	ApplicationTypeVersion win32.PWSTR
	ServiceTypeNameFilter  win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_DESCRIPTION
type FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_DESCRIPTION struct {
	ApplicationTypeName        win32.PWSTR
	ApplicationTypeVersion     win32.PWSTR
	ServiceGroupTypeNameFilter win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM
type FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM struct {
	ServiceTypeDescription *FABRIC_SERVICE_TYPE_DESCRIPTION
	ServiceManifestVersion win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM_EX1
type FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM_EX1 struct {
	ServiceManifestName win32.PWSTR
	Reserved            unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM_EX2
type FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM_EX2 struct {
	IsServiceGroup int8
	Reserved       unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_QUERY_RESULT_LIST
type FABRIC_SERVICE_TYPE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_TYPE_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_ITEM
type FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_ITEM struct {
	ServiceGroupMemberTypeDescription *FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION_LIST
	ServiceManifestVersion            win32.PWSTR
	ServiceManifestName               win32.PWSTR
	Reserved                          unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_LIST
type FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_ITEM
}

// struct FABRIC_APPLICATION_QUERY_DESCRIPTION
type FABRIC_APPLICATION_QUERY_DESCRIPTION struct {
	ApplicationNameFilter win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_APPLICATION_QUERY_DESCRIPTION_EX1
type FABRIC_APPLICATION_QUERY_DESCRIPTION_EX1 struct {
	ContinuationToken win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_APPLICATION_QUERY_DESCRIPTION_EX2
type FABRIC_APPLICATION_QUERY_DESCRIPTION_EX2 struct {
	ApplicationTypeNameFilter    win32.PWSTR
	ExcludeApplicationParameters int8
	Reserved                     unsafe.Pointer
}

// struct FABRIC_APPLICATION_QUERY_DESCRIPTION_EX3
type FABRIC_APPLICATION_QUERY_DESCRIPTION_EX3 struct {
	ApplicationDefinitionKindFilter uint32
	Reserved                        unsafe.Pointer
}

// struct FABRIC_APPLICATION_QUERY_DESCRIPTION_EX4
type FABRIC_APPLICATION_QUERY_DESCRIPTION_EX4 struct {
	MaxResults int32
	Reserved   unsafe.Pointer
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

// struct FABRIC_APPLICATION_QUERY_RESULT_ITEM_EX1
type FABRIC_APPLICATION_QUERY_RESULT_ITEM_EX1 struct {
	UpgradeTypeVersion win32.PWSTR
	UpgradeParameters  *FABRIC_APPLICATION_PARAMETER_LIST
	Reserved           unsafe.Pointer
}

// struct FABRIC_APPLICATION_QUERY_RESULT_ITEM_EX2
type FABRIC_APPLICATION_QUERY_RESULT_ITEM_EX2 struct {
	ApplicationDefinitionKind int32
	Reserved                  unsafe.Pointer
}

// struct FABRIC_APPLICATION_QUERY_RESULT_LIST
type FABRIC_APPLICATION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_QUERY_DESCRIPTION
type FABRIC_SERVICE_QUERY_DESCRIPTION struct {
	ApplicationName   win32.PWSTR
	ServiceNameFilter win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_QUERY_DESCRIPTION_EX1
type FABRIC_SERVICE_QUERY_DESCRIPTION_EX1 struct {
	ContinuationToken win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_QUERY_DESCRIPTION_EX2
type FABRIC_SERVICE_QUERY_DESCRIPTION_EX2 struct {
	ServiceTypeNameFilter win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_SERVICE_QUERY_DESCRIPTION_EX3
type FABRIC_SERVICE_QUERY_DESCRIPTION_EX3 struct {
	MaxResults int32
	Reserved   unsafe.Pointer
}

// struct FABRIC_SYSTEM_SERVICE_QUERY_DESCRIPTION
type FABRIC_SYSTEM_SERVICE_QUERY_DESCRIPTION struct {
	SystemServiceNameFilter win32.PWSTR
	Reserved                unsafe.Pointer
}

// struct FABRIC_SERVICE_QUERY_RESULT_ITEM
type FABRIC_SERVICE_QUERY_RESULT_ITEM struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SERVICE_QUERY_RESULT_LIST
type FABRIC_SERVICE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_QUERY_RESULT_ITEM
}

// struct FABRIC_STATELESS_SERVICE_QUERY_RESULT_ITEM
type FABRIC_STATELESS_SERVICE_QUERY_RESULT_ITEM struct {
	ServiceName            win32.PWSTR
	ServiceTypeName        win32.PWSTR
	ServiceManifestVersion win32.PWSTR
	HealthState            int32
	Reserved               unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_QUERY_RESULT_ITEM_EX1
type FABRIC_STATELESS_SERVICE_QUERY_RESULT_ITEM_EX1 struct {
	ServiceStatus int32
	Reserved      unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_QUERY_RESULT_ITEM_EX2
type FABRIC_STATELESS_SERVICE_QUERY_RESULT_ITEM_EX2 struct {
	IsServiceGroup int8
	Reserved       unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_QUERY_RESULT_ITEM
type FABRIC_STATEFUL_SERVICE_QUERY_RESULT_ITEM struct {
	ServiceName            win32.PWSTR
	ServiceTypeName        win32.PWSTR
	ServiceManifestVersion win32.PWSTR
	HasPersistedState      int8
	HealthState            int32
	Reserved               unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_QUERY_RESULT_ITEM_EX1
type FABRIC_STATEFUL_SERVICE_QUERY_RESULT_ITEM_EX1 struct {
	ServiceStatus int32
	Reserved      unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_QUERY_RESULT_ITEM_EX2
type FABRIC_STATEFUL_SERVICE_QUERY_RESULT_ITEM_EX2 struct {
	IsServiceGroup int8
	Reserved       unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_QUERY_DESCRIPTION
type FABRIC_SERVICE_GROUP_MEMBER_QUERY_DESCRIPTION struct {
	ApplicationName   win32.PWSTR
	ServiceNameFilter win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_ITEM
type FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_ITEM struct {
	ServiceType win32.PWSTR
	ServiceName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_LIST
type FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_ITEM
}

// struct FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_ITEM
type FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_ITEM struct {
	ServiceName win32.PWSTR
	Members     *FABRIC_SERVICE_GROUP_MEMBER_MEMBER_QUERY_RESULT_LIST
	Reserved    unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_LIST
type FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_ITEM
}

// struct FABRIC_PARTITION_LOAD_INFORMATION_QUERY_DESCRIPTION
type FABRIC_PARTITION_LOAD_INFORMATION_QUERY_DESCRIPTION struct {
	PartitionId syscall.GUID
	Reserved    unsafe.Pointer
}

// struct FABRIC_REPLICA_LOAD_INFORMATION_QUERY_DESCRIPTION
type FABRIC_REPLICA_LOAD_INFORMATION_QUERY_DESCRIPTION struct {
	PartitionId         syscall.GUID
	ReplicaOrInstanceId int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_UNPLACED_REPLICA_INFORMATION_QUERY_DESCRIPTION
type FABRIC_UNPLACED_REPLICA_INFORMATION_QUERY_DESCRIPTION struct {
	ServiceName        win32.PWSTR
	PartitionId        syscall.GUID
	OnlyQueryPrimaries int8
	Reserved           unsafe.Pointer
}

// struct FABRIC_SERVICE_PARTITION_QUERY_DESCRIPTION
type FABRIC_SERVICE_PARTITION_QUERY_DESCRIPTION struct {
	ServiceName       win32.PWSTR
	PartitionIdFilter syscall.GUID
	Reserved          unsafe.Pointer
}

// struct FABRIC_NODE_LOAD_INFORMATION_QUERY_DESCRIPTION
type FABRIC_NODE_LOAD_INFORMATION_QUERY_DESCRIPTION struct {
	NodeName win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_APPLICATION_LOAD_INFORMATION_QUERY_DESCRIPTION
type FABRIC_APPLICATION_LOAD_INFORMATION_QUERY_DESCRIPTION struct {
	ApplicationName win32.PWSTR
	Reserved        unsafe.Pointer
}

// struct FABRIC_SERVICE_PARTITION_QUERY_RESULT_ITEM
type FABRIC_SERVICE_PARTITION_QUERY_RESULT_ITEM struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SERVICE_PARTITION_QUERY_RESULT_LIST
type FABRIC_SERVICE_PARTITION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_PARTITION_QUERY_RESULT_ITEM
}

// struct FABRIC_STATELESS_SERVICE_PARTITION_QUERY_RESULT_ITEM
type FABRIC_STATELESS_SERVICE_PARTITION_QUERY_RESULT_ITEM struct {
	PartitionInformation *FABRIC_SERVICE_PARTITION_INFORMATION
	InstanceCount        uint32
	HealthState          int32
	PartitionStatus      int32
	Reserved             unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_PARTITION_QUERY_RESULT_ITEM
type FABRIC_STATEFUL_SERVICE_PARTITION_QUERY_RESULT_ITEM struct {
	PartitionInformation            *FABRIC_SERVICE_PARTITION_INFORMATION
	TargetReplicaSetSize            uint32
	MinReplicaSetSize               uint32
	HealthState                     int32
	PartitionStatus                 int32
	LastQuorumLossDurationInSeconds int64
	Reserved                        unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_PARTITION_QUERY_RESULT_ITEM_EX1
type FABRIC_STATEFUL_SERVICE_PARTITION_QUERY_RESULT_ITEM_EX1 struct {
	PrimaryEpoch FABRIC_EPOCH
	Reserved     unsafe.Pointer
}

// struct FABRIC_SERVICE_REPLICA_QUERY_DESCRIPTION
type FABRIC_SERVICE_REPLICA_QUERY_DESCRIPTION struct {
	PartitionId                 syscall.GUID
	ReplicaIdOrInstanceIdFilter int64
	Reserved                    unsafe.Pointer
}

// struct FABRIC_SERVICE_REPLICA_QUERY_DESCRIPTION_EX1
type FABRIC_SERVICE_REPLICA_QUERY_DESCRIPTION_EX1 struct {
	ReplicaStatusFilter uint32
	Reserved            unsafe.Pointer
}

// struct FABRIC_SERVICE_REPLICA_QUERY_RESULT_ITEM
type FABRIC_SERVICE_REPLICA_QUERY_RESULT_ITEM struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SERVICE_REPLICA_QUERY_RESULT_LIST
type FABRIC_SERVICE_REPLICA_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_REPLICA_QUERY_RESULT_ITEM
}

// struct FABRIC_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM
type FABRIC_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM struct {
	ReplicaId                    int64
	ReplicaRole                  int32
	ReplicaStatus                int32
	AggregatedHealthState        int32
	ReplicaAddress               win32.PWSTR
	NodeName                     win32.PWSTR
	LastInBuildDurationInSeconds int64
	Reserved                     unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM
type FABRIC_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM struct {
	InstanceId                   int64
	ReplicaStatus                int32
	AggregatedHealthState        int32
	ReplicaAddress               win32.PWSTR
	NodeName                     win32.PWSTR
	LastInBuildDurationInSeconds int64
	Reserved                     unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_APPLICATION_QUERY_DESCRIPTION struct {
	NodeName              win32.PWSTR
	ApplicationNameFilter win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_PAGED_DEPLOYED_APPLICATION_QUERY_DESCRIPTION
type FABRIC_PAGED_DEPLOYED_APPLICATION_QUERY_DESCRIPTION struct {
	NodeName              win32.PWSTR
	ApplicationNameFilter win32.PWSTR
	IncludeHealthState    int8
	PagingDescription     *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved              unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM struct {
	ApplicationName           win32.PWSTR
	ApplicationTypeName       win32.PWSTR
	DeployedApplicationStatus int32
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM_EX
type FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM_EX struct {
	WorkDirectory win32.PWSTR
	LogDirectory  win32.PWSTR
	TempDirectory win32.PWSTR
	Reserved      unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM_EX2
type FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM_EX2 struct {
	HealthState int32
	Reserved    unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_DESCRIPTION struct {
	NodeName                  win32.PWSTR
	ApplicationName           win32.PWSTR
	ServiceManifestNameFilter win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_ITEM struct {
	ServiceManifestName          win32.PWSTR
	ServiceManifestVersion       win32.PWSTR
	DeployedServicePackageStatus int32
	Reserved                     unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_ITEM_EX1
type FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_ITEM_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_DESCRIPTION struct {
	NodeName                  win32.PWSTR
	ApplicationName           win32.PWSTR
	ServiceManifestNameFilter win32.PWSTR
	ServiceTypeNameFilter     win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_ITEM struct {
	ServiceTypeName     win32.PWSTR
	CodePackageName     win32.PWSTR
	ServiceManifestName win32.PWSTR
	Status              int32
	Reserved            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_ITEM_EX1
type FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_ITEM_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_ITEM
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

// struct FABRIC_CODE_PACKAGE_ENTRY_POINT_EX1
type FABRIC_CODE_PACKAGE_ENTRY_POINT_EX1 struct {
	CodePackageInstanceId int64
	Reserved              unsafe.Pointer
}

// struct FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_DESCRIPTION struct {
	NodeName                  win32.PWSTR
	ApplicationName           win32.PWSTR
	ServiceManifestNameFilter win32.PWSTR
	CodePackageNameFilter     win32.PWSTR
	Reserved                  unsafe.Pointer
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

// struct FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_ITEM_EX1
type FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_ITEM_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	HostType                   int32
	HostIsolationMode          int32
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_DESCRIPTION struct {
	NodeName                  win32.PWSTR
	ApplicationName           win32.PWSTR
	ServiceManifestNameFilter win32.PWSTR
	PartitionIdFilter         syscall.GUID
	Reserved                  unsafe.Pointer
}

// struct FABRIC_RECONFIGURATION_INFORMATION_QUERY_RESULT
type FABRIC_RECONFIGURATION_INFORMATION_QUERY_RESULT struct {
	PreviousConfigurationRole   int32
	ReconfigurationPhase        int32
	ReconfigurationType         int32
	ReconfigurationStartTimeUtc FILETIME_
	Reserved                    unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_ITEM struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM struct {
	ServiceName            win32.PWSTR
	ServiceTypeName        win32.PWSTR
	ServiceManifestVersion win32.PWSTR
	CodePackageName        win32.PWSTR
	PartitionId            syscall.GUID
	ReplicaId              int64
	ReplicaRole            int32
	ReplicaStatus          int32
	Address                win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM_EX1
type FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM_EX1 struct {
	ServiceManifestName win32.PWSTR
	Reserved            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM_EX2
type FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM_EX2 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM_EX3
type FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM_EX3 struct {
	HostProcessId              int64
	ReconfigurationInformation *FABRIC_RECONFIGURATION_INFORMATION_QUERY_RESULT
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM struct {
	ServiceName            win32.PWSTR
	ServiceTypeName        win32.PWSTR
	ServiceManifestVersion win32.PWSTR
	CodePackageName        win32.PWSTR
	PartitionId            syscall.GUID
	InstanceId             int64
	ReplicaStatus          int32
	Address                win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM_EX1
type FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM_EX1 struct {
	ServiceManifestName win32.PWSTR
	Reserved            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM_EX2
type FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM_EX2 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM_EX3
type FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM_EX3 struct {
	HostProcessId int64
	Reserved      unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC_REPORT
type FABRIC_LOAD_METRIC_REPORT struct {
	Name            win32.PWSTR
	Value           uint32
	LastReportedUtc FILETIME_
	Reserved        unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC_REPORT_EX1
type FABRIC_LOAD_METRIC_REPORT_EX1 struct {
	CurrentValue float64
	Reserved     unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC_REPORT_LIST
type FABRIC_LOAD_METRIC_REPORT_LIST struct {
	Count uint32
	Items *FABRIC_LOAD_METRIC_REPORT
}

// struct FABRIC_UNPLACED_REPLICA_INFORMATION_LIST
type FABRIC_UNPLACED_REPLICA_INFORMATION_LIST struct {
	Count uint32
	Items *win32.PWSTR
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

// struct FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_DETAIL_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_DETAIL_QUERY_RESULT_ITEM struct {
	ServiceName                         win32.PWSTR
	PartitionId                         syscall.GUID
	InstanceId                          int64
	CurrentServiceOperation             int32
	CurrentServiceOperationStartTimeUtc FILETIME_
	ReportedLoad                        *FABRIC_LOAD_METRIC_REPORT_LIST
	Reserved                            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_DETAIL_QUERY_RESULT_ITEM_EX1
type FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_DETAIL_QUERY_RESULT_ITEM_EX1 struct {
	DeployedServiceReplica *FABRIC_DEPLOYED_STATELESS_SERVICE_INSTANCE_QUERY_RESULT_ITEM
	Reserved               unsafe.Pointer
}

// struct FABRIC_REMOTE_REPLICATOR_STATUS
type FABRIC_REMOTE_REPLICATOR_STATUS struct {
	ReplicaId                             int64
	LastAcknowledgementProcessedTimeUtc   FILETIME_
	LastReceivedReplicationSequenceNumber int64
	LastAppliedReplicationSequenceNumber  int64
	IsInBuild                             int8
	LastReceivedCopySequenceNumber        int64
	LastAppliedCopySequenceNumber         int64
	Reserved                              unsafe.Pointer
}

// struct FABRIC_REMOTE_REPLICATOR_ACKNOWLEDGEMENT_DETAIL
type FABRIC_REMOTE_REPLICATOR_ACKNOWLEDGEMENT_DETAIL struct {
	AverageReceiveDurationMilliseconds int64
	AverageApplyDurationMilliseconds   int64
	NotReceivedCount                   int64
	ReceivedAndNotAppliedCount         int64
	Reserved                           unsafe.Pointer
}

// struct FABRIC_REMOTE_REPLICATOR_ACKNOWLEDGEMENT_STATUS
type FABRIC_REMOTE_REPLICATOR_ACKNOWLEDGEMENT_STATUS struct {
	CopyStreamAcknowledgementDetails        *FABRIC_REMOTE_REPLICATOR_ACKNOWLEDGEMENT_DETAIL
	ReplicationStreamAcknowledgementDetails *FABRIC_REMOTE_REPLICATOR_ACKNOWLEDGEMENT_DETAIL
	Reserved                                unsafe.Pointer
}

// struct FABRIC_REMOTE_REPLICATOR_STATUS_LIST
type FABRIC_REMOTE_REPLICATOR_STATUS_LIST struct {
	Count uint32
	Items *FABRIC_REMOTE_REPLICATOR_STATUS
}

// struct FABRIC_REPLICATOR_QUEUE_STATUS
type FABRIC_REPLICATOR_QUEUE_STATUS struct {
	QueueUtilizationPercentage uint32
	QueueMemorySize            int64
	FirstSequenceNumber        int64
	CompletedSequenceNumber    int64
	CommittedSequenceNumber    int64
	LastSequenceNumber         int64
	Reserved                   unsafe.Pointer
}

// struct FABRIC_PRIMARY_REPLICATOR_STATUS_QUERY_RESULT
type FABRIC_PRIMARY_REPLICATOR_STATUS_QUERY_RESULT struct {
	ReplicationQueueStatus *FABRIC_REPLICATOR_QUEUE_STATUS
	RemoteReplicators      *FABRIC_REMOTE_REPLICATOR_STATUS_LIST
	Reserved               unsafe.Pointer
}

// struct FABRIC_SECONDARY_REPLICATOR_STATUS_QUERY_RESULT
type FABRIC_SECONDARY_REPLICATOR_STATUS_QUERY_RESULT struct {
	ReplicationQueueStatus                  *FABRIC_REPLICATOR_QUEUE_STATUS
	LastReplicationOperationReceivedTimeUtc FILETIME_
	IsInBuild                               int8
	CopyQueueStatus                         *FABRIC_REPLICATOR_QUEUE_STATUS
	LastCopyOperationReceivedTimeUtc        FILETIME_
	LastAcknowledgementSentTimeUtc          FILETIME_
	Reserved                                unsafe.Pointer
}

// struct FABRIC_REPLICATOR_STATUS_QUERY_RESULT
type FABRIC_REPLICATOR_STATUS_QUERY_RESULT struct {
	Role  int32
	Value unsafe.Pointer
}

// struct FABRIC_REPLICA_STATUS_QUERY_RESULT
type FABRIC_REPLICA_STATUS_QUERY_RESULT struct {
	Kind     int32
	Value    unsafe.Pointer
	Reserved unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM struct {
	ServiceName                         win32.PWSTR
	PartitionId                         syscall.GUID
	ReplicaId                           int64
	CurrentServiceOperation             int32
	CurrentServiceOperationStartTimeUtc FILETIME_
	CurrentReplicatorOperation          int32
	ReadStatus                          int32
	WriteStatus                         int32
	ReportedLoad                        *FABRIC_LOAD_METRIC_REPORT_LIST
	ReplicatorStatus                    *FABRIC_REPLICATOR_STATUS_QUERY_RESULT
	Reserved                            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM_EX1
type FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM_EX1 struct {
	ReplicaStatus *FABRIC_REPLICA_STATUS_QUERY_RESULT
	Reserved      unsafe.Pointer
}

// struct FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM_EX2
type FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM_EX2 struct {
	DeployedServiceReplica *FABRIC_DEPLOYED_STATEFUL_SERVICE_REPLICA_QUERY_RESULT_ITEM
	Reserved               unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_MIGRATION_QUERY_RESULT
type FABRIC_KEY_VALUE_STORE_MIGRATION_QUERY_RESULT struct {
	CurrentPhase int32
	State        int32
	NextPhase    int32
	Reserved     unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_STATUS_QUERY_RESULT
type FABRIC_KEY_VALUE_STORE_STATUS_QUERY_RESULT struct {
	DatabaseRowCountEstimate         int64
	DatabaseLogicalSizeEstimate      int64
	CopyNotificationCurrentKeyFilter win32.PWSTR
	CopyNotificationCurrentProgress  int64
	StatusDetails                    win32.PWSTR
	Reserved                         unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_STATUS_QUERY_RESULT_EX1
type FABRIC_KEY_VALUE_STORE_STATUS_QUERY_RESULT_EX1 struct {
	ProviderKind    int32
	MigrationStatus *FABRIC_KEY_VALUE_STORE_MIGRATION_QUERY_RESULT
	Reserved        unsafe.Pointer
}

// struct FABRIC_PROVISIONED_CODE_VERSION_QUERY_DESCRIPTION
type FABRIC_PROVISIONED_CODE_VERSION_QUERY_DESCRIPTION struct {
	CodeVersionFilter win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_ITEM
type FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_ITEM struct {
	CodeVersion win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_LIST
type FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_ITEM
}

// struct FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_DESCRIPTION
type FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_DESCRIPTION struct {
	ConfigVersionFilter win32.PWSTR
	Reserved            unsafe.Pointer
}

// struct FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_ITEM
type FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_ITEM struct {
	ConfigVersion win32.PWSTR
	Reserved      unsafe.Pointer
}

// struct FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_LIST
type FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_ITEM
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

// struct FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION
type FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION struct {
	NodeName              win32.PWSTR
	ApplicationName       win32.PWSTR
	ServiceManifestName   win32.PWSTR
	CodePackageName       win32.PWSTR
	CodePackageInstanceId int64
	Reserved              unsafe.Pointer
}

// struct FABRIC_RESTART_NODE_DESCRIPTION
type FABRIC_RESTART_NODE_DESCRIPTION struct {
	NodeName       win32.PWSTR
	NodeInstanceId uint64
	Reserved       unsafe.Pointer
}

// struct FABRIC_RESTART_NODE_DESCRIPTION_USING_NODE_NAME
type FABRIC_RESTART_NODE_DESCRIPTION_USING_NODE_NAME struct {
	NodeName               win32.PWSTR
	NodeInstanceId         uint64
	ShouldCreateFabricDump int8
	Reserved               unsafe.Pointer
}

// struct FABRIC_RESTART_NODE_DESCRIPTION2
type FABRIC_RESTART_NODE_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_RESTART_NODE_DESCRIPTION_EX1
type FABRIC_RESTART_NODE_DESCRIPTION_EX1 struct {
	CreateFabricDump int8
	Reserved         unsafe.Pointer
}

// struct FABRIC_STOP_NODE_DESCRIPTION
type FABRIC_STOP_NODE_DESCRIPTION struct {
	NodeName       win32.PWSTR
	NodeInstanceId uint64
	Reserved       unsafe.Pointer
}

// struct FABRIC_START_NODE_DESCRIPTION
type FABRIC_START_NODE_DESCRIPTION struct {
	NodeName              win32.PWSTR
	NodeInstanceId        uint64
	IpAddressOrFqdn       win32.PWSTR
	ClusterConnectionPort uint32
	Reserved              unsafe.Pointer
}

// struct FABRIC_START_NODE_DESCRIPTION2
type FABRIC_START_NODE_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_START_NODE_DESCRIPTION_USING_NODE_NAME
type FABRIC_START_NODE_DESCRIPTION_USING_NODE_NAME struct {
	NodeName              win32.PWSTR
	NodeInstanceId        uint64
	IpAddressOrFqdn       win32.PWSTR
	ClusterConnectionPort uint32
	Reserved              unsafe.Pointer
}

// struct FABRIC_STOP_NODE_DESCRIPTION2
type FABRIC_STOP_NODE_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_STOP_NODE_DESCRIPTION_USING_NODE_NAME
type FABRIC_STOP_NODE_DESCRIPTION_USING_NODE_NAME struct {
	NodeName       win32.PWSTR
	NodeInstanceId uint64
	Reserved       unsafe.Pointer
}

// struct FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION2
type FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_USING_NODE_NAME
type FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_USING_NODE_NAME struct {
	NodeName              win32.PWSTR
	ApplicationName       win32.PWSTR
	ServiceManifestName   win32.PWSTR
	CodePackageName       win32.PWSTR
	CodePackageInstanceId int64
	Reserved              unsafe.Pointer
}

// struct FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_USING_NODE_NAME_EX1
type FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_USING_NODE_NAME_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
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

// struct FABRIC_DEPLOYED_CODE_PACKAGE_RESULT_EX1
type FABRIC_DEPLOYED_CODE_PACKAGE_RESULT_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_MOVE_PRIMARY_DESCRIPTION2
type FABRIC_MOVE_PRIMARY_DESCRIPTION2 struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_MOVE_PRIMARY_DESCRIPTION_USING_NODE_NAME
type FABRIC_MOVE_PRIMARY_DESCRIPTION_USING_NODE_NAME struct {
	NodeName          win32.PWSTR
	ServiceName       win32.PWSTR
	PartitionId       syscall.GUID
	IgnoreConstraints int8
	Reserved          unsafe.Pointer
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

// struct FABRIC_MOVE_SECONDARY_DESCRIPTION_USING_NODE_NAME
type FABRIC_MOVE_SECONDARY_DESCRIPTION_USING_NODE_NAME struct {
	CurrentNodeName   win32.PWSTR
	NewNodeName       win32.PWSTR
	ServiceName       win32.PWSTR
	PartitionId       syscall.GUID
	IgnoreConstraints int8
	Reserved          unsafe.Pointer
}

// struct FABRIC_MOVE_SECONDARY_RESULT
type FABRIC_MOVE_SECONDARY_RESULT struct {
	CurrentNodeName win32.PWSTR
	NewNodeName     win32.PWSTR
	ServiceName     win32.PWSTR
	PartitionId     syscall.GUID
	Reserved        unsafe.Pointer
}

// struct FABRIC_RESTART_REPLICA_DESCRIPTION
type FABRIC_RESTART_REPLICA_DESCRIPTION struct {
	NodeName            win32.PWSTR
	PartitionId         syscall.GUID
	ReplicaOrInstanceId int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_REMOVE_REPLICA_DESCRIPTION
type FABRIC_REMOVE_REPLICA_DESCRIPTION struct {
	NodeName            win32.PWSTR
	PartitionId         syscall.GUID
	ReplicaOrInstanceId int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_REMOVE_REPLICA_DESCRIPTION_EX1
type FABRIC_REMOVE_REPLICA_DESCRIPTION_EX1 struct {
	ForceRemove int8
	Reserved    unsafe.Pointer
}

// struct FABRIC_HEALTH_REPORT_SEND_OPTIONS
type FABRIC_HEALTH_REPORT_SEND_OPTIONS struct {
	Immediate int8
	Reserved  unsafe.Pointer
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

// struct FABRIC_HEALTH_REPORT
type FABRIC_HEALTH_REPORT struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH_REPORT
type FABRIC_CLUSTER_HEALTH_REPORT struct {
	HealthInformation *FABRIC_HEALTH_INFORMATION
	Reserved          unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH_REPORT
type FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH_REPORT struct {
	PartitionId       syscall.GUID
	InstanceId        int64
	HealthInformation *FABRIC_HEALTH_INFORMATION
	Reserved          unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH_REPORT
type FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH_REPORT struct {
	PartitionId       syscall.GUID
	ReplicaId         int64
	HealthInformation *FABRIC_HEALTH_INFORMATION
	Reserved          unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_REPORT
type FABRIC_PARTITION_HEALTH_REPORT struct {
	PartitionId       syscall.GUID
	HealthInformation *FABRIC_HEALTH_INFORMATION
	Reserved          unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_REPORT
type FABRIC_NODE_HEALTH_REPORT struct {
	NodeName          win32.PWSTR
	HealthInformation *FABRIC_HEALTH_INFORMATION
	Reserved          unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_REPORT
type FABRIC_APPLICATION_HEALTH_REPORT struct {
	ApplicationName   win32.PWSTR
	HealthInformation *FABRIC_HEALTH_INFORMATION
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_REPORT
type FABRIC_SERVICE_HEALTH_REPORT struct {
	ServiceName       win32.PWSTR
	HealthInformation *FABRIC_HEALTH_INFORMATION
	Reserved          unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_REPORT
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_REPORT struct {
	ApplicationName     win32.PWSTR
	ServiceManifestName win32.PWSTR
	NodeName            win32.PWSTR
	HealthInformation   *FABRIC_HEALTH_INFORMATION
	Reserved            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_REPORT_EX1
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_REPORT_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_REPORT
type FABRIC_DEPLOYED_APPLICATION_HEALTH_REPORT struct {
	ApplicationName   win32.PWSTR
	NodeName          win32.PWSTR
	HealthInformation *FABRIC_HEALTH_INFORMATION
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_HEALTH_POLICY
type FABRIC_SERVICE_TYPE_HEALTH_POLICY struct {
	MaxPercentUnhealthyServices             byte
	MaxPercentUnhealthyPartitionsPerService byte
	MaxPercentUnhealthyReplicasPerPartition byte
	Reserved                                unsafe.Pointer
}

// struct FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP_ITEM
type FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP_ITEM struct {
	ServiceTypeName         win32.PWSTR
	ServiceTypeHealthPolicy *FABRIC_SERVICE_TYPE_HEALTH_POLICY
}

// struct FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP
type FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP struct {
	Count uint32
	Items *FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP_ITEM
}

// struct FABRIC_APPLICATION_HEALTH_POLICY
type FABRIC_APPLICATION_HEALTH_POLICY struct {
	ConsiderWarningAsError                  int8
	MaxPercentUnhealthyDeployedApplications byte
	DefaultServiceTypeHealthPolicy          *FABRIC_SERVICE_TYPE_HEALTH_POLICY
	ServiceTypeHealthPolicyMap              *FABRIC_SERVICE_TYPE_HEALTH_POLICY_MAP
	Reserved                                unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_POLICY_MAP_ITEM
type FABRIC_APPLICATION_HEALTH_POLICY_MAP_ITEM struct {
	ApplicationName win32.PWSTR
	HealthPolicy    *FABRIC_APPLICATION_HEALTH_POLICY
}

// struct FABRIC_APPLICATION_HEALTH_POLICY_MAP
type FABRIC_APPLICATION_HEALTH_POLICY_MAP struct {
	Count uint32
	Items *FABRIC_APPLICATION_HEALTH_POLICY_MAP_ITEM
}

// struct FABRIC_CLUSTER_HEALTH_POLICY
type FABRIC_CLUSTER_HEALTH_POLICY struct {
	ConsiderWarningAsError          int8
	MaxPercentUnhealthyNodes        byte
	MaxPercentUnhealthyApplications byte
	Reserved                        unsafe.Pointer
}

// struct FABRIC_CLUSTER_UPGRADE_HEALTH_POLICY
type FABRIC_CLUSTER_UPGRADE_HEALTH_POLICY struct {
	MaxPercentDeltaUnhealthyNodes              byte
	MaxPercentUpgradeDomainDeltaUnhealthyNodes byte
	Reserved                                   unsafe.Pointer
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

// struct FABRIC_LOAD_METRIC_INFORMATION_EX1
type FABRIC_LOAD_METRIC_INFORMATION_EX1 struct {
	ActivityThreshold uint32
	ClusterCapacity   int64
	ClusterLoad       int64
	Reserved          unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC_INFORMATION_EX2
type FABRIC_LOAD_METRIC_INFORMATION_EX2 struct {
	RemainingUnbufferedCapacity int64
	NodeBufferPercentage        float64
	BufferedCapacity            int64
	RemainingBufferedCapacity   int64
	IsClusterCapacityViolation  int8
	MinNodeLoadValue            int64
	MinNodeLoadNodeId           FABRIC_NODE_ID
	MaxNodeLoadValue            int64
	MaxNodeLoadNodeId           FABRIC_NODE_ID
	Reserved                    unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC_INFORMATION_EX3
type FABRIC_LOAD_METRIC_INFORMATION_EX3 struct {
	CurrentClusterLoad               float64
	BufferedClusterCapacityRemaining float64
	ClusterCapacityRemaining         float64
	MaximumNodeLoad                  float64
	MinimumNodeLoad                  float64
	Reserved                         unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC_INFORMATION_LIST
type FABRIC_LOAD_METRIC_INFORMATION_LIST struct {
	Count uint32
	Items *FABRIC_LOAD_METRIC_INFORMATION
}

// struct FABRIC_CLUSTER_LOAD_INFORMATION
type FABRIC_CLUSTER_LOAD_INFORMATION struct {
	LastBalancingStartTimeUtc FILETIME_
	LastBalancingEndTimeUtc   FILETIME_
	LoadMetricInformation     *FABRIC_LOAD_METRIC_INFORMATION_LIST
	Reserved                  unsafe.Pointer
}

// struct FABRIC_PARTITION_LOAD_INFORMATION
type FABRIC_PARTITION_LOAD_INFORMATION struct {
	PartitionId                syscall.GUID
	PrimaryLoadMetricReports   *FABRIC_LOAD_METRIC_REPORT_LIST
	SecondaryLoadMetricReports *FABRIC_LOAD_METRIC_REPORT_LIST
	Reserved                   unsafe.Pointer
}

// struct FABRIC_REPLICA_LOAD_INFORMATION
type FABRIC_REPLICA_LOAD_INFORMATION struct {
	PartitionId         syscall.GUID
	ReplicaOrInstanceId int64
	LoadMetricReports   *FABRIC_LOAD_METRIC_REPORT_LIST
	Reserved            unsafe.Pointer
}

// struct FABRIC_UNPLACED_REPLICA_INFORMATION
type FABRIC_UNPLACED_REPLICA_INFORMATION struct {
	ServiceName            win32.PWSTR
	PartitionId            syscall.GUID
	UnplacedReplicaReasons *FABRIC_STRING_LIST
	Reserved               unsafe.Pointer
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

// struct FABRIC_NODE_LOAD_METRIC_INFORMATION_EX1
type FABRIC_NODE_LOAD_METRIC_INFORMATION_EX1 struct {
	NodeBufferedCapacity          int64
	NodeRemainingBufferedCapacity int64
	Reserved                      unsafe.Pointer
}

// struct FABRIC_NODE_LOAD_METRIC_INFORMATION_EX2
type FABRIC_NODE_LOAD_METRIC_INFORMATION_EX2 struct {
	CurrentNodeLoad               float64
	NodeCapacityRemaining         float64
	BufferedNodeCapacityRemaining float64
	Reserved                      unsafe.Pointer
}

// struct FABRIC_NODE_LOAD_METRIC_INFORMATION_LIST
type FABRIC_NODE_LOAD_METRIC_INFORMATION_LIST struct {
	Count uint32
	Items *FABRIC_NODE_LOAD_METRIC_INFORMATION
}

// struct FABRIC_NODE_LOAD_INFORMATION
type FABRIC_NODE_LOAD_INFORMATION struct {
	NodeName                  win32.PWSTR
	NodeLoadMetricInformation *FABRIC_NODE_LOAD_METRIC_INFORMATION_LIST
	Reserved                  unsafe.Pointer
}

// struct FABRIC_APPLICATION_LOAD_METRIC_INFORMATION
type FABRIC_APPLICATION_LOAD_METRIC_INFORMATION struct {
	Name                win32.PWSTR
	ReservationCapacity int64
	ApplicationCapacity int64
	ApplicationLoad     int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_APPLICATION_LOAD_METRIC_INFORMATION_LIST
type FABRIC_APPLICATION_LOAD_METRIC_INFORMATION_LIST struct {
	Count       uint32
	LoadMetrics *FABRIC_APPLICATION_LOAD_METRIC_INFORMATION
	Reserved    unsafe.Pointer
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

// struct FABRIC_GATEWAY_INFORMATION
type FABRIC_GATEWAY_INFORMATION struct {
	NodeAddress    win32.PWSTR
	NodeId         FABRIC_NODE_ID
	NodeInstanceId uint64
	NodeName       win32.PWSTR
	Reserved       unsafe.Pointer
}

// struct FABRIC_START_APPROVED_UPGRADES_DESCRIPTION
type FABRIC_START_APPROVED_UPGRADES_DESCRIPTION struct {
	OperationId       syscall.GUID
	ClusterConfigPath win32.PWSTR
	RollbackOnFailure int8
	Reserved          unsafe.Pointer
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

// struct FABRIC_START_UPGRADE_DESCRIPTION_EX1
type FABRIC_START_UPGRADE_DESCRIPTION_EX1 struct {
	ApplicationHealthPolicyMap *FABRIC_APPLICATION_HEALTH_POLICY_MAP
	Reserved                   unsafe.Pointer
}

// struct FABRIC_UPGRADE_ORCHESTRATION_SERVICE_STATE
type FABRIC_UPGRADE_ORCHESTRATION_SERVICE_STATE struct {
	CurrentCodeVersion     win32.PWSTR
	CurrentManifestVersion win32.PWSTR
	TargetCodeVersion      win32.PWSTR
	TargetManifestVersion  win32.PWSTR
	PendingUpgradeType     win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_HEALTH_EVENT
type FABRIC_HEALTH_EVENT struct {
	HealthInformation        *FABRIC_HEALTH_INFORMATION
	SourceUtcTimestamp       FILETIME_
	LastModifiedUtcTimestamp FILETIME_
	IsExpired                int8
	Reserved                 unsafe.Pointer
}

// struct FABRIC_HEALTH_EVENT_EX1
type FABRIC_HEALTH_EVENT_EX1 struct {
	LastOkTransitionAt      FILETIME_
	LastWarningTransitionAt FILETIME_
	LastErrorTransitionAt   FILETIME_
	Reserved                unsafe.Pointer
}

// struct FABRIC_HEALTH_EVENT_LIST
type FABRIC_HEALTH_EVENT_LIST struct {
	Count uint32
	Items *FABRIC_HEALTH_EVENT
}

// struct FABRIC_HEALTH_EVALUATION
type FABRIC_HEALTH_EVALUATION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_HEALTH_EVALUATION_LIST
type FABRIC_HEALTH_EVALUATION_LIST struct {
	Count uint32
	Items *FABRIC_HEALTH_EVALUATION
}

// struct FABRIC_HEALTH_STATE_COUNT
type FABRIC_HEALTH_STATE_COUNT struct {
	OkCount      uint32
	WarningCount uint32
	ErrorCount   uint32
	Reserved     unsafe.Pointer
}

// struct FABRIC_ENTITY_KIND_HEALTH_STATE_COUNT
type FABRIC_ENTITY_KIND_HEALTH_STATE_COUNT struct {
	EntityKind       int32
	HealthStateCount *FABRIC_HEALTH_STATE_COUNT
	Reserved         unsafe.Pointer
}

// struct FABRIC_HEALTH_STATISTICS
type FABRIC_HEALTH_STATISTICS struct {
	Count    uint32
	Items    *FABRIC_ENTITY_KIND_HEALTH_STATE_COUNT
	Reserved unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH
type FABRIC_REPLICA_HEALTH struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH
type FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH struct {
	PartitionId           syscall.GUID
	ReplicaId             int64
	AggregatedHealthState int32
	HealthEvents          *FABRIC_HEALTH_EVENT_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH_EX1
type FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH
type FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH struct {
	PartitionId           syscall.GUID
	InstanceId            int64
	AggregatedHealthState int32
	HealthEvents          *FABRIC_HEALTH_EVENT_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH_EX1
type FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATE
type FABRIC_REPLICA_HEALTH_STATE struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATE_LIST
type FABRIC_REPLICA_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_REPLICA_HEALTH_STATE
}

// struct FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH_STATE
type FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH_STATE struct {
	PartitionId           syscall.GUID
	InstanceId            int64
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH_STATE_EX1
type FABRIC_STATELESS_SERVICE_INSTANCE_HEALTH_STATE_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH_STATE
type FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH_STATE struct {
	PartitionId           syscall.GUID
	ReplicaId             int64
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH_STATE_EX1
type FABRIC_STATEFUL_SERVICE_REPLICA_HEALTH_STATE_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH
type FABRIC_PARTITION_HEALTH struct {
	PartitionId           syscall.GUID
	AggregatedHealthState int32
	HealthEvents          *FABRIC_HEALTH_EVENT_LIST
	ReplicaHealthStates   *FABRIC_REPLICA_HEALTH_STATE_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_EX1
type FABRIC_PARTITION_HEALTH_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_EX2
type FABRIC_PARTITION_HEALTH_EX2 struct {
	HealthStatistics *FABRIC_HEALTH_STATISTICS
	Reserved         unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATE
type FABRIC_PARTITION_HEALTH_STATE struct {
	PartitionId           syscall.GUID
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATE_LIST
type FABRIC_PARTITION_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_PARTITION_HEALTH_STATE
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

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_EX1
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_EX2
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_EX2 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE struct {
	ApplicationName       win32.PWSTR
	ServiceManifestName   win32.PWSTR
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_EX1
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_LIST
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE
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

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_EX1
type FABRIC_DEPLOYED_APPLICATION_HEALTH_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_EX2
type FABRIC_DEPLOYED_APPLICATION_HEALTH_EX2 struct {
	HealthStatistics *FABRIC_HEALTH_STATISTICS
	Reserved         unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE struct {
	ApplicationName       win32.PWSTR
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_LIST
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE
}

// struct FABRIC_SERVICE_HEALTH
type FABRIC_SERVICE_HEALTH struct {
	ServiceName           win32.PWSTR
	AggregatedHealthState int32
	HealthEvents          *FABRIC_HEALTH_EVENT_LIST
	PartitionHealthStates *FABRIC_PARTITION_HEALTH_STATE_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_EX1
type FABRIC_SERVICE_HEALTH_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_EX2
type FABRIC_SERVICE_HEALTH_EX2 struct {
	HealthStatistics *FABRIC_HEALTH_STATISTICS
	Reserved         unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATE
type FABRIC_SERVICE_HEALTH_STATE struct {
	ServiceName           win32.PWSTR
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATE_LIST
type FABRIC_SERVICE_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_HEALTH_STATE
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

// struct FABRIC_APPLICATION_HEALTH_EX1
type FABRIC_APPLICATION_HEALTH_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_EX2
type FABRIC_APPLICATION_HEALTH_EX2 struct {
	HealthStatistics *FABRIC_HEALTH_STATISTICS
	Reserved         unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH
type FABRIC_NODE_HEALTH struct {
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	HealthEvents          *FABRIC_HEALTH_EVENT_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_EX1
type FABRIC_NODE_HEALTH_EX1 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH
type FABRIC_CLUSTER_HEALTH struct {
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATE
type FABRIC_NODE_HEALTH_STATE struct {
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATE_LIST
type FABRIC_NODE_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_NODE_HEALTH_STATE
}

// struct FABRIC_APPLICATION_HEALTH_STATE
type FABRIC_APPLICATION_HEALTH_STATE struct {
	ApplicationName       win32.PWSTR
	AggregatedHealthState int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_STATE_LIST
type FABRIC_APPLICATION_HEALTH_STATE_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_HEALTH_STATE
}

// struct FABRIC_CLUSTER_HEALTH_EX1
type FABRIC_CLUSTER_HEALTH_EX1 struct {
	NodeHealthStates        *FABRIC_NODE_HEALTH_STATE_LIST
	ApplicationHealthStates *FABRIC_APPLICATION_HEALTH_STATE_LIST
	HealthEvents            *FABRIC_HEALTH_EVENT_LIST
	Reserved                unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH_EX2
type FABRIC_CLUSTER_HEALTH_EX2 struct {
	UnhealthyEvaluations *FABRIC_HEALTH_EVALUATION_LIST
	Reserved             unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH_EX3
type FABRIC_CLUSTER_HEALTH_EX3 struct {
	HealthStatistics *FABRIC_HEALTH_STATISTICS
	Reserved         unsafe.Pointer
}

// struct FABRIC_EVENT_HEALTH_EVALUATION
type FABRIC_EVENT_HEALTH_EVALUATION struct {
	Description            win32.PWSTR
	AggregatedHealthState  int32
	UnhealthyEvent         *FABRIC_HEALTH_EVENT
	ConsiderWarningAsError int8
	Reserved               unsafe.Pointer
}

// struct FABRIC_REPLICAS_HEALTH_EVALUATION
type FABRIC_REPLICAS_HEALTH_EVALUATION struct {
	Description                             win32.PWSTR
	AggregatedHealthState                   int32
	UnhealthyEvaluations                    *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount                              uint32
	MaxPercentUnhealthyReplicasPerPartition byte
	Reserved                                unsafe.Pointer
}

// struct FABRIC_PARTITIONS_HEALTH_EVALUATION
type FABRIC_PARTITIONS_HEALTH_EVALUATION struct {
	Description                             win32.PWSTR
	AggregatedHealthState                   int32
	UnhealthyEvaluations                    *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount                              uint32
	MaxPercentUnhealthyPartitionsPerService byte
	Reserved                                unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGES_HEALTH_EVALUATION
type FABRIC_DEPLOYED_SERVICE_PACKAGES_HEALTH_EVALUATION struct {
	Description           win32.PWSTR
	AggregatedHealthState int32
	UnhealthyEvaluations  *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount            uint32
	Reserved              unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATIONS_HEALTH_EVALUATION
type FABRIC_DEPLOYED_APPLICATIONS_HEALTH_EVALUATION struct {
	Description                             win32.PWSTR
	AggregatedHealthState                   int32
	UnhealthyEvaluations                    *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount                              uint32
	MaxPercentUnhealthyDeployedApplications byte
	Reserved                                unsafe.Pointer
}

// struct FABRIC_SERVICES_HEALTH_EVALUATION
type FABRIC_SERVICES_HEALTH_EVALUATION struct {
	Description                 win32.PWSTR
	AggregatedHealthState       int32
	ServiceTypeName             win32.PWSTR
	UnhealthyEvaluations        *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount                  uint32
	MaxPercentUnhealthyServices byte
	Reserved                    unsafe.Pointer
}

// struct FABRIC_NODES_HEALTH_EVALUATION
type FABRIC_NODES_HEALTH_EVALUATION struct {
	Description              win32.PWSTR
	AggregatedHealthState    int32
	UnhealthyEvaluations     *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount               uint32
	MaxPercentUnhealthyNodes byte
	Reserved                 unsafe.Pointer
}

// struct FABRIC_APPLICATIONS_HEALTH_EVALUATION
type FABRIC_APPLICATIONS_HEALTH_EVALUATION struct {
	Description                     win32.PWSTR
	AggregatedHealthState           int32
	UnhealthyEvaluations            *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount                      uint32
	MaxPercentUnhealthyApplications byte
	Reserved                        unsafe.Pointer
}

// struct FABRIC_UPGRADE_DOMAIN_NODES_HEALTH_EVALUATION
type FABRIC_UPGRADE_DOMAIN_NODES_HEALTH_EVALUATION struct {
	Description              win32.PWSTR
	AggregatedHealthState    int32
	UpgradeDomainName        win32.PWSTR
	UnhealthyEvaluations     *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount               uint32
	MaxPercentUnhealthyNodes byte
	Reserved                 unsafe.Pointer
}

// struct FABRIC_UPGRADE_DOMAIN_DEPLOYED_APPLICATIONS_HEALTH_EVALUATION
type FABRIC_UPGRADE_DOMAIN_DEPLOYED_APPLICATIONS_HEALTH_EVALUATION struct {
	Description                             win32.PWSTR
	AggregatedHealthState                   int32
	UpgradeDomainName                       win32.PWSTR
	UnhealthyEvaluations                    *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount                              uint32
	MaxPercentUnhealthyDeployedApplications byte
	Reserved                                unsafe.Pointer
}

// struct FABRIC_SYSTEM_APPLICATION_HEALTH_EVALUATION
type FABRIC_SYSTEM_APPLICATION_HEALTH_EVALUATION struct {
	Description           win32.PWSTR
	AggregatedHealthState int32
	UnhealthyEvaluations  *FABRIC_HEALTH_EVALUATION_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_EVALUATION
type FABRIC_NODE_HEALTH_EVALUATION struct {
	Description           win32.PWSTR
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	UnhealthyEvaluations  *FABRIC_HEALTH_EVALUATION_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_EVALUATION
type FABRIC_REPLICA_HEALTH_EVALUATION struct {
	Description           win32.PWSTR
	PartitionId           syscall.GUID
	ReplicaOrInstanceId   int64
	AggregatedHealthState int32
	UnhealthyEvaluations  *FABRIC_HEALTH_EVALUATION_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_EVALUATION
type FABRIC_PARTITION_HEALTH_EVALUATION struct {
	Description           win32.PWSTR
	PartitionId           syscall.GUID
	AggregatedHealthState int32
	UnhealthyEvaluations  *FABRIC_HEALTH_EVALUATION_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_EVALUATION
type FABRIC_SERVICE_HEALTH_EVALUATION struct {
	Description           win32.PWSTR
	ServiceName           win32.PWSTR
	AggregatedHealthState int32
	UnhealthyEvaluations  *FABRIC_HEALTH_EVALUATION_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_EVALUATION
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_EVALUATION struct {
	Description           win32.PWSTR
	ApplicationName       win32.PWSTR
	ServiceManifestName   win32.PWSTR
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	UnhealthyEvaluations  *FABRIC_HEALTH_EVALUATION_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_EVALUATION_EX1
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_EVALUATION_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_EVALUATION
type FABRIC_DEPLOYED_APPLICATION_HEALTH_EVALUATION struct {
	Description           win32.PWSTR
	ApplicationName       win32.PWSTR
	NodeName              win32.PWSTR
	AggregatedHealthState int32
	UnhealthyEvaluations  *FABRIC_HEALTH_EVALUATION_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_EVALUATION
type FABRIC_APPLICATION_HEALTH_EVALUATION struct {
	Description           win32.PWSTR
	ApplicationName       win32.PWSTR
	AggregatedHealthState int32
	UnhealthyEvaluations  *FABRIC_HEALTH_EVALUATION_LIST
	Reserved              unsafe.Pointer
}

// struct FABRIC_X509_CREDENTIALS
type FABRIC_X509_CREDENTIALS struct {
	AllowedCommonNameCount uint32
	AllowedCommonNames     *win32.PWSTR
	FindType               int32
	FindValue              unsafe.Pointer
	StoreLocation          int32
	StoreName              win32.PWSTR
	ProtectionLevel        int32
	Reserved               unsafe.Pointer
}

// struct FABRIC_X509_CREDENTIALS_EX1
type FABRIC_X509_CREDENTIALS_EX1 struct {
	IssuerThumbprintCount uint32
	IssuerThumbprints     *win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_X509_NAME
type FABRIC_X509_NAME struct {
	Name                 win32.PWSTR
	IssuerCertThumbprint win32.PWSTR
	Reserved             unsafe.Pointer
}

// struct FABRIC_X509_CREDENTIALS_EX2
type FABRIC_X509_CREDENTIALS_EX2 struct {
	RemoteCertThumbprintCount uint32
	RemoteCertThumbprints     *win32.PWSTR
	RemoteX509NameCount       uint32
	RemoteX509Names           *FABRIC_X509_NAME
	FindValueSecondary        unsafe.Pointer
	Reserved                  unsafe.Pointer
}

// struct FABRIC_X509_ISSUER_NAME
type FABRIC_X509_ISSUER_NAME struct {
	Name             win32.PWSTR
	IssuerStoreCount uint32
	IssuerStores     *win32.PWSTR
	Reserved         unsafe.Pointer
}

// struct FABRIC_X509_CREDENTIALS_EX3
type FABRIC_X509_CREDENTIALS_EX3 struct {
	RemoteCertIssuerCount uint32
	RemoteCertIssuers     *FABRIC_X509_ISSUER_NAME
	Reserved              unsafe.Pointer
}

// struct FABRIC_X509_CREDENTIALS2
type FABRIC_X509_CREDENTIALS2 struct {
	CertLoadPath              win32.PWSTR
	RemoteCertThumbprintCount uint32
	RemoteCertThumbprints     *win32.PWSTR
	RemoteX509NameCount       uint32
	RemoteX509Names           *FABRIC_X509_NAME
	ProtectionLevel           int32
	Reserved                  unsafe.Pointer
}

// struct FABRIC_CLAIMS_CREDENTIALS
type FABRIC_CLAIMS_CREDENTIALS struct {
	ServerCommonNameCount uint32
	ServerCommonNames     *win32.PWSTR
	IssuerThumbprintCount uint32
	IssuerThumbprints     *win32.PWSTR
	LocalClaims           win32.PWSTR
	ProtectionLevel       int32
	Reserved              unsafe.Pointer
}

// struct FABRIC_CLAIMS_CREDENTIALS_EX1
type FABRIC_CLAIMS_CREDENTIALS_EX1 struct {
	ServerThumbprintCount uint32
	ServerThumbprints     *win32.PWSTR
	Reserved              unsafe.Pointer
}

// struct FABRIC_WINDOWS_CREDENTIALS
type FABRIC_WINDOWS_CREDENTIALS struct {
	RemoteSpn           win32.PWSTR
	RemoteIdentityCount uint32
	RemoteIdentities    *win32.PWSTR
	ProtectionLevel     int32
	Reserved            unsafe.Pointer
}

// struct FABRIC_CLAIMS_RETRIEVAL_METADATA
type FABRIC_CLAIMS_RETRIEVAL_METADATA struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_AAD_CLAIMS_RETRIEVAL_METADATA
type FABRIC_AAD_CLAIMS_RETRIEVAL_METADATA struct {
	Authority          win32.PWSTR
	TenantId           win32.PWSTR
	ClusterApplication win32.PWSTR
	ClientApplication  win32.PWSTR
	ClientRedirect     win32.PWSTR
	Reserved           unsafe.Pointer
}

// struct FABRIC_AAD_CLAIMS_RETRIEVAL_METADATA_EX1
type FABRIC_AAD_CLAIMS_RETRIEVAL_METADATA_EX1 struct {
	LoginEndpoint win32.PWSTR
	Reserved      unsafe.Pointer
}

// struct FABRIC_SECURITY_USER_DESCRIPTION
type FABRIC_SECURITY_USER_DESCRIPTION struct {
	Name                    win32.PWSTR
	Sid                     win32.PWSTR
	ParentSystemGroups      *FABRIC_STRING_LIST
	ParentApplicationGroups *FABRIC_STRING_LIST
	Reserved                unsafe.Pointer
}

// struct FABRIC_SECURITY_USER_DESCRIPTION_LIST
type FABRIC_SECURITY_USER_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SECURITY_USER_DESCRIPTION
}

// struct FABRIC_SECURITY_GROUP_DESCRIPTION
type FABRIC_SECURITY_GROUP_DESCRIPTION struct {
	Name               win32.PWSTR
	Sid                win32.PWSTR
	DomainGroupMembers *FABRIC_STRING_LIST
	SystemGroupMembers *FABRIC_STRING_LIST
	DomainUserMembers  *FABRIC_STRING_LIST
	Reserved           unsafe.Pointer
}

// struct FABRIC_SECURITY_GROUP_DESCRIPTION_LIST
type FABRIC_SECURITY_GROUP_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SECURITY_GROUP_DESCRIPTION
}

// struct FABRIC_APPLICATION_PRINCIPALS_DESCRIPTION
type FABRIC_APPLICATION_PRINCIPALS_DESCRIPTION struct {
	Users    *FABRIC_SECURITY_USER_DESCRIPTION_LIST
	Groups   *FABRIC_SECURITY_GROUP_DESCRIPTION_LIST
	Reserved unsafe.Pointer
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

// struct FABRIC_APPLICATION_UPGRADE_PROGRESS_EX1
type FABRIC_APPLICATION_UPGRADE_PROGRESS_EX1 struct {
	StartTimestampUtc              FILETIME_
	FailureTimestampUtc            FILETIME_
	FailureReason                  int32
	UpgradeDomainProgressAtFailure *FABRIC_UPGRADE_DOMAIN_PROGRESS
	Reserved                       unsafe.Pointer
}

// struct FABRIC_APPLICATION_UPGRADE_PROGRESS_EX2
type FABRIC_APPLICATION_UPGRADE_PROGRESS_EX2 struct {
	UpgradeStatusDetails win32.PWSTR
	Reserved             unsafe.Pointer
}

// struct FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION_BASE
type FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION_BASE struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION
type FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION struct {
	BuildPath win32.PWSTR
	Async     int8
	Reserved  unsafe.Pointer
}

// struct FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION_EX1
type FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION_EX1 struct {
	ApplicationPackageCleanupPolicy int32
	Reserved                        unsafe.Pointer
}

// struct FABRIC_EXTERNAL_STORE_PROVISION_APPLICATION_TYPE_DESCRIPTION
type FABRIC_EXTERNAL_STORE_PROVISION_APPLICATION_TYPE_DESCRIPTION struct {
	ApplicationTypeName           win32.PWSTR
	ApplicationTypeVersion        win32.PWSTR
	ApplicationPackageDownloadUri win32.PWSTR
	Async                         int8
	Reserved                      unsafe.Pointer
}

// struct FABRIC_UNPROVISION_APPLICATION_TYPE_DESCRIPTION
type FABRIC_UNPROVISION_APPLICATION_TYPE_DESCRIPTION struct {
	ApplicationTypeName    win32.PWSTR
	ApplicationTypeVersion win32.PWSTR
	Async                  int8
	Reserved               unsafe.Pointer
}

// struct FABRIC_DELETE_APPLICATION_DESCRIPTION
type FABRIC_DELETE_APPLICATION_DESCRIPTION struct {
	ApplicationName win32.PWSTR
	ForceDelete     int8
	Reserved        unsafe.Pointer
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

// struct FABRIC_UPGRADE_PROGRESS_EX1
type FABRIC_UPGRADE_PROGRESS_EX1 struct {
	StartTimestampUtc              FILETIME_
	FailureTimestampUtc            FILETIME_
	FailureReason                  int32
	UpgradeDomainProgressAtFailure *FABRIC_UPGRADE_DOMAIN_PROGRESS
	Reserved                       unsafe.Pointer
}

// struct FABRIC_PACKAGE_SHARING_POLICY
type FABRIC_PACKAGE_SHARING_POLICY struct {
	PackageName win32.PWSTR
	Scope       int32
	Reserved    unsafe.Pointer
}

// struct FABRIC_PACKAGE_SHARING_POLICY_LIST
type FABRIC_PACKAGE_SHARING_POLICY_LIST struct {
	Count uint32
	Items *FABRIC_PACKAGE_SHARING_POLICY
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

// struct FABRIC_REPLICA_HEALTH_STATES_FILTER
type FABRIC_REPLICA_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATES_FILTER
type FABRIC_PARTITION_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATES_FILTER
type FABRIC_SERVICE_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_STATES_FILTER
type FABRIC_APPLICATION_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATES_FILTER
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATES_FILTER
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATES_FILTER struct {
	HealthStateFilter uint32
	Reserved          unsafe.Pointer
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

// struct FABRIC_CLUSTER_HEALTH_STATISTICS_FILTER
type FABRIC_CLUSTER_HEALTH_STATISTICS_FILTER struct {
	ExcludeHealthStatistics                  int8
	IncludeSystemApplicationHealthStatistics int8
	Reserved                                 unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH_QUERY_DESCRIPTION_EX1
type FABRIC_CLUSTER_HEALTH_QUERY_DESCRIPTION_EX1 struct {
	HealthStatisticsFilter *FABRIC_CLUSTER_HEALTH_STATISTICS_FILTER
	Reserved               unsafe.Pointer
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

// struct FABRIC_APPLICATION_HEALTH_STATISTICS_FILTER
type FABRIC_APPLICATION_HEALTH_STATISTICS_FILTER struct {
	ExcludeHealthStatistics int8
	Reserved                unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_QUERY_DESCRIPTION_EX1
type FABRIC_APPLICATION_HEALTH_QUERY_DESCRIPTION_EX1 struct {
	HealthStatisticsFilter *FABRIC_APPLICATION_HEALTH_STATISTICS_FILTER
	Reserved               unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_QUERY_DESCRIPTION
type FABRIC_SERVICE_HEALTH_QUERY_DESCRIPTION struct {
	ServiceName      win32.PWSTR
	HealthPolicy     *FABRIC_APPLICATION_HEALTH_POLICY
	EventsFilter     *FABRIC_HEALTH_EVENTS_FILTER
	PartitionsFilter *FABRIC_PARTITION_HEALTH_STATES_FILTER
	Reserved         unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATISTICS_FILTER
type FABRIC_SERVICE_HEALTH_STATISTICS_FILTER struct {
	ExcludeHealthStatistics int8
	Reserved                unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_QUERY_DESCRIPTION_EX1
type FABRIC_SERVICE_HEALTH_QUERY_DESCRIPTION_EX1 struct {
	HealthStatisticsFilter *FABRIC_SERVICE_HEALTH_STATISTICS_FILTER
	Reserved               unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_QUERY_DESCRIPTION
type FABRIC_PARTITION_HEALTH_QUERY_DESCRIPTION struct {
	PartitionId    syscall.GUID
	HealthPolicy   *FABRIC_APPLICATION_HEALTH_POLICY
	EventsFilter   *FABRIC_HEALTH_EVENTS_FILTER
	ReplicasFilter *FABRIC_REPLICA_HEALTH_STATES_FILTER
	Reserved       unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATISTICS_FILTER
type FABRIC_PARTITION_HEALTH_STATISTICS_FILTER struct {
	ExcludeHealthStatistics int8
	Reserved                unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_QUERY_DESCRIPTION_EX1
type FABRIC_PARTITION_HEALTH_QUERY_DESCRIPTION_EX1 struct {
	HealthStatisticsFilter *FABRIC_PARTITION_HEALTH_STATISTICS_FILTER
	Reserved               unsafe.Pointer
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

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATISTICS_FILTER
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATISTICS_FILTER struct {
	ExcludeHealthStatistics int8
	Reserved                unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_QUERY_DESCRIPTION_EX1
type FABRIC_DEPLOYED_APPLICATION_HEALTH_QUERY_DESCRIPTION_EX1 struct {
	HealthStatisticsFilter *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATISTICS_FILTER
	Reserved               unsafe.Pointer
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

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_QUERY_DESCRIPTION_EX1
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_QUERY_DESCRIPTION_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
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

// struct FABRIC_REPAIR_NODE_IMPACT
type FABRIC_REPAIR_NODE_IMPACT struct {
	NodeName    win32.PWSTR
	ImpactLevel int32
	Reserved    unsafe.Pointer
}

// struct FABRIC_REPAIR_NODE_IMPACT_LIST
type FABRIC_REPAIR_NODE_IMPACT_LIST struct {
	Count uint32
	Items *FABRIC_REPAIR_NODE_IMPACT
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

// struct FABRIC_REPAIR_TASK_HISTORY_EX1
type FABRIC_REPAIR_TASK_HISTORY_EX1 struct {
	PreparingHealthCheckStartUtcTimestamp FILETIME_
	PreparingHealthCheckEndUtcTimestamp   FILETIME_
	RestoringHealthCheckStartUtcTimestamp FILETIME_
	RestoringHealthCheckEndUtcTimestamp   FILETIME_
	Reserved                              unsafe.Pointer
}

// struct FABRIC_REPAIR_EXECUTOR_STATE
type FABRIC_REPAIR_EXECUTOR_STATE struct {
	Executor     win32.PWSTR
	ExecutorData win32.PWSTR
	Reserved     unsafe.Pointer
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

// struct FABRIC_REPAIR_TASK_EX1
type FABRIC_REPAIR_TASK_EX1 struct {
	PerformPreparingHealthCheck int8
	PerformRestoringHealthCheck int8
	PreparingHealthCheckState   int32
	RestoringHealthCheckState   int32
	Reserved                    unsafe.Pointer
}

// struct FABRIC_REPAIR_TASK_LIST
type FABRIC_REPAIR_TASK_LIST struct {
	Count uint32
	Items *FABRIC_REPAIR_TASK
}

// struct FABRIC_REPAIR_CANCEL_DESCRIPTION
type FABRIC_REPAIR_CANCEL_DESCRIPTION struct {
	Scope        *FABRIC_REPAIR_SCOPE_IDENTIFIER
	RepairTaskId win32.PWSTR
	Version      int64
	RequestAbort int8
	Reserved     unsafe.Pointer
}

// struct FABRIC_REPAIR_DELETE_DESCRIPTION
type FABRIC_REPAIR_DELETE_DESCRIPTION struct {
	Scope        *FABRIC_REPAIR_SCOPE_IDENTIFIER
	RepairTaskId win32.PWSTR
	Version      int64
	Reserved     unsafe.Pointer
}

// struct FABRIC_REPAIR_APPROVE_DESCRIPTION
type FABRIC_REPAIR_APPROVE_DESCRIPTION struct {
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

// struct FABRIC_STORE_BACKUP_INFO
type FABRIC_STORE_BACKUP_INFO struct {
	BackupFolder win32.PWSTR
	BackupOption int32
	Reserved     unsafe.Pointer
}

// struct FABRIC_STORE_BACKUP_INFO_EX1
type FABRIC_STORE_BACKUP_INFO_EX1 struct {
	BackupChainId syscall.GUID
	BackupIndex   uint32
	Reserved      unsafe.Pointer
}

// struct FABRIC_DELTA_NODES_CHECK_HEALTH_EVALUATION
type FABRIC_DELTA_NODES_CHECK_HEALTH_EVALUATION struct {
	Description                   win32.PWSTR
	AggregatedHealthState         int32
	BaselineErrorCount            uint32
	BaselineTotalCount            uint32
	TotalCount                    uint32
	MaxPercentDeltaUnhealthyNodes byte
	UnhealthyEvaluations          *FABRIC_HEALTH_EVALUATION_LIST
	Reserved                      unsafe.Pointer
}

// struct FABRIC_UPGRADE_DOMAIN_DELTA_NODES_CHECK_HEALTH_EVALUATION
type FABRIC_UPGRADE_DOMAIN_DELTA_NODES_CHECK_HEALTH_EVALUATION struct {
	Description                                win32.PWSTR
	AggregatedHealthState                      int32
	UpgradeDomainName                          win32.PWSTR
	BaselineErrorCount                         uint32
	BaselineTotalCount                         uint32
	TotalCount                                 uint32
	MaxPercentUpgradeDomainDeltaUnhealthyNodes byte
	UnhealthyEvaluations                       *FABRIC_HEALTH_EVALUATION_LIST
	Reserved                                   unsafe.Pointer
}

// struct FABRIC_PAGING_STATUS
type FABRIC_PAGING_STATUS struct {
	ContinuationToken win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_DESCRIPTION_EX1
type FABRIC_NODE_QUERY_DESCRIPTION_EX1 struct {
	ContinuationToken win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_DESCRIPTION_EX2
type FABRIC_NODE_QUERY_DESCRIPTION_EX2 struct {
	NodeStatusFilter uint32
	Reserved         unsafe.Pointer
}

// struct FABRIC_NODE_QUERY_DESCRIPTION_EX3
type FABRIC_NODE_QUERY_DESCRIPTION_EX3 struct {
	MaxResults int32
	Reserved   unsafe.Pointer
}

// struct FABRIC_SERVICE_PARTITION_QUERY_DESCRIPTION_EX1
type FABRIC_SERVICE_PARTITION_QUERY_DESCRIPTION_EX1 struct {
	ContinuationToken win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_REPLICA_QUERY_DESCRIPTION_EX2
type FABRIC_SERVICE_REPLICA_QUERY_DESCRIPTION_EX2 struct {
	ContinuationToken win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_APPLICATION_TYPE_HEALTH_POLICY_MAP_ITEM
type FABRIC_APPLICATION_TYPE_HEALTH_POLICY_MAP_ITEM struct {
	ApplicationTypeName             win32.PWSTR
	MaxPercentUnhealthyApplications byte
}

// struct FABRIC_APPLICATION_TYPE_HEALTH_POLICY_MAP
type FABRIC_APPLICATION_TYPE_HEALTH_POLICY_MAP struct {
	Count uint32
	Items *FABRIC_APPLICATION_TYPE_HEALTH_POLICY_MAP_ITEM
}

// struct FABRIC_CLUSTER_HEALTH_POLICY_EX1
type FABRIC_CLUSTER_HEALTH_POLICY_EX1 struct {
	ApplicationTypeHealthPolicyMap *FABRIC_APPLICATION_TYPE_HEALTH_POLICY_MAP
	Reserved                       unsafe.Pointer
}

// struct FABRIC_APPLICATION_TYPE_APPLICATIONS_HEALTH_EVALUATION
type FABRIC_APPLICATION_TYPE_APPLICATIONS_HEALTH_EVALUATION struct {
	Description                     win32.PWSTR
	AggregatedHealthState           int32
	ApplicationTypeName             win32.PWSTR
	UnhealthyEvaluations            *FABRIC_HEALTH_EVALUATION_LIST
	TotalCount                      uint32
	MaxPercentUnhealthyApplications byte
	Reserved                        unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATE_FILTER
type FABRIC_NODE_HEALTH_STATE_FILTER struct {
	HealthStateFilter uint32
	NodeNameFilter    win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATE_FILTER_LIST
type FABRIC_NODE_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_NODE_HEALTH_STATE_FILTER
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER struct {
	HealthStateFilter         uint32
	ServiceManifestNameFilter win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER_EX1
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER_EX1 struct {
	ServicePackageActivationIdFilter win32.PWSTR
	Reserved                         unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER_LIST
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER struct {
	HealthStateFilter             uint32
	NodeNameFilter                win32.PWSTR
	DeployedServicePackageFilters *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_FILTER_LIST
	Reserved                      unsafe.Pointer
}

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER_LIST
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER
}

// struct FABRIC_REPLICA_HEALTH_STATE_FILTER
type FABRIC_REPLICA_HEALTH_STATE_FILTER struct {
	HealthStateFilter         uint32
	ReplicaOrInstanceIdFilter int64
	Reserved                  unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATE_FILTER_LIST
type FABRIC_REPLICA_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_REPLICA_HEALTH_STATE_FILTER
}

// struct FABRIC_PARTITION_HEALTH_STATE_FILTER
type FABRIC_PARTITION_HEALTH_STATE_FILTER struct {
	HealthStateFilter uint32
	PartitionIdFilter syscall.GUID
	ReplicaFilters    *FABRIC_REPLICA_HEALTH_STATE_FILTER_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_PARTITION_HEALTH_STATE_FILTER_LIST
type FABRIC_PARTITION_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_PARTITION_HEALTH_STATE_FILTER
}

// struct FABRIC_SERVICE_HEALTH_STATE_FILTER
type FABRIC_SERVICE_HEALTH_STATE_FILTER struct {
	HealthStateFilter uint32
	ServiceNameFilter win32.PWSTR
	PartitionFilters  *FABRIC_PARTITION_HEALTH_STATE_FILTER_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_SERVICE_HEALTH_STATE_FILTER_LIST
type FABRIC_SERVICE_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_HEALTH_STATE_FILTER
}

// struct FABRIC_APPLICATION_HEALTH_STATE_FILTER
type FABRIC_APPLICATION_HEALTH_STATE_FILTER struct {
	HealthStateFilter          uint32
	ApplicationNameFilter      win32.PWSTR
	ServiceFilters             *FABRIC_SERVICE_HEALTH_STATE_FILTER_LIST
	DeployedApplicationFilters *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_FILTER_LIST
	Reserved                   unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_STATE_FILTER_EX1
type FABRIC_APPLICATION_HEALTH_STATE_FILTER_EX1 struct {
	ApplicationTypeNameFilter win32.PWSTR
	Reserved                  unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_STATE_FILTER_LIST
type FABRIC_APPLICATION_HEALTH_STATE_FILTER_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_HEALTH_STATE_FILTER
}

// struct FABRIC_CLUSTER_HEALTH_CHUNK_QUERY_DESCRIPTION
type FABRIC_CLUSTER_HEALTH_CHUNK_QUERY_DESCRIPTION struct {
	ClusterHealthPolicy        *FABRIC_CLUSTER_HEALTH_POLICY
	ApplicationHealthPolicyMap *FABRIC_APPLICATION_HEALTH_POLICY_MAP
	ApplicationFilters         *FABRIC_APPLICATION_HEALTH_STATE_FILTER_LIST
	NodeFilters                *FABRIC_NODE_HEALTH_STATE_FILTER_LIST
	Reserved                   unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATE_CHUNK
type FABRIC_NODE_HEALTH_STATE_CHUNK struct {
	NodeName    win32.PWSTR
	HealthState int32
	Reserved    unsafe.Pointer
}

// struct FABRIC_NODE_HEALTH_STATE_CHUNK_LIST
type FABRIC_NODE_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_NODE_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATE_CHUNK
type FABRIC_REPLICA_HEALTH_STATE_CHUNK struct {
	ReplicaOrInstanceId int64
	HealthState         int32
	Reserved            unsafe.Pointer
}

// struct FABRIC_REPLICA_HEALTH_STATE_CHUNK_LIST
type FABRIC_REPLICA_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_REPLICA_HEALTH_STATE_CHUNK
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

// struct FABRIC_PARTITION_HEALTH_STATE_CHUNK_LIST
type FABRIC_PARTITION_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_PARTITION_HEALTH_STATE_CHUNK
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

// struct FABRIC_SERVICE_HEALTH_STATE_CHUNK_LIST
type FABRIC_SERVICE_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_SERVICE_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK struct {
	ServiceManifestName win32.PWSTR
	HealthState         int32
	Reserved            unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK_EX1
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK_EX1 struct {
	ServicePackageActivationId win32.PWSTR
	Reserved                   unsafe.Pointer
}

// struct FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK_LIST
type FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH_STATE_CHUNK
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

// struct FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_CHUNK_LIST
type FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_DEPLOYED_APPLICATION_HEALTH_STATE_CHUNK
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

// struct FABRIC_APPLICATION_HEALTH_STATE_CHUNK_EX1
type FABRIC_APPLICATION_HEALTH_STATE_CHUNK_EX1 struct {
	ApplicationTypeName win32.PWSTR
	Reserved            unsafe.Pointer
}

// struct FABRIC_APPLICATION_HEALTH_STATE_CHUNK_LIST
type FABRIC_APPLICATION_HEALTH_STATE_CHUNK_LIST struct {
	Count      uint32
	Items      *FABRIC_APPLICATION_HEALTH_STATE_CHUNK
	TotalCount uint32
	Reserved   unsafe.Pointer
}

// struct FABRIC_CLUSTER_HEALTH_CHUNK
type FABRIC_CLUSTER_HEALTH_CHUNK struct {
	HealthState                  int32
	NodeHealthStateChunks        *FABRIC_NODE_HEALTH_STATE_CHUNK_LIST
	ApplicationHealthStateChunks *FABRIC_APPLICATION_HEALTH_STATE_CHUNK_LIST
	Reserved                     unsafe.Pointer
}

// struct FABRIC_ROLLING_UPGRADE_POLICY_DESCRIPTION_EX3
type FABRIC_ROLLING_UPGRADE_POLICY_DESCRIPTION_EX3 struct {
	ApplicationHealthPolicyMap *FABRIC_APPLICATION_HEALTH_POLICY_MAP
	Reserved                   unsafe.Pointer
}

// struct FABRIC_EVENT_CONTEXT_MAP_ITEM
type FABRIC_EVENT_CONTEXT_MAP_ITEM struct {
	Key   win32.PWSTR
	Value win32.PWSTR
}

// struct FABRIC_EVENT_CONTEXT_MAP
type FABRIC_EVENT_CONTEXT_MAP struct {
	Count uint32
	Items *FABRIC_EVENT_CONTEXT_MAP_ITEM
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

// struct FABRIC_CHAOS_PARAMETERS_EX1
type FABRIC_CHAOS_PARAMETERS_EX1 struct {
	ClusterHealthPolicy *FABRIC_CLUSTER_HEALTH_POLICY
	Reserved            unsafe.Pointer
}

// struct FABRIC_CHAOS_TARGET_FILTER
type FABRIC_CHAOS_TARGET_FILTER struct {
	NodeTypeInclusionList    *FABRIC_STRING_LIST
	ApplicationInclusionList *FABRIC_STRING_LIST
	Reserved                 unsafe.Pointer
}

// struct FABRIC_CHAOS_PARAMETERS_EX2
type FABRIC_CHAOS_PARAMETERS_EX2 struct {
	ChaosTargetFilter *FABRIC_CHAOS_TARGET_FILTER
	Reserved          unsafe.Pointer
}

// struct FABRIC_START_CHAOS_DESCRIPTION
type FABRIC_START_CHAOS_DESCRIPTION struct {
	ChaosParameters *FABRIC_CHAOS_PARAMETERS
	Reserved        unsafe.Pointer
}

// struct FABRIC_CHAOS_EVENTS_SEGMENT_FILTER
type FABRIC_CHAOS_EVENTS_SEGMENT_FILTER struct {
	StartTimeUtc FILETIME_
	EndTimeUtc   FILETIME_
	Reserved     unsafe.Pointer
}

// struct FABRIC_CHAOS_EVENTS_SEGMENT_DESCRIPTION
type FABRIC_CHAOS_EVENTS_SEGMENT_DESCRIPTION struct {
	Filter            *FABRIC_CHAOS_EVENTS_SEGMENT_FILTER
	PagingDescription *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved          unsafe.Pointer
}

// struct FABRIC_CHAOS_REPORT_FILTER
type FABRIC_CHAOS_REPORT_FILTER struct {
	StartTimeUtc FILETIME_
	EndTimeUtc   FILETIME_
	Reserved     unsafe.Pointer
}

// struct FABRIC_GET_CHAOS_REPORT_DESCRIPTION
type FABRIC_GET_CHAOS_REPORT_DESCRIPTION struct {
	Filter            *FABRIC_CHAOS_REPORT_FILTER
	ContinuationToken win32.PWSTR
	Reserved          unsafe.Pointer
}

// struct FABRIC_CHAOS_EVENT
type FABRIC_CHAOS_EVENT struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_STARTED_EVENT
type FABRIC_STARTED_EVENT struct {
	TimeStampUtc    FILETIME_
	ChaosParameters *FABRIC_CHAOS_PARAMETERS
	Reserved        unsafe.Pointer
}

// struct FABRIC_EXECUTING_FAULTS_EVENT
type FABRIC_EXECUTING_FAULTS_EVENT struct {
	TimeStampUtc FILETIME_
	Faults       *FABRIC_STRING_LIST
	Reserved     unsafe.Pointer
}

// struct FABRIC_WAITING_EVENT
type FABRIC_WAITING_EVENT struct {
	TimeStampUtc FILETIME_
	Reason       win32.PWSTR
	Reserved     unsafe.Pointer
}

// struct FABRIC_VALIDATION_FAILED_EVENT
type FABRIC_VALIDATION_FAILED_EVENT struct {
	TimeStampUtc FILETIME_
	Reason       win32.PWSTR
	Reserved     unsafe.Pointer
}

// struct FABRIC_TEST_ERROR_EVENT
type FABRIC_TEST_ERROR_EVENT struct {
	TimeStampUtc FILETIME_
	Reason       win32.PWSTR
	Reserved     unsafe.Pointer
}

// struct FABRIC_STOPPED_EVENT
type FABRIC_STOPPED_EVENT struct {
	TimeStampUtc FILETIME_
	Reason       win32.PWSTR
	Reserved     unsafe.Pointer
}

// struct FABRIC_CHAOS_EVENT_LIST
type FABRIC_CHAOS_EVENT_LIST struct {
	Count uint32
	Items *FABRIC_CHAOS_EVENT
}

// struct FABRIC_CHAOS_REPORT
type FABRIC_CHAOS_REPORT struct {
	ChaosParameters   *FABRIC_CHAOS_PARAMETERS
	Status            int32
	ContinuationToken win32.PWSTR
	History           *FABRIC_CHAOS_EVENT_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_CHAOS_EVENTS_SEGMENT
type FABRIC_CHAOS_EVENTS_SEGMENT struct {
	ContinuationToken win32.PWSTR
	History           *FABRIC_CHAOS_EVENT_LIST
	Reserved          unsafe.Pointer
}

// struct FABRIC_CHAOS_DESCRIPTION
type FABRIC_CHAOS_DESCRIPTION struct {
	ChaosParameters *FABRIC_CHAOS_PARAMETERS
	Status          int32
	ScheduleStatus  int32
	Reserved        unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_TIME_UTC
type FABRIC_CHAOS_SCHEDULE_TIME_UTC struct {
	Hour     uint32
	Minute   uint32
	Reserved unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC
type FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC struct {
	StartTime *FABRIC_CHAOS_SCHEDULE_TIME_UTC
	EndTime   *FABRIC_CHAOS_SCHEDULE_TIME_UTC
	Reserved  unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC_LIST
type FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC_LIST struct {
	Count uint32
	Items *FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC
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

// struct FABRIC_CHAOS_SCHEDULE_JOB
type FABRIC_CHAOS_SCHEDULE_JOB struct {
	ChaosParameters win32.PWSTR
	Days            *FABRIC_CHAOS_SCHEDULE_JOB_ACTIVE_DAYS
	Times           *FABRIC_CHAOS_SCHEDULE_TIME_RANGE_UTC_LIST
	Reserved        unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_JOB_LIST
type FABRIC_CHAOS_SCHEDULE_JOB_LIST struct {
	Count uint32
	Items *FABRIC_CHAOS_SCHEDULE_JOB
}

// struct FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP_ITEM
type FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP_ITEM struct {
	Name       win32.PWSTR
	Parameters *FABRIC_CHAOS_PARAMETERS
}

// struct FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP
type FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP struct {
	Count uint32
	Items *FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP_ITEM
}

// struct FABRIC_CHAOS_SCHEDULE
type FABRIC_CHAOS_SCHEDULE struct {
	StartDate          FILETIME_
	ExpiryDate         FILETIME_
	ChaosParametersMap *FABRIC_CHAOS_SCHEDULE_CHAOS_PARAMETERS_MAP
	Jobs               *FABRIC_CHAOS_SCHEDULE_JOB_LIST
	Reserved           unsafe.Pointer
}

// struct FABRIC_CHAOS_SCHEDULE_DESCRIPTION
type FABRIC_CHAOS_SCHEDULE_DESCRIPTION struct {
	Version  uint32
	Schedule *FABRIC_CHAOS_SCHEDULE
	Reserved unsafe.Pointer
}

// struct FABRIC_CHAOS_SERVICE_SCHEDULE_DESCRIPTION
type FABRIC_CHAOS_SERVICE_SCHEDULE_DESCRIPTION struct {
	ChaosScheduleDescription *FABRIC_CHAOS_SCHEDULE_DESCRIPTION
	Reserved                 unsafe.Pointer
}

// struct FABRIC_SECRET
type FABRIC_SECRET struct {
	Name        win32.PWSTR
	Version     win32.PWSTR
	Value       win32.PWSTR
	Kind        win32.PWSTR
	ContentType win32.PWSTR
}

// struct FABRIC_SECRET_LIST
type FABRIC_SECRET_LIST struct {
	Count uint32
	Items *FABRIC_SECRET
}

// struct FABRIC_SECRET_REFERENCE
type FABRIC_SECRET_REFERENCE struct {
	Name    win32.PWSTR
	Version win32.PWSTR
}

// struct FABRIC_SECRET_REFERENCE_LIST
type FABRIC_SECRET_REFERENCE_LIST struct {
	Count uint32
	Items *FABRIC_SECRET_REFERENCE
}

// struct FABRIC_CODE_PACKAGE_EVENT_DESCRIPTION
type FABRIC_CODE_PACKAGE_EVENT_DESCRIPTION struct {
	CodePackageName   win32.PWSTR
	IsSetupEntryPoint int32
	IsContainerHost   int32
	EventType         int32
	TimeStampInTicks  int64
	SequenceNumber    int64
	Properties        *FABRIC_STRING_MAP
	Reserved          unsafe.Pointer
}

// struct FABRIC_NETWORK_DESCRIPTION
type FABRIC_NETWORK_DESCRIPTION struct {
	NetworkType int32
	Value       unsafe.Pointer
}

// struct FABRIC_LOCAL_NETWORK_CONFIGURATION_DESCRIPTION
type FABRIC_LOCAL_NETWORK_CONFIGURATION_DESCRIPTION struct {
	NetworkAddressPrefix win32.PWSTR
	Reserved             unsafe.Pointer
}

// struct FABRIC_LOCAL_NETWORK_DESCRIPTION
type FABRIC_LOCAL_NETWORK_DESCRIPTION struct {
	NetworkConfiguration *FABRIC_LOCAL_NETWORK_CONFIGURATION_DESCRIPTION
	Reserved             unsafe.Pointer
}

// struct FABRIC_DELETE_NETWORK_DESCRIPTION
type FABRIC_DELETE_NETWORK_DESCRIPTION struct {
	NetworkName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_NETWORK_INFORMATION
type FABRIC_NETWORK_INFORMATION struct {
	NetworkType int32
	Value       unsafe.Pointer
}

// struct FABRIC_NETWORK_QUERY_DESCRIPTION
type FABRIC_NETWORK_QUERY_DESCRIPTION struct {
	NetworkNameFilter   win32.PWSTR
	NetworkStatusFilter uint32
	PagingDescription   *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved            unsafe.Pointer
}

// struct FABRIC_LOCAL_NETWORK_INFORMATION
type FABRIC_LOCAL_NETWORK_INFORMATION struct {
	NetworkName          win32.PWSTR
	NetworkConfiguration *FABRIC_LOCAL_NETWORK_CONFIGURATION_DESCRIPTION
	NetworkStatus        int32
	Reserved             unsafe.Pointer
}

// struct FABRIC_NETWORK_QUERY_RESULT_LIST
type FABRIC_NETWORK_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_NETWORK_INFORMATION
}

// struct FABRIC_NETWORK_APPLICATION_QUERY_DESCRIPTION
type FABRIC_NETWORK_APPLICATION_QUERY_DESCRIPTION struct {
	NetworkName           win32.PWSTR
	ApplicationNameFilter win32.PWSTR
	PagingDescription     *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved              unsafe.Pointer
}

// struct FABRIC_NETWORK_APPLICATION_QUERY_RESULT_ITEM
type FABRIC_NETWORK_APPLICATION_QUERY_RESULT_ITEM struct {
	ApplicationName win32.PWSTR
	Reserved        unsafe.Pointer
}

// struct FABRIC_NETWORK_APPLICATION_QUERY_RESULT_LIST
type FABRIC_NETWORK_APPLICATION_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_NETWORK_APPLICATION_QUERY_RESULT_ITEM
}

// struct FABRIC_NETWORK_NODE_QUERY_DESCRIPTION
type FABRIC_NETWORK_NODE_QUERY_DESCRIPTION struct {
	NetworkName       win32.PWSTR
	NodeNameFilter    win32.PWSTR
	PagingDescription *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved          unsafe.Pointer
}

// struct FABRIC_NETWORK_NODE_QUERY_RESULT_ITEM
type FABRIC_NETWORK_NODE_QUERY_RESULT_ITEM struct {
	NodeName win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_NETWORK_NODE_QUERY_RESULT_LIST
type FABRIC_NETWORK_NODE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_NETWORK_NODE_QUERY_RESULT_ITEM
}

// struct FABRIC_APPLICATION_NETWORK_QUERY_DESCRIPTION
type FABRIC_APPLICATION_NETWORK_QUERY_DESCRIPTION struct {
	ApplicationName   win32.PWSTR
	PagingDescription *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved          unsafe.Pointer
}

// struct FABRIC_APPLICATION_NETWORK_QUERY_RESULT_ITEM
type FABRIC_APPLICATION_NETWORK_QUERY_RESULT_ITEM struct {
	NetworkName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_APPLICATION_NETWORK_QUERY_RESULT_LIST
type FABRIC_APPLICATION_NETWORK_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_APPLICATION_NETWORK_QUERY_RESULT_ITEM
}

// struct FABRIC_DEPLOYED_NETWORK_QUERY_DESCRIPTION
type FABRIC_DEPLOYED_NETWORK_QUERY_DESCRIPTION struct {
	NodeName          win32.PWSTR
	PagingDescription *FABRIC_QUERY_PAGING_DESCRIPTION
	Reserved          unsafe.Pointer
}

// struct FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_ITEM
type FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_ITEM struct {
	NetworkName win32.PWSTR
	Reserved    unsafe.Pointer
}

// struct FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_ITEM
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

// struct FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_LIST
type FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_LIST struct {
	Count uint32
	Items *FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_ITEM
}
