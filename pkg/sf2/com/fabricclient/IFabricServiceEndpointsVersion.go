package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 0A673DC5-2297-4FC5-A38F-482D29144FA5
var IID_IFabricServiceEndpointsVersion = syscall.GUID{Data1: 0x0A673DC5, Data2: 0x2297, Data3: 0x4FC5,
	Data4: [8]byte{0xA3, 0x8F, 0x48, 0x2D, 0x29, 0x14, 0x4F, 0xA5}}

type IFabricServiceEndpointsVersion struct {
	win32.IUnknown
}

func NewIFabricServiceEndpointsVersion(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricServiceEndpointsVersion {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricServiceEndpointsVersion)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricServiceEndpointsVersion) IID() *syscall.GUID {
	return &IID_IFabricServiceEndpointsVersion
}

func (this *IFabricServiceEndpointsVersion) Compare(other *IFabricServiceEndpointsVersion, compareResult *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(other)), uintptr(unsafe.Pointer(compareResult)))
	return com.Error(ret)
}
