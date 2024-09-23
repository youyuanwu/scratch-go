package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// EBD76F6F-508E-43EA-9CA2-A98EA2C0E846
var IID_IFabricGetDeployedApplicationPagedListResult = syscall.GUID{Data1: 0xEBD76F6F, Data2: 0x508E, Data3: 0x43EA,
	Data4: [8]byte{0x9C, 0xA2, 0xA9, 0x8E, 0xA2, 0xC0, 0xE8, 0x46}}

type IFabricGetDeployedApplicationPagedListResult struct {
	win32.IUnknown
}

func NewIFabricGetDeployedApplicationPagedListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetDeployedApplicationPagedListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetDeployedApplicationPagedListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetDeployedApplicationPagedListResult) IID() *syscall.GUID {
	return &IID_IFabricGetDeployedApplicationPagedListResult
}

func (this *IFabricGetDeployedApplicationPagedListResult) Get_DeployedApplicationPagedList() *FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_APPLICATION_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricGetDeployedApplicationPagedListResult) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
