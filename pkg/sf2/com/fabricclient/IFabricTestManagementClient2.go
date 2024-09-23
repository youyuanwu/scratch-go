package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 1222B1FF-AE51-43B3-BBDF-439E7F61CA1A
var IID_IFabricTestManagementClient2 = syscall.GUID{Data1: 0x1222B1FF, Data2: 0xAE51, Data3: 0x43B3,
	Data4: [8]byte{0xBB, 0xDF, 0x43, 0x9E, 0x7F, 0x61, 0xCA, 0x1A}}

type IFabricTestManagementClient2 struct {
	IFabricTestManagementClient
}

func NewIFabricTestManagementClient2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricTestManagementClient2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricTestManagementClient2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricTestManagementClient2) IID() *syscall.GUID {
	return &IID_IFabricTestManagementClient2
}

func (this *IFabricTestManagementClient2) BeginStartChaos(restartPartitionDescription *FABRIC_START_CHAOS_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(restartPartitionDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricTestManagementClient2) EndStartChaos(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricTestManagementClient2) BeginStopChaos(timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricTestManagementClient2) EndStopChaos(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

func (this *IFabricTestManagementClient2) BeginGetChaosReport(getChaosReportDescription *FABRIC_GET_CHAOS_REPORT_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(getChaosReportDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricTestManagementClient2) EndGetChaosReport(context *IFabricAsyncOperationContext, result **IFabricChaosReportResult) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
