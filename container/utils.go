package container

import (
	"courier/configs"
	"courier/namespace"
	"courier/utils"
)

func CheckAndFormatConfig(config *configs.ContainerConfig) error {
	// err
	if _, err := namespace.NewNSManager().IsSupported(config.Namespace); err != nil {
		return err
	}
	// warn
	if config.Other.Name == "" {
		config.Other.Name = "courier_" + utils.RandString(10)
	}
	return nil
}
