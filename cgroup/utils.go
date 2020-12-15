package cgroup

import (
	"courier/configs"
	"reflect"
)

func config2map(config *configs.CgroupConfig) map[string]string {
	if config == nil {
		return nil
	}
	var (
		configMap = map[string]string{}
		cfgKey    = reflect.TypeOf(*config).Elem()
		cfgVal    = reflect.ValueOf(*config).Elem()
	)
	for i := 0; i < cfgKey.NumField(); i++ {
		key := cfgKey.Field(i).Tag.Get("file")
		value := cfgVal.Field(i).String()
		configMap[key] = value
	}
	return configMap
}

func map2config(configMap map[string]string) *configs.CgroupConfig {
	config := &configs.CgroupConfig{}
	cfg := reflect.ValueOf(config).Elem()
	for i := 0; i < cfg.NumField(); i++ {

	}
	return nil
}
