package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 46F1A40C-A4F3-409E-A7EC-6FD115F7ACC7
var IID_IFabricGetPartitionLoadInformationResult = syscall.GUID{Data1: 0x46F1A40C, Data2: 0xA4F3, Data3: 0x409E,
	Data4: [8]byte{0xA7, 0xEC, 0x6F, 0xD1, 0x15, 0xF7, 0xAC, 0xC7}}

type IFabricGetPartitionLoadInformationResult struct {
	win32.IUnknown
}

func NewIFabricGetPartitionLoadInformationResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetPartitionLoadInformationResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetPartitionLoadInformationResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetPartitionLoadInformationResult) IID() *syscall.GUID {
	return &IID_IFabricGetPartitionLoadInformationResult
}

func (this *IFabricGetPartitionLoadInformationResult) Get_PartitionLoadInformation() *FABRIC_PARTITION_LOAD_INFORMATION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_PARTITION_LOAD_INFORMATION)(unsafe.Pointer(ret))
}
