package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 26617B63-1350-4D7F-830C-2200978D31BB
var IID_IFabricApplicationManagementClient9 = syscall.GUID{Data1: 0x26617B63, Data2: 0x1350, Data3: 0x4D7F,
	Data4: [8]byte{0x83, 0x0C, 0x22, 0x00, 0x97, 0x8D, 0x31, 0xBB}}

type IFabricApplicationManagementClient9 struct {
	IFabricApplicationManagementClient8
}

func NewIFabricApplicationManagementClient9(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationManagementClient9 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationManagementClient9)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationManagementClient9) IID() *syscall.GUID {
	return &IID_IFabricApplicationManagementClient9
}

func (this *IFabricApplicationManagementClient9) BeginUnprovisionApplicationType2(description *FABRIC_UNPROVISION_APPLICATION_TYPE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[37]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient9) EndUnprovisionApplicationType2(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[38]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
