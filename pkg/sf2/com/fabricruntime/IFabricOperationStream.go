package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// A98FB97A-D6B0-408A-A878-A9EDB09C2587
var IID_IFabricOperationStream = syscall.GUID{Data1: 0xA98FB97A, Data2: 0xD6B0, Data3: 0x408A,
	Data4: [8]byte{0xA8, 0x78, 0xA9, 0xED, 0xB0, 0x9C, 0x25, 0x87}}

type IFabricOperationStream struct {
	win32.IUnknown
}

func NewIFabricOperationStream(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricOperationStream {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricOperationStream)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricOperationStream) IID() *syscall.GUID {
	return &IID_IFabricOperationStream
}

func (this *IFabricOperationStream) BeginGetOperation(callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricOperationStream) EndGetOperation(context *IFabricAsyncOperationContext, operation **IFabricOperation) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(operation)))
	com.AddToScope(operation)
	return com.Error(ret)
}
