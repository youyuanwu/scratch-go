package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 5D8DDE9C-05E8-428D-B494-43873D7C2DB8
var IID_IFabricGetApplicationTypePagedListResult = syscall.GUID{Data1: 0x5D8DDE9C, Data2: 0x05E8, Data3: 0x428D,
	Data4: [8]byte{0xB4, 0x94, 0x43, 0x87, 0x3D, 0x7C, 0x2D, 0xB8}}

type IFabricGetApplicationTypePagedListResult struct {
	win32.IUnknown
}

func NewIFabricGetApplicationTypePagedListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetApplicationTypePagedListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetApplicationTypePagedListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetApplicationTypePagedListResult) IID() *syscall.GUID {
	return &IID_IFabricGetApplicationTypePagedListResult
}

func (this *IFabricGetApplicationTypePagedListResult) Get_ApplicationTypePagedList() *FABRIC_APPLICATION_TYPE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_APPLICATION_TYPE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricGetApplicationTypePagedListResult) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
