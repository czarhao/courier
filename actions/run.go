package actions

import (
	"courier/configs"
	"courier/container"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {
	cfg := configs.NewDefaultContainerConfig()
	proc, err := container.NewProc(cfg)
	if err != nil {
		return err
	}
	return container.RunProc(proc)
}
