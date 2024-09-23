package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 7D124A7D-258E-49F2-A9B0-E800406103FB
var IID_IFabricSecondaryEventHandler = syscall.GUID{Data1: 0x7D124A7D, Data2: 0x258E, Data3: 0x49F2,
	Data4: [8]byte{0xA9, 0xB0, 0xE8, 0x00, 0x40, 0x61, 0x03, 0xFB}}

type IFabricSecondaryEventHandler struct {
	win32.IUnknown
}

type IFabricSecondaryEventHandlerInterface interface {
	win32.IUnknownInterface
	OnCopyComplete(enumerator *IFabricKeyValueStoreEnumerator) com.Error
	OnReplicationOperation(enumerator *IFabricKeyValueStoreNotificationEnumerator) com.Error
}

type IFabricSecondaryEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricSecondaryEventHandlerInterface
}

func (this *IFabricSecondaryEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricSecondaryEventHandlerInterface)
}

func (this *IFabricSecondaryEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricSecondaryEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricSecondaryEventHandlerImpl) OnCopyComplete(enumerator *IFabricKeyValueStoreEnumerator) com.Error {
	var ret com.Error
	return ret
}
func (this *IFabricSecondaryEventHandlerImpl) OnReplicationOperation(enumerator *IFabricKeyValueStoreNotificationEnumerator) com.Error {
	var ret com.Error
	return ret
}

type IFabricSecondaryEventHandlerVtbl struct {
	win32.IUnknownVtbl
	OnCopyComplete         uintptr
	OnReplicationOperation uintptr
}

type IFabricSecondaryEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricSecondaryEventHandlerComObj) impl() IFabricSecondaryEventHandlerInterface {
	return this.Impl().(IFabricSecondaryEventHandlerInterface)
}

func (this *IFabricSecondaryEventHandlerComObj) OnCopyComplete(enumerator *IFabricKeyValueStoreEnumerator) uintptr {
	return (uintptr)(this.impl().OnCopyComplete(enumerator))
}

func (this *IFabricSecondaryEventHandlerComObj) OnReplicationOperation(enumerator *IFabricKeyValueStoreNotificationEnumerator) uintptr {
	return (uintptr)(this.impl().OnReplicationOperation(enumerator))
}

var _pIFabricSecondaryEventHandlerVtbl *IFabricSecondaryEventHandlerVtbl

func (this *IFabricSecondaryEventHandlerComObj) BuildVtbl(lock bool) *IFabricSecondaryEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricSecondaryEventHandlerVtbl != nil {
		return _pIFabricSecondaryEventHandlerVtbl
	}
	_pIFabricSecondaryEventHandlerVtbl = &IFabricSecondaryEventHandlerVtbl{
		IUnknownVtbl:           *this.IUnknownComObj.BuildVtbl(false),
		OnCopyComplete:         syscall.NewCallback((*IFabricSecondaryEventHandlerComObj).OnCopyComplete),
		OnReplicationOperation: syscall.NewCallback((*IFabricSecondaryEventHandlerComObj).OnReplicationOperation),
	}
	return _pIFabricSecondaryEventHandlerVtbl
}

func (this *IFabricSecondaryEventHandlerComObj) IFabricSecondaryEventHandler() *IFabricSecondaryEventHandler {
	return (*IFabricSecondaryEventHandler)(unsafe.Pointer(this))
}

func (this *IFabricSecondaryEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricSecondaryEventHandlerComObj(impl IFabricSecondaryEventHandlerInterface, scoped bool) *IFabricSecondaryEventHandlerComObj {
	comObj := com.NewComObj[IFabricSecondaryEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricSecondaryEventHandler(impl IFabricSecondaryEventHandlerInterface) *IFabricSecondaryEventHandler {
	return NewIFabricSecondaryEventHandlerComObj(impl, true).IFabricSecondaryEventHandler()
}
