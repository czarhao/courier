package subsystem

type memoryLimit struct {
	basic
}

func NewMemoryLimit() Subsystem {
	return memoryLimit{basic{
		name: "memory",
		file: "memory.limit_in_bytes",
	}}
}

type swappiness struct {
	basic
}

func NewSwappiness() Subsystem {
	return swappiness{basic{
		name: "memory",
		file: "memory.swappiness",
	}}
}
