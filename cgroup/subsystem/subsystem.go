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
	Create(config map[string]string, cgroup string) error
	Apply(cgroup string, pid int) error
	Remove(cgroup string) error
	Status(cgroup string) (string, error)
	IsEqual(value string, config map[string]string) bool
}

func NewSubsystem(name, file string) Subsystem {
	return &sub{
		name: name,
		file: file,
	}
}

type sub struct{
	name, file string
}

func (s *sub) Name() string {
	return s.name
}

func (s *sub) Status(cgroup string) (string, error) {
	croot, err := findCgroupMountPoint(s.Name())
	if err != nil {
		return "", err
	}
	cpath := path.Join(croot, cgroup)
	if _, err = os.Stat(cpath); os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return readFile(path.Join(cpath, s.file))
}

func (s *sub) Create(config map[string]string, cgroup string) error {
	if config[s.file] == "" {
		return nil
	}
	mut, err := getCgroupMountPoint(s.Name(), cgroup, true)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(mut, s.file), []byte(config[s.file]), 0644); err != nil {
		return fmt.Errorf("set subsystem s share fail %v", err)
	}
	return nil
}

func (s *sub) IsEqual(value string, config map[string]string) bool {
	value = strings.Replace(value, " ", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	return value == config[s.file]
}

func (s *sub) Apply(cgroup string, pid int) error {
	mut, err := getCgroupMountPoint(s.Name(), cgroup, false)
	if err != nil {
		return err
	}
	return writeCgroupProc(mut, pid)
}

func (s *sub) Remove(cgroup string) error {
	subsystemPath, err := getCgroupMountPoint(s.Name(), cgroup, false)
	if err != nil {
		return fmt.Errorf("remove subsystem %s fail %v", s.Name(), err)
	}
	return os.RemoveAll(subsystemPath)
}
