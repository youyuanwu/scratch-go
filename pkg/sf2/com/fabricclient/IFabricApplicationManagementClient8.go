package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 97B38E85-7329-47FF-A8D2-B7CBF1603689
var IID_IFabricApplicationManagementClient8 = syscall.GUID{Data1: 0x97B38E85, Data2: 0x7329, Data3: 0x47FF,
	Data4: [8]byte{0xA8, 0xD2, 0xB7, 0xCB, 0xF1, 0x60, 0x36, 0x89}}

type IFabricApplicationManagementClient8 struct {
	IFabricApplicationManagementClient7
}

func NewIFabricApplicationManagementClient8(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationManagementClient8 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationManagementClient8)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationManagementClient8) IID() *syscall.GUID {
	return &IID_IFabricApplicationManagementClient8
}

func (this *IFabricApplicationManagementClient8) BeginProvisionApplicationType2(description *FABRIC_PROVISION_APPLICATION_TYPE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[35]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient8) EndProvisionApplicationType2(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[36]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
