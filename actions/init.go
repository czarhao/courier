package actions

import (
	"courier/container"
	"courier/mount"
	"fmt"
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
	fmt.Println(mount.SetMount())
	defer fmt.Println(mount.ClearMount())
	path, err := exec.LookPath(cmds[0])
	if err != nil {
		return err
	}
	return syscall.Exec(path, cmds[0:], os.Environ())
}
