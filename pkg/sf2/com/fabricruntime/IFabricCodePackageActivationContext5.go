package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// FE45387E-8711-4949-AC36-31DC95035513
var IID_IFabricCodePackageActivationContext5 = syscall.GUID{Data1: 0xFE45387E, Data2: 0x8711, Data3: 0x4949,
	Data4: [8]byte{0xAC, 0x36, 0x31, 0xDC, 0x95, 0x03, 0x55, 0x13}}

type IFabricCodePackageActivationContext5 struct {
	IFabricCodePackageActivationContext4
}

func NewIFabricCodePackageActivationContext5(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricCodePackageActivationContext5 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricCodePackageActivationContext5)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricCodePackageActivationContext5) IID() *syscall.GUID {
	return &IID_IFabricCodePackageActivationContext5
}

func (this *IFabricCodePackageActivationContext5) Get_ServiceListenAddress() string {
	addr := (*this.LpVtbl)[36]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricCodePackageActivationContext5) Get_ServicePublishAddress() string {
	addr := (*this.LpVtbl)[37]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}
