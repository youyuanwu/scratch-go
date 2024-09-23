package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 77FF0C6B-6780-48EC-B4B0-61989327B0F2
var IID_IFabricStatefulServiceFactory = syscall.GUID{Data1: 0x77FF0C6B, Data2: 0x6780, Data3: 0x48EC,
	Data4: [8]byte{0xB4, 0xB0, 0x61, 0x98, 0x93, 0x27, 0xB0, 0xF2}}

type IFabricStatefulServiceFactory struct {
	win32.IUnknown
}

func NewIFabricStatefulServiceFactory(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricStatefulServiceFactory {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricStatefulServiceFactory)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricStatefulServiceFactory) IID() *syscall.GUID {
	return &IID_IFabricStatefulServiceFactory
}

func (this *IFabricStatefulServiceFactory) CreateReplica(serviceTypeName string, serviceName string, initializationDataLength uint32, initializationData *byte, partitionId syscall.GUID, replicaId int64, serviceReplica **IFabricStatefulServiceReplica) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceTypeName)), uintptr(win32.StrToPointer(serviceName)), uintptr(initializationDataLength), uintptr(unsafe.Pointer(initializationData)), (uintptr)(unsafe.Pointer(&partitionId)), uintptr(replicaId), uintptr(unsafe.Pointer(serviceReplica)))
	com.AddToScope(serviceReplica)
	return com.Error(ret)
}
