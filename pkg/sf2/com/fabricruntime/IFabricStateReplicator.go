package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 89E9A978-C771-44F2-92E8-3BF271CABE9C
var IID_IFabricStateReplicator = syscall.GUID{Data1: 0x89E9A978, Data2: 0xC771, Data3: 0x44F2,
	Data4: [8]byte{0x92, 0xE8, 0x3B, 0xF2, 0x71, 0xCA, 0xBE, 0x9C}}

type IFabricStateReplicator struct {
	win32.IUnknown
}

func NewIFabricStateReplicator(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStateReplicator {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStateReplicator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStateReplicator) IID() *syscall.GUID {
	return &IID_IFabricStateReplicator
}

func (this *IFabricStateReplicator) BeginReplicate(operationData *IFabricOperationData, callback *IFabricAsyncOperationCallback, sequenceNumber *int64, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(operationData)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(sequenceNumber)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricStateReplicator) EndReplicate(context *IFabricAsyncOperationContext, sequenceNumber *int64) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(sequenceNumber)))
	return com.Error(ret)
}

func (this *IFabricStateReplicator) GetReplicationStream(stream **IFabricOperationStream) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	com.AddToScope(stream)
	return com.Error(ret)
}

func (this *IFabricStateReplicator) GetCopyStream(stream **IFabricOperationStream) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	com.AddToScope(stream)
	return com.Error(ret)
}

func (this *IFabricStateReplicator) UpdateReplicatorSettings(replicatorSettings *FABRIC_REPLICATOR_SETTINGS) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(replicatorSettings)))
	return com.Error(ret)
}
