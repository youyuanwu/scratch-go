package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 3BA780E9-58EB-478D-BC89-42C89E19D083
var IID_IFabricGetNetworkNodeListResult = syscall.GUID{Data1: 0x3BA780E9, Data2: 0x58EB, Data3: 0x478D,
	Data4: [8]byte{0xBC, 0x89, 0x42, 0xC8, 0x9E, 0x19, 0xD0, 0x83}}

type IFabricGetNetworkNodeListResult struct {
	win32.IUnknown
}

func NewIFabricGetNetworkNodeListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetNetworkNodeListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetNetworkNodeListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetNetworkNodeListResult) IID() *syscall.GUID {
	return &IID_IFabricGetNetworkNodeListResult
}

func (this *IFabricGetNetworkNodeListResult) Get_NetworkNodeList() *FABRIC_NETWORK_NODE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NETWORK_NODE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricGetNetworkNodeListResult) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
