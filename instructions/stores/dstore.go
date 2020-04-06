package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type DSTORE struct{ base.Index8Instruction }
type DSTORE_0 struct{ base.NoOperandsInstruction }
type DSTORE_1 struct{ base.NoOperandsInstruction }
type DSTORE_2 struct{ base.NoOperandsInstruction }
type DSTORE_3 struct{ base.NoOperandsInstruction }

func dStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (ins *DSTORE) Execute(frame *rtda.Frame) {
	dStore(frame, ins.Index)
}

func (ins *DSTORE_0) Execute(frame *rtda.Frame) {
	dStore(frame, 0)
}

func (ins *DSTORE_1) Execute(frame *rtda.Frame) {
	dStore(frame, 1)
}

func (ins *DSTORE_2) Execute(frame *rtda.Frame) {
	dStore(frame, 2)
}

func (ins *DSTORE_3) Execute(frame *rtda.Frame) {
	dStore(frame, 3)
}
