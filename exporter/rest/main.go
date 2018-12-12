package main

import (
	"github.com/jsen-joker/gend/exporter/rest/controller"
	"github.com/jsen-joker/gend/exporter/rest/route"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	_, _ = fmt.Fprintln(writer, "Hello, ", html.EscapeString(request.URL.Path))
	//})
	controller.EmbedInitGend("")
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
