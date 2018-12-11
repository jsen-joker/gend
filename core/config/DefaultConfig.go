package config

import (
	"encoding/json"
)
type DefaultConfig struct {
	RpcAddress string
	// 默认字段值
	// 版本 扩展 0
	Version int64
	// 0 最大峰值 1 最小粒度
	IdType int64
	// 生成方式 0 嵌入式生成 1 服务器生成
	Gen int64
	// 机器ID
	Machine int64

	// 字段范围
	VersionBits byte
	TypeBits byte
	GenBits byte
	TimeBits byte
	SerialBits byte
	MachineBits byte
}

var m *DefaultConfig

func GetDefaultConfigInstance() *DefaultConfig {
	if m == nil {
		m = &DefaultConfig {}
		m.RpcAddress = "localhost:50051"
	}
	return m
}

func InitConfig(data string) {
	JsonParse := &JsonStruct{}
	v := GetDefaultConfigInstance()
	JsonParse.Load(data, &v)

}

type JsonStruct struct {
}

func (jst *JsonStruct) Load(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return
	}
}