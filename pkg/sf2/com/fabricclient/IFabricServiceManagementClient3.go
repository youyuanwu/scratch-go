package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 98EC1156-C249-4F66-8D7C-9A5FA88E8E6D
var IID_IFabricServiceManagementClient3 = syscall.GUID{Data1: 0x98EC1156, Data2: 0xC249, Data3: 0x4F66,
	Data4: [8]byte{0x8D, 0x7C, 0x9A, 0x5F, 0xA8, 0x8E, 0x8E, 0x6D}}

type IFabricServiceManagementClient3 struct {
	IFabricServiceManagementClient2
}

func NewIFabricServiceManagementClient3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceManagementClient3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceManagementClient3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceManagementClient3) IID() *syscall.GUID {
	return &IID_IFabricServiceManagementClient3
}

func (this *IFabricServiceManagementClient3) BeginRemoveReplica(description *FABRIC_REMOVE_REPLICA_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient3) EndRemoveReplica(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient3) BeginRestartReplica(description *FABRIC_RESTART_REPLICA_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient3) EndRestartReplica(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
