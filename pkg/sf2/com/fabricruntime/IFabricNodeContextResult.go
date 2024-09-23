package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 0952F885-6F5A-4ED3-ABE4-90C403D1E3CE
var IID_IFabricNodeContextResult = syscall.GUID{Data1: 0x0952F885, Data2: 0x6F5A, Data3: 0x4ED3,
	Data4: [8]byte{0xAB, 0xE4, 0x90, 0xC4, 0x03, 0xD1, 0xE3, 0xCE}}

type IFabricNodeContextResult struct {
	win32.IUnknown
}

func NewIFabricNodeContextResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricNodeContextResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricNodeContextResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricNodeContextResult) IID() *syscall.GUID {
	return &IID_IFabricNodeContextResult
}

func (this *IFabricNodeContextResult) Get_NodeContext() *FABRIC_NODE_CONTEXT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NODE_CONTEXT)(unsafe.Pointer(ret))
}
