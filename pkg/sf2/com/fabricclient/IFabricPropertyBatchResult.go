package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// EE747FF5-3FBB-46A8-ADBC-47CE09C48BBE
var IID_IFabricPropertyBatchResult = syscall.GUID{Data1: 0xEE747FF5, Data2: 0x3FBB, Data3: 0x46A8,
	Data4: [8]byte{0xAD, 0xBC, 0x47, 0xCE, 0x09, 0xC4, 0x8B, 0xBE}}

type IFabricPropertyBatchResult struct {
	win32.IUnknown
}

func NewIFabricPropertyBatchResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPropertyBatchResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPropertyBatchResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPropertyBatchResult) IID() *syscall.GUID {
	return &IID_IFabricPropertyBatchResult
}

func (this *IFabricPropertyBatchResult) GetProperty(operationIndexInRequest uint32, property **IFabricPropertyValueResult) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(operationIndexInRequest), uintptr(unsafe.Pointer(property)))
	com.AddToScope(property)
	return com.Error(ret)
}
