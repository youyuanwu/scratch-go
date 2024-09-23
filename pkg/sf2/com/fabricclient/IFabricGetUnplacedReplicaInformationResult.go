package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 9D86A611-3FD3-451B-9495-6A831F417473
var IID_IFabricGetUnplacedReplicaInformationResult = syscall.GUID{Data1: 0x9D86A611, Data2: 0x3FD3, Data3: 0x451B,
	Data4: [8]byte{0x94, 0x95, 0x6A, 0x83, 0x1F, 0x41, 0x74, 0x73}}

type IFabricGetUnplacedReplicaInformationResult struct {
	win32.IUnknown
}

func NewIFabricGetUnplacedReplicaInformationResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetUnplacedReplicaInformationResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetUnplacedReplicaInformationResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetUnplacedReplicaInformationResult) IID() *syscall.GUID {
	return &IID_IFabricGetUnplacedReplicaInformationResult
}

func (this *IFabricGetUnplacedReplicaInformationResult) Get_UnplacedReplicaInformation() *FABRIC_UNPLACED_REPLICA_INFORMATION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_UNPLACED_REPLICA_INFORMATION)(unsafe.Pointer(ret))
}
