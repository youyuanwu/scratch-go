package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 86B4F744-38C7-4DAB-B6B4-11C23734C269
var IID_IFabricServiceDescriptionResult = syscall.GUID{Data1: 0x86B4F744, Data2: 0x38C7, Data3: 0x4DAB,
	Data4: [8]byte{0xB6, 0xB4, 0x11, 0xC2, 0x37, 0x34, 0xC2, 0x69}}

type IFabricServiceDescriptionResult struct {
	win32.IUnknown
}

func NewIFabricServiceDescriptionResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceDescriptionResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceDescriptionResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceDescriptionResult) IID() *syscall.GUID {
	return &IID_IFabricServiceDescriptionResult
}

func (this *IFabricServiceDescriptionResult) Get_Description() *FABRIC_SERVICE_DESCRIPTION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_DESCRIPTION)(unsafe.Pointer(ret))
}
