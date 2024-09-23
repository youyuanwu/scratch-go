package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 26844276-25B1-4F8C-ADBE-B1B3A3083C17
var IID_IFabricApplicationManagementClient7 = syscall.GUID{Data1: 0x26844276, Data2: 0x25B1, Data3: 0x4F8C,
	Data4: [8]byte{0xAD, 0xBE, 0xB1, 0xB3, 0xA3, 0x08, 0x3C, 0x17}}

type IFabricApplicationManagementClient7 struct {
	IFabricApplicationManagementClient6
}

func NewIFabricApplicationManagementClient7(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationManagementClient7 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationManagementClient7)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationManagementClient7) IID() *syscall.GUID {
	return &IID_IFabricApplicationManagementClient7
}

func (this *IFabricApplicationManagementClient7) BeginDeleteApplication2(deleteDescription *FABRIC_DELETE_APPLICATION_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[33]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deleteDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient7) EndDeleteApplication2(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[34]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
