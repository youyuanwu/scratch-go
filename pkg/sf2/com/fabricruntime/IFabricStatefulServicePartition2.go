package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// DF27B476-FA25-459F-A7D3-87D3EEC9C73C
var IID_IFabricStatefulServicePartition2 = syscall.GUID{Data1: 0xDF27B476, Data2: 0xFA25, Data3: 0x459F,
	Data4: [8]byte{0xA7, 0xD3, 0x87, 0xD3, 0xEE, 0xC9, 0xC7, 0x3C}}

type IFabricStatefulServicePartition2 struct {
	IFabricStatefulServicePartition1
}

func NewIFabricStatefulServicePartition2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatefulServicePartition2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatefulServicePartition2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatefulServicePartition2) IID() *syscall.GUID {
	return &IID_IFabricStatefulServicePartition2
}

func (this *IFabricStatefulServicePartition2) ReportReplicaHealth(healthInfo *FABRIC_HEALTH_INFORMATION) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)))
	return com.Error(ret)
}

func (this *IFabricStatefulServicePartition2) ReportPartitionHealth(healthInfo *FABRIC_HEALTH_INFORMATION) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)))
	return com.Error(ret)
}
