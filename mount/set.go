package mount

import (
	"courier/configs"
	"fmt"
	"os"
	"syscall"
)

const (
	procMountFlags  = syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	tmpfsMountFlags = syscall.MS_NOSUID | syscall.MS_STRICTATIME
)

func SetMount(cfg *configs.MountConfig) error {
	if cfg.Path == "" {
		pwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get pwd fail, error: %v", err)
		}
		cfg.Path = pwd
	}

	if cfg.UsePivotRoot {
		if err := pivotRoot(cfg.Path); err != nil {
			return err
		}
	} else {
		if err := chroot(cfg.Path); err != nil {
			return err
		}
	}

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
