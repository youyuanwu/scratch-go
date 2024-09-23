package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 36D8E378-3706-403D-8D99-2AFD1A120687
var IID_IFabricPartitionQuorumLossProgressResult = syscall.GUID{Data1: 0x36D8E378, Data2: 0x3706, Data3: 0x403D,
	Data4: [8]byte{0x8D, 0x99, 0x2A, 0xFD, 0x1A, 0x12, 0x06, 0x87}}

type IFabricPartitionQuorumLossProgressResult struct {
	win32.IUnknown
}

func NewIFabricPartitionQuorumLossProgressResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPartitionQuorumLossProgressResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPartitionQuorumLossProgressResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPartitionQuorumLossProgressResult) IID() *syscall.GUID {
	return &IID_IFabricPartitionQuorumLossProgressResult
}

func (this *IFabricPartitionQuorumLossProgressResult) Get_Progress() *FABRIC_PARTITION_QUORUM_LOSS_PROGRESS {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PARTITION_QUORUM_LOSS_PROGRESS)(unsafe.Pointer(ret))
}
