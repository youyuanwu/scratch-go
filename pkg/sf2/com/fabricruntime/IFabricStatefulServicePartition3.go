package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 51F1269D-B061-4C1C-96CF-6508CECE813B
var IID_IFabricStatefulServicePartition3 = syscall.GUID{Data1: 0x51F1269D, Data2: 0xB061, Data3: 0x4C1C,
	Data4: [8]byte{0x96, 0xCF, 0x65, 0x08, 0xCE, 0xCE, 0x81, 0x3B}}

type IFabricStatefulServicePartition3 struct {
	IFabricStatefulServicePartition2
}

func NewIFabricStatefulServicePartition3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatefulServicePartition3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatefulServicePartition3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatefulServicePartition3) IID() *syscall.GUID {
	return &IID_IFabricStatefulServicePartition3
}

func (this *IFabricStatefulServicePartition3) ReportReplicaHealth2(healthInfo *FABRIC_HEALTH_INFORMATION, sendOptions *FABRIC_HEALTH_REPORT_SEND_OPTIONS) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)), uintptr(unsafe.Pointer(sendOptions)))
	return com.Error(ret)
}

func (this *IFabricStatefulServicePartition3) ReportPartitionHealth2(healthInfo *FABRIC_HEALTH_INFORMATION, sendOptions *FABRIC_HEALTH_REPORT_SEND_OPTIONS) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)), uintptr(unsafe.Pointer(sendOptions)))
	return com.Error(ret)
}
