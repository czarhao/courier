package mount

import (
	"fmt"
	"syscall"
)

func pivotRoot(rootfs string) error {
	oldroot, err := syscall.Open("/", syscall.O_DIRECTORY|syscall.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("get old rootfs failed, err: %v", err)
	}
	defer syscall.Close(oldroot)
	newroot, err := syscall.Open(rootfs, syscall.O_DIRECTORY|syscall.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("get new rootfs failed, err: %v", err)
	}
	defer syscall.Close(newroot)

	if err := syscall.Fchdir(newroot); err != nil {
		return fmt.Errorf("change to the new root failed, err: %v", err)
	}
	if err := syscall.PivotRoot(".", "."); err != nil {
		return fmt.Errorf("pivot_root the new root failed, err: %v", err)
	}
	if err := syscall.Fchdir(oldroot); err != nil {
		return fmt.Errorf("change to the old root failed, err: %v", err)
	}

	if err := syscall.Mount("", ".", "",
		syscall.MS_SLAVE|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("mount '.' failed, err: %v", err)
	}
	if err := syscall.Unmount(".", syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount '.' failed, err: %v", err)
	}

	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("switch to new root failed, err: %v", err)
	}
	return nil
}
