package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 886E4AD2-EDB8-4734-9DD4-0E9A2BE5238B
var IID_IFabricGetServiceTypeListResult = syscall.GUID{Data1: 0x886E4AD2, Data2: 0xEDB8, Data3: 0x4734,
	Data4: [8]byte{0x9D, 0xD4, 0x0E, 0x9A, 0x2B, 0xE5, 0x23, 0x8B}}

type IFabricGetServiceTypeListResult struct {
	win32.IUnknown
}

func NewIFabricGetServiceTypeListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetServiceTypeListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetServiceTypeListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetServiceTypeListResult) IID() *syscall.GUID {
	return &IID_IFabricGetServiceTypeListResult
}

func (this *IFabricGetServiceTypeListResult) Get_ServiceTypeList() *FABRIC_SERVICE_TYPE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_TYPE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
