package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type ISTORE struct{ base.Index8Instruction }
type ISTORE_0 struct{ base.NoOperandsInstruction }
type ISTORE_1 struct{ base.NoOperandsInstruction }
type ISTORE_2 struct{ base.NoOperandsInstruction }
type ISTORE_3 struct{ base.NoOperandsInstruction }

func iStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

func (ins *ISTORE) Execute(frame *rtda.Frame) {
	iStore(frame, ins.Index)
}

func (ins *ISTORE_0) Execute(frame *rtda.Frame) {
	iStore(frame, 0)
}

func (ins *ISTORE_1) Execute(frame *rtda.Frame) {
	iStore(frame, 1)
}

func (ins *ISTORE_2) Execute(frame *rtda.Frame) {
	iStore(frame, 2)
}

func (ins *ISTORE_3) Execute(frame *rtda.Frame) {
	iStore(frame, 3)
}
