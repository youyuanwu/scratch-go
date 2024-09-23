package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// AA3116FE-277D-482D-BD16-5366FA405757
var IID_IFabricReplicatorCatchupSpecificQuorum = syscall.GUID{Data1: 0xAA3116FE, Data2: 0x277D, Data3: 0x482D,
	Data4: [8]byte{0xBD, 0x16, 0x53, 0x66, 0xFA, 0x40, 0x57, 0x57}}

type IFabricReplicatorCatchupSpecificQuorum struct {
	win32.IUnknown
}

func NewIFabricReplicatorCatchupSpecificQuorum(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricReplicatorCatchupSpecificQuorum {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricReplicatorCatchupSpecificQuorum)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricReplicatorCatchupSpecificQuorum) IID() *syscall.GUID {
	return &IID_IFabricReplicatorCatchupSpecificQuorum
}
