package configs

type NamespaceConfig struct {
	UseNET, UsePID, UseNS, UseUTS, UseIPC, UseUSER bool
}

func NewDefaultNSConfig() *NamespaceConfig {
	return &NamespaceConfig{
		UseNET:  true,
		UsePID:  true,
		UseNS:   true,
		UseUTS:  true,
		UseIPC:  true,
		UseUSER: false,
	}
}
