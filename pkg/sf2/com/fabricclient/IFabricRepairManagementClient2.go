package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 5067D775-3BAA-48E4-8C72-BB5573CC3FB8
var IID_IFabricRepairManagementClient2 = syscall.GUID{Data1: 0x5067D775, Data2: 0x3BAA, Data3: 0x48E4,
	Data4: [8]byte{0x8C, 0x72, 0xBB, 0x55, 0x73, 0xCC, 0x3F, 0xB8}}

type IFabricRepairManagementClient2 struct {
	IFabricRepairManagementClient
}

func NewIFabricRepairManagementClient2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricRepairManagementClient2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricRepairManagementClient2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricRepairManagementClient2) IID() *syscall.GUID {
	return &IID_IFabricRepairManagementClient2
}

func (this *IFabricRepairManagementClient2) BeginUpdateRepairTaskHealthPolicy(updateDescription *FABRIC_REPAIR_TASK_HEALTH_POLICY_UPDATE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(updateDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient2) EndUpdateRepairTaskHealthPolicy(context *IFabricAsyncOperationContext, commitVersion *int64) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(commitVersion)))
	return com.Error(ret)
}
