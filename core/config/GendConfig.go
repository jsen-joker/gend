package config

const(
	// 最大峰值
	SECONDS = 0
	// 最小粒度
	MILLISECONDS = 1
)


type GendConfig struct {
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

	// 0 1
	GenType int
	// dynamic
	VersionMask int64
	TypeMask int64
	GenMask int64
	TimeMask int64
	SerialMask int64
	MachineMask int64

	VersionPos uint64
	TypePos uint64
	GenPos uint64
	TimePos uint64
	SerialPos uint64
	MachinePos uint64
}

func (gc *GendConfig) init() {
	gc.GenType = SECONDS
	gc.MachinePos = 0
	gc.SerialPos = gc.MachinePos + uint64(gc.MachineBits)
	gc.TimePos = gc.SerialPos + uint64(gc.SerialBits)
	gc.GenPos = gc.TimePos + uint64(gc.TimeBits)
	gc.TypePos = gc.GenPos + uint64(gc.GenBits)
	gc.VersionPos = gc.TypePos + uint64(gc.TypeBits)

	var mask = -1 ^ -1 << gc.VersionBits
	gc.VersionMask = int64(mask)
	mask = -1 ^ -1 << gc.TypeBits
	gc.TypeMask = int64(mask)
	mask = -1 ^ -1 << gc.GenBits
	gc.GenMask = int64(mask)
	mask = -1 ^ -1 << gc.TimeBits
	gc.TimeMask = int64(mask)
	mask = -1 ^ -1 << gc.SerialBits
	gc.SerialMask = int64(mask)
	mask = -1 ^ -1 << gc.MachineBits
	gc.MachineMask = int64(mask)
}



func NewGendConfig() *GendConfig {
	dc := GetDefaultConfigInstance()

	gc := &GendConfig{
		Version: 0,
		IdType: 0,
		Gen: 0,
		Machine: 1,

		VersionBits: 1,
		TypeBits: 1,
		GenBits: 1,
		TimeBits: 30,
		SerialBits: 20,
		MachineBits: 10,
	}
	if dc.Version != 0 {
		gc.Version = dc.Version
	}
	if dc.IdType != 0 {
		gc.IdType = dc.IdType
	}
	if dc.Gen != 0 {
		gc.Gen = dc.Gen
	}
	if dc.Machine != 0 {
		gc.Machine = dc.Machine
	}
	if dc.VersionBits != 0 {
		gc.VersionBits = dc.VersionBits
	}
	if dc.TypeBits != 0 {
		gc.TypeBits = dc.TypeBits
	}
	if dc.GenBits != 0 {
		gc.GenBits = dc.GenBits
	}
	if dc.TimeBits != 0 {
		gc.TimeBits = dc.TimeBits
	}
	if dc.SerialBits != 0 {
		gc.SerialBits = dc.SerialBits
	}
	if dc.MachineBits != 0 {
		gc.MachineBits = dc.MachineBits
	}
	gc.init()
	return gc
}