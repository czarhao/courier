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
		},

		&cli.Command{
			Name:   "init",
			Hidden: true,
			Action: actions.Init,
		},

		&cli.Command{
			Name:    "template",
			Aliases: []string{"t", "tmp", "temp"},
			Usage:   `Create a standard template configuration, example: "courier temp (filename)"`,
			Action:  actions.Temp,
		},
	)

	if err := courier.Run(os.Args); err != nil {
		utils.Logger.Fatalf("Courier have some trouble: %v", err)
	}
}
