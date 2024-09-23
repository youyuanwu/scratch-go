package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 32D656A1-7AD5-47B8-BD66-A2E302626B7E
var IID_IFabricTransactionBase = syscall.GUID{Data1: 0x32D656A1, Data2: 0x7AD5, Data3: 0x47B8,
	Data4: [8]byte{0xBD, 0x66, 0xA2, 0xE3, 0x02, 0x62, 0x6B, 0x7E}}

type IFabricTransactionBase struct {
	win32.IUnknown
}

func NewIFabricTransactionBase(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricTransactionBase {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricTransactionBase)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricTransactionBase) IID() *syscall.GUID {
	return &IID_IFabricTransactionBase
}

func (this *IFabricTransactionBase) Get_Id() *syscall.GUID {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*syscall.GUID)(unsafe.Pointer(ret))
}

func (this *IFabricTransactionBase) Get_IsolationLevel() int32 {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int32(ret)
}
