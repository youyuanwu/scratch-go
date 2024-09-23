package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// EF25BC08-BE76-43C7-ADAD-20F01FBA3399
var IID_IFabricKeyValueStoreNotificationEnumerator = syscall.GUID{Data1: 0xEF25BC08, Data2: 0xBE76, Data3: 0x43C7,
	Data4: [8]byte{0xAD, 0xAD, 0x20, 0xF0, 0x1F, 0xBA, 0x33, 0x99}}

type IFabricKeyValueStoreNotificationEnumerator struct {
	win32.IUnknown
}

func NewIFabricKeyValueStoreNotificationEnumerator(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreNotificationEnumerator {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreNotificationEnumerator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreNotificationEnumerator) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreNotificationEnumerator
}

func (this *IFabricKeyValueStoreNotificationEnumerator) MoveNext() com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreNotificationEnumerator) Get_Current() *IFabricKeyValueStoreNotification {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*IFabricKeyValueStoreNotification)(unsafe.Pointer(ret))
}

func (this *IFabricKeyValueStoreNotificationEnumerator) Reset() {
	addr := (*this.LpVtbl)[5]
	_, _, _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
}
