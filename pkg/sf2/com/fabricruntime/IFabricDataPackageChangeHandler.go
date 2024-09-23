package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 8D0A726F-BD17-4B32-807B-BE2A8024B2E0
var IID_IFabricDataPackageChangeHandler = syscall.GUID{Data1: 0x8D0A726F, Data2: 0xBD17, Data3: 0x4B32,
	Data4: [8]byte{0x80, 0x7B, 0xBE, 0x2A, 0x80, 0x24, 0xB2, 0xE0}}

type IFabricDataPackageChangeHandler struct {
	win32.IUnknown
}

type IFabricDataPackageChangeHandlerInterface interface {
	win32.IUnknownInterface
	OnPackageAdded(source *IFabricCodePackageActivationContext, dataPackage *IFabricDataPackage)
	OnPackageRemoved(source *IFabricCodePackageActivationContext, dataPackage *IFabricDataPackage)
	OnPackageModified(source *IFabricCodePackageActivationContext, previousDataPackage *IFabricDataPackage, dataPackage *IFabricDataPackage)
}

type IFabricDataPackageChangeHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricDataPackageChangeHandlerInterface
}

func (this *IFabricDataPackageChangeHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricDataPackageChangeHandlerInterface)
}

func (this *IFabricDataPackageChangeHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricDataPackageChangeHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricDataPackageChangeHandlerImpl) OnPackageAdded(source *IFabricCodePackageActivationContext, dataPackage *IFabricDataPackage) {
}
func (this *IFabricDataPackageChangeHandlerImpl) OnPackageRemoved(source *IFabricCodePackageActivationContext, dataPackage *IFabricDataPackage) {
}
func (this *IFabricDataPackageChangeHandlerImpl) OnPackageModified(source *IFabricCodePackageActivationContext, previousDataPackage *IFabricDataPackage, dataPackage *IFabricDataPackage) {
}

type IFabricDataPackageChangeHandlerVtbl struct {
	win32.IUnknownVtbl
	OnPackageAdded    uintptr
	OnPackageRemoved  uintptr
	OnPackageModified uintptr
}

type IFabricDataPackageChangeHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricDataPackageChangeHandlerComObj) impl() IFabricDataPackageChangeHandlerInterface {
	return this.Impl().(IFabricDataPackageChangeHandlerInterface)
}

func (this *IFabricDataPackageChangeHandlerComObj) OnPackageAdded(source *IFabricCodePackageActivationContext, dataPackage *IFabricDataPackage) {
	(this.impl().OnPackageAdded(source, dataPackage))
}

func (this *IFabricDataPackageChangeHandlerComObj) OnPackageRemoved(source *IFabricCodePackageActivationContext, dataPackage *IFabricDataPackage) {
	(this.impl().OnPackageRemoved(source, dataPackage))
}

func (this *IFabricDataPackageChangeHandlerComObj) OnPackageModified(source *IFabricCodePackageActivationContext, previousDataPackage *IFabricDataPackage, dataPackage *IFabricDataPackage) {
	(this.impl().OnPackageModified(source, previousDataPackage, dataPackage))
}

var _pIFabricDataPackageChangeHandlerVtbl *IFabricDataPackageChangeHandlerVtbl

func (this *IFabricDataPackageChangeHandlerComObj) BuildVtbl(lock bool) *IFabricDataPackageChangeHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricDataPackageChangeHandlerVtbl != nil {
		return _pIFabricDataPackageChangeHandlerVtbl
	}
	_pIFabricDataPackageChangeHandlerVtbl = &IFabricDataPackageChangeHandlerVtbl{
		IUnknownVtbl:      *this.IUnknownComObj.BuildVtbl(false),
		OnPackageAdded:    syscall.NewCallback((*IFabricDataPackageChangeHandlerComObj).OnPackageAdded),
		OnPackageRemoved:  syscall.NewCallback((*IFabricDataPackageChangeHandlerComObj).OnPackageRemoved),
		OnPackageModified: syscall.NewCallback((*IFabricDataPackageChangeHandlerComObj).OnPackageModified),
	}
	return _pIFabricDataPackageChangeHandlerVtbl
}

func (this *IFabricDataPackageChangeHandlerComObj) IFabricDataPackageChangeHandler() *IFabricDataPackageChangeHandler {
	return (*IFabricDataPackageChangeHandler)(unsafe.Pointer(this))
}

func (this *IFabricDataPackageChangeHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricDataPackageChangeHandlerComObj(impl IFabricDataPackageChangeHandlerInterface, scoped bool) *IFabricDataPackageChangeHandlerComObj {
	comObj := com.NewComObj[IFabricDataPackageChangeHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricDataPackageChangeHandler(impl IFabricDataPackageChangeHandlerInterface) *IFabricDataPackageChangeHandler {
	return NewIFabricDataPackageChangeHandlerComObj(impl, true).IFabricDataPackageChangeHandler()
}
