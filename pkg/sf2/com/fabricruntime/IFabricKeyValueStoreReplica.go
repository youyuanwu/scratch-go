package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 97DA35C4-38ED-4A2A-8F37-FBEB56382235
var IID_IFabricKeyValueStoreReplica = syscall.GUID{Data1: 0x97DA35C4, Data2: 0x38ED, Data3: 0x4A2A,
	Data4: [8]byte{0x8F, 0x37, 0xFB, 0xEB, 0x56, 0x38, 0x22, 0x35}}

type IFabricKeyValueStoreReplica struct {
	IFabricStatefulServiceReplica
}

func NewIFabricKeyValueStoreReplica(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricKeyValueStoreReplica {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricKeyValueStoreReplica)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricKeyValueStoreReplica) IID() *syscall.GUID {
	return &IID_IFabricKeyValueStoreReplica
}

func (this *IFabricKeyValueStoreReplica) GetCurrentEpoch(currentEpoch *FABRIC_EPOCH) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(currentEpoch)))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) UpdateReplicatorSettings(replicatorSettings *FABRIC_REPLICATOR_SETTINGS) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(replicatorSettings)))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) CreateTransaction(transaction **IFabricTransaction) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)))
	com.AddToScope(transaction)
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) Add(transaction *IFabricTransactionBase, key string, valueSizeInBytes int32, value *byte) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(win32.StrToPointer(key)), uintptr(valueSizeInBytes), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) Remove(transaction *IFabricTransactionBase, key string, checkSequenceNumber int64) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(win32.StrToPointer(key)), uintptr(checkSequenceNumber))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) Update(transaction *IFabricTransactionBase, key string, valueSizeInBytes int32, value *byte, checkSequenceNumber int64) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(win32.StrToPointer(key)), uintptr(valueSizeInBytes), uintptr(unsafe.Pointer(value)), uintptr(checkSequenceNumber))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) Get(transaction *IFabricTransactionBase, key string, result **IFabricKeyValueStoreItemResult) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(win32.StrToPointer(key)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) GetMetadata(transaction *IFabricTransactionBase, key string, result **IFabricKeyValueStoreItemMetadataResult) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(win32.StrToPointer(key)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) Contains(transaction *IFabricTransactionBase, key string, result *int8) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(win32.StrToPointer(key)), uintptr(unsafe.Pointer(result)))
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) Enumerate(transaction *IFabricTransactionBase, result **IFabricKeyValueStoreItemEnumerator) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) EnumerateByKey(transaction *IFabricTransactionBase, keyPrefix string, result **IFabricKeyValueStoreItemEnumerator) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(win32.StrToPointer(keyPrefix)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) EnumerateMetadata(transaction *IFabricTransactionBase, result **IFabricKeyValueStoreItemMetadataEnumerator) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}

func (this *IFabricKeyValueStoreReplica) EnumerateMetadataByKey(transaction *IFabricTransactionBase, keyPrefix string, result **IFabricKeyValueStoreItemMetadataEnumerator) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transaction)), uintptr(win32.StrToPointer(keyPrefix)), uintptr(unsafe.Pointer(result)))
	com.AddToScope(result)
	return com.Error(ret)
}
