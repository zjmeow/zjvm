package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type ASTORE struct{ base.Index8Instruction }
type ASTORE_0 struct{ base.NoOperandsInstruction }
type ASTORE_1 struct{ base.NoOperandsInstruction }
type ASTORE_2 struct{ base.NoOperandsInstruction }
type ASTORE_3 struct{ base.NoOperandsInstruction }

func aStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (ins *ASTORE) Execute(frame *rtda.Frame) {
	aStore(frame, ins.Index)
}

func (ins *ASTORE_0) Execute(frame *rtda.Frame) {
	aStore(frame, 0)
}

func (ins *ASTORE_1) Execute(frame *rtda.Frame) {
	aStore(frame, 1)
}

func (ins *ASTORE_2) Execute(frame *rtda.Frame) {
	aStore(frame, 2)
}

func (ins *ASTORE_3) Execute(frame *rtda.Frame) {
	aStore(frame, 3)
}
