package subsystem

type cpuPeriod struct {
	basic
}

type cpuQuota struct {
	basic
}

func NewCpuPeriod() Subsystem {
	return cpuPeriod{basic{
		name: "cpu",
		file: "cpu.cfs_period_us",
	}}
}

func NewCpuQuota() Subsystem {
	return cpuQuota{basic{
		name: "cpu",
		file: "cpu.cfs_quota_us",
	}}
}

func (c cpuQuota) IsDefault(value string) bool {
	return value == "-1"
}
