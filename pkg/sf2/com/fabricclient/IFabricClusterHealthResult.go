package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 6B9B0F2C-6782-4A31-A256-570FA8BA32D3
var IID_IFabricClusterHealthResult = syscall.GUID{Data1: 0x6B9B0F2C, Data2: 0x6782, Data3: 0x4A31,
	Data4: [8]byte{0xA2, 0x56, 0x57, 0x0F, 0xA8, 0xBA, 0x32, 0xD3}}

type IFabricClusterHealthResult struct {
	win32.IUnknown
}

func NewIFabricClusterHealthResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClusterHealthResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClusterHealthResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClusterHealthResult) IID() *syscall.GUID {
	return &IID_IFabricClusterHealthResult
}

func (this *IFabricClusterHealthResult) Get_ClusterHealth() *FABRIC_CLUSTER_HEALTH {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CLUSTER_HEALTH)(unsafe.Pointer(ret))
}
