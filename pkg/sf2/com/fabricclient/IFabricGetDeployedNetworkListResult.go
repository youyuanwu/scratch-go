package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 347F5D8C-1ABD-48E1-A7D1-9083556DAFD3
var IID_IFabricGetDeployedNetworkListResult = syscall.GUID{Data1: 0x347F5D8C, Data2: 0x1ABD, Data3: 0x48E1,
	Data4: [8]byte{0xA7, 0xD1, 0x90, 0x83, 0x55, 0x6D, 0xAF, 0xD3}}

type IFabricGetDeployedNetworkListResult struct {
	win32.IUnknown
}

func NewIFabricGetDeployedNetworkListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetDeployedNetworkListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetDeployedNetworkListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetDeployedNetworkListResult) IID() *syscall.GUID {
	return &IID_IFabricGetDeployedNetworkListResult
}

func (this *IFabricGetDeployedNetworkListResult) Get_DeployedNetworkList() *FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_NETWORK_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricGetDeployedNetworkListResult) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
