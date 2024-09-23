package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 6637A860-26BC-4F1A-902F-F418FCFE1E51
var IID_IFabricGetApplicationListResult2 = syscall.GUID{Data1: 0x6637A860, Data2: 0x26BC, Data3: 0x4F1A,
	Data4: [8]byte{0x90, 0x2F, 0xF4, 0x18, 0xFC, 0xFE, 0x1E, 0x51}}

type IFabricGetApplicationListResult2 struct {
	IFabricGetApplicationListResult
}

func NewIFabricGetApplicationListResult2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetApplicationListResult2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetApplicationListResult2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetApplicationListResult2) IID() *syscall.GUID {
	return &IID_IFabricGetApplicationListResult2
}

func (this *IFabricGetApplicationListResult2) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
