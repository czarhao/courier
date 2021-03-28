package cgroup

import (
	"courier/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig2map(t *testing.T) {
	cgroup := configs.NewDefaultCgroupConfig()
	_ = config2map(cgroup)
	cgroup.SetCpuCfsPeriodUs("666")
	config := config2map(cgroup)
	assert.Equal(t, "666", config["cpu.cfs_period_us"])
}