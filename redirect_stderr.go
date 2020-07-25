// +build !windows,!darwin,!dragonfly,!freebsd,!linux,!nacl,!netbsd,!openbsd

package logger

import (
	"log"
	"os"
)

func RedirectStderr(f *os.File) {

	log.Printf("Can't redirect stderr to file")
}
