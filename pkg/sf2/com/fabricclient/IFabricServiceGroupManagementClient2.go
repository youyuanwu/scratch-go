package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 4F0DC42D-8FEC-4EA9-A96B-5BE1FA1E1D64
var IID_IFabricServiceGroupManagementClient2 = syscall.GUID{Data1: 0x4F0DC42D, Data2: 0x8FEC, Data3: 0x4EA9,
	Data4: [8]byte{0xA9, 0x6B, 0x5B, 0xE1, 0xFA, 0x1E, 0x1D, 0x64}}

type IFabricServiceGroupManagementClient2 struct {
	IFabricServiceGroupManagementClient
}

func NewIFabricServiceGroupManagementClient2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceGroupManagementClient2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceGroupManagementClient2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceGroupManagementClient2) IID() *syscall.GUID {
	return &IID_IFabricServiceGroupManagementClient2
}

func (this *IFabricServiceGroupManagementClient2) BeginUpdateServiceGroup(name string, serviceGroupUpdateDescription *FABRIC_SERVICE_GROUP_UPDATE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(serviceGroupUpdateDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceGroupManagementClient2) EndUpdateServiceGroup(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
