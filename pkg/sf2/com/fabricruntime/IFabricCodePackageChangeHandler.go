package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// B90D36CD-ACB5-427A-B318-3B045981D0CC
var IID_IFabricCodePackageChangeHandler = syscall.GUID{Data1: 0xB90D36CD, Data2: 0xACB5, Data3: 0x427A,
	Data4: [8]byte{0xB3, 0x18, 0x3B, 0x04, 0x59, 0x81, 0xD0, 0xCC}}

type IFabricCodePackageChangeHandler struct {
	win32.IUnknown
}

type IFabricCodePackageChangeHandlerInterface interface {
	win32.IUnknownInterface
	OnPackageAdded(source *IFabricCodePackageActivationContext, codePackage *IFabricCodePackage)
	OnPackageRemoved(source *IFabricCodePackageActivationContext, codePackage *IFabricCodePackage)
	OnPackageModified(source *IFabricCodePackageActivationContext, previousCodePackage *IFabricCodePackage, codePackage *IFabricCodePackage)
}

type IFabricCodePackageChangeHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricCodePackageChangeHandlerInterface
}

func (this *IFabricCodePackageChangeHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricCodePackageChangeHandlerInterface)
}

func (this *IFabricCodePackageChangeHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricCodePackageChangeHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricCodePackageChangeHandlerImpl) OnPackageAdded(source *IFabricCodePackageActivationContext, codePackage *IFabricCodePackage) {
}
func (this *IFabricCodePackageChangeHandlerImpl) OnPackageRemoved(source *IFabricCodePackageActivationContext, codePackage *IFabricCodePackage) {
}
func (this *IFabricCodePackageChangeHandlerImpl) OnPackageModified(source *IFabricCodePackageActivationContext, previousCodePackage *IFabricCodePackage, codePackage *IFabricCodePackage) {
}

type IFabricCodePackageChangeHandlerVtbl struct {
	win32.IUnknownVtbl
	OnPackageAdded    uintptr
	OnPackageRemoved  uintptr
	OnPackageModified uintptr
}

type IFabricCodePackageChangeHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricCodePackageChangeHandlerComObj) impl() IFabricCodePackageChangeHandlerInterface {
	return this.Impl().(IFabricCodePackageChangeHandlerInterface)
}

func (this *IFabricCodePackageChangeHandlerComObj) OnPackageAdded(source *IFabricCodePackageActivationContext, codePackage *IFabricCodePackage) {
	(this.impl().OnPackageAdded(source, codePackage))
}

func (this *IFabricCodePackageChangeHandlerComObj) OnPackageRemoved(source *IFabricCodePackageActivationContext, codePackage *IFabricCodePackage) {
	(this.impl().OnPackageRemoved(source, codePackage))
}

func (this *IFabricCodePackageChangeHandlerComObj) OnPackageModified(source *IFabricCodePackageActivationContext, previousCodePackage *IFabricCodePackage, codePackage *IFabricCodePackage) {
	(this.impl().OnPackageModified(source, previousCodePackage, codePackage))
}

var _pIFabricCodePackageChangeHandlerVtbl *IFabricCodePackageChangeHandlerVtbl

func (this *IFabricCodePackageChangeHandlerComObj) BuildVtbl(lock bool) *IFabricCodePackageChangeHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricCodePackageChangeHandlerVtbl != nil {
		return _pIFabricCodePackageChangeHandlerVtbl
	}
	_pIFabricCodePackageChangeHandlerVtbl = &IFabricCodePackageChangeHandlerVtbl{
		IUnknownVtbl:      *this.IUnknownComObj.BuildVtbl(false),
		OnPackageAdded:    syscall.NewCallback((*IFabricCodePackageChangeHandlerComObj).OnPackageAdded),
		OnPackageRemoved:  syscall.NewCallback((*IFabricCodePackageChangeHandlerComObj).OnPackageRemoved),
		OnPackageModified: syscall.NewCallback((*IFabricCodePackageChangeHandlerComObj).OnPackageModified),
	}
	return _pIFabricCodePackageChangeHandlerVtbl
}

func (this *IFabricCodePackageChangeHandlerComObj) IFabricCodePackageChangeHandler() *IFabricCodePackageChangeHandler {
	return (*IFabricCodePackageChangeHandler)(unsafe.Pointer(this))
}

func (this *IFabricCodePackageChangeHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricCodePackageChangeHandlerComObj(impl IFabricCodePackageChangeHandlerInterface, scoped bool) *IFabricCodePackageChangeHandlerComObj {
	comObj := com.NewComObj[IFabricCodePackageChangeHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricCodePackageChangeHandler(impl IFabricCodePackageChangeHandlerInterface) *IFabricCodePackageChangeHandler {
	return NewIFabricCodePackageChangeHandlerComObj(impl, true).IFabricCodePackageChangeHandler()
}
