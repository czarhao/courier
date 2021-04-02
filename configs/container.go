package configs

type ContainerConfig struct {
	Cgroup    *CgroupConfig
	Namespace *NamespaceConfig
	Image     *ImageConfig
	Other     *OtherConfig
}

func NewDefaultContainerConfig() *ContainerConfig {
	return &ContainerConfig{
		Cgroup:    NewDefaultCgroupConfig(),
		Namespace: NewDefaultNSConfig(),
		Image:     NewDefaultImageConfig(),
		Other:     NewDefaultOtherConfig(),
	}
}
