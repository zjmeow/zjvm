package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type DoubleLoad struct {
	base.NoOperandsInstruction
	index uint
}
type DLOAD struct{ base.Index8Instruction }

func NewDoubleLoad(index uint) *DoubleLoad {
	return &DoubleLoad{
		index: index,
	}
}

func dLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
func (ins *DoubleLoad) Execute(frame *rtda.Frame) {
	iLoad(frame, ins.index)
}

func (ins *DLOAD) Execute(frame *rtda.Frame) {
	dLoad(frame, ins.Index)
}
