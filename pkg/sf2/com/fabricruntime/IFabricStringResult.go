package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 4AE69614-7D0F-4CD4-B836-23017000D132
var IID_IFabricStringResult = syscall.GUID{Data1: 0x4AE69614, Data2: 0x7D0F, Data3: 0x4CD4,
	Data4: [8]byte{0xB8, 0x36, 0x23, 0x01, 0x70, 0x00, 0xD1, 0x32}}

type IFabricStringResult struct {
	win32.IUnknown
}

func NewIFabricStringResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStringResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStringResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStringResult) IID() *syscall.GUID {
	return &IID_IFabricStringResult
}

func (this *IFabricStringResult) Get_String() string {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}
