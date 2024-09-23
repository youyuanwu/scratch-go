package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 10C9E99D-BB3F-4263-A7F7-ABBAF3C03576
var IID_IFabricPartitionHealthResult = syscall.GUID{Data1: 0x10C9E99D, Data2: 0xBB3F, Data3: 0x4263,
	Data4: [8]byte{0xA7, 0xF7, 0xAB, 0xBA, 0xF3, 0xC0, 0x35, 0x76}}

type IFabricPartitionHealthResult struct {
	win32.IUnknown
}

func NewIFabricPartitionHealthResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPartitionHealthResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPartitionHealthResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPartitionHealthResult) IID() *syscall.GUID {
	return &IID_IFabricPartitionHealthResult
}

func (this *IFabricPartitionHealthResult) Get_PartitionHealth() *FABRIC_PARTITION_HEALTH {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PARTITION_HEALTH)(unsafe.Pointer(ret))
}
