package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 8180DB27-7D0B-43B0-82E0-4A8E022FC238
var IID_IFabricServiceManagementClient4 = syscall.GUID{Data1: 0x8180DB27, Data2: 0x7D0B, Data3: 0x43B0,
	Data4: [8]byte{0x82, 0xE0, 0x4A, 0x8E, 0x02, 0x2F, 0xC2, 0x38}}

type IFabricServiceManagementClient4 struct {
	IFabricServiceManagementClient3
}

func NewIFabricServiceManagementClient4(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceManagementClient4 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceManagementClient4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceManagementClient4) IID() *syscall.GUID {
	return &IID_IFabricServiceManagementClient4
}

func (this *IFabricServiceManagementClient4) BeginRegisterServiceNotificationFilter(description *FABRIC_SERVICE_NOTIFICATION_FILTER_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient4) EndRegisterServiceNotificationFilter(context *IFabricAsyncOperationContext, filterId *int64) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(filterId)))
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient4) BeginUnregisterServiceNotificationFilter(filterId int64, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(filterId), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient4) EndUnregisterServiceNotificationFilter(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
