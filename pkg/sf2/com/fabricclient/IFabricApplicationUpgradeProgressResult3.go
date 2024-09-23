package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 1BC1D9C3-EEF5-41FE-B8A2-ABB97A8BA8E2
var IID_IFabricApplicationUpgradeProgressResult3 = syscall.GUID{Data1: 0x1BC1D9C3, Data2: 0xEEF5, Data3: 0x41FE,
	Data4: [8]byte{0xB8, 0xA2, 0xAB, 0xB9, 0x7A, 0x8B, 0xA8, 0xE2}}

type IFabricApplicationUpgradeProgressResult3 struct {
	IFabricApplicationUpgradeProgressResult2
}

func NewIFabricApplicationUpgradeProgressResult3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationUpgradeProgressResult3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationUpgradeProgressResult3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationUpgradeProgressResult3) IID() *syscall.GUID {
	return &IID_IFabricApplicationUpgradeProgressResult3
}

func (this *IFabricApplicationUpgradeProgressResult3) Get_UpgradeProgress() *FABRIC_APPLICATION_UPGRADE_PROGRESS {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_APPLICATION_UPGRADE_PROGRESS)(unsafe.Pointer(ret))
}
