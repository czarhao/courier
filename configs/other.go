package configs

import "courier/utils"

type OtherConfig struct {
	TTY     bool
	Name    string
	Image   string
	Command []string
}

func NewDefaultOtherConfig() *OtherConfig {
	return &OtherConfig{
		TTY:     true,
		Name:    "courier_" + utils.RandString(10),
		Image:   "busybox.tar",
		Command: []string{"/bin/sh"},
	}
}
