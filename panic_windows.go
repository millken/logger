// +build windows

package logger

import (
	"os"
	"syscall"
)

// PanicFile Can be used by other packages.
// gin.DefaultErrorWriter = logger.PanicFile
var PanicFile *os.File

const (
	kernel32dll = "kernel32.dll"
)

func InitPanicFile(fileLocation string) error {
	file, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_APPEND, 0666)
	PanicFile = file
	if err != nil {
		return err
	}
	kernel32 := syscall.NewLazyDLL(kernel32dll)
	setStdHandle := kernel32.NewProc("SetStdHandle")
	sh := syscall.STD_ERROR_HANDLE
	v, _, err := setStdHandle.Call(uintptr(sh), uintptr(file.Fd()))
	if v == 0 {
		return err
	}
	return nil
}
