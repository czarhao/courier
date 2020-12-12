package main

import (
	"courier/actions"
	"courier/utils"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	courier := &cli.App{
		Name:    "Courier",
		Usage:   "Courier is a new container engine",
		Version: actions.Version(),
		Action: func(c *cli.Context) error {
			utils.Logger.Info(`Courier is a new container engine, use "help" see more.`)
			return nil
		},
	}

	courier.Commands = append(courier.Commands,
		&cli.Command{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   `Run a container, example: "courier run template.yaml"`,
			Action:  actions.Run,
		})

	if err := courier.Run(os.Args); err != nil {
		utils.Logger.Fatal("Courier have some trouble:%v", err)
	}
}
