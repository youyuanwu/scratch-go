package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 108C7735-97E1-4AF8-8C2D-9080B1B29D33
var IID_IFabricApplicationManagementClient3 = syscall.GUID{Data1: 0x108C7735, Data2: 0x97E1, Data3: 0x4AF8,
	Data4: [8]byte{0x8C, 0x2D, 0x90, 0x80, 0xB1, 0xB2, 0x9D, 0x33}}

type IFabricApplicationManagementClient3 struct {
	IFabricApplicationManagementClient2
}

func NewIFabricApplicationManagementClient3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationManagementClient3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationManagementClient3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationManagementClient3) IID() *syscall.GUID {
	return &IID_IFabricApplicationManagementClient3
}

func (this *IFabricApplicationManagementClient3) BeginUpdateApplicationUpgrade(description *FABRIC_APPLICATION_UPGRADE_UPDATE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient3) EndUpdateApplicationUpgrade(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient3) BeginRestartDeployedCodePackage(restartCodePackageDescription *FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(restartCodePackageDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient3) EndRestartDeployedCodePackage(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient3) CopyApplicationPackage(imageStoreConnectionString string, applicationPackagePath string, applicationPackagePathInImageStore string) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(imageStoreConnectionString)), uintptr(win32.StrToPointer(applicationPackagePath)), uintptr(win32.StrToPointer(applicationPackagePathInImageStore)))
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient3) RemoveApplicationPackage(imageStoreConnectionString string, applicationPackagePathInImageStore string) com.Error {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(imageStoreConnectionString)), uintptr(win32.StrToPointer(applicationPackagePathInImageStore)))
	return com.Error(ret)
}
