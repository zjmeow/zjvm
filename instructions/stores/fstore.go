package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type FSTORE struct{ base.Index8Instruction }
type FSTORE_0 struct{ base.NoOperandsInstruction }
type FSTORE_1 struct{ base.NoOperandsInstruction }
type FSTORE_2 struct{ base.NoOperandsInstruction }
type FSTORE_3 struct{ base.NoOperandsInstruction }

func fStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

func (ins *FSTORE) Execute(frame *rtda.Frame) {
	fStore(frame, ins.Index)
}

func (ins *FSTORE_0) Execute(frame *rtda.Frame) {
	fStore(frame, 0)
}

func (ins *FSTORE_1) Execute(frame *rtda.Frame) {
	fStore(frame, 1)
}

func (ins *FSTORE_2) Execute(frame *rtda.Frame) {
	fStore(frame, 2)
}

func (ins *FSTORE_3) Execute(frame *rtda.Frame) {
	fStore(frame, 3)
}
