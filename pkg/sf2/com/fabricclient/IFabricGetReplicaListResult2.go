package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 0BC12F86-C157-4C0D-B274-01FB09145934
var IID_IFabricGetReplicaListResult2 = syscall.GUID{Data1: 0x0BC12F86, Data2: 0xC157, Data3: 0x4C0D,
	Data4: [8]byte{0xB2, 0x74, 0x01, 0xFB, 0x09, 0x14, 0x59, 0x34}}

type IFabricGetReplicaListResult2 struct {
	IFabricGetReplicaListResult
}

func NewIFabricGetReplicaListResult2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetReplicaListResult2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetReplicaListResult2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetReplicaListResult2) IID() *syscall.GUID {
	return &IID_IFabricGetReplicaListResult2
}

func (this *IFabricGetReplicaListResult2) Get_PagingStatus() *FABRIC_PAGING_STATUS {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PAGING_STATUS)(unsafe.Pointer(ret))
}
