package populate

import (
	"gend/core/config"
	"gend/core/entity"
	"gend/core/utils"
	"sync"
)

type LockPopulate struct {
	*BasePopulate
	lock *sync.Mutex
	LS uint64
}
func (lockPopulate *LockPopulate) Populate(timer *utils.Timer, gendConfig *config.GendConfig, id *entity.ID) error  {
	lockPopulate.lock.Lock()
	defer lockPopulate.lock.Unlock()

	err := lockPopulate.BasePopulate.Populate(timer, gendConfig, id)
	if err != nil {
		print("error")
		return err
	}
	return nil
}

func NewLockPopulate() IdPopulate  {
	println("create lock populate")
	lc := LockPopulate{BasePopulate: &BasePopulate{},}
	lc.lock = new(sync.Mutex)
	return &lc
}
