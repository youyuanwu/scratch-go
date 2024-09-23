package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 5E572763-29A9-463A-B602-1332C0F60E6B
var IID_IFabricGetServiceGroupMemberTypeListResult = syscall.GUID{Data1: 0x5E572763, Data2: 0x29A9, Data3: 0x463A,
	Data4: [8]byte{0xB6, 0x02, 0x13, 0x32, 0xC0, 0xF6, 0x0E, 0x6B}}

type IFabricGetServiceGroupMemberTypeListResult struct {
	win32.IUnknown
}

func NewIFabricGetServiceGroupMemberTypeListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetServiceGroupMemberTypeListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetServiceGroupMemberTypeListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetServiceGroupMemberTypeListResult) IID() *syscall.GUID {
	return &IID_IFabricGetServiceGroupMemberTypeListResult
}

func (this *IFabricGetServiceGroupMemberTypeListResult) Get_ServiceGroupMemberTypeList() *FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_GROUP_MEMBER_TYPE_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
