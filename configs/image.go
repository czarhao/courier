package configs

import "path"

const BASE = "/root/.courier"

type ImageConfig struct {
	ImageDir string
	LayerDir string
	WriteDir string
}

func NewDefaultImageConfig() *ImageConfig {
	return &ImageConfig{
		ImageDir: path.Join(BASE, "images"),
		LayerDir: path.Join(BASE, "layer"),
		WriteDir: path.Join(BASE, "write"),
	}
}