package constants

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type ACONST_NULL struct{ base.NoOperandsInstruction }
type DCONST_0 struct{ base.NoOperandsInstruction }
type DCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }
type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }
type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }

func (ins *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (ins *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0)
}

func (ins *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1)
}

func (ins *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0)
}

func (ins *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1)
}

func (ins *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2)
}

func (ins *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}
func (ins *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}
func (ins *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

func (ins *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}
func (ins *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}
func (ins *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}
func (ins *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}
func (ins *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}
func (ins *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
