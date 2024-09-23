package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 5722B789-3936-4C33-9F7A-342967457612
var IID_IFabricGetDeployedApplicationListResult = syscall.GUID{Data1: 0x5722B789, Data2: 0x3936, Data3: 0x4C33,
	Data4: [8]byte{0x9F, 0x7A, 0x34, 0x29, 0x67, 0x45, 0x76, 0x12}}

type IFabricGetDeployedApplicationListResult struct {
	win32.IUnknown
}

func NewIFabricGetDeployedApplicationListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetDeployedApplicationListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetDeployedApplicationListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetDeployedApplicationListResult) IID() *syscall.GUID {
	return &IID_IFabricGetDeployedApplicationListResult
}

func (this *IFabricGetDeployedApplicationListResult) Get_DeployedApplicationList() *FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
