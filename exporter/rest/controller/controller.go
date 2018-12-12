package controller

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/jsen-joker/gend/core"
	"github.com/jsen-joker/gend/core/config"
	"github.com/jsen-joker/gend/core/entity"
	"github.com/jsen-joker/gend/exporter/rest/utils"
	"net/http"
	"strconv"
)

var gendFacade core.GendFacade

func EmbedInitGend(data string)  {
	config.InitConfig(data)
	config.GetDefaultConfigInstance().Gen = 0
	gendFacade = core.GendFacade{}
	gendFacade.Init()
}

func embedGenId() int64  {
	return gendFacade.GenId()
}

func embedExpId(id int64) entity.ID {
	return gendFacade.ExpId(id)
}


func Echo(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintln(w, "Hello Gend")
}

func GenId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := utils.Succeed(embedGenId())
	if err := json.NewEncoder(w).Encode(resp); err != nil{
		panic(err)
	}
}

func ExpId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	vars := r.URL.Query()
	id, err := strconv.ParseInt(vars["id"][0], 10, 64)
	if err != nil {
		panic(err)
	}
	resp := utils.Succeed(embedExpId(id))
	if err := json.NewEncoder(w).Encode(resp); err != nil{
		panic(err)
	}
}