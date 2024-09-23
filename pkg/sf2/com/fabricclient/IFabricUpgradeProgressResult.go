package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 2ADB07DB-F7DB-4621-9AFC-DAABE1E53BF8
var IID_IFabricUpgradeProgressResult = syscall.GUID{Data1: 0x2ADB07DB, Data2: 0xF7DB, Data3: 0x4621,
	Data4: [8]byte{0x9A, 0xFC, 0xDA, 0xAB, 0xE1, 0xE5, 0x3B, 0xF8}}

type IFabricUpgradeProgressResult struct {
	win32.IUnknown
}

func NewIFabricUpgradeProgressResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricUpgradeProgressResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricUpgradeProgressResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricUpgradeProgressResult) IID() *syscall.GUID {
	return &IID_IFabricUpgradeProgressResult
}

func (this *IFabricUpgradeProgressResult) Get_TargetCodeVersion() string {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricUpgradeProgressResult) Get_TargetConfigVersion() string {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricUpgradeProgressResult) Get_UpgradeState() int32 {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int32(ret)
}

func (this *IFabricUpgradeProgressResult) GetUpgradeDomains(itemCount *uint32, bufferedItems **FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(itemCount)), uintptr(unsafe.Pointer(bufferedItems)))
	return com.Error(ret)
}

func (this *IFabricUpgradeProgressResult) GetChangedUpgradeDomains(previousProgress *IFabricUpgradeProgressResult, itemCount *uint32, bufferedItems **FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(previousProgress)), uintptr(unsafe.Pointer(itemCount)), uintptr(unsafe.Pointer(bufferedItems)))
	return com.Error(ret)
}
