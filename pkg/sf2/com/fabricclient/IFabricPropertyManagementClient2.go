package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 04991C28-3F9D-4A49-9322-A56D308965FD
var IID_IFabricPropertyManagementClient2 = syscall.GUID{Data1: 0x04991C28, Data2: 0x3F9D, Data3: 0x4A49,
	Data4: [8]byte{0x93, 0x22, 0xA5, 0x6D, 0x30, 0x89, 0x65, 0xFD}}

type IFabricPropertyManagementClient2 struct {
	IFabricPropertyManagementClient
}

func NewIFabricPropertyManagementClient2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPropertyManagementClient2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPropertyManagementClient2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPropertyManagementClient2) IID() *syscall.GUID {
	return &IID_IFabricPropertyManagementClient2
}

func (this *IFabricPropertyManagementClient2) BeginPutCustomPropertyOperation(name string, propertyOperation *FABRIC_PUT_CUSTOM_PROPERTY_OPERATION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[31]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(propertyOperation)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricPropertyManagementClient2) EndPutCustomPropertyOperation(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[32]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
