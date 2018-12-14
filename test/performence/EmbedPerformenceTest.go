package main

import (
	"github.com/jsen-joker/gend/core"
	"github.com/jsen-joker/gend/core/config"
	"strconv"
	"sync"
	"time"
)

var gendFacade core.GendFacade

func EmbedInitGend(data string)  {
	config.InitConfig(data)
	config.GetDefaultConfigInstance().Gen = 0
	gendFacade = core.GendFacade{}
	gendFacade.Init()
}

func EmbedGenId() int64  {
	return gendFacade.GenId()
}

func EmbedExpId(id int64) string  {
	return gendFacade.ExpId(id).ToJson()
}

func task(rg int64, n *sync.WaitGroup) {
	defer n.Done()
	var i int64 = 0
	for i < rg {
		i ++
		EmbedGenId()
	}
}

func main() {
	// 1000 100 26
	// 1000 500 18
	// 1000 1000 25
	// 1000 5000 17
	// 1000 10000 22
	var n sync.WaitGroup
	EmbedInitGend("")
	var group int64 = 5000
	start := time.Now()
	for i := 0; i < 1000; i++ {
		n.Add(1)
		go task(group, &n)
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
