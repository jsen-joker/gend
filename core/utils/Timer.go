package utils

import (
	"errors"
	"fmt"
	"github.com/jsen-joker/gend/core/config"
	"time"
)

const EPOCH = 1544433018041

type Timer struct {
	Config *config.GendConfig
	maxTime int64
}

func (timer *Timer) Init() error {
	timer.maxTime = (1 << timer.Config.TimeBits) - 1
	uT, err := timer.GenTime()
	println(uT)
	if err != nil {
		return err
	}
	return nil
}

// 获取当前时间戳
func (timer *Timer) GenTime() (int64, error) {
	var t = time.Now().UnixNano() / 1e6 - EPOCH
	if timer.Config.GenType == config.SECONDS {
		t /= 1000
	}
	err := timer.validateTimestamp(t)
	if err != nil {
		return 0, err
	}
	return t, nil
}

// 当发生serial不够用的时候，for循环直到下一个可用时间段
func (timer *Timer) TillNextTimeUnit(lastTimestamp int64) (int64, error)  {
	var uT, err = timer.GenTime()
	if err != nil {
		return 0, err
	}
	for ; uT <= lastTimestamp;  {
		uT, err = timer.GenTime()
		if err != nil {
			return 0, err
		}
	}
	return uT, nil
}

// 时间是否超出范围
func (timer *Timer) validateTimestamp(timestamp int64) error {
	if timestamp > timer.maxTime {
		err := fmt.Sprintf("The current timestamp (%d >= %d) has overflowed, Vesta Service will be terminate.", timestamp, timer.maxTime)
		return errors.New(err)
	}
	return nil
}

func NewTimer(Config *config.GendConfig) *Timer {
	return &Timer{
		Config: Config,
	}
}