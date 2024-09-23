package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// EE483BA5-9018-4C99-9804-BE6185DB88E6
var IID_IFabricRepairManagementClient = syscall.GUID{Data1: 0xEE483BA5, Data2: 0x9018, Data3: 0x4C99,
	Data4: [8]byte{0x98, 0x04, 0xBE, 0x61, 0x85, 0xDB, 0x88, 0xE6}}

type IFabricRepairManagementClient struct {
	win32.IUnknown
}

func NewIFabricRepairManagementClient(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricRepairManagementClient {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricRepairManagementClient)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricRepairManagementClient) IID() *syscall.GUID {
	return &IID_IFabricRepairManagementClient
}

func (this *IFabricRepairManagementClient) BeginCreateRepairTask(repairTask *FABRIC_REPAIR_TASK, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(repairTask)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) EndCreateRepairTask(context *IFabricAsyncOperationContext, commitVersion *int64) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(commitVersion)))
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) BeginCancelRepairTask(requestDescription *FABRIC_REPAIR_CANCEL_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(requestDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) EndCancelRepairTask(context *IFabricAsyncOperationContext, commitVersion *int64) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(commitVersion)))
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) BeginForceApproveRepairTask(requestDescription *FABRIC_REPAIR_APPROVE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(requestDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) EndForceApproveRepairTask(context *IFabricAsyncOperationContext, commitVersion *int64) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(commitVersion)))
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) BeginDeleteRepairTask(requestDescription *FABRIC_REPAIR_DELETE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(requestDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) EndDeleteRepairTask(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) BeginUpdateRepairExecutionState(repairTask *FABRIC_REPAIR_TASK, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(repairTask)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) EndUpdateRepairExecutionState(context *IFabricAsyncOperationContext, commitVersion *int64) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(commitVersion)))
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) BeginGetRepairTaskList(queryDescription *FABRIC_REPAIR_TASK_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRepairManagementClient) EndGetRepairTaskList(context *IFabricAsyncOperationContext, result **IFabricGetRepairTaskListResult) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
