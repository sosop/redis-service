package config

import (
	"github.com/sosop/libconfig"
)

var (
	iniConfig *libconfig.IniConfig
	mode      string
)

func init() {
	iniConfig = libconfig.NewIniConfig("config/config.cnf")
	mode = iniConfig.GetString("mode", "prod")
}

func GetString(key string, withMode bool, defaultValue ...string) string {
	if withMode {
		return iniConfig.GetString(mode+"::"+key, defaultValue...)
	} else {
		return iniConfig.GetString(key, defaultValue...)
	}
}
