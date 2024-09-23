package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// B01E63EE-1EA4-4181-95C7-983B32E16848
var IID_IFabricApplicationManagementClient6 = syscall.GUID{Data1: 0xB01E63EE, Data2: 0x1EA4, Data3: 0x4181,
	Data4: [8]byte{0x95, 0xC7, 0x98, 0x3B, 0x32, 0xE1, 0x68, 0x48}}

type IFabricApplicationManagementClient6 struct {
	IFabricApplicationManagementClient5
}

func NewIFabricApplicationManagementClient6(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationManagementClient6 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationManagementClient6)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationManagementClient6) IID() *syscall.GUID {
	return &IID_IFabricApplicationManagementClient6
}

func (this *IFabricApplicationManagementClient6) BeginUpdateApplication(applicationUpdateDescription *FABRIC_APPLICATION_UPDATE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[31]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(applicationUpdateDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient6) EndUpdateApplication(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[32]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
