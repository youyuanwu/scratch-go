package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F4AD6BFA-E23C-4A48-9617-C099CD59A23A
var IID_IFabricOperation = syscall.GUID{Data1: 0xF4AD6BFA, Data2: 0xE23C, Data3: 0x4A48,
	Data4: [8]byte{0x96, 0x17, 0xC0, 0x99, 0xCD, 0x59, 0xA2, 0x3A}}

type IFabricOperation struct {
	win32.IUnknown
}

func NewIFabricOperation(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricOperation {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricOperation)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricOperation) IID() *syscall.GUID {
	return &IID_IFabricOperation
}

func (this *IFabricOperation) Get_Metadata() *FABRIC_OPERATION_METADATA {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_OPERATION_METADATA)(unsafe.Pointer(ret))
}

func (this *IFabricOperation) GetData(count *uint32, buffers **FABRIC_OPERATION_DATA_BUFFER) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)), uintptr(unsafe.Pointer(buffers)))
	return com.Error(ret)
}

func (this *IFabricOperation) Acknowledge() com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}
