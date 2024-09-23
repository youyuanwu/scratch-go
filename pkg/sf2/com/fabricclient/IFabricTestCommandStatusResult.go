package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 87798F5C-E600-493A-A926-16B6807378E6
var IID_IFabricTestCommandStatusResult = syscall.GUID{Data1: 0x87798F5C, Data2: 0xE600, Data3: 0x493A,
	Data4: [8]byte{0xA9, 0x26, 0x16, 0xB6, 0x80, 0x73, 0x78, 0xE6}}

type IFabricTestCommandStatusResult struct {
	win32.IUnknown
}

func NewIFabricTestCommandStatusResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricTestCommandStatusResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricTestCommandStatusResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricTestCommandStatusResult) IID() *syscall.GUID {
	return &IID_IFabricTestCommandStatusResult
}

func (this *IFabricTestCommandStatusResult) Get_Result() *TEST_COMMAND_QUERY_RESULT_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*TEST_COMMAND_QUERY_RESULT_LIST)(unsafe.Pointer(ret))
}
