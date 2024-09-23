package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// AACE77AE-D8E1-4144-B1EE-5AC74FD54F65
var IID_IFabricEseLocalStoreSettingsResult = syscall.GUID{Data1: 0xAACE77AE, Data2: 0xD8E1, Data3: 0x4144,
	Data4: [8]byte{0xB1, 0xEE, 0x5A, 0xC7, 0x4F, 0xD5, 0x4F, 0x65}}

type IFabricEseLocalStoreSettingsResult struct {
	win32.IUnknown
}

func NewIFabricEseLocalStoreSettingsResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricEseLocalStoreSettingsResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricEseLocalStoreSettingsResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricEseLocalStoreSettingsResult) IID() *syscall.GUID {
	return &IID_IFabricEseLocalStoreSettingsResult
}

func (this *IFabricEseLocalStoreSettingsResult) Get_Settings() *FABRIC_ESE_LOCAL_STORE_SETTINGS {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_ESE_LOCAL_STORE_SETTINGS)(unsafe.Pointer(ret))
}
