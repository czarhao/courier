package container

import (
	"courier/cgroup"
	"courier/configs"
	"courier/namespace"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type proc struct {
	ns namespace.Manager
	cm cgroup.Manager
	cfg *configs.ContainerConfig

	wpipe *os.File
	init *exec.Cmd
}

func NewProc(config *configs.ContainerConfig) (*proc, error) {
	rpipe, wpipe, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create pipe, err: %v", err)
	}

	cmd := exec.Command("/proc/self/exe", "init")
	cmd.ExtraFiles = []*os.File{rpipe}

	return &proc{
		ns:  namespace.NewNSManager(),
		cm:  cgroup.NewManager(),
		cfg: config,
		wpipe: wpipe,
		init: cmd,
	}, err
}

func (p *proc) SetNamespace() error {
	flags, err := p.ns.GetCloneFlag(p.cfg.Namespace)
	if err != nil {
		return err
	}
	p.init.SysProcAttr = &syscall.SysProcAttr{Cloneflags: flags}
	return nil
}

func (p *proc) CreateCgroup() error {
	if p.cfg.Other.Name == ""{
		return fmt.Errorf("not set container path")
	}
	return p.cm.Create(p.cfg.Cgroup, p.cfg.Other.Name)
}

func (p *proc) SetCgroup() error {
	return p.cm.Apply(p.init.Process.Pid, p.cfg.Other.Name)
}

func (p *proc) DestroyCgroup() error {
	return p.cm.Destroy(p.cfg.Other.Name)
}

func (p *proc) StartInit() error {
	if err := p.init.Start(); err != nil {
		return fmt.Errorf("start init proc failed, err: %v", err)
	}
	return nil
}
