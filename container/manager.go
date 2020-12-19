package container

import (
	"courier/cgroup"
	"courier/configs"
	"courier/namespace"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type proc struct {
	ns  namespace.Manager
	cm  cgroup.Manager
	cfg *configs.ContainerConfig

	wpipe *os.File
	cmd   *exec.Cmd
}

func NewProc(config *configs.ContainerConfig) (*proc, error) {
	rpipe, wpipe, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create pipe, err: %v", err)
	}

	cmd := exec.Command("/proc/self/exe", "init")
	cmd.ExtraFiles = []*os.File{rpipe}

	if config.Other.TTY {
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	}

	return &proc{
		ns:    namespace.NewNSManager(),
		cm:    cgroup.NewManager(),
		cfg:   config,
		wpipe: wpipe,
		cmd:   cmd,
	}, err
}

func (p *proc) SetNamespace() error {
	flags, err := p.ns.GetCloneFlag(p.cfg.Namespace)
	if err != nil {
		return err
	}
	p.cmd.SysProcAttr = &syscall.SysProcAttr{Cloneflags: flags}
	return nil
}

func (p *proc) CreateCgroup() error {
	if p.cfg.Other.Name == "" {
		return fmt.Errorf("not set container path")
	}
	return p.cm.Create(p.cfg.Cgroup, p.cfg.Other.Name)
}

func (p *proc) SetCgroup() error {
	return p.cm.Apply(p.cmd.Process.Pid, p.cfg.Other.Name)
}

func (p *proc) DestroyCgroup() error {
	return p.cm.Destroy(p.cfg.Other.Name)
}

func (p *proc) Init() error {
	if err := p.cmd.Start(); err != nil {
		return fmt.Errorf("start cmd proc failed, err: %v", err)
	}
	return nil
}

func (p *proc) SendCmd() error {
	cmd := strings.Join(p.cfg.Other.Command, " ")
	if _, err := p.wpipe.WriteString(cmd); err != nil {
		return fmt.Errorf("write cmd failed, err: %v", err)
	}
	return p.wpipe.Close()
}

func (p *proc) Wait() error {
	return p.cmd.Wait()
}
