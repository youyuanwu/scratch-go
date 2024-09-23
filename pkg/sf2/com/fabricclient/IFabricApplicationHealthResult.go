package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 41612FAB-E615-4A48-98E7-4ABCC93B6049
var IID_IFabricApplicationHealthResult = syscall.GUID{Data1: 0x41612FAB, Data2: 0xE615, Data3: 0x4A48,
	Data4: [8]byte{0x98, 0xE7, 0x4A, 0xBC, 0xC9, 0x3B, 0x60, 0x49}}

type IFabricApplicationHealthResult struct {
	win32.IUnknown
}

func NewIFabricApplicationHealthResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationHealthResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationHealthResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationHealthResult) IID() *syscall.GUID {
	return &IID_IFabricApplicationHealthResult
}

func (this *IFabricApplicationHealthResult) Get_ApplicationHealth() *FABRIC_APPLICATION_HEALTH {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_APPLICATION_HEALTH)(unsafe.Pointer(ret))
}
