package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// BB8F69DE-F667-4FAB-820D-274CF4303AB4
var IID_IFabricSecretReferencesResult = syscall.GUID{Data1: 0xBB8F69DE, Data2: 0xF667, Data3: 0x4FAB,
	Data4: [8]byte{0x82, 0x0D, 0x27, 0x4C, 0xF4, 0x30, 0x3A, 0xB4}}

type IFabricSecretReferencesResult struct {
	win32.IUnknown
}

func NewIFabricSecretReferencesResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricSecretReferencesResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricSecretReferencesResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricSecretReferencesResult) IID() *syscall.GUID {
	return &IID_IFabricSecretReferencesResult
}

func (this *IFabricSecretReferencesResult) Get_SecretReferences() *FABRIC_SECRET_REFERENCE_LIST {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_SECRET_REFERENCE_LIST)(unsafe.Pointer(ret))
}
