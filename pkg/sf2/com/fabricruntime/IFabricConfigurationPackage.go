package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// AC4C3BFA-2563-46B7-A71D-2DCA7B0A8F4D
var IID_IFabricConfigurationPackage = syscall.GUID{Data1: 0xAC4C3BFA, Data2: 0x2563, Data3: 0x46B7,
	Data4: [8]byte{0xA7, 0x1D, 0x2D, 0xCA, 0x7B, 0x0A, 0x8F, 0x4D}}

type IFabricConfigurationPackage struct {
	win32.IUnknown
}

func NewIFabricConfigurationPackage(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricConfigurationPackage {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricConfigurationPackage)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricConfigurationPackage) IID() *syscall.GUID {
	return &IID_IFabricConfigurationPackage
}

func (this *IFabricConfigurationPackage) Get_Description() *FABRIC_CONFIGURATION_PACKAGE_DESCRIPTION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CONFIGURATION_PACKAGE_DESCRIPTION)(unsafe.Pointer(ret))
}

func (this *IFabricConfigurationPackage) Get_Path() string {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricConfigurationPackage) Get_Settings() *FABRIC_CONFIGURATION_SETTINGS {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CONFIGURATION_SETTINGS)(unsafe.Pointer(ret))
}

func (this *IFabricConfigurationPackage) GetSection(sectionName string, bufferedValue **FABRIC_CONFIGURATION_SECTION) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(sectionName)), uintptr(unsafe.Pointer(bufferedValue)))
	return com.Error(ret)
}

func (this *IFabricConfigurationPackage) GetValue(sectionName string, parameterName string, isEncrypted *int8, bufferedValue *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(sectionName)), uintptr(win32.StrToPointer(parameterName)), uintptr(unsafe.Pointer(isEncrypted)), uintptr(unsafe.Pointer(bufferedValue)))
	return com.Error(ret)
}

func (this *IFabricConfigurationPackage) DecryptValue(encryptedValue string, decryptedValue **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(encryptedValue)), uintptr(unsafe.Pointer(decryptedValue)))
	com.AddToScope(decryptedValue)
	return com.Error(ret)
}
