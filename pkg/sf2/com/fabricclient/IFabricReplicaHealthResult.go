package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// B4D5F2D9-E5CC-49AE-A6C8-89E8DF7B6C15
var IID_IFabricReplicaHealthResult = syscall.GUID{Data1: 0xB4D5F2D9, Data2: 0xE5CC, Data3: 0x49AE,
	Data4: [8]byte{0xA6, 0xC8, 0x89, 0xE8, 0xDF, 0x7B, 0x6C, 0x15}}

type IFabricReplicaHealthResult struct {
	win32.IUnknown
}

func NewIFabricReplicaHealthResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricReplicaHealthResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricReplicaHealthResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricReplicaHealthResult) IID() *syscall.GUID {
	return &IID_IFabricReplicaHealthResult
}

func (this *IFabricReplicaHealthResult) Get_ReplicaHealth() *FABRIC_REPLICA_HEALTH {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_REPLICA_HEALTH)(unsafe.Pointer(ret))
}
