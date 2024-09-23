package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// FE15A879-0DBE-4841-9CC6-6E92077CD669
var IID_IFabricSecretsResult = syscall.GUID{Data1: 0xFE15A879, Data2: 0x0DBE, Data3: 0x4841,
	Data4: [8]byte{0x9C, 0xC6, 0x6E, 0x92, 0x07, 0x7C, 0xD6, 0x69}}

type IFabricSecretsResult struct {
	win32.IUnknown
}

func NewIFabricSecretsResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricSecretsResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricSecretsResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricSecretsResult) IID() *syscall.GUID {
	return &IID_IFabricSecretsResult
}

func (this *IFabricSecretsResult) Get_Secrets() *FABRIC_SECRET_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SECRET_LIST)(unsafe.Pointer(ret))
}
