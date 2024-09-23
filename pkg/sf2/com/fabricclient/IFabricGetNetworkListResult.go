package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// BD777A0F-2020-40BB-8F23-8756649CCE47
var IID_IFabricGetNetworkListResult = syscall.GUID{Data1: 0xBD777A0F, Data2: 0x2020, Data3: 0x40BB,
	Data4: [8]byte{0x8F, 0x23, 0x87, 0x56, 0x64, 0x9C, 0xCE, 0x47}}

type IFabricGetNetworkListResult struct {
	win32.IUnknown
}

func NewIFabricGetNetworkListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetNetworkListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetNetworkListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetNetworkListResult) IID() *syscall.GUID {
	return &IID_IFabricGetNetworkListResult
}

func (this *IFabricGetNetworkListResult) Get_NetworkList() *FABRIC_NETWORK_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NETWORK_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricGetNetworkListResult) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
