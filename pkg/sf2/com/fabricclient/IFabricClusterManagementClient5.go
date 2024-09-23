package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// A6DDD816-A100-11E4-89D3-123B93F75CBA
var IID_IFabricClusterManagementClient5 = syscall.GUID{Data1: 0xA6DDD816, Data2: 0xA100, Data3: 0x11E4,
	Data4: [8]byte{0x89, 0xD3, 0x12, 0x3B, 0x93, 0xF7, 0x5C, 0xBA}}

type IFabricClusterManagementClient5 struct {
	IFabricClusterManagementClient4
}

func NewIFabricClusterManagementClient5(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClusterManagementClient5 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClusterManagementClient5)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClusterManagementClient5) IID() *syscall.GUID {
	return &IID_IFabricClusterManagementClient5
}

func (this *IFabricClusterManagementClient5) BeginResetPartitionLoad(partitionId syscall.GUID, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[43]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&partitionId)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient5) EndResetPartitionLoad(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[44]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
