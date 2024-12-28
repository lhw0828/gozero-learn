package main

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-server/greet"
	"log"
)

func main() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/client.yml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	client := greet.NewGreetClient(conn.Conn())
	resp, err := client.Ping(context.Background(), &greet.Request{Ping: "ping"})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resp)
}
