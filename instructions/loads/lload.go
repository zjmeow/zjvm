package loads

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type LLOAD struct{ base.Index8Instruction }
type LLOAD_0 struct{ base.NoOperandsInstruction }
type LLOAD_1 struct{ base.NoOperandsInstruction }
type LLOAD_2 struct{ base.NoOperandsInstruction }
type LLOAD_3 struct{ base.NoOperandsInstruction }

func lLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

func (ins *LLOAD) Execute(frame *rtda.Frame) {
	lLoad(frame, uint(ins.Index))
}

func (ins *LLOAD_0) Execute(frame *rtda.Frame) {
	lLoad(frame, 0)
}

func (ins *LLOAD_1) Execute(frame *rtda.Frame) {
	lLoad(frame, 1)
}

func (ins *LLOAD_2) Execute(frame *rtda.Frame) {
	lLoad(frame, 2)
}

func (ins *LLOAD_3) Execute(frame *rtda.Frame) {
	lLoad(frame, 3)
}
