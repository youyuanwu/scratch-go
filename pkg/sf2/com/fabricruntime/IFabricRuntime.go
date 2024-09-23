package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// CC53AF8E-74CD-11DF-AC3E-0024811E3892
var IID_IFabricRuntime = syscall.GUID{Data1: 0xCC53AF8E, Data2: 0x74CD, Data3: 0x11DF,
	Data4: [8]byte{0xAC, 0x3E, 0x00, 0x24, 0x81, 0x1E, 0x38, 0x92}}

type IFabricRuntime struct {
	win32.IUnknown
}

func NewIFabricRuntime(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricRuntime {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricRuntime)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricRuntime) IID() *syscall.GUID {
	return &IID_IFabricRuntime
}

func (this *IFabricRuntime) BeginRegisterStatelessServiceFactory(serviceTypeName string, factory *IFabricStatelessServiceFactory, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceTypeName)), uintptr(unsafe.Pointer(factory)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRuntime) EndRegisterStatelessServiceFactory(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricRuntime) RegisterStatelessServiceFactory(serviceTypeName string, factory *IFabricStatelessServiceFactory) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceTypeName)), uintptr(unsafe.Pointer(factory)))
	return com.Error(ret)
}

func (this *IFabricRuntime) BeginRegisterStatefulServiceFactory(serviceTypeName string, factory *IFabricStatefulServiceFactory, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceTypeName)), uintptr(unsafe.Pointer(factory)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRuntime) EndRegisterStatefulServiceFactory(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricRuntime) RegisterStatefulServiceFactory(serviceTypeName string, factory *IFabricStatefulServiceFactory) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceTypeName)), uintptr(unsafe.Pointer(factory)))
	return com.Error(ret)
}

func (this *IFabricRuntime) CreateServiceGroupFactoryBuilder(builder **IFabricServiceGroupFactoryBuilder) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(builder)))
	com.AddToScope(builder)
	return com.Error(ret)
}

func (this *IFabricRuntime) BeginRegisterServiceGroupFactory(groupServiceType string, factory *IFabricServiceGroupFactory, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(groupServiceType)), uintptr(unsafe.Pointer(factory)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricRuntime) EndRegisterServiceGroupFactory(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricRuntime) RegisterServiceGroupFactory(groupServiceType string, factory *IFabricServiceGroupFactory) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(groupServiceType)), uintptr(unsafe.Pointer(factory)))
	return com.Error(ret)
}
