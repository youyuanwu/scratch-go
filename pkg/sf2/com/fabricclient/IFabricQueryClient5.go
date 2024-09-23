package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 75C35E8C-87A2-4810-A401-B50DA858FE34
var IID_IFabricQueryClient5 = syscall.GUID{Data1: 0x75C35E8C, Data2: 0x87A2, Data3: 0x4810,
	Data4: [8]byte{0xA4, 0x01, 0xB5, 0x0D, 0xA8, 0x58, 0xFE, 0x34}}

type IFabricQueryClient5 struct {
	IFabricQueryClient4
}

func NewIFabricQueryClient5(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricQueryClient5 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricQueryClient5)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricQueryClient5) IID() *syscall.GUID {
	return &IID_IFabricQueryClient5
}

func (this *IFabricQueryClient5) BeginGetUnplacedReplicaInformation(queryDescription *FABRIC_UNPLACED_REPLICA_INFORMATION_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[45]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient5) EndGetUnplacedReplicaInformation(context *IFabricAsyncOperationContext, result **IFabricGetUnplacedReplicaInformationResult) com.Error {
	addr := (*this.LpVtbl)[46]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
