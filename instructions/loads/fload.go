package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type FloatLoad struct {
	base.NoOperandsInstruction
	index uint
}
type FLOAD struct{ base.Index8Instruction }

func NewFloatLoad(index uint) *FloatLoad {
	return &FloatLoad{
		index: index,
	}
}

func fLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
func (ins *FloatLoad) Execute(frame *rtda.Frame) {
	iLoad(frame, ins.index)
}

func (ins *FLOAD) Execute(frame *rtda.Frame) {
	fLoad(frame, ins.Index)
}
