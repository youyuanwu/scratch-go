package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 3C73B32E-9A08-48CA-B3A3-993A2029E37A
var IID_IFabricServiceGroupManagementClient4 = syscall.GUID{Data1: 0x3C73B32E, Data2: 0x9A08, Data3: 0x48CA,
	Data4: [8]byte{0xB3, 0xA3, 0x99, 0x3A, 0x20, 0x29, 0xE3, 0x7A}}

type IFabricServiceGroupManagementClient4 struct {
	IFabricServiceGroupManagementClient3
}

func NewIFabricServiceGroupManagementClient4(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceGroupManagementClient4 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceGroupManagementClient4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceGroupManagementClient4) IID() *syscall.GUID {
	return &IID_IFabricServiceGroupManagementClient4
}

func (this *IFabricServiceGroupManagementClient4) BeginCreateServiceGroupFromTemplate2(serviceGroupFromTemplateDescription *FABRIC_SERVICE_GROUP_FROM_TEMPLATE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceGroupFromTemplateDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceGroupManagementClient4) EndCreateServiceGroupFromTemplate2(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
