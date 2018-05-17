package consul

import (
	"aos/pkg/setting"
	"errors"
	"os"

	"github.com/aos-stack/golib/consul"
)

// GdConfig 配置文件
// var GdConsul, err = InitConfig()

// InitConfig 配置文件
func InitConfig() (map[string]string, error) {
	// env := os.Getenv("SYSTEM_ENV")
	if setting.CONSUL_URL == "" || setting.CONSUL_LIST_NAME == "" {
		return nil, errors.New("consul 配置为空")
	}
	config, err := consul.GetConf(setting.CONSUL_LIST_NAME, setting.CONSUL_URL)

	if err != nil || config == nil {
		return nil, err
	}
	return config, nil
}

//GetEnv 获取 consul address
func GetEnv() (string, error) {

	consuls := map[string]string{
		"dev":        "test.consul.xxxx.com",
		"test":       "t.consul.xxxx.com",
		"prepare":    "pre.consul.xxx.com",
		"production": "pro.consul.xxx.com",
	}

	env := os.Getenv("SYSTEM_ENV")
	if env == "" {
		return consuls["dev"], nil
	}

	return consuls[env], nil
}
