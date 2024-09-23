package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 067F144A-E5BE-4F5E-A181-8B5593E20242
var IID_IFabricReplicator = syscall.GUID{Data1: 0x067F144A, Data2: 0xE5BE, Data3: 0x4F5E,
	Data4: [8]byte{0xA1, 0x81, 0x8B, 0x55, 0x93, 0xE2, 0x02, 0x42}}

type IFabricReplicator struct {
	win32.IUnknown
}

func NewIFabricReplicator(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricReplicator {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricReplicator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricReplicator) IID() *syscall.GUID {
	return &IID_IFabricReplicator
}

func (this *IFabricReplicator) BeginOpen(callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricReplicator) EndOpen(context *IFabricAsyncOperationContext, replicationAddress **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(replicationAddress)))
	com.AddToScope(replicationAddress)
	return com.Error(ret)
}

func (this *IFabricReplicator) BeginChangeRole(epoch *FABRIC_EPOCH, role int32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(epoch)), uintptr(role), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricReplicator) EndChangeRole(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricReplicator) BeginUpdateEpoch(epoch *FABRIC_EPOCH, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(epoch)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricReplicator) EndUpdateEpoch(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricReplicator) BeginClose(callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricReplicator) EndClose(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricReplicator) Abort() {
	addr := (*this.LpVtbl)[11]
	_, _, _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
}

func (this *IFabricReplicator) GetCurrentProgress(lastSequenceNumber *int64) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lastSequenceNumber)))
	return com.Error(ret)
}

func (this *IFabricReplicator) GetCatchUpCapability(fromSequenceNumber *int64) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fromSequenceNumber)))
	return com.Error(ret)
}
