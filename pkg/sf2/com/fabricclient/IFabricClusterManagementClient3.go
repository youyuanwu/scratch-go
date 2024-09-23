package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// C3001D74-92B6-44CB-AC2F-2FFC4A56287C
var IID_IFabricClusterManagementClient3 = syscall.GUID{Data1: 0xC3001D74, Data2: 0x92B6, Data3: 0x44CB,
	Data4: [8]byte{0xAC, 0x2F, 0x2F, 0xFC, 0x4A, 0x56, 0x28, 0x7C}}

type IFabricClusterManagementClient3 struct {
	IFabricClusterManagementClient2
}

func NewIFabricClusterManagementClient3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricClusterManagementClient3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricClusterManagementClient3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricClusterManagementClient3) IID() *syscall.GUID {
	return &IID_IFabricClusterManagementClient3
}

func (this *IFabricClusterManagementClient3) BeginUpdateFabricUpgrade(description *FABRIC_UPGRADE_UPDATE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[31]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient3) EndUpdateFabricUpgrade(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[32]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient3) BeginStopNode(stopNodeDescription *FABRIC_STOP_NODE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[33]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stopNodeDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient3) EndStopNode(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[34]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient3) BeginRestartNode(restartNodeDescription *FABRIC_RESTART_NODE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[35]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(restartNodeDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient3) EndRestartNode(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[36]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient3) BeginStartNode(startNodeDescription *FABRIC_START_NODE_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[37]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(startNodeDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient3) EndStartNode(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[38]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient3) CopyClusterPackage(imageStoreConnectionString string, clusterManifestPath string, clusterManifestPathInImageStore string, codePackagePath string, codePackagePathInImageStore string) com.Error {
	addr := (*this.LpVtbl)[39]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(imageStoreConnectionString)), uintptr(win32.StrToPointer(clusterManifestPath)), uintptr(win32.StrToPointer(clusterManifestPathInImageStore)), uintptr(win32.StrToPointer(codePackagePath)), uintptr(win32.StrToPointer(codePackagePathInImageStore)))
	return com.Error(ret)
}

func (this *IFabricClusterManagementClient3) RemoveClusterPackage(imageStoreConnectionString string, clusterManifestPathInImageStore string, codePackagePathInImageStore string) com.Error {
	addr := (*this.LpVtbl)[40]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(imageStoreConnectionString)), uintptr(win32.StrToPointer(clusterManifestPathInImageStore)), uintptr(win32.StrToPointer(codePackagePathInImageStore)))
	return com.Error(ret)
}
