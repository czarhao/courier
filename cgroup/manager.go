package cgroup

import (
	"courier/cgroup/subsystem"
	"courier/configs"
	"fmt"
)

type Manager interface {
	// Apply 是根据 cgroup config 和进程 pid 加入到 cgroup 中
	// name 是当前这个 container 的 name
	Apply(config *configs.CgroupConfig, pid int, name string) error
	// Create 是根据 cgroup config 来创建 cgroup
	Create(config *configs.CgroupConfig, name string) error
	// Destroy 删掉这个 container 的 cgroup
	Destroy(name string) error
	// GetStat 读取这个 container 的 cgroup 配置
	GetStat(name string) (map[string]string, error)
}

// TODO 支持更多的 subsystem
type manager struct {
	subs []subsystem.Subsystem
}

func NewManager() *manager {
	return &manager{subs: []subsystem.Subsystem{
		subsystem.NewCpuPeriod(),
		subsystem.NewCpuQuota(),
		subsystem.NewCpuCpus(),
		subsystem.NewCpuMems(),
		subsystem.NewMemoryLimit(),
		subsystem.NewSwappiness(),
		subsystem.NewReadDevice(),
		subsystem.NewWriteDevice(),
	}}
}

func (m *manager) Create(config *configs.CgroupConfig, name string) error {
	if config == nil {
		return fmt.Errorf("not set config! ")
	}
	configMap := config2map(config)
	for _, sub := range m.subs {
		// 会出现 3 种情况
		// 1. 读取到了 当前 subsystem 的状态
		// 2. 读不到 subsystem 的状态：文件不存在(一般是这个)
		// 3. 发生 err
		status, err := sub.Status(name)
		if err != nil {
			return fmt.Errorf("create cgroup fail, err: %v", err)
		}
		// 判断一下是否已经设置了(文件存在与否) || 设置的结果与我们期望的是否一致
		if !sub.IsSet(configMap) || sub.IsEqual(status, configMap) {
			continue
		}
		if err := sub.Create(configMap, name); err != nil {
			return fmt.Errorf("create cgroup: %s subsystem: %s failed! err: %v", name, sub.Name(), err)
		}
	}
	return nil
}

func (m *manager) Apply(config *configs.CgroupConfig, pid int, name string) error {
	if config == nil {
		return fmt.Errorf("not set config! ")
	}
	configMap := config2map(config)
	for _, sub := range m.subs {
		if err := sub.Apply(configMap, name, pid); err != nil {
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
	stat := make(map[string]string, len(m.subs))
	for _, sub := range m.subs {
		status, err := sub.Status(name)
		if err != nil {
			return nil, err
		}
		if len(status) == 0 {
			continue
		}
		stat[sub.Name()] = status
	}
	return stat, nil
}
