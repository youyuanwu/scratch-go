package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 02139DA8-7140-42AE-8403-79A551600E63
var IID_IFabricQueryClient10 = syscall.GUID{Data1: 0x02139DA8, Data2: 0x7140, Data3: 0x42AE,
	Data4: [8]byte{0x84, 0x03, 0x79, 0xA5, 0x51, 0x60, 0x0E, 0x63}}

type IFabricQueryClient10 struct {
	IFabricQueryClient9
}

func NewIFabricQueryClient10(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricQueryClient10 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricQueryClient10)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricQueryClient10) IID() *syscall.GUID {
	return &IID_IFabricQueryClient10
}

func (this *IFabricQueryClient10) BeginGetDeployedApplicationPagedList(queryDescription *FABRIC_PAGED_DEPLOYED_APPLICATION_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[60]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient10) EndGetDeployedApplicationPagedList(context *IFabricAsyncOperationContext, result **IFabricGetDeployedApplicationPagedListResult) com.Error {
	addr := (*this.LpVtbl)[61]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
