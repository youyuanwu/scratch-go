package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 4F9D0390-AA08-4DEE-BA49-62891EB47C37
var IID_IFabricGetApplicationNetworkListResult = syscall.GUID{Data1: 0x4F9D0390, Data2: 0xAA08, Data3: 0x4DEE,
	Data4: [8]byte{0xBA, 0x49, 0x62, 0x89, 0x1E, 0xB4, 0x7C, 0x37}}

type IFabricGetApplicationNetworkListResult struct {
	win32.IUnknown
}

func NewIFabricGetApplicationNetworkListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetApplicationNetworkListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetApplicationNetworkListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetApplicationNetworkListResult) IID() *syscall.GUID {
	return &IID_IFabricGetApplicationNetworkListResult
}

func (this *IFabricGetApplicationNetworkListResult) Get_ApplicationNetworkList() *FABRIC_APPLICATION_NETWORK_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_APPLICATION_NETWORK_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricGetApplicationNetworkListResult) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
