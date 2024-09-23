package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 711D60A6-9623-476C-970C-83059A0B4D55
var IID_IFabricStopNodeResult = syscall.GUID{Data1: 0x711D60A6, Data2: 0x9623, Data3: 0x476C,
	Data4: [8]byte{0x97, 0x0C, 0x83, 0x05, 0x9A, 0x0B, 0x4D, 0x55}}

type IFabricStopNodeResult struct {
	win32.IUnknown
}

func NewIFabricStopNodeResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStopNodeResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStopNodeResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStopNodeResult) IID() *syscall.GUID {
	return &IID_IFabricStopNodeResult
}

func (this *IFabricStopNodeResult) Get_Result() *FABRIC_NODE_RESULT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NODE_RESULT)(unsafe.Pointer(ret))
}
