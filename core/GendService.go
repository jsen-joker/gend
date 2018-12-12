package core

import "github.com/jsen-joker/gend/core/entity"

// 服务接口
type GendService interface {
	Init() error
	GenId() int64
	ExpId(id int64) entity.ID
}
