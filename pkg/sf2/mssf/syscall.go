package mssf

// //#cgo LDFLAGS: -lFabricRuntime
// //#include <FabricCommonGo.h>
// import "C"
// import (
// 	"unsafe"

// 	"github.com/youyuanwu/scratch-go/pkg/sf2/com/fabriccommon"
// 	"github.com/zzl/go-com/com"
// )

// func FabricGetLastErrorMessage(ptr **fabriccommon.IFabricStringResult) com.Error {
// 	ec := C.FabricGetLastErrorMessage(uintptr(unsafe.Pointer(ptr)))
// 	com.Error(ec)
// }

// mkwinsyscall -output client2.go -systemdll 'C:/Program Files/Microsoft Service Fabric/bin/Fabric/Fabric.Code/FabricClient.dll'
//sys FabricCreateLocalClient(iid syscall.GUID, out *uintptr) () = FabricClient.FabricCreateLocalClient
