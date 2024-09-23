package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// C6FB97F7-82F3-4E6C-A80A-021E8FFCA425
var IID_IFabricClientSettings2 = syscall.GUID{Data1: 0xC6FB97F7, Data2: 0x82F3, Data3: 0x4E6C,
	Data4: [8]byte{0xA8, 0x0A, 0x02, 0x1E, 0x8F, 0xFC, 0xA4, 0x25}}

type IFabricClientSettings2 struct {
	IFabricClientSettings
}

func NewIFabricClientSettings2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClientSettings2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClientSettings2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClientSettings2) IID() *syscall.GUID {
	return &IID_IFabricClientSettings2
}

func (this *IFabricClientSettings2) GetSettings(result **IFabricClientSettingsResult) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricClientSettings2) SetSettings(fabricClientSettings *FABRIC_CLIENT_SETTINGS) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fabricClientSettings)))
	return com.Error(ret)
}
