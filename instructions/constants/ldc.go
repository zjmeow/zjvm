package constants

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type Ldc struct {
	base.Index8Instruction
}

type LdcW struct {
	base.Index16Instruction
}
type Ldc2W struct {
	base.Index16Instruction
}

func (ins *Ldc2W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	index := ins.Index
	c := cp.GetConstant(index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

func (ins *Ldc) Execute(frame *rtda.Frame) {
	ldc(frame, ins.Index)
}
func (ins *LdcW) Execute(frame *rtda.Frame) {
	ldc(frame, ins.Index)
}

func ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := heap.JString(frame.Method().Class().ClassLoader(), c.(string))
		stack.PushRef(internedStr)
	default:
		panic("todo")
	}
}
