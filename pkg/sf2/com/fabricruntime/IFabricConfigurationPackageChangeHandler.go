package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// C3954D48-B5EE-4FF4-9BC0-C30F6D0D3A85
var IID_IFabricConfigurationPackageChangeHandler = syscall.GUID{Data1: 0xC3954D48, Data2: 0xB5EE, Data3: 0x4FF4,
	Data4: [8]byte{0x9B, 0xC0, 0xC3, 0x0F, 0x6D, 0x0D, 0x3A, 0x85}}

type IFabricConfigurationPackageChangeHandler struct {
	win32.IUnknown
}

type IFabricConfigurationPackageChangeHandlerInterface interface {
	win32.IUnknownInterface
	OnPackageAdded(source *IFabricCodePackageActivationContext, configPackage *IFabricConfigurationPackage)
	OnPackageRemoved(source *IFabricCodePackageActivationContext, configPackage *IFabricConfigurationPackage)
	OnPackageModified(source *IFabricCodePackageActivationContext, previousConfigPackage *IFabricConfigurationPackage, configPackage *IFabricConfigurationPackage)
}

type IFabricConfigurationPackageChangeHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricConfigurationPackageChangeHandlerInterface
}

func (this *IFabricConfigurationPackageChangeHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricConfigurationPackageChangeHandlerInterface)
}

func (this *IFabricConfigurationPackageChangeHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricConfigurationPackageChangeHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricConfigurationPackageChangeHandlerImpl) OnPackageAdded(source *IFabricCodePackageActivationContext, configPackage *IFabricConfigurationPackage) {
}
func (this *IFabricConfigurationPackageChangeHandlerImpl) OnPackageRemoved(source *IFabricCodePackageActivationContext, configPackage *IFabricConfigurationPackage) {
}
func (this *IFabricConfigurationPackageChangeHandlerImpl) OnPackageModified(source *IFabricCodePackageActivationContext, previousConfigPackage *IFabricConfigurationPackage, configPackage *IFabricConfigurationPackage) {
}

type IFabricConfigurationPackageChangeHandlerVtbl struct {
	win32.IUnknownVtbl
	OnPackageAdded    uintptr
	OnPackageRemoved  uintptr
	OnPackageModified uintptr
}

type IFabricConfigurationPackageChangeHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricConfigurationPackageChangeHandlerComObj) impl() IFabricConfigurationPackageChangeHandlerInterface {
	return this.Impl().(IFabricConfigurationPackageChangeHandlerInterface)
}

func (this *IFabricConfigurationPackageChangeHandlerComObj) OnPackageAdded(source *IFabricCodePackageActivationContext, configPackage *IFabricConfigurationPackage) {
	(this.impl().OnPackageAdded(source, configPackage))
}

func (this *IFabricConfigurationPackageChangeHandlerComObj) OnPackageRemoved(source *IFabricCodePackageActivationContext, configPackage *IFabricConfigurationPackage) {
	(this.impl().OnPackageRemoved(source, configPackage))
}

func (this *IFabricConfigurationPackageChangeHandlerComObj) OnPackageModified(source *IFabricCodePackageActivationContext, previousConfigPackage *IFabricConfigurationPackage, configPackage *IFabricConfigurationPackage) {
	(this.impl().OnPackageModified(source, previousConfigPackage, configPackage))
}

var _pIFabricConfigurationPackageChangeHandlerVtbl *IFabricConfigurationPackageChangeHandlerVtbl

func (this *IFabricConfigurationPackageChangeHandlerComObj) BuildVtbl(lock bool) *IFabricConfigurationPackageChangeHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricConfigurationPackageChangeHandlerVtbl != nil {
		return _pIFabricConfigurationPackageChangeHandlerVtbl
	}
	_pIFabricConfigurationPackageChangeHandlerVtbl = &IFabricConfigurationPackageChangeHandlerVtbl{
		IUnknownVtbl:      *this.IUnknownComObj.BuildVtbl(false),
		OnPackageAdded:    syscall.NewCallback((*IFabricConfigurationPackageChangeHandlerComObj).OnPackageAdded),
		OnPackageRemoved:  syscall.NewCallback((*IFabricConfigurationPackageChangeHandlerComObj).OnPackageRemoved),
		OnPackageModified: syscall.NewCallback((*IFabricConfigurationPackageChangeHandlerComObj).OnPackageModified),
	}
	return _pIFabricConfigurationPackageChangeHandlerVtbl
}

func (this *IFabricConfigurationPackageChangeHandlerComObj) IFabricConfigurationPackageChangeHandler() *IFabricConfigurationPackageChangeHandler {
	return (*IFabricConfigurationPackageChangeHandler)(unsafe.Pointer(this))
}

func (this *IFabricConfigurationPackageChangeHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricConfigurationPackageChangeHandlerComObj(impl IFabricConfigurationPackageChangeHandlerInterface, scoped bool) *IFabricConfigurationPackageChangeHandlerComObj {
	comObj := com.NewComObj[IFabricConfigurationPackageChangeHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricConfigurationPackageChangeHandler(impl IFabricConfigurationPackageChangeHandlerInterface) *IFabricConfigurationPackageChangeHandler {
	return NewIFabricConfigurationPackageChangeHandlerComObj(impl, true).IFabricConfigurationPackageChangeHandler()
}
