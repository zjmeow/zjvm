package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type DLOAD struct{ base.Index8Instruction }
type DLOAD_0 struct{ base.NoOperandsInstruction }
type DLOAD_1 struct{ base.NoOperandsInstruction }
type DLOAD_2 struct{ base.NoOperandsInstruction }
type DLOAD_3 struct{ base.NoOperandsInstruction }

func dLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

func (ins *DLOAD) Execute(frame *rtda.Frame) {
	dLoad(frame, uint(ins.Index))
}

func (ins *DLOAD_0) Execute(frame *rtda.Frame) {
	dLoad(frame, 0)
}

func (ins *DLOAD_1) Execute(frame *rtda.Frame) {
	dLoad(frame, 1)
}

func (ins *DLOAD_2) Execute(frame *rtda.Frame) {
	dLoad(frame, 2)
}

func (ins *DLOAD_3) Execute(frame *rtda.Frame) {
	dLoad(frame, 3)
}
