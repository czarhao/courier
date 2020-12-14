package mount

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func pivotRoot(rootfs string) error {
	oldroot, err := unix.Open("/", unix.O_DIRECTORY|unix.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("get old rootfs failed, err: %v", err)
	}
	defer unix.Close(oldroot)
	newroot, err := unix.Open(rootfs, unix.O_DIRECTORY|unix.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("get new rootfs failed, err: %v", err)
	}
	defer unix.Close(newroot)

	if err := unix.Fchdir(newroot); err != nil {
		return fmt.Errorf("change to the new root failed, err: %v", err)
	}
	if err := unix.PivotRoot(".", "."); err != nil {
		return fmt.Errorf("pivot_root the new root failed, err: %v", err)
	}
	if err := unix.Fchdir(oldroot); err != nil {
		return fmt.Errorf("change to the old root failed, err: %v", err)
	}

	//
	if err := unix.Mount("", ".", "",
		unix.MS_SLAVE|unix.MS_REC, ""); err != nil {
		return fmt.Errorf("mount '.' failed, err: %v", err)
	}
	if err := unix.Unmount(".", unix.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount '.' failed, err: %v", err)
	}

	if err := unix.Chdir("/"); err != nil {
		return fmt.Errorf("switch to new root failed, err: %v", err)
	}
	return nil
}
