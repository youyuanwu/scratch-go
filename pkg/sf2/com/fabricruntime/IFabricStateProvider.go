package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 3EBFEC79-BD27-43F3-8BE8-DA38EE723951
var IID_IFabricStateProvider = syscall.GUID{Data1: 0x3EBFEC79, Data2: 0xBD27, Data3: 0x43F3,
	Data4: [8]byte{0x8B, 0xE8, 0xDA, 0x38, 0xEE, 0x72, 0x39, 0x51}}

type IFabricStateProvider struct {
	win32.IUnknown
}

func NewIFabricStateProvider(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStateProvider {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStateProvider)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStateProvider) IID() *syscall.GUID {
	return &IID_IFabricStateProvider
}

func (this *IFabricStateProvider) BeginUpdateEpoch(epoch *FABRIC_EPOCH, previousEpochLastSequenceNumber int64, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(epoch)), uintptr(previousEpochLastSequenceNumber), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricStateProvider) EndUpdateEpoch(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricStateProvider) GetLastCommittedSequenceNumber(sequenceNumber *int64) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sequenceNumber)))
	return com.Error(ret)
}

func (this *IFabricStateProvider) BeginOnDataLoss(callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricStateProvider) EndOnDataLoss(context *IFabricAsyncOperationContext, isStateChanged *int8) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(isStateChanged)))
	return com.Error(ret)
}

func (this *IFabricStateProvider) GetCopyContext(copyContextStream **IFabricOperationDataStream) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(copyContextStream)))
	com.AddToScope(copyContextStream)
	return com.Error(ret)
}

func (this *IFabricStateProvider) GetCopyState(uptoSequenceNumber int64, copyContextStream *IFabricOperationDataStream, copyStateStream **IFabricOperationDataStream) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(uptoSequenceNumber), uintptr(unsafe.Pointer(copyContextStream)), uintptr(unsafe.Pointer(copyStateStream)))
	com.AddToScope(copyStateStream)
	return com.Error(ret)
}
