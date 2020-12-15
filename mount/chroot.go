package mount

import (
	"fmt"
	"syscall"
)

func chroot(path string) error {
	if err := syscall.Chroot(path); err != nil {
		return fmt.Errorf("call chroot failed, err: %v", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chdir failed, err: ")
	}
	return nil
}
