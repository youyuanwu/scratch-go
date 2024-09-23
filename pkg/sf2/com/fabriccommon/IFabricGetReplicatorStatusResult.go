package fabriccommon

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 30E10C61-A710-4F99-A623-BB1403265186
var IID_IFabricGetReplicatorStatusResult = syscall.GUID{Data1: 0x30E10C61, Data2: 0xA710, Data3: 0x4F99,
	Data4: [8]byte{0xA6, 0x23, 0xBB, 0x14, 0x03, 0x26, 0x51, 0x86}}

type IFabricGetReplicatorStatusResult struct {
	win32.IUnknown
}

func NewIFabricGetReplicatorStatusResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetReplicatorStatusResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetReplicatorStatusResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetReplicatorStatusResult) IID() *syscall.GUID {
	return &IID_IFabricGetReplicatorStatusResult
}

func (this *IFabricGetReplicatorStatusResult) Get_ReplicatorStatus() *FABRIC_REPLICATOR_STATUS_QUERY_RESULT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_REPLICATOR_STATUS_QUERY_RESULT)(unsafe.Pointer(ret))
}
