package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// FA5FDA9B-472C-45A0-9B60-A374691227A4
var IID_IFabricCodePackageActivationContext6 = syscall.GUID{Data1: 0xFA5FDA9B, Data2: 0x472C, Data3: 0x45A0,
	Data4: [8]byte{0x9B, 0x60, 0xA3, 0x74, 0x69, 0x12, 0x27, 0xA4}}

type IFabricCodePackageActivationContext6 struct {
	IFabricCodePackageActivationContext5
}

func NewIFabricCodePackageActivationContext6(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricCodePackageActivationContext6 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricCodePackageActivationContext6)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricCodePackageActivationContext6) IID() *syscall.GUID {
	return &IID_IFabricCodePackageActivationContext6
}

func (this *IFabricCodePackageActivationContext6) GetDirectory(logicalDirectoryName string, directoryPath **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[38]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(logicalDirectoryName)), uintptr(unsafe.Pointer(directoryPath)))
	com.AddToScope(directoryPath)
	return com.Error(ret)
}
