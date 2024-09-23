package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 65851388-0421-4107-977B-39F4E15440D4
var IID_IFabricGetDeployedServicePackageListResult = syscall.GUID{Data1: 0x65851388, Data2: 0x0421, Data3: 0x4107,
	Data4: [8]byte{0x97, 0x7B, 0x39, 0xF4, 0xE1, 0x54, 0x40, 0xD4}}

type IFabricGetDeployedServicePackageListResult struct {
	win32.IUnknown
}

func NewIFabricGetDeployedServicePackageListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetDeployedServicePackageListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetDeployedServicePackageListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetDeployedServicePackageListResult) IID() *syscall.GUID {
	return &IID_IFabricGetDeployedServicePackageListResult
}

func (this *IFabricGetDeployedServicePackageListResult) Get_DeployedServicePackageList() *FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_SERVICE_PACKAGE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
