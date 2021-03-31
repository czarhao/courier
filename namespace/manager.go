package namespace

import (
	"courier/configs"
	"fmt"
	"os"
)

type Manager interface {
	// GetCloneFlag 是针对一个 namespace config 来生成 flag
	GetCloneFlag(cfg *configs.NamespaceConfig) (uintptr, error)

	// VerifyNSConfig 判断这个 namespace config 中设置的 namespace
	// 在当前的系统是否支持
	VerifyNSConfig(cfg *configs.NamespaceConfig) (bool, error)

	// 判断这个 namespace 当前系统是否是支持
	IsSupported(ns string) bool

	// GetProcNamespace 获取一个进程设置了哪些 namespace
	GetProcNamespace(pid int) []string
}

type nsManager struct{}

func NewNSManager() *nsManager {
	return &nsManager{}
}

func (m *nsManager) VerifyNSConfig(cfg *configs.NamespaceConfig) (bool, error) {
	nss := config2nss(cfg)
	for _, ns := range nss {
		if !ns.isSupported() {
			return false, fmt.Errorf("not supported by current os does not support %s namespace", ns)
		}
	}
	return true, nil
}

func (m *nsManager) GetCloneFlag(cfg *configs.NamespaceConfig) (uintptr, error) {
	if support, err := m.VerifyNSConfig(cfg); !support {
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

func (m *nsManager) GetProcNamespace(pid int) []string {
	nss := make([]string, 0, 6)
	if NET.isUsed(pid) {
		nss = append(nss, "net")
	}
	if PID.isUsed(pid) {
		nss = append(nss, "pid")
	}
	if NS.isUsed(pid) {
		nss = append(nss, "namespace")
	}
	if UTS.isUsed(pid) {
		nss = append(nss, "uts")
	}
	if IPC.isUsed(pid) {
		nss = append(nss, "ipc")
	}
	if USER.isUsed(pid) {
		nss = append(nss, "user")
	}
	return nss
}

func (m *nsManager) IsSupported(ns string) bool {
	path := fmt.Sprintf("/proc/self/namespace/%s", ns)
	_, err := os.Stat(path)
	return err == nil
}
