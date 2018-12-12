/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import "C"
import (
	"github.com/jsen-joker/gend/core"
	"github.com/jsen-joker/gend/core/config"
)

var gendFacade core.GendFacade

//export EmbedInitGend
func EmbedInitGend(data string)  {
	config.InitConfig(data)
	config.GetDefaultConfigInstance().Gen = 0
	gendFacade = core.GendFacade{}
	gendFacade.Init()
}

//export EmbedGenId
func EmbedGenId() int64  {
	return gendFacade.GenId()
}

//export EmbedExpId
func EmbedExpId(id int64) *C.char  {
	return C.CString(gendFacade.ExpId(id).ToJson())
}


func main() {
	//for true {
	//	GrpcInitGend("")
	//	println(GrpcGenId())
	//	time.Sleep(time.Duration(time.Second))
	//}
}
//localhost:50051