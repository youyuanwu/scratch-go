package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 86F08D7E-14DD-4575-8489-B1D5D679029C
var IID_IFabricAsyncOperationCallback = syscall.GUID{Data1: 0x86F08D7E, Data2: 0x14DD, Data3: 0x4575,
	Data4: [8]byte{0x84, 0x89, 0xB1, 0xD5, 0xD6, 0x79, 0x02, 0x9C}}

type IFabricAsyncOperationCallback struct {
	win32.IUnknown
}

func NewIFabricAsyncOperationCallback(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricAsyncOperationCallback {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricAsyncOperationCallback)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricAsyncOperationCallback) IID() *syscall.GUID {
	return &IID_IFabricAsyncOperationCallback
}

func (this *IFabricAsyncOperationCallback) Invoke(context *IFabricAsyncOperationContext) {
	addr := (*this.LpVtbl)[3]
	_, _, _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
}
