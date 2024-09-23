package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 19EE48B4-6D4D-470B-AC1E-2D3996A173C8
var IID_IFabricTransaction = syscall.GUID{Data1: 0x19EE48B4, Data2: 0x6D4D, Data3: 0x470B,
	Data4: [8]byte{0xAC, 0x1E, 0x2D, 0x39, 0x96, 0xA1, 0x73, 0xC8}}

type IFabricTransaction struct {
	IFabricTransactionBase
}

func NewIFabricTransaction(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricTransaction {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricTransaction)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricTransaction) IID() *syscall.GUID {
	return &IID_IFabricTransaction
}

func (this *IFabricTransaction) BeginCommit(timeoutMilliseconds uint32, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(timeoutMilliseconds), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(context)))
	com.AddToScope(context)
	return com.Error(ret)
}

func (this *IFabricTransaction) EndCommit(context *IFabricAsyncOperationContext, commitSequenceNumber *int64) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(commitSequenceNumber)))
	return com.Error(ret)
}

func (this *IFabricTransaction) Rollback() {
	addr := (*this.LpVtbl)[7]
	_, _, _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
}
