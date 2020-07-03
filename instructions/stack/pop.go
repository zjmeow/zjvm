package stack

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type Pop struct {
	base.NoOperandsInstruction
}

type Pop2 struct {
	base.NoOperandsInstruction
}

func (ins *Pop) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

func (ins *Pop2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}
