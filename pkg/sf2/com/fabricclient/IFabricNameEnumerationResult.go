package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 557E8105-F4F4-4FD3-9D21-075F34E2F98C
var IID_IFabricNameEnumerationResult = syscall.GUID{Data1: 0x557E8105, Data2: 0xF4F4, Data3: 0x4FD3,
	Data4: [8]byte{0x9D, 0x21, 0x07, 0x5F, 0x34, 0xE2, 0xF9, 0x8C}}

type IFabricNameEnumerationResult struct {
	win32.IUnknown
}

func NewIFabricNameEnumerationResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricNameEnumerationResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricNameEnumerationResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricNameEnumerationResult) IID() *syscall.GUID {
	return &IID_IFabricNameEnumerationResult
}

func (this *IFabricNameEnumerationResult) Get_EnumerationStatus() int32 {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int32(ret)
}

func (this *IFabricNameEnumerationResult) GetNames(itemCount *uint32, bufferedItems **win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(itemCount)), uintptr(unsafe.Pointer(bufferedItems)))
	return com.Error(ret)
}
