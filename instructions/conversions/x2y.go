package conversions

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type X2Y struct {
	base.NoOperandsInstruction
	castFn func(frame *rtda.Frame)
}

func (instr *X2Y) Execute(frame *rtda.Frame) {
	instr.castFn(frame)
}

func i2b(frame *rtda.Frame) { frame.OperandStack().PushInt(int32(int8(frame.OperandStack().PopInt()))) }
func i2c(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(uint16(frame.OperandStack().PopInt())))
}
func i2s(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(int16(frame.OperandStack().PopInt())))
}
func i2l(frame *rtda.Frame) { frame.OperandStack().PushLong(int64(frame.OperandStack().PopInt())) }
func i2f(frame *rtda.Frame) { frame.OperandStack().PushFloat(float32(frame.OperandStack().PopInt())) }
func i2d(frame *rtda.Frame) { frame.OperandStack().PushDouble(float64(frame.OperandStack().PopInt())) }
func l2i(frame *rtda.Frame) { frame.OperandStack().PushInt(int32(frame.OperandStack().PopLong())) }
func l2f(frame *rtda.Frame) { frame.OperandStack().PushFloat(float32(frame.OperandStack().PopLong())) }
func l2d(frame *rtda.Frame) { frame.OperandStack().PushDouble(float64(frame.OperandStack().PopLong())) }
func f2i(frame *rtda.Frame) { frame.OperandStack().PushInt(int32(frame.OperandStack().PopFloat())) }
func f2l(frame *rtda.Frame) { frame.OperandStack().PushLong(int64(frame.OperandStack().PopFloat())) }
func f2d(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(float64(frame.OperandStack().PopFloat()))
}
func d2i(frame *rtda.Frame) { frame.OperandStack().PushInt(int32(frame.OperandStack().PopDouble())) }
func d2l(frame *rtda.Frame) { frame.OperandStack().PushLong(int64(frame.OperandStack().PopDouble())) }
func d2f(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(float32(frame.OperandStack().PopDouble()))
}
