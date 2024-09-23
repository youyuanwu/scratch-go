package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 23E4EE1B-049A-48F5-8DD7-B601EACE47DE
var IID_IFabricServiceManagementClient6 = syscall.GUID{Data1: 0x23E4EE1B, Data2: 0x049A, Data3: 0x48F5,
	Data4: [8]byte{0x8D, 0xD7, 0xB6, 0x01, 0xEA, 0xCE, 0x47, 0xDE}}

type IFabricServiceManagementClient6 struct {
	IFabricServiceManagementClient5
}

func NewIFabricServiceManagementClient6(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceManagementClient6 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceManagementClient6)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceManagementClient6) IID() *syscall.GUID {
	return &IID_IFabricServiceManagementClient6
}

func (this *IFabricServiceManagementClient6) BeginCreateServiceFromTemplate2(serviceFromTemplateDescription *FABRIC_SERVICE_FROM_TEMPLATE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[29]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceFromTemplateDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient6) EndCreateServiceFromTemplate2(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[30]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
