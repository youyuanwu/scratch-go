package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 7FEFCF06-C840-4D8A-9CC7-36F080E0E121
var IID_IFabricGetClusterHealthChunkResult = syscall.GUID{Data1: 0x7FEFCF06, Data2: 0xC840, Data3: 0x4D8A,
	Data4: [8]byte{0x9C, 0xC7, 0x36, 0xF0, 0x80, 0xE0, 0xE1, 0x21}}

type IFabricGetClusterHealthChunkResult struct {
	win32.IUnknown
}

func NewIFabricGetClusterHealthChunkResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetClusterHealthChunkResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetClusterHealthChunkResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetClusterHealthChunkResult) IID() *syscall.GUID {
	return &IID_IFabricGetClusterHealthChunkResult
}

func (this *IFabricGetClusterHealthChunkResult) Get_ClusterHealthChunk() *FABRIC_CLUSTER_HEALTH_CHUNK {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CLUSTER_HEALTH_CHUNK)(unsafe.Pointer(ret))
}
