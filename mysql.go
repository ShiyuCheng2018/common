package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
)

type MySqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMysqlFromConsul(config config.Config, path ...string) *MySqlConfig {
	mysqlConfig := &MySqlConfig{}
	err := config.Get(path...).Scan(mysqlConfig)
	if err != nil {
		log.Error(err)
	}
	return mysqlConfig
}
