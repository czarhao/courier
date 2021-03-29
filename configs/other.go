package configs

import "courier/utils"

type OtherConfig struct {
	TTY     bool
	Name    string
	Command []string
}

func NewDefaultOtherConfig() *OtherConfig {
	return &OtherConfig{
		TTY:     true,
		Name:    "courier_" + utils.RandString(10),
		Command: []string{"/bin/bash"},
	}
}
