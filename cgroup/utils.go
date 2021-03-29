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
		cfgKey    = reflect.TypeOf(config).Elem()
		cfgVal    = reflect.ValueOf(config).Elem()
	)
	for i := 0; i < cfgKey.NumField(); i++ {
		if value := cfgVal.Field(i).String(); value != "" {
			key := cfgKey.Field(i).Tag.Get("file")
			configMap[key] = value
		}
	}
	return configMap
}
