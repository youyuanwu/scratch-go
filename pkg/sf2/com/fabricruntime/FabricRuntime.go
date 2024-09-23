package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

var CLSID_FabricRuntime = syscall.GUID{Data1: 0xCC53AF8C, Data2: 0x74CD, Data3: 0x11DF,
	Data4: [8]byte{0xAC, 0x3E, 0x00, 0x24, 0x81, 0x1E, 0x38, 0x92}}

type FabricRuntime struct {
	IFabricRuntime
}

func NewFabricRuntime(pUnk *win32.IUnknown, addRef bool, scoped bool) *FabricRuntime {
	if pUnk == nil {
		return nil
	}
	p := (*FabricRuntime)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewFabricRuntimeInstance(scoped bool) (*FabricRuntime, error) {
	var p *win32.IUnknown
	hr := win32.CoCreateInstance(&CLSID_FabricRuntime, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_IFabricRuntime, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewFabricRuntime(p, false, scoped), nil
}
