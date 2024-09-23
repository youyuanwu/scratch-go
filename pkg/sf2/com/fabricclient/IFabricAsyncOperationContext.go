package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 841720BF-C9E8-4E6F-9C3F-6B7F4AC73BCD
var IID_IFabricAsyncOperationContext = syscall.GUID{Data1: 0x841720BF, Data2: 0xC9E8, Data3: 0x4E6F,
	Data4: [8]byte{0x9C, 0x3F, 0x6B, 0x7F, 0x4A, 0xC7, 0x3B, 0xCD}}

type IFabricAsyncOperationContext struct {
	win32.IUnknown
}

func NewIFabricAsyncOperationContext(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricAsyncOperationContext {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricAsyncOperationContext)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricAsyncOperationContext) IID() *syscall.GUID {
	return &IID_IFabricAsyncOperationContext
}

func (this *IFabricAsyncOperationContext) IsCompleted() int8 {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int8(ret)
}

func (this *IFabricAsyncOperationContext) CompletedSynchronously() int8 {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int8(ret)
}

func (this *IFabricAsyncOperationContext) Get_Callback(callback **IFabricAsyncOperationCallback) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)))
	com.AddToScope(callback)
	return com.Error(ret)
}

func (this *IFabricAsyncOperationContext) Cancel() com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}
