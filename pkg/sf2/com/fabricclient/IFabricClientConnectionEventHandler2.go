package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 6B5DBD26-7D7A-4A3F-B8EA-1F049105E897
var IID_IFabricClientConnectionEventHandler2 = syscall.GUID{Data1: 0x6B5DBD26, Data2: 0x7D7A, Data3: 0x4A3F,
	Data4: [8]byte{0xB8, 0xEA, 0x1F, 0x04, 0x91, 0x05, 0xE8, 0x97}}

type IFabricClientConnectionEventHandler2 struct {
	IFabricClientConnectionEventHandler
}

func NewIFabricClientConnectionEventHandler2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClientConnectionEventHandler2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClientConnectionEventHandler2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClientConnectionEventHandler2) IID() *syscall.GUID {
	return &IID_IFabricClientConnectionEventHandler2
}

func (this *IFabricClientConnectionEventHandler2) OnClaimsRetrieval(metadata *FABRIC_CLAIMS_RETRIEVAL_METADATA, token **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(metadata)), uintptr(unsafe.Pointer(token)))
	com.AddToScope(token)
	return com.Error(ret)
}
