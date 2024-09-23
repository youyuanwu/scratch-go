package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// A42DA40D-A637-478D-83F3-2813871234CF
var IID_IFabricPropertyEnumerationResult = syscall.GUID{Data1: 0xA42DA40D, Data2: 0xA637, Data3: 0x478D,
	Data4: [8]byte{0x83, 0xF3, 0x28, 0x13, 0x87, 0x12, 0x34, 0xCF}}

type IFabricPropertyEnumerationResult struct {
	win32.IUnknown
}

func NewIFabricPropertyEnumerationResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPropertyEnumerationResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPropertyEnumerationResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPropertyEnumerationResult) IID() *syscall.GUID {
	return &IID_IFabricPropertyEnumerationResult
}

func (this *IFabricPropertyEnumerationResult) Get_EnumerationStatus() int32 {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return int32(ret)
}

func (this *IFabricPropertyEnumerationResult) Get_PropertyCount() uint32 {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return uint32(ret)
}

func (this *IFabricPropertyEnumerationResult) GetProperty(index uint32, property **IFabricPropertyValueResult) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(property)))
	com.AddToScope(property)
	return com.Error(ret)
}
