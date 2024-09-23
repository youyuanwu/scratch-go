package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// CC53AF90-74CD-11DF-AC3E-0024811E3892
var IID_IFabricStatelessServiceInstance = syscall.GUID{Data1: 0xCC53AF90, Data2: 0x74CD, Data3: 0x11DF,
	Data4: [8]byte{0xAC, 0x3E, 0x00, 0x24, 0x81, 0x1E, 0x38, 0x92}}

type IFabricStatelessServiceInstance struct {
	win32.IUnknown
}

func NewIFabricStatelessServiceInstance(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatelessServiceInstance {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatelessServiceInstance)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatelessServiceInstance) IID() *syscall.GUID {
	return &IID_IFabricStatelessServiceInstance
}

func (this *IFabricStatelessServiceInstance) BeginOpen(partition *IFabricStatelessServicePartition, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(partition)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricStatelessServiceInstance) EndOpen(context *IFabricAsyncOperationContext, serviceAddress **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(serviceAddress)))
	com.AddToScope(serviceAddress)
	return com.Error(ret)
}

func (this *IFabricStatelessServiceInstance) BeginClose(callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricStatelessServiceInstance) EndClose(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricStatelessServiceInstance) Abort() {
	addr := (*this.LpVtbl)[7]
	_, _, _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
}
