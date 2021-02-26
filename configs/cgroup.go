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
	writeDevice string `file:"blkio.throttle.write_bps_device"`
	ReadDevice  string `file:"blkio.throttle.read_bps_device"`
}

func NewDefaultCgroupConfig() *CgroupConfig {
	return &CgroupConfig{
		CpuCfsPeriodUs: "100000",
		CpuCfsQuotaUs:  "200000",
		CpuSetCpus:     "0,1,2",
		CpuSetMems:     "0",

		MemoryLimit: "100M",
		Swappiness:  "0",

		writeDevice: "8:0 102400",
		ReadDevice:  "8:0 102400",
	}
}
