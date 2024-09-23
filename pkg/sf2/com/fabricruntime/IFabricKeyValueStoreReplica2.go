package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// FEF805B2-5ACA-4CAA-9C51-FB3BD577A792
var IID_IFabricKeyValueStoreReplica2 = syscall.GUID{Data1: 0xFEF805B2, Data2: 0x5ACA, Data3: 0x4CAA,
	Data4: [8]byte{0x9C, 0x51, 0xFB, 0x3B, 0xD5, 0x77, 0xA7, 0x92}}

type IFabricKeyValueStoreReplica2 struct {
	IFabricKeyValueStoreReplica
}

func NewIFabricKeyValueStoreReplica2(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreReplica2 {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreReplica2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreReplica2) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreReplica2
}

func (this *IFabricKeyValueStoreReplica2) Backup(backupDirectory string) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(backupDirectory)))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica2) Restore(backupDirectory string) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(backupDirectory)))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica2) CreateTransaction2(settings *FABRIC_KEY_VALUE_STORE_TRANSACTION_SETTINGS, transaction **IFabricTransaction) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(transaction)))
	com.AddToScope(transaction)
	return com.Error(ret)
}
