package configs

import "path"

type RootfsConfig struct {
	BaseDir      string
	ImageStorage string
	UnzipPath    string
	ReadLayer    string
	WriteLayer   string
}

func NewDefaultMountConfig() *RootfsConfig {
	return &RootfsConfig{
		BaseDir:    "/root/.courier",
		ImageStorage: "image",
		UnzipPath:  "unzip",
		ReadLayer:  "/home/czarhao/tmp/busybox",
		WriteLayer: "/home/czarhao/tmp/write",
	}
}

func (rc *RootfsConfig) GetImageStorage() string {
	return rc.merge(rc.ImageStorage)
}

func (rc *RootfsConfig) GetUnzipPath() string {
	return rc.merge(rc.UnzipPath)
}

func (rc *RootfsConfig) merge(p string) string {
	if len(p) == 0 {
		return rc.BaseDir
	}
	if p[0] == '/' {
		return p
	}
	return path.Join(rc.BaseDir, p)
}