package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 02BD6674-9C5A-4262-89A8-AC1A6A2FB5E9
var IID_IFabricGetRollingUpgradeMonitoringPolicyResult = syscall.GUID{Data1: 0x02BD6674, Data2: 0x9C5A, Data3: 0x4262,
	Data4: [8]byte{0x89, 0xA8, 0xAC, 0x1A, 0x6A, 0x2F, 0xB5, 0xE9}}

type IFabricGetRollingUpgradeMonitoringPolicyResult struct {
	win32.IUnknown
}

func NewIFabricGetRollingUpgradeMonitoringPolicyResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetRollingUpgradeMonitoringPolicyResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetRollingUpgradeMonitoringPolicyResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetRollingUpgradeMonitoringPolicyResult) IID() *syscall.GUID {
	return &IID_IFabricGetRollingUpgradeMonitoringPolicyResult
}

func (this *IFabricGetRollingUpgradeMonitoringPolicyResult) Get_Policy() *FABRIC_ROLLING_UPGRADE_MONITORING_POLICY {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_ROLLING_UPGRADE_MONITORING_POLICY)(unsafe.Pointer(ret))
}
