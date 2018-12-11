package main

import (
	"gend/core"
)


func main() {
	var gs = core.NewDefaultGendService()
	id := gs.GenId()
	gs.ExpId(id).ToString()
	println(gs.GenId())
	println(gs.GenId())
	println(gs.GenId())
}
