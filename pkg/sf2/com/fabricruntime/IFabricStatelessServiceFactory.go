package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// CC53AF8F-74CD-11DF-AC3E-0024811E3892
var IID_IFabricStatelessServiceFactory = syscall.GUID{Data1: 0xCC53AF8F, Data2: 0x74CD, Data3: 0x11DF,
	Data4: [8]byte{0xAC, 0x3E, 0x00, 0x24, 0x81, 0x1E, 0x38, 0x92}}

type IFabricStatelessServiceFactory struct {
	win32.IUnknown
}

func NewIFabricStatelessServiceFactory(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatelessServiceFactory {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatelessServiceFactory)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatelessServiceFactory) IID() *syscall.GUID {
	return &IID_IFabricStatelessServiceFactory
}

func (this *IFabricStatelessServiceFactory) CreateInstance(serviceTypeName string, serviceName string, initializationDataLength uint32, initializationData *byte, partitionId syscall.GUID, instanceId int64, serviceInstance **IFabricStatelessServiceInstance) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceTypeName)), uintptr(win32.StrToPointer(serviceName)), uintptr(initializationDataLength), uintptr(unsafe.Pointer(initializationData)), (uintptr)(unsafe.Pointer(&partitionId)), uintptr(instanceId), uintptr(unsafe.Pointer(serviceInstance)))
	com.AddToScope(serviceInstance)
	return com.Error(ret)
}
