package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 8952E931-B2B3-470A-B982-6B415F30DBC0
var IID_IFabricChaosReportResult = syscall.GUID{Data1: 0x8952E931, Data2: 0xB2B3, Data3: 0x470A,
	Data4: [8]byte{0xB9, 0x82, 0x6B, 0x41, 0x5F, 0x30, 0xDB, 0xC0}}

type IFabricChaosReportResult struct {
	win32.IUnknown
}

func NewIFabricChaosReportResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricChaosReportResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricChaosReportResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricChaosReportResult) IID() *syscall.GUID {
	return &IID_IFabricChaosReportResult
}

func (this *IFabricChaosReportResult) Get_ChaosReportResult() *FABRIC_CHAOS_REPORT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CHAOS_REPORT)(unsafe.Pointer(ret))
}
