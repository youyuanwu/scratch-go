package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 33302306-FB8D-4831-B493-57EFCC772462
var IID_IFabricPropertyMetadataResult = syscall.GUID{Data1: 0x33302306, Data2: 0xFB8D, Data3: 0x4831,
	Data4: [8]byte{0xB4, 0x93, 0x57, 0xEF, 0xCC, 0x77, 0x24, 0x62}}

type IFabricPropertyMetadataResult struct {
	win32.IUnknown
}

func NewIFabricPropertyMetadataResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricPropertyMetadataResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricPropertyMetadataResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricPropertyMetadataResult) IID() *syscall.GUID {
	return &IID_IFabricPropertyMetadataResult
}

func (this *IFabricPropertyMetadataResult) Get_Metadata() *FABRIC_NAMED_PROPERTY_METADATA {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NAMED_PROPERTY_METADATA)(unsafe.Pointer(ret))
}
