package main

import (
	"syscall"
	"unsafe"
)

func main() {
	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ := syscall.GetProcAddress(kernel32, "GetModuleHandleW")
	if getModuleHandle == 0 {
		panic("no mod")
	}

	mod0, err := syscall.LoadLibrary("D:\\code\\rs\\service-fabric-apps-rs\\target\\debug\\cpp.dll")
	if err != nil {
		panic(err)
	}
	if mod0 == 0 {
		panic("no mod")
	}

	getFabricStringsApiTable, err := syscall.GetProcAddress(mod0, "GetFabricStringsApiTable")
	if err != nil {
		panic(err)
	}
	//var c *fabricclient.IFabricStringResult = nil
	var api uintptr = 0

	// only relavent return is ret, the hresult
	ret, _, _ := syscall.SyscallN(getFabricStringsApiTable,
		// uintptr(unsafe.Pointer(&fabricclient.IID_IFabricApplicationManagementClient)),
		uintptr(unsafe.Pointer(&api)))

	println("ret is ", ret)
}
