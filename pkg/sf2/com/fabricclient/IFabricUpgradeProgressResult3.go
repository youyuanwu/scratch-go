package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// DC3346EF-D2EF-40C1-807B-1CA8D2388B47
var IID_IFabricUpgradeProgressResult3 = syscall.GUID{Data1: 0xDC3346EF, Data2: 0xD2EF, Data3: 0x40C1,
	Data4: [8]byte{0x80, 0x7B, 0x1C, 0xA8, 0xD2, 0x38, 0x8B, 0x47}}

type IFabricUpgradeProgressResult3 struct {
	IFabricUpgradeProgressResult2
}

func NewIFabricUpgradeProgressResult3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricUpgradeProgressResult3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricUpgradeProgressResult3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricUpgradeProgressResult3) IID() *syscall.GUID {
	return &IID_IFabricUpgradeProgressResult3
}

func (this *IFabricUpgradeProgressResult3) Get_UpgradeProgress() *FABRIC_UPGRADE_PROGRESS {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_UPGRADE_PROGRESS)(unsafe.Pointer(ret))
}
