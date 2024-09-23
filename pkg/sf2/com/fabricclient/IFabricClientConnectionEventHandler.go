package fabricclient

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 2BD21F94-D962-4BB4-84B8-5A4B3E9D4D4D
var IID_IFabricClientConnectionEventHandler = syscall.GUID{Data1: 0x2BD21F94, Data2: 0xD962, Data3: 0x4BB4,
	Data4: [8]byte{0x84, 0xB8, 0x5A, 0x4B, 0x3E, 0x9D, 0x4D, 0x4D}}

type IFabricClientConnectionEventHandler struct {
	win32.IUnknown
}

type IFabricClientConnectionEventHandlerInterface interface {
	win32.IUnknownInterface
	OnConnected(__MIDL__IFabricClientConnectionEventHandler0000 *IFabricGatewayInformationResult) com.Error
	OnDisconnected(__MIDL__IFabricClientConnectionEventHandler0001 *IFabricGatewayInformationResult) com.Error
}

type IFabricClientConnectionEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricClientConnectionEventHandlerInterface
}

func (this *IFabricClientConnectionEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricClientConnectionEventHandlerInterface)
}

func (this *IFabricClientConnectionEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricClientConnectionEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricClientConnectionEventHandlerImpl) OnConnected(__MIDL__IFabricClientConnectionEventHandler0000 *IFabricGatewayInformationResult) com.Error {
	var ret com.Error
	return ret
}
func (this *IFabricClientConnectionEventHandlerImpl) OnDisconnected(__MIDL__IFabricClientConnectionEventHandler0001 *IFabricGatewayInformationResult) com.Error {
	var ret com.Error
	return ret
}

type IFabricClientConnectionEventHandlerVtbl struct {
	win32.IUnknownVtbl
	OnConnected    uintptr
	OnDisconnected uintptr
}

type IFabricClientConnectionEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricClientConnectionEventHandlerComObj) impl() IFabricClientConnectionEventHandlerInterface {
	return this.Impl().(IFabricClientConnectionEventHandlerInterface)
}

func (this *IFabricClientConnectionEventHandlerComObj) OnConnected(__MIDL__IFabricClientConnectionEventHandler0000 *IFabricGatewayInformationResult) uintptr {
	return (uintptr)(this.impl().OnConnected(__MIDL__IFabricClientConnectionEventHandler0000))
}

func (this *IFabricClientConnectionEventHandlerComObj) OnDisconnected(__MIDL__IFabricClientConnectionEventHandler0001 *IFabricGatewayInformationResult) uintptr {
	return (uintptr)(this.impl().OnDisconnected(__MIDL__IFabricClientConnectionEventHandler0001))
}

var _pIFabricClientConnectionEventHandlerVtbl *IFabricClientConnectionEventHandlerVtbl

func (this *IFabricClientConnectionEventHandlerComObj) BuildVtbl(lock bool) *IFabricClientConnectionEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricClientConnectionEventHandlerVtbl != nil {
		return _pIFabricClientConnectionEventHandlerVtbl
	}
	_pIFabricClientConnectionEventHandlerVtbl = &IFabricClientConnectionEventHandlerVtbl{
		IUnknownVtbl:   *this.IUnknownComObj.BuildVtbl(false),
		OnConnected:    syscall.NewCallback((*IFabricClientConnectionEventHandlerComObj).OnConnected),
		OnDisconnected: syscall.NewCallback((*IFabricClientConnectionEventHandlerComObj).OnDisconnected),
	}
	return _pIFabricClientConnectionEventHandlerVtbl
}

func (this *IFabricClientConnectionEventHandlerComObj) IFabricClientConnectionEventHandler() *IFabricClientConnectionEventHandler {
	return (*IFabricClientConnectionEventHandler)(unsafe.Pointer(this))
}

func (this *IFabricClientConnectionEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricClientConnectionEventHandlerComObj(impl IFabricClientConnectionEventHandlerInterface, scoped bool) *IFabricClientConnectionEventHandlerComObj {
	comObj := com.NewComObj[IFabricClientConnectionEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricClientConnectionEventHandler(impl IFabricClientConnectionEventHandlerInterface) *IFabricClientConnectionEventHandler {
	return NewIFabricClientConnectionEventHandlerComObj(impl, true).IFabricClientConnectionEventHandler()
}
