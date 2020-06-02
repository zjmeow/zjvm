package constants

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type ACONST_NULL struct{ base.NoOperandsInstruction }

type IntConst struct {
	base.NoOperandsInstruction
	num int32
}
type LongConst struct {
	base.NoOperandsInstruction
	num int64
}

type FloatConst struct {
	base.NoOperandsInstruction
	num float32
}

type DoubleConst struct {
	base.NoOperandsInstruction
	num float64
}

func NewConstInt(num int32) base.Instruction {
	return &IntConst{num: num}
}
func NewConstLong(num int64) base.Instruction {
	return &LongConst{num: num}
}
func NewConstFloat(num float32) base.Instruction {
	return &FloatConst{num: num}
}

func NewConstDouble(num float64) base.Instruction {
	return &DoubleConst{num: num}
}

func (ins *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (ins *IntConst) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(ins.num)
}
func (ins *LongConst) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(ins.num)
}
func (ins *FloatConst) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(ins.num)
}
func (ins *DoubleConst) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(ins.num)
}
