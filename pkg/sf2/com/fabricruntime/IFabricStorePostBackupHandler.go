package fabricruntime

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 2AF2E8A6-41DF-4E32-9D2A-D73A711E652A
var IID_IFabricStorePostBackupHandler = syscall.GUID{Data1: 0x2AF2E8A6, Data2: 0x41DF, Data3: 0x4E32,
	Data4: [8]byte{0x9D, 0x2A, 0xD7, 0x3A, 0x71, 0x1E, 0x65, 0x2A}}

type IFabricStorePostBackupHandler struct {
	win32.IUnknown
}

type IFabricStorePostBackupHandlerInterface interface {
	win32.IUnknownInterface
	BeginPostBackup(info *FABRIC_STORE_BACKUP_INFO, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error
	EndPostBackup(context *IFabricAsyncOperationContext, status *int8) com.Error
}

type IFabricStorePostBackupHandlerImpl struct {
	com.IUnknownImpl
	RealObject IFabricStorePostBackupHandlerInterface
}

func (this *IFabricStorePostBackupHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(IFabricStorePostBackupHandlerInterface)
}

func (this *IFabricStorePostBackupHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_IFabricStorePostBackupHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *IFabricStorePostBackupHandlerImpl) BeginPostBackup(info *FABRIC_STORE_BACKUP_INFO, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) com.Error {
	var ret com.Error
	return ret
}
func (this *IFabricStorePostBackupHandlerImpl) EndPostBackup(context *IFabricAsyncOperationContext, status *int8) com.Error {
	var ret com.Error
	return ret
}

type IFabricStorePostBackupHandlerVtbl struct {
	win32.IUnknownVtbl
	BeginPostBackup uintptr
	EndPostBackup   uintptr
}

type IFabricStorePostBackupHandlerComObj struct {
	com.IUnknownComObj
}

func (this *IFabricStorePostBackupHandlerComObj) impl() IFabricStorePostBackupHandlerInterface {
	return this.Impl().(IFabricStorePostBackupHandlerInterface)
}

func (this *IFabricStorePostBackupHandlerComObj) BeginPostBackup(info *FABRIC_STORE_BACKUP_INFO, callback *IFabricAsyncOperationCallback, context **IFabricAsyncOperationContext) uintptr {
	return (uintptr)(this.impl().BeginPostBackup(info, callback, context))
}

func (this *IFabricStorePostBackupHandlerComObj) EndPostBackup(context *IFabricAsyncOperationContext, status *int8) uintptr {
	return (uintptr)(this.impl().EndPostBackup(context, status))
}

var _pIFabricStorePostBackupHandlerVtbl *IFabricStorePostBackupHandlerVtbl

func (this *IFabricStorePostBackupHandlerComObj) BuildVtbl(lock bool) *IFabricStorePostBackupHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pIFabricStorePostBackupHandlerVtbl != nil {
		return _pIFabricStorePostBackupHandlerVtbl
	}
	_pIFabricStorePostBackupHandlerVtbl = &IFabricStorePostBackupHandlerVtbl{
		IUnknownVtbl:    *this.IUnknownComObj.BuildVtbl(false),
		BeginPostBackup: syscall.NewCallback((*IFabricStorePostBackupHandlerComObj).BeginPostBackup),
		EndPostBackup:   syscall.NewCallback((*IFabricStorePostBackupHandlerComObj).EndPostBackup),
	}
	return _pIFabricStorePostBackupHandlerVtbl
}

func (this *IFabricStorePostBackupHandlerComObj) IFabricStorePostBackupHandler() *IFabricStorePostBackupHandler {
	return (*IFabricStorePostBackupHandler)(unsafe.Pointer(this))
}

func (this *IFabricStorePostBackupHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewIFabricStorePostBackupHandlerComObj(impl IFabricStorePostBackupHandlerInterface, scoped bool) *IFabricStorePostBackupHandlerComObj {
	comObj := com.NewComObj[IFabricStorePostBackupHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewIFabricStorePostBackupHandler(impl IFabricStorePostBackupHandlerInterface) *IFabricStorePostBackupHandler {
	return NewIFabricStorePostBackupHandlerComObj(impl, true).IFabricStorePostBackupHandler()
}
