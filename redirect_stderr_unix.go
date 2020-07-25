// Log the panic under unix to the log file

// +build !windows,!solaris,!plan9

package logger

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func RedirectStderr(f *os.File) {
	err := unix.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}
