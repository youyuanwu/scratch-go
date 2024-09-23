package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 38C4C723-3815-49D8-BDF2-68BFB536B8C9
var IID_IFabricSecretStoreClient = syscall.GUID{Data1: 0x38C4C723, Data2: 0x3815, Data3: 0x49D8,
	Data4: [8]byte{0xBD, 0xF2, 0x68, 0xBF, 0xB5, 0x36, 0xB8, 0xC9}}

type IFabricSecretStoreClient struct {
	win32.IUnknown
}

func NewIFabricSecretStoreClient(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricSecretStoreClient {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricSecretStoreClient)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricSecretStoreClient) IID() *syscall.GUID {
	return &IID_IFabricSecretStoreClient
}

func (this *IFabricSecretStoreClient) BeginGetSecrets(secretReferences *FABRIC_SECRET_REFERENCE_LIST, includeValue int8, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(secretReferences)), uintptr(includeValue), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricSecretStoreClient) EndGetSecrets(context *IFabricAsyncOperationContext, result **IFabricSecretsResult) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricSecretStoreClient) BeginSetSecrets(secrets *FABRIC_SECRET_LIST, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(secrets)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricSecretStoreClient) EndSetSecrets(context *IFabricAsyncOperationContext, result **IFabricSecretsResult) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricSecretStoreClient) BeginRemoveSecrets(secretReferences *FABRIC_SECRET_REFERENCE_LIST, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(secretReferences)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricSecretStoreClient) EndRemoveSecrets(context *IFabricAsyncOperationContext, result **IFabricSecretReferencesResult) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricSecretStoreClient) BeginGetSecretVersions(secretReferences *FABRIC_SECRET_REFERENCE_LIST, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(secretReferences)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricSecretStoreClient) EndGetSecretVersions(context *IFabricAsyncOperationContext, result **IFabricSecretReferencesResult) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
