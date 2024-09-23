package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 16F563F3-4017-496E-B0E7-2650DE5774B3
var IID_IFabricQueryClient3 = syscall.GUID{Data1: 0x16F563F3, Data2: 0x4017, Data3: 0x496E,
	Data4: [8]byte{0xB0, 0xE7, 0x26, 0x50, 0xDE, 0x57, 0x74, 0xB3}}

type IFabricQueryClient3 struct {
	IFabricQueryClient2
}

func NewIFabricQueryClient3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricQueryClient3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricQueryClient3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricQueryClient3) IID() *syscall.GUID {
	return &IID_IFabricQueryClient3
}

func (this *IFabricQueryClient3) BeginGetNodeLoadInformation(queryDescription *FABRIC_NODE_LOAD_INFORMATION_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[37]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient3) EndGetNodeLoadInformation(context *IFabricAsyncOperationContext, result **IFabricGetNodeLoadInformationResult) com.Error {
	addr := (*this.LpVtbl)[38]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricQueryClient3) BeginGetReplicaLoadInformation(queryDescription *FABRIC_REPLICA_LOAD_INFORMATION_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[39]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient3) EndGetReplicaLoadInformation(context *IFabricAsyncOperationContext, result **IFabricGetReplicaLoadInformationResult) com.Error {
	addr := (*this.LpVtbl)[40]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
