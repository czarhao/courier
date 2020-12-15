package configs

type ContainerConfig struct {
	Cgroup    *CgroupConfig
	Mount     *MountConfig
	Namespace *NamespaceConfig
	Other     *OtherConfig
}

func NewDefaultContainerConfig() *ContainerConfig {
	return &ContainerConfig{
		Cgroup:    NewDefaultCgroupConfig(),
		Mount:     NewDefaultMountConfig(),
		Namespace: NewDefaultNSConfig(),
		Other:     NewDefaultOtherConfig(),
	}
}
