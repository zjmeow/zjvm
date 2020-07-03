package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type IntStore struct {
	base.NoOperandsInstruction
	index uint
}

type ISTORE struct{ base.Index8Instruction }

func NewStoreInt(index uint) *IntStore {
	return &IntStore{
		index: index,
	}
}

func iStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

func (ins *IntStore) Execute(frame *rtda.Frame) {
	iStore(frame, ins.index)
}

func (ins *ISTORE) Execute(frame *rtda.Frame) {
	iStore(frame, ins.Index)
}
