package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 20792B45-4D13-41A4-AF13-346E529F00C5
var IID_IFabricCodePackage = syscall.GUID{Data1: 0x20792B45, Data2: 0x4D13, Data3: 0x41A4,
	Data4: [8]byte{0xAF, 0x13, 0x34, 0x6E, 0x52, 0x9F, 0x00, 0xC5}}

type IFabricCodePackage struct {
	win32.IUnknown
}

func NewIFabricCodePackage(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricCodePackage {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricCodePackage)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricCodePackage) IID() *syscall.GUID {
	return &IID_IFabricCodePackage
}

func (this *IFabricCodePackage) Get_Description() *FABRIC_CODE_PACKAGE_DESCRIPTION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CODE_PACKAGE_DESCRIPTION)(unsafe.Pointer(ret))
}

func (this *IFabricCodePackage) Get_Path() string {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}
