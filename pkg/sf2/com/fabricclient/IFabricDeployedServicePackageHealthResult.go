package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 40991CE0-CDBB-44E9-9CDC-B14A5D5EA4C1
var IID_IFabricDeployedServicePackageHealthResult = syscall.GUID{Data1: 0x40991CE0, Data2: 0xCDBB, Data3: 0x44E9,
	Data4: [8]byte{0x9C, 0xDC, 0xB1, 0x4A, 0x5D, 0x5E, 0xA4, 0xC1}}

type IFabricDeployedServicePackageHealthResult struct {
	win32.IUnknown
}

func NewIFabricDeployedServicePackageHealthResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricDeployedServicePackageHealthResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricDeployedServicePackageHealthResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricDeployedServicePackageHealthResult) IID() *syscall.GUID {
	return &IID_IFabricDeployedServicePackageHealthResult
}

func (this *IFabricDeployedServicePackageHealthResult) Get_DeployedServicePackageHealth() *FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_SERVICE_PACKAGE_HEALTH)(unsafe.Pointer(ret))
}
