package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 70BE1B10-B259-46FC-B813-0B75720E7183
var IID_IFabricCodePackageActivator = syscall.GUID{Data1: 0x70BE1B10, Data2: 0xB259, Data3: 0x46FC,
	Data4: [8]byte{0xB8, 0x13, 0x0B, 0x75, 0x72, 0x0E, 0x71, 0x83}}

type IFabricCodePackageActivator struct {
	win32.IUnknown
}

func NewIFabricCodePackageActivator(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricCodePackageActivator {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricCodePackageActivator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricCodePackageActivator) IID() *syscall.GUID {
	return &IID_IFabricCodePackageActivator
}

func (this *IFabricCodePackageActivator) BeginActivateCodePackage(codePackageNames *FABRIC_STRING_LIST, environment *FABRIC_STRING_MAP, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(codePackageNames)), uintptr(unsafe.Pointer(environment)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricCodePackageActivator) EndActivateCodePackage(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivator) BeginDeactivateCodePackage(codePackageNames *FABRIC_STRING_LIST, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(codePackageNames)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricCodePackageActivator) EndDeactivateCodePackage(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivator) AbortCodePackage(codePackageNames *FABRIC_STRING_LIST) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(codePackageNames)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivator) RegisterCodePackageEventHandler(eventHandler *IFabricCodePackageEventHandler, callbackHandle *uint64) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(callbackHandle)))
	return com.Error(ret)
}

func (this *IFabricCodePackageActivator) UnregisterCodePackageEventHandler(callbackHandle uint64) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(callbackHandle))
	return com.Error(ret)
}
