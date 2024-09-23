package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 718954F3-DC1E-4060-9806-0CBF36F71051
var IID_IFabricReplicatorSettingsResult = syscall.GUID{Data1: 0x718954F3, Data2: 0xDC1E, Data3: 0x4060,
	Data4: [8]byte{0x98, 0x06, 0x0C, 0xBF, 0x36, 0xF7, 0x10, 0x51}}

type IFabricReplicatorSettingsResult struct {
	win32.IUnknown
}

func NewIFabricReplicatorSettingsResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricReplicatorSettingsResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricReplicatorSettingsResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricReplicatorSettingsResult) IID() *syscall.GUID {
	return &IID_IFabricReplicatorSettingsResult
}

func (this *IFabricReplicatorSettingsResult) Get_ReplicatorSettings() *FABRIC_REPLICATOR_SETTINGS {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_REPLICATOR_SETTINGS)(unsafe.Pointer(ret))
}
