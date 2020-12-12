package namespace

import (
	"courier/configs"
	"fmt"
	"os"
	"syscall"
)

type ns string

const (
	NET  ns = "net"
	PID  ns = "pid"
	NS   ns = "ns"
	UTS  ns = "uts"
	IPC  ns = "ipc"
	USER ns = "user"
)

func config2nss(cfg *configs.NamespaceConfig) []ns {
	var types = make([]ns, 0, 6)
	if cfg.SetIPC {
		types = append(types, IPC)
	}
	if cfg.SetNET {
		types = append(types, NET)
	}
	if cfg.SetNS {
		types = append(types, NS)
	}
	if cfg.SetPID {
		types = append(types, PID)
	}
	if cfg.SetUSER {
		types = append(types, USER)
	}
	if cfg.SetUTS {
		types = append(types, UTS)
	}
	return types
}

func (n ns) isSupported() bool {
	path := fmt.Sprintf("/proc/self/ns/%s", n)
	_, err := os.Stat(path)
	return err == nil
}

func (n ns) isSet(pid int) bool {
	_, err := os.Stat(n.getPath(pid))
	return err == nil
}

func (n ns) getFlag() int {
	switch n {
	case NET:
		return syscall.CLONE_NEWNET
	case PID:
		return syscall.CLONE_NEWPID
	case NS:
		return syscall.CLONE_NEWNS
	case UTS:
		return syscall.CLONE_NEWUTS
	case IPC:
		return syscall.CLONE_NEWIPC
	case USER:
		return syscall.CLONE_NEWUSER
	}
	return 0
}

func (n ns) getPath(pid int) string {
	return fmt.Sprintf("/proc/%d/n/%s", pid, n)
}
