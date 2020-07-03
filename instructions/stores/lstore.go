package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type LongStore struct {
	base.NoOperandsInstruction
	index uint
}
type LSTORE struct{ base.Index8Instruction }

func NewStoreLong(index uint) *LongStore {
	return &LongStore{
		index: index,
	}
}

func lStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (ins *LongStore) Execute(frame *rtda.Frame) {
	lStore(frame, ins.index)
}

func (ins *LSTORE) Execute(frame *rtda.Frame) {
	lStore(frame, ins.Index)
}
