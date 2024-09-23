package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F038C61E-7059-41B6-8DEA-D304A2080F46
var IID_IFabricGetApplicationListResult = syscall.GUID{Data1: 0xF038C61E, Data2: 0x7059, Data3: 0x41B6,
	Data4: [8]byte{0x8D, 0xEA, 0xD3, 0x04, 0xA2, 0x08, 0x0F, 0x46}}

type IFabricGetApplicationListResult struct {
	win32.IUnknown
}

func NewIFabricGetApplicationListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetApplicationListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetApplicationListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetApplicationListResult) IID() *syscall.GUID {
	return &IID_IFabricGetApplicationListResult
}

func (this *IFabricGetApplicationListResult) Get_ApplicationList() *FABRIC_APPLICATION_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_APPLICATION_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
