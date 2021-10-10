package functions

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32DLL           = windows.NewLazyDLL("user32.dll")
	procSystemParamInfo = user32DLL.NewProc("SystemParametersInfoW")
	procMessageBox      = user32DLL.NewProc("MessageBoxW")
)

func ChangeBackground(absolutePath string) {
	imagePathPtr, _ := windows.UTF16PtrFromString(absolutePath)
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(imagePathPtr)), 0x001A)
}

func MessageBox(title, message string) {
	titlePtr, _ := windows.UTF16PtrFromString(title)
	messagePtr, _ := windows.UTF16PtrFromString(message)
	procMessageBox.Call(
		0,
		uintptr(unsafe.Pointer(titlePtr)),
		uintptr(unsafe.Pointer(messagePtr)),
		0,
	)
}
