package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 52040BD9-A78E-4308-A30E-7114E3684E76
var IID_IFabricServiceHealthResult = syscall.GUID{Data1: 0x52040BD9, Data2: 0xA78E, Data3: 0x4308,
	Data4: [8]byte{0xA3, 0x0E, 0x71, 0x14, 0xE3, 0x68, 0x4E, 0x76}}

type IFabricServiceHealthResult struct {
	win32.IUnknown
}

func NewIFabricServiceHealthResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceHealthResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceHealthResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceHealthResult) IID() *syscall.GUID {
	return &IID_IFabricServiceHealthResult
}

func (this *IFabricServiceHealthResult) Get_ServiceHealth() *FABRIC_SERVICE_HEALTH {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SERVICE_HEALTH)(unsafe.Pointer(ret))
}
