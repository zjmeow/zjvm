package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type ALOAD struct{ base.Index8Instruction }
type ALOAD_0 struct{ base.NoOperandsInstruction }
type ALOAD_1 struct{ base.NoOperandsInstruction }
type ALOAD_2 struct{ base.NoOperandsInstruction }
type ALOAD_3 struct{ base.NoOperandsInstruction }

func aLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

func (ins *ALOAD) Execute(frame *rtda.Frame) {
	aLoad(frame, uint(ins.Index))
}

func (ins *ALOAD_0) Execute(frame *rtda.Frame) {
	aLoad(frame, 0)
}

func (ins *ALOAD_1) Execute(frame *rtda.Frame) {
	aLoad(frame, 1)
}

func (ins *ALOAD_2) Execute(frame *rtda.Frame) {
	aLoad(frame, 2)
}

func (ins *ALOAD_3) Execute(frame *rtda.Frame) {
	aLoad(frame, 3)
}
