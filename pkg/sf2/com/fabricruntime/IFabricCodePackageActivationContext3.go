package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 6EFEE900-F491-4B03-BC5B-3A70DE103593
var IID_IFabricCodePackageActivationContext3 = syscall.GUID{Data1: 0x6EFEE900, Data2: 0xF491, Data3: 0x4B03,
	Data4: [8]byte{0xBC, 0x5B, 0x3A, 0x70, 0xDE, 0x10, 0x35, 0x93}}

type IFabricCodePackageActivationContext3 struct {
	IFabricCodePackageActivationContext2
}

func NewIFabricCodePackageActivationContext3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricCodePackageActivationContext3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricCodePackageActivationContext3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricCodePackageActivationContext3) IID() *syscall.GUID {
	return &IID_IFabricCodePackageActivationContext3
}

func (this *IFabricCodePackageActivationContext3) ReportApplicationHealth(healthInfo *FABRIC_HEALTH_INFORMATION) com.Error {
	addr := (*this.LpVtbl)[30]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext3) ReportDeployedApplicationHealth(healthInfo *FABRIC_HEALTH_INFORMATION) com.Error {
	addr := (*this.LpVtbl)[31]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext3) ReportDeployedServicePackageHealth(healthInfo *FABRIC_HEALTH_INFORMATION) com.Error {
	addr := (*this.LpVtbl)[32]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)))
	return com.Error(ret)
}
