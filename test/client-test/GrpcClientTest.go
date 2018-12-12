// first run exporter/grpc/boot/main.go
package main

import (
	"github.com/jsen-joker/gend/core/config"
	"github.com/jsen-joker/gend/importer/grpc-client/service"
	"time"
)

var c *service.ClientService


func GrpcInitGend(data string)  {
	println(data)
	config.InitConfig(data)
	config.GetDefaultConfigInstance().Gen = 1
	c = service.GetClientService(config.GetDefaultConfigInstance().RpcAddress)
}

func GrpcGenId() int64  {
	return c.GenId()
}

func GrpcExpId(id int64) string  {
	return c.ExpId(id)
}


func main() {
	for true {
		GrpcInitGend("")
		println(GrpcGenId())
		println(GrpcExpId(GrpcGenId()))
		time.Sleep(time.Duration(time.Second))
	}
}
