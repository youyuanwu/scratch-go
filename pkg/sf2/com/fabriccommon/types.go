package fabriccommon

import (
	"unsafe"
)

// struct __MIDL___MIDL_itf_FabricCommon_0004_0001_0001
type MIDL___MIDL_itf_FabricCommon_0004_0001_0001__ struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// struct FABRIC_REPLICATOR_STATUS_QUERY_RESULT
type FABRIC_REPLICATOR_STATUS_QUERY_RESULT struct {
	Role  int32
	Value unsafe.Pointer
}
