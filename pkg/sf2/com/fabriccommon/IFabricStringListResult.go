package fabriccommon

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// AFAB1C53-757B-4B0E-8B7E-237AEEE6BFE9
var IID_IFabricStringListResult = syscall.GUID{Data1: 0xAFAB1C53, Data2: 0x757B, Data3: 0x4B0E,
	Data4: [8]byte{0x8B, 0x7E, 0x23, 0x7A, 0xEE, 0xE6, 0xBF, 0xE9}}

type IFabricStringListResult struct {
	win32.IUnknown
}

func NewIFabricStringListResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStringListResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStringListResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStringListResult) IID() *syscall.GUID {
	return &IID_IFabricStringListResult
}

func (this *IFabricStringListResult) GetStrings(itemCount *uint32, bufferedItems **win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(itemCount)), uintptr(unsafe.Pointer(bufferedItems)))
	return com.Error(ret)
}
