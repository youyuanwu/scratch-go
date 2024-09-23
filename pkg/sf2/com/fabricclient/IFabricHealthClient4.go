package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 9F0401AF-4909-404F-8696-0A71BD753E98
var IID_IFabricHealthClient4 = syscall.GUID{Data1: 0x9F0401AF, Data2: 0x4909, Data3: 0x404F,
	Data4: [8]byte{0x86, 0x96, 0x0A, 0x71, 0xBD, 0x75, 0x3E, 0x98}}

type IFabricHealthClient4 struct {
	IFabricHealthClient3
}

func NewIFabricHealthClient4(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricHealthClient4 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricHealthClient4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricHealthClient4) IID() *syscall.GUID {
	return &IID_IFabricHealthClient4
}

func (this *IFabricHealthClient4) ReportHealth2(healthReport *FABRIC_HEALTH_REPORT, sendOptions *FABRIC_HEALTH_REPORT_SEND_OPTIONS) com.Error {
	addr := (*this.LpVtbl)[38]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthReport)), uintptr(unsafe.Pointer(sendOptions)))
	return com.Error(ret)
}
