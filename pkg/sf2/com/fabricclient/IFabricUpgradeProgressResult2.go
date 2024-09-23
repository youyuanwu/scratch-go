package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 9CC0AAF3-0F6C-40A3-85AC-38338DD36D75
var IID_IFabricUpgradeProgressResult2 = syscall.GUID{Data1: 0x9CC0AAF3, Data2: 0x0F6C, Data3: 0x40A3,
	Data4: [8]byte{0x85, 0xAC, 0x38, 0x33, 0x8D, 0xD3, 0x6D, 0x75}}

type IFabricUpgradeProgressResult2 struct {
	IFabricUpgradeProgressResult
}

func NewIFabricUpgradeProgressResult2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricUpgradeProgressResult2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricUpgradeProgressResult2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricUpgradeProgressResult2) IID() *syscall.GUID {
	return &IID_IFabricUpgradeProgressResult2
}

func (this *IFabricUpgradeProgressResult2) Get_RollingUpgradeMode() int32 {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int32(ret)
}

func (this *IFabricUpgradeProgressResult2) Get_NextUpgradeDomain() string {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}
