package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 6D9D355E-89CF-4928-B758-B11CA4664FBE
var IID_IFabricGetDeployedServiceReplicaDetailResult = syscall.GUID{Data1: 0x6D9D355E, Data2: 0x89CF, Data3: 0x4928,
	Data4: [8]byte{0xB7, 0x58, 0xB1, 0x1C, 0xA4, 0x66, 0x4F, 0xBE}}

type IFabricGetDeployedServiceReplicaDetailResult struct {
	win32.IUnknown
}

func NewIFabricGetDeployedServiceReplicaDetailResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetDeployedServiceReplicaDetailResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetDeployedServiceReplicaDetailResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetDeployedServiceReplicaDetailResult) IID() *syscall.GUID {
	return &IID_IFabricGetDeployedServiceReplicaDetailResult
}

func (this *IFabricGetDeployedServiceReplicaDetailResult) Get_ReplicaDetail() *FABRIC_DEPLOYED_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_SERVICE_REPLICA_DETAIL_QUERY_RESULT_ITEM)(unsafe.Pointer(ret))
}
