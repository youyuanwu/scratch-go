package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// DBA68C7A-3F77-49BB-B611-FF94DF062B8D
var IID_IFabricGetDeployedServiceTypeListResult = syscall.GUID{Data1: 0xDBA68C7A, Data2: 0x3F77, Data3: 0x49BB,
	Data4: [8]byte{0xB6, 0x11, 0xFF, 0x94, 0xDF, 0x06, 0x2B, 0x8D}}

type IFabricGetDeployedServiceTypeListResult struct {
	win32.IUnknown
}

func NewIFabricGetDeployedServiceTypeListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetDeployedServiceTypeListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetDeployedServiceTypeListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetDeployedServiceTypeListResult) IID() *syscall.GUID {
	return &IID_IFabricGetDeployedServiceTypeListResult
}

func (this *IFabricGetDeployedServiceTypeListResult) Get_DeployedServiceTypeList() *FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_SERVICE_TYPE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
