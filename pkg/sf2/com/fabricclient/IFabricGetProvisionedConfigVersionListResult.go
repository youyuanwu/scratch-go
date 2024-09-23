package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 1BBB9F78-E883-49D1-A998-7EB864FD4A0E
var IID_IFabricGetProvisionedConfigVersionListResult = syscall.GUID{Data1: 0x1BBB9F78, Data2: 0xE883, Data3: 0x49D1,
	Data4: [8]byte{0xA9, 0x98, 0x7E, 0xB8, 0x64, 0xFD, 0x4A, 0x0E}}

type IFabricGetProvisionedConfigVersionListResult struct {
	win32.IUnknown
}

func NewIFabricGetProvisionedConfigVersionListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetProvisionedConfigVersionListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetProvisionedConfigVersionListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetProvisionedConfigVersionListResult) IID() *syscall.GUID {
	return &IID_IFabricGetProvisionedConfigVersionListResult
}

func (this *IFabricGetProvisionedConfigVersionListResult) Get_ProvisionedConfigVersionList() *FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PROVISIONED_CONFIG_VERSION_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
