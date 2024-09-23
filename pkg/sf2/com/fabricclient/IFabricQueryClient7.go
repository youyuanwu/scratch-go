package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 538BAA81-BA97-46DA-95AC-E1CDD184CC74
var IID_IFabricQueryClient7 = syscall.GUID{Data1: 0x538BAA81, Data2: 0xBA97, Data3: 0x46DA,
	Data4: [8]byte{0x95, 0xAC, 0xE1, 0xCD, 0xD1, 0x84, 0xCC, 0x74}}

type IFabricQueryClient7 struct {
	IFabricQueryClient6
}

func NewIFabricQueryClient7(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricQueryClient7 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricQueryClient7)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricQueryClient7) IID() *syscall.GUID {
	return &IID_IFabricQueryClient7
}

func (this *IFabricQueryClient7) BeginGetApplicationLoadInformation(queryDescription *FABRIC_APPLICATION_LOAD_INFORMATION_QUERY_DESCRIPTION, timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[52]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(queryDescription)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricQueryClient7) EndGetApplicationLoadInformation(context *IFabricAsyncOperationContext, result **IFabricGetApplicationLoadInformationResult) com.Error {
	addr := (*this.LpVtbl)[53]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
