package actions

import (
	"courier/configs"
	"courier/container"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {
	proc, err := container.NewProc(configs.NewDefaultContainerConfig())
	if err != nil {
		return err
	}
	return container.RunProc(proc)
}
