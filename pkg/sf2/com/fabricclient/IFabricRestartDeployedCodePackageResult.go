package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// FE087DC4-7A6A-41E3-90E9-B734A4CEF41F
var IID_IFabricRestartDeployedCodePackageResult = syscall.GUID{Data1: 0xFE087DC4, Data2: 0x7A6A, Data3: 0x41E3,
	Data4: [8]byte{0x90, 0xE9, 0xB7, 0x34, 0xA4, 0xCE, 0xF4, 0x1F}}

type IFabricRestartDeployedCodePackageResult struct {
	win32.IUnknown
}

func NewIFabricRestartDeployedCodePackageResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricRestartDeployedCodePackageResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricRestartDeployedCodePackageResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricRestartDeployedCodePackageResult) IID() *syscall.GUID {
	return &IID_IFabricRestartDeployedCodePackageResult
}

func (this *IFabricRestartDeployedCodePackageResult) Get_Result() *FABRIC_DEPLOYED_CODE_PACKAGE_RESULT {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_DEPLOYED_CODE_PACKAGE_RESULT)(unsafe.Pointer(ret))
}
