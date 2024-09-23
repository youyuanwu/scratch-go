package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 29E064BF-5D78-49E5-BAA6-ACFC24A4A8B5
var IID_IFabricGetDeployedReplicaListResult = syscall.GUID{Data1: 0x29E064BF, Data2: 0x5D78, Data3: 0x49E5,
	Data4: [8]byte{0xBA, 0xA6, 0xAC, 0xFC, 0x24, 0xA4, 0xA8, 0xB5}}

type IFabricGetDeployedReplicaListResult struct {
	win32.IUnknown
}

func NewIFabricGetDeployedReplicaListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetDeployedReplicaListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetDeployedReplicaListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetDeployedReplicaListResult) IID() *syscall.GUID {
	return &IID_IFabricGetDeployedReplicaListResult
}

func (this *IFabricGetDeployedReplicaListResult) Get_DeployedReplicaList() *FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_SERVICE_REPLICA_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
