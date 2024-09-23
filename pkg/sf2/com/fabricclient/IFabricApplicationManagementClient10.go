package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 67001225-D106-41AE-8BD4-5A0A119C5C01
var IID_IFabricApplicationManagementClient10 = syscall.GUID{Data1: 0x67001225, Data2: 0xD106, Data3: 0x41AE,
	Data4: [8]byte{0x8B, 0xD4, 0x5A, 0x0A, 0x11, 0x9C, 0x5C, 0x01}}

type IFabricApplicationManagementClient10 struct {
	IFabricApplicationManagementClient9
}

func NewIFabricApplicationManagementClient10(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationManagementClient10 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationManagementClient10)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationManagementClient10) IID() *syscall.GUID {
	return &IID_IFabricApplicationManagementClient10
}

func (this *IFabricApplicationManagementClient10) BeginProvisionApplicationType3(description *FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION_BASE, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[39]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient10) EndProvisionApplicationType3(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[40]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
