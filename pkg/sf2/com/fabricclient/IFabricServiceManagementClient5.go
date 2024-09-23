package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F9A70679-8CA3-4E27-9411-483E0C89B1FA
var IID_IFabricServiceManagementClient5 = syscall.GUID{Data1: 0xF9A70679, Data2: 0x8CA3, Data3: 0x4E27,
	Data4: [8]byte{0x94, 0x11, 0x48, 0x3E, 0x0C, 0x89, 0xB1, 0xFA}}

type IFabricServiceManagementClient5 struct {
	IFabricServiceManagementClient4
}

func NewIFabricServiceManagementClient5(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceManagementClient5 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceManagementClient5)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceManagementClient5) IID() *syscall.GUID {
	return &IID_IFabricServiceManagementClient5
}

func (this *IFabricServiceManagementClient5) BeginDeleteService2(deleteDescription *FABRIC_DELETE_SERVICE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deleteDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient5) EndDeleteService2(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[28]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
