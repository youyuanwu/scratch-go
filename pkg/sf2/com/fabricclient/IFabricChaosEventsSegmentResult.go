package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// DE148299-C48A-4540-877B-5B1DAA518D76
var IID_IFabricChaosEventsSegmentResult = syscall.GUID{Data1: 0xDE148299, Data2: 0xC48A, Data3: 0x4540,
	Data4: [8]byte{0x87, 0x7B, 0x5B, 0x1D, 0xAA, 0x51, 0x8D, 0x76}}

type IFabricChaosEventsSegmentResult struct {
	win32.IUnknown
}

func NewIFabricChaosEventsSegmentResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricChaosEventsSegmentResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricChaosEventsSegmentResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricChaosEventsSegmentResult) IID() *syscall.GUID {
	return &IID_IFabricChaosEventsSegmentResult
}

func (this *IFabricChaosEventsSegmentResult) Get_ChaosEventsSegmentResult() *FABRIC_CHAOS_EVENTS_SEGMENT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CHAOS_EVENTS_SEGMENT)(unsafe.Pointer(ret))
}
