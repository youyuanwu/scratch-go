package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 614921E6-75F1-44E7-9107-AB88819136B8
var IID_IFabricPartitionDataLossProgressResult = syscall.GUID{Data1: 0x614921E6, Data2: 0x75F1, Data3: 0x44E7,
	Data4: [8]byte{0x91, 0x07, 0xAB, 0x88, 0x81, 0x91, 0x36, 0xB8}}

type IFabricPartitionDataLossProgressResult struct {
	win32.IUnknown
}

func NewIFabricPartitionDataLossProgressResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPartitionDataLossProgressResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPartitionDataLossProgressResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPartitionDataLossProgressResult) IID() *syscall.GUID {
	return &IID_IFabricPartitionDataLossProgressResult
}

func (this *IFabricPartitionDataLossProgressResult) Get_Progress() *FABRIC_PARTITION_DATA_LOSS_PROGRESS {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PARTITION_DATA_LOSS_PROGRESS)(unsafe.Pointer(ret))
}
