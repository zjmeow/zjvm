package rtda

import "github.com/zjmeow/zjvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
