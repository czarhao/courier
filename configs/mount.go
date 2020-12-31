package configs

type MountConfig struct {
	Path         string
	ReadLayer    string
	WriteLayer   string
}

func NewDefaultMountConfig() *MountConfig {
	return &MountConfig{
		Path: "/home/czarhao/tmp/root",
		ReadLayer: "/home/czarhao/tmp/busybox",
		WriteLayer: "/home/czarhao/tmp/write",
	}
}
