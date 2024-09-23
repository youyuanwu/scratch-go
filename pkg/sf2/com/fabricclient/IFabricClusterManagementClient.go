package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// A3CF17E0-CF84-4AE0-B720-1785C0FB4ACE
var IID_IFabricClusterManagementClient = syscall.GUID{Data1: 0xA3CF17E0, Data2: 0xCF84, Data3: 0x4AE0,
	Data4: [8]byte{0xB7, 0x20, 0x17, 0x85, 0xC0, 0xFB, 0x4A, 0xCE}}

type IFabricClusterManagementClient struct {
	win32.IUnknown
}

func NewIFabricClusterManagementClient(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClusterManagementClient {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClusterManagementClient)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClusterManagementClient) IID() *syscall.GUID {
	return &IID_IFabricClusterManagementClient
}

func (this *IFabricClusterManagementClient) BeginNodeStateRemoved(nodeName string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(nodeName)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient) EndNodeStateRemoved(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient) BeginRecoverPartitions(timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient) EndRecoverPartitions(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
