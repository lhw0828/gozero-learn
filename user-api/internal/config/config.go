package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	// 自定义配置
	MysqlConfig MysqlConfig
	Auth        Auth
}

type Auth struct {
	AccessSecret string
	Expire       int64
}

type MysqlConfig struct {
	DataSource     string
	ConnectTimeout int64
}
