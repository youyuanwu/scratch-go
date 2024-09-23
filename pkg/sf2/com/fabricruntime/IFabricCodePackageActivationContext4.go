package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 99EFEBB6-A7B4-4D45-B45E-F191A66EEF03
var IID_IFabricCodePackageActivationContext4 = syscall.GUID{Data1: 0x99EFEBB6, Data2: 0xA7B4, Data3: 0x4D45,
	Data4: [8]byte{0xB4, 0x5E, 0xF1, 0x91, 0xA6, 0x6E, 0xEF, 0x03}}

type IFabricCodePackageActivationContext4 struct {
	IFabricCodePackageActivationContext3
}

func NewIFabricCodePackageActivationContext4(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricCodePackageActivationContext4 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricCodePackageActivationContext4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricCodePackageActivationContext4) IID() *syscall.GUID {
	return &IID_IFabricCodePackageActivationContext4
}

func (this *IFabricCodePackageActivationContext4) ReportApplicationHealth2(healthInfo *FABRIC_HEALTH_INFORMATION, sendOptions *FABRIC_HEALTH_REPORT_SEND_OPTIONS) com.Error {
	addr := (*this.LpVtbl)[33]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)), uintptr(unsafe.Pointer(sendOptions)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext4) ReportDeployedApplicationHealth2(healthInfo *FABRIC_HEALTH_INFORMATION, sendOptions *FABRIC_HEALTH_REPORT_SEND_OPTIONS) com.Error {
	addr := (*this.LpVtbl)[34]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)), uintptr(unsafe.Pointer(sendOptions)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext4) ReportDeployedServicePackageHealth2(healthInfo *FABRIC_HEALTH_INFORMATION, sendOptions *FABRIC_HEALTH_REPORT_SEND_OPTIONS) com.Error {
	addr := (*this.LpVtbl)[35]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthInfo)), uintptr(unsafe.Pointer(sendOptions)))
	return com.Error(ret)
}
