package control

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
)

type Return struct {
	base.NoOperandsInstruction
	returnType string
}

// 返回弹出栈即可
func (r *Return) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	if r.returnType == "v" {
		return
	}
	invokerFrame := thread.TopFrame()
	switch r.returnType {
	case "i":
		returnVal := currentFrame.OperandStack().PopInt()
		invokerFrame.OperandStack().PushInt(returnVal)
	case "l":
		returnVal := currentFrame.OperandStack().PopLong()
		invokerFrame.OperandStack().PushLong(returnVal)
	case "f":
		returnVal := currentFrame.OperandStack().PopFloat()
		invokerFrame.OperandStack().PushFloat(returnVal)
	case "d":
		returnVal := currentFrame.OperandStack().PopDouble()
		invokerFrame.OperandStack().PushDouble(returnVal)
	case "a":
		returnVal := currentFrame.OperandStack().PopRef()
		invokerFrame.OperandStack().PushRef(returnVal)
		// void直接返回
	default:
		panic("illegal return type")
	}

}
func NewReturn(returnType string) *Return {
	return &Return{returnType: returnType}
}
