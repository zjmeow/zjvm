package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type IntLoad struct {
	base.NoOperandsInstruction
	index uint
}
type ILOAD struct{ base.Index8Instruction }

func NewLoadInt(index uint) *IntLoad {
	return &IntLoad{
		index: index,
	}
}

func iLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (ins *ILOAD) Execute(frame *rtda.Frame) {
	iLoad(frame, ins.Index)
}

func (ins *IntLoad) Execute(frame *rtda.Frame) {
	iLoad(frame, ins.index)
}
