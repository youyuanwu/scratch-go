package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

var CLSID_FabricClient = syscall.GUID{Data1: 0x75F087FA, Data2: 0x37F1, Data3: 0x4984,
	Data4: [8]byte{0xB5, 0x13, 0x72, 0x4D, 0xA5, 0xB0, 0x21, 0x97}}

type FabricClient struct {
	IFabricClientSettings2
}

func NewFabricClient(pUnk *win32.IUnknown, addRef bool, scoped bool) *FabricClient {
	if pUnk == nil {
		return nil
	}
	p := (*FabricClient)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func NewFabricClientInstance(scoped bool) (*FabricClient, error) {
	var p *win32.IUnknown
	hr := win32.CoCreateInstance(&CLSID_FabricClient, nil,
		win32.CLSCTX_INPROC_SERVER|win32.CLSCTX_LOCAL_SERVER,
		&IID_IFabricClientSettings2, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return nil, com.NewError(hr)
	}
	return NewFabricClient(p, false, scoped), nil
}
