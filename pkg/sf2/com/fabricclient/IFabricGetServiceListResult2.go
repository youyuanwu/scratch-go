package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 30263683-4B25-4EC3-86D7-94ED86E7A8BF
var IID_IFabricGetServiceListResult2 = syscall.GUID{Data1: 0x30263683, Data2: 0x4B25, Data3: 0x4EC3,
	Data4: [8]byte{0x86, 0xD7, 0x94, 0xED, 0x86, 0xE7, 0xA8, 0xBF}}

type IFabricGetServiceListResult2 struct {
	IFabricGetServiceListResult
}

func NewIFabricGetServiceListResult2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetServiceListResult2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetServiceListResult2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetServiceListResult2) IID() *syscall.GUID {
	return &IID_IFabricGetServiceListResult2
}

func (this *IFabricGetServiceListResult2) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
