package main

import "C"
import (
	"github.com/jsen-joker/gend/core/config"
	"github.com/jsen-joker/gend/importer/grpc-client/service"
	"strconv"
	"sync"
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

func GrpcExpId(id int64) string {
	return c.ExpId(id)
}

func grpcTask(rg int64, n *sync.WaitGroup) {
	defer n.Done()
	var i int64 = 0
	for i < rg {
		i ++
		GrpcGenId()
	}
}

func main() {
	// 100 100 23
	// 100 500 18
	// 100 1000 19
	// 100 5000 19
	// 100 10000 18

	var n sync.WaitGroup
	GrpcInitGend("")
	var group int64 = 10000
	start := time.Now()
	for i := 0; i < 100; i++ {
		n.Add(1)
		go grpcTask(group, &n)
	}
	// for true {
	// 	EmbedInitGend("")
	// 	println(EmbedGenId())
	// 	println(EmbedExpId(EmbedGenId()))
	// 	time.Sleep(time.Duration(time.Second))
	// }
	n.Wait()
	println(strconv.FormatInt(time.Now().UnixNano() / 1e3, 10))
	ava := time.Now().UnixNano() / 1e3 - start.UnixNano() / 1e3
	println(ava / 100 / group)
}
