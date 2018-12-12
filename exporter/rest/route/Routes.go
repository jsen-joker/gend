package route

import (
	"github.com/jsen-joker/gend/exporter/rest/controller"
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Name:"Index",     Method:"GET",   Pattern:"/", HandlerFunc: controller.Echo},
	Route{Name:"GenId",     Method:"GET",   Pattern:"/genId", HandlerFunc: controller.GenId},
	Route{Name:"ExpId",     Method:"GET",   Pattern:"/expId", HandlerFunc: controller.ExpId},
}