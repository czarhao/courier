package namespace

import (
	"courier/configs"
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

type namespace string

// 六种 namespace
const (
	NET  namespace = "net"
	PID  namespace = "pid"
	NS   namespace = "mnt"
	UTS  namespace = "uts"
	IPC  namespace = "ipc"
	// 一般是不会使用 user namespace 的
	USER namespace = "user"
)

func config2nss(cfg *configs.NamespaceConfig) []namespace {
	var namespaces = make([]namespace, 0, 6)
	if cfg.UseIPC {
		namespaces = append(namespaces, IPC)
	}
	if cfg.UseNET {
		namespaces = append(namespaces, NET)
	}
	if cfg.UseNS {
		namespaces = append(namespaces, NS)
	}
	if cfg.UsePID {
		namespaces = append(namespaces, PID)
	}
	if cfg.UseUSER {
		namespaces = append(namespaces, USER)
	}
	if cfg.UseUTS {
		namespaces = append(namespaces, UTS)
	}
	return namespaces
}

func (n namespace) isSupported() bool {
	path := fmt.Sprintf("/proc/self/namespace/%s", n)
	_, err := os.Stat(path)
	return err == nil
}

func (n namespace) isUsed(pid int) bool {
	_, err := os.Stat(n.getPath(pid))
	return err == nil
}

func (n namespace) getFlag() int {
	switch n {
	case NET:
		return unix.CLONE_NEWNET
	case PID:
		return unix.CLONE_NEWPID
	case NS:
		return unix.CLONE_NEWNS
	case UTS:
		return unix.CLONE_NEWUTS
	case IPC:
		return unix.CLONE_NEWIPC
	case USER:
		return unix.CLONE_NEWUSER
	}
	return 0
}

func (n namespace) getPath(pid int) string {
	return fmt.Sprintf("/proc/%d/n/%s", pid, n)
}
