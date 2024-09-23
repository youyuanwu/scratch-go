package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 68A98626-6A1B-4DD8-AD93-74C0936E86AA
var IID_IFabricNodeTransitionProgressResult = syscall.GUID{Data1: 0x68A98626, Data2: 0x6A1B, Data3: 0x4DD8,
	Data4: [8]byte{0xAD, 0x93, 0x74, 0xC0, 0x93, 0x6E, 0x86, 0xAA}}

type IFabricNodeTransitionProgressResult struct {
	win32.IUnknown
}

func NewIFabricNodeTransitionProgressResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricNodeTransitionProgressResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricNodeTransitionProgressResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricNodeTransitionProgressResult) IID() *syscall.GUID {
	return &IID_IFabricNodeTransitionProgressResult
}

func (this *IFabricNodeTransitionProgressResult) Get_Progress() *FABRIC_NODE_TRANSITION_PROGRESS {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NODE_TRANSITION_PROGRESS)(unsafe.Pointer(ret))
}
