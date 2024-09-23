package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 899E0CA8-16DF-458E-8915-D0307B4AB101
var IID_IFabricCodePackageEventHandler = syscall.GUID{Data1: 0x899E0CA8, Data2: 0x16DF, Data3: 0x458E,
	Data4: [8]byte{0x89, 0x15, 0xD0, 0x30, 0x7B, 0x4A, 0xB1, 0x01}}

type IFabricCodePackageEventHandler struct {
	win32.IUnknown
}

type IFabricCodePackageEventHandlerInterface interface {
	win32.IUnknownInterface
	OnCodePackageEvent(source *IFabricCodePackageActivator, eventDesc *FABRIC_CODE_PACKAGE_EVENT_DESCRIPTION)
}

type IFabricCodePackageEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricCodePackageEventHandlerInterface
}

func (this *IFabricCodePackageEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricCodePackageEventHandlerInterface)
}

func (this *IFabricCodePackageEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricCodePackageEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricCodePackageEventHandlerImpl) OnCodePackageEvent(source *IFabricCodePackageActivator, eventDesc *FABRIC_CODE_PACKAGE_EVENT_DESCRIPTION) {
}

type IFabricCodePackageEventHandlerVtbl struct {
	win32.IUnknownVtbl
	OnCodePackageEvent uintptr
}

type IFabricCodePackageEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricCodePackageEventHandlerComObj) impl() IFabricCodePackageEventHandlerInterface {
	return this.Impl().(IFabricCodePackageEventHandlerInterface)
}

func (this *IFabricCodePackageEventHandlerComObj) OnCodePackageEvent(source *IFabricCodePackageActivator, eventDesc *FABRIC_CODE_PACKAGE_EVENT_DESCRIPTION) {
	(this.impl().OnCodePackageEvent(source, eventDesc))
}

var _pIFabricCodePackageEventHandlerVtbl *IFabricCodePackageEventHandlerVtbl

func (this *IFabricCodePackageEventHandlerComObj) BuildVtbl(lock bool) *IFabricCodePackageEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricCodePackageEventHandlerVtbl != nil {
		return _pIFabricCodePackageEventHandlerVtbl
	}
	_pIFabricCodePackageEventHandlerVtbl = &IFabricCodePackageEventHandlerVtbl{
		IUnknownVtbl:       *this.IUnknownComObj.BuildVtbl(false),
		OnCodePackageEvent: syscall.NewCallback((*IFabricCodePackageEventHandlerComObj).OnCodePackageEvent),
	}
	return _pIFabricCodePackageEventHandlerVtbl
}

func (this *IFabricCodePackageEventHandlerComObj) IFabricCodePackageEventHandler() *IFabricCodePackageEventHandler {
	return (*IFabricCodePackageEventHandler)(unsafe.Pointer(this))
}

func (this *IFabricCodePackageEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricCodePackageEventHandlerComObj(impl IFabricCodePackageEventHandlerInterface, scoped bool) *IFabricCodePackageEventHandlerComObj {
	comObj := com.NewComObj[IFabricCodePackageEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricCodePackageEventHandler(impl IFabricCodePackageEventHandlerInterface) *IFabricCodePackageEventHandler {
	return NewIFabricCodePackageEventHandlerComObj(impl, true).IFabricCodePackageEventHandler()
}
