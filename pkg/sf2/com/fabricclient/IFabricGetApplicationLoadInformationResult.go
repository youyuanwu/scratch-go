package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 38FD0512-7586-4BD5-9616-B7070CF025C0
var IID_IFabricGetApplicationLoadInformationResult = syscall.GUID{Data1: 0x38FD0512, Data2: 0x7586, Data3: 0x4BD5,
	Data4: [8]byte{0x96, 0x16, 0xB7, 0x07, 0x0C, 0xF0, 0x25, 0xC0}}

type IFabricGetApplicationLoadInformationResult struct {
	win32.IUnknown
}

func NewIFabricGetApplicationLoadInformationResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetApplicationLoadInformationResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetApplicationLoadInformationResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetApplicationLoadInformationResult) IID() *syscall.GUID {
	return &IID_IFabricGetApplicationLoadInformationResult
}

func (this *IFabricGetApplicationLoadInformationResult) Get_ApplicationLoadInformation() *FABRIC_APPLICATION_LOAD_INFORMATION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_APPLICATION_LOAD_INFORMATION)(unsafe.Pointer(ret))
}
