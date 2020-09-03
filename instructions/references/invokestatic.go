package references

import (
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/rtda"
	"github.com/zjmeow/zjvm/rtda/heap"
)

type InvokeStatic struct {
	base.Index16Instruction
}

func (ins *InvokeStatic) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	methodRef := cp.GetConstant(ins.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()
	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	base.InvokeMethod(frame, method)
}
