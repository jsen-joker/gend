package entity

import "fmt"

// ID模型
type ID struct {
	Version int64 `json:"version"`
	Type int64 `json:"type"`
	Gen int64 `json:"gen"`
	Time int64 `json:"time"`
	Serial int64 `json:"serial"`
	Machine int64 `json:"machine"`
}
// id 核心schema

func (id ID) ToString() {
	fmt.Printf(`

Version: %d,
Type: %d,
Gen: %d,
Time: %d,
Serial: %d,
Machine: %d,


`, id.Version, id.Type, id.Gen, id.Time, id.Serial, id.Machine)
}

func (id ID) ToJson() string {
	return fmt.Sprintf(`{"Version":%d, "Type":%d, "Gen":%d, "Time":%d, "Serial":%d, "Machine":%d, }`, id.Version, id.Type, id.Gen, id.Time, id.Serial, id.Machine)
}