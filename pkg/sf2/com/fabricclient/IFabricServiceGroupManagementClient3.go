package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// CBEE0E12-B5A0-44DC-8C3C-C067958F82F6
var IID_IFabricServiceGroupManagementClient3 = syscall.GUID{Data1: 0xCBEE0E12, Data2: 0xB5A0, Data3: 0x44DC,
	Data4: [8]byte{0x8C, 0x3C, 0xC0, 0x67, 0x95, 0x8F, 0x82, 0xF6}}

type IFabricServiceGroupManagementClient3 struct {
	IFabricServiceGroupManagementClient2
}

func NewIFabricServiceGroupManagementClient3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceGroupManagementClient3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceGroupManagementClient3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceGroupManagementClient3) IID() *syscall.GUID {
	return &IID_IFabricServiceGroupManagementClient3
}

func (this *IFabricServiceGroupManagementClient3) BeginCreateServiceGroupFromTemplate(applicationName string, serviceName string, serviceTypeName string, initializationDataSize uint32, initializationData *byte, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationName)), uintptr(win32.StrToPointer(serviceName)), uintptr(win32.StrToPointer(serviceTypeName)), uintptr(initializationDataSize), uintptr(unsafe.Pointer(initializationData)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceGroupManagementClient3) EndCreateServiceGroupFromTemplate(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
