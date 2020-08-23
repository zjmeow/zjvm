package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type InvokeVirtual struct {
	base.Index16Instruction
}

func (ins *InvokeVirtual) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(ins.Index).(*heap.MethodRef)
	stack := frame.OperandStack()
	if methodRef.Name() == "println" {

	}
	stack.PopRef()
}
