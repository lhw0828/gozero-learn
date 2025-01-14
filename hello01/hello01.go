package main

import (
	"flag"
	"fmt"

	"hello01/internal/config"
	"hello01/internal/handler"
	"hello01/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/hello01-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	// 读取配置文件
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 初始化服务上下文
	ctx := svc.NewServiceContext(c)
	// 注册路由处理函数
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
