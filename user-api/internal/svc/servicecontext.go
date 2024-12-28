package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user-api/internal/config"
	"user-api/internal/db"
)

// ServiceContext 服务上下文 依赖注入，需要用到的依赖都在此进行注入，比如配置，数据库连接，redis连接等
type ServiceContext struct {
	// 配置
	Config config.Config
	// 数据库连接
	Mysql sqlx.SqlConn
	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库连接
	mysqlConn := db.NewMysql(c.MysqlConfig)
	redisConn := db.NewRedis(c.RedisConfig)
	return &ServiceContext{
		Config: c,
		Mysql:  mysqlConn,
		Redis:  redisConn,
	}
}
