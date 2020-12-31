package mount

import (
	"fmt"
	"syscall"
)

const (
	procMountFlags  = syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	tmpfsMountFlags = syscall.MS_NOSUID | syscall.MS_STRICTATIME
)

func mountOther() error {
	if err := syscall.Mount("proc", "/proc", "proc",
		uintptr(procMountFlags), ""); err != nil {
		return fmt.Errorf("mount /proc failed, err: %v", err)
	}
	if err := syscall.Mount("tmpfs", "/dev", "tmpfs",
		uintptr(tmpfsMountFlags), "mode=755"); err != nil {
		return fmt.Errorf("mount /dev failed, err: %v", err)
	}
	return nil
}

func umountOther() error {
	if err := syscall.Unmount("/proc", syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("mount /proc failed, err: %v", err)
	}
	if err := syscall.Unmount("/dev", syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("mount /dev failed, err: %v", err)
	}
	return nil
}
