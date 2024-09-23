package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 3B825AFD-CB31-4589-961E-E3778AA23A60
var IID_IFabricClientSettingsResult = syscall.GUID{Data1: 0x3B825AFD, Data2: 0xCB31, Data3: 0x4589,
	Data4: [8]byte{0x96, 0x1E, 0xE3, 0x77, 0x8A, 0xA2, 0x3A, 0x60}}

type IFabricClientSettingsResult struct {
	win32.IUnknown
}

func NewIFabricClientSettingsResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClientSettingsResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClientSettingsResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClientSettingsResult) IID() *syscall.GUID {
	return &IID_IFabricClientSettingsResult
}

func (this *IFabricClientSettingsResult) Get_Settings() *FABRIC_CLIENT_SETTINGS {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CLIENT_SETTINGS)(unsafe.Pointer(ret))
}
