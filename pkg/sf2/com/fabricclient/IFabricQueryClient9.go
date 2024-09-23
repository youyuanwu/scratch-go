package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 76F0B4A5-4941-49D7-993C-AD7AFC37C6AF
var IID_IFabricQueryClient9 = syscall.GUID{Data1: 0x76F0B4A5, Data2: 0x4941, Data3: 0x49D7,
	Data4: [8]byte{0x99, 0x3C, 0xAD, 0x7A, 0xFC, 0x37, 0xC6, 0xAF}}

type IFabricQueryClient9 struct {
	IFabricQueryClient8
}

func NewIFabricQueryClient9(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricQueryClient9 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricQueryClient9)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricQueryClient9) IID() *syscall.GUID {
	return &IID_IFabricQueryClient9
}

func (this *IFabricQueryClient9) BeginGetApplicationTypePagedList(queryDescription *PAGED_FABRIC_APPLICATION_TYPE_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[58]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient9) EndGetApplicationTypePagedList(context *IFabricAsyncOperationContext, result **IFabricGetApplicationTypePagedListResult) com.Error {
	addr := (*this.LpVtbl)[59]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
