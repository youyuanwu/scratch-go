package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 173B2BB4-09C6-42FB-8754-CAA8D43CF1B2
var IID_IFabricQueryClient6 = syscall.GUID{Data1: 0x173B2BB4, Data2: 0x09C6, Data3: 0x42FB,
	Data4: [8]byte{0x87, 0x54, 0xCA, 0xA8, 0xD4, 0x3C, 0xF1, 0xB2}}

type IFabricQueryClient6 struct {
	IFabricQueryClient5
}

func NewIFabricQueryClient6(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricQueryClient6 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricQueryClient6)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricQueryClient6) IID() *syscall.GUID {
	return &IID_IFabricQueryClient6
}

func (this *IFabricQueryClient6) EndGetNodeList2(context *IFabricAsyncOperationContext, result **IFabricGetNodeListResult2) com.Error {
	addr := (*this.LpVtbl)[47]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricQueryClient6) EndGetApplicationList2(context *IFabricAsyncOperationContext, result **IFabricGetApplicationListResult2) com.Error {
	addr := (*this.LpVtbl)[48]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricQueryClient6) EndGetServiceList2(context *IFabricAsyncOperationContext, result **IFabricGetServiceListResult2) com.Error {
	addr := (*this.LpVtbl)[49]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricQueryClient6) EndGetPartitionList2(context *IFabricAsyncOperationContext, result **IFabricGetPartitionListResult2) com.Error {
	addr := (*this.LpVtbl)[50]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricQueryClient6) EndGetReplicaList2(context *IFabricAsyncOperationContext, result **IFabricGetReplicaListResult2) com.Error {
	addr := (*this.LpVtbl)[51]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
