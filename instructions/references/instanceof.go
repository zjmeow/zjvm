package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type InstanceOf struct {
	base.Index16Instruction
}

// 如果instance of 返回 true 则向 stack 里压入 true，否则压入false
func (ins *InstanceOf) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(ins.Index).(heap.ClassRef)
	class := classRef.ResolveClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	}
	stack.PushInt(0)
}
