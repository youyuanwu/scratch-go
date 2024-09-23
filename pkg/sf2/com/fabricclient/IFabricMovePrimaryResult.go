package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 66AC03F5-E61C-47A2-80FE-49309A02C92C
var IID_IFabricMovePrimaryResult = syscall.GUID{Data1: 0x66AC03F5, Data2: 0xE61C, Data3: 0x47A2,
	Data4: [8]byte{0x80, 0xFE, 0x49, 0x30, 0x9A, 0x02, 0xC9, 0x2C}}

type IFabricMovePrimaryResult struct {
	win32.IUnknown
}

func NewIFabricMovePrimaryResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricMovePrimaryResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricMovePrimaryResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricMovePrimaryResult) IID() *syscall.GUID {
	return &IID_IFabricMovePrimaryResult
}

func (this *IFabricMovePrimaryResult) Get_Result() *FABRIC_MOVE_PRIMARY_RESULT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_MOVE_PRIMARY_RESULT)(unsafe.Pointer(ret))
}
