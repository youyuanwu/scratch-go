package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// D3161F31-708A-4F83-91FF-F2AF15F74A2F
var IID_IFabricConfigurationPackage2 = syscall.GUID{Data1: 0xD3161F31, Data2: 0x708A, Data3: 0x4F83,
	Data4: [8]byte{0x91, 0xFF, 0xF2, 0xAF, 0x15, 0xF7, 0x4A, 0x2F}}

type IFabricConfigurationPackage2 struct {
	IFabricConfigurationPackage
}

func NewIFabricConfigurationPackage2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricConfigurationPackage2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricConfigurationPackage2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricConfigurationPackage2) IID() *syscall.GUID {
	return &IID_IFabricConfigurationPackage2
}

func (this *IFabricConfigurationPackage2) GetValues(sectionName string, parameterPrefix string, bufferedValue **FABRIC_CONFIGURATION_PARAMETER_LIST) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(sectionName)), uintptr(win32.StrToPointer(parameterPrefix)), uintptr(unsafe.Pointer(bufferedValue)))
	return com.Error(ret)
}
