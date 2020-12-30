package subsystem

type cpuCpus struct {
	basic
}

type cpuMems struct {
	basic
}

func NewCpuCpus() Subsystem {
	return cpuCpus{basic{
		name: "cpuset",
		file: "cpuset.cpus",
	}}
}

func NewCpuMems() Subsystem {
	return cpuMems{basic{
		name: "cpuset",
		file: "cpuset.mems",
	}}
}
