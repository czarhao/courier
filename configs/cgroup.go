package configs

type CgroupConfig struct {
	// cpu
	CpuCfsPeriodUs string `file:"cpu.cfs_period_us"`
	CpuCfsQuotaUs  string `file:"cpu.cfs_quota_us"`

	CpuSetCpus     string `file:"cpuset.cpus"`
	CpuSetMems     string `file:"cpuset.mems"`
	// mem
	MemoryLimit string `file:"memory.limit_in_bytes"`
	Swappiness  string `file:"memory.swappiness"`
	// blkio
	WriteDevice string `file:"blkio.throttle.write_bps_device"`
	ReadDevice  string `file:"blkio.throttle.read_bps_device"`
}

func NewDefaultCgroupConfig() *CgroupConfig {
	return &CgroupConfig{}
}

func (cfg *CgroupConfig) SetCpu(cpuNum int, ) {

}
