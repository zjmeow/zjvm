package stack

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type Swap struct {
	base.NoOperandsInstruction
}

func (ins *Swap) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	tmp1 := stack.PopSlot()
	tmp2 := stack.PopSlot()
	stack.PushSlot(tmp1)
	stack.PushSlot(tmp2)
}
