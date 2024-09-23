package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 8AE3BE0E-505D-4DC1-AD8F-0CB0F9576B8A
var IID_IFabricStatefulServiceReplica = syscall.GUID{Data1: 0x8AE3BE0E, Data2: 0x505D, Data3: 0x4DC1,
	Data4: [8]byte{0xAD, 0x8F, 0x0C, 0xB0, 0xF9, 0x57, 0x6B, 0x8A}}

type IFabricStatefulServiceReplica struct {
	win32.IUnknown
}

func NewIFabricStatefulServiceReplica(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatefulServiceReplica {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatefulServiceReplica)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatefulServiceReplica) IID() *syscall.GUID {
	return &IID_IFabricStatefulServiceReplica
}

func (this *IFabricStatefulServiceReplica) BeginOpen(openMode int32, partition *IFabricStatefulServicePartition, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(openMode), uintptr(unsafe.Pointer(partition)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricStatefulServiceReplica) EndOpen(context *IFabricAsyncOperationContext, replicator **IFabricReplicator) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(replicator)))
	com.AddToScope(replicator)
	return com.Error(ret)
}

func (this *IFabricStatefulServiceReplica) BeginChangeRole(newRole int32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(newRole), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricStatefulServiceReplica) EndChangeRole(context *IFabricAsyncOperationContext, serviceAddress **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(serviceAddress)))
	com.AddToScope(serviceAddress)
	return com.Error(ret)
}

func (this *IFabricStatefulServiceReplica) BeginClose(callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricStatefulServiceReplica) EndClose(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricStatefulServiceReplica) Abort() {
	addr := (*this.LpVtbl)[9]
	_, _, _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
}
