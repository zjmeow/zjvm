package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type FLOAD struct{ base.Index8Instruction }
type FLOAD_0 struct{ base.NoOperandsInstruction }
type FLOAD_1 struct{ base.NoOperandsInstruction }
type FLOAD_2 struct{ base.NoOperandsInstruction }
type FLOAD_3 struct{ base.NoOperandsInstruction }

func fLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

func (ins *FLOAD) Execute(frame *rtda.Frame) {
	fLoad(frame, ins.Index)
}

func (ins *FLOAD_0) Execute(frame *rtda.Frame) {
	fLoad(frame, 0)
}

func (ins *FLOAD_1) Execute(frame *rtda.Frame) {
	fLoad(frame, 1)
}

func (ins *FLOAD_2) Execute(frame *rtda.Frame) {
	fLoad(frame, 2)
}

func (ins *FLOAD_3) Execute(frame *rtda.Frame) {
	fLoad(frame, 3)
}
