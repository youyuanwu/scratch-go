package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 62707EE5-B625-4489-AA4D-2E54B06EA248
var IID_IFabricApplicationUpgradeProgressResult2 = syscall.GUID{Data1: 0x62707EE5, Data2: 0xB625, Data3: 0x4489,
	Data4: [8]byte{0xAA, 0x4D, 0x2E, 0x54, 0xB0, 0x6E, 0xA2, 0x48}}

type IFabricApplicationUpgradeProgressResult2 struct {
	IFabricApplicationUpgradeProgressResult
}

func NewIFabricApplicationUpgradeProgressResult2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationUpgradeProgressResult2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationUpgradeProgressResult2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationUpgradeProgressResult2) IID() *syscall.GUID {
	return &IID_IFabricApplicationUpgradeProgressResult2
}

func (this *IFabricApplicationUpgradeProgressResult2) Get_RollingUpgradeMode() int32 {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int32(ret)
}

func (this *IFabricApplicationUpgradeProgressResult2) Get_NextUpgradeDomain() string {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}
