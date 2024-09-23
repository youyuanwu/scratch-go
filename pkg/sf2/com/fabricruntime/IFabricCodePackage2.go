package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// CDF0A4E6-AD80-4CD6-B67E-E4C002428600
var IID_IFabricCodePackage2 = syscall.GUID{Data1: 0xCDF0A4E6, Data2: 0xAD80, Data3: 0x4CD6,
	Data4: [8]byte{0xB6, 0x7E, 0xE4, 0xC0, 0x02, 0x42, 0x86, 0x00}}

type IFabricCodePackage2 struct {
	IFabricCodePackage
}

func NewIFabricCodePackage2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricCodePackage2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricCodePackage2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricCodePackage2) IID() *syscall.GUID {
	return &IID_IFabricCodePackage2
}

func (this *IFabricCodePackage2) Get_SetupEntryPointRunAsPolicy() *FABRIC_RUNAS_POLICY_DESCRIPTION {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_RUNAS_POLICY_DESCRIPTION)(unsafe.Pointer(ret))
}

func (this *IFabricCodePackage2) Get_EntryPointRunAsPolicy() *FABRIC_RUNAS_POLICY_DESCRIPTION {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_RUNAS_POLICY_DESCRIPTION)(unsafe.Pointer(ret))
}
