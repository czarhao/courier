package configs

type OtherConfig struct {
	TTY bool
	Name string
}

func NewDefaultOtherConfig() *OtherConfig {
	return &OtherConfig{TTY: true}
}
