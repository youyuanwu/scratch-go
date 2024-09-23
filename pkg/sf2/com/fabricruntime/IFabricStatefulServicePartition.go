package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 5BECCC37-8655-4F20-BD43-F50691D7CD16
var IID_IFabricStatefulServicePartition = syscall.GUID{Data1: 0x5BECCC37, Data2: 0x8655, Data3: 0x4F20,
	Data4: [8]byte{0xBD, 0x43, 0xF5, 0x06, 0x91, 0xD7, 0xCD, 0x16}}

type IFabricStatefulServicePartition struct {
	win32.IUnknown
}

func NewIFabricStatefulServicePartition(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatefulServicePartition {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatefulServicePartition)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatefulServicePartition) IID() *syscall.GUID {
	return &IID_IFabricStatefulServicePartition
}

func (this *IFabricStatefulServicePartition) GetPartitionInfo(bufferedValue **FABRIC_SERVICE_PARTITION_INFORMATION) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bufferedValue)))
	return com.Error(ret)
}

func (this *IFabricStatefulServicePartition) GetReadStatus(readStatus *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(readStatus)))
	return com.Error(ret)
}

func (this *IFabricStatefulServicePartition) GetWriteStatus(writeStatus *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(writeStatus)))
	return com.Error(ret)
}

func (this *IFabricStatefulServicePartition) CreateReplicator(stateProvider *IFabricStateProvider, replicatorSettings *FABRIC_REPLICATOR_SETTINGS, replicator **IFabricReplicator, stateReplicator **IFabricStateReplicator) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stateProvider)), uintptr(unsafe.Pointer(replicatorSettings)), uintptr(unsafe.Pointer(replicator)), uintptr(unsafe.Pointer(stateReplicator)))
	com.AddToScope(replicator)
	com.AddToScope(stateReplicator)
	return com.Error(ret)
}

func (this *IFabricStatefulServicePartition) ReportLoad(metricCount uint32, metrics *FABRIC_LOAD_METRIC) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(metricCount), uintptr(unsafe.Pointer(metrics)))
	return com.Error(ret)
}

func (this *IFabricStatefulServicePartition) ReportFault(faultType int32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(faultType))
	return com.Error(ret)
}
