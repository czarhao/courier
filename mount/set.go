package mount

import (
	"fmt"
	"os"
	"syscall"
)

const (
	procMountFlags  = syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	tmpfsMountFlags = syscall.MS_NOSUID | syscall.MS_STRICTATIME
)

func SetMount() error {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get pwd fail, error: %v", err)
	}
	if err := pivotRoot(pwd); err != nil {
		fmt.Println(err)
		if err := chroot(pwd); err != nil {
			return err
		}
	}

	//if err := syscall.Mount("proc", "/proc", "proc",
	//	uintptr(procMountFlags), ""); err != nil {
	//	return fmt.Errorf("mount /proc failed, err: %v", err)
	//}
	//if err := syscall.Mount("tmpfs", "/dev", "tmpfs",
	//	uintptr(tmpfsMountFlags), "mode=755"); err != nil {
	//	return fmt.Errorf("mount /dev failed, err: %v", err)
	//}
	return nil
}
