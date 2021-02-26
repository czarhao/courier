package subsystem

type readDevice struct {
	basic
}

func NewReadDevice() Subsystem {
	return readDevice{basic{
		name: "blkio",
		file: "blkio.throttle.read_bps_device",
	}}
}

type writeDevice struct {
	basic
}

func NewWriteDevice() Subsystem {
	return writeDevice{basic{
		name: "blkio",
		file: "blkio.throttle.write_bps_device",
	}}
}
