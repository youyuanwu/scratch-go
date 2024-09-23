package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 3860D61D-1E51-4A65-B109-D93C11311657
var IID_IFabricServiceGroupFactory = syscall.GUID{Data1: 0x3860D61D, Data2: 0x1E51, Data3: 0x4A65,
	Data4: [8]byte{0xB1, 0x09, 0xD9, 0x3C, 0x11, 0x31, 0x16, 0x57}}

type IFabricServiceGroupFactory struct {
	win32.IUnknown
}

func NewIFabricServiceGroupFactory(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceGroupFactory {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceGroupFactory)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceGroupFactory) IID() *syscall.GUID {
	return &IID_IFabricServiceGroupFactory
}
