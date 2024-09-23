package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 6586D264-A96E-4F46-9388-189DE5D61D6D
var IID_IFabricGetDeployedNetworkCodePackageListResult = syscall.GUID{Data1: 0x6586D264, Data2: 0xA96E, Data3: 0x4F46,
	Data4: [8]byte{0x93, 0x88, 0x18, 0x9D, 0xE5, 0xD6, 0x1D, 0x6D}}

type IFabricGetDeployedNetworkCodePackageListResult struct {
	win32.IUnknown
}

func NewIFabricGetDeployedNetworkCodePackageListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetDeployedNetworkCodePackageListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetDeployedNetworkCodePackageListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetDeployedNetworkCodePackageListResult) IID() *syscall.GUID {
	return &IID_IFabricGetDeployedNetworkCodePackageListResult
}

func (this *IFabricGetDeployedNetworkCodePackageListResult) Get_DeployedNetworkCodePackageList() *FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_NETWORK_CODE_PACKAGE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricGetDeployedNetworkCodePackageListResult) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
