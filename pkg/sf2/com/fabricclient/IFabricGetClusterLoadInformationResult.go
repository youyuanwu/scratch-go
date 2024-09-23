package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 7CC3EB08-0E69-4E52-81FC-0190AB997DBF
var IID_IFabricGetClusterLoadInformationResult = syscall.GUID{Data1: 0x7CC3EB08, Data2: 0x0E69, Data3: 0x4E52,
	Data4: [8]byte{0x81, 0xFC, 0x01, 0x90, 0xAB, 0x99, 0x7D, 0xBF}}

type IFabricGetClusterLoadInformationResult struct {
	win32.IUnknown
}

func NewIFabricGetClusterLoadInformationResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetClusterLoadInformationResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetClusterLoadInformationResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetClusterLoadInformationResult) IID() *syscall.GUID {
	return &IID_IFabricGetClusterLoadInformationResult
}

func (this *IFabricGetClusterLoadInformationResult) Get_ClusterLoadInformation() *FABRIC_CLUSTER_LOAD_INFORMATION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CLUSTER_LOAD_INFORMATION)(unsafe.Pointer(ret))
}
