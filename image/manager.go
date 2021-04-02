package image

import (
	"courier/configs"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//go:embed busybox.tar
var busybox []byte

var (
	ErrNotFindImage = errors.New("can not find image")
)

type Manager interface {
	Init(cfg *configs.ImageConfig) error
	CreateLayer(image string) (string, error)
	CreateWriteDir(container string) (string, error)
	Create(config *configs.OtherConfig) error
	Destroy(container string) error
}

type manager struct {
	cfg *configs.ImageConfig
	cache map[string]string
}

func NewImageManager() Manager {
	return &manager{
		cache: map[string]string{},
	}
}

// Init 创建四个文件夹
// ImageDir: 存储 image
// LayerDir: 解压 image 后的位置
// WriteDir: 各个 container 的 write 层
// ContainerDir: 用来存放具体的
func (m *manager) Init(cfg *configs.ImageConfig) error {
	err := func() error {
		if err := os.MkdirAll(cfg.ImageDir, os.ModePerm); err != nil {
			return err
		}
		if err := os.MkdirAll(cfg.LayerDir, os.ModePerm); err != nil {
			return err
		}
		if err := os.MkdirAll(cfg.WriteDir, os.ModePerm); err != nil {
			return err
		}
		return os.MkdirAll(cfg.ContainerDir, os.ModePerm)
	}()
	if err != nil {
		return fmt.Errorf("failed to init the image storage directory, err: %v", err)
	}
	m.cfg = cfg
	// 至少要有一个 busybox.tar
	exist, path := fileExist(cfg.ImageDir, "busybox.tar")
	if !exist {
		return dumpFile(path, busybox)
	}
	return nil
}

// CreateLayer 是对 ImageDir 中对 image 进行解压
// 解压到 LayerDir 中, 得到一个一个 layer
func (m *manager) CreateLayer(image string) (string, error) {
	exist, layerPath := fileExist(m.cfg.LayerDir, image)
	if exist {
		return layerPath, nil
	}
	exist, imagePath := fileExist(m.cfg.ImageDir, image)
	if !exist {
		return "", ErrNotFindImage
	}
	if err := os.MkdirAll(layerPath, os.ModePerm); err != nil {
		return "", err
	}
	if _, err := exec.Command("tar", "-zxvf", imagePath, "-C", layerPath).CombinedOutput(); err != nil {
		return "", err
	}
	return layerPath, nil
}

// CreateWriteDir 是创建一个 container 的读写层
func (m *manager) CreateWriteDir(container string) (string, error) {
	exist, path := fileExist(m.cfg.WriteDir, container)
	if !exist {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return "", err
		}
	}
	return path, nil
}

func (m *manager) CreateContainerDir(container string) (string, error) {
	exist, path := fileExist(m.cfg.ContainerDir, container)
	if !exist {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return "", err
		}
	}
	return path, nil
}

func (m manager) Create(config *configs.OtherConfig) error {
	readLayer, err := m.CreateLayer(config.Image)
	if err != nil {
		return err
	}
	writeLayer, err := m.CreateWriteDir(config.Name)
	if err != nil {
		return err
	}
	container, err := m.CreateContainerDir(config.Name)
	if err != nil {
		return err
	}

	dirs := fmt.Sprintf("dirs=%s:%s", writeLayer, readLayer)
	if _, err := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", container).CombinedOutput(); err != nil {
		return fmt.Errorf("run command for creating mount point failed, err: %v", err)
	}
	m.cache[config.Name] = container
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