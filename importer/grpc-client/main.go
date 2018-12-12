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
	"github.com/jsen-joker/gend/core/config"
	"github.com/jsen-joker/gend/importer/grpc-client/service"
)

var c *service.ClientService


//export GrpcInitGend
func GrpcInitGend(data string)  {
	println(data)
	config.InitConfig(data)
	config.GetDefaultConfigInstance().Gen = 1
	c = service.GetClientService(config.GetDefaultConfigInstance().RpcAddress)
}

//export GrpcGenId
func GrpcGenId() int64  {
	return c.GenId()
}

//export GrpcExpId
func GrpcExpId(id int64) *C.char  {
	return C.CString(c.ExpId(id))
}
func main() {
	//for true {
	//	GrpcInitGend("")
	//	println(GrpcGenId())
	//	time.Sleep(time.Duration(time.Second))
	//}
}
//localhost:50051