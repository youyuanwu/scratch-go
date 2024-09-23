package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 0BC06AEE-FFFA-4450-9099-116A5F0E0B53
var IID_IFabricKeyValueStoreItemMetadataEnumerator = syscall.GUID{Data1: 0x0BC06AEE, Data2: 0xFFFA, Data3: 0x4450,
	Data4: [8]byte{0x90, 0x99, 0x11, 0x6A, 0x5F, 0x0E, 0x0B, 0x53}}

type IFabricKeyValueStoreItemMetadataEnumerator struct {
	win32.IUnknown
}

func NewIFabricKeyValueStoreItemMetadataEnumerator(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreItemMetadataEnumerator {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreItemMetadataEnumerator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreItemMetadataEnumerator) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreItemMetadataEnumerator
}

func (this *IFabricKeyValueStoreItemMetadataEnumerator) MoveNext() com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreItemMetadataEnumerator) Get_Current() *IFabricKeyValueStoreItemMetadataResult {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*IFabricKeyValueStoreItemMetadataResult)(unsafe.Pointer(ret))
}
