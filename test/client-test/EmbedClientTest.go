package main

import (
	"github.com/jsen-joker/gend/core"
	"github.com/jsen-joker/gend/core/config"
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


func main() {
	for true {
		EmbedInitGend("")
		println(EmbedGenId())
		println(EmbedExpId(EmbedGenId()))
		time.Sleep(time.Duration(time.Second))
	}
}
