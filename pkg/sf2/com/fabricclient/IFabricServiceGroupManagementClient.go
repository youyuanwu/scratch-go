package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 2061227E-0281-4BAF-9B19-B2DFB2E63BBE
var IID_IFabricServiceGroupManagementClient = syscall.GUID{Data1: 0x2061227E, Data2: 0x0281, Data3: 0x4BAF,
	Data4: [8]byte{0x9B, 0x19, 0xB2, 0xDF, 0xB2, 0xE6, 0x3B, 0xBE}}

type IFabricServiceGroupManagementClient struct {
	win32.IUnknown
}

func NewIFabricServiceGroupManagementClient(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceGroupManagementClient {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceGroupManagementClient)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceGroupManagementClient) IID() *syscall.GUID {
	return &IID_IFabricServiceGroupManagementClient
}

func (this *IFabricServiceGroupManagementClient) BeginCreateServiceGroup(description *FABRIC_SERVICE_GROUP_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceGroupManagementClient) EndCreateServiceGroup(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricServiceGroupManagementClient) BeginDeleteServiceGroup(name string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceGroupManagementClient) EndDeleteServiceGroup(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricServiceGroupManagementClient) BeginGetServiceGroupDescription(name string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceGroupManagementClient) EndGetServiceGroupDescription(context *IFabricAsyncOperationContext, result **IFabricServiceGroupDescriptionResult) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
