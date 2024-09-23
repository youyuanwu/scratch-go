package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// A57E7740-FA33-448E-9F35-8BF802A713AA
var IID_IFabricGatewayInformationResult = syscall.GUID{Data1: 0xA57E7740, Data2: 0xFA33, Data3: 0x448E,
	Data4: [8]byte{0x9F, 0x35, 0x8B, 0xF8, 0x02, 0xA7, 0x13, 0xAA}}

type IFabricGatewayInformationResult struct {
	win32.IUnknown
}

func NewIFabricGatewayInformationResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGatewayInformationResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGatewayInformationResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGatewayInformationResult) IID() *syscall.GUID {
	return &IID_IFabricGatewayInformationResult
}

func (this *IFabricGatewayInformationResult) Get_GatewayInformation() *FABRIC_GATEWAY_INFORMATION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_GATEWAY_INFORMATION)(unsafe.Pointer(ret))
}
