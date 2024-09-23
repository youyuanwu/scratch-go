package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 9FF35B6C-9D97-4312-93AD-7F34CBDB4CA4
var IID_IFabricStatelessServicePartition2 = syscall.GUID{Data1: 0x9FF35B6C, Data2: 0x9D97, Data3: 0x4312,
	Data4: [8]byte{0x93, 0xAD, 0x7F, 0x34, 0xCB, 0xDB, 0x4C, 0xA4}}

type IFabricStatelessServicePartition2 struct {
	IFabricStatelessServicePartition1
}

func NewIFabricStatelessServicePartition2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatelessServicePartition2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatelessServicePartition2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatelessServicePartition2) IID() *syscall.GUID {
	return &IID_IFabricStatelessServicePartition2
}

func (this *IFabricStatelessServicePartition2) ReportInstanceHealth(healthInfo *FABRIC_HEALTH_INFORMATION) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)))
	return com.Error(ret)
}

func (this *IFabricStatelessServicePartition2) ReportPartitionHealth(healthInfo *FABRIC_HEALTH_INFORMATION) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)))
	return com.Error(ret)
}
