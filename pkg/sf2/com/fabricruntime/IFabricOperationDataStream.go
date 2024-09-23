package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// C4E9084C-BE92-49C9-8C18-D44D088C2E32
var IID_IFabricOperationDataStream = syscall.GUID{Data1: 0xC4E9084C, Data2: 0xBE92, Data3: 0x49C9,
	Data4: [8]byte{0x8C, 0x18, 0xD4, 0x4D, 0x08, 0x8C, 0x2E, 0x32}}

type IFabricOperationDataStream struct {
	win32.IUnknown
}

func NewIFabricOperationDataStream(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricOperationDataStream {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricOperationDataStream)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricOperationDataStream) IID() *syscall.GUID {
	return &IID_IFabricOperationDataStream
}

func (this *IFabricOperationDataStream) BeginGetNext(callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricOperationDataStream) EndGetNext(context *IFabricAsyncOperationContext, operationData **IFabricOperationData) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(operationData)))
	com.AddToScope(operationData)
	return com.Error(ret)
}
