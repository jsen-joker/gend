package utils

import (
	"github.com/jsen-joker/gend/core/config"
	"github.com/jsen-joker/gend/core/entity"
)

// id 和 long 之间的编解码

func Encode(id entity.ID, config *config.GendConfig) int64 {
	var ret int64 = 0
	ret |= id.Machine << config.MachinePos
	ret |= id.Serial << config.SerialPos
	ret |= id.Time << config.TimePos
	ret |= id.Gen << config.GenPos
	ret |= id.Type << config.TypePos
	ret |= id.Version << config.VersionPos

	return ret
}

func Decode(id int64, config *config.GendConfig) *entity.ID  {
	return &entity.ID {
		Machine: (id >> config.MachinePos) & config.MachineMask,
		Serial: (id >> config.SerialPos) & config.SerialMask,
		Time: (id >> config.TimePos) & config.TimeMask,
		Gen: (id >> config.GenPos) & config.GenMask,
		Type: (id >> config.TypePos) & config.TypeMask,
		Version: (id >> config.VersionPos) & config.VersionMask,
	}
} 
