package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// CB660AA6-C51E-4F05-9526-93982B550E8F
var IID_IFabricKeyValueStoreNotification = syscall.GUID{Data1: 0xCB660AA6, Data2: 0xC51E, Data3: 0x4F05,
	Data4: [8]byte{0x95, 0x26, 0x93, 0x98, 0x2B, 0x55, 0x0E, 0x8F}}

type IFabricKeyValueStoreNotification struct {
	IFabricKeyValueStoreItemResult
}

func NewIFabricKeyValueStoreNotification(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreNotification {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreNotification)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreNotification) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreNotification
}

func (this *IFabricKeyValueStoreNotification) IsDelete() int8 {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int8(ret)
}
