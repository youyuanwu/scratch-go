package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 80D2155C-4FC2-4FDE-9696-C2F39B471C3D
var IID_IFabricAtomicGroupStateReplicator = syscall.GUID{Data1: 0x80D2155C, Data2: 0x4FC2, Data3: 0x4FDE,
	Data4: [8]byte{0x96, 0x96, 0xC2, 0xF3, 0x9B, 0x47, 0x1C, 0x3D}}

type IFabricAtomicGroupStateReplicator struct {
	win32.IUnknown
}

func NewIFabricAtomicGroupStateReplicator(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricAtomicGroupStateReplicator {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricAtomicGroupStateReplicator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricAtomicGroupStateReplicator) IID() *syscall.GUID {
	return &IID_IFabricAtomicGroupStateReplicator
}

func (this *IFabricAtomicGroupStateReplicator) CreateAtomicGroup(atomicGroupId *int64) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(atomicGroupId)))
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateReplicator) BeginReplicateAtomicGroupOperation(atomicGroupId int64, operationData *IFabricOperationData, callback *IFabricAsyncOperationCallback, operationSequenceNumber *int64, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(atomicGroupId), uintptr(unsafe.Pointer(operationData)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(operationSequenceNumber)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateReplicator) EndReplicateAtomicGroupOperation(context *IFabricAsyncOperationContext, operationSequenceNumber *int64) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(operationSequenceNumber)))
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateReplicator) BeginReplicateAtomicGroupCommit(atomicGroupId int64, callback *IFabricAsyncOperationCallback, commitSequenceNumber *int64, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(atomicGroupId), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(commitSequenceNumber)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateReplicator) EndReplicateAtomicGroupCommit(context *IFabricAsyncOperationContext, commitSequenceNumber *int64) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(commitSequenceNumber)))
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateReplicator) BeginReplicateAtomicGroupRollback(atomicGroupId int64, callback *IFabricAsyncOperationCallback, rollbackSequenceNumber *int64, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(atomicGroupId), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(rollbackSequenceNumber)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateReplicator) EndReplicateAtomicGroupRollback(context *IFabricAsyncOperationContext, rollbackSequenceNumber *int64) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(rollbackSequenceNumber)))
	return com.Error(ret)
}
