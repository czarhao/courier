package cgroup

import (
	"courier/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig2Map(t *testing.T) {
	cfg := configs.NewDefaultCgroupConfig()
	cfg.CpuShares = "512"
	m := config2map(cfg)
	assert.Equal(t, m["cpu.shares"], "512")
}
