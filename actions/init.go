package actions

import (
	"courier/container"
	"courier/mount"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"syscall"
)

func Init(c *cli.Context) error {
	cmds, err := container.ReadCommandFromPip()
	if err != nil {
		return err
	}
	if err := mount.SetMount(); err != nil {
		return err
	}
	path, err := exec.LookPath(cmds[0])
	if err != nil {
		return err
	}
	return syscall.Exec(path, cmds[0:], os.Environ())
}
