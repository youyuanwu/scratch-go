package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 3CA814D4-E067-48B7-9BDC-9BE33810416D
var IID_IFabricServiceGroupDescriptionResult = syscall.GUID{Data1: 0x3CA814D4, Data2: 0xE067, Data3: 0x48B7,
	Data4: [8]byte{0x9B, 0xDC, 0x9B, 0xE3, 0x38, 0x10, 0x41, 0x6D}}

type IFabricServiceGroupDescriptionResult struct {
	win32.IUnknown
}

func NewIFabricServiceGroupDescriptionResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceGroupDescriptionResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceGroupDescriptionResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceGroupDescriptionResult) IID() *syscall.GUID {
	return &IID_IFabricServiceGroupDescriptionResult
}

func (this *IFabricServiceGroupDescriptionResult) Get_Description() *FABRIC_SERVICE_GROUP_DESCRIPTION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_GROUP_DESCRIPTION)(unsafe.Pointer(ret))
}
