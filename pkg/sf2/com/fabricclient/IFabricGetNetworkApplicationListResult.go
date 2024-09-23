package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// AD1F51FF-E244-498E-9F72-609B01124B84
var IID_IFabricGetNetworkApplicationListResult = syscall.GUID{Data1: 0xAD1F51FF, Data2: 0xE244, Data3: 0x498E,
	Data4: [8]byte{0x9F, 0x72, 0x60, 0x9B, 0x01, 0x12, 0x4B, 0x84}}

type IFabricGetNetworkApplicationListResult struct {
	win32.IUnknown
}

func NewIFabricGetNetworkApplicationListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetNetworkApplicationListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetNetworkApplicationListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetNetworkApplicationListResult) IID() *syscall.GUID {
	return &IID_IFabricGetNetworkApplicationListResult
}

func (this *IFabricGetNetworkApplicationListResult) Get_NetworkApplicationList() *FABRIC_NETWORK_APPLICATION_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NETWORK_APPLICATION_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricGetNetworkApplicationListResult) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
