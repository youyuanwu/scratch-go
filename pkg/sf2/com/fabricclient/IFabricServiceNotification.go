package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 8222C825-08AD-4639-AFCE-A8988CBD6DB3
var IID_IFabricServiceNotification = syscall.GUID{Data1: 0x8222C825, Data2: 0x08AD, Data3: 0x4639,
	Data4: [8]byte{0xAF, 0xCE, 0xA8, 0x98, 0x8C, 0xBD, 0x6D, 0xB3}}

type IFabricServiceNotification struct {
	win32.IUnknown
}

func NewIFabricServiceNotification(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceNotification {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceNotification)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceNotification) IID() *syscall.GUID {
	return &IID_IFabricServiceNotification
}

func (this *IFabricServiceNotification) Get_Notification() *FABRIC_SERVICE_NOTIFICATION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_NOTIFICATION)(unsafe.Pointer(ret))
}

func (this *IFabricServiceNotification) GetVersion(result **IFabricServiceEndpointsVersion) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
