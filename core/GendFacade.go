package core

import (
	"gend/core/entity"
)

type GendFacade struct {
	gendService GendService
}

func (gf *GendFacade) Init()  {
	gf.gendService = NewDefaultGendService()
}
func (gf GendFacade) GenId() int64  {
	return gf.gendService.GenId()
}
func (gf GendFacade) ExpId(id int64) entity.ID  {
	return gf.gendService.ExpId(id)
}


