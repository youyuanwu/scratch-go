package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 60FE896A-B690-4ABB-94FD-86C615D29BEE
var IID_IFabricMoveSecondaryResult = syscall.GUID{Data1: 0x60FE896A, Data2: 0xB690, Data3: 0x4ABB,
	Data4: [8]byte{0x94, 0xFD, 0x86, 0xC6, 0x15, 0xD2, 0x9B, 0xEE}}

type IFabricMoveSecondaryResult struct {
	win32.IUnknown
}

func NewIFabricMoveSecondaryResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricMoveSecondaryResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricMoveSecondaryResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricMoveSecondaryResult) IID() *syscall.GUID {
	return &IID_IFabricMoveSecondaryResult
}

func (this *IFabricMoveSecondaryResult) Get_Result() *FABRIC_MOVE_SECONDARY_RESULT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_MOVE_SECONDARY_RESULT)(unsafe.Pointer(ret))
}
