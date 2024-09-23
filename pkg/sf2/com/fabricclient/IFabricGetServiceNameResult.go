package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// B64FB70C-FE53-4CA1-B6D9-23D1150FE76C
var IID_IFabricGetServiceNameResult = syscall.GUID{Data1: 0xB64FB70C, Data2: 0xFE53, Data3: 0x4CA1,
	Data4: [8]byte{0xB6, 0xD9, 0x23, 0xD1, 0x15, 0x0F, 0xE7, 0x6C}}

type IFabricGetServiceNameResult struct {
	win32.IUnknown
}

func NewIFabricGetServiceNameResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricGetServiceNameResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricGetServiceNameResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricGetServiceNameResult) IID() *syscall.GUID {
	return &IID_IFabricGetServiceNameResult
}

func (this *IFabricGetServiceNameResult) Get_ServiceName() *FABRIC_SERVICE_NAME_QUERY_RESULT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_NAME_QUERY_RESULT)(unsafe.Pointer(ret))
}
