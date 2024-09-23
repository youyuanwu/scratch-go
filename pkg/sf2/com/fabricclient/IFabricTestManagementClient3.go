package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// A4B94AFD-0CB5-4010-8995-E58E9B6CA373
var IID_IFabricTestManagementClient3 = syscall.GUID{Data1: 0xA4B94AFD, Data2: 0x0CB5, Data3: 0x4010,
	Data4: [8]byte{0x89, 0x95, 0xE5, 0x8E, 0x9B, 0x6C, 0xA3, 0x73}}

type IFabricTestManagementClient3 struct {
	IFabricTestManagementClient2
}

func NewIFabricTestManagementClient3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricTestManagementClient3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricTestManagementClient3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricTestManagementClient3) IID() *syscall.GUID {
	return &IID_IFabricTestManagementClient3
}

func (this *IFabricTestManagementClient3) BeginStartNodeTransition(description *FABRIC_NODE_TRANSITION_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricTestManagementClient3) EndStartNodeTransition(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricTestManagementClient3) BeginGetNodeTransitionProgress(operationId syscall.GUID, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&operationId)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricTestManagementClient3) EndGetNodeTransitionProgress(context *IFabricAsyncOperationContext, result **IFabricNodeTransitionProgressResult) com.Error {
	addr := (*this.LpVtbl)[28]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
