package dbconf

import (
	"aos/pkg/consul"
	"fmt"
	"os"
)

var dbConf = consul.GdConsul

type GdDb struct {
	DriverName string
	DriverDns  string
}

var PUBLIC_MYSQL_DB_HOST string

func init() {
	PUBLIC_MYSQL_DB_HOST = dbConf["PUBLIC_MYSQL_DB_HOST"]
}

func GetMySqlConfig() ([]GdDb, error) {
	if IsDev() {
		dbAllConfig := []GdDb{
			{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "gaodun_test", "414639e58d", "114.55.11.155", "3307", "resource")},
		}
		return dbAllConfig, nil
	}
	// 数据库配置
	dbAllConfig := []GdDb{
		{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf["UCENTER_MYSQL_USER"], dbConf["UCENTER_MYSQL_PASSWORD"], dbConf["PUBLIC_MYSQL_DB_HOST"], dbConf["UCENTER_MYSQL_DB_PORT"], "gaodun")},
	}
	return dbAllConfig, nil
}

// IsDev ...
func IsDev() bool {
	env := os.Getenv("SYSTEM_ENV")
	if env == "dev" || env == "" {
		return true
	}
	return false
}
