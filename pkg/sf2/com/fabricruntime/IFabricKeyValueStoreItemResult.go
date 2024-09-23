package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// C1F1C89D-B0B8-44DC-BC97-6C074C1A805E
var IID_IFabricKeyValueStoreItemResult = syscall.GUID{Data1: 0xC1F1C89D, Data2: 0xB0B8, Data3: 0x44DC,
	Data4: [8]byte{0xBC, 0x97, 0x6C, 0x07, 0x4C, 0x1A, 0x80, 0x5E}}

type IFabricKeyValueStoreItemResult struct {
	win32.IUnknown
}

func NewIFabricKeyValueStoreItemResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreItemResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreItemResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreItemResult) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreItemResult
}

func (this *IFabricKeyValueStoreItemResult) Get_Item() *FABRIC_KEY_VALUE_STORE_ITEM {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_KEY_VALUE_STORE_ITEM)(unsafe.Pointer(ret))
}
