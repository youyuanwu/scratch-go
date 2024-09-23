package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// FA8AA86E-F0FA-4A14-BED7-1DCFA0980B5B
var IID_IFabricChaosDescriptionResult = syscall.GUID{Data1: 0xFA8AA86E, Data2: 0xF0FA, Data3: 0x4A14,
	Data4: [8]byte{0xBE, 0xD7, 0x1D, 0xCF, 0xA0, 0x98, 0x0B, 0x5B}}

type IFabricChaosDescriptionResult struct {
	win32.IUnknown
}

func NewIFabricChaosDescriptionResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricChaosDescriptionResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricChaosDescriptionResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricChaosDescriptionResult) IID() *syscall.GUID {
	return &IID_IFabricChaosDescriptionResult
}

func (this *IFabricChaosDescriptionResult) Get_ChaosDescriptionResult() *FABRIC_CHAOS_DESCRIPTION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CHAOS_DESCRIPTION)(unsafe.Pointer(ret))
}
