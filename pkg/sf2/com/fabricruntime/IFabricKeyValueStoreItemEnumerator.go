package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// C202788F-54D3-44A6-8F3C-B4BBFCDB95D2
var IID_IFabricKeyValueStoreItemEnumerator = syscall.GUID{Data1: 0xC202788F, Data2: 0x54D3, Data3: 0x44A6,
	Data4: [8]byte{0x8F, 0x3C, 0xB4, 0xBB, 0xFC, 0xDB, 0x95, 0xD2}}

type IFabricKeyValueStoreItemEnumerator struct {
	win32.IUnknown
}

func NewIFabricKeyValueStoreItemEnumerator(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreItemEnumerator {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreItemEnumerator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreItemEnumerator) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreItemEnumerator
}

func (this *IFabricKeyValueStoreItemEnumerator) MoveNext() com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreItemEnumerator) Get_Current() *IFabricKeyValueStoreItemResult {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*IFabricKeyValueStoreItemResult)(unsafe.Pointer(ret))
}
