package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// D7490E43-2217-4158-93E1-9CE4DD6F724A
var IID_IFabricApplicationManagementClient5 = syscall.GUID{Data1: 0xD7490E43, Data2: 0x2217, Data3: 0x4158,
	Data4: [8]byte{0x93, 0xE1, 0x9C, 0xE4, 0xDD, 0x6F, 0x72, 0x4A}}

type IFabricApplicationManagementClient5 struct {
	IFabricApplicationManagementClient4
}

func NewIFabricApplicationManagementClient5(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationManagementClient5 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationManagementClient5)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationManagementClient5) IID() *syscall.GUID {
	return &IID_IFabricApplicationManagementClient5
}

func (this *IFabricApplicationManagementClient5) BeginRollbackApplicationUpgrade(applicationName string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[29]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationName)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient5) EndRollbackApplicationUpgrade(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[30]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
