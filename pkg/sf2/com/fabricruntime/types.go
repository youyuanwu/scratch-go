package fabricruntime

import (
	"github.com/zzl/go-win32api/v2/win32"
	"unsafe"
)

// struct __MIDL___MIDL_itf_FabricRuntime_0004_0001_0001
type MIDL___MIDL_itf_FabricRuntime_0004_0001_0001__ struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// struct FABRIC_SERVICE_PARTITION_INFORMATION
type FABRIC_SERVICE_PARTITION_INFORMATION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_LOAD_METRIC
type FABRIC_LOAD_METRIC struct {
	Name     win32.PWSTR
	Value    uint32
	Reserved unsafe.Pointer
}

// struct FABRIC_EPOCH
type FABRIC_EPOCH struct {
	DataLossNumber      int64
	ConfigurationNumber int64
	Reserved            unsafe.Pointer
}

// struct FABRIC_OPERATION_DATA_BUFFER
type FABRIC_OPERATION_DATA_BUFFER struct {
	BufferSize uint32
	Buffer     *byte
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

// struct FABRIC_SECURITY_CREDENTIALS
type FABRIC_SECURITY_CREDENTIALS struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_OPERATION_METADATA
type FABRIC_OPERATION_METADATA struct {
	Type           int32
	SequenceNumber int64
	AtomicGroupId  int64
	Reserved       unsafe.Pointer
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

// struct FABRIC_HEALTH_REPORT_SEND_OPTIONS
type FABRIC_HEALTH_REPORT_SEND_OPTIONS struct {
	Immediate int8
	Reserved  unsafe.Pointer
}

// struct FABRIC_REPLICA_SET_CONFIGURATION
type FABRIC_REPLICA_SET_CONFIGURATION struct {
	ReplicaCount uint32
	Replicas     *FABRIC_REPLICA_INFORMATION
	WriteQuorum  uint32
	Reserved     unsafe.Pointer
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

// struct FABRIC_SERVICE_TYPE_DESCRIPTION_LIST
type FABRIC_SERVICE_TYPE_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_TYPE_DESCRIPTION
}

// struct FABRIC_SERVICE_TYPE_DESCRIPTION
type FABRIC_SERVICE_TYPE_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION_LIST
type FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION
}

// struct FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION
type FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION struct {
	Description        *FABRIC_SERVICE_TYPE_DESCRIPTION
	Members            *FABRIC_SERVICE_GROUP_TYPE_MEMBER_DESCRIPTION_LIST
	UseImplicitFactory int8
	Reserved           unsafe.Pointer
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

// struct FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION
type FABRIC_SERVICE_LOAD_METRIC_DESCRIPTION struct {
	Name                 win32.PWSTR
	Weight               int32
	PrimaryDefaultLoad   uint32
	SecondaryDefaultLoad uint32
	Reserved             unsafe.Pointer
}

// struct FABRIC_APPLICATION_PRINCIPALS_DESCRIPTION
type FABRIC_APPLICATION_PRINCIPALS_DESCRIPTION struct {
	Users    *FABRIC_SECURITY_USER_DESCRIPTION_LIST
	Groups   *FABRIC_SECURITY_GROUP_DESCRIPTION_LIST
	Reserved unsafe.Pointer
}

// struct FABRIC_SECURITY_USER_DESCRIPTION_LIST
type FABRIC_SECURITY_USER_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SECURITY_USER_DESCRIPTION
}

// struct FABRIC_SECURITY_USER_DESCRIPTION
type FABRIC_SECURITY_USER_DESCRIPTION struct {
	Name                    win32.PWSTR
	Sid                     win32.PWSTR
	ParentSystemGroups      *FABRIC_STRING_LIST
	ParentApplicationGroups *FABRIC_STRING_LIST
	Reserved                unsafe.Pointer
}

// struct FABRIC_STRING_LIST
type FABRIC_STRING_LIST struct {
	Count uint32
	Items *win32.PWSTR
}

// struct FABRIC_SECURITY_GROUP_DESCRIPTION_LIST
type FABRIC_SECURITY_GROUP_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_SECURITY_GROUP_DESCRIPTION
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

// struct FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_LIST
type FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_LIST struct {
	Count uint32
	Items *FABRIC_ENDPOINT_RESOURCE_DESCRIPTION
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

// struct FABRIC_EXEHOST_ENTRY_POINT_DESCRIPTION
type FABRIC_EXEHOST_ENTRY_POINT_DESCRIPTION struct {
	Program       win32.PWSTR
	Arguments     win32.PWSTR
	WorkingFolder int32
	Reserved      unsafe.Pointer
}

// struct FABRIC_CODE_PACKAGE_ENTRY_POINT_DESCRIPTION
type FABRIC_CODE_PACKAGE_ENTRY_POINT_DESCRIPTION struct {
	Kind  int32
	Value unsafe.Pointer
}

// struct FABRIC_CONFIGURATION_PACKAGE_DESCRIPTION
type FABRIC_CONFIGURATION_PACKAGE_DESCRIPTION struct {
	Name                   win32.PWSTR
	Version                win32.PWSTR
	ServiceManifestName    win32.PWSTR
	ServiceManifestVersion win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_CONFIGURATION_SETTINGS
type FABRIC_CONFIGURATION_SETTINGS struct {
	Sections *FABRIC_CONFIGURATION_SECTION_LIST
	Reserved unsafe.Pointer
}

// struct FABRIC_CONFIGURATION_SECTION_LIST
type FABRIC_CONFIGURATION_SECTION_LIST struct {
	Count uint32
	Items *FABRIC_CONFIGURATION_SECTION
}

// struct FABRIC_CONFIGURATION_SECTION
type FABRIC_CONFIGURATION_SECTION struct {
	Name       win32.PWSTR
	Parameters *FABRIC_CONFIGURATION_PARAMETER_LIST
	Reserved   unsafe.Pointer
}

// struct FABRIC_CONFIGURATION_PARAMETER_LIST
type FABRIC_CONFIGURATION_PARAMETER_LIST struct {
	Count uint32
	Items *FABRIC_CONFIGURATION_PARAMETER
}

// struct FABRIC_CONFIGURATION_PARAMETER
type FABRIC_CONFIGURATION_PARAMETER struct {
	Name         win32.PWSTR
	Value        win32.PWSTR
	MustOverride int8
	IsEncrypted  int8
	Reserved     unsafe.Pointer
}

// struct FABRIC_DATA_PACKAGE_DESCRIPTION
type FABRIC_DATA_PACKAGE_DESCRIPTION struct {
	Name                   win32.PWSTR
	Version                win32.PWSTR
	ServiceManifestName    win32.PWSTR
	ServiceManifestVersion win32.PWSTR
	Reserved               unsafe.Pointer
}

// struct FABRIC_RUNAS_POLICY_DESCRIPTION
type FABRIC_RUNAS_POLICY_DESCRIPTION struct {
	UserName win32.PWSTR
	Reserved unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_ITEM
type FABRIC_KEY_VALUE_STORE_ITEM struct {
	Metadata *FABRIC_KEY_VALUE_STORE_ITEM_METADATA
	Value    *byte
	Reserved unsafe.Pointer
}

// struct FABRIC_KEY_VALUE_STORE_ITEM_METADATA
type FABRIC_KEY_VALUE_STORE_ITEM_METADATA struct {
	Key              win32.PWSTR
	ValueSizeInBytes int32
	SequenceNumber   int64
	LastModifiedUtc  FILETIME_
	Reserved         unsafe.Pointer
}

// struct _FILETIME
type FILETIME_ struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

// struct FABRIC_KEY_VALUE_STORE_TRANSACTION_SETTINGS
type FABRIC_KEY_VALUE_STORE_TRANSACTION_SETTINGS struct {
	SerializationBlockSize uint32
	Reserved               unsafe.Pointer
}

// struct FABRIC_STORE_BACKUP_INFO
type FABRIC_STORE_BACKUP_INFO struct {
	BackupFolder win32.PWSTR
	BackupOption int32
	Reserved     unsafe.Pointer
}

// struct FABRIC_NODE_CONTEXT
type FABRIC_NODE_CONTEXT struct {
	NodeName        win32.PWSTR
	NodeType        win32.PWSTR
	IPAddressOrFQDN win32.PWSTR
	NodeInstanceId  uint64
	NodeId          FABRIC_NODE_ID
	Reserved        unsafe.Pointer
}

// struct FABRIC_NODE_ID
type FABRIC_NODE_ID struct {
	Low      uint64
	High     uint64
	Reserved unsafe.Pointer
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

// struct FABRIC_STRING_MAP
type FABRIC_STRING_MAP struct {
	Count uint32
	Items *FABRIC_APPLICATION_PARAMETER
}

// struct FABRIC_APPLICATION_PARAMETER
type FABRIC_APPLICATION_PARAMETER struct {
	Name     win32.PWSTR
	Value    win32.PWSTR
	Reserved unsafe.Pointer
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

// struct FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS
type FABRIC_KEY_VALUE_STORE_REPLICA_SETTINGS struct {
	TransactionDrainTimeoutInSeconds uint32
	SecondaryNotificationMode        int32
	Reserved                         unsafe.Pointer
}
