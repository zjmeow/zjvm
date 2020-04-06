package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type ILOAD struct{ base.Index8Instruction }
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

func iLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (ins *ILOAD) Execute(frame *rtda.Frame) {
	iLoad(frame, ins.Index)
}

func (ins *ILOAD_0) Execute(frame *rtda.Frame) {
	iLoad(frame, 0)
}

func (ins *ILOAD_1) Execute(frame *rtda.Frame) {
	iLoad(frame, 1)
}

func (ins *ILOAD_2) Execute(frame *rtda.Frame) {
	iLoad(frame, 2)
}

func (ins *ILOAD_3) Execute(frame *rtda.Frame) {
	iLoad(frame, 3)
}
