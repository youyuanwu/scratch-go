package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 2B670953-6148-4F7D-A920-B390DE43D913
var IID_IFabricAtomicGroupStateProvider = syscall.GUID{Data1: 0x2B670953, Data2: 0x6148, Data3: 0x4F7D,
	Data4: [8]byte{0xA9, 0x20, 0xB3, 0x90, 0xDE, 0x43, 0xD9, 0x13}}

type IFabricAtomicGroupStateProvider struct {
	win32.IUnknown
}

func NewIFabricAtomicGroupStateProvider(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricAtomicGroupStateProvider {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricAtomicGroupStateProvider)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricAtomicGroupStateProvider) IID() *syscall.GUID {
	return &IID_IFabricAtomicGroupStateProvider
}

func (this *IFabricAtomicGroupStateProvider) BeginAtomicGroupCommit(atomicGroupId int64, commitSequenceNumber int64, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(atomicGroupId), uintptr(commitSequenceNumber), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateProvider) EndAtomicGroupCommit(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateProvider) BeginAtomicGroupRollback(atomicGroupId int64, rollbackequenceNumber int64, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(atomicGroupId), uintptr(rollbackequenceNumber), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateProvider) EndAtomicGroupRollback(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateProvider) BeginUndoProgress(fromCommitSequenceNumber int64, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(fromCommitSequenceNumber), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricAtomicGroupStateProvider) EndUndoProgress(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
