package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 68A971E2-F15F-4D95-A79C-8A257909659E
var IID_IFabricCodePackageActivationContext = syscall.GUID{Data1: 0x68A971E2, Data2: 0xF15F, Data3: 0x4D95,
	Data4: [8]byte{0xA7, 0x9C, 0x8A, 0x25, 0x79, 0x09, 0x65, 0x9E}}

type IFabricCodePackageActivationContext struct {
	win32.IUnknown
}

func NewIFabricCodePackageActivationContext(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricCodePackageActivationContext {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricCodePackageActivationContext)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricCodePackageActivationContext) IID() *syscall.GUID {
	return &IID_IFabricCodePackageActivationContext
}

func (this *IFabricCodePackageActivationContext) Get_ContextId() string {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricCodePackageActivationContext) Get_CodePackageName() string {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricCodePackageActivationContext) Get_CodePackageVersion() string {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricCodePackageActivationContext) Get_WorkDirectory() string {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricCodePackageActivationContext) Get_LogDirectory() string {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricCodePackageActivationContext) Get_TempDirectory() string {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return win32.BstrToStrAndFree(win32.BSTR(unsafe.Pointer(ret)))
}

func (this *IFabricCodePackageActivationContext) Get_ServiceTypes() *FABRIC_SERVICE_TYPE_DESCRIPTION_LIST {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_TYPE_DESCRIPTION_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricCodePackageActivationContext) Get_ServiceGroupTypes() *FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION_LIST {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_GROUP_TYPE_DESCRIPTION_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricCodePackageActivationContext) Get_ApplicationPrincipals() *FABRIC_APPLICATION_PRINCIPALS_DESCRIPTION {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_APPLICATION_PRINCIPALS_DESCRIPTION)(unsafe.Pointer(ret))
}

func (this *IFabricCodePackageActivationContext) Get_ServiceEndpointResources() *FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_LIST {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_ENDPOINT_RESOURCE_DESCRIPTION_LIST)(unsafe.Pointer(ret))
}

func (this *IFabricCodePackageActivationContext) GetServiceEndpointResource(serviceEndpointResourceName string, bufferedValue **FABRIC_ENDPOINT_RESOURCE_DESCRIPTION) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceEndpointResourceName)), uintptr(unsafe.Pointer(bufferedValue)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) GetCodePackageNames(names **IFabricStringListResult) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(names)))
	com.AddToScope(names)
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) GetConfigurationPackageNames(names **IFabricStringListResult) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(names)))
	com.AddToScope(names)
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) GetDataPackageNames(names **IFabricStringListResult) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(names)))
	com.AddToScope(names)
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) GetCodePackage(codePackageName string, codePackage **IFabricCodePackage) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(codePackageName)), uintptr(unsafe.Pointer(codePackage)))
	com.AddToScope(codePackage)
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) GetConfigurationPackage(configPackageName string, configPackage **IFabricConfigurationPackage) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(configPackageName)), uintptr(unsafe.Pointer(configPackage)))
	com.AddToScope(configPackage)
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) GetDataPackage(dataPackageName string, dataPackage **IFabricDataPackage) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(dataPackageName)), uintptr(unsafe.Pointer(dataPackage)))
	com.AddToScope(dataPackage)
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) RegisterCodePackageChangeHandler(callback *IFabricCodePackageChangeHandler, callbackHandle *int64) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(callbackHandle)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) UnregisterCodePackageChangeHandler(callbackHandle int64) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(callbackHandle))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) RegisterConfigurationPackageChangeHandler(callback *IFabricConfigurationPackageChangeHandler, callbackHandle *int64) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(callbackHandle)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) UnregisterConfigurationPackageChangeHandler(callbackHandle int64) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(callbackHandle))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) RegisterDataPackageChangeHandler(callback *IFabricDataPackageChangeHandler, callbackHandle *int64) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(callbackHandle)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivationContext) UnregisterDataPackageChangeHandler(callbackHandle int64) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(callbackHandle))
	return com.Error(ret)
}
