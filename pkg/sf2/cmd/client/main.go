package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/youyuanwu/scratch-go/pkg/sf2/com/fabricclient"
)

// func main() {
// 	var c *fabricclient.IFabricApplicationManagementClient
// 	err := mssf.FabricCreateLocalClient(&fabricclient.IID_IFabricApplicationManagementClient, &c)
// 	if err.FAILED() {
// 		panic(err)
// 	}

// 	if c == nil {
// 		panic("c create failed")
// 	}
// 	//fabricclient.IFabricApplicationManagementClient
// }

// func main() {
// 	var mod = syscall.NewLazyDLL("user32.dll")
// 	var proc = mod.NewProc("MessageBoxW")
// 	var MB_YESNOCANCEL = 0x00000003

// 	ret, _, _ := proc.Call(0,
// 		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("This test is Done."))),
// 		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
// 		uintptr(MB_YESNOCANCEL))
// 	fmt.Printf("Return: %d\n", ret)

// }

// func main() {
// 	var mod = syscall.NewLazyDLL("FabricClient.dll")
// 	var proc = mod.NewProc("FabricCreateLocalClient")

// 	var c *fabricclient.IFabricApplicationManagementClient = nil

// 	ret, _, callErr := proc.Call(0,
// 		uintptr(unsafe.Pointer(&fabricclient.IID_IFabricApplicationManagementClient)),
// 		uintptr(unsafe.Pointer(&c)))
// 	if callErr != nil {
// 		panic(callErr)
// 	}
// 	fmt.Printf("Return: %d\n", ret)

// }

// TODO try cgo.

func main() {

	// ole.Initialize()
	// defer ole.Uninitialize()

	// defer com.NewScope().Leave()
	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ := syscall.GetProcAddress(kernel32, "GetModuleHandleW")
	if getModuleHandle == 0 {
		panic("no mod")
	}

	mod0, err := syscall.LoadLibrary("FabricResources.dll")
	if err != nil {
		panic(err)
	}
	if mod0 == 0 {
		panic("no mod")
	}

	mod, err := syscall.LoadLibrary("FabricCommon.dll")
	if err != nil {
		panic(err)
	}
	fabricCreateLocalClient, err := syscall.GetProcAddress(mod, "FabricGetLastErrorMessage")
	if err != nil {
		panic(err)
	}
	var c *fabricclient.IFabricStringResult = nil

	ret, _, callErr := syscall.SyscallN(fabricCreateLocalClient,
		// uintptr(unsafe.Pointer(&fabricclient.IID_IFabricApplicationManagementClient)),
		uintptr(unsafe.Pointer(&c)))
	if callErr != 0 && callErr != 126 {
		panic(callErr)
	}

	if c != nil {
		fmt.Println("c is not nil")
		st := c.Get_String() // The generated api is freeing the string. generator is wrong.
		fmt.Println("str is ", st)
	}
	fmt.Printf("Return: %d\n", ret)
}

// func main() {
// 	var ptr *fabriccommon.IFabricStringResult = nil
// 	mssf.FabricGetLastErrorMessage(&ptr)
// }
