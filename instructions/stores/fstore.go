package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type FloatStore struct {
	base.NoOperandsInstruction
	index uint
}

type FSTORE struct{ base.Index8Instruction }

func NewStoreFloat(index uint) *FloatStore {
	return &FloatStore{
		index: index,
	}
}
func (ins *FloatStore) Execute(frame *rtda.Frame) {
	fStore(frame, ins.index)
}

func fStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

func (ins *FSTORE) Execute(frame *rtda.Frame) {
	fStore(frame, ins.Index)
}
