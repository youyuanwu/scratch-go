package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 9B9105AC-D595-4F59-9C94-1FFDBF92A876
var IID_IFabricClusterManagementClient7 = syscall.GUID{Data1: 0x9B9105AC, Data2: 0xD595, Data3: 0x4F59,
	Data4: [8]byte{0x9C, 0x94, 0x1F, 0xFD, 0xBF, 0x92, 0xA8, 0x76}}

type IFabricClusterManagementClient7 struct {
	IFabricClusterManagementClient6
}

func NewIFabricClusterManagementClient7(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClusterManagementClient7 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClusterManagementClient7)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClusterManagementClient7) IID() *syscall.GUID {
	return &IID_IFabricClusterManagementClient7
}

func (this *IFabricClusterManagementClient7) BeginUpgradeConfiguration(startUpgradeDescription *FABRIC_START_UPGRADE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[47]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(startUpgradeDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient7) EndUpgradeConfiguration(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[48]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient7) BeginGetClusterConfigurationUpgradeStatus(timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[49]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient7) EndGetClusterConfigurationUpgradeStatus(context *IFabricAsyncOperationContext, result **IFabricOrchestrationUpgradeStatusResult) com.Error {
	addr := (*this.LpVtbl)[50]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient7) BeginGetClusterConfiguration(timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[51]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient7) EndGetClusterConfiguration(context *IFabricAsyncOperationContext, result **IFabricStringResult) com.Error {
	addr := (*this.LpVtbl)[52]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient7) BeginGetUpgradesPendingApproval(timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[53]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient7) EndGetUpgradesPendingApproval(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[54]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient7) BeginStartApprovedUpgrades(timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[55]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient7) EndStartApprovedUpgrades(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[56]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
