package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 472BF2E1-D617-4B5C-A91D-FABED9FF3550
var IID_IFabricNodeContextResult2 = syscall.GUID{Data1: 0x472BF2E1, Data2: 0xD617, Data3: 0x4B5C,
	Data4: [8]byte{0xA9, 0x1D, 0xFA, 0xBE, 0xD9, 0xFF, 0x35, 0x50}}

type IFabricNodeContextResult2 struct {
	IFabricNodeContextResult
}

func NewIFabricNodeContextResult2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricNodeContextResult2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricNodeContextResult2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricNodeContextResult2) IID() *syscall.GUID {
	return &IID_IFabricNodeContextResult2
}

func (this *IFabricNodeContextResult2) GetDirectory(logicalDirectoryName string, directoryPath **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(logicalDirectoryName)), uintptr(unsafe.Pointer(directoryPath)))
	com.AddToScope(directoryPath)
	return com.Error(ret)
}
