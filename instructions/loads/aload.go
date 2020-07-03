package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type ALoad struct {
	base.NoOperandsInstruction
	index uint
}
type ALOAD struct{ base.Index8Instruction }

func NewALoad(index uint) *ALoad {
	return &ALoad{
		index: index,
	}
}

func aLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}
func (ins *ALoad) Execute(frame *rtda.Frame) {
	aLoad(frame, ins.index)
}

func (ins *ALOAD) Execute(frame *rtda.Frame) {
	aLoad(frame, ins.Index)
}
