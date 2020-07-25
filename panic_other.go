// +build darwin dragonfly freebsd linux nacl netbsd openbsd

package logger

import (
	"os"
	
	"golang.org/x/sys/unix"
)

// PanicFile Can be used by other packages.
// gin.DefaultErrorWriter = logger.PanicFile
var PanicFile *os.File

func InitPanicFile(fileLocation string) error {
	file, err := os.OpenFile(fileLocation,os.O_WRONLY | os.O_CREATE | os.O_SYNC, 0644)
	PanicFile = file
	if err != nil {
		println(err)
		return err
	}
	if err = unix.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
		return err
	}
	return nil
}
