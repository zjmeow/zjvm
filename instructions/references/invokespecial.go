package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type InvokeSpecial struct {
	base.Index16Instruction
}

func (ins *InvokeSpecial) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
