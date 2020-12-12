package configs

type NamespaceConfig struct {
	SetNET, SetPID, SetNS, SetUTS, SetIPC, SetUSER bool
}

func NewDefaultNSConfig() *NamespaceConfig {
	return &NamespaceConfig{
		SetNET:  true,
		SetPID:  true,
		SetNS:   true,
		SetUTS:  true,
		SetIPC:  true,
		SetUSER: true,
	}
}
