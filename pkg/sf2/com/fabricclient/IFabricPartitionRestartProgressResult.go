package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// D2CB2EE1-A1BA-4CBD-80F7-14FD3D55BB61
var IID_IFabricPartitionRestartProgressResult = syscall.GUID{Data1: 0xD2CB2EE1, Data2: 0xA1BA, Data3: 0x4CBD,
	Data4: [8]byte{0x80, 0xF7, 0x14, 0xFD, 0x3D, 0x55, 0xBB, 0x61}}

type IFabricPartitionRestartProgressResult struct {
	win32.IUnknown
}

func NewIFabricPartitionRestartProgressResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPartitionRestartProgressResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPartitionRestartProgressResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPartitionRestartProgressResult) IID() *syscall.GUID {
	return &IID_IFabricPartitionRestartProgressResult
}

func (this *IFabricPartitionRestartProgressResult) Get_Progress() *FABRIC_PARTITION_RESTART_PROGRESS {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PARTITION_RESTART_PROGRESS)(unsafe.Pointer(ret))
}
