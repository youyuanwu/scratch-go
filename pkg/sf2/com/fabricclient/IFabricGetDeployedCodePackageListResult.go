package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 3F390652-C0DC-4919-8A7F-8AE1E827DE0C
var IID_IFabricGetDeployedCodePackageListResult = syscall.GUID{Data1: 0x3F390652, Data2: 0xC0DC, Data3: 0x4919,
	Data4: [8]byte{0x8A, 0x7F, 0x8A, 0xE1, 0xE8, 0x27, 0xDE, 0x0C}}

type IFabricGetDeployedCodePackageListResult struct {
	win32.IUnknown
}

func NewIFabricGetDeployedCodePackageListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetDeployedCodePackageListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetDeployedCodePackageListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetDeployedCodePackageListResult) IID() *syscall.GUID {
	return &IID_IFabricGetDeployedCodePackageListResult
}

func (this *IFabricGetDeployedCodePackageListResult) Get_DeployedCodePackageList() *FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_CODE_PACKAGE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
