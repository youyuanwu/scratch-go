package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// AA67DE09-3657-435F-A2F6-B3A17A0A4371
var IID_IFabricDataPackage = syscall.GUID{Data1: 0xAA67DE09, Data2: 0x3657, Data3: 0x435F,
	Data4: [8]byte{0xA2, 0xF6, 0xB3, 0xA1, 0x7A, 0x0A, 0x43, 0x71}}

type IFabricDataPackage struct {
	win32.IUnknown
}

func NewIFabricDataPackage(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricDataPackage {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricDataPackage)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricDataPackage) IID() *syscall.GUID {
	return &IID_IFabricDataPackage
}

func (this *IFabricDataPackage) Get_Description() *FABRIC_DATA_PACKAGE_DESCRIPTION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DATA_PACKAGE_DESCRIPTION)(unsafe.Pointer(ret))
}

func (this *IFabricDataPackage) Get_Path() string {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}
