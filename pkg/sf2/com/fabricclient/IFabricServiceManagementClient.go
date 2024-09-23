package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F7368189-FD1F-437C-888D-8C89CECC57A0
var IID_IFabricServiceManagementClient = syscall.GUID{Data1: 0xF7368189, Data2: 0xFD1F, Data3: 0x437C,
	Data4: [8]byte{0x88, 0x8D, 0x8C, 0x89, 0xCE, 0xCC, 0x57, 0xA0}}

type IFabricServiceManagementClient struct {
	win32.IUnknown
}

func NewIFabricServiceManagementClient(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceManagementClient {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceManagementClient)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceManagementClient) IID() *syscall.GUID {
	return &IID_IFabricServiceManagementClient
}

func (this *IFabricServiceManagementClient) BeginCreateService(description *FABRIC_SERVICE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) EndCreateService(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) BeginCreateServiceFromTemplate(applicationName string, serviceName string, serviceTypeName string, initializationDataSize uint32, initializationData *byte, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationName)), uintptr(win32.StrToPointer(serviceName)), uintptr(win32.StrToPointer(serviceTypeName)), uintptr(initializationDataSize), uintptr(unsafe.Pointer(initializationData)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) EndCreateServiceFromTemplate(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) BeginDeleteService(name string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) EndDeleteService(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) BeginGetServiceDescription(name string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) EndGetServiceDescription(context *IFabricAsyncOperationContext, result **IFabricServiceDescriptionResult) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) RegisterServicePartitionResolutionChangeHandler(name string, keyType int32, partitionKey unsafe.Pointer, callback *IFabricServicePartitionResolutionChangeHandler, callbackHandle *int64) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(keyType), uintptr(partitionKey), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(callbackHandle)))
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) UnregisterServicePartitionResolutionChangeHandler(callbackHandle int64) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(callbackHandle))
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) BeginResolveServicePartition(name string, partitionKeyType int32, partitionKey unsafe.Pointer, previousResult *IFabricResolvedServicePartitionResult, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(partitionKeyType), uintptr(partitionKey), uintptr(unsafe.Pointer(previousResult)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricServiceManagementClient) EndResolveServicePartition(context *IFabricAsyncOperationContext, result **IFabricResolvedServicePartitionResult) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
