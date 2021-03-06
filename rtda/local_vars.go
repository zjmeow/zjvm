package rtda

import (
	"github.com/zjmeow/zjvm/rtda/heap"
	"math"
)

type LocalVars []Slot

func newLocalVars(size uint) LocalVars {
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

func (lv LocalVars) SetRef(index uint, val *heap.Object) {
	lv[index].ref = val
}
func (lv LocalVars) GetRef(index uint) *heap.Object {
	return lv[index].ref
}
func (lv LocalVars) SetSlot(index uint, slot Slot) {
	lv[index] = slot
}
