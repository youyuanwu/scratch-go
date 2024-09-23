package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 576B2462-5F69-4351-87C7-3EC2D1654A22
var IID_IFabricGetRepairTaskListResult = syscall.GUID{Data1: 0x576B2462, Data2: 0x5F69, Data3: 0x4351,
	Data4: [8]byte{0x87, 0xC7, 0x3E, 0xC2, 0xD1, 0x65, 0x4A, 0x22}}

type IFabricGetRepairTaskListResult struct {
	win32.IUnknown
}

func NewIFabricGetRepairTaskListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetRepairTaskListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetRepairTaskListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetRepairTaskListResult) IID() *syscall.GUID {
	return &IID_IFabricGetRepairTaskListResult
}

func (this *IFabricGetRepairTaskListResult) Get_Tasks() *FABRIC_REPAIR_TASK_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_REPAIR_TASK_LIST)(unsafe.Pointer(ret))
}
