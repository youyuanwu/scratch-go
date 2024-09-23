package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// CC53AF91-74CD-11DF-AC3E-0024811E3892
var IID_IFabricStatelessServicePartition = syscall.GUID{Data1: 0xCC53AF91, Data2: 0x74CD, Data3: 0x11DF,
	Data4: [8]byte{0xAC, 0x3E, 0x00, 0x24, 0x81, 0x1E, 0x38, 0x92}}

type IFabricStatelessServicePartition struct {
	win32.IUnknown
}

func NewIFabricStatelessServicePartition(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatelessServicePartition {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatelessServicePartition)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatelessServicePartition) IID() *syscall.GUID {
	return &IID_IFabricStatelessServicePartition
}

func (this *IFabricStatelessServicePartition) GetPartitionInfo(bufferedValue **FABRIC_SERVICE_PARTITION_INFORMATION) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bufferedValue)))
	return com.Error(ret)
}

func (this *IFabricStatelessServicePartition) ReportLoad(metricCount uint32, metrics *FABRIC_LOAD_METRIC) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(metricCount), uintptr(unsafe.Pointer(metrics)))
	return com.Error(ret)
}

func (this *IFabricStatelessServicePartition) ReportFault(faultType int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(faultType))
	return com.Error(ret)
}
