package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 9933ED08-5D0C-4AED-BAB6-F676BF5BE8AA
var IID_IFabricServiceManagementClient2 = syscall.GUID{Data1: 0x9933ED08, Data2: 0x5D0C, Data3: 0x4AED,
	Data4: [8]byte{0xBA, 0xB6, 0xF6, 0x76, 0xBF, 0x5B, 0xE8, 0xAA}}

type IFabricServiceManagementClient2 struct {
	IFabricServiceManagementClient
}

func NewIFabricServiceManagementClient2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceManagementClient2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceManagementClient2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceManagementClient2) IID() *syscall.GUID {
	return &IID_IFabricServiceManagementClient2
}

func (this *IFabricServiceManagementClient2) BeginGetServiceManifest(applicationTypeName string, applicationTypeVersion string, serviceManifestName string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationTypeName)), uintptr(win32.StrToPointer(applicationTypeVersion)), uintptr(win32.StrToPointer(serviceManifestName)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient2) EndGetServiceManifest(context *IFabricAsyncOperationContext, result **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient2) BeginUpdateService(name string, serviceUpdateDescription *FABRIC_SERVICE_UPDATE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(serviceUpdateDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient2) EndUpdateService(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
