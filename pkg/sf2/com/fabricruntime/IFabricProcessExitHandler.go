package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// C58D50A2-01F0-4267-BBE7-223B565C1346
var IID_IFabricProcessExitHandler = syscall.GUID{Data1: 0xC58D50A2, Data2: 0x01F0, Data3: 0x4267,
	Data4: [8]byte{0xBB, 0xE7, 0x22, 0x3B, 0x56, 0x5C, 0x13, 0x46}}

type IFabricProcessExitHandler struct {
	win32.IUnknown
}

type IFabricProcessExitHandlerInterface interface {
	win32.IUnknownInterface
	FabricProcessExited()
}

type IFabricProcessExitHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricProcessExitHandlerInterface
}

func (this *IFabricProcessExitHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricProcessExitHandlerInterface)
}

func (this *IFabricProcessExitHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricProcessExitHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricProcessExitHandlerImpl) FabricProcessExited() {
}

type IFabricProcessExitHandlerVtbl struct {
	win32.IUnknownVtbl
	FabricProcessExited uintptr
}

type IFabricProcessExitHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricProcessExitHandlerComObj) impl() IFabricProcessExitHandlerInterface {
	return this.Impl().(IFabricProcessExitHandlerInterface)
}

func (this *IFabricProcessExitHandlerComObj) FabricProcessExited() {
	(this.impl().FabricProcessExited())
}

var _pIFabricProcessExitHandlerVtbl *IFabricProcessExitHandlerVtbl

func (this *IFabricProcessExitHandlerComObj) BuildVtbl(lock bool) *IFabricProcessExitHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricProcessExitHandlerVtbl != nil {
		return _pIFabricProcessExitHandlerVtbl
	}
	_pIFabricProcessExitHandlerVtbl = &IFabricProcessExitHandlerVtbl{
		IUnknownVtbl:        *this.IUnknownComObj.BuildVtbl(false),
		FabricProcessExited: syscall.NewCallback((*IFabricProcessExitHandlerComObj).FabricProcessExited),
	}
	return _pIFabricProcessExitHandlerVtbl
}

func (this *IFabricProcessExitHandlerComObj) IFabricProcessExitHandler() *IFabricProcessExitHandler {
	return (*IFabricProcessExitHandler)(unsafe.Pointer(this))
}

func (this *IFabricProcessExitHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricProcessExitHandlerComObj(impl IFabricProcessExitHandlerInterface, scoped bool) *IFabricProcessExitHandlerComObj {
	comObj := com.NewComObj[IFabricProcessExitHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricProcessExitHandler(impl IFabricProcessExitHandlerInterface) *IFabricProcessExitHandler {
	return NewIFabricProcessExitHandlerComObj(impl, true).IFabricProcessExitHandler()
}
