package mssf

// var (
// 	fabricclient            syscall.Handle
// 	fabricCreateLocalClient uintptr

// 	// fabriccommon, _    = syscall.LoadLibrary("fabriccommon.dll")
// 	// fabricresources, _ = syscall.LoadLibrary("FabricResources.dll")
// )

// func init() {
// 	// Init ConfigMap here
// 	fabricclientlocal, err := syscall.LoadLibrary("fabricclient.dll")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fabricclient = fabricclientlocal
// 	fabricCreateLocalClientLocal, err := syscall.GetProcAddress(fabricclient, "FabricCreateLocalClient")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fabricCreateLocalClient = fabricCreateLocalClientLocal
// }

// func FabricCreateLocalClient[T interface{}](iid *syscall.GUID, ptr **T) com.Error {
// 	println("addr:", fabricCreateLocalClient)
// 	if ret, _, callErr := syscall.SyscallN(fabricCreateLocalClient, uintptr(unsafe.Pointer(iid)), uintptr(unsafe.Pointer(ptr))); callErr != 0 {
// 		panic(callErr)
// 	} else {
// 		e := com.Error(ret)
// 		return e
// 	}
// }
