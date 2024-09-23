package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 6C83D5C1-1954-4B80-9175-0D0E7C8715C9
var IID_IFabricCodePackageActivationContext2 = syscall.GUID{Data1: 0x6C83D5C1, Data2: 0x1954, Data3: 0x4B80,
	Data4: [8]byte{0x91, 0x75, 0x0D, 0x0E, 0x7C, 0x87, 0x15, 0xC9}}

type IFabricCodePackageActivationContext2 struct {
	IFabricCodePackageActivationContext
}

func NewIFabricCodePackageActivationContext2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricCodePackageActivationContext2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricCodePackageActivationContext2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricCodePackageActivationContext2) IID() *syscall.GUID {
	return &IID_IFabricCodePackageActivationContext2
}

func (this *IFabricCodePackageActivationContext2) Get_ApplicationName() string {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricCodePackageActivationContext2) Get_ApplicationTypeName() string {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricCodePackageActivationContext2) GetServiceManifestName(serviceManifestName **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[28]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceManifestName)))
	com.AddToScope(serviceManifestName)
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext2) GetServiceManifestVersion(serviceManifestVersion **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[29]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(serviceManifestVersion)))
	com.AddToScope(serviceManifestVersion)
	return com.Error(ret)
}
