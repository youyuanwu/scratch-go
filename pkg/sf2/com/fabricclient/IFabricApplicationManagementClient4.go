package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 82C41B22-DBCB-4F7A-8D2F-F9BB94ADD446
var IID_IFabricApplicationManagementClient4 = syscall.GUID{Data1: 0x82C41B22, Data2: 0xDBCB, Data3: 0x4F7A,
	Data4: [8]byte{0x8D, 0x2F, 0xF9, 0xBB, 0x94, 0xAD, 0xD4, 0x46}}

type IFabricApplicationManagementClient4 struct {
	IFabricApplicationManagementClient3
}

func NewIFabricApplicationManagementClient4(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricApplicationManagementClient4 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricApplicationManagementClient4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricApplicationManagementClient4) IID() *syscall.GUID {
	return &IID_IFabricApplicationManagementClient4
}

func (this *IFabricApplicationManagementClient4) BeginDeployServicePackageToNode(applicationTypeName string, applicationTypeVersion string, serviceManifestName string, sharingPolicy *FABRIC_PACKAGE_SHARING_POLICY_LIST, nodeName string, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationTypeName)), uintptr(win32.StrToPointer(applicationTypeVersion)), uintptr(win32.StrToPointer(serviceManifestName)), uintptr(unsafe.Pointer(sharingPolicy)), uintptr(win32.StrToPointer(nodeName)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricApplicationManagementClient4) EndDeployServicePackageToNode(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[28]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
