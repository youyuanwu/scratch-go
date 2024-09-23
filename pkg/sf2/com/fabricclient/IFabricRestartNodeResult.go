package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 2F7E9D57-FE07-4E34-93E1-01D5A6298CA9
var IID_IFabricRestartNodeResult = syscall.GUID{Data1: 0x2F7E9D57, Data2: 0xFE07, Data3: 0x4E34,
	Data4: [8]byte{0x93, 0xE1, 0x01, 0xD5, 0xA6, 0x29, 0x8C, 0xA9}}

type IFabricRestartNodeResult struct {
	win32.IUnknown
}

func NewIFabricRestartNodeResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricRestartNodeResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricRestartNodeResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricRestartNodeResult) IID() *syscall.GUID {
	return &IID_IFabricRestartNodeResult
}

func (this *IFabricRestartNodeResult) Get_Result() *FABRIC_NODE_RESULT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NODE_RESULT)(unsafe.Pointer(ret))
}
