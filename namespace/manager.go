package namespace

import (
	"courier/configs"
	"fmt"
)

type Manager interface {
	GetCloneFlag(config *configs.NamespaceConfig) (uintptr, error)
	IsSupported(cfg *configs.NamespaceConfig) (bool, error)
	GetProcNamespace(pid int) []string
}

type NSManager struct{}

func NewNSManager() NSManager {
	return NSManager{}
}

func (m NSManager) IsSupported(cfg *configs.NamespaceConfig) (bool, error) {
	nss := config2nss(cfg)
	for _, ns := range nss {
		if !ns.isSupported() {
			return false, fmt.Errorf("this os is not support %s namespace", ns)
		}
	}
	return true, nil
}

func (m NSManager) GetCloneFlag(cfg *configs.NamespaceConfig) (uintptr, error) {
	if support, err := m.IsSupported(cfg); !support {
		return 0, err
	}
	var (
		nss  = config2nss(cfg)
		flag = 0
	)
	for _, ns := range nss {
		flag |= ns.getFlag()
	}
	return uintptr(flag), nil
}

func (m NSManager) GetProcNamespace(pid int) []string {
	nss := make([]string, 0, 6)
	if NET.isSet(pid) {
		nss = append(nss, "net")
	}
	if PID.isSet(pid) {
		nss = append(nss, "pid")
	}
	if NS.isSet(pid) {
		nss = append(nss, "ns")
	}
	if UTS.isSet(pid) {
		nss = append(nss, "uts")
	}
	if IPC.isSet(pid) {
		nss = append(nss, "ipc")
	}
	if USER.isSet(pid) {
		nss = append(nss, "user")
	}
	return nss
}
