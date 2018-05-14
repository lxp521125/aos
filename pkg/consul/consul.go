package consul

import (
	"gitlab.gaodun.com/golib/consul"
)

// GdConfig 配置文件
var GdConsul, err = InitConfig()

// InitConfig 配置文件
func InitConfig() (map[string]string, error) {
	config, err := consul.GetConf("")
	if err != nil {
		return nil, err
	}
	return config, nil
}
