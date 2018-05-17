//Package consul package
package consul

import (
	"net/http"
	_ "reflect"
	"strings"
	"sync"
	"time"

	"github.com/armon/consul-api"
	"github.com/bernos/go-retry"
	// "gitlab.gaodun.com/golib/graylog"
)

var p sync.RWMutex

// AddGrayLog 日志
func AddGrayLog(info string) {
	// m := make(map[string]interface{})
	// p.Lock()
	// m["item"] = "consul"
	// p.Unlock()
	// graylog.GdLog(info, m)
}

//ConsulConfig 获取 consul 地址
func ConsulConfig(consulAddr string) *consulapi.Config {
	return &consulapi.Config{
		Address:    consulAddr,
		Scheme:     "http",
		HttpClient: http.DefaultClient,
	}
}

// RetryConsul 重试
func RetryConsul(projectName string, kv *consulapi.KV) func() (interface{}, error) {
	return func() (interface{}, error) {
		pair, _, err := kv.List(projectName, nil)
		if err != nil {
			AddGrayLog("consule first " + err.Error())
			return nil, err
		}
		return pair, nil
	}
}

// 获取consul数据 (key: value)
func getConsulVal(projectName string, consulConfig *consulapi.Config) (info map[string]string, err error) {
	info = make(map[string]string)
	client, err := consulapi.NewClient(consulConfig)

	if err != nil {
		AddGrayLog("consule first " + err.Error())
		return info, err
	}

	kv := client.KV()
	r := retry.Retry(RetryConsul(projectName, kv),
		retry.MaxRetries(3),
		retry.BaseDelay(time.Millisecond*1000))
	pair, err := r()
	if err != nil {
		return nil, err
	}

	if err != nil {
		AddGrayLog("consule first " + err.Error())
		return info, err
	}
	pa := pair.(consulapi.KVPairs)
	// 遍历key, value
	for _, item := range pa {
		keyList := strings.Split(item.Key, "/")
		keyNum := len(keyList)
		itemKey := keyList[keyNum-1]
		info[string(itemKey)] = string(item.Value)
	}

	return info, nil
}

// GetConf 获取配置项
// return map[string]string
func GetConf(projectName, url string) (minfo map[string]string, err error) {
	return getConsulVal(projectName, ConsulConfig(url))
}
