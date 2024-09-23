package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// A0CFBC71-184B-443B-B102-4B6D0A7CBC49
var IID_IFabricInfrastructureServiceClient = syscall.GUID{Data1: 0xA0CFBC71, Data2: 0x184B, Data3: 0x443B,
	Data4: [8]byte{0xB1, 0x02, 0x4B, 0x6D, 0x0A, 0x7C, 0xBC, 0x49}}

type IFabricInfrastructureServiceClient struct {
	win32.IUnknown
}

func NewIFabricInfrastructureServiceClient(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricInfrastructureServiceClient {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricInfrastructureServiceClient)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricInfrastructureServiceClient) IID() *syscall.GUID {
	return &IID_IFabricInfrastructureServiceClient
}

func (this *IFabricInfrastructureServiceClient) BeginInvokeInfrastructureCommand(serviceName string, command string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceName)), uintptr(win32.StrToPointer(command)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricInfrastructureServiceClient) EndInvokeInfrastructureCommand(context *IFabricAsyncOperationContext, result **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricInfrastructureServiceClient) BeginInvokeInfrastructureQuery(serviceName string, command string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceName)), uintptr(win32.StrToPointer(command)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricInfrastructureServiceClient) EndInvokeInfrastructureQuery(context *IFabricAsyncOperationContext, result **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
