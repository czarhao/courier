package actions

import (
	"courier/utils"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	DefaultTemplateName = "hello_world.yaml"
	YamlTemplate = `
container:
  name: dev_container
  tty: true
image:
  path: /home/czarhao/Code/go/porter/image/busybox.tar
  run: /bin/sh
  layer:
    root_url: /home/czarhao/tmp/root
    mnt_url: /home/czarhao/tmp/mnt
    writer_url: /home/czarhao/tmp/writer
limit:
  memory: 50m
  cpu_share: 512
`
)

func Temp(c *cli.Context) error {
	var (
		file *os.File
		err error
		filename = DefaultTemplateName
	)
	if c.Args().Len() > 0 {
		filename = c.Args().First()
	}
	if file, err = os.Create(filename); err != nil {
		utils.Logger.Fatalf("create %s failed, %v", filename, err)
	}
	defer file.Close()

	if _, err = file.Write([]byte(YamlTemplate)); err != nil {
		utils.Logger.Fatalf("create %s failed, %v", filename, err)
	}
	return nil
}
