package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F2FA2000-70A7-4ED5-9D3E-0B7DECA2433F
var IID_IFabricStatelessServicePartition3 = syscall.GUID{Data1: 0xF2FA2000, Data2: 0x70A7, Data3: 0x4ED5,
	Data4: [8]byte{0x9D, 0x3E, 0x0B, 0x7D, 0xEC, 0xA2, 0x43, 0x3F}}

type IFabricStatelessServicePartition3 struct {
	IFabricStatelessServicePartition2
}

func NewIFabricStatelessServicePartition3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatelessServicePartition3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatelessServicePartition3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatelessServicePartition3) IID() *syscall.GUID {
	return &IID_IFabricStatelessServicePartition3
}

func (this *IFabricStatelessServicePartition3) ReportInstanceHealth2(healthInfo *FABRIC_HEALTH_INFORMATION, sendOptions *FABRIC_HEALTH_REPORT_SEND_OPTIONS) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)), uintptr(unsafe.Pointer(sendOptions)))
	return com.Error(ret)
}

func (this *IFabricStatelessServicePartition3) ReportPartitionHealth2(healthInfo *FABRIC_HEALTH_INFORMATION, sendOptions *FABRIC_HEALTH_REPORT_SEND_OPTIONS) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)), uintptr(unsafe.Pointer(sendOptions)))
	return com.Error(ret)
}
