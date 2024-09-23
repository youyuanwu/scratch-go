package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 4A0F2DA7-F851-44E5-8E12-AA521076097A
var IID_IFabricGetNodeListResult2 = syscall.GUID{Data1: 0x4A0F2DA7, Data2: 0xF851, Data3: 0x44E5,
	Data4: [8]byte{0x8E, 0x12, 0xAA, 0x52, 0x10, 0x76, 0x09, 0x7A}}

type IFabricGetNodeListResult2 struct {
	IFabricGetNodeListResult
}

func NewIFabricGetNodeListResult2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetNodeListResult2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetNodeListResult2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetNodeListResult2) IID() *syscall.GUID {
	return &IID_IFabricGetNodeListResult2
}

func (this *IFabricGetNodeListResult2) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
