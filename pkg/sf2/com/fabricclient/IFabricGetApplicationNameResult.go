package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 258DBCC8-AC9A-47FF-838B-57FF506C73B1
var IID_IFabricGetApplicationNameResult = syscall.GUID{Data1: 0x258DBCC8, Data2: 0xAC9A, Data3: 0x47FF,
	Data4: [8]byte{0x83, 0x8B, 0x57, 0xFF, 0x50, 0x6C, 0x73, 0xB1}}

type IFabricGetApplicationNameResult struct {
	win32.IUnknown
}

func NewIFabricGetApplicationNameResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetApplicationNameResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetApplicationNameResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetApplicationNameResult) IID() *syscall.GUID {
	return &IID_IFabricGetApplicationNameResult
}

func (this *IFabricGetApplicationNameResult) Get_ApplicationName() *FABRIC_APPLICATION_NAME_QUERY_RESULT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_APPLICATION_NAME_QUERY_RESULT)(unsafe.Pointer(ret))
}
