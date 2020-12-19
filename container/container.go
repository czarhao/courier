package container

type Process interface {
	SendCmd() error
	Wait() error

	// init
	Init() error

	SetNamespace() error

	CreateCgroup() error
	SetCgroup() error
	DestroyCgroup() error
}

func RunProc(p Process) (err error) {
	if err := initProc(p); err != nil {
		return err
	}
	if err := p.SendCmd(); err != nil {
		return err
	}
	return p.Wait()
}

func initProc(p Process) (err error) {
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
	defer func() {
		err = p.DestroyCgroup()
	}()
	if err = p.SetCgroup(); err != nil {
		return err
	}
	return nil
}
