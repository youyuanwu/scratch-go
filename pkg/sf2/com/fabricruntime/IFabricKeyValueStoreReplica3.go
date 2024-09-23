package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// C1297172-A8AA-4096-BDCC-1ECE0C5D8C8F
var IID_IFabricKeyValueStoreReplica3 = syscall.GUID{Data1: 0xC1297172, Data2: 0xA8AA, Data3: 0x4096,
	Data4: [8]byte{0xBD, 0xCC, 0x1E, 0xCE, 0x0C, 0x5D, 0x8C, 0x8F}}

type IFabricKeyValueStoreReplica3 struct {
	IFabricKeyValueStoreReplica2
}

func NewIFabricKeyValueStoreReplica3(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreReplica3 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreReplica3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreReplica3) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreReplica3
}

func (this *IFabricKeyValueStoreReplica3) BeginBackup(backupDirectory string, backupOption int32, postBackupHandler *IFabricStorePostBackupHandler, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(backupDirectory)), uintptr(backupOption), uintptr(unsafe.Pointer(postBackupHandler)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica3) EndBackup(context *IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}
