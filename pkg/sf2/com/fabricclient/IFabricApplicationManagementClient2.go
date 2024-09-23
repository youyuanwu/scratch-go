package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F873516F-9BFE-47E5-93B9-3667AAF19324
var IID_IFabricApplicationManagementClient2 = syscall.GUID{Data1: 0xF873516F, Data2: 0x9BFE, Data3: 0x47E5,
	Data4: [8]byte{0x93, 0xB9, 0x36, 0x67, 0xAA, 0xF1, 0x93, 0x24}}

type IFabricApplicationManagementClient2 struct {
	IFabricApplicationManagementClient
}

func NewIFabricApplicationManagementClient2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationManagementClient2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationManagementClient2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationManagementClient2) IID() *syscall.GUID {
	return &IID_IFabricApplicationManagementClient2
}

func (this *IFabricApplicationManagementClient2) BeginGetApplicationManifest(applicationTypeName string, applicationTypeVersion string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationTypeName)), uintptr(win32.StrToPointer(applicationTypeVersion)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient2) EndGetApplicationManifest(context *IFabricAsyncOperationContext, result **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient2) BeginMoveNextApplicationUpgradeDomain2(applicationName string, nextUpgradeDomain string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationName)), uintptr(win32.StrToPointer(nextUpgradeDomain)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient2) EndMoveNextApplicationUpgradeDomain2(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
