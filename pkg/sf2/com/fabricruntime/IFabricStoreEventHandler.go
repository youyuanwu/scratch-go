package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 220E6DA4-985B-4DEE-8FE9-77521B838795
var IID_IFabricStoreEventHandler = syscall.GUID{Data1: 0x220E6DA4, Data2: 0x985B, Data3: 0x4DEE,
	Data4: [8]byte{0x8F, 0xE9, 0x77, 0x52, 0x1B, 0x83, 0x87, 0x95}}

type IFabricStoreEventHandler struct {
	win32.IUnknown
}

type IFabricStoreEventHandlerInterface interface {
	win32.IUnknownInterface
	OnDataLoss()
}

type IFabricStoreEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricStoreEventHandlerInterface
}

func (this *IFabricStoreEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricStoreEventHandlerInterface)
}

func (this *IFabricStoreEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricStoreEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricStoreEventHandlerImpl) OnDataLoss() {
}

type IFabricStoreEventHandlerVtbl struct {
	win32.IUnknownVtbl
	OnDataLoss uintptr
}

type IFabricStoreEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricStoreEventHandlerComObj) impl() IFabricStoreEventHandlerInterface {
	return this.Impl().(IFabricStoreEventHandlerInterface)
}

func (this *IFabricStoreEventHandlerComObj) OnDataLoss() {
	(this.impl().OnDataLoss())
}

var _pIFabricStoreEventHandlerVtbl *IFabricStoreEventHandlerVtbl

func (this *IFabricStoreEventHandlerComObj) BuildVtbl(lock bool) *IFabricStoreEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricStoreEventHandlerVtbl != nil {
		return _pIFabricStoreEventHandlerVtbl
	}
	_pIFabricStoreEventHandlerVtbl = &IFabricStoreEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		OnDataLoss:   syscall.NewCallback((*IFabricStoreEventHandlerComObj).OnDataLoss),
	}
	return _pIFabricStoreEventHandlerVtbl
}

func (this *IFabricStoreEventHandlerComObj) IFabricStoreEventHandler() *IFabricStoreEventHandler {
	return (*IFabricStoreEventHandler)(unsafe.Pointer(this))
}

func (this *IFabricStoreEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricStoreEventHandlerComObj(impl IFabricStoreEventHandlerInterface, scoped bool) *IFabricStoreEventHandlerComObj {
	comObj := com.NewComObj[IFabricStoreEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricStoreEventHandler(impl IFabricStoreEventHandlerInterface) *IFabricStoreEventHandler {
	return NewIFabricStoreEventHandlerComObj(impl, true).IFabricStoreEventHandler()
}
