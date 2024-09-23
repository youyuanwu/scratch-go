package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// FD0FE113-CDF8-4803-B4A0-32B1B3EF3716
var IID_IFabricResolvedServicePartitionResult = syscall.GUID{Data1: 0xFD0FE113, Data2: 0xCDF8, Data3: 0x4803,
	Data4: [8]byte{0xB4, 0xA0, 0x32, 0xB1, 0xB3, 0xEF, 0x37, 0x16}}

type IFabricResolvedServicePartitionResult struct {
	win32.IUnknown
}

func NewIFabricResolvedServicePartitionResult(pUnk *win32.IUnknown, addRef bool, scoped bool) *IFabricResolvedServicePartitionResult {
	if pUnk == nil {
		return nil
	}
	p := (*IFabricResolvedServicePartitionResult)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *IFabricResolvedServicePartitionResult) IID() *syscall.GUID {
	return &IID_IFabricResolvedServicePartitionResult
}

func (this *IFabricResolvedServicePartitionResult) Get_Partition() *FABRIC_RESOLVED_SERVICE_PARTITION {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return (*FABRIC_RESOLVED_SERVICE_PARTITION)(unsafe.Pointer(ret))
}

func (this *IFabricResolvedServicePartitionResult) GetEndpoint(endpoint **FABRIC_RESOLVED_SERVICE_ENDPOINT) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(endpoint)))
	return com.Error(ret)
}

func (this *IFabricResolvedServicePartitionResult) CompareVersion(other *IFabricResolvedServicePartitionResult, compareResult *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(other)), uintptr(unsafe.Pointer(compareResult)))
	return com.Error(ret)
}
