package configs

import (
	"courier/utils"
	"runtime"
	"strconv"
)

const (
	CgroupDefaultPeriod = 100000
)

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
// SetCpuUsage 传入的 usage 应该是 CPU核数*100 > usage > 0
func (cfg *CgroupConfig) SetCpuUsage(usage int) {
	if usage < 1 {
		utils.Logger.Warnf("set cpu usage failed, expect  usage > 0")
		return
	}
	sumCpu := runtime.NumCPU()
	useCpu := usage / 100
	if usage % 100 != 0 {
		useCpu++
	}

	if  sumCpu < useCpu {
		utils.Logger.Warnf("set cpu usage failed, usage < %d", sumCpu * 100)
		return
	}

	cfg.CpuCfsPeriodUs = strconv.Itoa(CgroupDefaultPeriod)
	cfg.CpuCfsQuotaUs = strconv.Itoa(CgroupDefaultPeriod / 100 * usage)

	return
}

// SetMemoryLimitMB 传入的是限制的最大内存数
func (cfg *CgroupConfig) SetMemoryLimitMB(mb int) {
	// mb => kb => b
	val := mb * 1024 * 1024
	cfg.MemoryLimit = strconv.Itoa(val)
}

