package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// E00D3761-3AC5-407D-A04F-1B59486217CF
var IID_IFabricGetReplicaListResult = syscall.GUID{Data1: 0xE00D3761, Data2: 0x3AC5, Data3: 0x407D,
	Data4: [8]byte{0xA0, 0x4F, 0x1B, 0x59, 0x48, 0x62, 0x17, 0xCF}}

type IFabricGetReplicaListResult struct {
	win32.IUnknown
}

func NewIFabricGetReplicaListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetReplicaListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetReplicaListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetReplicaListResult) IID() *syscall.GUID {
	return &IID_IFabricGetReplicaListResult
}

func (this *IFabricGetReplicaListResult) Get_ReplicaList() *FABRIC_SERVICE_REPLICA_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_REPLICA_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
