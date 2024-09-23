package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// BAB8AD87-37B7-482A-985D-BAF38A785DCD
var IID_IFabricOperationData = syscall.GUID{Data1: 0xBAB8AD87, Data2: 0x37B7, Data3: 0x482A,
	Data4: [8]byte{0x98, 0x5D, 0xBA, 0xF3, 0x8A, 0x78, 0x5D, 0xCD}}

type IFabricOperationData struct {
	win32.IUnknown
}

func NewIFabricOperationData(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricOperationData {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricOperationData)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricOperationData) IID() *syscall.GUID {
	return &IID_IFabricOperationData
}

func (this *IFabricOperationData) GetData(count *uint32, buffers **FABRIC_OPERATION_DATA_BUFFER) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)), uintptr(unsafe.Pointer(buffers)))
	return com.Error(ret)
}
