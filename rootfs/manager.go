package rootfs

import (
	"courier/configs"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

const Busybox = "busybox.tar"

//go:embed rootfs/busybox.tar
var busybox []byte

type Manager interface {
	InitEnv(config *configs.RootfsConfig) error
	Create(container string, config *configs.RootfsConfig) error
	Destroy(container string) error
}

type manager struct {
	cache map[string]string
}

func NewManager() Manager {
	return manager{
		cache: map[string]string{},
	}
}

func (m manager) InitEnv(config *configs.RootfsConfig) error {
	if err := m.initImageStorage(config.GetImageStorage()); err != nil {
		return err
	}
	if err := os.MkdirAll(config.GetUnzipPath(), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (m manager) Create(container string, config *configs.RootfsConfig) error {
	dirs := fmt.Sprintf("dirs=%s:%s", config.WriteLayer, config.ReadLayer)
	if _, err := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", config.BaseDir).CombinedOutput(); err != nil {
		return fmt.Errorf("run command for creating mount point failed, err: %v", err)
	}
	m.cache[container] = config.BaseDir
	return nil
}

func (m manager) Destroy(container string) error {
	path, ok := m.cache[container]
	if !ok {
		return fmt.Errorf("not find container %s", container)
	}
	if err := syscall.Unmount(path, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("mount %s failed, err: %v", path, err)
	}
	return nil
}

func (m manager) initImageStorage(storage string) error {
	if err := os.MkdirAll(storage, os.ModePerm); err != nil {
		return err
	}

	busyboxPath := storage + Busybox
	if _, err := os.Stat(busyboxPath); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		return dumpFile(busyboxPath, busybox)
	}
	return nil
}
