package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type AStore struct {
	base.NoOperandsInstruction
	index uint
}

type ASTORE struct{ base.Index8Instruction }

func NewStoreA(index uint) *AStore {
	return &AStore{
		index: index,
	}
}
func (ins *AStore) Execute(frame *rtda.Frame) {
	aStore(frame, ins.index)
}

func aStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (ins *ASTORE) Execute(frame *rtda.Frame) {
	aStore(frame, ins.Index)
}
