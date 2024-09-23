package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 3D00D0BE-7014-41DA-9C5B-0A9EF46E2A43
var IID_IFabricHealthClient = syscall.GUID{Data1: 0x3D00D0BE, Data2: 0x7014, Data3: 0x41DA,
	Data4: [8]byte{0x9C, 0x5B, 0x0A, 0x9E, 0xF4, 0x6E, 0x2A, 0x43}}

type IFabricHealthClient struct {
	win32.IUnknown
}

func NewIFabricHealthClient(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricHealthClient {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricHealthClient)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricHealthClient) IID() *syscall.GUID {
	return &IID_IFabricHealthClient
}

func (this *IFabricHealthClient) ReportHealth(healthReport *FABRIC_HEALTH_REPORT) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthReport)))
	return com.Error(ret)
}

func (this *IFabricHealthClient) BeginGetClusterHealth(healthPolicy *FABRIC_CLUSTER_HEALTH_POLICY, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(healthPolicy)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricHealthClient) EndGetClusterHealth(context *IFabricAsyncOperationContext, result **IFabricClusterHealthResult) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricHealthClient) BeginGetNodeHealth(nodeName string, healthPolicy *FABRIC_CLUSTER_HEALTH_POLICY, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(nodeName)), uintptr(unsafe.Pointer(healthPolicy)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricHealthClient) EndGetNodeHealth(context *IFabricAsyncOperationContext, result **IFabricNodeHealthResult) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricHealthClient) BeginGetApplicationHealth(applicationName string, healthPolicy *FABRIC_APPLICATION_HEALTH_POLICY, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationName)), uintptr(unsafe.Pointer(healthPolicy)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricHealthClient) EndGetApplicationHealth(context *IFabricAsyncOperationContext, result **IFabricApplicationHealthResult) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricHealthClient) BeginGetServiceHealth(serviceName string, healthPolicy *FABRIC_APPLICATION_HEALTH_POLICY, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(serviceName)), uintptr(unsafe.Pointer(healthPolicy)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricHealthClient) EndGetServiceHealth(context *IFabricAsyncOperationContext, result **IFabricServiceHealthResult) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricHealthClient) BeginGetPartitionHealth(partitionId syscall.GUID, healthPolicy *FABRIC_APPLICATION_HEALTH_POLICY, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&partitionId)), uintptr(unsafe.Pointer(healthPolicy)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricHealthClient) EndGetPartitionHealth(context *IFabricAsyncOperationContext, result **IFabricPartitionHealthResult) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricHealthClient) BeginGetReplicaHealth(partitionId syscall.GUID, replicaId int64, healthPolicy *FABRIC_APPLICATION_HEALTH_POLICY, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&partitionId)), uintptr(replicaId), uintptr(unsafe.Pointer(healthPolicy)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricHealthClient) EndGetReplicaHealth(context *IFabricAsyncOperationContext, result **IFabricReplicaHealthResult) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricHealthClient) BeginGetDeployedApplicationHealth(applicationName string, nodeName string, healthPolicy *FABRIC_APPLICATION_HEALTH_POLICY, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationName)), uintptr(win32.StrToPointer(nodeName)), uintptr(unsafe.Pointer(healthPolicy)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricHealthClient) EndGetDeployedApplicationHealth(context *IFabricAsyncOperationContext, result **IFabricDeployedApplicationHealthResult) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricHealthClient) BeginGetDeployedServicePackageHealth(applicationName string, serviceManifestName string, nodeName string, healthPolicy *FABRIC_APPLICATION_HEALTH_POLICY, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(applicationName)), uintptr(win32.StrToPointer(serviceManifestName)), uintptr(win32.StrToPointer(nodeName)), uintptr(unsafe.Pointer(healthPolicy)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricHealthClient) EndGetDeployedServicePackageHealth(context *IFabricAsyncOperationContext, result **IFabricDeployedServicePackageHealthResult) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
