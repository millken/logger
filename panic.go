// +build !windows

package logger

import (
	"os"
	"syscall"
)

// PanicFile Can be used by other packages.
// gin.DefaultErrorWriter = logger.PanicFile
var PanicFile *os.File

func InitPanicFile(fileLocation string) error {
	file, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	PanicFile = file
	if err != nil {
		println(err)
		return err
	}
	if err = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
		return err
	}
	return nil
}
