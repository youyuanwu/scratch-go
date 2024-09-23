package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// AB92081D-0D78-410B-9777-0846DBA24C10
var IID_IFabricQueryClient4 = syscall.GUID{Data1: 0xAB92081D, Data2: 0x0D78, Data3: 0x410B,
	Data4: [8]byte{0x97, 0x77, 0x08, 0x46, 0xDB, 0xA2, 0x4C, 0x10}}

type IFabricQueryClient4 struct {
	IFabricQueryClient3
}

func NewIFabricQueryClient4(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricQueryClient4 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricQueryClient4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricQueryClient4) IID() *syscall.GUID {
	return &IID_IFabricQueryClient4
}

func (this *IFabricQueryClient4) BeginGetServiceGroupMemberList(queryDescription *FABRIC_SERVICE_GROUP_MEMBER_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[41]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient4) EndGetServiceGroupMemberList(context *IFabricAsyncOperationContext, result **IFabricGetServiceGroupMemberListResult) com.Error {
	addr := (*this.LpVtbl)[42]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricQueryClient4) BeginGetServiceGroupMemberTypeList(queryDescription *FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[43]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient4) EndGetServiceGroupMemberTypeList(context *IFabricAsyncOperationContext, result **IFabricGetServiceGroupMemberTypeListResult) com.Error {
	addr := (*this.LpVtbl)[44]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
