package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// A04B7E9A-DAAB-45D4-8DA3-95EF3AB5DBAC
var IID_IFabricServiceNotificationEventHandler = syscall.GUID{Data1: 0xA04B7E9A, Data2: 0xDAAB, Data3: 0x45D4,
	Data4: [8]byte{0x8D, 0xA3, 0x95, 0xEF, 0x3A, 0xB5, 0xDB, 0xAC}}

type IFabricServiceNotificationEventHandler struct {
	win32.IUnknown
}

type IFabricServiceNotificationEventHandlerInterface interface {
	win32.IUnknownInterface
	OnNotification(__MIDL__IFabricServiceNotificationEventHandler0000 *IFabricServiceNotification) com.Error
}

type IFabricServiceNotificationEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricServiceNotificationEventHandlerInterface
}

func (this *IFabricServiceNotificationEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricServiceNotificationEventHandlerInterface)
}

func (this *IFabricServiceNotificationEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricServiceNotificationEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricServiceNotificationEventHandlerImpl) OnNotification(__MIDL__IFabricServiceNotificationEventHandler0000 *IFabricServiceNotification) com.Error {
	var ret com.Error
	return ret
}

type IFabricServiceNotificationEventHandlerVtbl struct {
	win32.IUnknownVtbl
	OnNotification uintptr
}

type IFabricServiceNotificationEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricServiceNotificationEventHandlerComObj) impl() IFabricServiceNotificationEventHandlerInterface {
	return this.Impl().(IFabricServiceNotificationEventHandlerInterface)
}

func (this *IFabricServiceNotificationEventHandlerComObj) OnNotification(__MIDL__IFabricServiceNotificationEventHandler0000 *IFabricServiceNotification) uintptr {
	return (uintptr)(this.impl().OnNotification(__MIDL__IFabricServiceNotificationEventHandler0000))
}

var _pIFabricServiceNotificationEventHandlerVtbl *IFabricServiceNotificationEventHandlerVtbl

func (this *IFabricServiceNotificationEventHandlerComObj) BuildVtbl(lock bool) *IFabricServiceNotificationEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricServiceNotificationEventHandlerVtbl != nil {
		return _pIFabricServiceNotificationEventHandlerVtbl
	}
	_pIFabricServiceNotificationEventHandlerVtbl = &IFabricServiceNotificationEventHandlerVtbl{
		IUnknownVtbl:   *this.IUnknownComObj.BuildVtbl(false),
		OnNotification: syscall.NewCallback((*IFabricServiceNotificationEventHandlerComObj).OnNotification),
	}
	return _pIFabricServiceNotificationEventHandlerVtbl
}

func (this *IFabricServiceNotificationEventHandlerComObj) IFabricServiceNotificationEventHandler() *IFabricServiceNotificationEventHandler {
	return (*IFabricServiceNotificationEventHandler)(unsafe.Pointer(this))
}

func (this *IFabricServiceNotificationEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricServiceNotificationEventHandlerComObj(impl IFabricServiceNotificationEventHandlerInterface, scoped bool) *IFabricServiceNotificationEventHandlerComObj {
	comObj := com.NewComObj[IFabricServiceNotificationEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricServiceNotificationEventHandler(impl IFabricServiceNotificationEventHandlerInterface) *IFabricServiceNotificationEventHandler {
	return NewIFabricServiceNotificationEventHandlerComObj(impl, true).IFabricServiceNotificationEventHandler()
}
