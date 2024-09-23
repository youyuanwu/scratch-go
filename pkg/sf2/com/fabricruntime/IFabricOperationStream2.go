package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 0930199B-590A-4065-BEC9-5F93B6AAE086
var IID_IFabricOperationStream2 = syscall.GUID{Data1: 0x0930199B, Data2: 0x590A, Data3: 0x4065,
	Data4: [8]byte{0xBE, 0xC9, 0x5F, 0x93, 0xB6, 0xAA, 0xE0, 0x86}}

type IFabricOperationStream2 struct {
	IFabricOperationStream
}

func NewIFabricOperationStream2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricOperationStream2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricOperationStream2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricOperationStream2) IID() *syscall.GUID {
	return &IID_IFabricOperationStream2
}

func (this *IFabricOperationStream2) ReportFault(faultType int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(faultType))
	return com.Error(ret)
}
