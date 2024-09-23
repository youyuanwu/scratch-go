package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// B131B99A-D251-47B2-9D08-24DDD6793206
var IID_IFabricGetPartitionListResult2 = syscall.GUID{Data1: 0xB131B99A, Data2: 0xD251, Data3: 0x47B2,
	Data4: [8]byte{0x9D, 0x08, 0x24, 0xDD, 0xD6, 0x79, 0x32, 0x06}}

type IFabricGetPartitionListResult2 struct {
	IFabricGetPartitionListResult
}

func NewIFabricGetPartitionListResult2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetPartitionListResult2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetPartitionListResult2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetPartitionListResult2) IID() *syscall.GUID {
	return &IID_IFabricGetPartitionListResult2
}

func (this *IFabricGetPartitionListResult2) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
