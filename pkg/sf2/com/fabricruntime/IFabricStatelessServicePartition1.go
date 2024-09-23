package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// BF6BB505-7BD0-4371-B6C0-CBA319A5E50B
var IID_IFabricStatelessServicePartition1 = syscall.GUID{Data1: 0xBF6BB505, Data2: 0x7BD0, Data3: 0x4371,
	Data4: [8]byte{0xB6, 0xC0, 0xCB, 0xA3, 0x19, 0xA5, 0xE5, 0x0B}}

type IFabricStatelessServicePartition1 struct {
	IFabricStatelessServicePartition
}

func NewIFabricStatelessServicePartition1(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatelessServicePartition1 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatelessServicePartition1)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatelessServicePartition1) IID() *syscall.GUID {
	return &IID_IFabricStatelessServicePartition1
}

func (this *IFabricStatelessServicePartition1) ReportMoveCost(moveCost int32) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(moveCost))
	return com.Error(ret)
}
