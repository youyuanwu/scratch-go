package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// DD3E4497-3373-458D-AD22-C88EBD27493E
var IID_IFabricHealthClient3 = syscall.GUID{Data1: 0xDD3E4497, Data2: 0x3373, Data3: 0x458D,
	Data4: [8]byte{0xAD, 0x22, 0xC8, 0x8E, 0xBD, 0x27, 0x49, 0x3E}}

type IFabricHealthClient3 struct {
	IFabricHealthClient2
}

func NewIFabricHealthClient3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricHealthClient3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricHealthClient3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricHealthClient3) IID() *syscall.GUID {
	return &IID_IFabricHealthClient3
}

func (this *IFabricHealthClient3) BeginGetClusterHealthChunk(queryDescription *FABRIC_CLUSTER_HEALTH_CHUNK_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[36]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricHealthClient3) EndGetClusterHealthChunk(context *IFabricAsyncOperationContext, result **IFabricGetClusterHealthChunkResult) com.Error {
	addr := (*this.LpVtbl)[37]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
