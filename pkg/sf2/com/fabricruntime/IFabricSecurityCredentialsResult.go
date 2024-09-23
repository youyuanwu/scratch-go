package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 049A111D-6A30-48E9-8F69-470760D3EFB9
var IID_IFabricSecurityCredentialsResult = syscall.GUID{Data1: 0x049A111D, Data2: 0x6A30, Data3: 0x48E9,
	Data4: [8]byte{0x8F, 0x69, 0x47, 0x07, 0x60, 0xD3, 0xEF, 0xB9}}

type IFabricSecurityCredentialsResult struct {
	win32.IUnknown
}

func NewIFabricSecurityCredentialsResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricSecurityCredentialsResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricSecurityCredentialsResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricSecurityCredentialsResult) IID() *syscall.GUID {
	return &IID_IFabricSecurityCredentialsResult
}

func (this *IFabricSecurityCredentialsResult) Get_SecurityCredentials() *FABRIC_SECURITY_CREDENTIALS {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SECURITY_CREDENTIALS)(unsafe.Pointer(ret))
}
