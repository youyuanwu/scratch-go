package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// E461F70B-51B8-4B73-9F35-E38E5AC68719
var IID_IFabricNodeHealthResult = syscall.GUID{Data1: 0xE461F70B, Data2: 0x51B8, Data3: 0x4B73,
	Data4: [8]byte{0x9F, 0x35, 0xE3, 0x8E, 0x5A, 0xC6, 0x87, 0x19}}

type IFabricNodeHealthResult struct {
	win32.IUnknown
}

func NewIFabricNodeHealthResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricNodeHealthResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricNodeHealthResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricNodeHealthResult) IID() *syscall.GUID {
	return &IID_IFabricNodeHealthResult
}

func (this *IFabricNodeHealthResult) Get_NodeHealth() *FABRIC_NODE_HEALTH {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NODE_HEALTH)(unsafe.Pointer(ret))
}
