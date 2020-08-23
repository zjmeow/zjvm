package heap

import (
	"math"
)

type LocalVars []Slot

func newSlots(size uint) LocalVars {
	if size > 0 {
		return make([]Slot, size)
	}
	return nil
}

func (lv LocalVars) SetInt(index uint, val int32) {
	lv[index].num = val
}
func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].num
}
func (lv LocalVars) SetFloat(index uint, val float32) {
	lv[index].num = int32(math.Float32bits(val))
}
func (lv LocalVars) GetFloat(index uint) float32 {
	bits := uint32(lv[index].num)
	return math.Float32frombits(bits)
}
func (lv LocalVars) SetLong(index uint, val int64) {
	lv[index].num = int32(val)
	lv[index+1].num = int32(val >> 32)
}
func (lv LocalVars) GetLong(index uint) int64 {
	low := lv[index].num
	high := lv[index+1].num
	return int64(high)<<32 | int64(low)
}
func (lv LocalVars) SetDouble(index uint, val float64) {
	bits := int64(math.Float64bits(val))
	lv.SetLong(index, bits)
}
func (lv LocalVars) GetDouble(index uint) float64 {
	bits := uint64(lv.GetLong(index))
	return math.Float64frombits(bits)
}

func (lv LocalVars) SetRef(index uint, val *Object) {
	lv[index].ref = val
}
func (lv LocalVars) GetRef(index uint) *Object {
	return lv[index].ref
}

func (lv LocalVars) Set(index uint, val interface{}) {
	switch val.(type) {
	case int32:
		lv.SetInt(index, val.(int32))
	case int64:
		lv.SetLong(index, val.(int64))
	case float32:
		lv.SetFloat(index, val.(float32))
	case float64:
		lv.SetDouble(index, val.(float64))
	case *Object:
		lv.SetRef(index, val.(*Object))
	default:
		panic("type not found")
	}
}
