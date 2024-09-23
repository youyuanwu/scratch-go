package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 3B93F0D9-C0A9-4DF5-9B09-B2365DE89D84
var IID_IFabricChaosScheduleDescriptionResult = syscall.GUID{Data1: 0x3B93F0D9, Data2: 0xC0A9, Data3: 0x4DF5,
	Data4: [8]byte{0x9B, 0x09, 0xB2, 0x36, 0x5D, 0xE8, 0x9D, 0x84}}

type IFabricChaosScheduleDescriptionResult struct {
	win32.IUnknown
}

func NewIFabricChaosScheduleDescriptionResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricChaosScheduleDescriptionResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricChaosScheduleDescriptionResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricChaosScheduleDescriptionResult) IID() *syscall.GUID {
	return &IID_IFabricChaosScheduleDescriptionResult
}

func (this *IFabricChaosScheduleDescriptionResult) Get_ChaosScheduleDescriptionResult() *FABRIC_CHAOS_SCHEDULE_DESCRIPTION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_CHAOS_SCHEDULE_DESCRIPTION)(unsafe.Pointer(ret))
}
