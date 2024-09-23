package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 9953E19A-EA1E-4A1F-BDA4-AB42FDB77185
var IID_IFabricGetServiceListResult = syscall.GUID{Data1: 0x9953E19A, Data2: 0xEA1E, Data3: 0x4A1F,
	Data4: [8]byte{0xBD, 0xA4, 0xAB, 0x42, 0xFD, 0xB7, 0x71, 0x85}}

type IFabricGetServiceListResult struct {
	win32.IUnknown
}

func NewIFabricGetServiceListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetServiceListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetServiceListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetServiceListResult) IID() *syscall.GUID {
	return &IID_IFabricGetServiceListResult
}

func (this *IFabricGetServiceListResult) Get_ServiceList() *FABRIC_SERVICE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
