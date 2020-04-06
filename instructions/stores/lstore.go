package stores

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type LSTORE struct{ base.Index8Instruction }
type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }

func lStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (ins *LSTORE) Execute(frame *rtda.Frame) {
	lStore(frame, ins.Index)
}

func (ins *LSTORE_0) Execute(frame *rtda.Frame) {
	lStore(frame, 0)
}

func (ins *LSTORE_1) Execute(frame *rtda.Frame) {
	lStore(frame, 1)
}

func (ins *LSTORE_2) Execute(frame *rtda.Frame) {
	lStore(frame, 2)
}

func (ins *LSTORE_3) Execute(frame *rtda.Frame) {
	lStore(frame, 3)
}
