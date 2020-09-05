package rtda

import "github.com/zjmeow/zjvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}

func (slot Slot) IntValue() int32 {
	return int32(slot.num)
}
