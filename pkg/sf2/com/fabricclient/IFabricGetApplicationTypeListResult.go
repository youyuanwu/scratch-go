package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 944F7A70-224E-4191-8DD1-BBA46DC88DD2
var IID_IFabricGetApplicationTypeListResult = syscall.GUID{Data1: 0x944F7A70, Data2: 0x224E, Data3: 0x4191,
	Data4: [8]byte{0x8D, 0xD1, 0xBB, 0xA4, 0x6D, 0xC8, 0x8D, 0xD2}}

type IFabricGetApplicationTypeListResult struct {
	win32.IUnknown
}

func NewIFabricGetApplicationTypeListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetApplicationTypeListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetApplicationTypeListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetApplicationTypeListResult) IID() *syscall.GUID {
	return &IID_IFabricGetApplicationTypeListResult
}

func (this *IFabricGetApplicationTypeListResult) Get_ApplicationTypeList() *FABRIC_APPLICATION_TYPE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_APPLICATION_TYPE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
