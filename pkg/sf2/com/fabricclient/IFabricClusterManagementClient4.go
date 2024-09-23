package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// B6B12671-F283-4D71-A818-0260549BC83E
var IID_IFabricClusterManagementClient4 = syscall.GUID{Data1: 0xB6B12671, Data2: 0xF283, Data3: 0x4D71,
	Data4: [8]byte{0xA8, 0x18, 0x02, 0x60, 0x54, 0x9B, 0xC8, 0x3E}}

type IFabricClusterManagementClient4 struct {
	IFabricClusterManagementClient3
}

func NewIFabricClusterManagementClient4(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClusterManagementClient4 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClusterManagementClient4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClusterManagementClient4) IID() *syscall.GUID {
	return &IID_IFabricClusterManagementClient4
}

func (this *IFabricClusterManagementClient4) BeginRollbackFabricUpgrade(timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[41]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient4) EndRollbackFabricUpgrade(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[42]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
