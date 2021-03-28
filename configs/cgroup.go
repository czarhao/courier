package configs

import (
	"courier/utils"
	"errors"
	"strconv"
)

const CgroupDefaultValue = -1

var ErrCpuUsageSet = errors.New("set cpu usage failed, usage must: 100 > usage > 0")

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

// 100 > usage > 0
func (cfg *CgroupConfig) SetCpuUsage(usage int) {
	if usage == CgroupDefaultValue {
		return
	}
	if usage > 100 || usage < 1 {
		utils.Logger.Warnf("set cpu usage failed, expect usage < 100 && usage > 0")
		return
	}
	// TODO
	return
}

func (cfg *CgroupConfig) SetCpuNum(num int) {
	cpus := "0"
	for i := 1; i < num; i++ {
		cpus += "," + strconv.Itoa(i)
	}
	cfg.CpuSetCpus = cpus
}

func (cfg *CgroupConfig) SetMemoryLimitMB(mb int) {
	// mb => kb => b
	val := mb * 1024 * 1024
	cfg.MemoryLimit = strconv.Itoa(val)
}

func (cfg *CgroupConfig)

