package entity

import "fmt"

// ID模型
type ID struct {
	Version int64
	Type int64
	Gen int64
	Time int64
	Serial int64
	Machine int64
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