package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// A9FE8B06-19B1-49E6-8911-41D9D9219E1C
var IID_IFabricServiceGroupFactoryBuilder = syscall.GUID{Data1: 0xA9FE8B06, Data2: 0x19B1, Data3: 0x49E6,
	Data4: [8]byte{0x89, 0x11, 0x41, 0xD9, 0xD9, 0x21, 0x9E, 0x1C}}

type IFabricServiceGroupFactoryBuilder struct {
	win32.IUnknown
}

func NewIFabricServiceGroupFactoryBuilder(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceGroupFactoryBuilder {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceGroupFactoryBuilder)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceGroupFactoryBuilder) IID() *syscall.GUID {
	return &IID_IFabricServiceGroupFactoryBuilder
}

func (this *IFabricServiceGroupFactoryBuilder) AddStatelessServiceFactory(memberServiceType string, factory *IFabricStatelessServiceFactory) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(memberServiceType)), uintptr(unsafe.Pointer(factory)))
	return com.Error(ret)
}

func (this *IFabricServiceGroupFactoryBuilder) AddStatefulServiceFactory(memberServiceType string, factory *IFabricStatefulServiceFactory) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(memberServiceType)), uintptr(unsafe.Pointer(factory)))
	return com.Error(ret)
}

func (this *IFabricServiceGroupFactoryBuilder) RemoveServiceFactory(memberServiceType string) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(memberServiceType)))
	return com.Error(ret)
}

func (this *IFabricServiceGroupFactoryBuilder) ToServiceGroupFactory(factory **IFabricServiceGroupFactory) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factory)))
	com.AddToScope(factory)
	return com.Error(ret)
}
