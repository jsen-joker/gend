package populate

import (
	"gend/core/config"
	"gend/core/entity"
	"gend/core/utils"
)

// 核心，计算获取ID动态部分
type IdPopulate interface {
	Populate(timer *utils.Timer, gendConfig *config.GendConfig, id *entity.ID) error
}
