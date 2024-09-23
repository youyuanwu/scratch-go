package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F495715D-8E03-4232-B8D6-1227B39984FC
var IID_IFabricServicePartitionResolutionChangeHandler = syscall.GUID{Data1: 0xF495715D, Data2: 0x8E03, Data3: 0x4232,
	Data4: [8]byte{0xB8, 0xD6, 0x12, 0x27, 0xB3, 0x99, 0x84, 0xFC}}

type IFabricServicePartitionResolutionChangeHandler struct {
	win32.IUnknown
}

type IFabricServicePartitionResolutionChangeHandlerInterface interface {
	win32.IUnknownInterface
	OnChange(source *IFabricServiceManagementClient, handlerId int64, partition *IFabricResolvedServicePartitionResult, error com.Error)
}

type IFabricServicePartitionResolutionChangeHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricServicePartitionResolutionChangeHandlerInterface
}

func (this *IFabricServicePartitionResolutionChangeHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricServicePartitionResolutionChangeHandlerInterface)
}

func (this *IFabricServicePartitionResolutionChangeHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricServicePartitionResolutionChangeHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricServicePartitionResolutionChangeHandlerImpl) OnChange(source *IFabricServiceManagementClient, handlerId int64, partition *IFabricResolvedServicePartitionResult, error com.Error) {
}

type IFabricServicePartitionResolutionChangeHandlerVtbl struct {
	win32.IUnknownVtbl
	OnChange uintptr
}

type IFabricServicePartitionResolutionChangeHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricServicePartitionResolutionChangeHandlerComObj) impl() IFabricServicePartitionResolutionChangeHandlerInterface {
	return this.Impl().(IFabricServicePartitionResolutionChangeHandlerInterface)
}

func (this *IFabricServicePartitionResolutionChangeHandlerComObj) OnChange(source *IFabricServiceManagementClient, handlerId int64, partition *IFabricResolvedServicePartitionResult, error com.Error) {
	(this.impl().OnChange(source, handlerId, partition, error))
}

var _pIFabricServicePartitionResolutionChangeHandlerVtbl *IFabricServicePartitionResolutionChangeHandlerVtbl

func (this *IFabricServicePartitionResolutionChangeHandlerComObj) BuildVtbl(lock bool) *IFabricServicePartitionResolutionChangeHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricServicePartitionResolutionChangeHandlerVtbl != nil {
		return _pIFabricServicePartitionResolutionChangeHandlerVtbl
	}
	_pIFabricServicePartitionResolutionChangeHandlerVtbl = &IFabricServicePartitionResolutionChangeHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		OnChange:     syscall.NewCallback((*IFabricServicePartitionResolutionChangeHandlerComObj).OnChange),
	}
	return _pIFabricServicePartitionResolutionChangeHandlerVtbl
}

func (this *IFabricServicePartitionResolutionChangeHandlerComObj) IFabricServicePartitionResolutionChangeHandler() *IFabricServicePartitionResolutionChangeHandler {
	return (*IFabricServicePartitionResolutionChangeHandler)(unsafe.Pointer(this))
}

func (this *IFabricServicePartitionResolutionChangeHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricServicePartitionResolutionChangeHandlerComObj(impl IFabricServicePartitionResolutionChangeHandlerInterface, scoped bool) *IFabricServicePartitionResolutionChangeHandlerComObj {
	comObj := com.NewComObj[IFabricServicePartitionResolutionChangeHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricServicePartitionResolutionChangeHandler(impl IFabricServicePartitionResolutionChangeHandlerInterface) *IFabricServicePartitionResolutionChangeHandler {
	return NewIFabricServicePartitionResolutionChangeHandlerComObj(impl, true).IFabricServicePartitionResolutionChangeHandler()
}
