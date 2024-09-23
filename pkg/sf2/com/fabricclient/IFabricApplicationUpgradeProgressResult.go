package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 1E4670F8-EDE5-48AB-881F-C45A0F38413A
var IID_IFabricApplicationUpgradeProgressResult = syscall.GUID{Data1: 0x1E4670F8, Data2: 0xEDE5, Data3: 0x48AB,
	Data4: [8]byte{0x88, 0x1F, 0xC4, 0x5A, 0x0F, 0x38, 0x41, 0x3A}}

type IFabricApplicationUpgradeProgressResult struct {
	win32.IUnknown
}

func NewIFabricApplicationUpgradeProgressResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationUpgradeProgressResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationUpgradeProgressResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationUpgradeProgressResult) IID() *syscall.GUID {
	return &IID_IFabricApplicationUpgradeProgressResult
}

func (this *IFabricApplicationUpgradeProgressResult) Get_ApplicationName() string {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricApplicationUpgradeProgressResult) Get_ApplicationTypeName() string {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricApplicationUpgradeProgressResult) Get_TargetApplicationTypeVersion() string {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricApplicationUpgradeProgressResult) Get_UpgradeState() int32 {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int32(ret)
}

func (this *IFabricApplicationUpgradeProgressResult) GetUpgradeDomains(itemCount *uint32, bufferedItems **FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(itemCount)), uintptr(unsafe.Pointer(bufferedItems)))
	return com.Error(ret)
}

func (this *IFabricApplicationUpgradeProgressResult) GetChangedUpgradeDomains(previousProgress *IFabricApplicationUpgradeProgressResult, itemCount *uint32, bufferedItems **FABRIC_UPGRADE_DOMAIN_STATUS_DESCRIPTION) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(previousProgress)), uintptr(unsafe.Pointer(itemCount)), uintptr(unsafe.Pointer(bufferedItems)))
	return com.Error(ret)
}
