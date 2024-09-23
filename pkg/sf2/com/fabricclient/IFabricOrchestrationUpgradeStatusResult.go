package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 95A56E4A-490D-445E-865C-EF0A62F15504
var IID_IFabricOrchestrationUpgradeStatusResult = syscall.GUID{Data1: 0x95A56E4A, Data2: 0x490D, Data3: 0x445E,
	Data4: [8]byte{0x86, 0x5C, 0xEF, 0x0A, 0x62, 0xF1, 0x55, 0x04}}

type IFabricOrchestrationUpgradeStatusResult struct {
	win32.IUnknown
}

func NewIFabricOrchestrationUpgradeStatusResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricOrchestrationUpgradeStatusResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricOrchestrationUpgradeStatusResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricOrchestrationUpgradeStatusResult) IID() *syscall.GUID {
	return &IID_IFabricOrchestrationUpgradeStatusResult
}

func (this *IFabricOrchestrationUpgradeStatusResult) Get_Progress() *FABRIC_ORCHESTRATION_UPGRADE_PROGRESS {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_ORCHESTRATION_UPGRADE_PROGRESS)(unsafe.Pointer(ret))
}
