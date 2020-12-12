package configs

type CgroupConfig struct {
	CpuShares string `file:"cpu.shares"`
}

func NewDefaultCgroupConfig() *CgroupConfig {
	return &CgroupConfig{}
}
