package rootfs

import (
	"courier/configs"
	"fmt"
	"os/exec"
)

type Manager interface {
	Create(container string, config *configs.MountConfig) error
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

func (m manager) Create(container string, config *configs.MountConfig) error {
	dirs := fmt.Sprintf("dirs=%s:%s", config.WriteLayer, config.ReadLayer)
	if _, err := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", config.Path).CombinedOutput(); err != nil {
		return fmt.Errorf("run command for creating mount point failed, err: %v", err)
	}
	m.cache[container] = config.Path
	return nil
}

func (m manager) Destroy(container string) error {
	path, ok := m.cache[container]
	if !ok {
		return fmt.Errorf("not find container %s", container)
	}
	_, err := exec.Command("umount", path).CombinedOutput()
	if err != nil {
		return fmt.Errorf("unmount %s point failed, err: %v", path, err)
	}
	return nil
}
