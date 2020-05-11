package stack

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type Dup struct {
	base.NoOperandsInstruction
}

func (ins *Dup) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	tmp := stack.PopSlot()
	stack.PushSlot(tmp)
	stack.PushSlot(tmp)
}

type DupX1 struct {
	base.NoOperandsInstruction
}

func (ins *DupX1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	tmp1 := stack.PopSlot()
	tmp2 := stack.PopSlot()
	stack.PushSlot(tmp1)
	stack.PushSlot(tmp2)
	stack.PushSlot(tmp1)
}

type DupX2 struct {
	base.NoOperandsInstruction
}

func (ins *DupX2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	tmp1 := stack.PopSlot()
	tmp2 := stack.PopSlot()
	tmp3 := stack.PopSlot()
	stack.PushSlot(tmp1)
	stack.PushSlot(tmp3)
	stack.PushSlot(tmp2)
	stack.PushSlot(tmp1)
}

type Dup2 struct {
	base.NoOperandsInstruction
}

func (ins *Dup2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	tmp1 := stack.PopSlot()
	tmp2 := stack.PopSlot()
	stack.PushSlot(tmp2)
	stack.PushSlot(tmp1)
	stack.PushSlot(tmp2)
	stack.PushSlot(tmp1)
}

type Dup2X1 struct {
	base.NoOperandsInstruction
}

func (ins *Dup2X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	tmp1 := stack.PopSlot()
	tmp2 := stack.PopSlot()
	tmp3 := stack.PopSlot()
	stack.PushSlot(tmp2)
	stack.PushSlot(tmp1)
	stack.PushSlot(tmp3)
	stack.PushSlot(tmp2)
	stack.PushSlot(tmp1)
}

type Dup2X2 struct {
	base.NoOperandsInstruction
}

func (ins *Dup2X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	tmp1 := stack.PopSlot()
	tmp2 := stack.PopSlot()
	tmp3 := stack.PopSlot()
	tmp4 := stack.PopSlot()
	stack.PushSlot(tmp2)
	stack.PushSlot(tmp1)
	stack.PushSlot(tmp4)
	stack.PushSlot(tmp3)
	stack.PushSlot(tmp2)
	stack.PushSlot(tmp1)
}
