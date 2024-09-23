package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// AFC1266C-967B-4769-9F8A-B249C5887EE6
var IID_IFabricGetPartitionListResult = syscall.GUID{Data1: 0xAFC1266C, Data2: 0x967B, Data3: 0x4769,
	Data4: [8]byte{0x9F, 0x8A, 0xB2, 0x49, 0xC5, 0x88, 0x7E, 0xE6}}

type IFabricGetPartitionListResult struct {
	win32.IUnknown
}

func NewIFabricGetPartitionListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetPartitionListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetPartitionListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetPartitionListResult) IID() *syscall.GUID {
	return &IID_IFabricGetPartitionListResult
}

func (this *IFabricGetPartitionListResult) Get_PartitionList() *FABRIC_SERVICE_PARTITION_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_PARTITION_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
