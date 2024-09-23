package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// C9C66F2F-9DFF-4C87-BBE4-A08B4C4074CF
var IID_IFabricStatefulServicePartition1 = syscall.GUID{Data1: 0xC9C66F2F, Data2: 0x9DFF, Data3: 0x4C87,
	Data4: [8]byte{0xBB, 0xE4, 0xA0, 0x8B, 0x4C, 0x40, 0x74, 0xCF}}

type IFabricStatefulServicePartition1 struct {
	IFabricStatefulServicePartition
}

func NewIFabricStatefulServicePartition1(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatefulServicePartition1 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatefulServicePartition1)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatefulServicePartition1) IID() *syscall.GUID {
	return &IID_IFabricStatefulServicePartition1
}

func (this *IFabricStatefulServicePartition1) ReportMoveCost(moveCost int32) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(moveCost))
	return com.Error(ret)
}
