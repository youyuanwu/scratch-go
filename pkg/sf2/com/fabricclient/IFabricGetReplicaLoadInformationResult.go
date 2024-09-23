package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// E4190CA0-225C-11E4-8C21-0800200C9A66
var IID_IFabricGetReplicaLoadInformationResult = syscall.GUID{Data1: 0xE4190CA0, Data2: 0x225C, Data3: 0x11E4,
	Data4: [8]byte{0x8C, 0x21, 0x08, 0x00, 0x20, 0x0C, 0x9A, 0x66}}

type IFabricGetReplicaLoadInformationResult struct {
	win32.IUnknown
}

func NewIFabricGetReplicaLoadInformationResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetReplicaLoadInformationResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetReplicaLoadInformationResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetReplicaLoadInformationResult) IID() *syscall.GUID {
	return &IID_IFabricGetReplicaLoadInformationResult
}

func (this *IFabricGetReplicaLoadInformationResult) Get_ReplicaLoadInformation() *FABRIC_REPLICA_LOAD_INFORMATION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_REPLICA_LOAD_INFORMATION)(unsafe.Pointer(ret))
}
