package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 7E9F51A5-88AC-49B8-958D-329E3334802E
var IID_IFabricStartNodeResult = syscall.GUID{Data1: 0x7E9F51A5, Data2: 0x88AC, Data3: 0x49B8,
	Data4: [8]byte{0x95, 0x8D, 0x32, 0x9E, 0x33, 0x34, 0x80, 0x2E}}

type IFabricStartNodeResult struct {
	win32.IUnknown
}

func NewIFabricStartNodeResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStartNodeResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStartNodeResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStartNodeResult) IID() *syscall.GUID {
	return &IID_IFabricStartNodeResult
}

func (this *IFabricStartNodeResult) Get_Result() *FABRIC_NODE_RESULT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NODE_RESULT)(unsafe.Pointer(ret))
}
