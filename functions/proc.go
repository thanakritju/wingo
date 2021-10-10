package functions

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32DLL           = windows.NewLazyDLL("user32.dll")
	procSystemParamInfo = user32DLL.NewProc("SystemParametersInfoW")
)

func ChangeBackground(absolutePath string) {
	imagePath, _ := windows.UTF16PtrFromString(absolutePath)
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(imagePath)), 0x001A)
}
