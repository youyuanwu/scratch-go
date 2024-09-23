package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// E9F7F574-FD07-4A71-9F22-9CF9CCF3C166
var IID_IFabricGetServiceGroupMemberListResult = syscall.GUID{Data1: 0xE9F7F574, Data2: 0xFD07, Data3: 0x4A71,
	Data4: [8]byte{0x9F, 0x22, 0x9C, 0xF9, 0xCC, 0xF3, 0xC1, 0x66}}

type IFabricGetServiceGroupMemberListResult struct {
	win32.IUnknown
}

func NewIFabricGetServiceGroupMemberListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetServiceGroupMemberListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetServiceGroupMemberListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetServiceGroupMemberListResult) IID() *syscall.GUID {
	return &IID_IFabricGetServiceGroupMemberListResult
}

func (this *IFabricGetServiceGroupMemberListResult) Get_ServiceGroupMemberList() *FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_GROUP_MEMBER_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
