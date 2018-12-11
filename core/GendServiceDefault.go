package core

import (
	"gend/core/config"
	"gend/core/entity"
	"gend/core/populate"
	"gend/core/utils"
)

type DefaultGendService struct {
	populate populate.IdPopulate
	config *config.GendConfig
	timer *utils.Timer
}

func (gs *DefaultGendService) Init() error {
	gs.populate = populate.NewLockPopulate()
	gs.config = config.NewGendConfig()
	gs.timer = utils.NewTimer(gs.config)
	return gs.timer.Init()
}
// 获取id
func (gs DefaultGendService) GenId() int64 {
	id := &entity.ID {
		Version: gs.config.Version,
		Type: gs.config.IdType,
		Gen: gs.config.Gen,
		Machine: gs.config.Machine,
	}
	// lock 计算 serial和 Time
	gs.populateId(id)
	return utils.Encode(*id, gs.config)
}
// 解码 id
func (gs DefaultGendService) ExpId(id int64) entity.ID {
	return *utils.Decode(id, gs.config)
}
func (gs DefaultGendService) populateId(id *entity.ID) {
	err := gs.populate.Populate(gs.timer, gs.config, id)
	if err != nil {
		print("error")
	}
}

func NewDefaultGendService() GendService  {
	gs := &DefaultGendService{}
	err := gs.Init()
	if err != nil {
		println(err)
	}
	return gs
}