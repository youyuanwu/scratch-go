package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// B0E7DEE0-CF64-11E0-9572-0800200C9A66
var IID_IFabricClientSettings = syscall.GUID{Data1: 0xB0E7DEE0, Data2: 0xCF64, Data3: 0x11E0,
	Data4: [8]byte{0x95, 0x72, 0x08, 0x00, 0x20, 0x0C, 0x9A, 0x66}}

type IFabricClientSettings struct {
	win32.IUnknown
}

func NewIFabricClientSettings(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClientSettings {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClientSettings)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClientSettings) IID() *syscall.GUID {
	return &IID_IFabricClientSettings
}

func (this *IFabricClientSettings) SetSecurityCredentials(securityCredentials *FABRIC_SECURITY_CREDENTIALS) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(securityCredentials)))
	return com.Error(ret)
}

func (this *IFabricClientSettings) SetKeepAlive(keepAliveIntervalInSeconds uint32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(keepAliveIntervalInSeconds))
	return com.Error(ret)
}
