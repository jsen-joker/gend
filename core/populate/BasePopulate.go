package populate

import (
	"github.com/jsen-joker/gend/core/config"
	"github.com/jsen-joker/gend/core/entity"
	"github.com/jsen-joker/gend/core/utils"
)

// 核心，计算获取ID动态部分

type BasePopulate struct {
	LastTimestamp int64
	serial int64
}
func (basePopulate *BasePopulate) Populate(timer *utils.Timer, gendConfig *config.GendConfig, id *entity.ID) error  {
	var timestamp, err = timer.GenTime()
	if err != nil {
		return err
	}
	if timestamp == basePopulate.LastTimestamp {
		basePopulate.serial ++
		basePopulate.serial &= gendConfig.SerialMask
		if basePopulate.serial == 0 {
			timestamp, err = timer.TillNextTimeUnit(basePopulate.LastTimestamp)
			if err != nil {
				return err
			}
		}
	} else {
		basePopulate.LastTimestamp = timestamp
		basePopulate.serial = 0
	}
	// just mock
	id.Serial = basePopulate.serial
	id.Time = timestamp
	return nil
}