package cgroup

import (
	"courier/cgroup/subsystem"
	"courier/configs"
	"fmt"
)

type Manager interface {
	Apply(pid int, name string) error
	Create(config *configs.CgroupConfig, name string) error
	Destroy(name string) error
	GetStat(name string) (map[string]string, error)
}

type manager struct {
	subs []subsystem.Subsystem
}

func NewManager() *manager {
	return &manager{subs: []subsystem.Subsystem{
		subsystem.NewSubsystem("cpu", "cpu.shares"),
	}}
}

func (m *manager) Create(config *configs.CgroupConfig, name string) error {
	if config == nil {
		return fmt.Errorf("not set config! ")
	}
	configMap := config2map(config)
	for _, sub := range m.subs {
		status, err := sub.Status(name)
		if err != nil {
			return fmt.Errorf("create cgroup fail, err: %v", err)
		}
		if len(status) != 0 && !sub.IsEqual(status, configMap) {
			return fmt.Errorf("cgroup: %s is exited, but is not equaled our config", name)
		}
		if err := sub.Create(configMap, name); err != nil {
			return fmt.Errorf("create cgroup: %s subsystem: %s failed! err: %v", name, sub.Name(), err)
		}
	}
	return nil
}

func (m *manager) Apply(pid int, name string) error {
	for _, sub := range m.subs {
		if err := sub.Apply(name, pid); err != nil {
			return err
		}
	}
	return nil
}

func (m *manager) Destroy(name string) error {
	for _, sub := range m.subs {
		if err := sub.Remove(name); err != nil {
			return err
		}
	}
	return nil
}

func (m *manager) GetStat(name string) (map[string]string, error) {
	smap := map[string]string{}
	for _, sub := range m.subs {
		status, err := sub.Status(name)
		if err != nil {
			return nil, err
		}
		if len(status) == 0 {
			continue
		}
		smap[sub.Name()] = status
	}
	return smap, nil
}
