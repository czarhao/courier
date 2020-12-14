package configs

type MountConfig struct {
	UsePivotRoot bool
	Path string
}

func NewDefaultMountConfig() *MountConfig {
	return &MountConfig{
		UsePivotRoot: true,
	}
}