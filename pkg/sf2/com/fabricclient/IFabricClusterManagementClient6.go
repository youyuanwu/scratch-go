package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 9E454AE8-4B8C-4136-884A-37B0B92CC855
var IID_IFabricClusterManagementClient6 = syscall.GUID{Data1: 0x9E454AE8, Data2: 0x4B8C, Data3: 0x4136,
	Data4: [8]byte{0x88, 0x4A, 0x37, 0xB0, 0xB9, 0x2C, 0xC8, 0x55}}

type IFabricClusterManagementClient6 struct {
	IFabricClusterManagementClient5
}

func NewIFabricClusterManagementClient6(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClusterManagementClient6 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClusterManagementClient6)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClusterManagementClient6) IID() *syscall.GUID {
	return &IID_IFabricClusterManagementClient6
}

func (this *IFabricClusterManagementClient6) BeginToggleVerboseServicePlacementHealthReporting(enabled int8, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[45]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(enabled), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient6) EndToggleVerboseServicePlacementHealthReporting(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[46]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
