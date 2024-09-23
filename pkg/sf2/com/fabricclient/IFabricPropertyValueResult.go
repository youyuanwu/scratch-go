package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 9A518B49-9903-4B8F-834E-1979E9C6745E
var IID_IFabricPropertyValueResult = syscall.GUID{Data1: 0x9A518B49, Data2: 0x9903, Data3: 0x4B8F,
	Data4: [8]byte{0x83, 0x4E, 0x19, 0x79, 0xE9, 0xC6, 0x74, 0x5E}}

type IFabricPropertyValueResult struct {
	win32.IUnknown
}

func NewIFabricPropertyValueResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPropertyValueResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPropertyValueResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPropertyValueResult) IID() *syscall.GUID {
	return &IID_IFabricPropertyValueResult
}

func (this *IFabricPropertyValueResult) Get_Property() *FABRIC_NAMED_PROPERTY {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NAMED_PROPERTY)(unsafe.Pointer(ret))
}

func (this *IFabricPropertyValueResult) GetValueAsBinary(byteCount *uint32, bufferedValue **byte) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(byteCount)), uintptr(unsafe.Pointer(bufferedValue)))
	return com.Error(ret)
}

func (this *IFabricPropertyValueResult) GetValueAsInt64(value *int64) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *IFabricPropertyValueResult) GetValueAsDouble(value *float64) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *IFabricPropertyValueResult) GetValueAsWString(bufferedValue *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bufferedValue)))
	return com.Error(ret)
}

func (this *IFabricPropertyValueResult) GetValueAsGuid(value *syscall.GUID) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}
