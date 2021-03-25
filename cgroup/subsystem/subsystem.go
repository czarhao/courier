package subsystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Subsystem interface {
	Name() string
	// Create 根据
	Create(config map[string]string, cgroup string) error
	Apply(config map[string]string, cgroup string, pid int) error
	Remove(cgroup string) error
	Status(cgroup string) (string, error)
	IsEqual(value string, config map[string]string) bool
	IsDefault(value string) bool
}

// name: 这个 Subsystem 的名字
// file: 这个 Subsystem 具体的文件名
type basic struct {
	name, file string
}

func (b basic) Name() string {
	return b.name
}

func (b basic) Status(cgroup string) (string, error) {
	croot, err := findCgroupMountPoint(b.Name())
	if err != nil {
		return "", err
	}
	cpath := path.Join(croot, cgroup)
	if _, err = os.Stat(cpath); os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return readFile(path.Join(cpath, b.file))
}

func (b basic) Create(config map[string]string, cgroup string) error {
	if config[b.file] == "" {
		return nil
	}
	mut, err := getCgroupMountPoint(b.Name(), cgroup, true)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(mut, b.file), []byte(config[b.file]), 0644); err != nil {
		return fmt.Errorf("set subsystem %s fail %v", b.Name(), err)
	}
	return nil
}

func (b basic) IsEqual(value string, config map[string]string) bool {
	value = strings.Replace(value, " ", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	return value == config[b.file]
}

func (b basic) Apply(config map[string]string, cgroup string, pid int) error {
	if config[b.file] == "" {
		return nil
	}
	mut, err := getCgroupMountPoint(b.Name(), cgroup, false)
	if err != nil {
		return err
	}
	return writeCgroupProc(mut, pid)
}

func (b basic) Remove(cgroup string) error {
	subsystemPath, err := getCgroupMountPoint(b.Name(), cgroup, false)
	if err != nil {
		return fmt.Errorf("remove subsystem %s fail %v", b.Name(), err)
	}
	return os.RemoveAll(subsystemPath)
}

func (b basic) IsDefault(value string) bool {
	return value == ""
}
