package configs

type NamespaceConfig struct {
	SetNET, SetPID, SetNS, SetUTS, SetIPC, SetUSER bool
}

func NewDefaultNSConfig() *NamespaceConfig {
	return &NamespaceConfig{
		SetNET:  false,
		SetPID:  false,
		SetNS:   false,
		SetUTS:  false,
		SetIPC:  false,
		SetUSER: false,
	}
}
