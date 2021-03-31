package configs

type ContainerConfig struct {
	Cgroup    *CgroupConfig
	Mount     *RootfsConfig
	Namespace *NamespaceConfig
	Image     *ImageConfig
	Other     *OtherConfig
}

func NewDefaultContainerConfig() *ContainerConfig {
	return &ContainerConfig{
		Cgroup:    NewDefaultCgroupConfig(),
		Mount:     NewDefaultMountConfig(),
		Namespace: NewDefaultNSConfig(),
		Image:     NewDefaultImageConfig(),
		Other:     NewDefaultOtherConfig(),
	}
}
