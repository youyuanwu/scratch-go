package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// D042BDB6-4364-4818-B395-0E6B1A22CB11
var IID_IFabricGetProvisionedCodeVersionListResult = syscall.GUID{Data1: 0xD042BDB6, Data2: 0x4364, Data3: 0x4818,
	Data4: [8]byte{0xB3, 0x95, 0x0E, 0x6B, 0x1A, 0x22, 0xCB, 0x11}}

type IFabricGetProvisionedCodeVersionListResult struct {
	win32.IUnknown
}

func NewIFabricGetProvisionedCodeVersionListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetProvisionedCodeVersionListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetProvisionedCodeVersionListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetProvisionedCodeVersionListResult) IID() *syscall.GUID {
	return &IID_IFabricGetProvisionedCodeVersionListResult
}

func (this *IFabricGetProvisionedCodeVersionListResult) Get_ProvisionedCodeVersionList() *FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PROVISIONED_CODE_VERSION_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
