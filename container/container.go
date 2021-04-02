package container

type Process interface {
	SendCmd() error
	Wait() error

	// init
	Init() error

	SetNamespace() error
	// cgroup
	CreateCgroup() error
	SetCgroup() error
	DestroyCgroup() error
	// rootfs
	Mount() error
	Umount() error
}

func RunProc(p Process) (err error) {
	if err = p.Mount(); err != nil {
		return err
	}
	defer p.Umount()

	if err = p.SetNamespace(); err != nil {
		return err
	}

	if err = p.Init(); err != nil {
		return err
	}
	// cgroup
	if err = p.CreateCgroup(); err != nil {
		return err
	}
	defer p.DestroyCgroup()

	if err = p.SetCgroup(); err != nil {
		return err
	}

	if err := p.SendCmd(); err != nil {
		return err
	}
	return p.Wait()
}
