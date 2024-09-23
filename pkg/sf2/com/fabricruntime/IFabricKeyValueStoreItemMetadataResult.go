package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 17C483A1-69E6-4BDC-A058-54FD4A1839FD
var IID_IFabricKeyValueStoreItemMetadataResult = syscall.GUID{Data1: 0x17C483A1, Data2: 0x69E6, Data3: 0x4BDC,
	Data4: [8]byte{0xA0, 0x58, 0x54, 0xFD, 0x4A, 0x18, 0x39, 0xFD}}

type IFabricKeyValueStoreItemMetadataResult struct {
	win32.IUnknown
}

func NewIFabricKeyValueStoreItemMetadataResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreItemMetadataResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreItemMetadataResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreItemMetadataResult) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreItemMetadataResult
}

func (this *IFabricKeyValueStoreItemMetadataResult) Get_Metadata() *FABRIC_KEY_VALUE_STORE_ITEM_METADATA {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_KEY_VALUE_STORE_ITEM_METADATA)(unsafe.Pointer(ret))
}
