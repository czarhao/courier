package configs

type ImageConfig struct {
	BasePath string
	ImageDir string
	LayerDir string
	WriteDir string
}

func NewDefaultImageConfig() *ImageConfig {
	return &ImageConfig{
		BasePath: "/root/.courier",
		ImageDir: "/root/.courier/images",
		LayerDir: "/root/.courier/layers",
		WriteDir: "/root/.courier/write",
	}
}
