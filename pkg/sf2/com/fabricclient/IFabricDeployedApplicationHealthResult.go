package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 4DF50BF4-7C28-4210-94F7-50625DF6C942
var IID_IFabricDeployedApplicationHealthResult = syscall.GUID{Data1: 0x4DF50BF4, Data2: 0x7C28, Data3: 0x4210,
	Data4: [8]byte{0x94, 0xF7, 0x50, 0x62, 0x5D, 0xF6, 0xC9, 0x42}}

type IFabricDeployedApplicationHealthResult struct {
	win32.IUnknown
}

func NewIFabricDeployedApplicationHealthResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricDeployedApplicationHealthResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricDeployedApplicationHealthResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricDeployedApplicationHealthResult) IID() *syscall.GUID {
	return &IID_IFabricDeployedApplicationHealthResult
}

func (this *IFabricDeployedApplicationHealthResult) Get_DeployedApplicationHealth() *FABRIC_DEPLOYED_APPLICATION_HEALTH {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_APPLICATION_HEALTH)(unsafe.Pointer(ret))
}
