package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type Process interface {
}

type proc struct {
	cmd    *exec.Cmd
	writer *os.File
}

func NewProc() (Process, error) {
	rpipe, wpipe, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("failed create linux pipe, error: %v", err)
	}

	p := exec.Command("/proc/self/exe", "init")
	p.SysProcAttr = &syscall.SysProcAttr{Cloneflags: CLONE_FLAGS}

	p.Stdin, p.Stdout, p.Stderr = os.Stdin, os.Stdout, os.Stderr

	p.ExtraFiles = []*os.File{rpipe}

	return &proc{
		cmd:    p,
		writer: wpipe,
	}, err
}
