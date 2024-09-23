package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 4332EB3A-AED6-86FE-C2FA-653123DEA09B
var IID_IFabricGetNodeLoadInformationResult = syscall.GUID{Data1: 0x4332EB3A, Data2: 0xAED6, Data3: 0x86FE,
	Data4: [8]byte{0xC2, 0xFA, 0x65, 0x31, 0x23, 0xDE, 0xA0, 0x9B}}

type IFabricGetNodeLoadInformationResult struct {
	win32.IUnknown
}

func NewIFabricGetNodeLoadInformationResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetNodeLoadInformationResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetNodeLoadInformationResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetNodeLoadInformationResult) IID() *syscall.GUID {
	return &IID_IFabricGetNodeLoadInformationResult
}

func (this *IFabricGetNodeLoadInformationResult) Get_NodeLoadInformation() *FABRIC_NODE_LOAD_INFORMATION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_NODE_LOAD_INFORMATION)(unsafe.Pointer(ret))
}
