package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 564E50DD-C3A4-4600-A60E-6658874307AE
var IID_IFabricPrimaryReplicator = syscall.GUID{Data1: 0x564E50DD, Data2: 0xC3A4, Data3: 0x4600,
	Data4: [8]byte{0xA6, 0x0E, 0x66, 0x58, 0x87, 0x43, 0x07, 0xAE}}

type IFabricPrimaryReplicator struct {
	IFabricReplicator
}

func NewIFabricPrimaryReplicator(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPrimaryReplicator {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPrimaryReplicator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPrimaryReplicator) IID() *syscall.GUID {
	return &IID_IFabricPrimaryReplicator
}

func (this *IFabricPrimaryReplicator) BeginOnDataLoss(callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricPrimaryReplicator) EndOnDataLoss(context *IFabricAsyncOperationContext, isStateChanged *int8) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(isStateChanged)))
	return com.Error(ret)
}

func (this *IFabricPrimaryReplicator) UpdateCatchUpReplicaSetConfiguration(currentConfiguration *FABRIC_REPLICA_SET_CONFIGURATION, previousConfiguration *FABRIC_REPLICA_SET_CONFIGURATION) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(currentConfiguration)), uintptr(unsafe.Pointer(previousConfiguration)))
	return com.Error(ret)
}

func (this *IFabricPrimaryReplicator) BeginWaitForCatchUpQuorum(catchUpMode int32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(catchUpMode), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricPrimaryReplicator) EndWaitForCatchUpQuorum(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricPrimaryReplicator) UpdateCurrentReplicaSetConfiguration(currentConfiguration *FABRIC_REPLICA_SET_CONFIGURATION) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(currentConfiguration)))
	return com.Error(ret)
}

func (this *IFabricPrimaryReplicator) BeginBuildReplica(replica *FABRIC_REPLICA_INFORMATION, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(replica)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricPrimaryReplicator) EndBuildReplica(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricPrimaryReplicator) RemoveReplica(replicaId int64) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(replicaId))
	return com.Error(ret)
}
