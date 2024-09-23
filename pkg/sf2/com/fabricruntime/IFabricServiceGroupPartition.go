package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 2B24299A-7489-467F-8E7F-4507BFF73B86
var IID_IFabricServiceGroupPartition = syscall.GUID{Data1: 0x2B24299A, Data2: 0x7489, Data3: 0x467F,
	Data4: [8]byte{0x8E, 0x7F, 0x45, 0x07, 0xBF, 0xF7, 0x3B, 0x86}}

type IFabricServiceGroupPartition struct {
	win32.IUnknown
}

func NewIFabricServiceGroupPartition(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceGroupPartition {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceGroupPartition)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceGroupPartition) IID() *syscall.GUID {
	return &IID_IFabricServiceGroupPartition
}

func (this *IFabricServiceGroupPartition) ResolveMember(name string, riid *syscall.GUID, member unsafe.Pointer) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(riid)), uintptr(member))
	return com.Error(ret)
}
