package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 2C850629-6A83-4FC3-8468-C868B87E9A17
var IID_IFabricQueryClient8 = syscall.GUID{Data1: 0x2C850629, Data2: 0x6A83, Data3: 0x4FC3,
	Data4: [8]byte{0x84, 0x68, 0xC8, 0x68, 0xB8, 0x7E, 0x9A, 0x17}}

type IFabricQueryClient8 struct {
	IFabricQueryClient7
}

func NewIFabricQueryClient8(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricQueryClient8 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricQueryClient8)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricQueryClient8) IID() *syscall.GUID {
	return &IID_IFabricQueryClient8
}

func (this *IFabricQueryClient8) BeginGetServiceName(queryDescription *FABRIC_SERVICE_NAME_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[54]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient8) EndGetServiceName(context *IFabricAsyncOperationContext, result **IFabricGetServiceNameResult) com.Error {
	addr := (*this.LpVtbl)[55]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricQueryClient8) BeginGetApplicationName(queryDescription *FABRIC_APPLICATION_NAME_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[56]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient8) EndGetApplicationName(context *IFabricAsyncOperationContext, result **IFabricGetApplicationNameResult) com.Error {
	addr := (*this.LpVtbl)[57]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
