package populate

import (
	"github.com/jsen-joker/gend/core/config"
	"github.com/jsen-joker/gend/core/entity"
	"github.com/jsen-joker/gend/core/utils"
)

// 核心，计算获取ID动态部分
type IdPopulate interface {
	Populate(timer *utils.Timer, gendConfig *config.GendConfig, id *entity.ID) error
}
