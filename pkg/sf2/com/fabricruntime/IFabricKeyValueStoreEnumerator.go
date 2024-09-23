package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 6722B848-15BB-4528-BF54-C7BBE27B6F9A
var IID_IFabricKeyValueStoreEnumerator = syscall.GUID{Data1: 0x6722B848, Data2: 0x15BB, Data3: 0x4528,
	Data4: [8]byte{0xBF, 0x54, 0xC7, 0xBB, 0xE2, 0x7B, 0x6F, 0x9A}}

type IFabricKeyValueStoreEnumerator struct {
	win32.IUnknown
}

func NewIFabricKeyValueStoreEnumerator(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreEnumerator {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreEnumerator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreEnumerator) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreEnumerator
}

func (this *IFabricKeyValueStoreEnumerator) EnumerateByKey(keyPrefix string, result **IFabricKeyValueStoreItemEnumerator) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(keyPrefix)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreEnumerator) EnumerateMetadataByKey(keyPrefix string, result **IFabricKeyValueStoreItemMetadataEnumerator) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(keyPrefix)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
