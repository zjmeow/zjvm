package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type DoubleStore struct {
	base.NoOperandsInstruction
	index uint
}

type DSTORE struct{ base.Index8Instruction }

func NewStoreDouble(index uint) *DoubleStore {
	return &DoubleStore{
		index: index,
	}
}

func (ins *DoubleStore) Execute(frame *rtda.Frame) {
	dStore(frame, ins.index)
}

func dStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (ins *DSTORE) Execute(frame *rtda.Frame) {
	dStore(frame, ins.Index)
}
